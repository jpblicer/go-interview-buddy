package company

import "testing"

func TestAdd(t *testing.T) {
	company := Company{Name: "MegaCorp K.K."}

	got := Add("MegaCorp K.K.")
	want := company.Name

	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}

}
