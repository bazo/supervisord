package slack

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

type SlackFormatter struct {
}

func NewSlackFormatter() *SlackFormatter {
	return &SlackFormatter{}
}

//https://github.com/johntdyer/slackrus/blob/master/slackrus.go
func (f *SlackFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Note this doesn't include Time, Level and Message which are available on
	// the Entry. Consult `godoc` on information about those fields or read the
	// source of the official loggers.
	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func (f *SlackFormatter) Fire(entry *logrus.Entry) error {
	log.Println(entry)
	return nil
}

func (f *SlackFormatter) Levels() []logrus.Level {
	return logrus.AllLevels
}
