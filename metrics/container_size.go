package metrics

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetContainerSize() (int64, error) {
	cmd := exec.Command("docker", "exec", "dockerpg", "du", "-sb", "/var/lib/postgresql/data")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	parts := strings.Fields(string(output))
	if len(parts) < 1 {
		return 0, err
	}

	sizeBytes, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return sizeBytes, nil
}
