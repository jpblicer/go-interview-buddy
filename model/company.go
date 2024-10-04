package model

import (
	"fmt"
	"strings"
)

type Company struct {
	Name          string
	Website       string
	InterviewStep int
}

var CompanyList = make([]*Company, 0)

// Create
func AddCompany(name, website string) *Company {
	company := &Company{
		Name:          name,
		Website:       website,
		InterviewStep: 0,
	}

	CompanyList = append(CompanyList, company)
	return company
}

// Read All
func ListCompanies() string {
	var names []string

	for index, company := range CompanyList {
		names = append(names, fmt.Sprintf("%d: %s", index, company.Name))
	}

	return strings.Join(names, "\n")
}

// Read One
func ListCompany(companies string) string {
	return "Testing K.K."
}

// Delete
func DeleteCompany() {}
