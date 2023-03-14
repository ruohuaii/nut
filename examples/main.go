package main

import (
	"log"

	"github.com/ruohuaii/nut"
	"github.com/ruohuaii/nut/examples/types"
)

func main() {
	log.Println(nut.Generate(&types.Person{}, "types/types.go", true))
}
