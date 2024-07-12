package main

import (
	"atp/sandi/repository/morse"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		log.Println("Usage:")
		log.Println("example: go run . encode")
		log.Println("example: go run . decode")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var data string
	fmt.Print("type here:")
	fmt.Scanf("%s", &data)
	length := len(data)
	if length <= 0 {
		log.Fatal("data is empty")
	}
	log.Printf("data:[%s]", data)

	ctx := context.Background()
	repo := morse.NewRepository()

	var result []string
	methode := flag.Args()[0]
	switch methode {
	case "encode":
		log.Println(" ========== time to encode ========== ")
		data = strings.ToUpper(data)
		log.Printf("word:[%s]", data)
		result = repo.Encode(ctx, data)

	case "decode":
		log.Println(" ========== time to decode ========== ")
		data = strings.ToUpper(data)
		log.Printf("word:[%s]", data)
		result = repo.Decode(ctx, data)

	default:
		log.Fatalf("failed cause methode not found")
	}

	conclusion := ""
	for i, val := range result {
		log.Printf("[%d] [%s]", i, val)
		conclusion += val
		if (i + 1) < len(result) {
			conclusion += "|"
		}

	}
	log.Printf("conclusion:[%s]", conclusion)
}
