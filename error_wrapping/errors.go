package main

import (
	"errors"
	"fmt"
)

var (
	ErrRedis       = errors.New("레디스 관련 에러")
	ErrRedisGet    = fmt.Errorf("조회 도중 에러: %w", ErrRedis)
	ErrRedisDelete = fmt.Errorf("삭제 도중 에러: %w", ErrRedis)

	ErrRedisHashGet = fmt.Errorf("해쉬 조회 도중 에러: %w", ErrRedisGet)
)

// 상위에 해당하는 것과 비교해도 에러에 걸림
// 즉, 상위의 wrapping chain에 존재한다면 처리 가능.
// 다만, 동등 비교(==)로는 불가

// 이를 통해, 에러 핸들러를 만들 수 있음

func main() {
	err := ErrRedisHashGet
	//fmt.Println(err)
	if err != nil {
		if errors.Is(err, ErrRedisGet) {
			fmt.Println(ErrRedisGet)
		}
	}
}
