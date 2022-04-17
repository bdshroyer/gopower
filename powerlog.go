package gopower

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type PowerLogReader struct {
	Cmd      *exec.Cmd
	FilePath string
}

func NewPowerLogReader() (*PowerLogReader, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Unable to obtain current working directory")
	}

	pl := &PowerLogReader{
		Cmd:      exec.Command("PowerLog", "-duration", "1"),
		FilePath: filepath.Join(workDir, "PowerLog.csv"),
	}

	return pl, nil
}
