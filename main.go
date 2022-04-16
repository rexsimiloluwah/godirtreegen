package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rexsimiloluwah/godirtreegen/cmd"
	"github.com/rexsimiloluwah/godirtreegen/utils"
)

func main() {
	// Reading the command line arguments
	styleInput := flag.String("style", "plain", "Output Style {plain|classic}")
	pathInput := flag.String("path", ".", "A valid path, default is the current working directory => .")
	mdOut := flag.String("o", "", "Output filename for the generated file tree diagram i.e. filetree.md, filetree.docx etc.")
	ignoreDirsInput := flag.String("ignore", "", "Comma separated list of folders to ignore when traversing.\n This could typically include large folders i.e. node_modules,.git etc.")
	showFileSize := flag.Bool("size", false, "Display file size")
	// Parsing the command line arguments
	flag.Parse()

	// Reading the default folders to ignore (from .folderignore)
	ignoreDirs := utils.ReadFolderIgnore(".folderignore")
	ignoreDirs = append(ignoreDirs, strings.Split(strings.Trim(*ignoreDirsInput, " "), ",")...)

	if *styleInput != "plain" && *styleInput != "classic" {
		fmt.Printf("Style %s is invalid.", *styleInput)
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		dirTreeDiagram := cmd.DirectoryTreeDiagram(*pathInput, *styleInput, ignoreDirs, *showFileSize)
		for _, d := range dirTreeDiagram {
			// print the directory tree to stdout
			fmt.Println(d)
		}
		if *mdOut != "" {
			utils.WriteToMd(*mdOut, dirTreeDiagram)
		}
	}
}
