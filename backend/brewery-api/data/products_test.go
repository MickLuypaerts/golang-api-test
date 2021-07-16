package data

import (
	"testing"
)

func TestChecksValidationWithValidData(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.00,
		SKU:   "abc-abc-3",
	}
	actual := p.Validate()
	if actual != nil {
		t.Errorf("Expected nil but got %v", actual)
	}
}

func TestChecksValidationWithInvalidData(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.00,
		SKU:   "abc-abc-3",
	}

	err := p.Validate()

	if err == nil {
		t.Errorf("Expected Fail got pass")
	}
}
