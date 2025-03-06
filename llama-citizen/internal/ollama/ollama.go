package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/pkg/structure"
)

type Response struct {
	Model    string `json:"model"`
	Response string `json:"response"`
}

func GetResponse(msg structure.Message) (*Response, error) {
	payload := map[string]interface{}{
		"model":    msg.Model,
		"messages": msg.Messages,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	resp, err := http.Post(
		"http://localhost:11434/v1/chat/completions",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Ollama Server: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	//TODO: delete bodyString
	bodyString := string(bodyBytes)
	fmt.Println("Response from Ollama Server:", bodyString)

	var ollamaResp Response
	if err := json.Unmarshal(bodyBytes, &ollamaResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &ollamaResp, nil
}
