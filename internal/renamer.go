package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
)

func GetFile() {

	err := filepath.Walk("./sample", func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			fmt.Println("Error Occured")
			return err
		}

		if !info.IsDir() {
			fmt.Println("File Name : ", info.Name())
			new_name := Match(info.Name())
			if new_name != "" {
				fmt.Println("New File Name : ", new_name)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("some error occured ", err)

	}

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
