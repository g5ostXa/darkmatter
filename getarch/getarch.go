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

const (
	mirror           = "https://mirror.quantum5.ca/archlinux/iso/latest/"
	bufferSize       = 64 * 1024
	progressBarWidth = 50
)

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

	m := newModel(cancel)
	p := tea.NewProgram(m)

	go downloadFiles(ctx, p, downloads)

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func newModel(cancel context.CancelFunc) model {

	m := model{
		bars:       make([]progress.Model, len(files)),
		downloaded: make([]int64, len(files)),
		totalBytes: make([]int64, len(files)),
		cancel:     cancel,
	}
	for i := range m.bars {
		m.bars[i] = progress.New(
			progress.WithDefaultBlend(),
			progress.WithWidth(progressBarWidth),
		)
	}
	return m
}

func downloadFiles(ctx context.Context, p *tea.Program, downloadDir string) {

	for i, name := range files {
		if err := download(ctx, p, i, mirror+name, filepath.Join(downloadDir, name)); err != nil {
			p.Send(doneMsg{err: err})
			return
		}
	}
	p.Send(doneMsg{})
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

func download(ctx context.Context, p *tea.Program, index int, url, dest string) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	return writeDownload(ctx, p, index, resp.Body, out, resp.ContentLength)
}

func writeDownload(ctx context.Context, p *tea.Program, index int, src io.Reader, dst io.Writer, total int64) error {

	buf := make([]byte, bufferSize)
	downloaded := int64(0)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		n, err := src.Read(buf)
		if n > 0 {
			if _, writeErr := dst.Write(buf[:n]); writeErr != nil {
				return writeErr
			}
			downloaded += int64(n)
			p.Send(progressMsg{index: index, done: downloaded, total: total})
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
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
			return m, m.bars[msg.index].SetPercent(float64(msg.done) / float64(msg.total))
		}
	case progress.FrameMsg:
		cmds := make([]tea.Cmd, len(m.bars))
		for i := range m.bars {
			m.bars[i], cmds[i] = m.bars[i].Update(msg)
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
		fmt.Fprintf(&b, "  %s\n  %s", name, m.bars[i].View())
		if m.totalBytes[i] > 0 {
			fmt.Fprintf(&b, "  %s / %s", formatBytes(m.downloaded[i]), formatBytes(m.totalBytes[i]))
		}
		b.WriteString("\n\n")
	}

	if m.cancelled {
		b.WriteString("  Cancelled.\n")
	} else if m.err != nil && m.err != context.Canceled {
		fmt.Fprintf(&b, "  Error: %s\n", m.err.Error())
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

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
