package tools

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"strings"
)

func Retimegit(path string, duration string, flags map[string]bool) {

	fmt.Printf("Dormammu, I've come to bargain.\n")

	commitNum := getTotalCommits(path, flags)
	durationOffset, err := time.ParseDuration("-" + duration)
	if err != nil {
		log.Fatal(err)
	}
	nowTime := time.Now().Add(durationOffset).Unix()

	durationInt, err := strconv.Atoi(strings.TrimSuffix(duration, "h"))
	if err != nil {
		log.Fatal(err)
	}

	if(flags["verbose"]){
		fmt.Printf("Start time in Unix seconds is %d\n",durationInt)
	}

	intervalHop := durationInt * 3600 / commitNum

	if(flags["verbose"]){
		fmt.Printf("Interval hop is %d\n",intervalHop)
	}

	retimeCmd := []string{
		"--commit-callback",
		fmt.Sprintf(`
		curr_time = ["%d", "+0530"]
		curr_time[0] = str(int(curr_time[0]) + (%d * commit.id))
		changed_time = " ".join(curr_time).encode()
		commit.author_date = changed_time
		commit.committer_date = changed_time
		`, nowTime, intervalHop),
		"--force",
	}
	ExecuteRewrite(path, retimeCmd, flags)
}
