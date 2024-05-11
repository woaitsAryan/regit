package scripts

import (
	"bytes"
	"fmt"
	"log"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/woaitsAryan/regit/internal/initializers"
	"github.com/woaitsAryan/regit/internal/helpers"
	"github.com/woaitsAryan/regit/internal/models"
)

func Recommitgit(flags models.Flags) {
	initializers.LoadEnv()
	commitDetails := getCommitData(flags)
	newCommitDetails := sendOpenAIMessage(commitDetails, flags)
	arrLength := len(newCommitDetails)
	jsonCommitDetails, _ := json.Marshal(newCommitDetails)

	recommitCmd := []string{
		"--commit-callback",
		fmt.Sprintf(`
		new_commit_arr = %s
		arr_length = %d
		commit.message = new_commit_arr[arr_length - commit.old_id].encode()
		`, string(jsonCommitDetails), arrLength),
		"--force",
	}

	helpers.ExecuteRewrite(recommitCmd, flags)
}

func getCommitData(flags models.Flags) []string {
	var commitDetails []string

	out, _ := exec.Command("git", "-C", flags.Source, "log", "--pretty=format:%H").Output()
	commits := strings.Split(string(out), "\n")

	for _, commit := range commits {
		out, _ := exec.Command("git", "-C", flags.Source, "show", "--no-color", commit).Output()
		commitDetails = append(commitDetails, string(out))
	}

	if flags.Verbose {
		fmt.Println("Captured commit details!")
	}

	return commitDetails
}

func sendOpenAIMessage(commitDetails []string, flags models.Flags) []string {

	fmt.Println("Processing commit details, this might take some time...")

	var respBody models.Response
	var commitResponseDetails []string

	for i, commit  := range commitDetails {

		body := map[string]interface{}{
			"model": initializers.OPENAI_MODEL,
			"messages": []map[string]interface{}{
				{
					"role":    "system",
					"content": "You are given commit message details and diffs. modify it what the commit message should be with proper formatting like using feat, fix or chore, you MUST always use these at the start. Output just the commit message and nothing else. If there's not enough information then just try to guess, never ask for more information.",
				}, 
				{
					"role":    "user",
					"content": string(commit),
				},
			},
		}

		bodyBytes, _ := json.Marshal(body)
		OPENAI_API_KEY := initializers.CONFIG.OPENAI_API_KEY

		req, _ := http.NewRequest("POST", initializers.OPENAI_URL, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)

		client := &http.Client{}
		resp, _ := client.Do(req)

		if resp == nil {
			log.Fatalln("Response is nil, fatal error, probably internet connectivity issue.")
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		bodyString := string(bodyBytes)

		err = json.Unmarshal([]byte(bodyString), &respBody)
		if err != nil {
			log.Fatalln(err)
		}
		content := respBody.Choices[0].Message.Content
		if flags.Verbose {
			fmt.Printf("Commit %d processed! Generated commit message is %s", i, content )
		}
		commitResponseDetails = append(commitResponseDetails, content)
	}
	return commitResponseDetails
}
