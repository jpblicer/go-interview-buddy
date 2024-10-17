package company

import "testing"

func TestAdd(t *testing.T) {
	company := Company{Name: "MegaCorp K.K."}

	got := Add(company)
	want := company

	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}

}
