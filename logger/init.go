package logger

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"
)

type envLog struct {
	Type  string `json:"type"`
	Level string `json:"level"`
	Path  string `json:"path"`
}

func (e *envLog) isValid() bool {
	if ParseType(e.Type) == TypeUnknown {
		return false
	}
	if ParseLevel(e.Level) == LevelUnknown {
		return false
	}
	if (ParseType(e.Type) == TypeStandardOut || ParseType(e.Type) == TypeStandardError) && e.Path != "" {
		return false
	}
	if ParseType(e.Type) == TypeFile && e.Path == "" {
		return false
	}

	return true
}

func (e *envLog) parseLog() (Log, error) {
	level := ParseLevel(e.Level)
	switch ParseType(e.Type) {
	case TypeNone:
		return NewLogConsole(level, ioutil.Discard), nil
	case TypeStandardOut:
		return NewLogConsole(level, os.Stdout), nil
	case TypeStandardError:
		return NewLogConsole(level, os.Stderr), nil
	case TypeFile:
		x, err := os.OpenFile(e.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			return nil, err
		}
		return NewLogFile(level, x), nil
	}

	return nil, xerrors.New("不正なロガータイプです")
}

func init() {
	v := os.Getenv("LOGGER")
	if v == "" {
		return
	}
	err := Initialize(v)
	if err != nil {
		panic(err)
	}
}

func Initialize(v string) error {
	logs, err := jsonParse(v)
	if err != nil {
		return err
	}
	for _, x := range logs {
		if x.isValid() {
			return xerrors.New("バリデーションチェックでエラーが発生しました")
		}
		y, err := x.parseLog()
		if err != nil {
			return err
		}
		Add(y)
	}

	return nil
}

func jsonParse(v string) ([]*envLog, error) {
	var envLogs []*envLog
	err := json.Unmarshal([]byte(v), &envLogs)
	if err != nil {
		return nil, err
	}
	return envLogs, nil
}
