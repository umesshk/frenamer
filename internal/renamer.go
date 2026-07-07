package internal

import (
	"fmt"
	"io/fs"
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
	FileName string
	FilePath string
}

func RenameFiles() error {

	var FileList []FiletoRename
	var FilestoRename []FileInfo

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

			new_name := Match(info.Name())

			if new_name != "" {

				current_file := FileInfo{FileName: path, FilePath: dirName}
				FilestoRename = append(FilestoRename, current_file)

			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Printing File...")

	total_file_count := len(FilestoRename)

	fmt.Println(FilestoRename)
	fmt.Println(total_file_count)

	for _, f := range FileList {
		fmt.Printf("%s => %s ", f.OldPath, f.NewPath)
	}

	fmt.Println("All Files are succefully renamed")

	return nil

}

func Match(fileName string) string {

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

	fileNum, err := strconv.Atoi(fileNumString)

	if err != nil {
		fmt.Printf("Inavlid File Format : %s\n", fileName)
		return ""
	}

	newName := fmt.Sprintf("%s-%d of 4.%s", nameString, fileNum, ext)

	return newName
}
