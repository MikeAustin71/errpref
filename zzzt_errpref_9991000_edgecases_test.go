package errpref

import (
	"strings"
	"testing"
)

func TestErrPref_EPrefFile_000100(t *testing.T) {
	/*
		func (ePref ErrPref) SetCtxt(
			oldErrPref string,
			newErrContext string) string
	*/

	ePref := ErrPref{}

	_ = ePref.SetCtxt(
		"oldErrPref",
		"")

	ePref = ErrPref{}

	_ = ePref.SetCtxt(
		"",
		"newErrContext")

}

func TestErrPref_Mechanics_000100(t *testing.T) {

	/*	func (ePrefMech *errPrefMechanics) assembleErrPrefix(
		oldErrPrefix string,
		newErrPrefix string,
		newErrContext string,
		maxErrStringLength uint,
		delimiters ErrPrefixDelimiters) string
	*/

	ePrefMech := errPrefMechanics{}

	delimiters := ErrPrefixDelimiters{}

	delimiters.SetToDefault()

	oldErrPrefix := "OldStuff()"
	newErrPrefix := "Tx2.DoSomething()"
	newErrContext := ""
	maxErrStringLength := uint(0)

	outputStr :=
		ePrefMech.assembleErrPrefix(
			oldErrPrefix,
			newErrPrefix,
			newErrContext,
			maxErrStringLength,
			delimiters)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.assembleErrPrefix() " +
			"returned a zero length output string!\n")
		return
	}

	/*	func (ePrefMech *errPrefMechanics) formatErrPrefix(
		maxErrStringLength uint,
		delimiters ErrPrefixDelimiters,
		errPrefix string) string
	*/

	maxErrStringLength = 0

	ePrefMech = errPrefMechanics{}

	outputStr =
		ePrefMech.formatErrPrefix(
			maxErrStringLength,
			delimiters,
			newErrPrefix)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.formatErrPrefix() " +
			"returned a zero length output string!\n")
		return
	}

	/*
		func (ePrefMech *errPrefMechanics) setErrorContext(
			oldErrPref string,
			newErrContext string,
			maxErrStringLength uint,
			delimiters ErrPrefixDelimiters) string
	*/

	maxErrStringLength = 0
	newErrContext = "X=7; Y=92"

	ePrefMech = errPrefMechanics{}

	outputStr =
		ePrefMech.setErrorContext(
			oldErrPrefix,
			newErrContext,
			maxErrStringLength,
			delimiters)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.setErrorContext() " +
			"returned a zero length output string!\n")
	}

}

func TestErrPref_Molecule_000100(t *testing.T) {

	ePrefMolecule := errPrefMolecule{}

	/*
		func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithContext(
			strBuilder *strings.Builder,
		lineLenCalc *EPrefixLineLenCalc) error
	*/

	var strBuilder *strings.Builder

	var lineLenCalc *EPrefixLineLenCalc

	xLineLenCalc := EPrefixLineLenCalc{}.New()

	lineLenCalc = &xLineLenCalc

	err := ePrefMolecule.writeNewEPrefWithContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithContext()\n" +
			"because strBuilder is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefMolecule = errPrefMolecule{}

	err = ePrefMolecule.writeNewEPrefWithOutContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithOutContext()\n" +
			"because strBuilder is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	strBuilder = &strings.Builder{}

	lineLenCalc = nil

	ePrefMolecule = errPrefMolecule{}

	err = ePrefMolecule.writeNewEPrefWithContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithContext()\n" +
			"because lineLenCalc is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefMolecule = errPrefMolecule{}

	err = ePrefMolecule.writeNewEPrefWithOutContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithOutContext()\n" +
			"because lineLenCalc is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPref_Nanobot_000100(t *testing.T) {

	/*
		func (ePrefNanobot *errPrefNanobot) formatErrPrefixComponents(
			leadingTextStr string,
			trailingTextStr string,
			maxErrPrefixTextLineLength uint,
			isLastLineTerminatedWithNewLine bool,
			delimiters ErrPrefixDelimiters,
			prefixContextCol []ErrorPrefixInfo) string
	*/

	ePrefNanobot := errPrefNanobot{}

	errPrefDelimiters := getValidErrPrefixDelimiters()

	var prefixContextCol []ErrorPrefixInfo

	prefixContextCol = nil

	_ = ePrefNanobot.formatErrPrefixComponents(
		"",
		"",
		70,
		false,
		errPrefDelimiters,
		prefixContextCol)

	prefixContextCol = getValidErrorPrefixInfoArray()
	errPrefDelimiters = getValidErrPrefixDelimiters()

	errPrefDelimiters.newLinePrefixDelimiter = ""

	errPrefDelimiters.newLineContextDelimiter = ""

	_ = ePrefNanobot.formatErrPrefixComponents(
		"",
		"",
		70,
		false,
		errPrefDelimiters,
		prefixContextCol)

	prefixContextCol = getValidErrorPrefixInfoArray()
	errPrefDelimiters = getValidErrPrefixDelimiters()

	_ = ePrefNanobot.formatErrPrefixComponents(
		"",
		"",
		0,
		false,
		errPrefDelimiters,
		prefixContextCol)

	/*
		func (ePrefNanobot *errPrefNanobot) setLastCtx(
			newErrContext string,
			errPrefixCollection []ErrorPrefixInfo)

	*/

	ePrefNanobot = errPrefNanobot{}

	prefixContextCol = getValidErrorPrefixInfoArray()

	ePrefNanobot.setLastCtx(
		"",
		prefixContextCol)

	prefixContextCol = nil

	ePrefNanobot.setLastCtx(
		"Tx92.Awesome()",
		prefixContextCol)

	prefixContextCol = make([]ErrorPrefixInfo, 0)

	ePrefNanobot.setLastCtx(
		"Tx92.Awesome()",
		prefixContextCol)

	/*
		func (ePrefNanobot *errPrefNanobot) extractLastErrPrfInfo(
			errPref string) ErrorPrefixInfo
	*/

	ePrefNanobot = errPrefNanobot{}

	_ = ePrefNanobot.extractLastErrPrfInfo(
		"")

}

func TestErrPref_Neutron_000100(t *testing.T) {

	funcName := "TestErrPref_Neutron_000100() "
	/*
		func (ePrefNeutron *errPrefNeutron) writeCurrentLineStr(
			strBuilder *strings.Builder,
		ePrefLineLenCalc *EPrefixLineLenCalc)
	*/

	ePrefNeutron := errPrefNeutron{}

	strBuilder := &strings.Builder{}

	ePrefLineLenCalc,
		err := getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Errorf("Error returned by getValidEPrefixLineLenCalc()\n"+
			"%v\n", err.Error())
		return
	}

	ePrefNeutron.writeCurrentLineStr(
		strBuilder,
		nil)

	ePrefNeutron.writeCurrentLineStr(
		nil,
		&ePrefLineLenCalc)

	ePrefLineLenCalc,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Errorf("Error returned by getValidEPrefixLineLenCalc()\n"+
			"%v\n", err.Error())
		return
	}

	ePrefLineLenCalc.currentLineStr = ""

	ePrefNeutron.writeCurrentLineStr(
		nil,
		&ePrefLineLenCalc)

}

func TestErrPref_Quark_000100(t *testing.T) {

	funcName := "TestErrPref_Quark_000100()"

	ePrefQuark := errPrefQuark{}

	/*
		func (ePrefQuark *errPrefQuark) convertPrintableChars(
			printableChars string,
			ePrefix string) (
			nonPrintableChars []rune,
			err error)
	*/

	_,
		err :=
		ePrefQuark.convertPrintableChars(
			"",
			funcName)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefQuark.convertPrintableChars()\n" +
			"because the printableChars parameter is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefQuark = errPrefQuark{}

	/*
		func (ePrefQuark *errPrefQuark) setErrPrefDisplayLineLength(
			maxErrPrefixStringLength uint)
	*/

	ePrefQuark.setErrPrefDisplayLineLength(0)

}

func TestErrPrefixDelimiters_MainFile_000100(t *testing.T) {

	/*
		func (ePrefDelims *ErrPrefixDelimiters) Equal(
			incomingDelimiters *ErrPrefixDelimiters) (
			areEqual bool)
	*/

	ePrefDelims := ErrPrefixDelimiters{}

	delimitersOne := getValidErrPrefixDelimiters()

	_ = ePrefDelims.Equal(&delimitersOne)

	ePrefDelims = ErrPrefixDelimiters{}

	_ = ePrefDelims.GetInLinePrefixDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.lenInLineContextDelimiter = 0

	_ = ePrefDelims.GetLengthInLineContextDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.lenNewLineContextDelimiter = 0

	_ = ePrefDelims.GetLengthNewLineContextDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.lenNewLinePrefixDelimiter = 0

	_ = ePrefDelims.GetLengthNewLinePrefixDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	_ = ePrefDelims.GetNewLineContextDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	_ = ePrefDelims.GetNewLinePrefixDelimiter()

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.SetInLineContextDelimiter("Xray")

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.SetNewLineContextDelimiter("Xray")

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.SetNewLinePrefixDelimiter("Xray")

	ePrefDelims = ErrPrefixDelimiters{}

	ePrefDelims.SetLineLengthValues()

	ePrefDelims = ErrPrefixDelimiters{}

	_ = ePrefDelims.String()

}

func TestErrPrefixDelimiters_Mechanics(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_Mechanics()"

	ePrefDelimitersMech := errPrefixDelimitersMechanics{}

	err := ePrefDelimitersMech.setErrPrefDelimiters(
		nil,
		"X",
		"Y",
		"Z",
		"A",
		funcName)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimitersMech.setErrPrefDelimiters()\n" +
			"because the ePrefDelimiters parameter is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefDelimitersMech = errPrefixDelimitersMechanics{}

	err = ePrefDelimitersMech.setToDefault(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimitersMech.setToDefault()\n" +
			"because the ePrefDelimiters parameter is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDelimiters_Electron_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_Electron_000100() "

	/*
		func (ePrefDelimsElectron *errPrefixDelimitersElectron) copyIn(
			targetDelimiters *ErrPrefixDelimiters,
			incomingDelimiters *ErrPrefixDelimiters,
			ePrefix string) error
	*/

	delimiters := getValidErrPrefixDelimiters()

	ePrefDelimsElectron := errPrefixDelimitersElectron{}

	err :=
		ePrefDelimsElectron.copyIn(
			nil,
			&delimiters,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsElectron.copyIn()\n" +
			"because the 'targetDelimiters' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefDelimsElectron.copyIn(
			&delimiters,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsElectron.copyIn()\n" +
			"because the 'incomingDelimiters' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	/*
		func (ePrefDelimsElectron *errPrefixDelimitersElectron) copyOut(
			delimiters *ErrPrefixDelimiters,
		ePrefix string) (
			ErrPrefixDelimiters,
			error)
	*/

	delimiters = getValidErrPrefixDelimiters()

	ePrefDelimsElectron = errPrefixDelimitersElectron{}

	_,
		err =
		ePrefDelimsElectron.copyOut(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsElectron.copyOut()\n" +
			"because the 'delimiters' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters = getValidErrPrefixDelimiters()

	delimiters2 := getValidErrPrefixDelimiters()

	ePrefDelimsElectron = errPrefixDelimitersElectron{}

	_,
		err =
		ePrefDelimsElectron.equal(
			nil,
			&delimiters2,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsElectron.equal()\n" +
			"because the 'delimitersOne' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDelimsElectron.equal()\n" +
			"because the 'delimitersTwo' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	var areEqual bool

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.inLinePrefixDelimiter = "9999!XXX"

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.inLinePrefixDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER, areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.lenInLinePrefixDelimiter = 469521

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.lenInLinePrefixDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.newLinePrefixDelimiter = "FX91!!!!"

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.newLinePrefixDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.lenNewLinePrefixDelimiter = 6712385

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.lenNewLinePrefixDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.inLineContextDelimiter = "!x%^#"

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.inLineContextDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.lenInLineContextDelimiter = 5894

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.lenInLineContextDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.newLineContextDelimiter = "!@!#$%^"

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.newLineContextDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

	delimiters2 = getValidErrPrefixDelimiters()

	delimiters2.lenNewLineContextDelimiter = 578

	areEqual,
		err =
		ePrefDelimsElectron.equal(
			&delimiters,
			&delimiters2,
			funcName)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefDelimsElectron.equal() would return areEqual=='false'\n" +
			"because the 'delimitersTwo.lenNewLineContextDelimiter' parameter is NOT EQUAL.\n" +
			"HOWEVER,  areEqual=='true'!\n")
		return
	}

}

func TestErrPrefixDelimiters_Quark_000100(t *testing.T) {
	/*
		func (ePrefDelimsQuark *errPrefixDelimitersQuark) empty(
			delimiters *ErrPrefixDelimiters,
			ePrefix string) (err error)
	*/

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	var delimiters *ErrPrefixDelimiters
	ePrefix := "TestCallingMethod()"

	err := ePrefDelimsQuark.empty(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.empty()\n" +
			"because 'delimiters' is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefDelimsQuark = errPrefixDelimitersQuark{}

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters' is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters = &ErrPrefixDelimiters{}

	delimiters.SetToDefault()

	delimiters.inLinePrefixDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.inLinePrefixDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.newLinePrefixDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.newLinePrefixDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.inLineContextDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.inLineContextDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.newLineContextDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.newLineContextDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestEPrefixLineLenCalc_000100(t *testing.T) {

	ePrefLineLenCalc := EPrefixLineLenCalc{}

	ePrefLineLenCalc.currentLineStr =
		"Now is the time for all good men to enlist!"

	ePrefLineLenCalc.maxErrStringLength = 11

	if ePrefLineLenCalc.CurrLineLenExceedsMaxLineLen() == false {
		t.Error("ERROR:\n" +
			"Expected ePrefLineLenCalc.CurrLineLenExceedsMaxLineLen()\n" +
			"to return 'true'.\n" +
			"HOWEVER, IT RETURNED FALSE INSTEAD!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_,
		_ = ePrefLineLenCalc.CopyOut("")

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	ePrefLineLenCalc.Empty()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ =
		ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen()

	ePrefLineLenCalc = EPrefixLineLenCalc{}.New()

	ePrefLineLenCalc.currentLineStr =
		"Now is the time for all good men to come to the aid"

	ePrefLineLenCalc.maxErrStringLength = 11

	if ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == false {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}.New()

	ePrefLineLenCalc.currentLineStr =
		"Now is the time for all good men to come to the aid"

	ePrefLineLenCalc.maxErrStringLength = 11

	if ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == false {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}.New()

	ePrefLineLenCalc.currentLineStr =
		"Now is the time for all good men to come to the aid"

	ePrefLineLenCalc.maxErrStringLength = 11

	if ePrefLineLenCalc.EPrefixWithContextExceedsRemainLineLen() == false {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	if ePrefLineLenCalc.EPrefixWithContextExceedsRemainLineLen() == true {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.EPrefWithoutContextExceedsRemainLineLen() == 'false'\n" +
			"HOWEVER, IT RETURNED TRUE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	if ePrefLineLenCalc.ErrPrefixHasContext() == true {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.ErrPrefixHasContext() == 'false'\n" +
			"HOWEVER, IT RETURNED TRUE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	if ePrefLineLenCalc.ErrorContextIsEmpty() == false {
		t.Error("ERROR\n" +
			"Expected #1 ePrefLineLenCalc.ErrorContextIsEmpty() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}.New()

	if ePrefLineLenCalc.ErrorContextIsEmpty() == false {
		t.Error("ERROR\n" +
			"Expected #2 ePrefLineLenCalc.ErrorContextIsEmpty() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	if ePrefLineLenCalc.ErrorPrefixIsEmpty() == false {
		t.Error("ERROR\n" +
			"Expected #1 ePrefLineLenCalc.ErrorPrefixIsEmpty() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}.New()

	if ePrefLineLenCalc.ErrorPrefixIsEmpty() == false {
		t.Error("ERROR\n" +
			"Expected #2 ePrefLineLenCalc.ErrorPrefixIsEmpty() == 'true'\n" +
			"HOWEVER, IT RETURNED FALSE!\n")
		return
	}

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ = ePrefLineLenCalc.GetErrorContextStr()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ = ePrefLineLenCalc.GetCurrLineStr()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ = ePrefLineLenCalc.GetCurrLineStrLength()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ = ePrefLineLenCalc.GetDelimiterInLineErrContext()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	_ = ePrefLineLenCalc.GetErrorPrefixStr()

	ePrefLineLenCalc = EPrefixLineLenCalc{}

	if ePrefLineLenCalc.IsErrPrefixLastIndex() {
		t.Error("ERROR\n" +
			"Expected ePrefLineLenCalc.IsErrPrefixLastIndex() == 'false'\n" +
			"HOWEVER, IT RETURNED 'true'!\n")
		return
	}

}

func TestEPrefixLineLenCalc_Electron_000100(t *testing.T) {

	funcName := "TestEPrefixLineLenCalc_Electron_000100() "

	ePrefLineLenCalcElectron := ePrefixLineLenCalcElectron{}

	var targetLineLenCalc, incomingLineLenCalc *EPrefixLineLenCalc

	incomingLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	ePrefix := "TestCallingMethod"

	err :=
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because targetLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	incomingLineLenCalc = nil

	err =
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because incomingLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	incomingLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	incomingLineLenCalc.ePrefDelimiters.newLinePrefixDelimiter = ""

	err =
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because incomingLineLenCalc is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = nil

	ePrefLineLenCalcElectron = ePrefixLineLenCalcElectron{}

	_,
		err = ePrefLineLenCalcElectron.copyOut(
		targetLineLenCalc,
		ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyOut()\n" +
			"because targetLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	targetLineLenCalc.ePrefDelimiters.SetToDefault()

	targetLineLenCalc.ePrefDelimiters.newLinePrefixDelimiter = ""

	_,
		err = ePrefLineLenCalcElectron.copyOut(
		targetLineLenCalc,
		ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyOut()\n" +
			"because targetLineLenCalc is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	/*
		func (ePrefLineLenCalcElectron *ePrefixLineLenCalcElectron) equal(
			lineLenCalcOne *EPrefixLineLenCalc,
			lineLenCalcTwo *EPrefixLineLenCalc,
			ePrefix string) (
			areEqual bool,
			err error)
	*/

	var lineLenCalcOne, lineLenCalcTwo EPrefixLineLenCalc

	lineLenCalcOne,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	ePrefLineLenCalcElectron = ePrefixLineLenCalcElectron{}

	_,
		err = ePrefLineLenCalcElectron.equal(
		nil,
		&lineLenCalcTwo,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.equal()\n" +
			"because lineLenCalcOne is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.equal()\n" +
			"because lineLenCalcTwo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	var areEqual bool

	lineLenCalcTwo.ePrefDelimiters.newLinePrefixDelimiter = "@@@@@@@@@"

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because ePrefDelimiters.newLinePrefixDelimiter\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.errorPrefixInfo.errorPrefixStr = ""

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because errorPrefixInfo.errorPrefixStr\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.currentLineStr = "999"

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because lineLenCalcTwo.currentLineStr\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.maxErrStringLength = 15

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because lineLenCalcTwo.maxErrStringLength\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

}

func TestEPrefixLineLenCalc_Quark_000100(t *testing.T) {

	funcName := "TestErrorPrefixInfo_Electron_000100() "

	ePrefLineLenCalcQuark := ePrefixLineLenCalcQuark{}

	_,
		err := ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc()\n" +
			"because ePrefLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	var ePrefLineLenCalc EPrefixLineLenCalc

	ePrefLineLenCalc,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Errorf("Error returned by getValidEPrefixLineLenCalc()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	ePrefLineLenCalcQuark = ePrefixLineLenCalcQuark{}

	ePrefLineLenCalc.errorPrefixInfo = ErrorPrefixInfo{}
	_,
		err = ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		&ePrefLineLenCalc,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc()\n" +
			"because ePrefLineLenCalc.errorPrefixInfo == nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefLineLenCalc,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Errorf("Error returned by getValidEPrefixLineLenCalc()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	ePrefLineLenCalc.maxErrStringLength = 0
	_,
		err = ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		&ePrefLineLenCalc,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc()\n" +
			"because ePrefLineLenCalc.maxErrStringLength == 0.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrorPrefixInfo_Electron_000100(t *testing.T) {

	funcName := "TestErrorPrefixInfo_Electron_000100() "

	ePrefInfoElectron := errorPrefixInfoElectron{}

	targetErrPrefixInfo := ErrorPrefixInfo{}

	inComingErrPrefixInfo := getValidErrorPrefixInfo()

	err :=
		ePrefInfoElectron.copyIn(
			nil,
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because targetErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefInfoElectron.copyIn(
			&targetErrPrefixInfo,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because inComingErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo.errorPrefixStr = ""

	err =
		ePrefInfoElectron.copyIn(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because inComingErrPrefixInfo is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefInfoElectron = errorPrefixInfoElectron{}
	_,
		err =
		ePrefInfoElectron.copyOut(
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyOut()\n" +
			"because inComingErrPrefixInfo is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err =
		ePrefInfoElectron.copyOut(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyOut()\n" +
			"because inComingErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetErrPrefixInfo = getValidErrorPrefixInfo()
	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	ePrefInfoElectron = errorPrefixInfoElectron{}

	areEqual :=
		ePrefInfoElectron.equal(
			nil,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because ePrefixInfo01 is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			nil)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because ePrefixInfo02 is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo.isFirstIdx = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isFirstIdx is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.isLastIdx = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isLastIdx is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.isPopulated = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isPopulated is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.errorPrefixStr = "Nothing"

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the errorPrefixStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.lenErrorPrefixStr = 9999999

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the lenErrorPrefixStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.errorContextStr = "XXX-NOTHING-XXX"

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the errorContextStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.lenErrorContextStr = 55998823

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the lenErrorContextStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrorPrefixInfo_Quark_000100(t *testing.T) {

	funcName := "TestErrorPrefixInfo_Quark_000100()"

	ePrefInfoQuark := errorPrefixInfoQuark{}

	/*
		func (ePrefDtoQuark *errorPrefixInfoQuark) testValidityOfErrorPrefixInfo(
			errPrefixInfo *ErrorPrefixInfo,
		ePrefix string) (
			isValid bool,
			err error)
	*/

	_,
		err := ePrefInfoQuark.testValidityOfErrorPrefixInfo(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoQuark.testValidityOfErrorPrefixInfo()\n" +
			"because the 'errPrefixInfo' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDto_MainFile_0001000(t *testing.T) {

	ePrefDto := ErrPrefixDto{}

	twoDSlice := getValidTwoDStrArray()

	ePrefDto.AddEPrefStrings(twoDSlice)

	colStr := getValidErrorPrefixCollectionStr()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.AddEPrefCollectionStr(colStr)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.ClearLeadingTextStr()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.ClearTrailingTextStr()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.Copy()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.CopyPtr()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.CopyIn(nil, "")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.CopyInFromIBuilder(
		nil, "")

	ePrefDto = ErrPrefixDto{}

	ePrefDto2 := ErrPrefixDto{}

	ePrefDto.CopyOutToIBuilder(
		&ePrefDto2)

	ePrefDto = ErrPrefixDto{}

	_,
		_ = ePrefDto.CopyOut("")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.DeleteLastErrPrefix()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.Empty()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.EmptyEPrefCollection()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.Equal(nil)

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetInputStringDelimiters()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetIsLastLineTerminatedWithNewLine()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetIsLastLineTerminatedWithNewLine()

	ePrefDto = ErrPrefixDto{}

	_,
		_,
		_ = ePrefDto.GetLastErrPrefix("")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetLeadingTextStr()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetLeftMarginChar()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetEPrefCollectionLen()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetEPrefStrings()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetMaxTextLineLenDefault()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetOutputStringDelimiters()

	ePrefDto = ErrPrefixDto{}

	_,
		_ = ePrefDto.GetStrDelimiters()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetTrailingTextStr()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetTurnOffTextDisplay()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.MergeErrPrefixDto(nil)

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.GetTurnOffTextDisplay()
	/*
		func (ePrefDto ErrPrefixDto) NewFromErrPrefDto(
			dto *ErrPrefixDto,
		newErrPrefix string,
			newErrContext string) (
			newErrPrefDto *ErrPrefixDto,
			err error)
	*/

	ePrefDto = ErrPrefixDto{}

	_,
		_ = ePrefDto.NewFromErrPrefDto(
		nil,
		"TX2.Awesome() ",
		"")
	/*
		func (ePrefDto ErrPrefixDto) NewIBasicErrorPrefix(
			iEPref IBasicErrorPrefix,
			newErrPrefix string,
			newErrContext string) (
			*ErrPrefixDto,
			error)
	*/

	ePrefDto = ErrPrefixDto{}

	ePrefDto2 = getValidErrorPrefixDto()

	ePrefDto2.inputStrDelimiters.newLinePrefixDelimiter = ""
	ePrefDto2.outputStrDelimiters.newLinePrefixDelimiter = ""

	_,
		_ = ePrefDto.NewIBasicErrorPrefix(
		&ePrefDto2,
		"TX2.Awesome() ",
		"")

}

func TestErrPrefixDto_MainFile_000200(t *testing.T) {

	ePrefDto := ErrPrefixDto{}

	_ =
		ePrefDto.ReplaceLastErrPrefix(
			"",
			"",
			"")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetCtx(
		"")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetCtxEmpty()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetEPrefCollection(nil)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetEPrefCtx("", "")

	ePrefDto = ErrPrefixDto{}

	_ =
		ePrefDto.SetInputStringDelimiters(
			ErrPrefixDelimiters{}, "")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetIsLastLineTermWithNewLine(true)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetLeadingTextStr("")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetLeftMarginChar(rune(0))

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetLeftMarginLength(-99)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetMaxTextLineLenToDefault()

	ePrefDto = ErrPrefixDto{}

	_ =
		ePrefDto.SetOutputStringDelimiters(
			ErrPrefixDelimiters{}, "")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetStrDelimitersToDefault()

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetTrailingTextStr("Hello")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.SetTurnOffTextDisplay(true)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.maxErrPrefixTextLineLength = 2

	_ = ePrefDto.String()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.StrMaxLineLen(2)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.maxErrPrefixTextLineLength = 2

	_ = ePrefDto.StrMaxLineLen(2)

	ePrefDto = ErrPrefixDto{}

	ePrefDto.ePrefCol = nil

	_ = ePrefDto.XCtxEmpty()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.XEPref("")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.XEPrefCtx("", "")

	/*
		func (ePrefDto *ErrPrefixDto) XSetFromStrings(
			oldErrPrefix string,
			newErrPrefix string,
			newErrContext string,
			inputStrDelimiters ErrPrefixDelimiters,
			outputStrDelimiters ErrPrefixDelimiters,
			ePrefix string) (*ErrPrefixDto, error)
	*/

	ePrefDto = ErrPrefixDto{}

	_,
		_ = ePrefDto.XSetFromStrings(
		"",
		"",
		"",
		ErrPrefixDelimiters{},
		ErrPrefixDelimiters{},
		"")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.ZCtx("NeverMind()")

	ePrefDto = ErrPrefixDto{}

	ePrefDto.ZCtxEmpty()

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.ZEPref("")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.ZEPrefCtx("", "")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.ZEPrefCtx("Hello()", "E=MC^2")

	ePrefDto = ErrPrefixDto{}

	_ = ePrefDto.ZEPrefOld("Hello()\nGoodbye()")

}

func TestErrPrefixDto_Mechanics_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Mechanics_000100() "

	/*
		func (ePrefDtoMech *errPrefixDtoMechanics) get2dEPrefStrings(
			errPrefDto *ErrPrefixDto,
		errorPrefStr string) (
			twoDStrArray [][2]string,
			err error)
	*/

	ePrefDtoMech := errPrefixDtoMechanics{}

	var ePDto ErrPrefixDto

	_,
		err := ePrefDtoMech.get2dEPrefStrings(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDtoMech.get2dEPrefStrings()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto = getValidErrorPrefixDto()

	ePDto.ePrefCol = nil

	_,
		err = ePrefDtoMech.get2dEPrefStrings(
		&ePDto,
		funcName)

	if err != nil {
		t.Errorf("ERROR:\n"+
			"ePrefDtoMech.get2dEPrefStrings() SHOULD NOT HAVE RETURNED AND Error.\n"+
			"'ePDto.ePrefCol' parameter is 'nil'.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	/*
		func (ePrefDtoMech *errPrefixDtoMechanics) setFromEmptyInterface(
			errPrefDto *ErrPrefixDto,
		iEPref interface{},
		errorPrefStr string) error
	*/

	ePDto = getValidErrorPrefixDto()

	var twoDStrings [][2]string

	twoDStrings = ePDto.GetEPrefStrings()

	ePrefDtoMech = errPrefixDtoMechanics{}

	err = ePrefDtoMech.setFromEmptyInterface(
		nil,
		twoDStrings,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDtoMech.setFromEmptyInterface()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDto_Nanobot_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Nanobot_000100() "

	ePrefixDtoNanobot := errPrefixDtoNanobot{}

	_,
		err :=
		ePrefixDtoNanobot.copyOutErrPrefDtoPtr(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.copyOutErrPrefDtoPtr()\n" +
			"because the 'ePrefixDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	ePrefixDtoNanobot.deleteLastErrContext(nil)

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	ePDto2 := getValidErrorPrefixDto()

	err =
		ePrefixDtoNanobot.setFromIBuilder(
			nil,
			&ePDto2,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBuilder()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromIBuilder(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBuilder()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	err =
		ePrefixDtoNanobot.setFromIBasicErrorPrefix(
			nil,
			&ePDto2,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBasicErrorPrefix()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromIBasicErrorPrefix(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBasicErrorPrefix()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	err =
		ePrefixDtoNanobot.setFromString(
			nil,
			"Hello",
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromString()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromString(
			&ePDto2,
			"",
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromString() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"an empty string.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	strBuilder := strings.Builder{}

	err =
		ePrefixDtoNanobot.setFromStringBuilder(
			nil,
			&strBuilder,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringBuilder()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()

	err =
		ePrefixDtoNanobot.setFromStringBuilder(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringBuilder()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	strArray := []string{"Hello", "Goodbye", "Come to dinner"}

	err =
		ePrefixDtoNanobot.setFromStringArray(
			nil,
			strArray,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringArray()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()
	strArray = nil

	_ =
		ePrefixDtoNanobot.setFromStringArray(
			&ePDto2,
			strArray,
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromStringArray() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"nil.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	str2DArray := make([][2]string, 3)

	str2DArray[0][0] = "Hello()"
	str2DArray[1][0] = "GoodBye()"
	str2DArray[2][0] = "ComeAgain()"

	err =
		ePrefixDtoNanobot.setFromTwoDStrArray(
			nil,
			str2DArray,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromTwoDStrArray()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()

	_ =
		ePrefixDtoNanobot.setFromTwoDStrArray(
			&ePDto2,
			nil,
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromTwoDStrArray() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"nil.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

}

func TestErrPrefixDto_Atom_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Atom_000100() "

	/*
		func (ePrefixDtoAtom *errPrefixDtoAtom) addTwoDimensionalStringArray(
			errPrefixDto *ErrPrefixDto,
			twoDStrArray [][2]string,
			errPrefStr string) error
	*/

	ePrefixDtoAtom := errPrefixDtoAtom{}

	twoDSlice := getValidTwoDStrArray()

	err := ePrefixDtoAtom.addTwoDimensionalStringArray(
		nil,
		twoDSlice,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoAtom.addTwoDimensionalStringArray()\n" +
			"because the 'errPrefixDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	nPDtoOne := getValidErrorPrefixDto()

	err = ePrefixDtoAtom.addTwoDimensionalStringArray(
		&nPDtoOne,
		nil,
		funcName)

	if err != nil {
		t.Errorf("ERROR:\n"+
			"Did NOT expect an error return from ePrefixDtoAtom.addTwoDimensionalStringArray()\n"+
			"even though the 'twoDStrArray' parameter is 'nil'.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n"+
			"%v", err.Error())
		return
	}

	ePrefixDtoAtom = errPrefixDtoAtom{}

	areEqual := ePrefixDtoAtom.areEqualErrPrefDtos(
		nil,
		&nPDtoOne)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto1' parameter is 'nil'.\n" +
			"HOWEVER, areEqual=='true'!\n")
		return
	}

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		nil)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2' parameter is 'nil'.\n" +
			"HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoOne = getValidErrorPrefixDto()

	nPDtoTwo := getValidErrorPrefixDto()

	nPDtoTwo.leftMarginLength = 99999999

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.leftMarginLength'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoOne = getValidErrorPrefixDto()

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.leftMarginChar = '%'

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.leftMarginChar'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.isLastLineTerminatedWithNewLine =
		!nPDtoOne.isLastLineTerminatedWithNewLine

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.isLastLineTerminatedWithNewLine'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.turnOffTextDisplay =
		!nPDtoOne.turnOffTextDisplay

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.turnOffTextDisplay'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.inputStrDelimiters.newLinePrefixDelimiter =
		"!@$%^&*"

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.inputStrDelimiters.newLinePrefixDelimiter'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.outputStrDelimiters.newLinePrefixDelimiter =
		"!@$%^&*"

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.outputStrDelimiters.newLinePrefixDelimiter'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.maxErrPrefixTextLineLength =
		99999999

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.maxErrPrefixTextLineLength'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.leadingTextStr =
		"!@#$%^&*()"

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.leadingTextStr'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.trailingTextStr =
		"!@#$%^&*()"

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.trailingTextStr'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

	nPDtoTwo = getValidErrorPrefixDto()

	nPDtoTwo.ePrefCol = nil

	areEqual = ePrefixDtoAtom.areEqualErrPrefDtos(
		&nPDtoOne,
		&nPDtoTwo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoAtom.areEqualErrPrefDtos() would return\n" +
			"areEqual=='false' because the 'errPrefixDto2.ePrefCol'\n" +
			"is NOT equal. HOWEVER, areEqual=='true'!\n")
		return
	}

}

func TestErrPrefixDto_Electron_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Electron_000100"

	/*
		func (ePrefDtoElectron errPrefixDtoElectron) emptyErrorPrefixDto(
			ePrefixDto *ErrPrefixDto,
			errPrefStr string) error
	*/

	ePrefDtoElectron := errPrefixDtoElectron{}

	err :=
		ePrefDtoElectron.emptyErrorPrefixDto(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoElectron.emptyErrorPrefixDto()\n" +
			"because parameter ePrefixDto is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestErrPrefixDto_Quark_000100(t *testing.T) {

	funcName := "ePrefDtoQuark.errPrefixDtoQuark()"

	ePrefDtoQuark := errPrefixDtoQuark{}

	_,
		err := ePrefDtoQuark.deleteLastErrPrefixInfo(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoQuark.deleteLastErrPrefixInfo()\n" +
			"because parameter ePrefixDto is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	ePrefDtoQuark = errPrefixDtoQuark{}

	err = ePrefDtoQuark.emptyErrPrefInfoCollection(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoQuark.deleteLastErrPrefixInfo()\n" +
			"because parameter ePrefixDto is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	/*
		func (ePrefDtoQuark *errPrefixDtoQuark) normalizeErrPrefixDto(
			ePrefixDto *ErrPrefixDto)
	*/

	ePrefDtoQuark = errPrefixDtoQuark{}

	ePrefDtoQuark.normalizeErrPrefixDto(nil)

	/*
		func (ePrefDtoQuark *errPrefixDtoQuark) testValidityOfErrPrefixDto(
			ePrefixDto *ErrPrefixDto,
		errPrefStr string) (
			isValid bool,
			err error)
	*/

	ePrefDtoQuark = errPrefixDtoQuark{}

	_,
		err = ePrefDtoQuark.testValidityOfErrPrefixDto(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoQuark.deleteLastErrPrefixInfo()\n" +
			"because parameter ePrefixDto is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	epDto := getValidErrorPrefixDto()

	epDto.inputStrDelimiters.newLinePrefixDelimiter = ""
	epDto.inputStrDelimiters.inLinePrefixDelimiter = ""

	_,
		err = ePrefDtoQuark.testValidityOfErrPrefixDto(
		&epDto,
		funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoQuark.testValidityOfErrPrefixDto()\n" +
			"because parameter epDto.inputStrDelimiters are invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

	epDto = getValidErrorPrefixDto()

	epDto.outputStrDelimiters.newLinePrefixDelimiter = ""
	epDto.outputStrDelimiters.inLinePrefixDelimiter = ""

	_,
		err = ePrefDtoQuark.testValidityOfErrPrefixDto(
		&epDto,
		funcName)

	if err == nil {
		t.Error("ERROR\n" +
			"Expected an error return from ePrefDtoQuark.testValidityOfErrPrefixDto()\n" +
			"because parameter epDto.outputStrDelimiters are invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}
