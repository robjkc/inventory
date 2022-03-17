package main

import (
	"encoding/json"
	"fmt"
	"inventory/handler"
	"inventory/types"
	"io/ioutil"
	"os"
)

func main() {
	data := loadData()
	displayMostExpensive(data)

	displayLongCds(data)

}

func displayLongCds(data []types.Item) {
	fmt.Println("Long cds..")
	cds := handler.LongCds(data)
	for _, item := range cds {
		fmt.Printf("Item: %+v\n", item)
	}
	fmt.Println()
}

func displayMostExpensive(data []types.Item) {
	fmt.Println("Top 5 most expensive books..")
	mostExpensive := handler.MostExpensive(data, "book")
	for _, item := range mostExpensive {
		fmt.Printf("Item: %+v\n", item)
	}
	fmt.Println()

	fmt.Println("Top 5 most expensive cds..")
	mostExpensive = handler.MostExpensive(data, "cd")
	for _, item := range mostExpensive {
		fmt.Printf("Item: %+v\n", item)
	}
	fmt.Println()

	fmt.Println("Top 5 most expensive dvds..")
	mostExpensive = handler.MostExpensive(data, "dvd")
	for _, item := range mostExpensive {
		fmt.Printf("Item: %+v\n", item)
	}
	fmt.Println()
}

func loadData() []types.Item {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Items array
	var items []types.Item

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &items)
	fmt.Println(len(items))
	/*for _, item := range items {
		fmt.Println("Title: " + item.Title)
		fmt.Printf("Price: %f\n", item.Price)
	}*/
	return items
}
