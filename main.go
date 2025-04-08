package main

import (
	"fmt"

	godiacritics "gopkg.in/Regis24GmbH/go-diacritics.v2"
)

func main() {
	noDiacrits := godiacritics.Normalize("än éᶍample")
	fmt.Println(noDiacrits)
}
