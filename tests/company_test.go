package tests

import (
	"testing"

	"github.com/jpblicer/go-interview-buddy/model"
)

func TestAddCompany(t *testing.T) {
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
	model.AddCompany("Testing K.K.", "https://www.testcompany.com")
	model.AddCompany("MegaCorp K.K.", "https://www.corp.com")

	companies := model.ListCompanies()
	want := `Testing K.K.\nMegaCorp K.K.`

	if companies != want {
		t.Errorf("Got %q, want %q", companies, want)
	}

}

func TestListCompany(t *testing.T) {
	company := "Testing K.K."

	want := "Testing K.K."
	got := model.ListCompany(company)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
