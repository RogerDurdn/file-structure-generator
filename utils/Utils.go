package utils

import (
	"file-structure-generator/model"
	"fmt"
	"log"
	"os"
)

func Manage(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func FinalName(record *model.Record) string {
	fName := os.Getenv("F_NAME")
	return fmt.Sprintf("%s_%s_%s_"+fName, record.Topic, record.Unit, record.Name)
}
