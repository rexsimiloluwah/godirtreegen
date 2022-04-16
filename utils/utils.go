package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Write the generated file tree diagram to a markdown file
func WriteToMd(fileName string, contents []string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(fileName), 0700)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	for _, c := range contents {
		_, _ = w.WriteString(c + "\n")
	}
	fmt.Printf("\nYipee! Directory Tree Diagram successfully written to %s 🎉🎉\n", fileName)
	w.Flush()
}

// Read the contents of .folderignore
func ReadFolderIgnore(filePath string) []string {
	resp, _ := ioutil.ReadFile(filePath)
	return strings.Split(strings.Trim(string(resp), " \n"), "\n")
}
