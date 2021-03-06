package master

import (
	"encoding/json"
	"fmt"
)

func Master() {
	fmt.Println("this is Master Branch")
}

const ChangeVersion = "2"

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
