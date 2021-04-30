package errpref

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrPrefixDto_NewIBasicErrorPrefix_000100(t *testing.T) {

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

	iBasicPref := testIBasicErrPref{}

	var err error

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto *ErrPrefixDto

	ePDto,
		err = ErrPrefixDto{}.NewIBasicErrorPrefix(
		&iBasicPref,
		"",
		"")

	if err != nil {
		t.Errorf("Error returned by NewIBasicErrorPrefix()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.Ptr()

	ePDto2.SetEPrefStrings(twoDSlice)

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIBasicErrorPrefix_000200(t *testing.T) {

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

	iBasicPref := testIBasicErrPref{}

	var err error

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto *ErrPrefixDto

	newErrPrefix := "Tx15.TomorrowIsABetterDay()"

	ePDto,
		err = ErrPrefixDto{}.NewIBasicErrorPrefix(
		&iBasicPref,
		newErrPrefix,
		"")

	if err != nil {
		t.Errorf("Error returned by NewIBasicErrorPrefix()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.Ptr()

	ePDto2.SetEPrefStrings(twoDSlice)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPref(newErrPrefix)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}
}

func TestErrPrefixDto_NewIBasicErrorPrefix_000300(t *testing.T) {

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

	iBasicPref := testIBasicErrPref{}

	var err error

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto *ErrPrefixDto

	newErrPrefix := "Tx15.TomorrowIsABetterDay()"

	ePDto,
		err = ErrPrefixDto{}.NewIBasicErrorPrefix(
		&iBasicPref,
		newErrPrefix,
		"")

	if err != nil {
		t.Errorf("Error returned by NewIBasicErrorPrefix()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.Ptr()

	ePDto2.SetEPrefStrings(twoDSlice)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPref(newErrPrefix)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}
}

func TestErrPrefixDto_NewIBasicErrorPrefix_000400(t *testing.T) {

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

	iBasicPref := testIBasicErrPref{}

	var err error

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto *ErrPrefixDto

	newErrPrefix := "Tx15.TomorrowIsABetterDay()"
	newErrContext := "X=7/Y  Y==0"

	ePDto,
		err = ErrPrefixDto{}.NewIBasicErrorPrefix(
		&iBasicPref,
		newErrPrefix,
		newErrContext)

	if err != nil {
		t.Errorf("Error returned by NewIBasicErrorPrefix()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.Ptr()

	ePDto2.SetEPrefStrings(twoDSlice)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefCtx(newErrPrefix, newErrContext)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}
}

func TestErrPrefixDto_NewEPrefCollection_000100(t *testing.T) {

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDtoNoOfElements,
		ePDto := ErrPrefixDto{}.NewEPrefCollection(
		expectedStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefOld(expectedStr)

	ePDto2NoOfElements := ePDto2.GetEPrefCollectionLen()

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if ePDtoNoOfElements != ePDto2NoOfElements {
		t.Errorf("Error: Expected ePDtoNoOfElements==ePDto2NoOfElements.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDtoNoOfElements=\n%v\n\nePDto2NoOfElements=\n%v\n\n",
			ePDtoNoOfElements,
			ePDto2NoOfElements)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoNoePDtoStrOfElements==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDtoStr=\n%v\n\nePDto2Str=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}
}

func TestErrPrefixDto_NewEPrefCtx_000100(t *testing.T) {

	newErrPrefix := "Tx5.BrandNewMethod()"
	newErrContext := "X->G"

	expectedDto := new(ErrPrefixDto)

	expectedDto.SetMaxTextLineLen(40)

	expectedDto.SetEPrefCtx(newErrPrefix, newErrContext)

	actualDto := ErrPrefixDto{}.NewEPrefCtx(
		newErrPrefix,
		newErrContext)

	actualDto.SetMaxTextLineLen(40)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedDto.String()),
		true)

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualDto.String()),
		true)

	if !actualDto.Equal(expectedDto) {
		t.Errorf("ERROR: Expected expectedDto==actualDto.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDto='%v'\n"+
			"actualDto='%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected expectedStr==actualStr.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedStr='%v'\n"+
			"actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewEPrefCtx_000200(t *testing.T) {

	newErrPrefix := "Tx5.BrandNewMethod()"
	newErrContext := ""

	expectedDto := new(ErrPrefixDto)

	expectedDto.SetMaxTextLineLen(40)

	expectedDto.SetEPrefCtx(newErrPrefix, newErrContext)

	actualDto := ErrPrefixDto{}.NewEPrefCtx(
		newErrPrefix,
		newErrContext)

	actualDto.SetMaxTextLineLen(40)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedDto.String()),
		true)

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualDto.String()),
		true)

	if !actualDto.Equal(expectedDto) {
		t.Errorf("ERROR: Expected expectedDto==actualDto.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDto='%v'\n"+
			"actualDto='%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected expectedStr==actualStr.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedStr='%v'\n"+
			"actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewEPrefCtx_000300(t *testing.T) {

	newErrPrefix := ""
	newErrContext := ""

	expectedDto := new(ErrPrefixDto)

	expectedDto.SetMaxTextLineLen(40)

	expectedDto.SetEPrefCtx(newErrPrefix, newErrContext)

	actualDto := ErrPrefixDto{}.NewEPrefCtx(
		newErrPrefix,
		newErrContext)

	actualDto.SetMaxTextLineLen(40)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedDto.String()),
		true)

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualDto.String()),
		true)

	if !actualDto.Equal(expectedDto) {
		t.Errorf("ERROR: Expected expectedDto==actualDto.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDto='%v'\n"+
			"actualDto='%v'\n",
			expectedStr,
			actualStr)

		return
	}

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected expectedStr==actualStr.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedStr='%v'\n"+
			"actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewEPrefOld_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\nTx3.DoSomething() - Tx4() - Tx5()\nTx6.DoSomethingElse()"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewEPrefOld_000200(t *testing.T) {

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewFromIErrorPrefix_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2 := ErrPrefixDto{}.NewFromIErrorPrefix(
		&ePDto)

	ePDto2.SetMaxTextLineLen(40)

	areEqual := ePDto.Equal(&ePDto2)

	if !areEqual {
		t.Error("Error: Expected ePDto and ePDto2 would be equal.\n" +
			"THEY ARE NOT EQUAL!\n")
	}

}

func TestErrPrefixDto_NewFromIErrorPrefix_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto2 := ErrPrefixDto{}.NewFromIErrorPrefix(
		&ePDto)

	ePDto2.SetMaxTextLineLen(40)

	ePDto2Str := ePDto2.String()

	if len(ePDto2Str) > 0 {
		t.Errorf("ERROR:\n"+
			"Expected ePDto2 string would be empty or zero length.\n"+
			"IT IS NOT! String length is greater than zero!\n"+
			"ePDto2Str='%v'\n",
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewFromStrings_000100(t *testing.T) {
	funcName := "TestErrPrefixDto_NewFromStrings_000100"

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

	var ePDto2 ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewFromStrings(
		initialStr,
		newErrPrefix,
		newErrCtx,
		outputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.NewFromStrings()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDto.String()),
			false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

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

func TestErrPrefixDto_NewFromStrings_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_NewFromStrings_000100"

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

	var ePDto2 ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewFromStrings(
		initialStr,
		"",
		"",
		outputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.NewFromStrings()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDto.String()),
			false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

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

func TestErrPrefixDto_NewFromStrings_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_NewFromStrings000300"

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

	outputDelimiters,
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
			outputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetInputStringDelimiters()\n"+
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

	inputDelimiters := ErrPrefixDelimiters{}

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	initialStr := ePDto.String()

	_,
		err = ErrPrefixDto{}.NewFromStrings(
		initialStr,
		newErrPrefix,
		newErrCtx,
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return fromErrPrefixDto{}.NewFromStrings()\n" +
			"because 'inputDelimiters' is empty and invalid!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!")
	}
}

func TestErrPrefixDto_NewFromStrings_000400(t *testing.T) {

	funcName := "TestErrPrefixDto_NewFromStrings000400"

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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	outputDelimiters := ErrPrefixDelimiters{}

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	initialStr := ePDto.String()

	_,
		err = ErrPrefixDto{}.NewFromStrings(
		initialStr,
		newErrPrefix,
		newErrCtx,
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return fromErrPrefixDto{}.NewFromStrings()\n" +
			"because 'outputDelimiters' is empty and invalid!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!")
	}

}

func TestErrPrefixDto_NewFromStrings_000500(t *testing.T) {

	funcName := "TestErrPrefixDto_NewFromStrings000400"

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

	var ePDto ErrPrefixDto

	ePDto,
		err = ErrPrefixDto{}.NewFromStrings(
		"",
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error:\n"+
			"Should NOT HAVE received an error return fromErrPrefixDto{}.NewFromStrings()\n"+
			"because both 'inputDelimiters' and 'outputDelimiters' are valid!\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!\n"+
			"Error='%v'\n", err.Error())

		return
	}

	actualInputDelims,
		actualOutputDelims := ePDto.GetStrDelimiters()

	if !actualInputDelims.Equal(&inputDelimiters) {
		t.Errorf("Error:\n"+
			"Original input delimiters not equal to "+
			"final input delimiters!\n"+
			"Original input delimiters='%v'\n"+
			"   Final input delimiters='%v'\n",
			inputDelimiters.String(),
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

func TestErrPrefixDto_NewIEmpty_000100(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000200(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	funcName := "TestErrPrefixDto_NewIEmpty_000200"

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		funcName,
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	if ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto!=ePDto2.\n"+
			"However, THEY ARE EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
	}

}

func TestErrPrefixDto_NewIEmpty_000300(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto.SetMaxTextLineLen(40)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 13)

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

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000400(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew()\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness()"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	var oneDSlice []string

	oneDSlice = make([]string, 14)

	oneDSlice[0] = "Tx1.Something()"

	oneDSlice[1] = "Tx2.SomethingElse()"

	oneDSlice[2] = "Tx3.DoSomething()"

	oneDSlice[3] = "Tx4()"

	oneDSlice[4] = "Tx5()"

	oneDSlice[5] = "Tx6.DoSomethingElse()"

	oneDSlice[6] = "Tx7.TrySomethingNew()"

	oneDSlice[7] = "Tx8.TryAnyCombination()"

	oneDSlice[8] = "Tx9.TryAHammer()"

	oneDSlice[9] = "Tx10.X()"

	oneDSlice[10] = "Tx11.TryAnything()"

	oneDSlice[11] = "Tx12.TryASalad()"

	oneDSlice[12] = "Tx13.SomeFabulousAndComplexStuff()"

	oneDSlice[13] = "Tx14.MoreAwesomeGoodness()"

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		oneDSlice,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000500(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		initialStr,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000600(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()" +
			"Tx7.TrySomethingNew() : something->newSomething" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()" +
			"Tx11.TryAnything() - Tx12.TryASalad()" +
			"Tx13.SomeFabulousAndComplexStuff()" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		initialStr,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000700(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error
	var ePDtoCopy ErrPrefixDto
	var ePDto2 *ErrPrefixDto

	funcName :=
		"TestErrPrefixDto_NewIEmpty_000700()"

	ePDtoCopy,
		err = ePDto.CopyOut(funcName)

	if err != nil {
		t.Errorf("Error from ePDto.CopyOut(funcName)\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDtoCopy,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000800(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error
	var ePDtoCopy *ErrPrefixDto
	var ePDto2 *ErrPrefixDto

	ePDtoCopy = ePDto.CopyPtr()

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDtoCopy,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000900(t *testing.T) {

	var err error
	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		nil,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	if ePDto2.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: Expected Collection Length==0\n"+
			"Instead Collection Length=='%v'\n",
			ePDto2.GetEPrefCollectionLen())
	}
}

func TestErrPrefixDto_NewIEmpty_001000(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	iBasicPref := testIBasicErrPref{}

	var err error

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		&iBasicPref,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_001100(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	iStringerEPref := testIStringerErrPref{}

	var err error

	iStringerEPref.Set(initialStr)

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		&iStringerEPref,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty(iStringerEPref)\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_001200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	sb := strings.Builder{}

	sb.WriteString(initialStr)

	var err error

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		sb,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty(strings.Builder{})\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_001300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	sb := strings.Builder{}

	sb.WriteString(initialStr)

	var sbPtr *strings.Builder

	sbPtr = &sb

	var err error

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		sbPtr,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty(strings.Builder{})\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_001400(t *testing.T) {

	var sbPtr *strings.Builder

	var err error
	_,
		err = ErrPrefixDto{}.NewIEmpty(
		sbPtr,
		"",
		"")

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}.NewIEmpty(sbPtr)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmpty_001500(t *testing.T) {

	var err error
	var ePDto1 *ErrPrefixDto

	_,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDto1,
		"",
		"")

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}.NewIEmpty(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmpty_001600(t *testing.T) {

	var err error
	var ePDto2 *ErrPrefixDto

	_,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDto2,
		"",
		"")

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}.NewIEmpty(sbPtr)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000100(t *testing.T) {
	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000100"

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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	ePDto.SetEPrefCtx(newErrPrefix, newErrCtx)

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

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		twoDSlice,
		newErrPrefix,
		newErrCtx,
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDto.String()),
			false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			expectedOutputStr,
			ePDto2Str)

		return
	}

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

	if !actualInputDelims.Equal(&inputDelimiters) {
		t.Errorf("Error:\n"+
			"Original input delimiters not equal to "+
			"final input delimiters!\n"+
			"Original input delimiters='%v'\n"+
			"   Final input delimiters='%v'\n",
			inputDelimiters.String(),
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

func TestErrPrefixDto_NewIEmptyWithDelimiters_000200(t *testing.T) {
	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000100"

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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	ePDto.SetEPrefCtx(newErrPrefix, "")

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

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		twoDSlice,
		newErrPrefix,
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	expectedOutputStr :=
		ErrPref{}.ConvertNonPrintableChars(
			[]rune(ePDto.String()),
			false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			expectedOutputStr,
			ePDto2Str)

		return
	}

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

	if !actualInputDelims.Equal(&inputDelimiters) {
		t.Errorf("Error:\n"+
			"Original input delimiters not equal to "+
			"final input delimiters!\n"+
			"Original input delimiters='%v'\n"+
			"   Final input delimiters='%v'\n",
			inputDelimiters.String(),
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

func TestErrPrefixDto_NewIEmptyWithDelimiters_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000300"
	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew()\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness()"

	ePDto.SetEPrefOld(initialStr)

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

	ePDto.SetMaxTextLineLen(40)

	var oneDSlice []string

	oneDSlice = make([]string, 14)

	oneDSlice[0] = "Tx1.Something()"

	oneDSlice[1] = "Tx2.SomethingElse()"

	oneDSlice[2] = "Tx3.DoSomething()"

	oneDSlice[3] = "Tx4()"

	oneDSlice[4] = "Tx5()"

	oneDSlice[5] = "Tx6.DoSomethingElse()"

	oneDSlice[6] = "Tx7.TrySomethingNew()"

	oneDSlice[7] = "Tx8.TryAnyCombination()"

	oneDSlice[8] = "Tx9.TryAHammer()"

	oneDSlice[9] = "Tx10.X()"

	oneDSlice[10] = "Tx11.TryAnything()"

	oneDSlice[11] = "Tx12.TryASalad()"

	oneDSlice[12] = "Tx13.SomeFabulousAndComplexStuff()"

	oneDSlice[13] = "Tx14.MoreAwesomeGoodness()"

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		oneDSlice,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000400(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000400"
	maxLineLen := 60

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(maxLineLen)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoCopy := ePDto.Copy()

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

	initialStr = ePDtoCopy.String()

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		initialStr,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000500(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000500"
	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	ePDto.SetMaxTextLineLen(60)

	initialStr = ePDto.String()

	var ePDtoCopy ErrPrefixDto

	ePDtoCopy,
		err = ePDto.CopyOut(funcName)

	if err != nil {
		t.Errorf("Error from ePDto.CopyOut(funcName)\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		ePDtoCopy,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(60)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDtoStr=\n%v\n\nePDto2Str=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000600(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000600"
	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

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

	ePDto.SetMaxTextLineLen(60)

	var ePDtoCopy *ErrPrefixDto

	ePDtoCopy = ePDto.CopyPtr()

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		ePDtoCopy,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(60)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDtoStr=\n%v\n\nePDto2Str=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000700(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000700"

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

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		nil,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	if ePDto2.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: Expected Collection Length==0\n"+
			"Instead Collection Length=='%v'\n",
			ePDto2.GetEPrefCollectionLen())
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000800(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000800()"

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 60

	ePDto.SetMaxTextLineLen(maxLineLen)

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

	ePDto.SetEPrefStrings(twoDSlice)

	var err error
	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
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

	iBasicPref := testIBasicErrPref{}

	err = iBasicPref.SetEPrefStrings(
		twoDSlice)

	if err != nil {
		t.Errorf("Error from iBasicPref.SetEPrefStrings()\n"+
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

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		&iBasicPref,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_000900(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000900()"

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 60

	ePDto.SetMaxTextLineLen(maxLineLen)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error

	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoCopy := ePDto.Copy()

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

	initialStr = ePDtoCopy.String()

	iStringerEPref := testIStringerErrPref{}

	iStringerEPref.Set(initialStr)

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		&iStringerEPref,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001000(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001000()"

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 60

	ePDto.SetMaxTextLineLen(maxLineLen)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error

	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoCopy := ePDto.Copy()

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

	initialStr = ePDtoCopy.String()

	sb := strings.Builder{}

	sb.WriteString(initialStr)

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		sb,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001100(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001100()"

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 60

	ePDto.SetMaxTextLineLen(maxLineLen)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error

	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
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

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		t.Errorf("Error from ePDto.SetOutputStringDelimiters() #1\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoCopy := ePDto.Copy()

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

	initialStr = ePDtoCopy.String()

	sb := strings.Builder{}

	sb.WriteString(initialStr)

	var sbPtr *strings.Builder

	sbPtr = &sb

	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		sbPtr,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)

		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001200(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001200()"

	var err error
	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			" - inputDelimiters\n"+
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
			"- outputDelimiters\n"+
			"%v\n", err.Error())
		return
	}

	var sbPtr *strings.Builder

	_,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		sbPtr,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}." +
			"NewIEmptyWithDelimiters(sbPtr) because\n" +
			"sbPtr is nil!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001300(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001200()"

	var err error
	var inputDelimiters ErrPrefixDelimiters

	inputDelimiters,
		err = ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDelimiters{}.New()\n"+
			" - inputDelimiters\n"+
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
			"- outputDelimiters\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto1 *ErrPrefixDto

	_,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		ePDto1,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}." +
			"NewIEmptyWithDelimiters(ePDto1) because\n" +
			"ePDto1 is nil!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001400(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001400()"

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

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	inputDelimiters.Empty()

	_,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		twoDSlice,
		newErrPrefix,
		newErrCtx,
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}." +
			"NewIEmptyWithDelimiters(twoDSlice) because\n" +
			"inputDelimiters is empty and therefore invalid!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestErrPrefixDto_NewIEmptyWithDelimiters_001500(t *testing.T) {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_001500()"

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

	newErrPrefix := "Tx15.TomorrowWillBeBetter()"

	newErrCtx := "7/X=Z  X=0"

	outputDelimiters.Empty()

	_,
		err = ErrPrefixDto{}.NewIEmptyWithDelimiters(
		twoDSlice,
		newErrPrefix,
		newErrCtx,
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}." +
			"NewIEmptyWithDelimiters(twoDSlice) because\n" +
			"inputDelimiters is empty and therefore invalid!\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

type testIBuilderErrPref struct {
	errPrefInfo [][2]string
}

func (tIBuilder *testIBuilderErrPref) GetEPrefStrings() [][2]string {

	if tIBuilder.errPrefInfo == nil {
		return nil
	}

	lenTwoDArray := len(tIBuilder.errPrefInfo)

	result := make([][2]string, lenTwoDArray)

	if lenTwoDArray == 0 {
		return result
	}

	copy(result, tIBuilder.errPrefInfo)

	return result
}

func (tIBuilder *testIBuilderErrPref) SetEPrefStrings(twoDStrArray [][2]string) {

	lenTwoDArray := len(twoDStrArray)

	if lenTwoDArray == 0 {
		return
	}

	tIBuilder.errPrefInfo = make([][2]string, lenTwoDArray)

	copy(tIBuilder.errPrefInfo, twoDStrArray)

}

func (tIBuilder *testIBuilderErrPref) String() string {

	str := ""

	lenTwoDArray := len(tIBuilder.errPrefInfo)

	if lenTwoDArray == 0 {
		return str
	}

	for i := 0; i < lenTwoDArray; i++ {

		if len(tIBuilder.errPrefInfo[i][0]) == 0 {
			continue
		}

		str += tIBuilder.errPrefInfo[i][0]

		if len(tIBuilder.errPrefInfo[i][1]) == 0 {
			str += "\n"
		} else {
			str += " : " + tIBuilder.errPrefInfo[i][1] +
				"\n"
		}
	}

	return str
}

type testIStringerErrPref struct {
	locString string
}

func (tIStr *testIStringerErrPref) Set(str string) {
	tIStr.locString = str
}

func (tIStr *testIStringerErrPref) String() string {
	return tIStr.locString
}

type testIBasicErrPref struct {
	input [][2]string
}

func (testIBPref *testIBasicErrPref) GetEPrefStrings() [][2]string {

	return testIBPref.input
}

func (testIBPref *testIBasicErrPref) SetEPrefStrings(
	strArray [][2]string) error {

	if strArray == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strArray' is nil pointer!\n",
			"testIBasicErrPref.SetEPrefStrings()")
	}

	lenStrArray := len(strArray)

	if lenStrArray == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strArray' is a ZERO Length Array!\n",
			"testIBasicErrPref.SetEPrefStrings()")
	}

	testIBPref.input =
		make([][2]string, lenStrArray)

	elements :=
		copy(testIBPref.input, strArray)

	if elements != lenStrArray {
		return fmt.Errorf("%v\n"+
			"Error: Expected to copy %v elemetns.\n"+
			"Instead, only copied %v elements.\n",
			"testIBasicErrPref.SetEPrefStrings()",
			lenStrArray,
			elements)
	}

	return nil
}
