// main.go
package main

import (
	"task7/Repositories"
	"task7/Delivery/routers"

)

func main() {
	db := Repositories.ConnectDB("mongodb://localhost:27017")

	r := routers.SetupRouter(db)
	r.Run(":8080")

}
