package errpref

import (
	"strings"
	"testing"
)

func TestErrPrefixDto_SetIEmpty_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(expectedStr)

	bogusStr :=
		"Xt1.Something() - Xt2.SomethingElse() - Xt3.DoSomething()\n" +
			"Xt4() - Xt5() - Xt6.DoSomethingElse()\n" +
			"Xt7.TrySomethingNew() : something->newSomething\n" +
			"Xt8.TryAnyCombination() - Xt9.TryAHammer() : x->y - Xt10.X()\n" +
			"Xt11.TryAnything() - Xt12.TryASalad()\n" +
			"Xt13.SomeFabulousAndComplexStuff()\n" +
			"Xt14.MoreAwesomeGoodness\n" +
			"Xt15.MustardSandwiches()\n" +
			"Xt16.TomatoSandwiches()\n" +
			"Xt17.Benitos() : Z=4 Y=3 X=9"

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetEPrefOld(bogusStr)

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

	err := ePDto2.SetIEmpty(
		twoDSlice,
		"TestErrPrefixDto_SetIEmpty_000100")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew()\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness()"

	ePDto.SetEPrefOld(expectedStr)

	ePDto.SetMaxTextLineLen(40)

	bogusStr :=
		"Xt1.Something() - Xt2.SomethingElse() - Xt3.DoSomething()\n" +
			"Xt4() - Xt5() - Xt6.DoSomethingElse()\n" +
			"Xt7.TrySomethingNew() : something->newSomething\n" +
			"Xt8.TryAnyCombination() - Xt9.TryAHammer() : x->y - Xt10.X()\n" +
			"Xt11.TryAnything() - Xt12.TryASalad()\n" +
			"Xt13.SomeFabulousAndComplexStuff()\n" +
			"Xt14.MoreAwesomeGoodness\n" +
			"Xt15.MustardSandwiches()\n" +
			"Xt16.TomatoSandwiches()\n" +
			"Xt17.Benitos() : Z=4 Y=3 X=9"

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetEPrefOld(bogusStr)

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

	err := ePDto2.SetIEmpty(
		oneDSlice,
		"TestErrPrefixDto_SetIEmpty_000200")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000300(t *testing.T) {

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

	bogusStr :=
		"Xt1.Something() - Xt2.SomethingElse() - Xt3.DoSomething()\n" +
			"Xt4() - Xt5() - Xt6.DoSomethingElse()\n" +
			"Xt7.TrySomethingNew() : something->newSomething\n" +
			"Xt8.TryAnyCombination() - Xt9.TryAHammer() : x->y - Xt10.X()\n" +
			"Xt11.TryAnything() - Xt12.TryASalad()\n" +
			"Xt13.SomeFabulousAndComplexStuff()\n" +
			"Xt14.MoreAwesomeGoodness\n" +
			"Xt15.MustardSandwiches()\n" +
			"Xt16.TomatoSandwiches()\n" +
			"Xt17.Benitos() : Z=4 Y=3 X=9"

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetEPrefOld(bogusStr)

	err := ePDto2.SetIEmpty(
		initialStr,
		"TestErrPrefixDto_SetIEmpty_000300")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000400(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}

	funcName :=
		"TestErrPrefixDto_SetIEmpty_000400()"

	ePDtoCopy,
		err = ePDto.CopyOut(funcName)

	if err != nil {
		t.Errorf("Error from ePDto.CopyOut(funcName)\n"+
			"%v\n", err.Error())
		return
	}

	err = ePDto2.SetIEmpty(
		ePDtoCopy,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000500(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}

	funcName :=
		"TestErrPrefixDto_SetIEmpty_000500()"

	ePDtoCopy = ePDto.CopyPtr()

	err = ePDto2.SetIEmpty(
		ePDtoCopy,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000600(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIEmpty_000600()"

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(expectedStr)

	err := ePDto.SetIEmpty(
		nil,
		funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	if ePDto.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: Expected Collection Length==0\n"+
			"Instead Collection Length=='%v'\n",
			ePDto.GetEPrefCollectionLen())
	}

}

func TestErrPrefixDto_SetIEmpty_000700(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		twoDSlice,
		"TestErrPrefixDto_SetIEmpty_000700")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000800(t *testing.T) {

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

	iStringerEPref.Set(initialStr)

	var err error

	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		&iStringerEPref,
		"TestErrPrefixDto_SetIEmpty_000800")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_000900(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		sb,
		"TestErrPrefixDto_SetIEmpty_000900")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_001000(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		sbPtr,
		"TestErrPrefixDto_SetIEmpty_001000")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
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

	if !ePDto.Equal(&ePDto2) {
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

func TestErrPrefixDto_SetIEmpty_001100(t *testing.T) {

	var sbPtr *strings.Builder

	var err error

	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		sbPtr,
		"TestErrPrefixDto_SetIEmpty_001100")

	if err == nil {
		t.Error("Expected an error return from ErrPrefixDto{}.SetIEmpty(sbPtr)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestErrPrefixDto_SetIEmpty_001200(t *testing.T) {

	var err error
	var ePDto1 *ErrPrefixDto
	ePDto2 := ErrPrefixDto{}

	err = ePDto2.SetIEmpty(
		ePDto1,
		"TestErrPrefixDto_SetIEmpty_001200")

	if err == nil {
		t.Error("Expected an error return from ePDto2.SetIEmpty(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}
