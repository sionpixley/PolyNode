package models

import "os/exec"

type ExecWrap struct{}

func (_ ExecWrap) Output(command *exec.Cmd) ([]byte, error) {
	return command.Output()
}

func (_ ExecWrap) Run(command *exec.Cmd) error {
	return command.Run()
}
