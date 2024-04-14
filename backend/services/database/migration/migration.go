package migration

import (
	"fmt"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/fizz/translators"
)

func Migrate() {

	table := fizz.NewTable("test", map[string]interface{}{})
	table.DisableTimestamps()

	t := translators.NewSQLite("")

	fmt.Println(fizz.AString(table.String(), t))

}
