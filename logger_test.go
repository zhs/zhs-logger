package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	l := New()
	f := Formatter{}

	e := l.WithField("", "")
	e.Message = "Test Message"
	e.Level = logrus.WarnLevel
	e.Time = time.Now()
	b, _ := f.Format(e)

	expected := strings.Join([]string{e.Time.Format(time.RFC3339), "[WARN ]:", "Test Message\n"}, " ")
	if string(b) != expected {
		t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
	}
}

func TestFormatterDefaultFormat(t *testing.T) {
	f := Formatter{}

	e := logrus.WithField("", "")
	e.Message = "Test Message"
	e.Level = logrus.WarnLevel
	e.Time = time.Now()

	b, _ := f.Format(e)

	expected := strings.Join([]string{e.Time.Format(time.RFC3339), "[WARN ]:", "Test Message\n"}, " ")
	if string(b) != expected {
		t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
	}
}
