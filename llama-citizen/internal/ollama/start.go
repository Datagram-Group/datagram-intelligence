package ollama

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func StartOllamaServer(ollamaDir string) error {
	err := os.Chdir(ollamaDir)
	if err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	cmd := exec.Command("./ollama", "serve")
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start Ollama server: %v", err)
	}

	log.Println("Ollama Server started.")
	time.Sleep(5 * time.Second)
	return nil
}
