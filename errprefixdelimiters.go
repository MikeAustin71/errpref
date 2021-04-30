package errpref

import (
	"fmt"
	"sync"
)

// ErrPrefixDelimiters - This type is used to store and transport
// the string delimiters used in parsing error prefix and error
// context strings.
//
type ErrPrefixDelimiters struct {
	inLinePrefixDelimiter      string // Error Prefix in-line string delimiters
	lenInLinePrefixDelimiter   uint   // Length of Error Prefix in-line string delimiter
	newLinePrefixDelimiter     string // Error Prefix new Line string delimiters
	lenNewLinePrefixDelimiter  uint   // Length of Error Prefix new Line string delimiters
	inLineContextDelimiter     string // Error Context in-line string delimiters
	lenInLineContextDelimiter  uint   // Length of Error Context in-line string delimiters
	newLineContextDelimiter    string // Error Context new line string delimiters
	lenNewLineContextDelimiter uint   // Length of Error Context new line string delimiters
	lock                       *sync.Mutex
}

// CopyIn - Receives an instance of type ErrPrefixDelimiters and
// proceeds to copy the internal member data variable values to the
// current ErrPrefixDelimiters instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  incomingDelimiters         *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrPrefixDelimiters instance will
//       be copied to current ErrPrefixDelimiters instance
//       ('ePrefDelims').
//
//       If this ErrPrefixDelimiters instance proves to be invalid,
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
func (ePrefDelims *ErrPrefixDelimiters) CopyIn(
	incomingDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.CopyIn() "

	ePrefDelimsElectron := errPrefixDelimitersElectron{}

	return ePrefDelimsElectron.copyIn(
		ePrefDelims,
		incomingDelimiters,
		ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// ErrPrefixDelimiters. After completion of this operation, the
// returned copy and the current ErrPrefixDelimiters instance are
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
//  ErrPrefixDelimiters
//     - If this method completes successfully, a deep copy of the
//       current ErrPrefixDelimiters instance will be returned through
//       this parameter as a completely new instance of
//       ErrPrefixDelimiters.
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
func (ePrefDelims *ErrPrefixDelimiters) CopyOut(
	ePrefix string) (
	ErrPrefixDelimiters,
	error) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.CopyOut() "

	ePrefDelimsElectron := errPrefixDelimitersElectron{}

	return ePrefDelimsElectron.copyOut(
		ePrefDelims,
		ePrefix+
			"ePrefDelims\n")
}

// Empty - This method will overwrite and reset all internal
// member variables to their zero values.
//
func (ePrefDelims *ErrPrefixDelimiters) Empty() {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	_ = errPrefixDelimitersQuark{}.ptr().empty(
		ePrefDelims,
		"")

	ePrefDelims.lock.Unlock()

	ePrefDelims.lock = nil
}

// Equal - Receives a pointer to an instance of ErrPrefixDelimiters
// and proceeds to determine whether the data values encapsulated in
// that instance are equal to those in the current
// ErrPrefixDelimiters instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  incomingDelimiters  *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. The
//       internal data values encapsulated by this object will be
//       compared to those contained in the current
//       ErrPrefixDelimiters instance. If the data values are
//       equal, a boolean flag of 'true' will be returned.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  areEqual            bool
//     - This method will compare the data values encapsulated by
//       the input parameter, 'incomingDelimiters' and the current
//       ErrPrefixDelimiters instance. If the data values are
//       equivalent, this boolean flag will be set to 'true'.
//
//       If the data values are NOT equivalent, this parameter will
//       be set to 'false'.
//
func (ePrefDelims *ErrPrefixDelimiters) Equal(
	incomingDelimiters *ErrPrefixDelimiters) (
	areEqual bool) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	areEqual,
		_ = errPrefixDelimitersElectron{}.ptr().equal(
		ePrefDelims,
		incomingDelimiters,
		"")

	return areEqual
}

// GetInLineContextDelimiter - Returns ePrefDelims.inLineContextDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetInLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.inLineContextDelimiter
}

// GetInLinePrefixDelimiter - Returns ePrefDelims.inLinePrefixDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetInLinePrefixDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.inLinePrefixDelimiter
}

// GetLengthInLineContextDelimiter - Returns the number of
// characters in the 'In Line Context Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthInLineContextDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenInLineContextDelimiter == 0 {

		ePrefDelims.lenInLineContextDelimiter =
			uint(len(ePrefDelims.inLineContextDelimiter))

	}

	return ePrefDelims.lenInLineContextDelimiter
}

// GetLengthInLinePrefixDelimiter - Returns the number of
// characters in the 'In Line Prefix Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthInLinePrefixDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenInLinePrefixDelimiter == 0 {

		ePrefDelims.lenInLinePrefixDelimiter =
			uint(len(ePrefDelims.inLinePrefixDelimiter))

	}

	return ePrefDelims.lenInLinePrefixDelimiter
}

// GetLengthNewLineContextDelimiter - Returns the number of
// characters in the 'New Line Context Delimiter' string as an
// unsigned integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthNewLineContextDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenNewLineContextDelimiter == 0 {

		ePrefDelims.lenNewLineContextDelimiter =
			uint(len(ePrefDelims.newLineContextDelimiter))

	}

	return ePrefDelims.lenNewLineContextDelimiter
}

// GetLengthNewLinePrefixDelimiter - Returns the number of
// characters in the 'New Line Prefix Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthNewLinePrefixDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenNewLinePrefixDelimiter == 0 {

		ePrefDelims.lenNewLinePrefixDelimiter =
			uint(len(ePrefDelims.newLinePrefixDelimiter))

	}

	return ePrefDelims.lenNewLinePrefixDelimiter
}

// GetNewLineContextDelimiter - Returns ePrefDelims.newLineContextDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetNewLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLineContextDelimiter
}

// GetNewLinePrefixDelimiter - Returns ePrefDelims.newLinePrefixDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetNewLinePrefixDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLinePrefixDelimiter
}

// IsValidInstance - Returns a boolean flag signalling whether the
// current ErrPrefixDelimiters instance is valid, or not.
//
// If this method returns a boolean value of 'false', it signals
// that the current ErrPrefixDelimiters instance is invalid.
//
// If this method returns a boolean value of 'true', it signals
// that the current ErrPrefixDelimiters instance is valid in all
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
//       ErrPrefixDelimiters instance is valid.
//
//       If this method returns a value of 'false', it signals that
//       the current ErrPrefixDelimiters instance is invalid.
//
//       If this method returns a value of 'true', it signals that
//       the current ErrPrefixDelimiters instance is valid in all
//       respects.
//
func (ePrefDelims *ErrPrefixDelimiters) IsValidInstance() bool {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	isValid,
		_ := errPrefixDelimitersQuark{}.ptr().
		testValidityOfErrPrefixDelimiters(
			ePrefDelims,
			"")

	return isValid
}

// IsValidInstanceError - Returns an error type signalling whether
// the current ErrPrefixDelimiters instance is valid, or not.
//
// If this method returns an error value NOT equal to 'nil', it
// signals that the current ErrPrefixDelimiters instance is
// invalid.
//
// If this method returns an error value which IS equal to 'nil',
// it signals that the current ErrPrefixDelimiters instance is
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
//       signals that the current ErrPrefixDelimiters is valid in
//       all respects.
//
//       If this returned error type is NOT equal to 'nil', it
//       signals that the current ErrPrefixDelimiters is invalid.
//
func (ePrefDelims *ErrPrefixDelimiters) IsValidInstanceError(
	ePrefix string) error {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.IsValidInstanceError() "

	_,
		err := errPrefixDelimitersQuark{}.ptr().
		testValidityOfErrPrefixDelimiters(
			ePrefDelims,
			ePrefix)

	return err
}

// New - Returns a new instance of ErrPrefixDelimiters generated
// from the string values passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newLinePrefixDelimiters    string
//     - The contents of this string will be used to parse error
//       prefix strings on separate lines of text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  inLinePrefixDelimiters     string
//     - The contents of this string will be used to separate
//       multiple error prefix elements within a single line of
//       text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  newLineContextDelimiters   string
//     - The contents of this string will be used to parse error
//       context elements on separate lines of text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  inLineContextDelimiters    string
//     - The contents of this string will be used to separate
//       multiple error context elements within a single line of
//       text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  ePrefix                    string
//     - A string containing the name of the function which called
//       this method. If an error occurs this string will be
//       prefixed to the beginning of the returned error message.
//
//       This parameter is optional. If an error prefix is not
//       required, submit an empty string for this parameter ("").
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDelimiters
//     - If this method completes successfully, this parameter will
//       return a new fully populated instance of
//       ErrPrefixDelimiters.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message.
//
//       In the event of an error, the value of parameter
//       'ePrefix' will be prefixed and attached to the beginning
//       of the error message.
//
func (ePrefDelims ErrPrefixDelimiters) New(
	newLinePrefixDelimiters string,
	inLinePrefixDelimiters string,
	newLineContextDelimiters string,
	inLineContextDelimiters string,
	ePrefix string) (
	ErrPrefixDelimiters,
	error) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += " ErrPrefixDelimiters.New()"

	newEPrefDelimiters := ErrPrefixDelimiters{
		lock: new(sync.Mutex),
	}

	err := errPrefixDelimitersMechanics{}.ptr().
		setErrPrefDelimiters(
			&newEPrefDelimiters,
			newLinePrefixDelimiters,
			inLinePrefixDelimiters,
			newLineContextDelimiters,
			inLineContextDelimiters,
			ePrefix)

	return newEPrefDelimiters, err
}

// NewDefaults - Returns a new instance of ErrPrefixDelimiters
// populated with system default values for string delimiters.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// String delimiters are used to parse raw input strings and
// separate error prefix and error context elements. In addition,
// string delimiters are used to join error prefix and error
// context elements for output or presentation text.
//
func (ePrefDelims ErrPrefixDelimiters) NewDefaults() ErrPrefixDelimiters {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	newEPrefDelimiters := ErrPrefixDelimiters{
		lock: new(sync.Mutex),
	}

	_ = errPrefixDelimitersMechanics{}.ptr().
		setToDefault(
			&newEPrefDelimiters,
			"")

	return newEPrefDelimiters
}

// SetDelimiters - Overwrites and replaces the data values for all
// internal member variables in the current ErrPrefixDelimiters
// instance.
//
// The new data values are generated from string values submitted
// as input parameters.
//
// IMPORTANT
// All existing delimiter information in this ErrPrefixDelimiters
// instance will be overwritten, deleted and replaced.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newLinePrefixDelimiters    string
//     - The contents of this string will be used to parse error
//       prefix strings on separate lines of text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  inLinePrefixDelimiters     string
//     - The contents of this string will be used to separate
//       multiple error prefix elements within a single line of
//       text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  newLineContextDelimiters   string
//     - The contents of this string will be used to parse error
//       context elements on separate lines of text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  inLineContextDelimiters    string
//     - The contents of this string will be used to separate
//       multiple error context elements within a single line of
//       text.
//
//       If an empty string (string length zero) is passed for this
//       parameter, an error will be returned.
//
//
//  ePrefix                    string
//     - A string containing the name of the function which called
//       this method. If an error occurs this string will be
//       prefixed to the beginning of the returned error message.
//
//       This parameter is optional. If an error prefix is not
//       required, submit an empty string for this parameter ("").
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
//       error Type will encapsulate an error message.
//
//       In the event of an error, the value of parameter
//       'ePrefix' will be prefixed and attached to the beginning
//       of the error message.
//
func (ePrefDelims *ErrPrefixDelimiters) SetDelimiters(
	newLinePrefixDelimiters string,
	inLinePrefixDelimiters string,
	newLineContextDelimiters string,
	inLineContextDelimiters string,
	ePrefix string) error {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += " ErrPrefixDelimiters.SetDelimiters()"

	return errPrefixDelimitersMechanics{}.ptr().
		setErrPrefDelimiters(
			ePrefDelims,
			newLinePrefixDelimiters,
			inLinePrefixDelimiters,
			newLineContextDelimiters,
			inLineContextDelimiters,
			ePrefix)
}

// SetInLineContextDelimiter - Sets ePrefDelims.inLineContextDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetInLineContextDelimiter(
	inLineContextDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.inLineContextDelimiter = inLineContextDelimiter

	ePrefDelims.lenInLineContextDelimiter =
		uint(len(ePrefDelims.inLineContextDelimiter))

	return
}

// SetInLinePrefixDelimiter - Sets ePrefDelims.inLinePrefixDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetInLinePrefixDelimiter(
	inLinePrefixDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.inLinePrefixDelimiter = inLinePrefixDelimiter

	ePrefDelims.lenInLinePrefixDelimiter =
		uint(len(ePrefDelims.inLinePrefixDelimiter))

	return
}

// SetNewLineContextDelimiter - Sets ePrefDelims.newLineContextDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetNewLineContextDelimiter(
	newLineContextDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.newLineContextDelimiter = newLineContextDelimiter

	ePrefDelims.lenNewLineContextDelimiter =
		uint(len(ePrefDelims.newLineContextDelimiter))

	return
}

// SetNewLinePrefixDelimiter - Sets ePrefDelims.newLinePrefixDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetNewLinePrefixDelimiter(
	newLinePrefixDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.newLinePrefixDelimiter = newLinePrefixDelimiter

	ePrefDelims.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelims.newLinePrefixDelimiter))

	return
}

// SetLineLengthValues - Automatically calculates, sets and stores
// the line lengths for all delimiters. These values are stored in
// internal member variables and may be accessed using 'Getter'
// methods.
//
func (ePrefDelims *ErrPrefixDelimiters) SetLineLengthValues() {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.lenInLineContextDelimiter =
		uint(len(ePrefDelims.inLineContextDelimiter))

	ePrefDelims.lenInLinePrefixDelimiter =
		uint(len(ePrefDelims.inLinePrefixDelimiter))

	ePrefDelims.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelims.newLinePrefixDelimiter))

	ePrefDelims.lenNewLineContextDelimiter =
		uint(len(ePrefDelims.newLineContextDelimiter))

	return
}

// String - Returns a string listing the values of the member
// variables for the current ErrPrefixDelimiters instance.
//
func (ePrefDelims ErrPrefixDelimiters) String() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ep := ErrPref{}

	inLinePrefixDelimiter :=
		ep.ConvertNonPrintableChars(
			[]rune(ePrefDelims.inLinePrefixDelimiter),
			true)

	newLinePrefixDelimiter :=
		ep.ConvertNonPrintableChars(
			[]rune(ePrefDelims.newLinePrefixDelimiter),
			true)

	inLineContextDelimiter :=
		ep.ConvertNonPrintableChars(
			[]rune(ePrefDelims.inLineContextDelimiter),
			true)

	newLineContextDelimiter :=
		ep.ConvertNonPrintableChars(
			[]rune(ePrefDelims.newLineContextDelimiter),
			true)

	str := fmt.Sprintf(
		"inLinePrefixDelimiter: %v\n"+
			"lenInLinePrefixDelimiter: %v\n"+
			"newLinePrefixDelimiter: %v\n"+
			"lenNewLinePrefixDelimiter: %v\n"+
			"inLineContextDelimiter: %v\n"+
			"lenInLineContextDelimiter: %v\n"+
			"newLineContextDelimiter: %v\n"+
			"lenNewLineContextDelimiter: %v\n",
		inLinePrefixDelimiter,
		ePrefDelims.lenInLinePrefixDelimiter,
		newLinePrefixDelimiter,
		ePrefDelims.lenNewLinePrefixDelimiter,
		inLineContextDelimiter,
		ePrefDelims.lenInLineContextDelimiter,
		newLineContextDelimiter,
		ePrefDelims.lenNewLineContextDelimiter)

	return str
}

// SetToDefault - Sets the values of the current
// ErrPrefixDelimiters instance to those of the system defaults.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
func (ePrefDelims *ErrPrefixDelimiters) SetToDefault() {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	_ = errPrefixDelimitersMechanics{}.ptr().
		setToDefault(
			ePrefDelims,
			"")
}

// SetToDefaultIfEmpty - Sets the values of the current
// ErrPrefixDelimiters instance to those of the system defaults if
// the current instance is empty or invalid.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
func (ePrefDelims *ErrPrefixDelimiters) SetToDefaultIfEmpty() {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	isValid,
		_ := errPrefixDelimitersQuark{}.ptr().
		testValidityOfErrPrefixDelimiters(
			ePrefDelims,
			"")

	if isValid {
		return
	}

	_ = errPrefixDelimitersMechanics{}.ptr().
		setToDefault(
			ePrefDelims,
			"")

}
