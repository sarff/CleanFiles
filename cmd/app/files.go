package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/user"
	"path/filepath"
)

func appdata(log logrus.FieldLogger) chan string {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	chann := make(chan string)
	go func() {
		filepath.Walk(user.HomeDir, func(path string, info os.FileInfo, _ error) (err error) {
			if info.IsDir() && containsProgs(info.Name()) {
				log.Errorf("funded: %s", path)
				chann <- path
			}
			return
		})
		defer close(chann)
	}()
	return chann

}

func downloads(log logrus.FieldLogger, ctlg string) chan string {

	chann := make(chan string)
	go func() {
		filepath.Walk(ctlg, func(path string, info os.FileInfo, _ error) (err error) {
			if info.IsDir() && containsProgs(info.Name()) {
				log.Errorf("funded: %s", path)
				chann <- path
			}
			return
		})
		defer close(chann)
	}()
	return chann

}
