package processes

import (
	"github.com/mitchellh/go-ps"
	"github.com/sirupsen/logrus"
	"os"
)

func whatever(log logrus.FieldLogger) {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}
}

func logKill(p *os.Process, log logrus.FieldLogger) error {
	log.Printf("killing PID %d", p.Pid)
	err := p.Kill()
	if err != nil {
		log.Print(err)
	}
	return err
}
