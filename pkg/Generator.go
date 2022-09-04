package pkg

import (
	"file-structure-generator/model"
	"file-structure-generator/utils"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"path/filepath"
)

var (
	currentDir  = os.Getenv("CURR_DIR")
	templateDir = os.Getenv("TEMP_DIR")
)

func Generate(pathData string) {
	records := obtainRecordsFromFile(pathData)
	tree := createTree(records)
	createFiles(tree)
}

func createFiles(tree *model.TreeTopic) {
	for topicName, topic := range tree.Topics {
		topicDir := createDir(currentDir, topicName)
		for unitName, unit := range topic.Units {
			unitDir := createDir(topicDir, unitName)
			for _, record := range unit {
				createFileFromTemplate(unitDir, record)
			}
		}
	}
}

func createDir(dirPath, dirName string) string {
	dir := filepath.Join(dirPath, dirName)
	utils.Manage(os.MkdirAll(dir, os.ModePerm))
	return dir
}

func createFileFromTemplate(dirPath string, record *model.Record) {
	srcFile, err := os.Open(filepath.Join(templateDir, "temp-"+record.Topic+".docx"))
	utils.Manage(err)
	defer srcFile.Close()

	destFile, err := os.Create(filepath.Join(dirPath, utils.FinalName(record)+".docx")) // creates if file doesn't exist
	utils.Manage(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	utils.Manage(err)

	err = destFile.Sync()
	utils.Manage(err)
}

func createTree(records []*model.Record) *model.TreeTopic {
	tree := model.NewTreeTopic()
	for _, record := range records {
		iterateTree(record, tree)
	}
	return tree
}

func iterateTree(record *model.Record, tree *model.TreeTopic) {
	topic, k := tree.Topics[record.Topic]
	if k {
		tasks, ok := topic.Units[record.Unit]
		if ok {
			topic.Units[record.Unit] = append(tasks, record)
		} else {
			topic.Units[record.Unit] = []*model.Record{}
			iterateTree(record, tree)
		}
	} else {
		tree.Topics[record.Topic] = model.NewTopic()
		iterateTree(record, tree)
	}
}

func obtainRecordsFromFile(pathFile string) []*model.Record {
	f, err := os.Open(pathFile)
	utils.Manage(err)
	defer f.Close()
	records := []*model.Record{}
	err = gocsv.UnmarshalFile(f, &records)
	utils.Manage(err)
	return records
}
