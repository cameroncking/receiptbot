package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func CreateFolder(userID int64, currentTime time.Time) string {
	folderName := fmt.Sprintf("./images/%d-%s", userID, currentTime.Format("20060102-150405"))
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return ""
	}
	return folderName
}
func DownloadFile(filepath string, url string) error {
	log.Printf("DownloadFile: %s, %s\n", filepath, url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
func AppendToFile(filepath string, text string) error {
	log.Printf("AppendToFile: %s, %s\n", filepath, text)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text + "\n")
	return err
}
