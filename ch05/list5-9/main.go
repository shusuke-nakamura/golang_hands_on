package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	p := "https://golang.org"

	// doc, er := goquery.NewDocument(p)
	// if er != nil {
	// 	panic(er)
	// }

	// https://pkg.go.dev/github.com/PuerkitoBio/goquery#example-package
	// 上記を参考してサンプルコードを変更

	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}

	defer re.Body.Close()

	doc, er := goquery.NewDocumentFromReader(re.Body)

	if er != nil {
		panic(er)
	}

	doc.Find("a").Each(func(n int, sel *goquery.Selection) {
		lk, _ := sel.Attr("href")
		println(n, sel.Text(), "(", lk, ")")
	})
}
