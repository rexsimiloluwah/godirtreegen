<h1 align="center">
    GODIRTREEGEN 📁🚀
</h1>
<h4 align="center">
    A Golang CLI for automatically generating directory tree diagrams
</h4>

## Description 
This is a CLI tool for generating customizable directory tree diagrams. Similar to what the `tree` command achieves on Unix-like systems. 

## How does this work? 
It recursively traverses through all sub-directories and files under the directory, and prints them out in a tree-like, readable format.

## Installation 
### Install by downloading pre-compiled binaries 
You can easily download the pre-compiled binaries for your Linux, Windows, FreeBSD, or MAC OS from the [releases](https://github.com/rexsimiloluwah/godirtreegen/releases) section. 

### Installing globally using Go 
```bash 
$ go install github.com/rexsimiloluwah/godirtreegen@latest
```

### Install via docker 
```bash
$ docker pull similoluwaokunowo/godirtreegen@latest 
```

### Compile from source 
1. Clone the repository 
```bash
$ git clone https://github.com/rexsimiloluwah/godirtreegen
$ cd godirtreegen 
```

2. Install using make 
```bash
$ make install
```

## Usage 
```
Usage of godirtreegen:
  -ignore string
        Comma separated list of folders to ignore when traversing.
         This could typically include large folders i.e. node_modules,.git etc.
  -o string
        Output filename for the generated file tree diagram i.e. filetree.md, filetree.docx etc.
  -path string
        A valid path, default is the current working directory => . (default ".")
  --size boolean
        Display file size
  -style string
        Output Style {plain|classic} (default "plain")
```

[View Examples Commands](./examples/commands.md)


