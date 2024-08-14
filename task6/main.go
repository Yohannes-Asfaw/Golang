package main

import (
    "task6/router"

)

func main() {
    r := router.SetupRouter()
    r.Run("localhost:8080")
}
