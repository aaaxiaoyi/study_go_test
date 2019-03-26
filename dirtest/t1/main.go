package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

const dataFile = "../t2/test.go"

func main() {

	_, filename, _, _ := runtime.Caller(0)
	datapath := path.Join(path.Dir(filename), dataFile)

	fmt.Println(datapath)

	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("open1", dataFile, "failed", err)
		return
	}
	fmt.Println("open", dataFile, "scuess")
	defer file.Close()

	file, err = os.Open(datapath)
	if err != nil {
		fmt.Println("open2", datapath, "failed")
		return
	}
	fmt.Println("open", datapath, "scuess")
	defer file.Close()
}
