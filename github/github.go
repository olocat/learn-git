package github

import (
	"encoding/json"
	"fmt"
)

func GithubBranch() {
	fmt.Println("this is second branch")
}

const ChangeVersion = "1"

func Content() {
	content := map[string]string{
		"1": "the first change",
	}

	jsonBytes, err := json.Marshal(content)
	jsonString := string(jsonBytes)
	if err != nil {
		return
	}

	fmt.Println("this is modify content")
	fmt.Println("the change version is", ChangeVersion)
	fmt.Println("the content is", jsonString)
}
