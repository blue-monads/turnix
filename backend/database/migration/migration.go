package migration

import (
	"fmt"

	"github.com/gobuffalo/fizz"
)

func Migrate() {

	table := fizz.NewTable("test", map[string]interface{}{})

	fmt.Println(table.Fizz())

}
