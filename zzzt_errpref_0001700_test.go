package errpref

import "testing"

func TestErrPrefixDelimiters_CopyIn_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn_000100()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne,
		err := ErrPrefixDelimiters{}.New(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ErrPrefixDelimiters{}.New()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsTwo := ErrPrefixDelimiters{}

	err = ePrefDelimsTwo.CopyIn(
		&ePrefDelimsOne,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsTwo.CopyIn()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsTwo.IsValidInstance() {
		t.Error("ERROR: " +
			"Expected ePrefDelimsTwo.IsValidInstance()=='true'\n" +
			"Instead, ePrefDelimsTwo.IsValidInstance()=='false'\n")

		return
	}

	if ePrefDelimsTwo.GetNewLinePrefixDelimiter() !=
		newLinePrefixDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetNewLinePrefixDelimiter()=="+
			"newLinePrefixDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetNewLinePrefixDelimiter()='%v'\n"+
			"newLinePrefixDelimiters='%v'\n",
			ePrefDelimsTwo.GetNewLinePrefixDelimiter(),
			newLinePrefixDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter() !=
		uint(len(newLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter()=="+
			"len(newLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter()='%v'\n"+
			"len(newLinePrefixDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter(),
			uint(len(newLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetInLinePrefixDelimiter() !=
		inLinePrefixDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetInLinePrefixDelimiter()=="+
			"inLinePrefixDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetInLinePrefixDelimiter()='%v'\n"+
			"inLinePrefixDelimiters='%v'\n",
			ePrefDelimsTwo.GetInLinePrefixDelimiter(),
			inLinePrefixDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthInLinePrefixDelimiter() !=
		uint(len(inLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthInLinePrefixDelimiter()=="+
			"len(inLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthInLinePrefixDelimiter()='%v'\n"+
			"len(inLinePrefixDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthInLinePrefixDelimiter(),
			uint(len(inLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetNewLineContextDelimiter() !=
		newLineContextDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetNewLineContextDelimiter()=="+
			"newLineContextDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetNewLineContextDelimiter()='%v'\n"+
			"newLineContextDelimiters='%v'\n",
			ePrefDelimsTwo.GetNewLineContextDelimiter(),
			newLineContextDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthNewLineContextDelimiter() !=
		uint(len(newLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthNewLineContextDelimiter()=="+
			"len(newLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthNewLineContextDelimiter()='%v'\n"+
			"len(newLineContextDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthNewLineContextDelimiter(),
			uint(len(newLineContextDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetInLineContextDelimiter() !=
		inLineContextDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetInLineContextDelimiter()=="+
			"inLineContextDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetInLineContextDelimiter()='%v'\n"+
			"inLineContextDelimiters='%v'\n",
			ePrefDelimsTwo.GetInLineContextDelimiter(),
			inLineContextDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthInLineContextDelimiter() !=
		uint(len(inLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthInLineContextDelimiter()=="+
			"len(inLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthInLineContextDelimiter()='%v'\n"+
			"len(inLineContextDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthInLineContextDelimiter(),
			uint(len(inLineContextDelimiters)))

		return
	}

	if !ePrefDelimsOne.Equal(&ePrefDelimsTwo) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne==ePrefDelimsTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")
	}

}

func TestErrPrefixDelimiters_CopyIn_000200(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn_000200()"

	ePrefDelimsTwo := ErrPrefixDelimiters{}

	err := ePrefDelimsTwo.CopyIn(
		nil,
		funcName)

	if err == nil {
		t.Errorf("ERROR:\n" +
			"Expected an error return from ePrefDelimsTwo.CopyIn()\n" +
			"because incomingDelimiters is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}
}

func TestErrPrefixDelimiters_CopyIn_000300(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn_000300()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne,
		err := ErrPrefixDelimiters{}.New(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ErrPrefixDelimiters{}.New()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsTwo := ErrPrefixDelimiters{}

	ePrefDelimsOne.Empty()

	err = ePrefDelimsTwo.CopyIn(
		&ePrefDelimsOne,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsTwo.CopyIn(&ePrefDelimsOne)\n" +
			"because 'ePrefDelimsOne' is empty an invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
	}
}

func TestErrPrefixDelimiters_CopyOut_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyOut()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.SetDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	var ePrefDelimsTwo ErrPrefixDelimiters

	ePrefDelimsTwo,
		err = ePrefDelimsOne.CopyOut(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsOne.Equal(&ePrefDelimsTwo) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne==ePrefDelimsTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")
	}
}

func TestErrPrefixDelimiters_CopyOut_000200(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyOut_000200()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.SetDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsOne.Empty()

	_,
		err = ePrefDelimsOne.CopyOut(funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			" Expected error return from ePrefDelimsOne.CopyOut()\n" +
			"because ePrefDelimsOne is empty.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestErrPrefixDelimiters_Empty_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne,
		err := ErrPrefixDelimiters{}.New(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ErrPrefixDelimiters{}.New()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsOne.IsValidInstance() {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne.IsValidInstance()=='true'\n" +
			"Instead, ePrefDelimsOne.IsValidInstance()=='false'\n")

		return
	}

	ePrefDelimsOne.Empty()

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err == nil {
		t.Error("ERROR: " +
			"Expected error return from ePrefDelimsOne.IsValidInstanceError()\n" +
			"after call to ePrefDelimsOne.Empty()\n" +
			"HOWEVER NO ERROR WAS RETURNED!\n")
	}

}

func TestErrPrefixDelimiters_Empty_000200(t *testing.T) {

	ePrefDelimsOne := ErrPrefixDelimiters{}

	ePrefDelimsOne.Empty()

	if ePrefDelimsOne.IsValidInstance() {
		t.Error("ERROR:\n" +
			"ePrefDelimsOne.IsValidInstance()==true after\n" +
			"Empty() was called on it!\n")
	}

}

func TestErrPrefixDelimiters_SetLineLengthValues_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_SetLineLengthValues()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.SetDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsOne.SetLineLengthValues()

	if ePrefDelimsOne.GetLengthNewLinePrefixDelimiter() !=
		uint(len(newLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthNewLinePrefixDelimiter()=="+
			"len(newLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthNewLinePrefixDelimiter()='%v'\n"+
			"len(newLinePrefixDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthNewLinePrefixDelimiter(),
			uint(len(newLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthInLinePrefixDelimiter() !=
		uint(len(inLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthInLinePrefixDelimiter()=="+
			"len(inLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthInLinePrefixDelimiter()='%v'\n"+
			"len(inLinePrefixDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthInLinePrefixDelimiter(),
			uint(len(inLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthNewLineContextDelimiter() !=
		uint(len(newLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthNewLineContextDelimiter()=="+
			"len(newLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthNewLineContextDelimiter()='%v'\n"+
			"len(newLineContextDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthNewLineContextDelimiter(),
			uint(len(newLineContextDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthInLineContextDelimiter() !=
		uint(len(inLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthInLineContextDelimiter()=="+
			"len(inLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthInLineContextDelimiter()='%v'\n"+
			"len(inLineContextDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthInLineContextDelimiter(),
			uint(len(inLineContextDelimiters)))
	}

}

func TestErrPrefixDelimiters_SetDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_SetLineLengthValues()"

	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		"",
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsOne.SetDelimiters()\n" +
			"because 'newLinePrefixDelimiters' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	newLinePrefixDelimiters := "\n"

	err = ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		"",
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsOne.SetDelimiters()\n" +
			"because 'inLinePrefixDelimiters' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	err = ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		"",
		inLineContextDelimiters,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsOne.SetDelimiters()\n" +
			"because 'newLineContextDelimiters' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	err = ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		"",
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsOne.SetDelimiters()\n" +
			"because 'inLineContextDelimiters' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestErrPrefixDelimiters_SetToDefault_000100(t *testing.T) {
	ePrefDelims := ErrPrefixDelimiters{}

	ePrefDelims.SetToDefault()

	newLineErrPrefix := ePrefDelims.GetNewLinePrefixDelimiter()

	actNewLinePrefixDelimiters := "\n"
	//actInLinePrefixDelimiters := " - "
	//actNewLineContextDelimiters := "\n :  "
	//actInLineContextDelimiters := " : "

	if newLineErrPrefix != actNewLinePrefixDelimiters {

		newLineErrPrefix =
			ErrPref{}.ConvertNonPrintableChars(
				[]rune(newLineErrPrefix),
				true)

		t.Errorf("ERROR:\n"+
			"Default New Line Prefix Delimiter is incorrect!\n"+
			"New Line Error Prefix Delimiters = '%v'\n",
			newLineErrPrefix)
	}

}
