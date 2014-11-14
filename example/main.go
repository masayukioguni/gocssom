package main

import (
	"github.com/masayukioguni/gocssom/cssom"
	"io/ioutil"
	"os"
)

func main() {

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	css := cssom.Parse(string(b))
	css.Print()

}
