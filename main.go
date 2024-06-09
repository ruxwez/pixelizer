package main

import (
	"os"
	"pixelizer/transformers"
	"pixelizer/util"
)

func main() {
	os.Mkdir("./in", os.ModeAppend)

	files, err := util.GetAllFiles("./in")

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		transformers.Apng(file, 5)
	}
}
