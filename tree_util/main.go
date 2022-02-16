package main

import (
	"fmt"
	"io/fs"
	"os"
)

func lastDirNumber(files []fs.DirEntry) int {
	for i := len(files) - 1; i >= 0; i-- {
		if files[i].IsDir() {
			return i
		}
	}
	return 0
}

func getFileSize(file fs.DirEntry) int {
	info, _ := file.Info()
	return int(info.Size())
}

func listDir(out *os.File, path string, printFiles bool, level int, tab string) {
	files, _ := os.ReadDir(path)
	fileSize := ""
	symb := "├───"
	add_tab := "│\t"
	for i := 0; i < len(files); i++ {
		if (i == len(files)-1) || (!printFiles && i == lastDirNumber(files)) {
			symb = "└───"
			add_tab = "\t"
		}
		if !files[i].IsDir() && printFiles {
			if getFileSize(files[i]) == 0 {
				fileSize = "empty"
			} else {
				fileSize = fmt.Sprintf("%db", getFileSize(files[i]))
			}
			fmt.Fprintf(out, "%s%s%s (%s)\n", tab, symb, files[i].Name(), fileSize)
		}
		if files[i].IsDir() {
			fmt.Fprintf(out, "%s%s%s\n", tab, symb, files[i].Name())
			listDir(out, path+"/"+files[i].Name(), printFiles, level+1, tab+add_tab)
		}
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	listDir(out, path, printFiles, 0, "")
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
