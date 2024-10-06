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

func AddCompany(name, website string) *Company {
	company := &Company{
		Name:          name,
		Website:       website,
		InterviewStep: 0,
	}

	CompanyList = append(CompanyList, company)
	return company
}

func ListCompanies() string {
	var names []string

	for index, company := range CompanyList {
		names = append(names, fmt.Sprintf("%d: %s", index, company.Name))
	}

	return strings.Join(names, "\n")
}

func ListCompany(index int) string {

	if index < 0 || index >= len(CompanyList) {
		return "Requested index is not valid"
	}

	company := CompanyList[index]

	companyDetails := fmt.Sprintf("%s\nInterviews: %d\nWebsite: %s", company.Name, company.InterviewStep, company.Website)

	return companyDetails
}

func DeleteCompany(index int) string {
	company := CompanyList[index]

	CompanyList = append(CompanyList[:index], CompanyList[index+1:]...)

	confirm := fmt.Sprintf("The element %s was deleted.", company.Name)

	return confirm
}
