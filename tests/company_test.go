package tests

import (
	"testing"

	"github.com/jpblicer/go-interview-buddy/model"
)

func TestAddCompany(t *testing.T) {
	name := "Testing K.K."
	website := "https://testcompany.com"

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
	company := "Testing K.K."

	want := "Testing K.K."
	got := model.ListCompanies(company)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
