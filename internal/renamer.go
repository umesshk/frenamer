package internal

import (
	"fmt"
	"io/fs"
	"os"
	_ "os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type FiletoRename struct {
	OldPath string
	NewPath string
}

type FileInfo struct {
	DirName  string
	FileName string
}

func RenameFiles() error {

	var FileList []FiletoRename
	var FilestoRename []FileInfo

	FileMap := make(map[string][]string)

	err := filepath.Walk("./sample", func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			fmt.Println("Error Occured")
			return err
		}

		if !info.IsDir() {

			var dirName string

			splitPath := strings.Split(path, "/")

			splitPath = splitPath[:len(splitPath)-1]

			dirName = strings.Join(splitPath, "/")

			new_name := Match(info.Name(), 0, 0)

			if new_name != "" {

				FileMap[dirName] = append(FileMap[dirName], path)
				current_file := FileInfo{FileName: path, DirName: dirName}
				FilestoRename = append(FilestoRename, current_file)

			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Printing File...")

	for _, list := range FileMap {

		count := len(list)

		for i, filename := range list {
			new_file_name := Match(filename, count, i+1)
			FileList = append(FileList, FiletoRename{OldPath: filename, NewPath: new_file_name})
		}

	}

	for _, f := range FileList {
		fmt.Printf("%s => %s\n", f.OldPath, f.NewPath)
		os.Rename(f.OldPath, f.NewPath)

	}

	fmt.Println("All Files are succefully renamed")

	return nil

}

func Match(fileName string, count, index int) string {

	Filetmp := strings.Split(fileName, ".")
	ext := Filetmp[len(Filetmp)-1]

	Filetmp = Filetmp[:len(Filetmp)-1]

	filelNameWithoutext := strings.Join(Filetmp, ".")

	splitedFileName := strings.Split(filelNameWithoutext, "_")

	nameStringSplit := splitedFileName[:len(splitedFileName)-1]

	nameString := strings.Join(nameStringSplit, "_")

	if nameString == "" {
		return ""
	}

	fileNumString := splitedFileName[len(splitedFileName)-1]

	_, err := strconv.Atoi(fileNumString)

	if err != nil {
		fmt.Printf("Inavlid File Format : %s\n", fileName)
		return ""
	}

	newName := fmt.Sprintf("%s-%d of %d.%s", nameString, index, count, ext)

	return newName
}
