package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", b)
	type Comics struct {
		Alt string
		Img string
	}
	var comics Comics

	if err = json.Unmarshal(b, &comics); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(comics)
	fmt.Println(comics.Img)
	fmt.Println(comics.Alt)

	resp, err = http.Get(comics.Img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	// img, _, err := image.Decode(bytes.NewReader(b))
	// if err != nil {
	// 	fmt.Println("Decode")
	// 	log.Fatalln(err)
	// }

	out, _ := os.Create("./img.jpeg")
	out.Write(b)
	defer out.Close()

	// var opts jpeg.Options
	// opts.Quality = 1

	// err = jpeg.Encode(out, img, &opts)
	// //jpeg.Encode(out, img, nil)
	// if err != nil {
	// 	fmt.Println("Encode")
	// 	log.Println(err)
	// }
	fmt.Printf("%s\n", b)
}
