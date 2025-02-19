package main

import (
	"fmt"
	"go-file/entity"
	"go-file/utils"
)

func main() {
	var ds utils.Datasource[entity.Fruit] = &utils.FruitDatasource{Path: "./db/", Name: "fruits.txt"}

	// write
	fruits := []entity.Fruit{
		{Id: 1, Name: "Banana", Price: 20.5, Color: "yellow"},
		{Id: 2, Name: "Apple", Price: 30.7, Color: "red"},
		{Id: 3, Name: "Kiwi", Price: 50.3, Color: "green"},
	}

	err := ds.Write(fruits)
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

	fmt.Println("succesfully written to file")

	// read
	readFruits, err := ds.Read()
	if err != nil {
		fmt.Println("read error: ", err)
		return
	}

	fmt.Println("Fruits in file:")
	for _, f := range readFruits {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Color: %s\n", f.Id, f.Name, f.Price, f.Color)
	}
}