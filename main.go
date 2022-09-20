package main

import (
	"fmt"

	"personal/guclient"
)

func main() {

	var responseObject, err = guclient.GetCards()

	if err != nil {
		fmt.Print(err.Error())
	}

	for i := 0; i < len(responseObject.Records); i++ {
		fmt.Println(responseObject.Records[i])
	}
}
