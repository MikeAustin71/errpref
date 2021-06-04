package errpref

import (
	"fmt"
	"strings"
	"testing"
)

type testFuncAlpha01 struct {
	input string
}

func (tFuncAlpha01 *testFuncAlpha01) Tx1DoSomething(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha01.Tx1DoSomething()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha2 := testFuncAlpha02{}

	return tFuncAlpha2.Tx2DoSomethingElse(
		ePrefix)
}

type testFuncAlpha02 struct {
	input string
}

func (tFuncAlpha02 *testFuncAlpha02) Tx2DoSomethingElse(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha02.Tx2DoSomethingElse()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha03 := testFuncAlpha03{}

	err = tFuncAlpha03.Tx3DoAnything(
		ePrefix)

	return err
}

type testFuncAlpha03 struct {
	input string
}

func (tFuncAlpha03 *testFuncAlpha03) Tx3DoAnything(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha03.Tx3DoAnything()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha04 := testFuncAlpha04{}

	err = tFuncAlpha04.Tx4DoNothing(
		ePrefix)

	return err
}

type testFuncAlpha04 struct {
	input string
}

func (tFuncAlpha04 *testFuncAlpha04) Tx4DoNothing(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha04.Tx4DoNothing()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha05 := testFuncAlpha05{}

	err = tFuncAlpha05.Tx5DoSomethingBig(
		ePrefix.XCtx(
			"A/B==4"))

	return err
}

type testFuncAlpha05 struct {
	input string
}

func (tFuncAlpha05 *testFuncAlpha05) Tx5DoSomethingBig(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha05.Tx5DoSomethingBig()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha06 := testFuncAlpha06{}

	err = tFuncAlpha06.Tx6GiveUp(
		ePrefix.XCtx(
			"A->B"))

	return err
}

type testFuncAlpha06 struct {
	input string
}

func (tFuncAlpha06 *testFuncAlpha06) Tx6GiveUp(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncAlpha06.Tx6GiveUp()",
		"")

	if err != nil {
		return err
	}

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}

type testFuncBravo01 struct {
	input string
}

func (tFuncBravo01 *testFuncBravo01) Tx1DoSomethingSpecial(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncBravo01.Tx1DoSomethingSpecial()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo02 := testFuncBravo02{}

	return tFuncBravo02.Tx2DoSomethingElse(ePrefix)
}

type testFuncBravo02 struct {
	input string
}

func (tFuncBravo02 *testFuncBravo02) Tx2DoSomethingElse(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncBravo02.Tx2DoSomethingElse()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo03 := testFuncBravo03{}

	err = tFuncBravo03.Tx3DoAnything(ePrefix)

	return err
}

type testFuncBravo03 struct {
	input string
}

func (tFuncBravo03 *testFuncBravo03) Tx3DoAnything(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncBravo03.Tx3DoAnything()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo04 := testFuncBravo04{}

	err = tFuncBravo04.Tx4DoNothing(ePrefix)

	return err
}

type testFuncBravo04 struct {
	input string
}

func (tFuncBravo04 *testFuncBravo04) Tx4DoNothing(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncBravo04.Tx4DoNothing()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo05 := testFuncBravo05{}

	ePrefix.SetCtx("A/B==4")

	err = tFuncBravo05.Tx5DoSomethingBig(ePrefix)

	return err
}

type testFuncBravo05 struct {
	input string
}

func (tFuncBravo05 *testFuncBravo05) Tx5DoSomethingBig(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncBravo05.Tx5DoSomethingBig()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo06 := testFuncBravo06{}

	ePrefix.SetCtx("A->B")

	err = tFuncBravo06.Tx6GiveUp(ePrefix)

	return err
}

type testFuncBravo06 struct {
	input string
}

func (tFuncY6 *testFuncBravo06) Tx6GiveUp(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"Tx6GiveUp()",
		"")

	if err != nil {
		return err
	}

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}

type testFuncCharlie01 struct {
	input string
}

func (tFuncCharlie01 *testFuncCharlie01) Tx1DoStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie01.Tx1DoStuff()",
		"X->Y")

	if err != nil {
		return err
	}

	tFuncCharlie02 := testFuncCharlie02{}

	return tFuncCharlie02.Tx2DoMoreStuff(
		ePrefix)
}

type testFuncCharlie02 struct {
	input string
}

func (tFuncCharlie02 *testFuncCharlie02) Tx2DoMoreStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie02.Tx2DoMoreStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie03 := testFuncCharlie03{}

	return tFuncCharlie03.Tx3DoLessStuff(
		ePrefix.XCtx(
			"B->C"))
}

type testFuncCharlie03 struct {
	input string
}

func (tFuncCharlie03 *testFuncCharlie03) Tx3DoLessStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie03.Tx3DoLessStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie04 := testFuncCharlie04{}

	return tFuncCharlie04.Tx4DoFunStuff(
		ePrefix)
}

type testFuncCharlie04 struct {
	input string
}

func (tFuncCharlie04 *testFuncCharlie04) Tx4DoFunStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie04.Tx4DoFunStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie05 := testFuncCharlie05{}

	return tFuncCharlie05.Tx5DoExcitingStuff(
		ePrefix)

}

type testFuncCharlie05 struct {
	input string
}

func (tFuncCharlie05 *testFuncCharlie05) Tx5DoExcitingStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie05.Tx5DoExcitingStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie06 := testFuncCharlie06{}

	return tFuncCharlie06.Tx6DoSpaceStuff(
		ePrefix.XCtx(
			"X*Y"))

}

type testFuncCharlie06 struct {
	input string
}

func (tFuncCharlie06 *testFuncCharlie06) Tx6DoSpaceStuff(
	errPrefix *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"testFuncCharlie06.Tx6DoSpaceStuff()",
		"")

	if err != nil {
		return err
	}

	ePrefix.SetCtx(
		"Asteroid Collision!")

	return fmt.Errorf("%v\n",
		ePrefix.String())
}

func TestErrPrefixDto_CallSeries_000100(t *testing.T) {

	funcName := "StartingMethod()"
	twoDAry := make([][2]string, 7)

	twoDAry[0][0] = funcName
	twoDAry[0][1] = ""

	twoDAry[1][0] = "testFuncCharlie01.Tx1DoStuff()"
	twoDAry[1][1] = "X->Y"

	twoDAry[2][0] = "testFuncCharlie02.Tx2DoMoreStuff()"
	twoDAry[2][1] = "B->C"

	twoDAry[3][0] = "testFuncCharlie03.Tx3DoLessStuff()"
	twoDAry[3][1] = ""

	twoDAry[4][0] = "testFuncCharlie04.Tx4DoFunStuff()"
	twoDAry[4][1] = ""

	twoDAry[5][0] = "testFuncCharlie05.Tx5DoExcitingStuff()"
	twoDAry[5][1] = "X*Y"

	twoDAry[6][0] = "testFuncCharlie06.Tx6DoSpaceStuff()"
	twoDAry[6][1] = "Asteroid Collision!"

	ePDtoBaseLine,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDAry,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoBaseLine.SetMaxTextLineLen(40)

	ePDto1 := ErrPrefixDto{}.Ptr()

	ePDto1.SetEPref(funcName)

	ePDto1.SetMaxTextLineLen(40)

	tFuncAlpha01 := testFuncAlpha01{}

	err =
		tFuncAlpha01.Tx1DoSomething(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err =
		tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncCharlie01 := testFuncCharlie01{}

	err2 :=
		tFuncCharlie01.Tx1DoStuff(ePDto1)

	if err2 == nil {
		t.Error("Expected error return from\n" +
			"tFuncCharlie01.Tx1DoStuff(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")

		return
	}

	if ePDto1.String() != funcName {
		t.Errorf("Error: Expected ePDto1.String() != funcName\n"+
			"Instead:\n"+
			"ePDto1.String()='%v'\n"+
			"funcName='%v'\n\n",
			ePDto1.String(),
			funcName)
		return
	}

	var ePDtoFinal *ErrPrefixDto

	ePDtoFinal,
		err = ErrPrefixDto{}.NewIEmpty(
		err2.Error(),
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(err.Error())\n"+
			"%v\n", err.Error())

		return
	}

	ePDtoFinal.SetMaxTextLineLen(40)

	ePDtoBaseLineStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoBaseLine.String()),
		false)

	ePDtoFinalStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoFinal.String()),
		false)

	if ePDtoBaseLineStr != ePDtoFinalStr {
		t.Errorf("Error: ePDtoBaseLineStr != ePDtoFinalStr\n"+
			"ePDtoBaseLineStr=\n%v\n\n"+
			"ePDtoFinalStr=\n%v\n\n",
			ePDtoBaseLineStr,
			ePDtoFinalStr)
	}

}

func TestErrPrefixDto_CallSeries_000200(t *testing.T) {

	funcName := "StartingMethod()"

	ePDto1 := new(ErrPrefixDto)

	ePDto1.SetEPref(funcName)

	ePDto1.SetMaxTextLineLen(40)

	turnOffTextDisplay := ePDto1.GetTurnOffTextDisplay()

	if turnOffTextDisplay == true {

		t.Error("ERROR:\n" +
			"Expected turnOffTextDisplay == false\n" +
			"Instead, turnOffTextDisplay == true\n")

		return

	}

	tFuncAlpha01 := testFuncAlpha01{}

	err :=
		tFuncAlpha01.Tx1DoSomething(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err =
		tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncCharlie01 := testFuncCharlie01{}

	err2 :=
		tFuncCharlie01.Tx1DoStuff(ePDto1)

	if err2 == nil {
		t.Error("Expected error return from\n" +
			"tFuncCharlie01.Tx1DoStuff(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")

		return
	}

	turnOffTextDisplay = ePDto1.GetTurnOffTextDisplay()

	if turnOffTextDisplay == true {

		t.Error("ERROR:\n" +
			"Expected turnOffTextDisplay == false\n" +
			"Instead, turnOffTextDisplay == true\n")

		return

	}

	ePDtoErrStr := err2.Error()

	if !strings.Contains(ePDtoErrStr, funcName) {
		t.Error("ERROR: turnOffTextDisplay == false\n" +
			"HOWEVER, NO ERROR PREFIX INFO WAS RETURNED!\n")
	}

}

func TestErrPrefixDto_CallSeries_000300(t *testing.T) {

	funcName := "StartingMethod()"

	ePDto1 := new(ErrPrefixDto)

	ePDto1.SetEPref(funcName)

	ePDto1.SetMaxTextLineLen(40)

	turnOffTextDisplay := ePDto1.GetTurnOffTextDisplay()

	if turnOffTextDisplay == true {

		t.Error("ERROR:\n" +
			"Expected turnOffTextDisplay == false\n" +
			"Instead, turnOffTextDisplay == true\n")

		return

	}

	tFuncAlpha01 := testFuncAlpha01{}

	err :=
		tFuncAlpha01.Tx1DoSomething(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err =
		tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	ePDto1.SetTurnOffTextDisplay(true)

	turnOffTextDisplay = ePDto1.GetTurnOffTextDisplay()

	if turnOffTextDisplay == false {

		t.Error("ERROR:\n" +
			"Attempted to turn off text display. Attempt FAILED!\n" +
			"Expected turnOffTextDisplay == true\n" +
			"Instead, turnOffTextDisplay == false\n")

		return

	}

	tFuncCharlie01 := testFuncCharlie01{}

	err2 :=
		tFuncCharlie01.Tx1DoStuff(ePDto1)

	if err2 == nil {
		t.Error("Expected error return from\n" +
			"tFuncCharlie01.Tx1DoStuff(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")

		return
	}

	turnOffTextDisplay = ePDto1.GetTurnOffTextDisplay()

	if turnOffTextDisplay == false {

		t.Error("ERROR: turnOffTextDisplay verification.\n" +
			"Expected turnOffTextDisplay == true\n" +
			"Instead, turnOffTextDisplay == false\n")

		return

	}

	ePDtoErrStr := err2.Error()

	if strings.Contains(ePDtoErrStr, funcName) {
		t.Error("ERROR: turnOffTextDisplay == true\n" +
			"HOWEVER, ERROR PREFIX INFO WAS RETURNED!\n")
	}

}
