module main

go 1.19

require (
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	my v0.0.0
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
)

replace my => ./my
