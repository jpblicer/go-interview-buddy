package model

type Company struct {
	Name          string
	Website       string
	InterviewStep int
}

// Create
func AddCompany(name, website string) *Company {
	return &Company{
		Name:          name,
		Website:       website,
		InterviewStep: 0,
	}
}

// Read All
func ListCompanies(companies string) string {
	return "Testing K.K."
}

// Read One

// Delete
