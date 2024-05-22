package initializers

import (
    "os/exec"
    "strings"
)

func GetLatestTag() string {
    cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
    output, err := cmd.Output()
    if err != nil {
        return "v0.0.0"
    }
    return strings.TrimSpace(string(output))
}