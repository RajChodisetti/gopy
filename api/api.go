package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func Api(folder string) {

	fileChan := make(chan string)

	// Use WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start a goroutine to send relative file paths to the channel
	wg.Add(1)
	go func() {
		defer close(fileChan)
		defer wg.Done()
		err := listFilesInFolder(folder, fileChan)
		if err != nil {
			log.Println("Error:", err)
		}
	}()

	// Start a goroutine to send relative file paths to the Python program
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendFilesToPython(folder, fileChan)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

func listFilesInFolder(folder string, fileChan chan<- string) error {
	log.Println("Starting directory traversal:", folder)
	defer log.Println("Finished directory traversal:", folder)

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				log.Printf("Permission denied to traverse %s: %s\n", path, err)
				return nil
			}
			log.Printf("Error encountered while traversing %s: %s\n", path, err)
			return err
		}
		if !info.IsDir() {
			fileChan <- path
		}
		return nil
	})

	return err
}

func sendFilesToPython(folder string, fileChan <-chan string) {
	var relativePaths []string

	for path := range fileChan {
		relativePath, err := filepath.Rel(folder, path)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		relativePaths = append(relativePaths, relativePath)
	}

	jsonData, err := json.Marshal(relativePaths)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	endpoint := "http://localhost:8000/receive_filenames"
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return
	}

}
func Pwd(r *http.Request) string {
	directory := r.URL.Query().Get("directory")
	if directory == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return pwd
		}
	}
	return directory
}
