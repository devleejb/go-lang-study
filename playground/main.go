package main

import "example.com/hello/print"

func main() {
	go print.Quote2()

	print.Quote()
}
