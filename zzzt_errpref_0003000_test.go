package errpref

import (
	"fmt"
	"testing"
)

func TestErrPrefixDto_XMethods_000100(t *testing.T) {

	tFuncX1 := testFuncX1{}

	err := tFuncX1.Tx1DoSomething()

	if err == nil {
		t.Error("ERROR: Expected an error return from\n" +
			"tFuncB1.Tx1DoSomething(). However, no error was returned!\n")
		return
	}

	expectedStr :=
		"Tx1DoSomething()[SPACE]-[SPACE]Tx2DoSomethingElse()\\n" +
			"Tx3DoAnything()\\n" +
			"Tx4DoNothing()[SPACE]:[SPACE]A/B==4\\n" +
			"Tx5DoSomethingBig()[SPACE]:[SPACE]A->B\\n" +
			"Tx6GiveUp()[SPACE]:[SPACE]A/B[SPACE]=[SPACE]C[SPACE]B==0\\n" +
			"Error:[SPACE]Divide[SPACE]By[SPACE]Zero!\\n"

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(err.Error()),
		true)

	if expectedStr != actualStr {
		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr, actualStr)
	}

}

func TestErrPrefixDto_YMethods_000200(t *testing.T) {

	tFuncY1 := testFuncY1{}

	err := tFuncY1.Tx1DoSomething()

	if err == nil {
		t.Error("ERROR: Expected an error return from\n" +
			"tFuncB1.Tx1DoSomething(). However, no error was returned!\n")
		return
	}

	expectedStr :=
		"Tx1DoSomething()[SPACE]-[SPACE]Tx2DoSomethingElse()\\n" +
			"Tx3DoAnything()\\n" +
			"Tx4DoNothing()[SPACE]:[SPACE]A/B==4\\n" +
			"Tx5DoSomethingBig()[SPACE]:[SPACE]A->B\\n" +
			"Tx6GiveUp()[SPACE]:[SPACE]A/B[SPACE]=[SPACE]C[SPACE]B==0\\n" +
			"Error:[SPACE]Divide[SPACE]By[SPACE]Zero!\\n"

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(err.Error()),
		true)

	if expectedStr != actualStr {
		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr, actualStr)
	}
}

type testFuncX1 struct {
	input string
}

func (tFuncX1 *testFuncX1) Tx1DoSomething() error {

	tFuncX2 := testFuncX2{}

	ePrefix := ErrPrefixDto{}

	err := tFuncX2.Tx2DoSomethingElse(
		ePrefix.XEPrefOld("Tx1DoSomething()"))

	return err
}

type testFuncX2 struct {
	input string
}

func (tFuncX2 *testFuncX2) Tx2DoSomethingElse(
	ePrefix *ErrPrefixDto) error {

	tFuncX3 := testFuncX3{}

	err := tFuncX3.Tx3DoAnything(
		ePrefix.XEPref("Tx2DoSomethingElse()"))

	return err
}

type testFuncX3 struct {
	input string
}

func (tFuncX3 *testFuncX3) Tx3DoAnything(
	ePrefix *ErrPrefixDto) error {

	tFuncX4 := testFuncX4{}

	err := tFuncX4.Tx4DoNothing(
		ePrefix.XEPref("Tx3DoAnything()"))

	return err
}

type testFuncX4 struct {
	input string
}

func (tFuncX4 *testFuncX4) Tx4DoNothing(
	ePrefix *ErrPrefixDto) error {

	tFuncX5 := testFuncX5{}

	err := tFuncX5.Tx5DoSomethingBig(
		ePrefix.XEPrefCtx(
			"Tx4DoNothing()",
			"A/B==4"))

	return err
}

type testFuncX5 struct {
	input string
}

func (tFuncX5 *testFuncX5) Tx5DoSomethingBig(
	ePrefix *ErrPrefixDto) error {

	tFuncX6 := testFuncX6{}

	err := tFuncX6.Tx6GiveUp(
		ePrefix.XEPrefCtx(
			"Tx5DoSomethingBig()",
			"A->B"))

	return err
}

type testFuncX6 struct {
	input string
}

func (tFuncX6 *testFuncX6) Tx6GiveUp(
	ePrefix *ErrPrefixDto) error {

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	var err = fmt.Errorf("%v\n"+
		"Error: Divide By Zero!\n",
		ePrefix.String())

	return err
}

type testFuncY1 struct {
	input string
}

func (tFuncY1 *testFuncY1) Tx1DoSomething() error {

	tFuncY2 := testFuncY2{}

	ePrefix := ErrPrefixDto{}

	ePrefix.SetEPrefOld("Tx1DoSomething()")

	err := tFuncY2.Tx2DoSomethingElse(ePrefix)

	return err
}

type testFuncY2 struct {
	input string
}

func (tFuncY2 *testFuncY2) Tx2DoSomethingElse(
	ePrefix ErrPrefixDto) error {

	ePrefix.SetEPref("Tx2DoSomethingElse()")

	tFuncY3 := testFuncY3{}

	err := tFuncY3.Tx3DoAnything(ePrefix)

	return err
}

type testFuncY3 struct {
	input string
}

func (tFuncY3 *testFuncY3) Tx3DoAnything(
	ePrefix ErrPrefixDto) error {

	ePrefix.SetEPref("Tx3DoAnything()")

	tFuncY4 := testFuncY4{}

	err := tFuncY4.Tx4DoNothing(ePrefix)

	return err
}

type testFuncY4 struct {
	input string
}

func (tFuncY4 *testFuncY4) Tx4DoNothing(
	ePrefix ErrPrefixDto) error {

	ePrefix.SetEPref("Tx4DoNothing()")

	tFuncY5 := testFuncY5{}

	ePrefix.SetCtx("A/B==4")

	err := tFuncY5.Tx5DoSomethingBig(ePrefix)

	return err
}

type testFuncY5 struct {
	input string
}

func (tFuncY5 *testFuncY5) Tx5DoSomethingBig(
	ePrefix ErrPrefixDto) error {

	ePrefix.SetEPref("Tx5DoSomethingBig()")

	tFuncY6 := testFuncY6{}

	ePrefix.SetCtx("A->B")

	err := tFuncY6.Tx6GiveUp(ePrefix)

	return err
}

type testFuncY6 struct {
	input string
}

func (tFuncY6 *testFuncY6) Tx6GiveUp(
	ePrefix ErrPrefixDto) error {

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	var err = fmt.Errorf("%s\n"+
		"Error: Divide By Zero!\n",
		ePrefix)

	return err
}
