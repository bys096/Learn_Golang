package main

import (
	"fmt"
)

func main() {

	//b := make([]byte, 12)

	tmp := "가나다"
	for i := 0; i < len(tmp); i++ {
		//fmt.Println("i: " + strconv.Itoa(i) + ", " + string(tmp[i]))
	}

	slice := []byte(tmp)
	fmt.Println(string(slice))

}
