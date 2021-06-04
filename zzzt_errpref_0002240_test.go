package errpref

import "testing"

func TestErrPrefixDto_CopyInFromIBuilder_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_CopyInFromIBuilder"
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

	iBuilder := testIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

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

	ePDto2.SetMaxTextLineLen(40)

	err :=
		ePDto2.CopyInFromIBuilder(
			&iBuilder,
			funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.CopyInFromIBuilder()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}

func TestErrPrefixDto_CopyOutToIBuilder_000100(t *testing.T) {

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

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

	iBuilder := testIBuilderErrPref{}

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

	ePDto1.CopyOutToIBuilder(&iBuilder)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefStrings(
		iBuilder.GetEPrefStrings())

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}

func TestErrPrefixDto_ReplaceLastErrPrefix_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_ReplaceLastErrPrefix_000100()"

	ePDto := ErrPrefixDto{}.New()

	maxTextLineLen := 50

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

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	replacementMethod := "Rx14.ReplacementMethod()"
	replacementContext := "X=5 Y=900"

	err :=
		ePDto.ReplaceLastErrPrefix(
			replacementMethod,
			replacementContext,
			funcName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by ePDto.ReplaceLastErrPrefix()\n"+
			"Error= '%v'\n",
			funcName,
			err.Error())

		return
	}

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 14 {
		t.Errorf("%v\n"+
			"Error: Expected that collection length = '14'.\n"+
			"Instead, colletion length = '%v'\n",
			funcName,
			collectionLen)

		return
	}

	var lastErrPrefInfo ErrorPrefixInfo
	var collectionIsEmpty bool

	lastErrPrefInfo,
		collectionIsEmpty,
		err = ePDto.GetLastErrPrefix(funcName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by ePDto.GetLastErrPrefix(funcName).\n"+
			"Error= '%v'\n",
			funcName,
			err.Error())

		return
	}

	if collectionIsEmpty {
		t.Errorf("%v\n"+
			"Error: Expected that collectionIsEmpty=='false'.\n"+
			"Instead, collectionIsEmpty=='true'.\n",
			funcName)
		return
	}

	if lastErrPrefInfo.GetErrPrefixStr() != replacementMethod {
		t.Errorf("%v\n"+
			"Error: Expected Last Error Prefix == '%v'\n"+
			"Instead, Last Error Prefix == '%v'\n",
			funcName,
			replacementMethod,
			lastErrPrefInfo.GetErrPrefixStr())

		return
	}

	if lastErrPrefInfo.GetErrContextStr() != replacementContext {
		t.Errorf("%v\n"+
			"Error: Expected Last Error Context == '%v'\n"+
			"Instead, Last Error Context == '%v'\n",
			funcName,
			replacementContext,
			lastErrPrefInfo.GetErrContextStr())
	}

}

func TestErrPrefixDto_ReplaceLastErrPrefix_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_ReplaceLastErrPrefix_000200()"

	ePDto := ErrPrefixDto{}.New()

	maxTextLineLen := 50

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

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	replacementMethod := "Rx14.ReplacementMethod()"
	replacementContext := ""

	err :=
		ePDto.ReplaceLastErrPrefix(
			replacementMethod,
			replacementContext,
			funcName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by ePDto.ReplaceLastErrPrefix()\n"+
			"Error= '%v'\n",
			funcName,
			err.Error())

		return
	}

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 14 {
		t.Errorf("%v\n"+
			"Error: Expected that collection length = '14'.\n"+
			"Instead, colletion length = '%v'\n",
			funcName,
			collectionLen)

		return
	}

	var lastErrPrefInfo ErrorPrefixInfo
	var collectionIsEmpty bool

	lastErrPrefInfo,
		collectionIsEmpty,
		err = ePDto.GetLastErrPrefix(funcName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by ePDto.GetLastErrPrefix(funcName).\n"+
			"Error= '%v'\n",
			funcName,
			err.Error())

		return
	}

	if collectionIsEmpty {
		t.Errorf("%v\n"+
			"Error: Expected that collectionIsEmpty=='false'.\n"+
			"Instead, collectionIsEmpty=='true'.\n",
			funcName)
		return
	}

	if lastErrPrefInfo.GetErrPrefixStr() != replacementMethod {
		t.Errorf("%v\n"+
			"Error: Expected Last Error Prefix == '%v'\n"+
			"Instead, Last Error Prefix == '%v'\n",
			funcName,
			replacementMethod,
			lastErrPrefInfo.GetErrPrefixStr())

		return
	}

	if lastErrPrefInfo.GetErrContextStr() != replacementContext {
		t.Errorf("%v\n"+
			"Error: Expected Last Error Context == '%v'\n"+
			"Instead, Last Error Context == '%v'\n",
			funcName,
			replacementContext,
			lastErrPrefInfo.GetErrContextStr())
	}

}

func TestErrPrefixDto_ReplaceLastErrPrefix_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_ReplaceLastErrPrefix_000300()"

	ePDto := ErrPrefixDto{}.New()

	maxTextLineLen := 50

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

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	replacementMethod := ""
	replacementContext := ""

	err :=
		ePDto.ReplaceLastErrPrefix(
			replacementMethod,
			replacementContext,
			funcName)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from ePDto.ReplaceLastErrPrefix() because\n"+
			"input parameter 'replacementMethod' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			funcName)
	}
}

func TestErrPrefixDto_SetIBasic_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBasic_000100()"

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

	ePDto := ErrPrefixDto{}

	err =
		ePDto.SetIBasic(
			&iBasicPref,
			funcName)

	if err != nil {
		t.Errorf("Error returned by ePDto.SetIBasic()\n"+
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

func TestErrPrefixDto_SetIBasic_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBasic_000200()"

	iBasicPref := testIBasicErrPref{}

	var err error

	ePDto := ErrPrefixDto{}

	err =
		ePDto.SetIBasic(
			&iBasicPref,
			funcName)

	if err != nil {
		t.Errorf("Error returned by ePDto.SetIBasic()\n"+
			"%v\n",
			err.Error())

		return
	}

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("ERROR:\n"+
			"Expected final Collection Length == 0\n"+
			"HOWEVER, final Collection Length IS NOT ZERO!!\n"+
			"Final Colleciton Length='%v'\n",
			collectionLen)
	}

}

func TestErrPrefixDto_SetIBuilder_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBuilder_000100"

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

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefOld(bogusStr)

	ePDto1.SetMaxTextLineLen(40)

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

	iBuilder := testIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefStrings(twoDSlice)

	err := ePDto1.SetIBuilder(
		&iBuilder,
		funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetIBuilder()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}

func TestErrPrefixDto_SetIBuilder_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBuilder_000200"

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

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefOld(bogusStr)

	ePDto1.SetMaxTextLineLen(40)

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

	iBuilder := testIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefStrings(twoDSlice)

	err := ePDto1.SetIBuilder(
		nil,
		funcName)

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return from ePDto1.SetIBuilder()\n" +
			"because iEPref is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
	}

}

func TestErrPrefixDto_SetIBuilder_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBuilder_000300"

	iBuilder := testIBuilderErrPref{}

	ePDto := ErrPrefixDto{}

	err := ePDto.SetIBuilder(
		&iBuilder,
		funcName)

	if err != nil {
		t.Errorf("Error:\n"+
			"Error return from ePDto.SetIBuilder()\n"+
			"Error='%v'\n",
			err.Error())
	}

}
