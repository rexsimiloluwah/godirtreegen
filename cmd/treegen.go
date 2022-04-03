package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var PIPE string = "│"
var ELBOW string = "└──"
var TEE string = "├──"
var PIPE_PREFIX string = "│   "
var SPACE_PREFIX string = strings.Repeat(" ", 4)

// Directory Tree Node
type DirectoryNode struct {
	Name     string
	Path     string
	Children []*DirectoryNode
}

// Directory Tree
type DirectoryTree struct {
	ParentNode *DirectoryNode
	Diagram    []string
	Style      string
	IgnoreDirs []string
}

type DirectoryTreeInterface interface {
	BuildTree()
	GenerateTreeDiagramHead()
	GenerateTreeDiagramBody(node *DirectoryNode, prefix string)
	AddFile(c *DirectoryNode, connector string, prefix string)
	AddFolderNested(c *DirectoryNode, connector string, prefix string, depth int, numEntries int)
	AddFolder(c *DirectoryNode, connector string, prefix string)
	DirectoryTreeDiagram() []string
}

// Initialize a new DirectoryTree struct/object
func NewDirectoryTree(rootDir string, style string, ignoreDirs []string) DirectoryTreeInterface {
	parentNode := &DirectoryNode{
		Name:     rootDir,
		Path:     rootDir,
		Children: make([]*DirectoryNode, 0),
	}
	return &DirectoryTree{
		parentNode,
		make([]string, 0),
		style,
		ignoreDirs,
	}
}

// Build the tree from the parent directory node
func (dt *DirectoryTree) BuildTree() {
	processNode(dt.ParentNode)
}

// Some of the logic for generating the file tree diagram was gleaned from a blog post by RealPython
// Credits: https://realpython.com/directory-tree-generator-python/
func (dt *DirectoryTree) GenerateTreeDiagramHead() {
	head := []string{dt.ParentNode.Name + "/", PIPE}
	dt.Diagram = append(dt.Diagram, head...)
}

// Generate the body of the directory tree diagram
func (dt *DirectoryTree) GenerateTreeDiagramBody(node *DirectoryNode, prefix string) {
	var connector string
	numEntries := len(node.Children)
	for idx, c := range node.Children {
		if idx != numEntries-1 { // check if we have gotten to the last item
			connector = TEE
		} else {
			connector = ELBOW
		}
		if len(c.Children) != 0 {
			if !contains(dt.IgnoreDirs, c.Name) {
				dt.AddFolderNested(c, connector, prefix, idx, numEntries)
			} else {
				dt.AddFolder(c, connector, prefix)
			}
		} else {
			dt.AddFile(c, connector, prefix)
		}
	}
}

// Add only the folder to the directory tree diagram (useful for large or trivial folders)
func (dt *DirectoryTree) AddFolder(c *DirectoryNode, connector string, prefix string) {
	if dt.Style == "classic" {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s 📂 %s", prefix, connector, c.Name))
	} else {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s %s", prefix, connector, c.Name))
	}
}

// Add a folder and its nested children to the directory tree diagram
func (dt *DirectoryTree) AddFolderNested(c *DirectoryNode, connector string, prefix string, depth int, numEntries int) {
	if dt.Style == "classic" {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s 📂 %s", prefix, connector, c.Name))
	} else {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s %s", prefix, connector, c.Name))
	}
	if depth != numEntries-1 {
		prefix += PIPE_PREFIX
	} else {
		prefix += SPACE_PREFIX
	}
	dt.GenerateTreeDiagramBody(c, prefix)
	dt.Diagram = append(dt.Diagram, PIPE)
}

// Add a file to the directory tree diagram
func (dt *DirectoryTree) AddFile(c *DirectoryNode, connector string, prefix string) {
	if dt.Style == "classic" {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s 📜 %s", prefix, connector, c.Name))
	} else {
		dt.Diagram = append(dt.Diagram, fmt.Sprintf("%s%s %s", prefix, connector, c.Name))
	}
}

// Returns a slice containing the directory tree diagram
func (dt *DirectoryTree) DirectoryTreeDiagram() []string {
	dt.GenerateTreeDiagramHead()
	dt.BuildTree()
	dt.GenerateTreeDiagramBody(dt.ParentNode, "")
	return dt.Diagram
}

// Check if a given path is folder/directory or file
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// Recursively traverse a given node in the file directory tree
func processNode(node *DirectoryNode) {
	entries := processEntries(node.Path)
	for _, entry := range entries {
		newDirectoryNode := &DirectoryNode{
			Name: entry.Name(),
			Path: filepath.Join(node.Path, entry.Name()),
		}
		if isDir, _ := isDirectory(newDirectoryNode.Path); isDir {
			node.Children = append(node.Children, newDirectoryNode)
			processNode(newDirectoryNode)
		} else {
			node.Children = append(node.Children, newDirectoryNode)
		}
	}
}

// Process all the files and directories in a given path
func processEntries(path string) []fs.FileInfo {
	// sort the directories and file path names
	var dirs []fs.FileInfo
	var files []fs.FileInfo
	var sortedEntries []fs.FileInfo
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("An error occured while reading directory: %s", err.Error())
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		} else {
			files = append(files, entry)
		}
	}
	sortedEntries = append(sortedEntries, dirs...)
	sortedEntries = append(sortedEntries, files...)
	return sortedEntries
}

// Utility function to check if a slice contains an element
func contains(slice []string, el string) bool {
	for _, s := range slice {
		if s == el {
			return true
		}
	}
	return false
}
