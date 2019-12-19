package main

import (
	"fmt"
	"github.com/go-template/architecture/infrastructure/repository"
)

func main() {
	repositoryImpl := repository.NewRepositoryImpl()
	repositoryImpl.Run(func (con repository.Connectable) {
		connect := con.Connect()
		fmt.Printf("connect: %v\n", connect)
	})
}
