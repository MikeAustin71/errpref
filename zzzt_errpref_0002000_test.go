package errpref

import "testing"

func TestErrPrefixDto_AddEPrefCollectionStr_000100(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}.New()

	numOfItemsAdded := ePDto2.AddEPrefCollectionStr(
		initialStr)

	if numOfItemsAdded != ePDto.GetEPrefCollectionLen() {
		t.Errorf("Error: Number of Items Added NOT equal to Collection Items!\n"+
			"numOfItemsAdded=='%v'\n"+
			"ePDto.GetEPrefCollectionLen()=='%v'\n",
			numOfItemsAdded,
			ePDto.GetEPrefCollectionLen())

		return
	}

	ePDto2.SetMaxTextLineLen(40)

	if ePDto.GetEPrefCollectionLen() != ePDto2.GetEPrefCollectionLen() {
		t.Errorf("Expected ePDto.GetEPrefCollectionLen() == ePDto2.GetEPrefCollectionLen()\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto.GetEPrefCollectionLen()=='%v'\n"+
			"ePDto2.GetEPrefCollectionLen()=='%v'\n",
			ePDto.GetEPrefCollectionLen(),
			ePDto2.GetEPrefCollectionLen())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
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

func TestErrPrefixDto_DeleteLastErrPrefix_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_DeleteLastErrPrefix_000100"

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

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 14 {
		t.Errorf("%v\n"+
			"Error: Expected initial collection length to equal '14'\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)

		return
	}

	isCollectionEmpty := ePDto.DeleteLastErrPrefix()

	if isCollectionEmpty {
		t.Error("After deleting the Last Error Prefix\n" +
			"Information object, the returned flag shows that\n" +
			"the collection is empty. This is an error.\n" +
			"THE COLLECTION SHOULD NOT BE EMPTY!\n")
		return
	}

	var lastErrInfo ErrorPrefixInfo
	var err error

	lastErrInfo,
		isCollectionEmpty,
		err = ePDto.GetLastErrPrefix(funcName)

	if err != nil {
		t.Errorf("Error returned by ePDto.GetLastErrPrefix(funcName).\n"+
			"Error= '%v'\n",
			err.Error())
		return
	}

	expectedLastErrorPrefix := twoDSlice[12][0]

	if expectedLastErrorPrefix !=
		lastErrInfo.GetErrPrefixStr() {
		t.Errorf("Error: Expected new Last Error Prefix ='%v'\n"+
			"Instead, Last Error Prefix = '%v'\n",
			expectedLastErrorPrefix,
			lastErrInfo.GetErrPrefixStr())

		return
	}

	if !lastErrInfo.GetIsLastIndex() {
		t.Error("ERROR: Expected final lastErrInfo.GetIsLastIndex()=='true'\n" +
			"Instead, lastErrInfo.GetIsLastIndex()='false'\n")
		return
	}

	collectionLen = ePDto.GetEPrefCollectionLen()

	if collectionLen != 13 {
		t.Errorf("%v\n"+
			"Error: Expected final collection length to equal '13'\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)
	}

}

func TestErrPrefixDto_DeleteLastErrPrefix_000200(t *testing.T) {

	funcName := "TestErrPrefixDto_DeleteLastErrPrefix_000200"

	ePDto := ErrPrefixDto{}.New()

	maxTextLineLen := 50

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 1 {
		t.Errorf("%v\n"+
			"Error: Expected initial collection length to equal '1'\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)

		return
	}

	isCollectionEmpty := ePDto.DeleteLastErrPrefix()

	if !isCollectionEmpty {
		t.Error("After deleting the Last Error Prefix\n" +
			"Information object in the collection, the returned\n" +
			"flag shows that the collection is NOT empty.\n" +
			"This is an error. THE COLLECTION SHOULD BE EMPTY!\n")
		return
	}

	var err error

	_,
		isCollectionEmpty,
		err = ePDto.GetLastErrPrefix(funcName)

	if err != nil {
		t.Errorf("Error returned by ePDto.GetLastErrPrefix(funcName).\n"+
			"Error= '%v'\n",
			err.Error())
		return
	}

	if isCollectionEmpty == false {
		t.Error("ERROR: After deleting the last Error Prefix\n" +
			"Information object, the call to ePDto.GetLastErrPrefix(funcName)\n" +
			"should have returned isCollectionEmpty=='true'.\n" +
			"Instead, isCollectionEmpty=='false'\n")

		return
	}

	collectionLen = ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("%v\n"+
			"Error: Expected final collection length to equal zero.\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)
	}

}

func TestErrPrefixDto_DeleteLastErrPrefix_000300(t *testing.T) {

	funcName := "TestErrPrefixDto_DeleteLastErrPrefix_000300"

	ePDto := ErrPrefixDto{}.New()

	maxTextLineLen := 50

	ePDto.SetMaxTextLineLen(maxTextLineLen)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("%v\n"+
			"Error: Expected initial collection length to equal zero.\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)

		return
	}

	isCollectionEmpty := ePDto.DeleteLastErrPrefix()

	if !isCollectionEmpty {
		t.Error("After deleting the Last Error Prefix\n" +
			"Information object in an empty collection, the returned\n" +
			"flag shows that the collection is NOT empty.\n" +
			"This is an error. THE COLLECTION SHOULD BE EMPTY!\n")
		return
	}

	var err error

	_,
		isCollectionEmpty,
		err = ePDto.GetLastErrPrefix(funcName)

	if err != nil {
		t.Errorf("Error returned by ePDto.GetLastErrPrefix(funcName).\n"+
			"Error= '%v'\n",
			err.Error())
		return
	}

	if isCollectionEmpty == false {
		t.Error("ERROR: After deleting the last Error Prefix\n" +
			"Information object in an empty collection, the call to\n" +
			"ePDto.GetLastErrPrefix(funcName) should have returned\n" +
			"isCollectionEmpty=='true'.\n" +
			"Instead, isCollectionEmpty=='false'\n")

		return
	}

	collectionLen = ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("%v\n"+
			"Error: Expected final collection length to equal zero.\n"+
			"Instead, collection length = '%v'\n",
			funcName,
			collectionLen)
	}

}

func TestErrPrefixDto_CopyPtr_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto2 := ePDto.CopyPtr()

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
	}

}

func TestErrPrefixDto_CopyIn_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		&ePDto,
		"TestErrPrefixDto_CopyIn_000100")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
	}
}

func TestErrPrefixDto_CopyIn_000200(t *testing.T) {
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

	ePDto.SetLeftMarginLength(5)

	ePDto.SetLeftMarginChar('*')

	ePDto.SetIsLastLineTermWithNewLine(true)

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		&ePDto,
		"TestErrPrefixDto_CopyIn_000200")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if ePDto2.GetLeftMarginLength() != 5 {
		t.Errorf("Error: After Copy expected Left Margin Length==5.\n"+
			"Instead, Left Margin Length = %v\n",
			ePDto2.GetLeftMarginLength())
		return
	}

	if ePDto2.GetLeftMarginChar() != '*' {
		t.Errorf("Error: After Copy expected Left Margin Char=='*'.\n"+
			"Instead, Left Margin Char == %v\n",
			string(ePDto2.GetLeftMarginChar()))
		return
	}

	if !ePDto2.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Copy expected GetIsLastLineTerminatedWithNewLine()=='true'.\n" +
			"Instead, GetIsLastLineTerminatedWithNewLine()=='false'.\n")
		return
	}

	if ePDto2.GetMaxTextLineLen() != 40 {
		t.Errorf("Error: After Copy expected GetMaxTextLineLen=='40'.\n"+
			"Instead, GetMaxTextLineLen == %v\n",
			ePDto2.GetMaxTextLineLen())
		return
	}

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
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

func TestErrPrefixDto_CopyIn_000300(t *testing.T) {
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

	ePDto.SetLeftMarginLength(5)

	ePDto.SetLeftMarginChar('*')

	ePDto.SetIsLastLineTermWithNewLine(true)

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		nil,
		"TestErrPrefixDto_CopyIn_000300")

	if err == nil {
		t.Error("Error:\n" +
			"Expected an error return from ePDto2.CopyIn()\n" +
			"because inComingErrPrefixDto is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestErrPrefixDto_CopyIn_000400(t *testing.T) {

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		nil,
		"TestErrPrefixDto_CopyIn_000200")

	if err == nil {
		t.Error("Error:\n" +
			"Expected error return from ePDto2.CopyIn(&ePDto)\n" +
			"because ePDto is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
	}
}

func TestErrPrefixDto_CopyOut_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto2,
		err := ePDto.CopyOut("TestErrPrefixDto_CopyOut_000100")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
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

func TestErrPrefixDto_CopyOut_000200(t *testing.T) {
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

	ePDto.SetLeftMarginLength(5)

	ePDto.SetLeftMarginChar('*')

	ePDto.SetIsLastLineTermWithNewLine(true)

	ePDto2,
		err := ePDto.CopyOut("TestErrPrefixDto_CopyOut_000200")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if ePDto2.GetLeftMarginLength() != 5 {
		t.Errorf("Error: After Copy expected Left Margin Length==5.\n"+
			"Instead, Left Margin Length = %v\n",
			ePDto2.GetLeftMarginLength())
		return
	}

	if ePDto2.GetLeftMarginChar() != '*' {
		t.Errorf("Error: After Copy expected Left Margin Char=='*'.\n"+
			"Instead, Left Margin Char == %v\n",
			string(ePDto2.GetLeftMarginChar()))
		return
	}

	if !ePDto2.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Copy expected GetIsLastLineTerminatedWithNewLine()=='true'.\n" +
			"Instead, GetIsLastLineTerminatedWithNewLine()=='false'.\n")
		return
	}

	if ePDto2.GetMaxTextLineLen() != 40 {
		t.Errorf("Error: After Copy expected GetMaxTextLineLen=='40'.\n"+
			"Instead, GetMaxTextLineLen == %v\n",
			ePDto2.GetMaxTextLineLen())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
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

func TestErrPrefixDto_Empty_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetLeftMarginLength(3)
	ePDto.SetLeftMarginChar('*')
	ePDto.SetMaxTextLineLen(40)
	ePDto.SetIsLastLineTermWithNewLine(true)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen == 0 {
		t.Error("Error: After Initialization with data,\n" +
			"Error Prefix Collection Length==0\n")
		return
	}

	ePDto.Empty()

	if ePDto.GetLeftMarginLength() != 0 {
		t.Errorf("Error: After Empty() expected Left Margin\n"+
			"Length == 0.\n"+
			"Instead, ePDto.GetLeftMarginLength()=='%v'\n",
			ePDto.GetLeftMarginLength())
		return
	}

	if ePDto.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: After Empty() expected Collection\n"+
			"Length == 0.\n"+
			"Instead, ePDto.GetEPrefCollectionLen()=='%v'\n",
			ePDto.GetEPrefCollectionLen())
		return
	}

	if ePDto.GetLeftMarginChar() != 0 {
		t.Errorf("Error: After Empty() expected Left Margin\n"+
			"Character == 0.\n"+
			"Instead, ePDto.GetLeftMarginChar()=='%v'\n",
			string(ePDto.GetLeftMarginChar()))
		return
	}

	if ePDto.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Empty() expected isLastLineTerminatedWithNewLine==false\n" +
			"Instead, isLastLineTerminatedWithNewLine==true")
	}

}

func TestErrPrefixDto_EmptyEPrefCollection_000100(t *testing.T) {

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

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen == 0 {
		t.Error("Error: After Initialization with data,\n" +
			"Error Prefix Collection Length==0\n")
		return
	}

	ePDto.EmptyEPrefCollection()

	collectionLen = ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("Error: After EmptyEPrefCollection(),\n"+
			"Error Prefix Collection Length!=0\n"+
			"Error Prefix Collection Length=='%v'\n",
			ePDto.GetEPrefCollectionLen())

	}

}

func TestErrPrefixDto_Equal_000100(t *testing.T) {

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

	ePDto.SetCtx("A!=B")

	actualStr := ePDto.String()

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

	ePDto2 := ePDto.Copy()

	if !ePDto.Equal(&ePDto2) {

		t.Error("Error:\n" +
			"Expected ePDto to Equal ePDto2.\n" +
			"However, THEY ARE NOT EQUAL!\n")

		return
	}

	ePDto2.SetMaxTextLineLen(60)

	if ePDto.Equal(&ePDto2) {

		t.Error("Error:\n" +
			"Expected ePDto to be UNEqual to ePDto2.\n" +
			"However, THEY ARE EQUAL!\n")

		return
	}
}

func TestErrPrefixDto_GetDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_GetDelimiters_000100"

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

	errPrefDelims := ePDto.GetInputStringDelimiters()

	errPrefDelimsOne := ErrPrefixDelimiters{}

	err := errPrefDelimsOne.CopyIn(
		&errPrefDelims,
		funcName)

	if err != nil {
		t.Errorf("Error Returned by errPrefDelimsOne.CopyIn()\n"+
			"%v\n",
			err.Error())

		return
	}

	ep := ErrPref{}

	epDelims := ep.GetDelimiters()

	var errPrefDelimsTwo ErrPrefixDelimiters

	errPrefDelimsTwo,
		err = epDelims.CopyOut(
		funcName)

	if err != nil {
		t.Errorf("Error Returned by epDelims.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	errPrefDelimsOneStr := errPrefDelimsOne.String()

	errPrefDelimsTwoStr := errPrefDelimsTwo.String()

	if !errPrefDelimsOne.Equal(&errPrefDelimsTwo) {
		t.Errorf("Error: Expected errPrefDelimsOne==errPrefDelimTwo\n"+
			"HOWEVER, THE TWO ARE NOT EQUAL\n"+
			"errPrefDelimsOne=\n%v\n"+
			"errPrefDelimsTwo=\n%v\n",
			errPrefDelimsOneStr,
			errPrefDelimsTwoStr)

		return
	}

	if errPrefDelimsOneStr != errPrefDelimsTwoStr {
		t.Errorf("Error: Expected errPrefDelimsOneStr==errPrefDelimsTwoStr\n"+
			"HOWEVER, THE TWO ARE NOT EQUAL\n"+
			"errPrefDelimsOne=\n%v\n"+
			"errPrefDelimsTwo=\n%v\n",
			errPrefDelimsOneStr,
			errPrefDelimsTwoStr)
	}

}

func TestErrPrefixDto_GetEPrefCollection(t *testing.T) {

	ePDto := ErrPrefixDto{}

	ePrefCol := ePDto.GetEPrefCollection()

	if len(ePrefCol) != 0 {
		t.Errorf("ERROR:\n"+
			"Expected collection length == 0, because the\n"+
			"ePDto instance was created with ErrPrefixDto{}\n"+
			"HOWEVER, THE COLLECTION LENGTH IS NOT EQUAL TO ZERO!!\n"+
			"len(ePrefCol)=='%v'\n",
			len(ePrefCol))
	}

}

func TestErrPrefixDto_GetEPrefStrings_000100(t *testing.T) {

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

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	err := ePDto.SetIEmpty(
		twoDSlice,
		"TestErrPrefixDto_GetEPrefStrings_000100")

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.SetIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	twoDSlice2 := ePDto.GetEPrefStrings()

	if len(twoDSlice) != len(twoDSlice2) {
		t.Errorf("ERROR: Expected len(twoDSlice) == len(twoDSlice2)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"len(twoDSlice) = '%v'\n"+
			"len(twoDSlice2)= '%v'\n",
			len(twoDSlice),
			len(twoDSlice2))
		return
	}

	for i := 0; i < len(twoDSlice2); i++ {

		if twoDSlice[i][0] != twoDSlice2[i][0] {
			t.Errorf("ERROR: twoDSlice[i][0] != twoDSlice2[i][0]\n"+
				"twoDSlice[i][0] ='%v'\n"+
				"twoDSlice2[i][0]='%v'\n",
				twoDSlice[i][0],
				twoDSlice2[i][0])
			return
		}

		if twoDSlice[i][1] != twoDSlice2[i][1] {
			t.Errorf("ERROR: twoDSlice[i][1] != twoDSlice2[i][1]\n"+
				"twoDSlice[i][1] ='%v'\n"+
				"twoDSlice2[i][1]='%v'\n",
				twoDSlice[i][1],
				twoDSlice2[i][1])
			return
		}

	}

}

func TestErrPrefixDto_GetMaxTextLineLen_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}

	maxTextLineLen := ePDto.GetMaxTextLineLen()

	defaultMaxLineLen := ErrPref{}.GetMaxErrPrefTextLineLength()

	if maxTextLineLen != defaultMaxLineLen {
		t.Errorf("ERROR:"+
			"Expected call to ePDto.GetMaxTextLineLen() to yield\n"+
			"a value of '%v' because 'ePDto' is invalid.\n"+
			"%v is the default Maximum Text Line Length.\n"+
			"HOWEVER, THE RETURNED VALUE WAS NOT THE DEFAULT\n"+
			"Maximum Text Line Length!\n"+
			"The return value was '%v'\n",
			defaultMaxLineLen,
			defaultMaxLineLen,
			maxTextLineLen)
	}

}

func TestErrPrefixDto_IsValidInstance_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}

	if ePDto.IsValidInstance() {
		t.Error("ERROR:\n" +
			"Expected ePDto.IsValidInstance() to return 'false'\n" +
			"because the ErrPrefixDto string delimiters are empty\n" +
			"and therefore invalid.\n" +
			"HOWEVER, THE RETURN VALUE WAS 'true'!!!\n")
	}

}

func TestErrPrefixDto_IsValidInstance_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	if !ePDto.IsValidInstance() {
		t.Error("ERROR:\n" +
			"Expected ePDto.IsValidInstance() to return a value of 'true'\n" +
			"because the ErrPrefixDto was created with New() and the\n" +
			"instance is therefore valid.\n" +
			"HOWEVER, THE RETURN VALUE WAS 'false'!!!\n")
	}

}

func TestErrPrefixDto_IsValidInstanceError_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}

	err := ePDto.IsValidInstanceError("")

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected ePDto.IsValidInstanceError(\"\") to return an error\n" +
			"because the string delimiters are empty and invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!!")
	}

}

func TestErrPrefixDto_IsValidInstanceError_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	err := ePDto.IsValidInstanceError("")

	if err != nil {
		t.Errorf("ERROR:\n"+
			"Expected ePDto.IsValidInstanceError(\"\") to return an\n"+
			"error value 'nil' because the 'ePDto' instance is valid.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n"+
			"Error='%v'\n",
			err.Error())
	}

}

func TestErrPrefixDto_MergeErrPrefixDto_000100(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

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

	ePDto2.SetEPrefOld(bogusStr)

	expectedCompositeStr := initialStr +
		"\n" +
		bogusStr

	expectedEPDto := ErrPrefixDto{}.New()

	expectedEPDto.SetMaxTextLineLen(40)

	expectedEPDto.SetEPrefOld(expectedCompositeStr)

	ePDto.MergeErrPrefixDto(
		&ePDto2)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedEPDto.String()),
		false)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	if !expectedEPDto.Equal(&ePDto) {
		t.Errorf("Error: Expected expectedEPDto==ePDto.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			expectedStr,
			ePDtoStr)
		return
	}

	if expectedStr != ePDtoStr {
		t.Errorf("Error: Expected expectedStr==ePDtoStr.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"expectedStr=\n%v\n\nePDtoStr=\n%v\n\n",
			expectedStr,
			ePDtoStr)
	}

}

func TestErrPrefixDto_MergeErrPrefixDto_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto2 := ErrPrefixDto{}

	ePDto.MergeErrPrefixDto(&ePDto2)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("Error:\n"+
			"Expected collectionLen == 0 after processing\n"+
			"an empty ErrPrefixDto instance.\n"+
			"HOWEVER, collectionLen > 0\n"+
			"collectionLen='%v'\n",
			collectionLen)
	}

}

func TestErrPrefixDto_MergeErrPrefixDto_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

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

	ePDto.MergeErrPrefixDto(nil)

	newCollectionLen := ePDto.GetEPrefCollectionLen()

	if newCollectionLen != 14 {
		t.Errorf("Error:\n"+
			"Expected newCollectionLen == 14 (the old value) "+
			"after processing ePDto.MergeErrPrefixDto(nil).\n"+
			"HOWEVER, newCollectionLen != 14 \n"+
			"newCollectionLen='%v'\n",
			newCollectionLen)
	}

}

func TestErrPrefixDto_Multiple_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew() : something->newSomething\\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\\n" +
			"Tx13.SomeFabulousAndComplexStuff()\\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxTextLineLenToDefault()
}

func TestErrPrefixDto_Multiple_000200(t *testing.T) {

	initialStr :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 40-Characters
	ePDto.SetMaxTextLineLen(40)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto.SetEPrefCtx(
		"Tx15.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	expectedStr :=
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

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxTextLineLenToDefault()
}

func TestErrPrefixDto_Multiple_000300(t *testing.T) {

	initialStr :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 40-Characters
	ePDto.SetMaxTextLineLen(40)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPrefCtx(
		"Tx9.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	ePDto.SetEPref("Tx10.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx11.X()")

	ePDto.SetEPrefCtx(
		"Tx12.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx13.TryASalad()",
		"")

	ePDto.SetEPref("Tx14.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx15.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedStr :=
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

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxTextLineLenToDefault()
}
