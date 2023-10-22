package main

import (
	"consumer/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func decodeData(payload []byte) (models.QueuePayload, error) {

	productWrapper := models.QueuePayload{}

	err := json.Unmarshal(payload, &productWrapper)
	if err != nil {
		return models.QueuePayload{}, errors.New("error unmarshalling message:" + err.Error())
	}

	return productWrapper, nil
}

func downloadImage(url, destinationFolder string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download image: HTTP status code %d", response.StatusCode)
	}

	// Extract the file name from the URL
	fileName := filepath.Base(url)
	fmt.Println(fileName)

	// Create the destination file
	filePath := filepath.Join(destinationFolder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the image content to the local file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
