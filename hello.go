package main

// 표준 출력을 위한 기본 라이브러리
import (
	"fmt"

	"rsc.io/quote"
)

// 외부 라이브러리 (visit pkg.go.dev and search for a "quote")

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
}
