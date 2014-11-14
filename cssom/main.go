package main

import (
	"github.com/masayukioguni/gocssom/cssom"
)

func main() {
	css := cssom.Parse(`div .a { 
					a: rgb(1,2,3)
					font-size: 150% !important
		 }`)
	css.Print()

}
