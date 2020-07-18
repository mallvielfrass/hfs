package main

import (
	"./hfs"
)

func printP() {
	print(" ")
}
func main() {
	fs := hfs.API("localhost:9999")
	data := fs.Ls()
	for i := 0; i < len(data.Item); i++ {
		print(data.Item[i].Name)
		printP()

		print(data.Item[i].IsDir)
		printP()
		print(data.Item[i].Size)
	}
}
