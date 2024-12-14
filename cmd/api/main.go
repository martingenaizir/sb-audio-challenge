package main

import "log"

func main() {
	if err := newApplication(); err != nil {
		log.Fatal(err)
	}
}
