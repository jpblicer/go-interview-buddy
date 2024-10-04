package model

type Company struct {
	Name          string
	Website       string
	InterviewStep int
}

var CompanyList = make([]*Company, 0)

// Create
func AddCompany(name, website string) *Company {
	return &Company{
		Name:          name,
		Website:       website,
		InterviewStep: 0,
	}
}

// Read All
func ListCompanies(){}

// Read One
func ListCompany(companies string) string {
	return "Testing K.K."
}

// Delete
func DeleteCompany(){}
