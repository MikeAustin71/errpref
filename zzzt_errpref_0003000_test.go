package errpref

import "testing"

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
