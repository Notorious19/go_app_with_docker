package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoriment = 2
const delay = 5

func main() {

	intro()

	for {
		displayMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func intro() {
	name := "Victor"
	version := 1.2
	fmt.Println("Hello, mr.", name)
	fmt.Println("This program is on version", version)
}

func displayMenu() {
	fmt.Println("1- Starting monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit program")
	fmt.Println("")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("")
	fmt.Println("You've selected option", readCommand)
	fmt.Println("")

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{"https://www.google.com", "https://appv2.industrycare.com.br/signin",
		"https://www.alura.com.br", "https://www.youtube.com.br"}

	sites = readFileSites()

	for i := 0; i < monitoriment; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", sites)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("---------------------------------------------------------------------------")
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	reply, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	if reply.StatusCode == 200 {
		fmt.Println(site, "loaded succesfully!")
		fmt.Println("")
		registerLogs(site, true)

	} else {
		fmt.Println(site, "could not be loaded. Status Code:", reply.StatusCode)
		registerLogs(site, false)
	}

}

func readFileSites() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLogs(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("01/02/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
