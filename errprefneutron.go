package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefNeutron struct {
	lock *sync.Mutex
}

// createNewEPrefInfo - Creates a single ErrorPrefixInfo object
// from an error prefix and an error context string. This method is
// only designed to operate on a single pair of pair of strings
// containing one error prefix and one error context.
//
// The error context string is optional and empty error strings
// will still result in the generation and return of a valid
// ErrorPrefixInfo object.
//
// If the input parameter, 'newErrPrefix' is an empty string, this
// method will generate an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - This is the new error prefix string which will be
//       configured in the ErrorPrefixInfo object returned by this
//       method.
//
//
//  newErrContext       string
//     - An optional error context description. This is the error
//       context information associated with the new error prefix
//       ('newErrPrefix'). Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of the function
//       documented by error prefix 'newErrPrefix'.
//
//
//  ePrefixStr          string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  errPrefixInfo       ErrorPrefixInfo
//     - This returned ErrorPrefixInfo will encapsulate the error
//       prefix and error context strings passed through input
//       parameters 'newErrPrefix' and 'newErrContext'.
//
//
//  err                 error
//     - If this method encounters a processing error during
//       execution, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (ePrefNeutron *errPrefNeutron) createNewEPrefInfo(
	newErrPrefix string,
	newErrContext string,
	ePrefixStr string) (
	errPrefixInfo ErrorPrefixInfo,
	err error) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	ePrefixStr += "errPrefNeutron.createNewEPrefInfo()"

	errPrefixInfo = ErrorPrefixInfo{}

	var (
		lenPrefixStr,
		lenContextStr int
	)

	ePrefElectron := errPrefElectron{}

	newErrPrefix,
		lenPrefixStr = ePrefElectron.cleanErrorPrefixStr(newErrPrefix)

	if lenPrefixStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Cleaned Error Prefix String is EMPTY!\n",
			ePrefixStr)

		return errPrefixInfo, err
	}

	newErrContext,
		lenContextStr = ePrefElectron.cleanErrorContextStr(newErrContext)

	errPrefixInfo.SetIsFirstIndex(false)

	errPrefixInfo.SetIsLastIndex(false)

	errPrefixInfo.SetErrPrefixStr(newErrPrefix)

	if lenContextStr > 0 {
		errPrefixInfo.SetErrPrefixHasContext(true)
		errPrefixInfo.SetErrContextStr(newErrContext)
	} else {
		errPrefixInfo.SetErrPrefixHasContext(false)
	}

	return errPrefixInfo, err
}

// ptr() - Returns a pointer to a new instance of
// errPrefNeutron.
//
func (ePrefNeutron errPrefNeutron) ptr() *errPrefNeutron {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	return &errPrefNeutron{}
}

// writeCurrentLineStr - Writes the contents of a current
// line of error prefix and error context characters to
// a sting builder for text display output.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder          *strings.Builder
//     - A pointer to a string builder. The contents of the current
//       line of text will be written to this string builder.
//
//
//  ePrefLineLenCalc    *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc, the Error
//       Prefix Line Length Calculator. This types encapsulates all
//       the data necessary to perform line length calculations and
//       format the error prefix and error context strings for
//       text display output.
//
//       type EPrefixLineLenCalc struct {
//         ePrefDelimiters    ErrPrefixDelimiters
//         errorPrefixInfo    *ErrorPrefixInfo
//         currentLineStr     string
//         maxErrStringLength uint
//       }
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefNeutron *errPrefNeutron) writeCurrentLineStr(
	strBuilder *strings.Builder,
	ePrefLineLenCalc *EPrefixLineLenCalc) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	if strBuilder == nil ||
		ePrefLineLenCalc == nil ||
		!ePrefLineLenCalc.IsValidInstance() {
		return
	}

	if ePrefLineLenCalc.GetCurrLineStrLength() == 0 {
		return
	}

	if strBuilder.Len() > 0 {
		strBuilder.WriteString(
			ePrefLineLenCalc.GetDelimiterNewLineErrPrefix())
	}

	strBuilder.WriteString(
		ePrefLineLenCalc.GetCurrLineStr())

	//if !ePrefLineLenCalc.IsErrPrefixLastIndex() {
	//	strBuilder.WriteString(
	//		ePrefLineLenCalc.GetDelimiterNewLineErrPrefix())
	//}

	ePrefLineLenCalc.SetCurrentLineStr("")

	return
}
