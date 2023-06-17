package main

import (
	"log"
	"net/http"

	EmployeeController "github.com/ahk4815/Go-Rest-Service/Controller"
)

func handleRequests() {
	http.HandleFunc("/", EmployeeController.HomePage)
	http.HandleFunc("/employee", EmployeeController.GetAllEmployees)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
