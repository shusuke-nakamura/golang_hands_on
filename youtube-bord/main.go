package main

import (
	"my"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	my.Migrate()
}
