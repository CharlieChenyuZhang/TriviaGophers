package trivia

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAQuestion(category string) string {
	response, err := http.Get("https://www.randomtriviagenerator.com/questions?limit=1&page=1&category=" + category)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Successfully retrieved a trivia question %s\n", string(data))
	return string(data)
}
