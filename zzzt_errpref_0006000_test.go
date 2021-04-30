package errpref

import "testing"

func TestErrPrefixDto_SetLeftMarginLength_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(3)

	ePDto.SetLeftMarginChar(' ')

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	expectedStr := "[SPACE][SPACE][SPACE]Tx1.Something()[SPACE]-" +
		"[SPACE]Tx2.SomethingElse()\\n[SPACE][SPACE][SPACE]" +
		"Tx3.DoSomething()[SPACE]-[SPACE]Tx4()[SPACE]-" +
		"[SPACE]Tx5()\\n[SPACE][SPACE][SPACE]Tx6.DoSomethingElse()" +
		"\\n[SPACE][SPACE][SPACE]Tx7.TrySomethingNew()\\n" +
		"[SPACE][SPACE][SPACE][SPACE]:[SPACE][SPACE]" +
		"something->newSomething\\n[SPACE][SPACE][SPACE]" +
		"Tx8.TryAnyCombination()\\n[SPACE][SPACE][SPACE]" +
		"Tx9.TryAHammer()[SPACE]:[SPACE]x->y[SPACE]-[SPACE]Tx10.X()\\n" +
		"[SPACE][SPACE][SPACE]Tx11.TryAnything()[SPACE]-" +
		"[SPACE]Tx12.TryASalad()\\n[SPACE][SPACE][SPACE]" +
		"Tx13.SomeFabulousAndComplexStuff()\\n[SPACE][SPACE]" +
		"[SPACE]Tx14.MoreAwesomeGoodness\\n[SPACE][SPACE][SPACE]" +
		"[SPACE]:[SPACE][SPACE]A=7[SPACE]B=8[SPACE]C=9"

	if expectedStr != actualConvertedStr {
		t.Errorf("Error expectedStr != actualConvertedStr\n"+
			"expectedStr=\n%v\n\n"+
			"actualStr=\n%v\n\n",
			expectedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_SetLeftMarginLength_000200(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(0)

	ePDto.SetLeftMarginChar('*')

	ePDto2 := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetLeftMarginChar('*')

	if !ePDto.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
		return
	}

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_SetLeftMarginChar_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(3)

	ePDto.SetLeftMarginChar('*')

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	expectedStr := "***Tx1.Something()[SPACE]-[SPACE]Tx2.SomethingElse()\n" +
		"***Tx3.DoSomething()[SPACE]-[SPACE]Tx4()[SPACE]-[SPACE]Tx5()\n" +
		"***Tx6.DoSomethingElse()\n" +
		"***Tx7.TrySomethingNew()\n" +
		"***[SPACE]:[SPACE][SPACE]something->newSomething\n" +
		"***Tx8.TryAnyCombination()\n" +
		"***Tx9.TryAHammer()[SPACE]:[SPACE]x->y[SPACE]-[SPACE]Tx10.X()\n" +
		"***Tx11.TryAnything()[SPACE]-[SPACE]Tx12.TryASalad()\n" +
		"***Tx13.SomeFabulousAndComplexStuff()\n" +
		"***Tx14.MoreAwesomeGoodness\n" +
		"***[SPACE]:[SPACE][SPACE]A=7[SPACE]B=8[SPACE]C=9"

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	if expectedStr != actualConvertedStr {
		t.Errorf("Error expectedStr != actualConvertedStr\n"+
			"expectedStr=\n%v\n\n"+
			"actualStr=\n%v\n\n",
			expectedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_SetLeftMarginChar_000200(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	// Left Margin Length defaults to zer

	ePDto.SetLeftMarginChar('*')

	ePDto2 := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetLeftMarginChar('*')

	if !ePDto.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
		return
	}

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_SetMaxTextLineLen_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(5)

	defaultMaxTextLineLen := ePDto.GetMaxTextLineLenDefault()

	actualMaxTextLineLen := ePDto.GetMaxTextLineLen()

	if actualMaxTextLineLen != defaultMaxTextLineLen {
		t.Errorf("ERROR:\n"+
			"Expected actualMaxTextLineLen == defaultMaxTextLineLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			" actualMaxTextLineLen='%v'\n"+
			"defaultMaxTextLineLen='%v'\n",
			actualMaxTextLineLen,
			defaultMaxTextLineLen)
	}

}
