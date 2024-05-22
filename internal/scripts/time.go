package scripts

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/woaitsAryan/regit/internal/helpers"
	"github.com/woaitsAryan/regit/internal/models"
)

func Rewind(duration string, flags models.Flags) {

	fmt.Printf("Dormammu, I've come to bargain.\n")

	durationOffset, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatal(err)
	}

	unixOffset := int64(durationOffset.Seconds())

	if flags.Verbose {
		fmt.Printf("Rewinding by %d\n", durationOffset)
	}
	fmt.Println(unixOffset)

	rewindCmd := []string{
		"--commit-callback",
		fmt.Sprintf(`
		time, timezone = commit.author_date.decode().split(" ")
		fastforward = %d
		new_time = [str(int(time) - fastforward), timezone]
		changed_time = " ".join(new_time).encode()
		commit.committer_date = changed_time
		commit.author_date = changed_time
		`, unixOffset),
		"--force",
	}
	helpers.ExecuteRewrite(rewindCmd, flags)
}


func Retimegit(duration string, flags models.Flags) {

	fmt.Printf("Dormammu, I've come to bargain.\n")

	commitNum := helpers.GetTotalCommits(flags)
	durationOffset, err := time.ParseDuration("-" + duration)
	if err != nil {
		log.Fatal(err)
	}
	nowTime := time.Now().Add(durationOffset).Unix()

	durationInt, err := strconv.Atoi(strings.TrimSuffix(duration, "h"))
	if err != nil {
		log.Fatal(err)
	}

	if(flags.Verbose){
		fmt.Printf("Start time in Unix seconds is %d\n",durationInt)
	}

	intervalHop := durationInt * 3600 / commitNum

	if(flags.Verbose){
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
	helpers.ExecuteRewrite(retimeCmd, flags)
}

func FastForward(duration string, flags models.Flags) {

	fmt.Printf("Dormammu, I've come to bargain.\n")

	durationOffset, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatal(err)
	}

	unixOffset := int64(durationOffset.Seconds())

	if flags.Verbose {
		fmt.Printf("Fast forwarding by %d\n", durationOffset)
	}
	fmt.Println(unixOffset)

	fastForwardCmd := []string{
		"--commit-callback",
		fmt.Sprintf(`
		time, timezone = commit.author_date.decode().split(" ")
		fastforward = %d
		new_time = [str(int(time) + fastforward), timezone]
		changed_time = " ".join(new_time).encode()
		commit.committer_date = changed_time
		commit.author_date = changed_time
		`, unixOffset),
		"--force",
	}
	helpers.ExecuteRewrite(fastForwardCmd, flags)
}
