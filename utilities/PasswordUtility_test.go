package utilities

import "testing"

func TestHash(t *testing.T) {
	input := "My fancy text"

	hashedInput, err := HashInput(input)

	if err != nil {
		t.Fatalf("Hash test failed with error: %v", err)
	}

	compareCheck, err := DoesHashedInputCompare(hashedInput, input)

	if err != nil {
		t.Fatalf("Hash test failed with error: %v", err)
	} else if !compareCheck {
		t.Fatal("Hash compare failed.")
	}
}
