package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type Verification struct {
	Email string
	Hash  string
}

func SaveInFile(path, email, hash string) error {
	v := Verification{
		Email: email,
		Hash:  hash,
	}

	file, err := os.Create(path) // создаём или перезаписываем файл
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // красиво отформатированный JSON
	return encoder.Encode(v)
}

func CompareText(path, textForComparing string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return false
	}

	var v Verification
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return false
	}

	return v.Hash == textForComparing
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
