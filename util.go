package sidecarbackup

import (
	"os"
)

func Exists(target string) bool {
	var err error
	_, err = os.Stat(target)
	return !os.IsNotExist(err)

}