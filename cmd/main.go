package main

import (
	"log"
	todo "todoAppGo"
)

func main() {
	srv:= new(todo.Server)
	if err:= srv.Run("9999"); err !=nil {
		log.Fatalf("error occured when run http server: %s", err.Error())
	}

}
