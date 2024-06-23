package must

import "github.com/k0kubun/pp"

func Handle(err error) {
	if err != nil {
		pp.Println(err.Error())
		panic(err)
	}
}
