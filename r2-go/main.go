package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/gopacket/examples/util"
	"github.com/radare/r2pipe-go"
)

var inputFolder = flag.String("i", "/Users/yogesh/Work/Data/binaries/input", "Folder to watch for binaries")
var processedFolder = flag.String("p", "/Users/yogesh/Work/Data/binaries/processed", "Folder to move binaries after processing")
var outputFolder = flag.String("o", "/Users/yogesh/Work/Data/binaries/output", "Folder to publish cmd output")

var fileToWrite string

var config Config

func main() {
	defer util.Run()()

	loadConfig("./config.yml")

	prepareOutput()

	for true {

		//iterate input directory
		files, err := ioutil.ReadDir(*inputFolder)
		if err != nil {
			log.Printf("Error reading directory: %s, Err: %v", *inputFolder, err)
			break
		}

		//kep trying afger every 5 secs
		if len(files) == 0 {
			print(*inputFolder, "  is Empty. Retrying after 5 secs.\n")
			time.Sleep(5 * time.Second)
			continue
		}
		for _, file := range files {
			filename := *inputFolder + "/" + file.Name()

			log.Printf("Processing file: %s\n", filename)
			processFile(filename)

			//move processed file to processed folder
			processedfile := *processedFolder + "/" + file.Name()
			os.Rename(filename, processedfile)
		}
	}
}

func processFile(filename string) {

	r2p, err := r2pipe.NewPipe(filename)
	if err != nil {
		print("ERROR: ", err)
		return
	}
	defer r2p.Close()

	for _, c := range config.R2Commands {
		ret, err := r2p.Cmd(c.Cmd)
		if err != nil {
			print("ERROR: ", err)
		} else {
			if _, err := c.File.WriteString(ret + "\n"); err != nil {
				log.Println("Failed to write string to file.", err)
			}
			print(ret, "\n")
		}
	}
}

func prepareOutput() {
	for i, c := range config.R2Commands {
		var err error
		filename := *outputFolder + "/" + c.Idx + ".json"
		config.R2Commands[i].File, err = os.OpenFile(filename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Filed to create a file. Error: ", err)
		}
	}
}
