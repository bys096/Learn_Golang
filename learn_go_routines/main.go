package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool) // 채널 생성
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isCool(person, c) // nico, flynn이 각각 호출한 총 2개의 함수로 채널을 보냄
	}
	fmt.Println("doing 1")
	result := <-c // 채널로부터 메시지를 받을 때까지, main function이 wait 한다
	fmt.Println(result)
	fmt.Println("doing 2")
	fmt.Println(<-c) // 위 코드와 같음
	// time.Sleep(time.Second * 5)
}

func isCool(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true // 채널로 true라는 메시지를 보냄(병렬로 처리됨으로 2개의 true 메시지가 채널로 도착)
	// return true
}
