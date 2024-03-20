package tools

import (
	"bytes"
	"fmt"
	"log"

	// "log"
	// "os"

	"os/exec"
	// "strconv"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/woaitsAryan/regit/initializers"
)

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

var respBody Response
var commitResponsedetails map[string]string

func Recommitgit(source string, flags map[string]bool) {
	initializers.LoadEnv()
	commitDetails := getCommitData(source, flags)
	newCommitDetails := sendOpenAIMessage(commitDetails, flags)
	finalCommitDetails := orderCommitData(source, newCommitDetails)
	fmt.Println(finalCommitDetails)
}

func orderCommitData(source string, commitDetails map[string]string) []string {
	finalCommitDetails := []string{}

	out, _ := exec.Command("git", "-C", source, "log", "--pretty=format:%H").Output()
	commits := strings.Split(string(out), "\n")

	fmt.Println(commits)
	fmt.Println(commitDetails)

	for _, commit := range commits {

		commitAppend := commitDetails[commit]

		if(commitAppend == ""){
			fmt.Println(commit)
			log.Fatalln("ChatGPT messed up!")
		}

		finalCommitDetails = append(finalCommitDetails, commitAppend)
	}

	return finalCommitDetails
}

func getCommitData(source string, flags map[string]bool) map[string]string {
	commitDetails := make(map[string]string)

	out, _ := exec.Command("git", "-C", source, "log", "--pretty=format:%H").Output()
	commits := strings.Split(string(out), "\n")

	for _, commit := range commits {
		// Get the commit details
		out, _ := exec.Command("git", "-C", source, "show", "--no-color", commit).Output()
		// Store the commit details in the map
		commitDetails[commit] = string(out)
	}

	if flags["verbose"] {
		fmt.Println("Captured commit details!")
	}

	return commitDetails
}

func sendOpenAIMessage(commitDetails map[string]string, flags map[string]bool) map[string]string {

	jsonStr, _ := json.Marshal(commitDetails)

	fmt.Println("Processing commit details, this might take some time...")


	body := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"response_format": map[string]string{
			"type": "json_object",
		},
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "You are given commit message details in a format of key:commit hash and value: commit details, you should ONLY return a object with a key: commit hash and value: new commit message and modify it what the commit message should be with proper formatting like using feat, fix and so on. Output in json",
			},
			{
				"role":    "user",
				"content": string(jsonStr),
			},
		},
	}

	bodyBytes, _ := json.Marshal(body)
	OPENAI_API_KEY := initializers.CONFIG.OPENAI_API_KEY

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)

	client := &http.Client{}
	resp, _ := client.Do(req)

	if(resp == nil){
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

	err = json.Unmarshal([]byte(content), &commitResponsedetails)
	if err != nil {
		log.Fatalln("Error unmarshaling JSON:", err)
	}

	return commitResponsedetails
}
