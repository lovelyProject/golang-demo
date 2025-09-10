package file

import (
	"fmt"
	"os"
)

func SaveInFile(path, text string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func CompareText(path, textForComparing string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return false
	}

	return string(data) == textForComparing
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
