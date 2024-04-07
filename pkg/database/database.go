package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type ComicInfo struct {
	URL      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

type ComicsDatabase map[string]ComicInfo

// SaveToFile сериализует данные в JSON и сохраняет их в указанный файл.
func SaveToFile(filename string, data interface{}) error {
	// Открываем файл с возможностью добавления в конец и создания, если файл не существует
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Создаем новый JSON encoder, который пишет в файл
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // Устанавливаем отступы для красивого форматирования

	// Записываем данные в файл
	if err := encoder.Encode(data); err != nil {
		return err
	}

	// Если необходимо добавить разделители между объектами JSON, можно записать разделитель в файл
	if _, err := file.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

// LoadFromFile читает данные из указанного файла и десериализует их из JSON.
func LoadFromFile(filename string, data interface{}) error {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}

// readFirstNComics читает данные из указанного файла и десериализует их из JSON.
func ReadFirstNComics(filepath string, n int) {
	jsonData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Десериализуем данные из JSON
	var comics map[string]ComicInfo
	if err := json.Unmarshal(jsonData, &comics); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Получаем и сортируем ключи для гарантии порядка обхода
	var keys []string
	for k := range comics {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Это важно, если ключи не являются строго упорядоченными числами

	// Выводим первые n объектов
	for i, k := range keys {
		if i >= n {
			break
		}
		fmt.Printf("=============\n")
		fmt.Printf("ID: %s\n URL: %s\n Keywords: %v\n", k, comics[k].URL, comics[k].Keywords)
		fmt.Printf("=============\n")
	}
	if err != nil {
		log.Fatalf("Error reading database: %v", err)
	}
	return
}
