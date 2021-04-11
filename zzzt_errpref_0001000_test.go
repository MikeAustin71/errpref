package errpref

import (
	"testing"
)

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
