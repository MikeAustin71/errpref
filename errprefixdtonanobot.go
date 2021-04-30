package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefixDtoNanobot struct {
	lock *sync.Mutex
}

// deleteLastErrContext - Deletes the Last Context for the
// last error prefix in the sequence.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. The last error
//       context for the last error prefix contained in this
//       ErrPrefixDto object will be deleted.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) deleteLastErrContext(
	ePrefixDto *ErrPrefixDto) {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	if ePrefixDto == nil {
		return
	}

	ePrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	collectionIdx := len(ePrefixDto.ePrefCol)

	if collectionIdx == 0 {
		return
	}

	collectionIdx--

	ePrefixDto.ePrefCol[collectionIdx].
		SetErrContextStr("")

	return
}

// ptr - Returns a pointer to a new instance of errPrefixDtoNanobot.
//
func (ePrefixDtoNanobot errPrefixDtoNanobot) ptr() *errPrefixDtoNanobot {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	return &errPrefixDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// setFromIBuilder - Receives a pointer to an ErrPrefixDto instance
// plus an object implementing the IBuilderErrorPrefix interface.
//
// This method then proceeds to copy an transfer all error prefix
// and error context information from the IBuilderErrorPrefix
// object to the ErrPrefixDto instance.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              IBuilderErrorPrefix
//     - An object which implements the IBuilderErrorPrefix
//       interface.
//
//       This interface contains a method named,
//       'GetEPrefStrings()'. This method returns a two dimensional
//       array of strings. Each 2-Dimensional array element holds a
//       separate string for error prefix and error context.
//
//       Information extracted from this object will generate error
//       prefix information used to overwrite and replace the error
//       prefix and error context information in the input
//       parameter, 'errPrefDto'.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromIBuilder(
	errPrefDto *ErrPrefixDto,
	iEPref IBuilderErrorPrefix,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBuilder()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == (IBuilderErrorPrefix)(nil) {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' has a nil value.\n",
			methodName)
	}

	err := errPrefixDtoQuark{}.ptr().emptyErrPrefInfoCollection(
		errPrefDto,
		methodName)

	if err != nil {
		return err
	}

	err = errPrefixDtoAtom{}.ptr().addTwoDimensionalStringArray(
		errPrefDto,
		iEPref.GetEPrefStrings(),
		methodName)

	return err
}

// setFromIBasicErrorPrefix - Receives an ErrPrefixDto object and
// then proceeds to overwrite and reset the ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from input
// parameter, 'iEPref' which implements the IBasicErrorPrefix
// interface.
//
// The IBasicErrorPrefix interface contains the method:
//  GetEPrefStrings() [][2]string
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance (errPrefDto) will be overwritten and
// deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              IBasicErrorPrefix
//     - An object which implements the IBasicErrorPrefix
//       interface.
//
//       This interface contains one method, 'GetEPrefStrings()'.
//       This single method returns a two dimensional array of
//       strings. Each 2-Dimensional array element holds a separate
//       string for error prefix and error context.
//
//       Information extracted from this object will generate error
//       prefix information used to overwrite and replace the error
//       prefix and error context information in the input
//       parameter, 'errPrefDto'.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromIBasicErrorPrefix(
	errPrefDto *ErrPrefixDto,
	iEPref IBasicErrorPrefix,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == (IBasicErrorPrefix)(nil) {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' has a nil value.\n",
			methodName)
	}

	errPrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	errPrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	twoDSlice := iEPref.GetEPrefStrings()

	lenTwoDSlice := len(twoDSlice)

	if lenTwoDSlice == 0 {
		errPrefDto.ePrefCol =
			make([]ErrorPrefixInfo, 0)
		return nil
	}

	errPrefDto.ePrefCol =
		make([]ErrorPrefixInfo, lenTwoDSlice)

	var ePrefInfo ErrorPrefixInfo

	for i := 0; i < lenTwoDSlice; i++ {

		ePrefInfo = ErrorPrefixInfo{
			isFirstIdx:             false,
			isLastIdx:              false,
			isPopulated:            true,
			errorPrefixStr:         twoDSlice[i][0],
			lenErrorPrefixStr:      uint(len(twoDSlice[i][0])),
			errPrefixHasContextStr: len(twoDSlice[i][1]) > 0,
			errorContextStr:        twoDSlice[i][1],
			lenErrorContextStr:     uint(len(twoDSlice[i][1])),
			lock:                   nil,
		}

		errPrefDto.ePrefCol[i] = ePrefInfo
	}

	errPrefixDtoAtom{}.ptr().
		setFlagsErrorPrefixInfoArray(
			errPrefDto.ePrefCol)

	return nil
}

// setFromString - Receives an ErrPrefixDto object and then
// proceeds to overwrite and reset the ErrorPrefixInfo collection
// containing error prefix and error context information. New
// error prefix information will be extracted from input
// parameter, 'iEPref', a string which contains error prefix
// information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              string
//     - A string containing error prefix information which will be
//       extracted and used to overwrite the existing error prefix
//       information contained in input parameter, 'errPrefDto'.
//
//       If this string is delimited by new line characters ("\n")
//       an array of strings will be created and used to create
//       multiple error prefix entries in the 'errPrefDto'
//       ErrorPrefixInfo collection.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromString(
	errPrefDto *ErrPrefixDto,
	iEPref string,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if len(iEPref) == 0 {

		errPrefDto.ePrefCol = nil

		return nil
	}

	ePrefAtom := errPrefixDtoAtom{}

	errPrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	errPrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom.getEPrefContextArray(
		iEPref,
		errPrefDto.inputStrDelimiters,
		&errPrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		errPrefDto.ePrefCol)

	return nil
}

// setFromStringBuilder - Receives an ErrPrefixDto object and then
// proceeds to overwrite and reset the ErrorPrefixInfo collection
// containing error prefix and error context information. New
// error prefix information will be extracted from input parameter,
// 'iEPref', a pointer to a string builder which contains error
// prefix information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              *strings.Builder
//     - A pointer to a string Builder object. The object's
//       internal string value contains error prefix information
//       which will be extracted and used to overwrite the existing
//       error prefix information contained in input parameter,
//       'errPrefDto'.
//
//       If this string is delimited by new line characters ("\n")
//       an array of strings will be created and used to create
//       multiple error prefix entries in the 'errPrefDto'
//       ErrorPrefixInfo collection.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromStringBuilder(
	errPrefDto *ErrPrefixDto,
	iEPref *strings.Builder,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' is a 'nil' pointer.\n",
			methodName)
	}

	if (*strings.Builder)(nil) == iEPref {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"Pointer 'iEPref' points to a nil value.\n",
			methodName)
	}

	strVal := iEPref.String()

	if len(strVal) == 0 {

		errPrefDto.ePrefCol = nil

		return nil

	}

	ePrefAtom := errPrefixDtoAtom{}

	errPrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	errPrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom.getEPrefContextArray(
		strVal,
		errPrefDto.inputStrDelimiters,
		&errPrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		errPrefDto.ePrefCol)

	return nil
}

// setFromStringArray - Receives an ErrPrefixDto object and then
// proceeds to overwrite and reset the ErrorPrefixInfo collection
// containing error prefix and error context information. New
// error prefix information will be extracted from input
// parameter, 'iEPref', a string array which contains error prefix
// information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              []string
//     - A string containing error prefix information which will be
//       extracted and used to overwrite the existing error prefix
//       information contained in input parameter, 'errPrefDto'.
//
//       If this string is delimited by new line characters ("\n")
//       an array of strings will be created and used to create
//       multiple error prefix entries in the 'errPrefDto'
//       ErrorPrefixInfo collection.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromStringArray(
	errPrefDto *ErrPrefixDto,
	iEPref []string,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	errPrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	errPrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	lenEPrefStrs := len(iEPref)

	if lenEPrefStrs == 0 {

		errPrefDto.ePrefCol = nil

		return nil
	}

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenEPrefStrs)

	for i := 0; i < lenEPrefStrs; i++ {

		errPrefDto.ePrefCol[i] =
			ErrorPrefixInfo{
				isFirstIdx:             false,
				isLastIdx:              false,
				isPopulated:            true,
				errorPrefixStr:         iEPref[i],
				lenErrorPrefixStr:      uint(len(iEPref[i])),
				errPrefixHasContextStr: false,
				errorContextStr:        "",
				lenErrorContextStr:     0,
				lock:                   nil,
			}
	}

	errPrefixDtoAtom{}.ptr().
		setFlagsErrorPrefixInfoArray(
			errPrefDto.ePrefCol)

	return nil
}

// setFromTwoDStrArray - Receives an ErrPrefixDto object and
// then proceeds to overwrite and reset the ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from the
// two two dimensional string array contained in input parameter,
// 'iEPref'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              [][2]string
//     - A two dimensional string array containing error prefix and
//       error context strings in pairs.
//
//       iEPref[x][0] - contains an error prefix string.
//       iEPref[x][0] - contains an error context string if one
//                      exists. Otherwise, it is populated with an
//                      empty string.
//
//       Information extracted from this array will be used to
//       generate and transfer error prefix and error context
//       information to the ErrPrefixDto input parameter,
//       'errPrefDto'.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefixDtoNanobot *errPrefixDtoNanobot) setFromTwoDStrArray(
	errPrefDto *ErrPrefixDto,
	iEPref [][2]string,
	errorPrefStr string) error {

	if ePrefixDtoNanobot.lock == nil {
		ePrefixDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefixDtoNanobot.lock.Lock()

	defer ePrefixDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromTwoDStrArray()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	errPrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	errPrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	if iEPref == nil {
		errPrefDto.ePrefCol = nil
		return nil
	}

	lenTwoDSlice := len(iEPref)

	if lenTwoDSlice == 0 {
		errPrefDto.ePrefCol = nil
		return nil
	}

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenTwoDSlice)

	var ePrefInfo ErrorPrefixInfo

	for i := 0; i < lenTwoDSlice; i++ {

		ePrefInfo = ErrorPrefixInfo{
			isFirstIdx:             false,
			isLastIdx:              false,
			isPopulated:            true,
			errorPrefixStr:         iEPref[i][0],
			lenErrorPrefixStr:      uint(len(iEPref[i][0])),
			errPrefixHasContextStr: len(iEPref[i][1]) > 0,
			errorContextStr:        iEPref[i][1],
			lenErrorContextStr:     uint(len(iEPref[i][1])),
			lock:                   nil,
		}

		errPrefDto.ePrefCol[i] = ePrefInfo
	}

	errPrefixDtoAtom{}.ptr().
		setFlagsErrorPrefixInfoArray(
			errPrefDto.ePrefCol)

	return nil
}
