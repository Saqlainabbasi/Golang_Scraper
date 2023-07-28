package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// define a product struct
type Product struct {
	name, image, url, price string
}

func main() {
	// declare a product slice of type Product struct
	products := []Product{}

	//create a new Intene of Colly Pkg....
	c := colly.NewCollector()

	url := "https://scrapeme.live/shop/"

	// get the data form the web Page ....
	// using an other colly method.....
	// this method takes string and a callback func as prameters.....second parameter is pointer
	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		// define a product
		product := Product{}

		product.url = h.ChildAttr("a", "href")
		product.image = h.ChildAttr("img", "src")
		product.name = h.ChildText("h2")
		product.price = h.ChildText(".price")
		// append add the element to the end of the slice......
		// fmt.Printf("product found: %q -> %s\n", product.price, product.name)
		products = append(products, product)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//use colly method to visit the url......
	c.Visit(url)
	// as the product is added to the slice we need to store it to the csv

	//open file using the Golang OS pkg....
	file, err := os.Create("product.csv")
	if err != nil {
		log.Fatalln("Failed to create the product.csv file", err)
	}
	defer file.Close()

	//initialize a file writer usign the csv pkg.....

	writer := csv.NewWriter(file)

	// make the headers for the csv file....
	headers := []string{
		"name", "image", "url", "price",
	}
	// set the headers...
	writer.Write(headers)

	// now we have to loop(range ) overt the products and save it to the csv...
	for _, p := range products {
		// make an slice of strings
		record := []string{
			p.name,
			p.image,
			p.url,
			p.price,
		}
		// for every itration of the loop we will save the reacord to the file
		writer.Write(record)
	}
	defer writer.Flush()
}
