package tests

import (
	"testing"

	"github.com/jpblicer/go-interview-buddy/model"
)

func TestListCompanies(t *testing.T) {
	company := "Testing K.K."

	want := "Testing K.K."
	got := model.ListCompanies(company)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
