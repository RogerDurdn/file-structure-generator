package main

import (
	"file-structure-generator/pkg"
	"os"
)

func main() {
	dataPath := os.Getenv("DATA_PATH")
	pkg.Generate(dataPath)
}
