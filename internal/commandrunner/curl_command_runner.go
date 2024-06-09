package commandrunner

import (
	"fmt"
	"os/exec"
)

var _ CommandRunner = (*curlCommandRunner)(nil)

type curlCommandRunner struct{}

func (c *curlCommandRunner) Run(curlCmd string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", curlCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte(""), fmt.Errorf("Failed to execute command: %s\nOutput: %s", cmd.String(), out)
	}
	return out, nil
}
