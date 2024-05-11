package initializers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CheckCommand() {
    cmd := exec.Command("git", "filter-repo", "--version")
    err := cmd.Run()
    if err != nil {
        fmt.Println("git filter-repo is not installed")
        switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("Please run the following command to install git filter-repo on MacOS: \n\n    brew install git-filter-repo")
        case "linux":
            fmt.Println("Please run the following command to install git filter-repo on Linux: \n\n    sudo apt-get install git-filter-repo")
        case "windows":
            fmt.Println("Please run the following command to install git filter-repo on Windows: \n\n  python3 -m pip install --user git-filter-repo")
        default:
            fmt.Println("Please run the following command to install git filter-repo: \n\n  python3 -m pip install --user git-filter-repo")
        }
        os.Exit(1)
    } 
}