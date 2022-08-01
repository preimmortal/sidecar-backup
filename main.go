package main

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)
type Person struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
}
// Main function
func main() {
  log.Info("Starting Sidecar-Backup")
	var y []byte
	var err error
	p := &Person{
		"abc",
		3,
	}
	if y, err = yaml.Marshal(p); err != nil {
		log.Error(err)
	} 
	log.Info(string(y))
}
