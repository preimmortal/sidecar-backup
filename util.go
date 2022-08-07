package sidecarbackup

import (
	"os"
	"os/exec"
)

var execCommand = exec.Command

var existsCommand = Exists
var removeCommand = Remove

var utilStatCommand = os.Stat
var utilIsNotExistCommand = os.IsNotExist
var utilRemoveCommand = os.Remove


func Exists(target string) bool {
	var err error
	_, err = utilStatCommand(target)
	return !utilIsNotExistCommand(err)
}

func Remove(target string) error {
	return utilRemoveCommand(target)
}