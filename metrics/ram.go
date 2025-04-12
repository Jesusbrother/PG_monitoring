package metrics

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetContainerRAMUsage() (int64, int64, error) {
	// Used memory
	usedCmd := exec.Command("docker", "exec", "dockerpg", "cat", "/sys/fs/cgroup/memory/memory.usage_in_bytes")
	usedOutput, err := usedCmd.Output()
	if err != nil {
		return 0, 0, err
	}

	usedMemory, err := strconv.ParseInt(strings.TrimSpace(string(usedOutput)), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	// Total memory limit
	limitCmd := exec.Command("docker", "exec", "dockerpg", "cat", "/sys/fs/cgroup/memory/memory.limit_in_bytes")
	limitOutput, err := limitCmd.Output()
	if err != nil {
		return 0, 0, err
	}

	totalMemory, err := strconv.ParseInt(strings.TrimSpace(string(limitOutput)), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return usedMemory, totalMemory, nil
}
