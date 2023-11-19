package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	fmt.Println("log path:", os.TempDir())
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Mastering Go, 3rd edition!")

	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")

}
