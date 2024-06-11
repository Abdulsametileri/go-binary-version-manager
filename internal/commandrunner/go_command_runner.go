package commandrunner

import (
	"fmt"
	"os/exec"
	"strings"
)

var _ CommandRunner = (*goCommandRunner)(nil)

type goCommandRunner struct{}

func (c *goCommandRunner) RunWith(args ...string) (string, error) {
	cmd := exec.Command("go", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %s\noutput: %s", cmd.String(), out)
	}
	return strings.TrimSpace(strings.TrimRight(string(out), "\r\n")), nil
}
