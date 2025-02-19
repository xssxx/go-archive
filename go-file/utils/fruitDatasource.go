package utils

import (
	"bufio"
	"fmt"
	"go-file/entity"
	"log"
	"os"
	"strconv"
	"strings"
)

type FruitDatasource struct {
	Path string
	Name string
}

func (f *FruitDatasource) Read() ([]entity.Fruit, error) {
	file, err := os.Open(f.Path + f.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fruits []entity.Fruit

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fruit, err := parseFruit(line)
		if err != nil {
			log.Println("skipping invalid line: ", line)
			continue
		}
		fruits = append(fruits, fruit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fruits, nil
}

func (f *FruitDatasource) Write(fruits []entity.Fruit) error {
	file, err := os.Open(f.Path + f.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, fruit := range fruits {
		line := fmt.Sprintf("%d,%s,%.2f,%s\n", fruit.Id, fruit.Name, fruit.Price, fruit.Color)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}

func parseFruit(line string) (entity.Fruit, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 4 {
		return entity.Fruit{}, fmt.Errorf("invalid format")
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return entity.Fruit{}, err
	}

	price, err := strconv.ParseFloat(parts[2], 32)
	if err != nil {
		return entity.Fruit{}, err
	}

	return entity.Fruit{
		Id: uint16(id),
		Name: parts[1],
		Price: float32(price),
		Color: parts[3],
	}, nil
}