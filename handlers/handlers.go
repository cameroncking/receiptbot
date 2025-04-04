package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"path/filepath"
	"receiptbot/utils"
	"strings"
	"time"
)

var userFolders = make(map[int64]string)
var userLastImageTimes = make(map[int64]time.Time)
var userHaveGotMemos = make(map[int64]bool)

func HandlePhoto(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.From.ID
	currentTime := time.Now()
	// Check if a new folder should be created
	if userFolders[userID] == "" || userHaveGotMemos[userID] || currentTime.Sub(userLastImageTimes[userID]) > 15*time.Minute {
		userFolders[userID] = utils.CreateFolder(userID, currentTime)
		userHaveGotMemos[userID] = false
	}
	var fileID string
	var fileName string
	if message.Photo != nil {
		fileID = message.Photo[len(message.Photo)-1].FileID
		fileName = fmt.Sprintf("photo_%d.jpg", time.Now().Unix())
		log.Printf("HandlePhoto: found photo %s, %s", fileID, fileName)
	} else if message.Document != nil && isImage(message.Document.MimeType) {
		fileID = message.Document.FileID
		fileName = message.Document.FileName
		log.Printf("HandlePhoto: found file %s, %s", fileID, fileName)
	} else {
		return
	}
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		fmt.Println("Error getting file:", err)
		return
	}
	filePath := filepath.Join(userFolders[userID], fileName)
	utils.DownloadFile(filePath, file.Link(bot.Token))
	userLastImageTimes[userID] = currentTime
	// Process the caption if it exists
	if message.Caption != "" {
		ProcessMemo(message.From.ID, message.Caption)
	}
}
func HandleText(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.From.ID
	ProcessMemo(userID, message.Text)
}
func ProcessMemo(userID int64, msg string) {
	if userFolders[userID] == "" {
		return
	}
	userHaveGotMemos[userID] = true
	memoPath := filepath.Join(userFolders[userID], "memo.txt")
	utils.AppendToFile(memoPath, msg)
}
func isImage(mimeType string) bool {
	return strings.HasPrefix(mimeType, "image/") || mimeType == "application/pdf"
}
