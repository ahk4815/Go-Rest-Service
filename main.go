package main

import (
	"log"
	"net/http"

	EmployeeController "github.com/ahk4815/Go-Rest-Service/Controller"
)

/*
This function creates a http router and binds the functions
to the url endpoints for the web service.
*/
func handleRequests() {
	http.HandleFunc("/", EmployeeController.HomePage)
	http.HandleFunc("/employee", EmployeeController.GetAllEmployees)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*
Main function.
*/
func main() {
	handleRequests()
}
