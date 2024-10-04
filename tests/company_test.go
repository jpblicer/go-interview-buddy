package tests

import (
	"testing"

	"github.com/jpblicer/go-interview-buddy/model"
)

func TestAddCompany(t *testing.T) {
	resetCompanyList()
	name := "Testing K.K."
	website := "https://www.testcompany.com"

	company := model.AddCompany(name, website)

	if company.Name != name {
		t.Errorf("Got %q but Want %q", company.Name, name)
	}
	if company.Website != website {
		t.Errorf("Got %q but Want %q", company.Website, website)
	}
	if company.InterviewStep != 0 {
		t.Errorf("Got %d but Want %d", company.InterviewStep, 0)
	}
}

func TestListCompanies(t *testing.T) {
	resetCompanyList()
	model.AddCompany("Testing K.K.", "https://www.testcompany.com")
	model.AddCompany("MegaCorp K.K.", "https://www.corp.com")

	companies := model.ListCompanies()
	want := "0: Testing K.K.\n1: MegaCorp K.K."

	if companies != want {
		t.Errorf("Got %q, want %q", companies, want)
	}

}

func TestListCompany(t *testing.T) {
	resetCompanyList()
	model.AddCompany("Testing K.K.", "https://www.testcompany.com")
	model.AddCompany("MegaCorp K.K.", "https://www.corp.com")

	want := "Testing K.K./nInterviews : 0/nWebsite: https://www.corp.com"
	got := model.ListCompany(0)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func resetCompanyList() {
	model.CompanyList = make([]*model.Company, 0)
}
