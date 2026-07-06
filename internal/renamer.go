package internal

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func RenameFiles() {

	err := filepath.Walk("./sample", func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			fmt.Println("Error Occured")
			return err
		}

		if !info.IsDir() {

			new_name := Match(info.Name())

			if new_name != "" {
				new_file_path := fmt.Sprintf("sample/%s", new_name)

				fmt.Println("Executing Commands ... ")

				cmd := exec.Command("mv", path, new_file_path)

				err := cmd.Run()

				if err != nil {
					fmt.Println("Error Occured ", err)
				}

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
