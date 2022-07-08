package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
)

//10gb //5Gb //1Gb
//https://speed.hetzner.de/10GB.bin
//http://speedtest-sgp1.digitalocean.com/5gb.test
//https://speed.hetzner.de/1GB.bin
//http://ipv4.download.thinkbroadband.com/1GB.zip
//http://ipv4.download.thinkbroadband.com/512MB.zip

//"fsutil file createnew E:\File.txt 1000000000" - 1Gb

func createNullFile(log logrus.FieldLogger, path, name string) {
	cmd := fmt.Sprintf("%s\\%s.txt", path, name)
	cmdprepare := exec.Command("fsutil", "file", "createnew", cmd, "1000000000")
	errcmd := cmdprepare.Run()
	if errcmd != nil {
		log.Println(errcmd)
	} else {
		log.Errorf("created fake: %s", cmd)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
