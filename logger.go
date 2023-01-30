package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "%time% [%lvl%]: %msg%\n"
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

// New returns Logrus with default formatter
func New() *logrus.Logger {
	return &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &Formatter{
			UseColor: true,
		},
	}
}

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	TimestampFormat string
	LogFormat       string
	UseColor        bool
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	output = strings.Replace(output, "%lvl%", convertLevelToText(entry.Level, f.UseColor), 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}

func convertLevelToText(level logrus.Level, useColor bool) string {
	switch level {
	case logrus.TraceLevel:
		return color("TRACE", useColor)
	case logrus.DebugLevel:
		return color("DEBUG", useColor)
	case logrus.InfoLevel:
		return color("INFO ", useColor)
	case logrus.WarnLevel:
		return color("WARN ", useColor)
	case logrus.ErrorLevel:
		return color("ERROR", useColor)
	case logrus.FatalLevel:
		return color("FATAL", useColor)
	case logrus.PanicLevel:
		return color("PANIC", useColor)
	}

	return "-----"
}

func color(level string, useColor bool) string {
	if !useColor {
		return level
	}

	switch level {
	case "TRACE":
		return Purple + level + Reset
	case "DEBUG":
		return Blue + level + Reset
	case "INFO ":
		return Green + level + Reset
	case "WARN ":
		return Yellow + level + Reset
	case "ERROR":
		return Red + level + Reset
	case "FATAL":
		return Red + level + Reset
	case "PANIC":
		return Red + level + Reset
	}

	return "-----"
}
