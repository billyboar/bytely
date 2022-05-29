package schema

import (
	"testing"
)

func TestValidationOK(t *testing.T) {
	u := URL{
		OriginalURL: "https://www.google.com",
	}

	err := u.Validate()
	if err != nil {
		t.Log("validation should not have errors")
		t.Fail()
	}
}

func TestValidationEmptyURL(t *testing.T) {
	u := URL{
		OriginalURL: "",
	}

	err := u.Validate()
	if err == nil {
		t.Log("Validation should have an error")
		t.FailNow()
	}

	if err != ErrOriginalURLMissing {
		t.Log("Validation should have `original_url is missing` error")
		t.Fail()
	}
}

func TestValidationInvalidURL(t *testing.T) {
	u := URL{
		OriginalURL: "https://hello:hello",
	}

	err := u.Validate()
	if err == nil {
		t.Log("Validation should have an error")
		t.FailNow()
	}
}
