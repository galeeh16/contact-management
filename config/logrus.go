package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

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

	return logger
}
