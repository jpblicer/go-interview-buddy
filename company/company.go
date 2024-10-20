package company

type Company struct {
	Name       string
	Applied    bool
	Interviews int
}

var companies = []Company{}

func Add(name string) Company {
	company := Company{Name: name}
	companies = append(companies, company)
	return company
}

func List() []Company {
	return companies
}
