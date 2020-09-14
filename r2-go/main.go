package main

import (
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/radare/r2pipe-go"
)

var inputFolder = flag.String("i", "/Users/yogesh/Work/Data/binaries/input", "Folder to watch for binaries")
var processedFolder = flag.String("p", "/Users/yogesh/Work/Data/binaries/processed", "Folder to move binaries after processing")
var outputFolder = flag.String("o", "/Users/yogesh/Work/Data/binaries/output", "Folder to publish cmd output")

var fileToWrite string

func main() {
	for true {
		files, err := ioutil.ReadDir(*inputFolder)
		if err != nil {
			print("Err: %v", *inputFolder, err)
			break
		}
		if len(files) == 0 {
			print(*inputFolder, "  is Empty. Retrying after 5 secs.\n")
			time.Sleep(5 * time.Second)
			continue
		}
		for _, file := range files {
			filename := *inputFolder + "/" + file.Name()
			processFile(filename)

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

	// disasm, err := r2p.Cmd("pd 20")
	// if err != nil {
	// 	print("ERROR: ", err)
	// } else {
	// 	print(disasm, "\n")
	// }

	ret, err := r2p.Cmd("ij")
	if err != nil {
		print("ERROR: ", err)
	} else {
		print(ret, "\n")
	}

	ret, err = r2p.Cmd("fs strings; fj")
	if err != nil {
		print("ERROR: ", err)
	} else {
		print(ret, "\n")
	}
}
