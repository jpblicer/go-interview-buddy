package company

import "testing"

func TestAdd(t *testing.T) {
	companies = []Company{}

	company := Company{Name: "MegaCorp K.K."}

	got := Add("MegaCorp K.K.")
	want := company

	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}

}

func TestList(t *testing.T) {
	companies = []Company{}

	companyA := Add("MegaCorp K.K.")
	companyB := Add("Acme Corporation")

	companies := []Company{companyA, companyB}

	got := len(List())
	want := len(companies)

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}

}
