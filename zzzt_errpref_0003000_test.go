package errpref

import "testing"

func TestErrPrefixDto_XCpy_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCpy("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: TestErrPrefixDto_XCpy_000100\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	return
}

func TestErrPrefixDto_XCpy_000200(t *testing.T) {

	ePrefix := ErrPrefixDto{}.NewEPrefCtx(
		"TestErrPrefixDto_XCpy_000200()",
		"")

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCpy("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: %v\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	xEPDto = xEPDto.XCpy("")

	actualStr = xEPDto.String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)

		return
	}

	return
}

func TestErrPrefixDto_XCpy_000300(t *testing.T) {

	ePrefix := ErrPrefixDto{}.NewEPrefCtx(
		"TestErrPrefixDto_XCpy_000300()",
		"")

	ePDto := ErrPrefixDto{}

	xEPDto := ePDto.XCpy("A!=B")

	expectedStr := ""

	actualStr := xEPDto.String()

	if expectedStr != actualStr {

		t.Errorf("Error: %v\n"+
			"Expected actualStr= ' ' (Empty String)\n"+
			"Instead, actualStr= '%v'\n",
			ePrefix.String(),
			actualStr)

		return
	}

	return
}

func TestErrPrefixDto_XCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtx("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_XCtx_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtx("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	xEPDto = xEPDto.XCtx("")

	actualStr = xEPDto.String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)
	}

}

func TestErrPrefixDto_XCtx_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}

	xEPDto := ePDto.XCtx("A!=B")

	expectedStr := ""

	actualStr := xEPDto.String()

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= ' ' (Empty String)\n"+
			"Instead, actualStr= '%v'\n",
			actualStr)
	}
}

func TestErrPrefixDto_XCtxEmpty_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtx("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	xEPDto = xEPDto.XCtxEmpty()

	actualStr = xEPDto.String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)
	}

}

func TestErrPrefixDto_XCtxEmpty_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtxEmpty()

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPrefixDto_XCtxEmpty_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtx("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	actualStr = xEPDto.XCtxEmpty().String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)
	}

}

func TestErrPrefixDto_XEPref_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XEPref("Tx5()")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if !ePDto.Equal(xEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_XEPrefCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()[SPACE]:[SPACE]X->G"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XEPrefCtx(
		"Tx5.BrandNewMethod()",
		"X->G")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_XEPrefOld_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	xEPDto := ePDto.XEPrefOld(initialStr)

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_XSetFromStrings_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_XSetFromStrings_000100()"

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	inputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	var outputDelimiters ErrPrefixDelimiters

	outputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n  &",
		" *&* ",
		"\n      %",
		" %%% ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetOutputStringDelimiters(
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	initialStr := ePDto.String()

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	ePDto.SetEPrefCtx(newErrPrefix, newErrCtx)

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.Ptr().
		XSetFromStrings(
			initialStr,
			newErrPrefix,
			newErrCtx,
			outputDelimiters,
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.Ptr().XSetFromStrings()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDto.String()),
			true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if expectedOutputStr != ePDto2Str {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputStr == ePDto1St\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto2Str = '%v'\n",
			expectedOutputStr,
			ePDto2Str)
		return
	}

	actualInputDelims,
		actualOutputDelims := ePDto2.GetStrDelimiters()

	if !actualInputDelims.Equal(&outputDelimiters) {
		t.Errorf("Error:\n"+
			"Original input delimiters not equal to "+
			"final input delimiters!\n"+
			"Original input delimiters='%v'\n"+
			"   Final input delimiters='%v'\n",
			outputDelimiters.String(),
			actualInputDelims.String())
	}

	if !actualOutputDelims.Equal(&outputDelimiters) {
		t.Errorf("Error:\n"+
			"Original output delimiters not equal to "+
			"final actual output delimiters!\n"+
			"Original output delimiters='%v'\n"+
			"   Final output delimiters='%v'\n",
			outputDelimiters.String(),
			actualOutputDelims.String())
	}

}

func TestErrPrefixDto_XSetFromStrings_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_XSetFromStrings_000200()"

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	inputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	var outputDelimiters ErrPrefixDelimiters

	outputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n  &",
		" *&* ",
		"\n      %",
		" %%% ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetOutputStringDelimiters(
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	initialStr := ePDto.String()

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	ePDto.SetEPrefCtx(newErrPrefix, newErrCtx)

	inputDelimiters.Empty()

	_,
		err = ErrPrefixDto{}.Ptr().
		XSetFromStrings(
			initialStr,
			newErrPrefix,
			newErrCtx,
			inputDelimiters,
			outputDelimiters,
			funcName)

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return from ErrPrefixDto{}.Ptr().XSetFromStrings()" +
			"because input parameter 'inputDelimiters' is empty and invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestErrPrefixDto_XSetFromStrings_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_XSetFromStrings_000300()"

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	inputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	var outputDelimiters ErrPrefixDelimiters

	outputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n  &",
		" *&* ",
		"\n      %",
		" %%% ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetOutputStringDelimiters(
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	initialStr := ePDto.String()

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	ePDto.SetEPrefCtx(newErrPrefix, newErrCtx)

	outputDelimiters.Empty()

	_,
		err = ErrPrefixDto{}.Ptr().
		XSetFromStrings(
			initialStr,
			newErrPrefix,
			newErrCtx,
			inputDelimiters,
			outputDelimiters,
			funcName)

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return from ErrPrefixDto{}.Ptr().XSetFromStrings()" +
			"because input parameter 'outputDelimiters' is empty and invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestErrPrefixDto_ZCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	zEPDto2 := ePDto.ZCtx("A!=B")

	actualStr := zEPDto2.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtx_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	zEPDto2 := ePDto.ZCtx("A!=B")

	actualStr := zEPDto2.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtx_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	ePDto2 := ePDto.ZCtx("A!=B")

	actualStr := ePDto2.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal ePDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtx_000400(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	ePDto2 := ePDto.ZCtx("A!=B")

	actualStr := ePDto2.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal ePDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

	ePDto3 := ePDto.ZCtx("")

	expectedStr = "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	actualStr = ePDto3.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error Series #3:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&ePDto3) {
		t.Error("Error Series #3:\n" +
			"Expected ePDto to Equal ePDto3.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtxEmpty_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	xEPDto := ePDto.XCtx("A!=B")

	actualStr := xEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	zEPDto2 := xEPDto.ZCtxEmpty()

	actualStr = zEPDto2.String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtxEmpty_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	ePDto.SetEPrefOld(initialStr)

	zEPDto2 := ePDto.ZCtxEmpty()

	actualStr := zEPDto2.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto2) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZCtxEmpty_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	zEPDto := ePDto.ZCtx("A!=B")

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	expectedStr2 := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	zEPDto2 := zEPDto.ZCtxEmpty()

	actualStr = zEPDto2.String()

	expectedStr2 = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr2),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr2 != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr2,
			actualStr)
	}

	if !zEPDto.Equal(&zEPDto2) {
		t.Error("Error:\n" +
			"Expected zEPDto to Equal zEPDto2.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZEPref_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()"

	ePDto.SetEPrefOld(initialStr)

	zEPDto := ePDto.ZEPref("Tx5()")

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZEPref_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4"

	ePDto.SetEPrefOld(initialStr)

	zEPDto := ePDto.ZEPref("")

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZEPrefCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()[SPACE]:[SPACE]X->G"

	ePDto.SetEPrefOld(initialStr)

	zEPDto := ePDto.ZEPrefCtx(
		"Tx5.BrandNewMethod()",
		"X->G")

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}
}

func TestErrPrefixDto_ZEPrefCtx_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()"

	ePDto.SetEPrefOld(initialStr)

	zEPDto := ePDto.ZEPrefCtx(
		"Tx5.BrandNewMethod()",
		"")

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZEPrefOld_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	zEPDto := ePDto.ZEPrefOld(initialStr)

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}

func TestErrPrefixDto_ZEPrefOld_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr := ""

	expectedStr := ""

	zEPDto := ePDto.ZEPrefOld(initialStr)

	actualStr := zEPDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto.Equal(&zEPDto) {
		t.Error("Error:\n" +
			"Expected ePDto to Equal zEPDto.\n" +
			"However, THEY ARE NOT EQUAL!!!\n")
	}

}
