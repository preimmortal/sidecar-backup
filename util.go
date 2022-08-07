package sidecarbackup

import (
	"os"
)

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