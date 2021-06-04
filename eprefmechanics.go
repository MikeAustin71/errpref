package errpref

import (
	"sync"
)

type errPrefMechanics struct {
	lock *sync.Mutex
}

// assembleErrPrefix - Assembles, consolidates and formats an error
// prefix string from a previous error prefix, a new error prefix
// and the new error context.
//
// This method applies default string delimiters when parsing error
// prefix strings.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix                  string
//     - The existing or previous error prefix string. This text
//       string usually consists of a series of function names and
//       associated error context strings.
//
//
//  newErrPrefix                  string
//     - This is the new error prefix string which will be added to
//       existing error prefix string represented by input
//       parameter, 'oldErrPref'.
//
//
//  newErrContext                 string
//     - An optional error context description. This is the error
//       context information associated with the new error prefix
//       ('newErrPref'). Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of the function
//       documented by error prefix 'newErrPrefix'.
//
//
//  maxErrPrefixTextLineLength    uint
//      - Specifies the maximum number of text characters which can
//        be on a single line for a new line character ('\n') is
//        inserted. If this value is zero, it will be set to the
//        default value of 40.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will concatenate the old error prefix, new
//       error prefix and the new error context. This consolidated
//       error prefix will then be returned to the calling
//       function.
//
func (ePrefMech *errPrefMechanics) assembleErrPrefix(
	oldErrPrefix string,
	newErrPrefix string,
	newErrContext string,
	maxErrStringLength uint,
	delimiters ErrPrefixDelimiters) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	if maxErrStringLength == 0 {
		maxErrStringLength =
			errPrefQuark{}.ptr().getErrPrefDisplayLineLength()
	}

	var lenOldErrPrefCleanStr int
	var errPrefixInfo ErrorPrefixInfo

	eMsgPrefix := "errPrefMechanics.assembleErrPrefix() "

	oldErrPrefix,
		lenOldErrPrefCleanStr =
		errPrefElectron{}.ptr().cleanErrorPrefixStr(oldErrPrefix)

	errPrefixInfo,
		_ = errPrefNeutron{}.ptr().createNewEPrefInfo(
		newErrPrefix,
		newErrContext,
		eMsgPrefix)

	var lenPrefixContextCol int

	prefixContextCol := make([]ErrorPrefixInfo, 0)

	lenPrefixContextCol = 0

	if lenOldErrPrefCleanStr > 0 {

		errPrefixDtoAtom{}.ptr().getEPrefContextArray(
			oldErrPrefix,
			delimiters,
			&prefixContextCol)

		lenPrefixContextCol = len(prefixContextCol)

		if errPrefixInfo.GetIsPopulated() &&
			lenPrefixContextCol > 0 {

			prefixContextCol[lenPrefixContextCol-1].
				SetIsLastIndex(false)
		}
	}

	if errPrefixInfo.GetIsPopulated() {

		errPrefixInfo.SetIsLastIndex(true)

		prefixContextCol = append(prefixContextCol, errPrefixInfo)
	}

	if len(prefixContextCol) == 0 {
		return ""
	}

	return errPrefNanobot{}.ptr().formatErrPrefixComponents(
		"", // Leading Text String
		"", // Trailing Text String
		maxErrStringLength,
		false,
		delimiters,
		prefixContextCol)
}

// formatErrPrefix - Returns a string of formatted error prefix information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  maxErrStringLength            uint
//      - Specifies the maximum number of text characters which can
//        be on a single line for a new line character ('\n') is
//        inserted. If this value is zero, it will be set to the
//        default value of 40.
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
//  errPrefix           string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'errPrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, a formatted string
//       containing error prefix and context information will be
//       returned to the caller. If an error occurs, this string
//       will be populated with an error message.
//
func (ePrefMech *errPrefMechanics) formatErrPrefix(
	maxErrStringLength uint,
	delimiters ErrPrefixDelimiters,
	errPrefix string) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	if maxErrStringLength == 0 {
		maxErrStringLength = errPrefQuark{}.ptr().getErrPrefDisplayLineLength()
	}

	localErrPrefix := "errPrefMechanics.formatErrPrefix() "

	prefixContextCol := make([]ErrorPrefixInfo, 0)

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		errPrefix,
		delimiters,
		&prefixContextCol)

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		prefixContextCol)

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			"", // Leading Text String
			"", // Trailing Text String
			maxErrStringLength,
			false,
			delimiters,
			prefixContextCol)
}

// extractLastErrPrefCtxPair - Extracts the last error prefix and
// error context combination from a series of error prefix/error
// context pairs contained in input parameter 'oldErrPref'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  maxErrStringLength            uint
//      - Specifies the maximum number of text characters which can
//        be on a single line for a new line character ('\n') is
//        inserted. If this value is zero, it will be set to the
//        default value of 40.
//
//
//  oldErrPref                    string
//     - The existing or previous error prefix string. This text
//       string usually consists of a series of function names and
//       associated error context strings.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, the last error
//       prefix and error context combination extracted from a
//       series of error prefix/error context pairs contained in
//       input parameter 'oldErrPref' will be returned to the
//       caller. If an error occurs, this string will be populated
//       with an error message.
//
func (ePrefMech *errPrefMechanics) extractLastErrPrefCtxPair(
	maxErrStringLength uint,
	oldErrPref string,
	delimiters ErrPrefixDelimiters) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	errPrefInfo := errPrefNanobot{}.ptr().
		extractLastErrPrfInfo(oldErrPref)

	if !errPrefInfo.GetIsPopulated() {
		return ""
	}

	prefixContextCol := make([]ErrorPrefixInfo, 0)

	prefixContextCol = append(prefixContextCol, errPrefInfo)

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			"", // Leading Text String
			"", // Trailing Text String
			maxErrStringLength,
			false,
			delimiters,
			prefixContextCol)
}

// ptr() - Returns a pointer to a new instance of
// errPrefMechanics.
//
func (ePrefMech errPrefMechanics) ptr() *errPrefMechanics {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	return &errPrefMechanics{}
}

// setErrorContext - Parses strings to generated formatted
// error prefix and context information.
//
// This method applies system default string delimiters
// when parsing error prefix and context strings.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix                  string
//     - The existing or previous error prefix string. This text
//       string usually consists of a series of function names and
//
//
//  newErrContext                 string
//     - An optional error context description. This is the error
//       context information associated with the last error prefix
//       element in parameter, 'oldErrPref'. Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of the function
//       documented by error prefix 'newErrPrefix'.
func (ePrefMech *errPrefMechanics) setErrorContext(
	oldErrPref string,
	newErrContext string,
	maxErrStringLength uint,
	delimiters ErrPrefixDelimiters) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	if maxErrStringLength == 0 {
		maxErrStringLength =
			errPrefQuark{}.ptr().getErrPrefDisplayLineLength()
	}

	var (
		lenOldErrPrefCleanStr,
		lenNewErrContextCleanStr,
		lenPrefixContextCol int
	)

	ePrefElectron := errPrefElectron{}

	oldErrPref,
		lenOldErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	if lenOldErrPrefCleanStr == 0 {
		return ""
	}

	newErrContext,
		lenNewErrContextCleanStr =
		ePrefElectron.cleanErrorContextStr(newErrContext)

	if lenNewErrContextCleanStr == 0 {
		return oldErrPref
	}

	prefixContextCol := make([]ErrorPrefixInfo, 0)

	errPrefixDtoAtom{}.ptr().getEPrefContextArray(
		oldErrPref,
		delimiters,
		&prefixContextCol)

	lenPrefixContextCol = len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return ""
	}

	lenPrefixContextCol--

	prefixContextCol[lenPrefixContextCol].
		SetErrPrefixHasContext(true)

	prefixContextCol[lenPrefixContextCol].
		SetErrContextStr(newErrContext)

	return errPrefNanobot{}.ptr().formatErrPrefixComponents(
		"", // Leading Text String
		"", // Trailing Text String
		maxErrStringLength,
		false,
		delimiters,
		prefixContextCol)
}
