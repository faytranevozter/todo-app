package main

import (
"log"
"net/http"

"github.com/faytranevozter/todo-app/backend"
)

func main() {
router := backend.SetupRouter()
log.Println("Server starting on :8080")
if err := http.ListenAndServe(":8080", router); err != nil {
log.Fatal(err)
}
}
