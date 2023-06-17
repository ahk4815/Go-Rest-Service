package EmployeeController

import (
	"encoding/json"
	"fmt"
	"net/http"

	EmployeeModel "github.com/ahk4815/Go-Rest-Service/Model"
)

var employees = []EmployeeModel.Employee{
	{ID: 1, Name: "John", Designation: "CEO"},
	{ID: 2, Name: "David", Designation: "CTO"},
	{ID: 3, Name: "Mary", Designation: "CFO"},
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Providing all employees - ")
	json.NewEncoder(w).Encode(employees)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
