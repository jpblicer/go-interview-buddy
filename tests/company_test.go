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
	addCompanyList()

	companies := model.ListCompanies()
	want := "0: Testing K.K.\n1: MegaCorp K.K."

	if companies != want {
		t.Errorf("Got %q, want %q", companies, want)
	}

}

func TestListCompany(t *testing.T) {
	t.Run("Should return a company based on index", func(t *testing.T) {
		resetCompanyList()
		addCompanyList()

		want := "Testing K.K.\nInterviews: 0\nWebsite: https://www.testcompany.com"
		got := model.ListCompany(0)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should return an error if the company id is not valid", func(t *testing.T) {
		resetCompanyList()
		addCompanyList()

		want := "Requested index is not valid"
		got := model.ListCompany(9001)

		if got != want {
			t.Errorf("got %q want%q", got, want)
		}
	})

}

func TestDeleteCompany(t *testing.T) {
	resetCompanyList()
	addCompanyList()

	want := "The element Testing K.K. was deleted."
	got := model.DeleteCompany(0)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func resetCompanyList() {
	model.CompanyList = make([]*model.Company, 0)
}

func addCompanyList() {
	model.AddCompany("Testing K.K.", "https://www.testcompany.com")
	model.AddCompany("MegaCorp K.K.", "https://www.corp.com")
}
