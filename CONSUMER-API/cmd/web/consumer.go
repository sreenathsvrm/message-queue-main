package main

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func (app *application) consume(ctx context.Context) {

	for {

		// read message from kafka
		msg, err := app.reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		// decode encoded data
		productWrapper, err := decodeData(msg.Value)
		if err != nil {
			fmt.Println(err.Error())
		}

		// get all image url from database
		imageUrls, err := app.products.GetProductUrls(productWrapper.ProductID)
		if err != nil {
			log.Println(err)
		}

		// get all individual url 
		urlString := strings.Trim(imageUrls, "{}")
		urls := strings.Split(urlString, ",")

		// Set image destination
		destinationFolder := "./images"

		// download each of the image and save the paths
		var compressedImage []string
		for _, url := range urls {
			imagepath, err := downloadImage(url, destinationFolder)
			if err != nil {
				fmt.Println(err.Error())
			}

			compressedImage = append(compressedImage, imagepath)
		}

		// update compressed product urls to db
		err = app.products.UpdateProductUrls(compressedImage, productWrapper.ProductID)
		if err != nil {
			log.Println(err)
		}
	}
}
