package errpref

import (
	"strings"
	"sync"
)

type errPrefNanobot struct {
	lock *sync.Mutex
}

// addEPrefInfo - Receives new error prefix and context strings.
// The method then proceeds to convert this pair of strings to an
// ErrorPrefixInfo object before adding that object to the end of
// an array of ErrorPrefixInfo objects.
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix string. Typically this is the
//       name of the function or method associated which is
//       currently executing.
//
//
//  newErrContext       string
//     - This is the error context information associated with the
//       new error prefix ('newErrPrefix'). This parameter is
//       optional and will accept an empty string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefNanobot *errPrefNanobot) addEPrefInfo(
	newErrPrefix string,
	newErrContext string,
	errPrefixCollection *[]ErrorPrefixInfo) {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	var (
		lenPrefixCleanStr,
		lenContextCleanStr int
	)

	ePrefElectron := errPrefElectron{}

	newErrPrefix,
		lenPrefixCleanStr =
		ePrefElectron.cleanErrorPrefixStr(
			newErrPrefix)

	if lenPrefixCleanStr == 0 {
		return
	}

	newErrContext,
		lenContextCleanStr =
		ePrefElectron.cleanErrorContextStr(
			newErrContext)

	newErrorPrefixInfo := ErrorPrefixInfo{}

	newErrorPrefixInfo.SetErrPrefixStr(
		newErrPrefix)

	if lenContextCleanStr > 0 {

		newErrorPrefixInfo.SetErrContextStr(
			newErrContext)

	}

	if *errPrefixCollection == nil {
		*errPrefixCollection = make([]ErrorPrefixInfo, 0)
	}

	*errPrefixCollection = append(
		*errPrefixCollection,
		newErrorPrefixInfo)

	return
}

// extractLastErrPrfInfo - Extracts the last error prefix element
// from a string comprised of a series of error prefix elements.
//
// This method applies system default string delimiters when
// parsing error prefix strings.
//
func (ePrefNanobot *errPrefNanobot) extractLastErrPrfInfo(
	errPref string) ErrorPrefixInfo {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	prefixContextCol := make([]ErrorPrefixInfo, 0)

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	errPrefixDtoAtom{}.
		ptr().getEPrefContextArray(
		errPref,
		delimiters,
		&prefixContextCol)

	lenCollection := len(prefixContextCol)

	if lenCollection == 0 {
		return ErrorPrefixInfo{}
	}

	lastEPref := prefixContextCol[lenCollection-1]

	lastEPref.SetIsLastIndex(false)
	lastEPref.SetIsFirstIndex(false)

	return lastEPref
}

// formatErrPrefixComponents - Returns a string of formatted error
// prefix information based on an array of error prefix elements.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  maxErrPrefixTextLineLength         uint
//     - This unsigned integer value will be used to set the
//       maximum number of characters allowed in a text display
//       line for error prefix information.
//
//       If 'maxErrPrefixTextLineLength' is set to a value of zero
//       (0), this method automatically set this value to the
//       current default maximum line length.
//
//
//  isLastLineTerminatedWithNewLine    bool
//     - By default, the last line of error prefix strings ARE NOT
//       terminated with a new line character ('\n'). In other
//       words, by default, the last line of returned error prefix
//       strings do not end with a new line character ('\n').
//
//       If this parameter is set to 'true', the last line of the
//       error prefix strings returned by this method WILL BE
//       terminated with a new line character ('\n').
//
//
//  delimiters                    ErrPrefixDelimiters
//     - An instance of ErrPrefixDelimiters containing string
//       delimiters used to join error prefix and error context
//       elements.
//
//       The key components of an ErrPrefixDelimiters object are
//       listed as follows:
//         New Line Error Prefix Delimiter
//         In-Line Error Prefix Delimiter
//         New Line Error Context Delimiter
//         In-Line Error Context Delimiter
//
//
//  prefixContextCol                   []ErrorPrefixInfo
//     - An array of ErrorPrefixInfo objects containing the error
//       prefix and error context information which will be
//       converted to a single string returned by this method.
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return a string containing formatted
//       error prefix and error context information.
//
func (ePrefNanobot *errPrefNanobot) formatErrPrefixComponents(
	maxErrPrefixTextLineLength uint,
	isLastLineTerminatedWithNewLine bool,
	delimiters ErrPrefixDelimiters,
	prefixContextCol []ErrorPrefixInfo) string {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	localErrPrefix := "errPrefNanobot.formatErrPrefixComponents() "

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	err := delimiters.IsValidInstanceError(
		localErrPrefix)

	if err != nil {
		return err.Error()
	}

	if maxErrPrefixTextLineLength == 0 {
		maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().
				getErrPrefDisplayLineLength()
	}

	lineLenCalculator := EPrefixLineLenCalc{}

	err =
		lineLenCalculator.SetEPrefDelimiters(
			delimiters,
			localErrPrefix)

	if err != nil {
		return err.Error()
	}

	lineLenCalculator.SetMaxErrStringLength(
		maxErrPrefixTextLineLength)

	var b1 strings.Builder

	b1.Grow(1024)

	lineLenCalculator.SetCurrentLineStr("")

	lastIdx := lenPrefixContextCol - 1

	ePrefMolecule := errPrefMolecule{}

	for i := 0; i < lenPrefixContextCol; i++ {

		if i == lastIdx {
			prefixContextCol[i].SetIsLastIndex(true)
		}

		err =
			lineLenCalculator.SetErrPrefixInfo(
				&prefixContextCol[i],
				localErrPrefix)

		if err != nil {
			return err.Error()
		}

		err =
			lineLenCalculator.IsValidInstanceError(
				localErrPrefix)

		if err != nil {
			return err.Error()
		}

		if lineLenCalculator.ErrPrefixHasContext() {

			err = ePrefMolecule.writeNewEPrefWithContext(
				&b1,
				&lineLenCalculator)

			if err != nil {
				return err.Error()
			}

			continue
		} else {

			err = ePrefMolecule.writeNewEPrefWithOutContext(
				&b1,
				&lineLenCalculator)

			if err != nil {
				return err.Error()
			}
		}

	}

	if lineLenCalculator.GetCurrLineStrLength() > 0 {

		errPrefNeutron{}.ptr().writeCurrentLineStr(
			&b1,
			&lineLenCalculator)
	}

	if isLastLineTerminatedWithNewLine {
		b1.WriteRune('\n')
	}

	lineLenCalculator.Empty()

	return b1.String()
}

// ptr() - Returns a pointer to a new instance of
// errPrefNanobot.
//
func (ePrefNanobot errPrefNanobot) ptr() *errPrefNanobot {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	return &errPrefNanobot{}
}

// setLastCtx - Sets or resets the error context for the last
// error prefix in the error prefix collection passed as an input
// parameter. This operation either adds, or replaces, the error
// context string associated with the last error prefix the
// current list of error prefixes contained in the error prefix
// collection.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in the 'errPrefixCollection' already
//       has an associated error context, that context will be
//       deleted and replaced by 'newErrContext'. If, however, the
//       last error prefix does NOT have an associated error
//       context, this 'newErrContext' string will be added and
//       associated with that last error prefix.
//
//       If 'newErrContext' is 'empty', this method will take no
//       action and exit.
//
//
//  errPrefixCollection []ErrorPrefixInfo
//      - An array of ErrorPrefixInfo objects. Each object defines
//        an error prefix/context pair. The 'newErrContext' string
//        will be configured as the error context data for the last
//        ErrorPrefixInfo object in this collection.
//
//        If the 'errPrefixCollection' is empty, this method will
//        take no action and exit.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//
func (ePrefNanobot *errPrefNanobot) setLastCtx(
	newErrContext string,
	errPrefixCollection []ErrorPrefixInfo) {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	var (
		lenCleanStr,
		lenCollection int
	)

	newErrContext,
		lenCleanStr =
		errPrefElectron{}.ptr().cleanErrorContextStr(
			newErrContext)

	if lenCleanStr == 0 {
		return
	}

	if errPrefixCollection == nil {
		errPrefixCollection = make([]ErrorPrefixInfo, 0)
	}

	lenCollection = len(errPrefixCollection)

	if lenCollection == 0 {
		return
	}

	lenCollection--

	errPrefixCollection[lenCollection].
		SetErrContextStr(newErrContext)

	return
}
