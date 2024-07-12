package morse_test

import (
	"atp/sandi/repository/morse"
	"context"
	"fmt"
	"log"
	"testing"
)

func Test_DecodexE(t *testing.T) {
	repo := morse.NewRepository()
	ctx := context.Background()

	word := "Hello World 123" //"aBcD Ef01"
	log.Printf("encode this  :[%s]", word)
	result := repo.Encode(ctx, word)
	encode := ""
	for j, val := range result {
		//log.Printf("[%d] %s\n", j, val)
		encode += val
		if (j + 1) < len(result) {
			encode += "|"
		}
	}
	log.Printf("encode:[%s]", encode)

	fmt.Println()
	log.Printf("decode this :[%s]", encode)
	decoder := repo.Decode(ctx, encode)
	decode := ""
	for _, val := range decoder {
		encode += val
	}
	log.Printf("decode:[%s]", decode)
}
