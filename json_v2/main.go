package main

import (
	jsonv2 "encoding/json/v2"
	"fmt"
	"log"
)

type UserS struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	data := []byte(`{"id": 1, "name": "Rafael", "email": "me@rflpazini.sh"}`)

	var u UserS
	if err := jsonv2.Unmarshal(data, &u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Usuário: %+v\n", u)

	out, _ := jsonv2.Marshal(u)
	fmt.Println(string(out))
}
