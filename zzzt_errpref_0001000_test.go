package errpref

import (
	"testing"
)

func TestErrPrefixDto_AddEPrefStrings_000100(t *testing.T) {

	var twoDSlice01 [][2]string

	twoDSlice01 = make([][2]string, 14)

	twoDSlice01[0][0] = "Tx1.Something()"
	twoDSlice01[0][1] = ""

	twoDSlice01[1][0] = "Tx2.SomethingElse()"
	twoDSlice01[1][1] = ""

	twoDSlice01[2][0] = "Tx3.DoSomething()"
	twoDSlice01[2][1] = ""

	twoDSlice01[3][0] = "Tx4()"
	twoDSlice01[3][1] = ""

	twoDSlice01[4][0] = "Tx5()"
	twoDSlice01[4][1] = ""

	twoDSlice01[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice01[5][1] = ""

	twoDSlice01[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice01[6][1] = "something->newSomething"

	twoDSlice01[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice01[7][1] = ""

	twoDSlice01[8][0] = "Tx9.TryAHammer()"
	twoDSlice01[8][1] = "x->y"

	twoDSlice01[9][0] = "Tx10.X()"
	twoDSlice01[9][1] = ""

	twoDSlice01[10][0] = "Tx11.TryAnything()"
	twoDSlice01[10][1] = ""

	twoDSlice01[11][0] = "Tx12.TryASalad()"
	twoDSlice01[11][1] = ""

	twoDSlice01[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice01[12][1] = ""

	twoDSlice01[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice01[13][1] = "A=7 B=8 C=9"

	var twoDSlice02 [][2]string

	twoDSlice02 = make([][2]string, 14)

	twoDSlice02[0][0] = "Tx1.TryBeer()"
	twoDSlice02[0][1] = ""

	twoDSlice02[1][0] = "Tx2.TryMescal()"
	twoDSlice02[1][1] = ""

	twoDSlice02[2][0] = "Tx3.TryWhiskey()"
	twoDSlice02[2][1] = ""

	twoDSlice02[3][0] = "Tx4.TryMoonShine()"
	twoDSlice02[3][1] = ""

	twoDSlice02[4][0] = "Tx5.TryScotch()"
	twoDSlice02[4][1] = ""

	twoDSlice02[5][0] = "Tx6.TryWine()"
	twoDSlice02[5][1] = ""

	twoDSlice02[6][0] = "Tx7.TryBrandy()"
	twoDSlice02[6][1] = "something->newSomething"

	twoDSlice02[7][0] = "Tx8.TryCognac()"
	twoDSlice02[7][1] = ""

	twoDSlice02[8][0] = "Tx9.TryAle()"
	twoDSlice02[8][1] = "x->y"

	twoDSlice02[9][0] = "Tx10.Vodka()"
	twoDSlice02[9][1] = ""

	twoDSlice02[10][0] = "Tx11.TryRum()"
	twoDSlice02[10][1] = ""

	twoDSlice02[11][0] = "Tx12.TryBourbon()"
	twoDSlice02[11][1] = ""

	twoDSlice02[12][0] = "Tx13.TryTequila()"
	twoDSlice02[12][1] = ""

	twoDSlice02[13][0] = "Tx14.TryWater"
	twoDSlice02[13][1] = "A=7 B=8 C=9"

	var ePDto1, ePDto3 *ErrPrefixDto
	var err error

	ePDto1,
		err = ErrPrefixDto{}.NewIEmpty(
		twoDSlice01,
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice01)\n"+
			"%v\n", err.Error())

		return
	}

	ePDto1.AddEPrefStrings(twoDSlice02)

	ePDto1.SetMaxTextLineLen(40)

	var twoDSlice03 [][2]string

	twoDSlice03 = make([][2]string,
		len(twoDSlice01)+len(twoDSlice02))

	copy(twoDSlice03[:len(twoDSlice01)],
		twoDSlice01[:])

	copy(twoDSlice03[len(twoDSlice01):],
		twoDSlice02[:])

	ePDto3,
		err = ErrPrefixDto{}.NewIEmpty(
		twoDSlice03,
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice03)\n"+
			"%v\n", err.Error())

		return
	}

	ePDto3.SetMaxTextLineLen(40)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto3.String()),
		true)

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		true)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected expectedStr == actualStr\n"+
			"Instead:\n"+
			"expectedStr=\n%v\n\n"+
			"actualStr=\n%v\n\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto1.Equal(ePDto3) {
		t.Errorf("Error: Expected ePDto1 == ePDto3\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"ePDto1.String()=\n%v\n\n"+
			"ePDto3.String()=\n%v\n\n",
			ePDto1.String(),
			ePDto3.String())
	}

}

func TestErrPref_ConvertNonPrintableChars_000100(t *testing.T) {

	tRunes := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	expectedStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	printableChars :=
		ErrPref{}.ConvertNonPrintableChars(
			tRunes,
			true)

	if printableChars != expectedStr {
		t.Errorf("ERROR:\n"+
			"Expected printableChars == expectedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"printableChars='%v'\n"+
			"expectedStr='%v'\n",
			printableChars,
			expectedStr)
	}

}

func TestErrPref_ConvertPrintableChars_000100(t *testing.T) {

	funcName := "TestErrPref_ConvertPrintableChars_000100"

	nonPrintableRuneArray := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	printableCharsStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	runeArray,
		err :=
		ErrPref{}.ConvertPrintableChars(
			printableCharsStr,
			funcName)

	if err != nil {
		t.Errorf("Error:\n"+
			"Error returned from ErrPref{}.ConvertPrintableChars()\n"+
			"Error = '%v'\n",
			err.Error())
		return
	}

	lenExpectedRuneArray := len(nonPrintableRuneArray)

	if lenExpectedRuneArray != len(runeArray) {
		t.Errorf("Error:\n"+
			"Expected lenExpectedRuneArray == len(runeArray).\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"lenExpectedRuneArray='%v'\n"+
			"      len(runeArray)='%v'\n",
			lenExpectedRuneArray,
			len(runeArray))
		return
	}

	for i := 0; i < len(nonPrintableRuneArray); i++ {
		if nonPrintableRuneArray[i] != runeArray[i] {
			t.Errorf("ERROR:\n"+
				"nonPrintableRuneArray[%v] != runeArray[%v]\n"+
				"nonPrintableRuneArray[%v]='%v'\n"+
				"runeArray[%v]='%v'\n",
				i,
				i,
				i,
				nonPrintableRuneArray[i],
				i,
				runeArray[i])
		}
	}

}

func TestErrPref_FmtStr_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\\nTx3.DoSomething() - Tx4() - Tx5()\\nTx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something()[SPACE]:[SPACE]A->B\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\n" +
		"[SPACE]:[SPACE][SPACE]A->B\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000400(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something() : A->B\nTx2.AVeryVeryLongMethodNameCalledSomething() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something() : A->B\n" +
		"Tx2.AVeryVeryLongMethodNameCalledSomething()\n" +
		"[SPACE]:[SPACE][SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPref_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()"

	actualStr := ErrPref{}.EPref(
		initialStr,
		"Tx5.BrandNewMethod()")

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPref_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr := ""

	expectedStr := "Tx5.BrandNewMethod()"

	actualStr := ErrPref{}.EPref(
		initialStr,
		"Tx5.BrandNewMethod()")

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPref_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	actualStr := ErrPref{}.EPref(
		initialStr,
		"")

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPref_000400(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	actualStr := ErrPref{}.EPref(
		initialStr,
		"")

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPrefCtx_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()[SPACE]:[SPACE]X->G"

	actualStr := ErrPref{}.EPrefCtx(
		initialStr,
		"Tx5.BrandNewMethod()",
		"X->G")

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

func TestErrPref_EPrefCtx_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()"

	actualStr := ErrPref{}.EPrefCtx(
		initialStr,
		"Tx5.BrandNewMethod()",
		"")

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

func TestErrPref_EPrefCtx_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	actualStr := ErrPref{}.EPrefCtx(
		initialStr,
		"",
		"G->X")

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

func TestErrPref_EPrefCtx_000400(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	actualStr := ErrPref{}.EPrefCtx(
		initialStr,
		"",
		"")

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

func TestErrPref_EPrefOld_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\nTx3.DoSomething() - Tx4() - Tx5()\nTx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPrefOld_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	actualStr := ErrPref{}.EPrefOld(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_EPrefOld_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr := "Tx1.StartThis()\n"

	expectedStr := "Tx1.StartThis()"

	actualStr := ErrPref{}.EPrefOld(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_GetLastEPref_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx5()"

	actualStr := ErrPref{}.GetLastEPref(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_GetMaxErrPrefTextLineLength_000100(t *testing.T) {

	ePref := ErrPref{}

	ePref.SetMaxErrPrefTextLineLengthToDefault()

	maxLineLen01 := ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen01 != 40 {
		t.Errorf("Error:\n"+
			"Expected Max== 40\n"+
			"Instead, Max== %v\n",
			maxLineLen01)
	}

}

func TestErrPref_SetMaxErrPrefTextLineLength_000100(t *testing.T) {

	ePref := ErrPref{}

	ePref.SetMaxErrPrefTextLineLengthToDefault()

	maxLineLen := ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 40 {
		t.Errorf("Error:\n"+
			"Expected Max== 40\n"+
			"Instead, Max== %v\n",
			maxLineLen)
		return
	}

	newMaxLineLen := uint(60)

	ePref.SetMaxErrPrefTextLineLength(newMaxLineLen)

	maxLineLen = ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 60 {
		t.Errorf("Error:\n"+
			"Expected Max== 60\n"+
			"Instead, Max== %v\n",
			maxLineLen)
		return
	}

	ePref.SetMaxErrPrefTextLineLengthToDefault()

	maxLineLen = ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 40 {
		t.Errorf("Error: Attempted to Reset Max Line Length to Default!\n"+
			"Expected Max== 40\n"+
			"Instead, Max== %v\n",
			maxLineLen)
	}

}

func TestErrPref_SetMaxErrPrefTextLineLengthToDefault_000100(t *testing.T) {

	ePref := ErrPref{}

	ePref.SetMaxErrPrefTextLineLengthToDefault()

	maxLineLen := ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 40 {
		t.Errorf("Error:\n"+
			"Expected Max== 40\n"+
			"Instead, Max== %v\n",
			maxLineLen)
		return
	}

	newMaxLineLen := uint(100)

	ePref.SetMaxErrPrefTextLineLength(newMaxLineLen)

	maxLineLen = ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 100 {
		t.Errorf("Error:\n"+
			"Expected Max== 100\n"+
			"Instead, Max== %v\n",
			maxLineLen)
		return
	}

	ePref.SetMaxErrPrefTextLineLengthToDefault()

	maxLineLen = ePref.GetMaxErrPrefTextLineLength()

	if maxLineLen != 40 {
		t.Errorf("Error: Attempted to Reset Max Line Length to Default!\n"+
			"Expected Max== 40\n"+
			"Instead, Max== %v\n",
			maxLineLen)
	}

}

func TestErrPref_SetCtxt_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	actualStr := ErrPref{}.SetCtxt(
		initialStr,
		"A!=B")

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_SetCtxt_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	actualStr := ErrPref{}.SetCtxt(
		initialStr,
		"A!=B")

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPref_SetCtxt_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		""

	actualStr := ErrPref{}.SetCtxt(
		initialStr,
		"A!=B")

	expectedStr := ""

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestNewErrPref_Multiple_000100(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ErrPref{}.SetMaxErrPrefTextLineLengthToDefault()

	// Setting Line Length to 60-Characters
	ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx7.TrySomethingNew()",
		"")

	ePrefix = ErrPref{}.SetCtxt(
		ePrefix,
		"something->newSomething")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx8.TryAnyCombination()",
		"")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx9.TryAHammer()")

	ePrefix = ErrPref{}.SetCtxt(ePrefix, "x->y")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx10.X()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx11.TryAnything()",
		"")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx12.TryASalad()",
		"")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx13.SomeFabulousAndComplexStuff()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew() : something->newSomething\\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\\n" +
			"Tx13.SomeFabulousAndComplexStuff()\\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePrefix = ErrPref{}.FmtStr(ePrefix)

	ErrPref{}.SetMaxErrPrefTextLineLengthToDefault()

	actualResult := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePrefix),
		false)

	if expectedResult != actualResult {
		t.Errorf("Error:\n"+
			"Expected Result        == '%v'\n"+
			"Instead, Actual Result == '%v'\n",
			expectedResult,
			actualResult)
	}

}

func TestNewErrPref_Multiple_000200(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	ePrefix :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx7.TrySomethingNew()",
		"")

	ePrefix = ErrPref{}.SetCtxt(
		ePrefix,
		"something->newSomething")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx8.TryAnyCombination()",
		"")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx9.TryAHammer()")

	ePrefix = ErrPref{}.SetCtxt(ePrefix, "x->y")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx10.X()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx11.TryAnything()",
		"")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx12.TryASalad()",
		"")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx13.SomeFabulousAndComplexStuff()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx15.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse()\\n" +
			"Tx3.DoSomething() - Tx4() - Tx5()\\n" +
			"Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew()\\n" +
			" :  something->newSomething\\n" +
			"Tx8.TryAnyCombination()\\n" +
			"Tx9.TryAHammer() : x->y - Tx10.X()\\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\\n" +
			"Tx13.SomeFabulousAndComplexStuff()\\n" +
			"Tx14.MoreAwesomeGoodness\\n" +
			" :  A=7 B=8 C=9\\n" +
			"Tx15.AVeryVeryLongMethodNameCalledSomething()\\n" +
			" :  A=7 B=8 C=9"

	ePrefix = ErrPref{}.FmtStr(ePrefix)

	ErrPref{}.SetMaxErrPrefTextLineLengthToDefault()

	actualResult := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePrefix),
		false)

	if expectedResult != actualResult {
		t.Errorf("Error:\n"+
			"Expected Result        == '%v'\n"+
			"Instead, Actual Result == '%v'\n",
			expectedResult,
			actualResult)
	}

}

func TestNewErrPref_Multiple_000300(t *testing.T) {

	ErrPref{}.SetMaxErrPrefTextLineLength(40)

	ePrefix :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx7.TrySomethingNew()",
		"")

	ePrefix = ErrPref{}.SetCtxt(
		ePrefix,
		"something->newSomething")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx8.TryAnyCombination()",
		"")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx9.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx10.TryAHammer()")

	ePrefix = ErrPref{}.SetCtxt(ePrefix, "x->y")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx11.X()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx12.TryAnything()",
		"")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx13.TryASalad()",
		"")

	ePrefix = ErrPref{}.EPref(
		ePrefix,
		"Tx14.SomeFabulousAndComplexStuff()")

	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx15.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse()\\n" +
			"Tx3.DoSomething() - Tx4() - Tx5()\\n" +
			"Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew()\\n" +
			" :  something->newSomething\\n" +
			"Tx8.TryAnyCombination()\\n" +
			"Tx9.AVeryVeryLongMethodNameCalledSomething()\\n" +
			" :  A=7 B=8 C=9\\n" +
			"Tx10.TryAHammer() : x->y - Tx11.X()\\n" +
			"Tx12.TryAnything() - Tx13.TryASalad()\\n" +
			"Tx14.SomeFabulousAndComplexStuff()\\n" +
			"Tx15.MoreAwesomeGoodness\\n" +
			" :  A=7 B=8 C=9"

	ePrefix = ErrPref{}.FmtStr(ePrefix)

	ErrPref{}.SetMaxErrPrefTextLineLengthToDefault()

	actualResult := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePrefix),
		false)

	if expectedResult != actualResult {
		t.Errorf("Error:\n"+
			"Expected Result        == '%v'\n"+
			"Instead, Actual Result == '%v'\n",
			expectedResult,
			actualResult)
	}

}

func TestErrPrefixDto_SetEPrefStrings_000100(t *testing.T) {

	var twoDSlice01 [][2]string

	twoDSlice01 = make([][2]string, 14)

	twoDSlice01[0][0] = "Tx1.Something()"
	twoDSlice01[0][1] = ""

	twoDSlice01[1][0] = "Tx2.SomethingElse()"
	twoDSlice01[1][1] = ""

	twoDSlice01[2][0] = "Tx3.DoSomething()"
	twoDSlice01[2][1] = ""

	twoDSlice01[3][0] = "Tx4()"
	twoDSlice01[3][1] = ""

	twoDSlice01[4][0] = "Tx5()"
	twoDSlice01[4][1] = ""

	twoDSlice01[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice01[5][1] = ""

	twoDSlice01[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice01[6][1] = "something->newSomething"

	twoDSlice01[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice01[7][1] = ""

	twoDSlice01[8][0] = "Tx9.TryAHammer()"
	twoDSlice01[8][1] = "x->y"

	twoDSlice01[9][0] = "Tx10.X()"
	twoDSlice01[9][1] = ""

	twoDSlice01[10][0] = "Tx11.TryAnything()"
	twoDSlice01[10][1] = ""

	twoDSlice01[11][0] = "Tx12.TryASalad()"
	twoDSlice01[11][1] = ""

	twoDSlice01[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice01[12][1] = ""

	twoDSlice01[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice01[13][1] = "A=7 B=8 C=9"

	var ePDto1, ePDtoExpected *ErrPrefixDto
	var err error

	ePDto1,
		err = ErrPrefixDto{}.NewIEmpty(
		twoDSlice01,
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice01)\n"+
			"%v\n", err.Error())

		return
	}

	var twoDSlice02 [][2]string

	twoDSlice02 = make([][2]string, 14)

	twoDSlice02[0][0] = "Tx1.TryBeer()"
	twoDSlice02[0][1] = ""

	twoDSlice02[1][0] = "Tx2.TryMescal()"
	twoDSlice02[1][1] = ""

	twoDSlice02[2][0] = "Tx3.TryWhiskey()"
	twoDSlice02[2][1] = ""

	twoDSlice02[3][0] = "Tx4.TryMoonShine()"
	twoDSlice02[3][1] = ""

	twoDSlice02[4][0] = "Tx5.TryScotch()"
	twoDSlice02[4][1] = ""

	twoDSlice02[5][0] = "Tx6.TryWine()"
	twoDSlice02[5][1] = ""

	twoDSlice02[6][0] = "Tx7.TryBrandy()"
	twoDSlice02[6][1] = "something->newSomething"

	twoDSlice02[7][0] = "Tx8.TryCognac()"
	twoDSlice02[7][1] = ""

	twoDSlice02[8][0] = "Tx9.TryAle()"
	twoDSlice02[8][1] = "x->y"

	twoDSlice02[9][0] = "Tx10.Vodka()"
	twoDSlice02[9][1] = ""

	twoDSlice02[10][0] = "Tx11.TryRum()"
	twoDSlice02[10][1] = ""

	twoDSlice02[11][0] = "Tx12.TryBourbon()"
	twoDSlice02[11][1] = ""

	twoDSlice02[12][0] = "Tx13.TryTequila()"
	twoDSlice02[12][1] = ""

	twoDSlice02[13][0] = "Tx14.TryWater"
	twoDSlice02[13][1] = "A=7 B=8 C=9"

	ePDtoExpected,
		err = ErrPrefixDto{}.NewIEmpty(
		twoDSlice02,
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice02)\n"+
			"%v\n", err.Error())

		return
	}

	ePDtoExpected.SetMaxTextLineLen(40)

	ePDto1.SetEPrefStrings(twoDSlice02)

	ePDto1.SetMaxTextLineLen(40)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoExpected.String()),
		true)

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		true)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected expectedStr == actualStr\n"+
			"Instead:\n"+
			"expectedStr=\n%v\n\n"+
			"actualStr=\n%v\n\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto1.Equal(ePDtoExpected) {
		t.Errorf("Error: Expected ePDto1 == ePDtoExpected\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"ePDto1.String()=\n%v\n\n"+
			"ePDtoExpected.String()=\n%v\n\n",
			ePDto1.String(),
			ePDtoExpected.String())
	}

}

func TestErrPrefixDto_SetEPrefStrings_000200(t *testing.T) {

	ePDto1 := ErrPrefixDto{}

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

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto1.SetEPrefStrings(twoDSlice)

	oldCollectionLen := ePDto1.GetEPrefCollectionLen()

	var twoDSlice02 [][2]string

	ePDto1.SetEPrefStrings(twoDSlice02)

	newCollectionLen := ePDto1.GetEPrefCollectionLen()

	if oldCollectionLen != newCollectionLen {
		t.Errorf("ERROR:\n"+
			"Expected oldCollectionLen == newCollectionLen because\n"+
			"in the call to ePDto1.SetEPrefStrings(twoDSlice02),\n"+
			"twoDSlice02 is nil and therefore new collection length\n"+
			"should be unchanged.\n"+
			"HOWEVER, oldCollectionLen != newCollectionLen!!!\n"+
			"oldCollectionLen='%v'\n"+
			"newCollectionLen='%v'\n",
			oldCollectionLen,
			newCollectionLen)
	}
}
