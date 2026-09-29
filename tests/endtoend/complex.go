package main

import "log"

type Complex64Validation struct {
	Required complex64 `valid:"required"`
	Eq       complex64 `valid:"eq=1+2i"`
	Neq      complex64 `valid:"neq=3+4i"`
	In       complex64 `valid:"in=1+2i 5+6i"`
	Nin      complex64 `valid:"nin=7+8i 9+0i"`
}

type Complex128Validation struct {
	Required complex128 `valid:"required"`
	Eq       complex128 `valid:"eq=1+2i"`
	Neq      complex128 `valid:"neq=3+4i"`
	In       complex128 `valid:"in=1+2i 5+6i"`
	Nin      complex128 `valid:"nin=7+8i 9+0i"`
}

func complexTests() {
	log.Println("starting complex64 tests")
	complex64Tests()
	log.Println("complex64 tests ok")

	log.Println("starting complex128 tests")
	complex128Tests()
	log.Println("complex128 tests ok")
}

func complex64Tests() {
	invalid := &Complex64Validation{
		Required: 0,
		Eq:       1 + 3i,
		Neq:      3 + 4i,
		In:       1 + 6i,
		Nin:      7 + 8i,
	}
	wantInvalid := []string{
		"Required is required",
		"Eq must be equal to 1+2i",
		"Neq must not be equal to 3+4i",
		"In must be one of '1+2i' '5+6i'",
		"Nin must not be one of '7+8i' '9+0i'",
	}
	errs := Complex64ValidationValidate(invalid)
	if !expectedMsgErrorsOk(errs, wantInvalid) {
		log.Fatalf("error = %v, wantErr %v", errs, wantInvalid)
	}

	// Nonzero with a zero real part, so required is not a real-part check.
	valid := &Complex64Validation{
		Required: 0 + 1i,
		Eq:       1 + 2i,
		Neq:      3 + 5i,
		In:       5 + 6i,
		Nin:      1 + 0i,
	}
	errs = Complex64ValidationValidate(valid)
	if !expectedMsgErrorsOk(errs, nil) {
		log.Fatalf("error = %v, wantErr %v", errs, nil)
	}
}

func complex128Tests() {
	invalid := &Complex128Validation{
		Required: 0,
		Eq:       9 + 2i,
		Neq:      3 + 4i,
		In:       0,
		Nin:      9 + 0i,
	}
	wantInvalid := []string{
		"Required is required",
		"Eq must be equal to 1+2i",
		"Neq must not be equal to 3+4i",
		"In must be one of '1+2i' '5+6i'",
		"Nin must not be one of '7+8i' '9+0i'",
	}
	errs := Complex128ValidationValidate(invalid)
	if !expectedMsgErrorsOk(errs, wantInvalid) {
		log.Fatalf("error = %v, wantErr %v", errs, wantInvalid)
	}

	// Nonzero with a zero imaginary part.
	valid := &Complex128Validation{
		Required: 1 + 0i,
		Eq:       1 + 2i,
		Neq:      0,
		In:       1 + 2i,
		Nin:      2 + 2i,
	}
	errs = Complex128ValidationValidate(valid)
	if !expectedMsgErrorsOk(errs, nil) {
		log.Fatalf("error = %v, wantErr %v", errs, nil)
	}
}
