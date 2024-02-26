package sidecarbackup

import (
	"os"
	"os/exec"
	"time"
)

var execCommand = exec.Command

var existsCommand = Exists
var removeCommand = Remove

var utilStatCommand = os.Stat
var utilIsNotExistCommand = os.IsNotExist
var utilRemoveCommand = os.Remove

var utilCreateCommand = os.OpenFile

func Exists(target string) bool {
	var err error
	_, err = utilStatCommand(target)
	return !utilIsNotExistCommand(err)
}

func CreateFile(target string) error {
	var f *os.File
	var err error
	f, err = utilCreateCommand(target, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	now := time.Now()
	if err = os.Chtimes(target, now, now); err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func Remove(target string) error {
	return utilRemoveCommand(target)
}
