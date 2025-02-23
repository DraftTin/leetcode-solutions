package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {

	}
	defer file.Close()
	file.WriteString("inasdiansdiansd\n")
	file.WriteString("woieuqoiwueqoiweukjL:w\n")

}
