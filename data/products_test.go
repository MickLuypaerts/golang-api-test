package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.00,
		SKU:   "abc-abc-1",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}