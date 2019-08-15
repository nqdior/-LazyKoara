package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wesovilabs/koazee"
)

var numbers = []int{}

func main() {

	maxValue := 100
	dummyCount := 5 * 10000000

	fmt.Printf("1. 0～100までのランダムな数値%d件を生成。\n", dummyCount)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dummyCount; i++ {
		numbers = append(numbers, rand.Intn(maxValue))
	}

	fmt.Printf("2. %d件生成完了。Koazeeに読み込み。\n", dummyCount)
	stream := koazee.StreamOf(numbers)

	/*
		このコメントアウトを外したら一覧が出力できます。超重いよ。
		fmt.Print("Koazeeに読み込み完了。一覧はコチラ ↓ \n")
		fmt.Println(stream.Out().Val())
		fmt.Println("\n")
	*/

	start := time.Now() // filter start
	count, _ := stream.Filter(func(value int) bool {
		return value == 10
	}).Count()
	end := time.Now() // filter end
	fmt.Printf("値が10のものだけを%f秒で抽出しました。\n", (end.Sub(start)).Seconds())
	fmt.Printf("対象件数は %d 件です。\n", count)
}

/**
go run koazeeFilter.go

1. 0～100までのランダムな数値50000000件を生成。
2. 50000000件生成完了。Koazeeに読み込み。
値が10のものだけを0.179520秒で抽出しました。
対象件数は 499951 件です。
*/
