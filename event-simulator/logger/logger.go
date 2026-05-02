package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Time    string         `json:"time"`
	Level   string         `json:"level"`
	Message string         `json:"message"`
	Fields  map[string]any `json:"fields,omitempty"`
}

type Logger struct {
	out *log.Logger
}

func New() *Logger {
	return &Logger{
		out: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) log(level, msg string, fields map[string]any) error {
	entry := LogEntry{
		Time:    time.Now().UTC().Format(time.RFC3339),
		Level:   level,
		Message: msg,
		Fields:  fields,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	l.out.Println(string(data))
	return nil
}

func (l *Logger) Info(msg string, fields map[string]any) {
	l.log("INFO", msg, fields)
}

func (l *Logger) Error(msg string, fields map[string]any) {
	l.log("ERROR", msg, fields)
}
