package main

import (
    "task5/router"

)

func main() {
    r := router.SetupRouter()
    r.Run("localhost:8080")
}
