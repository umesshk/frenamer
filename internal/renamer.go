package internal

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
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
	var rename bool

	flag.BoolVar(&rename, "rename", false, "Pass the flag to actually rename the files")
	flag.Parse()

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

			dirName = filepath.Dir(path)

			new_name := Match(info.Name())

			if new_name != nil {

				key := fmt.Sprintf("%s/%s", dirName, new_name.Base)

				FileMap[key] = append(FileMap[key], path)
				current_file := FileInfo{FileName: path, DirName: dirName}
				FilestoRename = append(FilestoRename, current_file)

			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	for _, list := range FileMap {

		count := len(list)

		for i, filename := range list {
			nf := Match(filename)

			new_file_name := fmt.Sprintf("%s-%d of %d.%s", nf.Base, i+1, count, nf.Ext)

			FileList = append(FileList, FiletoRename{OldPath: filename, NewPath: new_file_name})
		}

	}
	for v, f := range FileMap {
		fmt.Printf("%s \n", v)
		fmt.Println(f)
	}

	for _, f := range FileList {
		fmt.Println("Renaming...")
		fmt.Printf("%s => %s\n", f.OldPath, f.NewPath)
		if rename == true {
			err := os.Rename(f.OldPath, f.NewPath)

			if err != nil {
				fmt.Printf("Error Renaming : %s to %s ", f.OldPath, f.NewPath)
			}
		}

	}

	fmt.Println("All Files are succefully renamed...")

	return nil

}

type FileMatch struct {
	Base string
	Ext  string
}

func Match(fileName string) *FileMatch {

	Filetmp := strings.Split(fileName, ".")
	ext := Filetmp[len(Filetmp)-1]

	Filetmp = Filetmp[:len(Filetmp)-1]

	filelNameWithoutext := strings.Join(Filetmp, ".")

	splitedFileName := strings.Split(filelNameWithoutext, "_")

	nameStringSplit := splitedFileName[:len(splitedFileName)-1]

	nameString := strings.Join(nameStringSplit, "_")

	if nameString == "" {
		return nil
	}

	fileNumString := splitedFileName[len(splitedFileName)-1]

	_, err := strconv.Atoi(fileNumString)

	if err != nil {
		fmt.Printf("Inavlid File Format : %s\n", fileName)
		return nil
	}

	return &FileMatch{Base: nameString, Ext: ext}
}
