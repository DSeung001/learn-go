package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APIKey      = "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"
	ModelID     = "gpt-3.5-turbo"
	Prompt      = "Tell me a joke about cats:"
	MaxTokens   = 50
	Temperature = 0.7
)

type RequestBody struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type ResponseBody struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {

	// request 설정
	url := fmt.Sprintf("https://api.openai.com/v1/engines/%s/completions", ModelID)
	requestBody := RequestBody{
		Prompt:      Prompt,
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
	}

	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+APIKey)

	// request 요청
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	var responseBody ResponseBody

	// response 받음
	json.Unmarshal(respBody, &responseBody)
	fmt.Println(responseBody.Choices)

	if len(responseBody.Choices) > 0 {
		joke := responseBody.Choices[0].Text
		fmt.Println("Generated Joke:", joke)
	} else {
		fmt.Println("No joke generated.")
	}
}
