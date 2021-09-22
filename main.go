package main

import (
	"go-db2entity/to_entity"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := ""
	if len(os.Args) != 2 {
		log.Fatal("参数错误")
		return
	}
	filePath = os.Args[1]
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	to_entity.ToEntity(filePath, string(b))
}



