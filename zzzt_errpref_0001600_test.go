package errpref

import "testing"

func TestEPrefixLineLenCalc_CopyIn_000100(t *testing.T) {

	funcName := "TestEPrefixLineLenCalc_CopyIn_000100"

	ePrefLineLenOne := EPrefixLineLenCalc{}.Ptr()

	ePrefInfoOne := ErrorPrefixInfo{}.Ptr()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	err :=
		ePrefLineLenOne.SetErrPrefixInfo(
			ePrefInfoOne,
			funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenOne.SetErrPrefixInfo()\n"+
			"%v\n",
			err.Error())

		return
	}

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDelimiters := ePDto.GetInputStringDelimiters()

	err =
		ePrefLineLenOne.SetEPrefDelimiters(
			ePDelimiters,
			funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenOne.SetEPrefDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	maxTextLineLen := uint(70)

	ePrefLineLenOne.SetMaxErrStringLength(maxTextLineLen)

	currentLineStr := "Tx1.Something() - Tx2.SomethingElse()"

	currentLineStrLen := uint(len(currentLineStr))

	ePrefLineLenOne.SetCurrentLineStr(currentLineStr)

	if !ePrefLineLenOne.IsValidInstance() {
		t.Error("ERROR: Expected ePrefLineLenOne.IsValidInstance() " +
			"== 'true'\n" +
			"Instead, ePrefLineLenOne.IsValidInstance() " +
			"== 'false'\n")

		return
	}

	ePrefLineLenTwo := EPrefixLineLenCalc{}.New()

	err =
		ePrefLineLenTwo.CopyIn(
			ePrefLineLenOne,
			funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenTwo.CopyIn()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefLineLenTwo.Equal(ePrefLineLenOne) {
		t.Error("ERROR: Expected ePrefLineLenOne==ePrefLineLenTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")

		return
	}

	if ePrefLineLenTwo.GetMaxErrStringLength() != maxTextLineLen {
		t.Errorf("ERROR: "+
			"Expected ePrefLineLenTwo.GetMaxErrStringLength() == "+
			"maxTextLineLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefLineLenTwo.GetMaxErrStringLength()='%v'\n"+
			"maxTextLineLen='%v'\n",
			ePrefLineLenTwo.GetMaxErrStringLength(),
			maxTextLineLen)

		return
	}

	if !ePrefLineLenTwo.IsErrPrefixLastIndex() {
		t.Error("ERROR: " +
			"Expected ePrefLineLenTwo.IsErrPrefixLastIndex()=='true'\n" +
			"Instead, ePrefLineLenTwo.IsErrPrefixLastIndex()=='false'\n")

		return
	}

	remainingLineLen := maxTextLineLen - currentLineStrLen

	if ePrefLineLenTwo.GetRemainingLineLength() != remainingLineLen {
		t.Errorf("ERROR: "+
			"Expected ePrefLineLenTwo.GetRemainingLineLength() "+
			"== remainingLineLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefLineLenTwo.GetRemainingLineLength()='%v'\n"+
			"remainingLineLen='%v'\n",
			ePrefLineLenTwo.GetRemainingLineLength(),
			remainingLineLen)

	}

}

func TestEPrefixLineLenCalc_CopyOut_000100(t *testing.T) {

	funcName := "TestEPrefixLineLenCalc_CopyOut_000100"

	ePrefInfoOne := ErrorPrefixInfo{}.Ptr()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	ePrefLineLenOne := EPrefixLineLenCalc{}.Ptr()

	err :=
		ePrefLineLenOne.SetErrPrefixInfo(
			ePrefInfoOne,
			funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenOne.SetErrPrefixInfo()\n"+
			"%v\n",
			err.Error())

		return
	}

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDelimiters := ePDto.GetInputStringDelimiters()

	err =
		ePrefLineLenOne.SetEPrefDelimiters(
			ePDelimiters,
			funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenOne.SetEPrefDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	maxTextLineLen := uint(70)

	ePrefLineLenOne.SetMaxErrStringLength(maxTextLineLen)

	currentLineStr := "Tx1.Something()"

	ePrefLineLenOne.SetCurrentLineStr(currentLineStr)

	err = ePrefLineLenOne.IsValidInstanceError(
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	var ePrefLineLenTwo EPrefixLineLenCalc

	ePrefLineLenTwo,
		err = ePrefLineLenOne.CopyOut(
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefLineLenTwo.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefLineLenTwo.Equal(ePrefLineLenOne) {
		t.Error("ERROR: Expected ePrefLineLenOne==ePrefLineLenTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")

		return
	}

	if ePrefLineLenTwo.ErrorPrefixIsEmpty() {
		t.Error("ERROR: " +
			"Expected ePrefLineLenTwo.ErrorPrefixIsEmpty()=='true'\n" +
			"Instead, ePrefLineLenTwo.ErrorPrefixIsEmpty()=='false'\n")

		return
	}

	if ePrefLineLenTwo.ErrorContextIsEmpty() {
		t.Error("ERROR: " +
			"Expected ePrefLineLenTwo.ErrorContextIsEmpty()=='true'\n" +
			"Instead, ePrefLineLenTwo.ErrorContextIsEmpty()=='false'\n")
	}

}
