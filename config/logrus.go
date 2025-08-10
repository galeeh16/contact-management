package config

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Telegram Hooks
type TelegramHook struct {
	Token  string
	ChatID string
}

// level2 yang dikirim hook
func (hook *TelegramHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// fire dieksekusi saat log dengan levels di method Levels dipanggil
func (hook *TelegramHook) Fire(entry *logrus.Entry) error {
	// format pesan notifnya
	msg := fmt.Sprintf("App: %s\nLevel: %s\nMessage: %s\nTime: %s\n",
		os.Getenv("APP_NAME"),
		entry.Level,
		entry.Message,
		entry.Time.Format("2006-01-02 15:04:05"),
	)

	var msgBuilder strings.Builder
	msgBuilder.WriteString(msg)

	// tambahkan semua fields dari entry.Data
	if len(entry.Data) > 0 {
		msgBuilder.WriteString("*Fields*:\n")
		for key, value := range entry.Data {
			msgBuilder.WriteString(fmt.Sprintf(" - %s: %v\n", key, value))
		}
	}

	// URL API Telegram untuk send messsage
	apiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", hook.Token)

	// data form untuk POST request
	data := url.Values{}
	data.Set("chat_id", hook.ChatID)
	data.Set("text", msgBuilder.String())
	data.Set("parse_mode", "Markdown") // Optional: untuk format teks

	// kirim http post
	resp, err := http.PostForm(apiUrl, data)

	if err != nil {
		return fmt.Errorf("gagal kirim notifikasi ke telegram: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("gagal membaca response body telegram: %s", err.Error())
		}
		return fmt.Errorf("respons telegram tidak OK: status: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

// Initialize logger
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Membuat direktori logs jika belum ada
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		logger.Fatalf("Gagal membuat direktori logs: %v", err)
	}

	// Membuat nama file log berdasarkan tanggal saat ini
	logFileName := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")

	// Mengatur output logger ke lumberjack buat rotating
	logger.SetOutput(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    50,   // Max size in MB
		MaxBackups: 7,    // Max number of old log files to keep
		MaxAge:     7,    // Max age in days to keep a log file
		Compress:   true, // Compress old log files
	})

	// Mengatur level log
	logger.SetLevel(logrus.InfoLevel)

	logHook := os.Getenv("LOG_HOOK")

	// add hook telegram
	if logHook == "telegram" {
		hook := &TelegramHook{
			Token:  os.Getenv("TELEGRAM_TOKEN"),
			ChatID: os.Getenv("TELEGRAM_CHAT_ID"),
		}

		logger.AddHook(hook)
	}

	return logger
}
