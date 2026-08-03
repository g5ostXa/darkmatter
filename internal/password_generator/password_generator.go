package password_generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"charm.land/lipgloss/v2"
)

var (
	simpleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff004c"))
)

const defaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func generatePassword(length int, chars string) ([]byte, error) {

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return nil, err
		}
		password[i] = chars[num.Int64()]
	}
	return password, nil
}

func Gen() {

	passwordLength := 16
	password, err := generatePassword(passwordLength, defaultChars)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to generate password:", err)
		return
	}

	defer func() {

		for i := range password {
			password[i] = 0
		}
	}()

	lipgloss.Println(simpleStyle.Render(string(password)))
}
