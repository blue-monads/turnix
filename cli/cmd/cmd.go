package main

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
)

func ReadableInt64(rawcode int64) {

	code := fmt.Sprint(rawcode)
	final := ""

	for idx, digit := range code {

		i, err := strconv.Atoi(string(digit))

		if idx%2 == 0 {
			final = final + string(digit)
		} else {

			final = final + string(digit+49)
		}

		pp.Println("@curr", digit, i, err)
	}

	pp.Println("@final", final)

	//"@final" "9f6h2d6f"

}

func main() {

	ReadableInt64(95672368)

}
