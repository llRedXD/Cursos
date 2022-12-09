// - Partindo do código abaixo, utilize marshal para transformar []user em JSON.
package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}
    u := []user{u1,u2,u3}
    
    sb, err := json.Marshal(u)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(sb))

}