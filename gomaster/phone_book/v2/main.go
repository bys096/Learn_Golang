package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
	}

}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(0, 64)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func random(min, max int) int {
	return rand.Intn(max-min) + min

}

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			fmt.Println("key, ", key)
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	// 커맨드 라인 인수가 없다면
	fmt.Println("args[2]: ", arguments[2], "length:", len(arguments))

	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	populate(100)
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			list()
			return
		}
		fmt.Println(*result)

	case "list":
		list()

	case "populate":

		list()
	default:
		fmt.Println("Not a valid option")

	}

}
