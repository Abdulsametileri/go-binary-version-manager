package commandrunner

import (
	"fmt"
	"os/exec"
	"strings"
)

var _ CommandRunner = (*curlCommandRunner)(nil)

type curlCommandRunner struct{}

func (c *curlCommandRunner) RunWith(args ...string) (string, error) {
	args = append([]string{"-c"}, args...)
	cmd := exec.Command("sh", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to execute command: %s\nOutput: %s", cmd.String(), out)
	}
	return strings.TrimSpace(strings.TrimRight(string(out), "\r\n")), nil
}
