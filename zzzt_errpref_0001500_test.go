package errpref

import "testing"

func TestErrorPrefixInfo_CopyIn_000100(t *testing.T) {

	ePrefInfoOne := ErrorPrefixInfo{}.New()

	testMethod := "TxDoSomethingNow()"
	lenTestMethod := uint(len(testMethod))

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	lenErrContext := uint(len(errContext))

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	ePrefInfoTwo := ErrorPrefixInfo{}.Ptr()

	err :=
		ePrefInfoTwo.CopyIn(
			&ePrefInfoOne,
			"TestErrorPrefixInfo_CopyIn()")

	if err != nil {
		t.Errorf("ERROR returned by ePrefInfoTwo.CopyIn(&ePrefInfoOne)\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefInfoTwo.IsValidInstance() {
		t.Error("ERROR: ePrefInfoTwo.IsValidInstance() " +
			"returned 'false'\n")

		return
	}

	if !ePrefInfoOne.Equal(ePrefInfoTwo) {
		t.Error("ERROR: Expected ePrefInfoOne==ePrefInfoTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!\n ")
		return
	}

	if !ePrefInfoTwo.GetIsPopulated() {
		t.Error("ERROR: Expected ePrefInfoTwo.GetIsPopulated()=='true'\n" +
			"HOWEVER, IT IS FALSE!!\n ")
		return
	}

	if !ePrefInfoTwo.GetIsLastIndex() {
		t.Error("ERROR: Expected ePrefInfoTwo.GetIsLastIndex()=='true'\n" +
			"HOWEVER, IT IS FALSE!!\n ")
		return
	}

	if !ePrefInfoTwo.GetIsFirstIndex() {
		t.Error("ERROR: Expected ePrefInfoTwo.GetIsFirstIndex()=='true'\n" +
			"HOWEVER, IT IS FALSE!!\n ")
		return
	}

	if ePrefInfoTwo.GetLengthErrPrefixStr() != lenTestMethod {
		t.Errorf("ERROR: ePrefInfoTwo.GetLengthErrPrefixStr() "+
			"!= lenTestMethod\n"+
			"ePrefInfoTwo.GetLengthErrPrefixStr()='%v'\n"+
			"lenTestMethod='%v'\n",
			ePrefInfoTwo.GetLengthErrPrefixStr(),
			lenTestMethod)
		return
	}

	if ePrefInfoTwo.GetLengthErrContextStr() != lenErrContext {
		t.Errorf("ERROR: ePrefInfoTwo.GetLengthErrContextStr() "+
			"!= lenErrContext\n"+
			"ePrefInfoTwo.GetLengthErrContextStr()='%v'\n"+
			"lenErrContext='%v'\n",
			ePrefInfoTwo.GetLengthErrContextStr(),
			lenTestMethod)
		return
	}

	if ePrefInfoTwo.GetErrPrefixStr() != testMethod {
		t.Errorf("ERROR: ePrefInfoTwo.GetErrPrefixStr() "+
			"!= testMethod\n"+
			"ePrefInfoTwo.GetErrPrefixStr()='%v'\n"+
			"testMethod='%v'\n",
			ePrefInfoTwo.GetErrPrefixStr(),
			testMethod)
		return
	}

	if ePrefInfoTwo.GetErrContextStr() != errContext {
		t.Errorf("ERROR: ePrefInfoTwo.GetErrContextStr() "+
			"!= errContext\n"+
			"ePrefInfoTwo.GetErrContextStr()='%v'\n"+
			"errContext='%v'\n",
			ePrefInfoTwo.GetErrContextStr(),
			errContext)
		return
	}
}

func TestErrorPrefixInfo_CopyOut_000100(t *testing.T) {

	ePrefInfoOne := ErrorPrefixInfo{}.New()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoTwo,
		err :=
		ePrefInfoOne.CopyOut(
			"TestErrorPrefixInfo_CopyOut()")

	if err != nil {
		t.Errorf("ERROR returned by ePrefInfoTwo.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefInfoTwo.IsValidInstanceError("TestErrorPrefixInfo_CopyOut()")

	if err != nil {
		t.Errorf("ERROR returned by ePrefInfoTwo.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefInfoOne.Equal(&ePrefInfoTwo) {
		t.Error("ERROR: Expected ePrefInfoOne==ePrefInfoTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!\n ")
		return
	}

}

func TestErrorPrefixInfo_IsValidInstanceError_000100(t *testing.T) {

	ePrefInfoOne := ErrorPrefixInfo{}.New()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	err := ePrefInfoOne.IsValidInstanceError(
		"TestErrorPrefixInfo_IsValidInstanceError")

	if err != nil {
		t.Errorf("ERROR returned by ePrefInfoTwo.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefInfoOne.Empty()

	err = ePrefInfoOne.IsValidInstanceError(
		"TestErrorPrefixInfo_IsValidInstanceError #2")

	if err == nil {
		t.Error("ERROR Expected error return from ePrefInfoOne.IsValidInstanceError()\n" +
			"After ePrefInfoOne.Empty().\n" +
			"HOWEVER, NO ERROR WAS RETURNED.")
	}

}
