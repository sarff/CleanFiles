package main

import (
	"github.com/mitchellh/go-ps"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
)

func searchProcess(log logrus.FieldLogger) {
	inputStream := make(chan int)
	outputStream := make(chan error)

	processList, err := ps.Processes()
	if err != nil {
		log.Fatal("ps.Processes() Failed")
	}

	go sendPidToKill(inputStream, outputStream)

	go func() {
		defer close(inputStream)
		var process ps.Process
		for v := range processList {
			process = processList[v]
			if containsProgs(process.Executable()) {
				inputStream <- process.Pid()
			}
		}
	}()

	for x := range outputStream {
		if x != nil {
			log.Println(x)
		}
	}

}

func sendPidToKill(inputStream chan int, outputStream chan error) {
	for {
		val, ok := <-inputStream
		if !ok {
			break
		} else {
			outputStream <- pidKill(val)
		}
	}
	close(outputStream)
}

func pidKill(p int) error {

	kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(p))
	err := kill.Run()
	if err != nil {
		return err
	}
	return nil
}
