## Example commands 

1. Display help information 
```bash
$ godirtreegen --help 
```

2. Generate the file tree diagram (default directory = root path = ".")
```bash
$ godirtreegen  
```

**Output**
```markdown
./
│
├── .git
├── .github
│   └── workflows
│       ├── lint.yml
│       ├── publish-docker.yml
│       └── release.yml
│
│
├── cmd
│   └── treegen.go
│
├── scripts
│   └── install.sh
│
├── .dockerignore
├── .folderignore
├── .goreleaser.yml
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── godirgen.rb
├── main.go
└── run.sh 
```

> To specify a different path, run `godirtreegen -path </path/to/dir>` 

4. Generate the directory tree diagram with a fancy style 
```bash
$ godirtreegen -style classic 
```

**Output**
```markdown
./
│
├── 📂 .git
├── 📂 .github
│   └── 📂 workflows
│       ├── 📜 lint.yml
│       ├── 📜 publish-docker.yml
│       └── 📜 release.yml
│
│
├── 📂 cmd
│   └── 📜 treegen.go
│
├── 📂 scripts
│   └── 📜 install.sh
│
├── 📜 .dockerignore
├── 📜 .folderignore
├── 📜 .goreleaser.yml
├── 📜 Dockerfile
├── 📜 LICENSE
├── 📜 Makefile
├── 📜 README.md
├── 📜 go.mod
├── 📜 go.sum
├── 📜 godirgen.rb
├── 📜 main.go
└── 📜 run.sh 
```

5. Generate the directory tree diagram with the file sizes 
```bash
# Specify the `--size` flag 
$ godirtreegen --size 
```

**Output** 
```markdown
./
│
├── .git
├── .github
│   └── workflows
│       ├── lint.yml  (460.00B)
│       ├── publish-docker.yml  (601.00B)
│       └── release.yml  (994.00B)
│
│
├── cmd
│   └── treegen.go  (5.82KB)
│
├── scripts
│   └── install.sh  (1.41KB)
│
├── .dockerignore  (22.00B)
├── .folderignore  (50.00B)
├── .goreleaser.yml  (573.00B)
├── Dockerfile  (243.00B)
├── LICENSE  (11.09KB)
├── Makefile  (0.00B)
├── README.md  (2.33KB)
├── go.mod  (51.00B)
├── go.sum  (0.00B)
├── godirgen.rb  (1.08KB)
├── main.go  (2.48KB)
└── run.sh  (61.00B) 
```

6. Output the generated directory tree diagram to a file i.e. {.md, .doc, .txt, .docx} etc. 
```bash
$ godirtreegen -o <path/to/file>
```

7. Ignore specific sub-folders (this typically includes folders with many trivial files i.e. node_modules, .git etc, the [.folderignore](../.folderignore) file contains the folders excluded by default). Hence, those folders will not be further traversed in the directory tree diagram. 
```bash
# to add other folders to ignore (using the -ignore command with comma seperated folder names)
$ godirtreegen -ignore  <folder1>,<folder2>
# <folder1> and <folder2> will be ignored in the generated directory tree diagram
```
   


Enjoy 🎉🎉
