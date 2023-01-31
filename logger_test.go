package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	format = "2006-01-02 15:04:05"
)

func TestNew(t *testing.T) {
	l := New(nil)
	f := Formatter{}

	e := l.WithField("", "")
	e.Message = "Test Message"
	e.Level = logrus.WarnLevel
	e.Time = time.Now()
	b, _ := f.Format(e)

	expected := fmt.Sprintf("[%s]  WARN Test Message\n", e.Time.Format(format))
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

	expected := fmt.Sprintf("[%s]  WARN Test Message\n", e.Time.Format(format))
	if string(b) != expected {
		t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
	}
}
