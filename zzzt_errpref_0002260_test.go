package errpref

import "testing"

func TestErrPrefixDto_GetInputStringDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_GetInputStringDelimiters_000100() "

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	expectedDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto1.SetInputStringDelimiters(
			expectedDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	inputDelimiters :=
		ePDto1.GetInputStringDelimiters()

	if !expectedDelimiters.Equal(&inputDelimiters) {
		t.Errorf("ERROR:\n"+
			"Expected expectedDelimiters==inputDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDelimiters='%v'\n"+
			"inputDelimiters='%v'\n",
			expectedDelimiters.String(),
			inputDelimiters.String())
	}

}

func TestErrPrefixDto_GetOutputStringDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_GetOutputStringDelimiters_000100() "

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	expectedInputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"expectedInputDelimiters\n"+
			"%v\n", err.Error())
		return
	}

	var expectedOutputDelimiters ErrPrefixDelimiters

	expectedOutputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n |",
		" $ ",
		"\n      #",
		" ### ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"'expectedOutputDelimiters'\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto1.SetInputStringDelimiters(
			expectedInputDelimiters,
			funcName+"expectedInputDelimiters")

	if err != nil {
		t.Errorf("Error from ePDto1.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto1.SetOutputStringDelimiters(
			expectedOutputDelimiters,
			funcName+"expectedOutputDelimiters")

	if err != nil {
		t.Errorf("Error from ePDto1.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	inputStrDelimiters,
		outputStrDelimiters :=
		ePDto1.GetStrDelimiters()

	if !expectedInputDelimiters.Equal(&inputStrDelimiters) {
		t.Errorf("ERROR:\n"+
			"Expected expectedInputDelimiters==inputStrDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedInputDelimiters='%v'\n"+
			"     inputStrDelimiters='%v'\n",
			expectedInputDelimiters.String(),
			inputStrDelimiters.String())

		return
	}

	if !expectedOutputDelimiters.Equal(&outputStrDelimiters) {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputDelimiters==outputStrDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputDelimiters='%v'\n"+
			"     outputStrDelimiters='%v'\n",
			expectedOutputDelimiters.String(),
			outputStrDelimiters.String())

		return
	}

}

func TestErrPrefixDto_GetStrDelimiters_0100(t *testing.T) {

	funcName := "TestErrPrefixDto_GetStrDelimiters_0100() "

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	expectedDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto1.SetOutputStringDelimiters(
			expectedDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	outputStrDelimiters :=
		ePDto1.GetOutputStringDelimiters()

	if !expectedDelimiters.Equal(&outputStrDelimiters) {
		t.Errorf("ERROR:\n"+
			"Expected expectedDelimiters==outputStrDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDelimiters='%v'\n"+
			"outputStrDelimiters='%v'\n",
			expectedDelimiters.String(),
			outputStrDelimiters.String())
	}

}

func TestErrPrefixDto_SetOutputStringDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_SetOutputStringDelimiters_000100() "

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	outputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

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

	err =
		ePDto1.SetOutputStringDelimiters(
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutput := "Tx1.Something() -*- Tx2.SomethingElse()\n" +
		"  *Tx3.DoSomething() -*- Tx4() -*- Tx5()\n" +
		"  *Tx6.DoSomethingElse()\n  *Tx7.TrySomethingNew()\n" +
		"      %something->newSomething\n" +
		"  *Tx8.TryAnyCombination()\n" +
		"  *Tx9.TryAHammer() &&& x->y -*- Tx10.X()\n" +
		"  *Tx11.TryAnything() -*- Tx12.TryASalad()\n" +
		"  *Tx13.SomeFabulousAndComplexStuff()\n" +
		"  *Tx14.MoreAwesomeGoodness\n" +
		"      %A=7 B=8 C=9"

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(expectedOutput),
			false)

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	if expectedOutputStr != ePDto1Str {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputStr == ePDto1St\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto1Str = '%v'\n",
			expectedOutputStr,
			ePDto1Str)
	}

}

func TestErrPrefixDto_SetOutputStringDelimiters_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_SetOutputStringDelimiters_000200"

	ePDtoOne := ErrPrefixDto{}.New()

	ePDtoOne.SetMaxTextLineLen(60)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDtoOne.SetEPrefOld(initialStr)

	outputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDtoOne.SetOutputStringDelimiters(
			outputDelimiters,
			funcName+"outputDelimiters")

	if err != nil {
		t.Errorf("Error from ePDtoOne.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return

	}

	outputStrOne := ePDtoOne.String()

	ePDtoTwo := ErrPrefixDto{}.New()

	ePDtoTwo.SetMaxTextLineLen(60)

	err =
		ePDtoTwo.SetInputStringDelimiters(
			outputDelimiters,
			funcName+"outputDelimiters ")

	if err != nil {
		t.Errorf("Error from ePDtoOne.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDtoTwo.SetOutputStringDelimiters(
			outputDelimiters,
			funcName+"outputDelimiters ")

	if err != nil {
		t.Errorf("Error from ePDtoOne.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoTwo.SetEPrefOld(outputStrOne)

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDtoOne.String()),
			true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoTwo.String()),
		true)

	if expectedOutputStr != ePDto2Str {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputStr == ePDto2Str\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto2Str = '%v'\n",
			expectedOutputStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_SetInputStringDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_SetInputStringDelimiters_000100() "

	inputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	err =
		ePDto1.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	rawInput := "Tx1.Something() -*- Tx2.SomethingElse()\n" +
		"  *Tx3.DoSomething() -*- Tx4() -*- Tx5()\n" +
		"  *Tx6.DoSomethingElse()\n  *Tx7.TrySomethingNew()\n" +
		"      %something->newSomething\n" +
		"  *Tx8.TryAnyCombination()\n" +
		"  *Tx9.TryAHammer() &&& x->y -*- Tx10.X()\n" +
		"  *Tx11.TryAnything() -*- Tx12.TryASalad()\n" +
		"  *Tx13.SomeFabulousAndComplexStuff()\n" +
		"  *Tx14.MoreAwesomeGoodness\n" +
		"      %A=7 B=8 C=9"

	ePDto1.SetEPrefOld(rawInput)

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

	ePDtoBase := ErrPrefixDto{}.New()

	ePDtoBase.SetMaxTextLineLen(40)

	ePDtoBase.SetEPrefStrings(twoDSlice)

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDtoBase.String()),
			false)

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	if expectedOutputStr != ePDto1Str {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputStr == ePDto1St\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto1Str = '%v'\n",
			expectedOutputStr,
			ePDto1Str)
	}

}

func TestErrPrefixDto_SetStrDelimitersToDefault_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_SetStrDelimitersToDefault_000100() "

	inputDelimiters,
		err := ErrPrefixDelimiters{}.New(
		"\n  *",
		" -*- ",
		"\n      %",
		" &&& ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetMaxTextLineLen(40)

	err =
		ePDto1.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	rawInput := "Tx1.Something() -*- Tx2.SomethingElse()\n" +
		"  *Tx3.DoSomething() -*- Tx4() -*- Tx5()\n" +
		"  *Tx6.DoSomethingElse()\n  *Tx7.TrySomethingNew()\n" +
		"      %something->newSomething\n" +
		"  *Tx8.TryAnyCombination()\n" +
		"  *Tx9.TryAHammer() &&& x->y -*- Tx10.X()\n" +
		"  *Tx11.TryAnything() -*- Tx12.TryASalad()\n" +
		"  *Tx13.SomeFabulousAndComplexStuff()\n" +
		"  *Tx14.MoreAwesomeGoodness\n" +
		"      %A=7 B=8 C=9"

	ePDto1.SetEPrefOld(rawInput)

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

	ePDtoBase := ErrPrefixDto{}.New()

	ePDtoBase.SetMaxTextLineLen(40)

	ePDtoBase.SetEPrefStrings(twoDSlice)

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDtoBase.String()),
			false)

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	if expectedOutputStr != ePDto1Str {
		t.Errorf("ERROR:\n"+
			"Expected expectedOutputStr == ePDto1St\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto1Str = '%v'\n",
			expectedOutputStr,
			ePDto1Str)
	}

	defaultStrDelimiters := ErrPrefixDelimiters{}

	defaultStrDelimiters.SetToDefault()

	err =
		ePDtoBase.SetInputStringDelimiters(
			defaultStrDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDtoBase.SetInputStringDelimiters"+
			"(defaultStrDelimiters)\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDtoBase.SetOutputStringDelimiters(
			defaultStrDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDtoBase.SetOutputStringDelimiters"+
			"(defaultStrDelimiters)\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1.SetStrDelimitersToDefault()

	expectedOutputStr =
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDtoBase.String()),
			false)

	ePDto1Str = ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	if expectedOutputStr != ePDto1Str {
		t.Errorf("ERROR Default Series #2:\n"+
			"Expected expectedOutputStr == ePDto1St\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedOutputStr = '%v'\n"+
			"        ePDto1Str = '%v'\n",
			expectedOutputStr,
			ePDto1Str)
	}
}
