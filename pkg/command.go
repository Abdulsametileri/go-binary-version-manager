package pkg

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %s\noutput: %s", cmd.String(), out)
	}

	return strings.TrimSpace(strings.TrimRight(string(out), "\r\n")), nil
}
