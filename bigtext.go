package bigtext

import (
	"fmt"
	"os/exec"
)

// Use the Quicksilver app to display the text in large type.
func Display(text string) error {
	args := []string{
		"-e",
		fmt.Sprintf("tell application \"Quicksilver\" to show large type \"%s\"", text),
	}
	_, err := exec.Command("osascript", args...).Output()
	return err
}
