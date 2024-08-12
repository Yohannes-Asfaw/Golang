package task4

import (
    "task4/router"
)

func main() {
    r := router.SetupRouter()
    r.Run("localhost:8080")
}
