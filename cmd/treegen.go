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
type DirNode struct {
	Name     string
	Path     string
	Size     float64
	Children []*DirNode
}

// Directory Tree
type DirTree struct {
	ParentNode *DirNode
}

// Directory Tree Printer
type DirTreePrinter struct {
	Tree         *DirTree
	Diagram      []string
	Style        string
	IgnoreDirs   []string
	ShowFileSize bool
}

type DirTreeInterface interface {
	BuildTree()
}

type DirTreePrinterInterface interface {
	GenerateTreeDiagramHead()
	GenerateTreeDiagramBody(node *DirNode, prefix string)
	AddFile(c *DirNode, connector string, prefic string)
	AddFolderNested(c *DirNode, connector string, prefix string, depth int, numEntries int)
	AddFolder(c *DirNode, connector string, prefix string)
	GenerateDirTreeDiagram() []string
}

// Initialize a new DirectoryTree struct/object
func NewDirectoryTree(rootDir string) *DirTree {
	parentNode := &DirNode{
		Name:     rootDir,
		Path:     rootDir,
		Children: make([]*DirNode, 0),
	}
	return &DirTree{
		parentNode,
	}
}

// Initialize a new DirectoryTreePrinter struct/object
func NewDirectoryTreePrinter(tree *DirTree, style string, ignoreDirs []string, showFileSize bool) DirTreePrinterInterface {
	return &DirTreePrinter{
		tree,
		make([]string, 0),
		style,
		ignoreDirs,
		showFileSize,
	}
}

// Build the tree from the parent directory node
func (dt *DirTree) BuildTree() {
	processNode(dt.ParentNode)
}

// Some of the logic for generating the file tree diagram was gleaned from a blog post by RealPython
// Credits: https://realpython.com/directory-tree-generator-python/
func (p *DirTreePrinter) GenerateTreeDiagramHead() {
	head := []string{p.Tree.ParentNode.Name + "/", PIPE}
	p.Diagram = append(p.Diagram, head...)
}

// // Generate the body of the directory tree diagram
func (p *DirTreePrinter) GenerateTreeDiagramBody(node *DirNode, prefix string) {
	var connector string
	numEntries := len(node.Children)
	for idx, c := range node.Children {
		if idx != numEntries-1 { // check if we have gotten to the last item
			connector = TEE
		} else {
			connector = ELBOW
		}
		if len(c.Children) != 0 {
			if !contains(p.IgnoreDirs, c.Name) {
				p.AddFolderNested(c, connector, prefix, idx, numEntries)
			} else {
				p.AddFolder(c, connector, prefix)
			}
		} else {
			p.AddFile(c, connector, prefix)
		}
	}
}

// // Add only the folder to the directory tree diagram (useful for large or trivial folders)
func (p *DirTreePrinter) AddFolder(c *DirNode, connector string, prefix string) {
	if p.Style == "classic" {
		p.Diagram = append(p.Diagram, fmt.Sprintf("%s%s 📂 %s", prefix, connector, c.Name))
	} else {
		p.Diagram = append(p.Diagram, fmt.Sprintf("%s%s %s", prefix, connector, c.Name))
	}
}

// // Add a folder and its nested children to the directory tree diagram
func (p *DirTreePrinter) AddFolderNested(c *DirNode, connector string, prefix string, depth int, numEntries int) {
	if p.Style == "classic" {
		p.Diagram = append(p.Diagram, fmt.Sprintf("%s%s 📂 %s", prefix, connector, c.Name))
	} else {
		p.Diagram = append(p.Diagram, fmt.Sprintf("%s%s %s", prefix, connector, c.Name))
	}
	if depth != numEntries-1 {
		prefix += PIPE_PREFIX
	} else {
		prefix += SPACE_PREFIX
	}
	p.GenerateTreeDiagramBody(c, prefix)
	p.Diagram = append(p.Diagram, PIPE)
}

// // Add a file to the directory tree diagram
func (p *DirTreePrinter) AddFile(c *DirNode, connector string, prefix string) {
	if p.Style == "classic" {
		s := fmt.Sprintf("%s%s 📜 %s", prefix, connector, c.Name)
		if p.ShowFileSize {
			s += fmt.Sprintf("  (%s)", convertBytes(c.Size))
		}
		p.Diagram = append(p.Diagram, s)
	} else {
		s := fmt.Sprintf("%s%s %s", prefix, connector, c.Name)
		if p.ShowFileSize {
			s += fmt.Sprintf("  (%s)", convertBytes(c.Size))
		}
		p.Diagram = append(p.Diagram, s)
	}
}

// Return the directory tree diagram elements
func (p *DirTreePrinter) GenerateDirTreeDiagram() []string {
	return p.Diagram
}

// Returns a slice containing the directory tree diagram
func DirectoryTreeDiagram(path string, style string, ignoreDirs []string, showFileSize bool) []string {
	// Initialize a new directory tree
	dirTree := NewDirectoryTree(path)
	// Traverse the sub-directories and files and build the tree
	dirTree.BuildTree()
	// Initialize a new directory tree printer
	dirTreePrinter := NewDirectoryTreePrinter(
		dirTree,
		style,
		ignoreDirs,
		showFileSize,
	)
	dirTreePrinter.GenerateTreeDiagramHead()
	dirTreePrinter.GenerateTreeDiagramBody(dirTree.ParentNode, "")
	// Return the directory tree diagram
	diagram := dirTreePrinter.GenerateDirTreeDiagram()
	return diagram
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
func processNode(node *DirNode) {
	entries := processEntries(node.Path)
	for _, entry := range entries {
		newDirectoryNode := &DirNode{
			Name: entry.Name(),
			Size: float64(entry.Size()),
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

// Utility function to convert file size (bytes) to human readable format
func convertBytes(size float64) string {
	units := []string{"B", "KB", "MB", "GB", "PB", "TB"}
	idx := 0
	for size > 1000 {
		size = size / 1024
		idx++
	}
	return fmt.Sprintf("%.2f%s", size, units[idx])
}
