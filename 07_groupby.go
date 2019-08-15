package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Printf("input: %v\n", animals)
	stream := koazee.StreamOf(animals)
	fmt.Print("stream.GroupBy(strings.Len): ")
	out, _ := stream.GroupBy(func(val string)int{return len(val)})
	fmt.Println(out)
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.GroupBy(strings.Len): map[5:[tiger] 4:[lynx lion] 3:[dog cat dog fox] 6:[monkey]]
*/
