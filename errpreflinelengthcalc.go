package errpref

import (
	"fmt"
	"sync"
)

// EPrefixLineLenCalc - Error Prefix Line Length Calculator
// This type is used to perform calculations on the line length of
// error prefix text output strings. Among the calculations
// performed, these associated methods determine how many error
// prefix and error context strings can be accommodated on the same
// line of text give a maximum line length limit.
//
type EPrefixLineLenCalc struct {
	ePrefDelimiters    ErrPrefixDelimiters
	errorPrefixInfo    ErrorPrefixInfo
	currentLineStr     string
	maxErrStringLength uint
	lock               *sync.Mutex
}

// CurrLineLenExceedsMaxLineLen - If the length of the Current Line
// String (EPrefixLineLenCalc.currentLineStr) is greater than the
// Maximum Error String Length (EPrefixLineLenCalc.maxErrStringLength),
// this method returns 'true'.
//
//  currentLineStrLen > maxErrStringLength == true
//  currentLineStrLen <= maxErrStringLength == false
//
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) CurrLineLenExceedsMaxLineLen() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	if uint(len(ePrefixLineLenCalc.currentLineStr)) >
		ePrefixLineLenCalc.maxErrStringLength {
		return true
	}

	return false
}

// CopyIn - Receives an instance of type EPrefixLineLenCalc and
// proceeds to copy the internal member data variable values to the
// current EPrefixLineLenCalc instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  incomingLineLenCalc        *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this EPrefixLineLenCalc instance will
//       be copied to current EPrefixLineLenCalc instance
//       ('ePrefixLineLenCalc').
//
//       If this EPrefixLineLenCalc instance proves to be invalid,
//       an error will be returned.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) CopyIn(
	incomingLineLenCalc *EPrefixLineLenCalc,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.CopyIn() "

	ePrefLineLenCalcElectron := ePrefixLineLenCalcElectron{}

	return ePrefLineLenCalcElectron.copyIn(
		ePrefixLineLenCalc,
		incomingLineLenCalc,
		ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// EPrefixLineLenCalc. After completion of this operation, the
// returned copy and the current EPrefixLineLenCalc instance are
// identical in all respects.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  EPrefixLineLenCalc
//     - If this method completes successfully, a deep copy of the
//       current EPrefixLineLenCalc instance will be returned through
//       this parameter as a completely new instance of
//       EPrefixLineLenCalc.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be prefixed to the beginning of the returned
//       error message.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) CopyOut(
	ePrefix string) (
	EPrefixLineLenCalc,
	error) {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.CopyOut() "

	ePrefLineLenCalcElectron := ePrefixLineLenCalcElectron{}

	return ePrefLineLenCalcElectron.copyOut(
		ePrefixLineLenCalc,
		ePrefix)
}

// Empty - Sets all internal variables to their zero or
// uninitialized values.
//
// IMPORTANT
// This method will DELETE ALL VALID DATA contained in this
// instance of EPrefixLineLenCalc.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) Empty() {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefixLineLenCalc.ePrefDelimiters = ErrPrefixDelimiters{}

	ePrefixLineLenCalc.errorPrefixInfo.Empty()

	ePrefixLineLenCalc.currentLineStr = ""
	ePrefixLineLenCalc.maxErrStringLength = 0

}

// Equal - Returns a boolean flag signaling whether the data values
// contained in the current EPrefixLineLenCalc instance are equal
// to those contained in input parameter, 'lineLenCalcTwo'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  lineLenCalcTwo      *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. The data
//       values contained in this instance will be compared to
//       those contained in the current EPrefixLineLenCalc instance
//       (ePrefixLineLenCalc) to determine equality.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - A boolean flag signaling whether the data values contained
//       in the current EPrefixLineLenCalc instance are equal to
//       those contained in input parameter 'lineLenCalcTwo'. If
//       the data values are equal in all respects, this returned
//       boolean value will be set to 'true'.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) Equal(
	lineLenCalcTwo *EPrefixLineLenCalc) bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	areEqual,
		_ := ePrefixLineLenCalcElectron{}.ptr().equal(
		ePrefixLineLenCalc,
		lineLenCalcTwo,
		"")

	return areEqual
}

// EPrefWithoutContextExceedsRemainLineLen - Returns 'true' if the
// length of the in-line of in-line error prefix delimiter plus the
// length of the error prefix string exceeds the remaining unused
// line length.
//
//   in-line error prefix delimiter +
//   error prefix   > Remaining line Length
//    This method returns 'true'.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) EPrefWithoutContextExceedsRemainLineLen() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var errPrefixLen uint

	errPrefixLen =
		ePrefixLineLenCalc.ePrefDelimiters.GetLengthInLinePrefixDelimiter() +
			ePrefixLineLenCalc.errorPrefixInfo.GetLengthErrPrefixStr()

	currLineLen :=
		uint(len(ePrefixLineLenCalc.currentLineStr))

	if currLineLen > ePrefixLineLenCalc.maxErrStringLength {
		return true
	}

	remainingLineLen :=
		ePrefixLineLenCalc.maxErrStringLength -
			currLineLen

	if errPrefixLen > remainingLineLen {
		return true
	}

	return false
}

// EPrefixWithContextExceedsRemainLineLen - Returns 'true' if the
// combination of in-line error prefix delimiter plus error prefix
// plus in-line error context delimiter plus error context string
// exceeds the remaining unused line length.
//
//   in-line error prefix delimiter +
//   error prefix +
//   in-line error context delimiter +
//   error context                     > Remaining line Length
//    This method returns 'true'.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) EPrefixWithContextExceedsRemainLineLen() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var prefixWithContextLen uint

	prefixWithContextLen =
		ePrefixLineLenCalc.ePrefDelimiters.GetLengthInLinePrefixDelimiter() +
			ePrefixLineLenCalc.errorPrefixInfo.GetLengthErrPrefixStr()

	if ePrefixLineLenCalc.errorPrefixInfo.GetErrPrefixHasContextStr() {

		prefixWithContextLen +=
			ePrefixLineLenCalc.ePrefDelimiters.
				GetLengthInLineContextDelimiter() +
				ePrefixLineLenCalc.errorPrefixInfo.
					GetLengthErrContextStr()
	}

	currLineLen :=
		uint(len(ePrefixLineLenCalc.currentLineStr))

	if currLineLen > ePrefixLineLenCalc.maxErrStringLength {
		return true
	}

	remainingLineLen :=
		ePrefixLineLenCalc.maxErrStringLength -
			currLineLen

	if prefixWithContextLen > remainingLineLen {
		return true
	}

	return false
}

// ErrPrefixHasContext - Returns a boolean flag signaling whether
// the error prefix has an associated error context.
//
// If this method returns 'true', it means that the encapsulated
// error prefix DOES HAVE an associated error context string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) ErrPrefixHasContext() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.
		errorPrefixInfo.
		GetErrPrefixHasContextStr()
}

// ErrorContextIsEmpty - Returns a boolean flag signalling whether
// the Error Context String is an empty string.
//
// If this method returns 'true', it means that the Error Context
// String is empty and has a zero string length.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) ErrorContextIsEmpty() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	if ePrefixLineLenCalc.
		errorPrefixInfo.GetLengthErrContextStr() == 0 {
		return true
	}

	return false
}

// ErrorPrefixIsEmpty - Returns a boolean flag signalling whether
// the Error Prefix String is an empty string.
//
// If this method returns 'true', it means that the Error Prefix
// String is empty and has a zero string length.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) ErrorPrefixIsEmpty() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	if ePrefixLineLenCalc.
		errorPrefixInfo.GetLengthErrPrefixStr() == 0 {
		return true
	}

	return false
}

// GetErrorContextStr - Returns the error context string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetErrorContextStr() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.errorPrefixInfo.GetErrContextStr()
}

// GetCurrLineStr - Return the current line string. This string
// includes the characters which have been formatted and included
// in a single text line but which have not yet been written out
// text display. As soon as the current line string fills up with
// characters to the maximum allowed line length, the text line
// will be written out to the display device.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetCurrLineStr() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.currentLineStr
}

// GetCurrLineStrLength - Returns an unsigned integer value
// representing the number of characters in or string length of
// the Current Line String.
//
// The Current Line String contains the characters which have been
// collected thus far from error prefix and error context
// information. The current line string is used to control maximum
// line length and stores the characters which have not yet been
// written out to the text display.
//
// As soon as the Current Line String fills up with characters to
// the maximum allowed line length, the text line will be written
// out to the display device.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetCurrLineStrLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var lenCurrentLineStr uint

	lenCurrentLineStr =
		uint(len(ePrefixLineLenCalc.currentLineStr))

	return lenCurrentLineStr
}

// GetDelimiterInLineErrContext - Returns the In-Line Error Context
// delimiter as a string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetDelimiterInLineErrContext() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.
		ePrefDelimiters.
		GetInLineContextDelimiter()
}

// GetDelimiterNewLineErrContext - Returns the New Line Error
// Context delimiter as a string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetDelimiterNewLineErrContext() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.
		ePrefDelimiters.
		GetNewLineContextDelimiter()
}

// GetDelimiterInLineErrPrefix - Returns the In-Line Error Prefix
// delimiter as a string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetDelimiterInLineErrPrefix() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.
		ePrefDelimiters.
		GetInLinePrefixDelimiter()
}

// GetDelimiterNewLineErrPrefix - Returns the New Line Error Prefix
// delimiter as a string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetDelimiterNewLineErrPrefix() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.
		ePrefDelimiters.
		GetNewLinePrefixDelimiter()
}

// GetErrorPrefixStr - Returns the error prefix string.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetErrorPrefixStr() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.errorPrefixInfo.GetErrPrefixStr()
}

// GetMaxErrStringLength - Returns the current the value for
// maximum error string length. This limit controls the line length
// for text displays of error prefix strings.
//
// The value of maximum error string length is returned as an
// unsigned integer.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetMaxErrStringLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.maxErrStringLength
}

// GetRemainingLineLength - Returns the remaining line length. This is
// the difference between current line length and Maximum Error
// String Length.
//
//   remainingLineLen = maxErrStringLen - currentLineStringLen
//
// If currentLineStringLen is greater than Maximum Error String
// Length, the Remaining String Length is zero.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetRemainingLineLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var (
		lenCurrentLineStr,
		remainingLineLength uint
	)

	lenCurrentLineStr =
		uint(len(ePrefixLineLenCalc.currentLineStr))

	if lenCurrentLineStr >
		ePrefixLineLenCalc.maxErrStringLength {

		remainingLineLength = 0

	} else {

		remainingLineLength =
			ePrefixLineLenCalc.maxErrStringLength -
				lenCurrentLineStr

	}

	return remainingLineLength
}

// IsErrPrefixLastIndex - Returns a boolean flag signalling whether
// the encapsulated ErrorPrefixInfo object is the last element in
// an array.
//
// If this method returns 'true', it means that the encapsulated
// ErrorPrefixInfo object is the last element in an array of
// ErrorPrefixInfo objects.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) IsErrPrefixLastIndex() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.errorPrefixInfo.GetIsLastIndex()
}

// IsValidInstance - Returns a boolean flag signalling whether the
// current EPrefixLineLenCalc instance is valid, or not.
//
// If this method returns a boolean value of 'false', it signals
// that the current EPrefixLineLenCalc instance is invalid.
//
// If this method returns a boolean value of 'true', it signals
// that the current EPrefixLineLenCalc instance is valid in all
// respects.
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
//  bool
//     - This boolean flag signals whether the current
//       EPrefixLineLenCalc instance is valid.
//
//       If this method returns a value of 'false', it signals that
//       the current EPrefixLineLenCalc instance is invalid.
//
//       If this method returns a value of 'true', it signals that
//       the current EPrefixLineLenCalc instance is valid in all
//       respects.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) IsValidInstance() bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	isValid,
		_ := ePrefixLineLenCalcQuark{}.ptr().
		testValidityOfEPrefixLineLenCalc(
			ePrefixLineLenCalc,
			"")

	return isValid
}

// IsValidInstanceError - Returns an error type signalling whether
// the current EPrefixLineLenCalc instance is valid, or not.
//
// If this method returns an error value NOT equal to 'nil', it
// signals that the current EPrefixLineLenCalc instance is
// invalid.
//
// If this method returns an error value which IS equal to 'nil',
// it signals that the current EPrefixLineLenCalc instance is
// valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
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
//  error
//     - If this returned error type is set equal to 'nil', it
//       signals that the current EPrefixLineLenCalc is valid in
//       all respects.
//
//       If this returned error type is NOT equal to 'nil', it
//       signals that the current EPrefixLineLenCalc is invalid.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) IsValidInstanceError(
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.IsValidInstanceError() "

	_,
		err := ePrefixLineLenCalcQuark{}.ptr().
		testValidityOfEPrefixLineLenCalc(
			ePrefixLineLenCalc,
			ePrefix+
				"ePrefixLineLenCalc\n")

	return err
}

// New - Returns a new, empty instance of type EPrefixLineLenCalc.
//
func (ePrefixLineLenCalc EPrefixLineLenCalc) New() EPrefixLineLenCalc {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	newEPrefixLineLenCalc := EPrefixLineLenCalc{}

	return newEPrefixLineLenCalc
}

// Ptr - Returns a pointer to a new, empty instance of type
// EPrefixLineLenCalc.
//
func (ePrefixLineLenCalc EPrefixLineLenCalc) Ptr() *EPrefixLineLenCalc {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	newEPrefixLineLenCalc := EPrefixLineLenCalc{}

	return &newEPrefixLineLenCalc
}

// SetCurrentLineStr - Sets the Current Line String. This string
// represents the characters which have been collected thus far
// from error prefix and error context information. The current
// line string is used to control maximum line length and stores
// the characters which have not yet been written out to the
// text display.
//
// Be sure to set the Maximum Error String Length. Both the Current
// Line String and the  Maximum Error String Length are essential
// to line length calculations.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetCurrentLineStr(
	currentLineStr string) {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefixLineLenCalc.currentLineStr = currentLineStr

	return
}

// SetEPrefDelimiters - Sets the Error Prefix Delimiters member
// variable.
//
// The Error Prefix Delimiters object stores string delimiters used
// to terminate error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefDelimiters     ErrPrefixDelimiters
//     - An Error Prefix Delimiters object. This delimiters object
//       contains information on the delimiter strings used to
//       terminate error prefix and error context strings.
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
//
//  ePrefix             string
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message. Note that this error message
//       will incorporate the method chain and text passed by input
//       parameter, 'ePrefix'. The 'ePrefix' text will be prefixed
//       to the beginning of the error message.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetEPrefDelimiters(
	ePrefDelimiters ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.SetEPrefDelimiters() "

	err := ePrefDelimiters.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefDelimiters' validity check\n"+
			"returned an error message!\n"+
			"%v\n",
			ePrefix,
			err.Error())
	}

	return ePrefixLineLenCalc.ePrefDelimiters.CopyIn(
		&ePrefDelimiters,
		ePrefix+
			"ePrefDelimiters\n")

}

// SetErrPrefixInfo - Sets the Error Prefix Information Object member
// variable.
//
// The Error Prefix Information Object stores information on the
// error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixInfo       *ErrorPrefixInfo
//     - This Error Prefix Information Object stores information on
//       the error prefix and error context strings.
//
//       type ErrorPrefixInfo struct {
//         isValid                bool
//         isLastIdx              bool // Signals the last index in the array
//         errorPrefixStr         string
//         lenErrorPrefixStr      uint
//         errPrefixHasContextStr bool
//         errorContextStr        string
//         lenErrorContextStr     uint
//       }
//
//
//  ePrefix             string
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message. Note that this error message
//       will incorporate the method chain and text passed by input
//       parameter, 'ePrefix'. The 'ePrefix' text will be prefixed
//       to the beginning of the error message.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetErrPrefixInfo(
	errPrefixInfo *ErrorPrefixInfo,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.SetErrPrefixInfo() "

	if errPrefixInfo == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefixInfo' is a "+
			"'nil' pointer\n",
			ePrefix)
	}

	err := errPrefixInfo.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefixInfo' validity check\n"+
			"returned an error message!\n"+
			"%v\n",
			ePrefix,
			err.Error())
	}

	err = ePrefixLineLenCalc.errorPrefixInfo.CopyIn(
		errPrefixInfo,
		ePrefix)

	return err
}

// SetMaxErrStringLength - Sets EPrefixLineLenCalc.maxErrStringLength
//
// This method sets the value for maximum error string length. This
// limit controls the line length for text displays of error prefix
// strings.
//
// Set this value first, before setting Current Line Length
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetMaxErrStringLength(
	maxErrStringLength uint) {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefixLineLenCalc.maxErrStringLength = maxErrStringLength

}
