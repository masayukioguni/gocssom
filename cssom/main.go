package cssom

import (
	"github.com/masayukioguni/gocssom"
)

func main() {
	css := cssom.Parse(`div .a { 
					a: rgb(1,2,3)
					font-size: 150% !important
		 }`)
	css.Print()

}
