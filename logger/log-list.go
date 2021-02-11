package logger

import (
	"encoding/json"
	"strings"
	"time"
)

type logList []Log

func (l *logList) log(level Level, v interface{}) error {
	s, err := l.generateLog(v)
	if err != nil {
		return err
	}
	for _, x := range logger {
		if x.IsLog(level) {
			x.Log(s)
		}
	}

	return nil
}

func (l *logList) generateLog(v interface{}) (string, error) {
	x, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	s := `{`
	s = s + strings.Join(
		[]string{
			`"time":"` + time.Now().Format("2006/01/02/ 15:04:05.000000Z07:00") + `"`,
			`"message":` + string(x),
		},
		",",
	)
	s = s + `}`

	return s, nil
}
