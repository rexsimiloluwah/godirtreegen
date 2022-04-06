/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rexsimiloluwah/godirgen/cmd"
)

func main() {
	styleInput := flag.String("style", "plain", "Output Style {plain|classic}")
	pathInput := flag.String("path", ".", "A valid path, default is the current working directory => .")
	mdOut := flag.String("o", "", "Output filename for the generated file tree diagram i.e. filetree.md, filetree.docx etc.")
	ignoreDirsInput := flag.String("ignore", "", "Comma separated list of folders to ignore when traversing.\n This could typically include large folders i.e. node_modules,.git etc.")
	showFileSize := flag.Bool("size", false, "Display file size")
	flag.Parse()

	ignoreDirs := readFolderIgnore()
	ignoreDirs = append(ignoreDirs, strings.Split(strings.Trim(*ignoreDirsInput, " "), ",")...)

	if *styleInput != "plain" && *styleInput != "classic" {
		fmt.Printf("Style %s is invalid.", *styleInput)
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		dirTree := cmd.NewDirectoryTree(*pathInput, *styleInput, ignoreDirs, *showFileSize)
		dirTreeDiagram := dirTree.DirectoryTreeDiagram()
		for _, d := range dirTreeDiagram {
			fmt.Println(d)
		}
		if *mdOut != "" {
			WriteToMd(*mdOut, dirTreeDiagram)
		}
	}

}

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
func readFolderIgnore() []string {
	resp, _ := ioutil.ReadFile(".folderignore")
	return strings.Split(strings.Trim(string(resp), " \n"), "\n")
}
