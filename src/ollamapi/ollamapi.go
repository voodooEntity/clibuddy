package ollamapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// OllamApi struct to hold API URL and model name
type OllamApi struct {
	ApiURL string
}

// New function to create a new instance of OllamApi
func New(apiURL string) *OllamApi {
	return &OllamApi{
		ApiURL: apiURL,
	}
}

// RequestData struct for the POST request payload
type RequestData struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// Ask method to send the request to the API and get a response
func (o *OllamApi) Ask(model string, message string) (string, error) {
	// Create the data for the POST request
	data := RequestData{
		Model:  model,
		Prompt: message,
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error encoding JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", o.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Parse the result
	return parseResult(string(body)), nil
}

// parseResult function to parse the JSON response
func parseResult(result string) string {
	lines := strings.Split(result, "\n")
	var r string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			var response map[string]interface{}
			if err := json.Unmarshal([]byte(line), &response); err == nil {
				r += response["response"].(string)
			}
		}
	}
	return r
}
