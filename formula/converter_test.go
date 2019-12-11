package formula

import (
	"fmt"
	"os"
	"testing"
)

func TestConvImply(t *testing.T) {

	var f string
	var r string
	var err error

	// Convert A to A
	f = "A"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "A" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert A&A to (A&A)
	f = "A&A"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(A&A)" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert A>B to (~A|B)
	f = "A>B"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|B)" {
		t.Errorf("(Convert A>B to (~A|B)): Failed. The result is %v", r)
	}

	// Convert A>(A>B) to (~A|(~A|B))
	f = "A>(A>B)"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(~A|B))" {
		t.Errorf("(Convert A>(A>B) to (~A|(~A|B)): Failed. The result is %v", r)
	}

	// Convert (A>(A>B)) to (~A|(~A|B))
	f = "(A>(A>B))"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(~A|B))" {
		t.Errorf("(Convert (A>(A>B)) to (~A|(~A|B)): Failed. The result is %v", r)
	}

	// Convert A&B&C to (A&B&C)
	f = "A&B&C"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(A&B&C)" {
		t.Errorf("(Convert A&B&C to A&B&C: Failed. The result is %v", r)
	}

	// Convert (A&B&C|D&E&F) to ((A&B&C)|(D&E&F))
	f = "(A&B&C)|(D&E&F)"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "((A&B&C)|(D&E&F))" {
		t.Errorf("(Convert (A&B&C|D&E&F) to ((A&B&C)|(D&E&F)): Failed. The result is %v", r)
	}

}

func TestConvNeg(t *testing.T) {

	var f string
	var r string
	var err error

	// Convert A to A
	f = "A"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "A" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert A&A to (A&A)
	f = "A&A"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(A&A)" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert ~(A&B&C) to (~A|~B|~C)
	f = "~(A&B&C)"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|~B|~C)" {
		t.Errorf("(Convert ~(A&B&C) to (~A|~B|~C): Failed. The result is %v", r)
	}

	// Convert ~(~A&B&C) to (A|~B|~C)
	f = "~(~A&B&C)"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(A|~B|~C)" {
		t.Errorf("(Convert ~(~A&B&C) to (A|~B|~C): Failed. The result is %v", r)
	}

	// Convert ~(~A&~(~B&C)) to (A|~B|~C)
	f = "~(~A&~(~B&C))"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(A|(~B&C))" {
		t.Errorf("(Convert ~(~A&~(~B&C)) to (A|(~B&C)): Failed. The result is %v", r)
	}

	// Convert ~(A&(B|C)) to (~A|(~B&~C))
	f = "~(A&(B|C))"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(~B&~C))" {
		t.Errorf("(Convert ~(A&(B|C)) to (~A|(~B&~C)): Failed. The result is %v", r)
	}

	// Convert ~(A&B&C)|~(A&B&C) to ((~A|~B|~C)|(~A|~B|~C))
	f = "~(A&B&C)|~(A&B&C)"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "((~A|~B|~C)|(~A|~B|~C))" {
		t.Errorf("(Convert ~(A&B&C)|~(A&B&C) to ((~A|~B|~C)|(~A|~B|~C)): Failed. The result is %v", r)
	}

	// Convert ~~A to A
	f = "~~A"

	r, err = ConvNeg(f)
	if err != nil {
		printError(err)
	}

	if r != "A" {
		t.Errorf("(Convert ~~A to A: Failed. The result is %v", r)
	}

}

// printErrorはエラーメッセージ出力を統一する.
func printError(err /* error */ error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}
