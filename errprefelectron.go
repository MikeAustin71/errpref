package errpref

import (
	"strings"
	"sync"
)

type errPrefElectron struct {
	lock *sync.Mutex
}

// cleanErrorPrefixStr - Receives error prefix strings (not error
// context strings) and proceeds to scrub and remove invalid
// sub-strings and characters.
//
// This is a 'trim' operation which removes the target strings from
// both sides of the 'errPref' string.
//
// After this process is completed, the 'cleaned' string along with
// the length of the newly cleaned string is returned to the
// calling function.
//
func (ePrefElectron *errPrefElectron) cleanErrorPrefixStr(
	errPref string) (
	cleanStr string,
	lenCleanStr int) {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	lenCleanStr = 0
	cleanStr = ""

	ePrefQuark := errPrefQuark{}

	if ePrefQuark.isEmptyOrWhiteSpace(errPref) {
		return cleanStr, lenCleanStr
	}

	dirtyChars := []string{
		"\n",
		" ",
		"-",
		":",
	}

	lenDirtyChars := len(dirtyChars)

	for i := 0; i < 3; i++ {
		for j := 0; j < lenDirtyChars; j++ {

			errPref = strings.TrimLeft(strings.TrimRight(errPref, dirtyChars[j]), dirtyChars[j])

		}
	}

	if ePrefQuark.isEmptyOrWhiteSpace(errPref) {
		return cleanStr, lenCleanStr
	}

	cleanStr = errPref
	lenCleanStr = len(errPref)
	return cleanStr, lenCleanStr
}

// cleanErrorContextStr - Receives error context strings (not error
// prefix strings) and proceeds to scrub and remove invalid
// sub-strings and characters.
//
// This is a 'trim' operation which removes the target strings from
// both sides of the 'errContext' string.
//
// After this process is completed, the 'cleaned' string along with
// the length of the newly cleaned string is returned to the
// calling function.
//
func (ePrefElectron *errPrefElectron) cleanErrorContextStr(
	errContext string) (
	cleanStr string,
	lenCleanStr int) {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	lenCleanStr = 0
	cleanStr = ""

	ePrefQuark := errPrefQuark{}

	if ePrefQuark.isEmptyOrWhiteSpace(errContext) {
		return cleanStr, lenCleanStr
	}

	dirtyChars := []string{
		"\n",
		" ",
		" : ",
	}

	lenDirtyChars := len(dirtyChars)

	for i := 0; i < 3; i++ {
		for j := 0; j < lenDirtyChars; j++ {
			errContext = strings.TrimLeft(strings.TrimRight(errContext, dirtyChars[j]), dirtyChars[j])
		}
	}

	if ePrefQuark.isEmptyOrWhiteSpace(errContext) {
		return cleanStr, lenCleanStr
	}

	cleanStr = errContext
	lenCleanStr = len(errContext)
	return cleanStr, lenCleanStr
}

// getDelimiters - Returns the four delimiter strings used to
// delimit error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  delimiters                 ErrPrefixDelimiters
//     - A structure containing delimiter strings used to terminate
//       error prefix and error context lines of text.
//
//       type ErrPrefixDelimiters struct {
//         inLinePrefixDelimiter      string
//         lenInLinePrefixDelimiter   uint
//         newLinePrefixDelimiter     string
//         lenNewLinePrefixDelimiter  uint
//         inLineContextDelimiter     string
//         lenInLineContextDelimiter  uint
//         newLineContextDelimiter    string
//         lenNewLineContextDelimiter uint
//         maxErrStringLength         uint
//       }
//
func (ePrefElectron *errPrefElectron) getDelimiters() (
	delimiters ErrPrefixDelimiters) {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	delimiters.SetInLinePrefixDelimiter(" - ")

	delimiters.SetNewLinePrefixDelimiter("\n")

	delimiters.SetInLineContextDelimiter(" : ")

	delimiters.SetNewLineContextDelimiter("\n :  ")

	return delimiters
}

// ptr() - Returns a pointer to a new instance of
// errPrefElectron.
//
func (ePrefElectron errPrefElectron) ptr() *errPrefElectron {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	return &errPrefElectron{}
}
