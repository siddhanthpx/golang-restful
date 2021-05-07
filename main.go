package main

import (
	"golang-restful/client"
	"log"
)

func main() {
	_, err := client.NewClient()
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
