package getarch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
)

const mirror = "https://mirror.quantum5.ca/archlinux/iso/latest/"

var files = []string{
	"archlinux-x86_64.iso.sig",
	"archlinux-x86_64.iso",
}

type progressMsg struct {
	index int
	done  int64
	total int64
}

type doneMsg struct {
	err error
}

type model struct {
	bars       []progress.Model
	downloaded []int64
	totalBytes []int64
	err        error
	cancel     context.CancelFunc
	cancelled  bool
}

func Latest() {

	downloads, err := downloadsDir()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := model{
		bars:       make([]progress.Model, len(files)),
		downloaded: make([]int64, len(files)),
		totalBytes: make([]int64, len(files)),
		cancel:     cancel,
	}

	for i := range files {
		m.bars[i] = progress.New(
			progress.WithDefaultBlend(),
			progress.WithWidth(50),
		)
	}

	p := tea.NewProgram(m)

	go func() {

		for i, name := range files {
			err := download(
				ctx,
				p,
				i,
				mirror+name,
				filepath.Join(downloads, name),
			)

			if err != nil {
				p.Send(doneMsg{err: err})
				return
			}
		}

		p.Send(doneMsg{})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func downloadsDir() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, "Downloads")

	info, err := os.Stat(dir)
	if err != nil {
		return "", fmt.Errorf("%s does not exist", dir)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("%s exists but is not a directory", dir)
	}

	return dir, nil
}

func download(

	ctx context.Context,
	p *tea.Program,
	index int,
	url string,
	dest string,
) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", url, resp.Status)
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	total := resp.ContentLength
	var downloaded int64

	buf := make([]byte, 64*1024)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		n, readErr := resp.Body.Read(buf)

		if n > 0 {
			if _, err := out.Write(buf[:n]); err != nil {
				return err
			}

			downloaded += int64(n)

			p.Send(progressMsg{
				index: index,
				done:  downloaded,
				total: total,
			})
		}

		if readErr == io.EOF {
			break
		}

		if readErr != nil {
			return readErr
		}
	}

	return nil
}

func (m model) Init() tea.Cmd {

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.cancelled = true
			m.cancel()
			return m, tea.Quit
		}

	case progressMsg:
		m.downloaded[msg.index] = msg.done
		m.totalBytes[msg.index] = msg.total

		if msg.total > 0 {
			percent := float64(msg.done) / float64(msg.total)
			cmd := m.bars[msg.index].SetPercent(percent)
			return m, cmd
		}

	case progress.FrameMsg:
		var cmds []tea.Cmd

		for i := range m.bars {
			var cmd tea.Cmd
			m.bars[i], cmd = m.bars[i].Update(msg)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)

	case doneMsg:
		m.err = msg.err
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() tea.View {

	var b strings.Builder

	b.WriteString("\n  Downloading latest Arch ISO\n\n")

	for i, name := range files {
		b.WriteString("  ")
		b.WriteString(name)
		b.WriteString("\n")

		b.WriteString("  ")
		b.WriteString(m.bars[i].View())

		if m.totalBytes[i] > 0 {
			fmt.Fprintf(
				&b,
				"  %s / %s",
				formatBytes(m.downloaded[i]),
				formatBytes(m.totalBytes[i]),
			)
		}

		b.WriteString("\n\n")
	}

	if m.cancelled {
		b.WriteString("  Cancelled.\n")
	} else if m.err != nil && m.err != context.Canceled {
		b.WriteString("  Error: ")
		b.WriteString(m.err.Error())
		b.WriteString("\n")
	} else {
		b.WriteString("  Saving to ~/Downloads\n")
	}

	return tea.NewView(b.String())
}

func formatBytes(bytes int64) string {

	const unit = 1024

	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div := int64(unit)
	exp := 0

	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.1f %ciB",
		float64(bytes)/float64(div),
		"KMGTPE"[exp],
	)
}
