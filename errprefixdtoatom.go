package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefixDtoAtom struct {
	lock *sync.Mutex
}

// addTwoDimensionalStringArray - Adds error prefix info contained
// in a 2-Dimensional string array.
func (ePrefixDtoAtom *errPrefixDtoAtom) addTwoDimensionalStringArray(
	errPrefixDto *ErrPrefixDto,
	twoDStrArray [][2]string,
	errPrefStr string) error {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	errPrefStr = errPrefStr +
		"\nerrPrefixDtoAtom." +
		"addTwoDimensionalStringArray()"

	if errPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefixDto' is invalid!\n"+
			"'' is a 'nil' pointer.\n",
			errPrefStr)
	}

	errPrefixDtoQuark{}.ptr().normalizeErrPrefixDto(errPrefixDto)

	if twoDStrArray == nil {
		return nil
	}

	lenTwoDStrAry := len(twoDStrArray)

	if lenTwoDStrAry == 0 {
		return nil
	}

	var ePrefInfo ErrorPrefixInfo

	for i := 0; i < lenTwoDStrAry; i++ {

		if len(twoDStrArray[i][0]) == 0 {
			continue
		}

		ePrefInfo = ErrorPrefixInfo{
			isFirstIdx:             false,
			isLastIdx:              false,
			isPopulated:            true,
			errorPrefixStr:         twoDStrArray[i][0],
			lenErrorPrefixStr:      uint(len(twoDStrArray[i][0])),
			errPrefixHasContextStr: len(twoDStrArray[i][1]) > 0,
			errorContextStr:        twoDStrArray[i][1],
			lenErrorContextStr:     uint(len(twoDStrArray[i][1])),
			lock:                   nil,
		}

		errPrefixDto.ePrefCol =
			append(errPrefixDto.ePrefCol,
				ePrefInfo)
	}

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		errPrefixDto.ePrefCol)

	return nil
}

// areEqualErrPrefDtos - Receives pointers to two ErrPrefixDto
// objects and proceeds to compare the internal data values.
// If the ErrPrefixDto objects contain data values which ARE EQUAL
// in all respects, this method returns 'true'.
//
// If the data values for the two ErrPrefixDto objects ARE NOT
// EQUAL, this method will return 'false'.
//
func (ePrefixDtoAtom *errPrefixDtoAtom) areEqualErrPrefDtos(
	errPrefixDto1 *ErrPrefixDto,
	errPrefixDto2 *ErrPrefixDto) bool {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	if errPrefixDto1 == nil ||
		errPrefixDto2 == nil {
		return false
	}

	errPrefDtoQuark := errPrefixDtoQuark{}

	errPrefDtoQuark.normalizeErrPrefixDto(errPrefixDto1)

	errPrefDtoQuark.normalizeErrPrefixDto(errPrefixDto2)

	if errPrefixDto1.leftMarginLength !=
		errPrefixDto2.leftMarginLength {
		return false
	}

	if errPrefixDto1.leftMarginChar !=
		errPrefixDto2.leftMarginChar {
		return false
	}

	if errPrefixDto1.isLastLineTerminatedWithNewLine !=
		errPrefixDto2.isLastLineTerminatedWithNewLine {
		return false
	}

	if errPrefixDto1.turnOffTextDisplay !=
		errPrefixDto2.turnOffTextDisplay {
		return false
	}

	if !errPrefixDto1.inputStrDelimiters.Equal(
		&errPrefixDto2.inputStrDelimiters) {
		return false
	}

	if !errPrefixDto1.outputStrDelimiters.Equal(
		&errPrefixDto2.outputStrDelimiters) {
		return false
	}

	if errPrefixDto1.maxErrPrefixTextLineLength !=
		errPrefixDto2.maxErrPrefixTextLineLength {
		return false
	}

	if errPrefixDto1.leadingTextStr !=
		errPrefixDto2.leadingTextStr {
		return false
	}

	if errPrefixDto1.trailingTextStr !=
		errPrefixDto2.trailingTextStr {
		return false
	}

	if errPrefixDto1.ePrefCol == nil {
		errPrefixDto1.ePrefCol =
			make([]ErrorPrefixInfo, 0)
	}

	if errPrefixDto2.ePrefCol == nil {
		errPrefixDto2.ePrefCol =
			make([]ErrorPrefixInfo, 0)
	}

	lenIncomingEPrefCol01 :=
		len(errPrefixDto1.ePrefCol)

	lenIncomingEPrefCol02 :=
		len(errPrefixDto2.ePrefCol)

	if lenIncomingEPrefCol01 != lenIncomingEPrefCol02 {
		return false
	}

	if lenIncomingEPrefCol01 == 0 {
		return true
	}

	for i := 0; i < lenIncomingEPrefCol01; i++ {
		if !errPrefixDto1.ePrefCol[i].Equal(&errPrefixDto2.ePrefCol[i]) {
			return false
		}
	}

	return true
}

// copyInErrPrefDto - Copies the data fields from an incoming
// instance of ErrPrefixDto ('inComingErrPrefixDto') to the data
// fields of the target ErrPrefixDto instance ('targetErrPrefixDto')
//
// All of the data fields in the 'targetErrPrefixDto' instance will
// be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetErrPrefixDto         *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       WILL CHANGE AND OVERWRITE all values of internal member
//       variables encapsulated in this object to achieve the
//       method's objectives.
//
//       This ErrPrefixDto instance will receive all the data
//       values contained from input parameter
//       'inComingErrPrefixDto'. When the copy operation is
//       completed, the data values in 'targetErrPrefixDto' will be
//       identical to those contained in input parameter,
//       'inComingErrPrefixDto'.
//
//
//  inComingErrPrefixDto       *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrPrefixDto instance will be
//       copied to input parameter 'targetErrPrefixDto'.
//
//
//  eMsg                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'eMsg'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'eMsg'.
//
func (ePrefixDtoAtom *errPrefixDtoAtom) copyInErrPrefDto(
	targetErrPrefixDto *ErrPrefixDto,
	inComingErrPrefixDto *ErrPrefixDto,
	eMsg string) error {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	eMsg += "errPrefixDtoAtom.copyInErrPrefDto() "

	if targetErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetErrPrefixDto' is INVALID!\n"+
			"'targetErrPrefixDto' is a nil pointer!\n",
			eMsg)

	}

	if inComingErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'inComingErrPrefixDto' is INVALID!\n"+
			"'inComingErrPrefixDto' is a nil pointer.\n",
			eMsg)
	}

	ePrefDtoQuark := errPrefixDtoQuark{}

	ePrefDtoQuark.normalizeErrPrefixDto(inComingErrPrefixDto)

	ePrefDtoQuark.normalizeErrPrefixDto(targetErrPrefixDto)

	_,
		err := ePrefDtoQuark.
		testValidityOfErrPrefixDto(
			inComingErrPrefixDto,
			eMsg+"Testing validity of inComingErrPrefixDto ")

	if err != nil {
		return err
	}

	targetErrPrefixDto.isLastLineTerminatedWithNewLine =
		inComingErrPrefixDto.isLastLineTerminatedWithNewLine

	targetErrPrefixDto.turnOffTextDisplay =
		inComingErrPrefixDto.turnOffTextDisplay

	err = targetErrPrefixDto.inputStrDelimiters.CopyIn(
		&inComingErrPrefixDto.inputStrDelimiters,
		eMsg+
			"inComingErrPrefixDto.inputStrDelimiters->"+
			"targetErrPrefixDto.inputStrDelimiters ")

	if err != nil {
		return err
	}

	err = targetErrPrefixDto.outputStrDelimiters.CopyIn(
		&inComingErrPrefixDto.outputStrDelimiters,
		eMsg+
			"inComingErrPrefixDto.outputStrDelimiters->"+
			"targetErrPrefixDto.outputStrDelimiters ")

	if err != nil {
		return err
	}

	targetErrPrefixDto.maxErrPrefixTextLineLength =
		inComingErrPrefixDto.maxErrPrefixTextLineLength

	targetErrPrefixDto.leftMarginLength =
		inComingErrPrefixDto.leftMarginLength

	targetErrPrefixDto.leftMarginChar =
		inComingErrPrefixDto.leftMarginChar

	targetErrPrefixDto.leadingTextStr =
		inComingErrPrefixDto.leadingTextStr

	targetErrPrefixDto.trailingTextStr =
		inComingErrPrefixDto.trailingTextStr

	lenIncomingEPrefCol :=
		len(inComingErrPrefixDto.ePrefCol)

	if lenIncomingEPrefCol == 0 {
		targetErrPrefixDto.ePrefCol = nil
		return nil
	}

	targetErrPrefixDto.ePrefCol =
		make(
			[]ErrorPrefixInfo,
			lenIncomingEPrefCol)

	copy(
		targetErrPrefixDto.ePrefCol,
		inComingErrPrefixDto.ePrefCol)

	return nil
}

// copyOutErrPrefDto - Creates a deep copy of the data fields contained in
// input parameter 'ePrefixDto' and returns that data as a new
// instance ErrPrefixDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this object.
//
//       If this ErrPrefixDto instance proves to be invalid, an
//       error will be returned.
//
//
//  eMsg                string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'eMsg'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newEPrefixDto       ErrPrefixDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'ePrefixDto' will be returned in a new
//       instance of ErrPrefixDto
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'eMsg'. The
//       'eMsg' text will be prefixed to the beginning of the returned
//       error message.
//
func (ePrefixDtoAtom *errPrefixDtoAtom) copyOutErrPrefDto(
	ePrefixDto *ErrPrefixDto,
	eMsg string) (
	newEPrefixDto ErrPrefixDto,
	err error) {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	eMsg += "errPrefixDtoAtom.copyOutErrPrefDto() "

	ePrefDtoQuark := errPrefixDtoQuark{}

	newEPrefixDto = ePrefDtoQuark.
		newZeroErrPrefixDto()

	if ePrefixDto == nil {
		err =
			fmt.Errorf("%v\n"+
				"\nInput parameter 'ePrefixDto' is INVALID!\n"+
				"'ePrefixDto' is a nil pointer!\n",
				eMsg)

		return newEPrefixDto, err
	}

	ePrefDtoQuark.normalizeErrPrefixDto(ePrefixDto)

	newEPrefixDto.isLastLineTerminatedWithNewLine =
		ePrefixDto.isLastLineTerminatedWithNewLine

	newEPrefixDto.turnOffTextDisplay =
		ePrefixDto.turnOffTextDisplay

	err =
		newEPrefixDto.inputStrDelimiters.CopyIn(
			&ePrefixDto.inputStrDelimiters,
			eMsg+"ePrefixDto.inputStrDelimiters->"+
				"newEPrefixDto.inputStrDelimiters ")

	if err != nil {
		return newEPrefixDto, err
	}

	err =
		newEPrefixDto.outputStrDelimiters.CopyIn(
			&ePrefixDto.outputStrDelimiters,
			eMsg+"ePrefixDto.outputStrDelimiters->"+
				"newEPrefixDto.outputStrDelimiters ")

	if err != nil {
		return newEPrefixDto, err
	}

	newEPrefixDto.maxErrPrefixTextLineLength =
		ePrefixDto.maxErrPrefixTextLineLength

	newEPrefixDto.leftMarginLength =
		ePrefixDto.leftMarginLength

	newEPrefixDto.leftMarginChar =
		ePrefixDto.leftMarginChar

	newEPrefixDto.leadingTextStr =
		ePrefixDto.leadingTextStr

	newEPrefixDto.trailingTextStr =
		ePrefixDto.trailingTextStr

	lenIncomingEPrefCol :=
		len(ePrefixDto.ePrefCol)

	if lenIncomingEPrefCol == 0 {

		newEPrefixDto.ePrefCol = nil

		return newEPrefixDto, err
	}

	newEPrefixDto.ePrefCol =
		make(
			[]ErrorPrefixInfo,
			lenIncomingEPrefCol)

	copy(
		newEPrefixDto.ePrefCol,
		ePrefixDto.ePrefCol)

	return newEPrefixDto, err
}

// getEPrefContextArray - Receives a string containing a series of
// error prefix strings and error context strings. This method then
// proceeds to parse the error prefix and error context elements
// returning the results in an array of ErrorPrefixInfo objects
// passed as input parameter, 'prefixContextCol' .
//
// All error prefix elements parsed from the 'errPrefix' string
// will be appended to 'prefixContextCol'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefix           string
//     - This string contains a series of error prefix and error
//       context strings. Input parameter 'errPrefix' will be
//       parsed into error prefix and error context components
//       before being appended to array of ErrorPrefixInfo objects
//       passed in input parameter, 'prefixContextCol'
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited by either a
//       new line character ('\n') or the in-line delimiter string,
//       " - ".
//
//       If this string contains associated error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
//  prefixContextCol    *[]ErrorPrefixInfo
//     - A pointer to an array of ErrorPrefixInfo objects. The
//       error prefix and error context data elements parsed from
//       input parameter, 'errPrefix', will be appended to this
//       array.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefixDtoAtom *errPrefixDtoAtom) getEPrefContextArray(
	errPrefix string,
	delimiters ErrPrefixDelimiters,
	prefixContextCol *[]ErrorPrefixInfo) {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	if *prefixContextCol == nil {
		*prefixContextCol = make([]ErrorPrefixInfo, 0)
	}

	if len(errPrefix) == 0 {
		return
	}

	ePrefElectron := errPrefElectron{}

	var lenCleanPrefixStr int

	errPrefix,
		lenCleanPrefixStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanPrefixStr == 0 {
		return
	}

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetNewLineContextDelimiter(),
		delimiters.GetInLineContextDelimiter())

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetInLinePrefixDelimiter(),
		delimiters.GetNewLinePrefixDelimiter())

	errPrefix += delimiters.GetNewLinePrefixDelimiter()

	errPrefixContextCollection := strings.Split(
		errPrefix,
		delimiters.GetNewLinePrefixDelimiter())

	lCollection := len(errPrefixContextCollection)

	if lCollection == 0 {
		return
	}

	var contextIdx int
	var idxLenInLineContextDelimiter = int(delimiters.GetLengthInLineContextDelimiter() - 1)
	var (
		errPrefixStr,
		errCtxStr string
	)

	ePrefNeutron := errPrefNeutron{}

	for i := 0; i < lCollection; i++ {

		s := errPrefixContextCollection[i]

		contextIdx = strings.Index(s,
			delimiters.GetInLineContextDelimiter())

		if contextIdx == -1 {

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				s,
				"",
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		} else {

			errPrefixStr =
				s[0:contextIdx]

			errCtxStr = s[contextIdx+
				idxLenInLineContextDelimiter:]

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				errPrefixStr,
				errCtxStr,
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		}
	}

	return
}

// ptr() - Returns a pointer to a new instance of
// errPrefixDtoAtom.
//
func (ePrefixDtoAtom errPrefixDtoAtom) ptr() *errPrefixDtoAtom {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	return &errPrefixDtoAtom{}
}

// setFlagsErrorPrefixInfoArray - Sets internal flags in an array
// of ErrorPrefixInfo objects
//
func (ePrefixDtoAtom *errPrefixDtoAtom) setFlagsErrorPrefixInfoArray(
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefixDtoAtom.lock == nil {
		ePrefixDtoAtom.lock = new(sync.Mutex)
	}

	ePrefixDtoAtom.lock.Lock()

	defer ePrefixDtoAtom.lock.Unlock()

	if prefixContextCol == nil {
		prefixContextCol = make([]ErrorPrefixInfo, 0)
	}

	lenCollection := len(prefixContextCol)

	if lenCollection == 0 {
		return
	}

	if lenCollection == 1 {
		prefixContextCol[0].SetIsFirstIndex(true)
		prefixContextCol[0].SetIsLastIndex(true)
		return
	}

	lastIdx := lenCollection - 1

	for i := 0; i < lenCollection; i++ {

		_ = prefixContextCol[i].GetIsPopulated()

		switch i {
		case 0:
			prefixContextCol[i].SetIsFirstIndex(true)
			prefixContextCol[i].SetIsLastIndex(false)
		case lastIdx:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(true)
		default:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(false)
		}
	}

	return
}
