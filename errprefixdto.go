package errpref

import (
	"strings"
	"sync"
)

// ErrPrefixDto - Error Prefix Data Transfer Object. This type
// encapsulates and formats error prefix and error context strings.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to
// document the thread of code execution by listing the calling
// sequence for specific functions and methods.
//
// The error context string is designed to provide additional
// error context information associated with the currently
// executing function or method. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// ErrPrefixDto is similar in function to type 'ErrPref'. However,
// unlike 'ErrPref', ErrPrefixDto is configured encapsulate and
// transmit error prefix information between methods. 'ErrPref'
// only returns strings. This type, ErrPrefixDto, returns an object
// which may be exchanged between methods.
//
type ErrPrefixDto struct {
	ePrefCol                        []ErrorPrefixInfo
	leftMarginLength                int
	leftMarginChar                  rune
	isLastLineTerminatedWithNewLine bool
	turnOffTextDisplay              bool
	inputStrDelimiters              ErrPrefixDelimiters
	outputStrDelimiters             ErrPrefixDelimiters
	maxErrPrefixTextLineLength      uint
	lock                            *sync.Mutex
}

// AddEPrefStrings - Adds error prefix information extracted
// from a passed two-dimensional string array to the internal
// ErrorPrefixInfo collection managed by this ErrPrefixDto
// instance.
//
// The incoming collection of error prefix information is added to
// the END of the collection maintained by the current ErrPrefixDto
// instance.
//
// If input parameter 'twoDStrArray' is 'nil' or empty, this method
// will take no action an exit.
//
func (ePrefDto *ErrPrefixDto) AddEPrefStrings(
	twoDStrArray [][2]string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	_ = errPrefixDtoAtom{}.ptr().addTwoDimensionalStringArray(
		ePrefDto,
		twoDStrArray,
		"")
}

// AddEPrefCollectionStr - Receives a string containing one or more
// error prefix and error context pairs. This error prefix
// information is parsed and added to the internal store of
// error prefix elements.
//
// The incoming collection of error prefix information is added to
// the END of the collection maintained by the current ErrPrefixDto
// instance.
//
// Upon completion this method returns an integer value identifying
// the number of error prefix elements successfully parsed and
// added to internal error prefix storage.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefixCollectionStr      string
//     - A string consisting of a series of error prefix strings.
//       Error prefixes should be delimited by either a new line
//       character ('\n') or the in-line delimiter string, " - ".
//
//       If the collection string contains error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  numberOfCollectionItemsParsed int
//     - This returned integer value contains the number of error
//       prefix elements parsed from input parameter
//       'errorPrefixCollection' and added to the internal store
//       of error prefix elements.
//
func (ePrefDto *ErrPrefixDto) AddEPrefCollectionStr(
	errorPrefixCollectionStr string) (
	numberOfCollectionItemsParsed int) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	previousCollectionLen := len(ePrefDto.ePrefCol)

	ePrefixAtom := errPrefixDtoAtom{}

	ePrefixAtom.getEPrefContextArray(
		errorPrefixCollectionStr,
		ePrefDto.inputStrDelimiters,
		&ePrefDto.ePrefCol)

	ePrefixAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	numberOfCollectionItemsParsed =
		len(ePrefDto.ePrefCol) -
			previousCollectionLen

	return numberOfCollectionItemsParsed
}

// Copy - Creates a deep copy of the data fields contained in
// the current ErrPrefixDto instance, and returns that data as a
// as a new instance of ErrPrefixDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - If this method completes successfully, a deep copy of the
//       current ErrPrefixDto instance will be returned through
//       this parameter as a pointer to a new instance of
//       ErrPrefixDto.
//
func (ePrefDto *ErrPrefixDto) Copy() ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}

// CopyPtr - Creates a deep copy of the data fields contained in
// the current ErrPrefixDto instance, and returns that data as a
// pointer to a new instance of ErrPrefixDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - If this method completes successfully, a deep copy of the
//       current ErrPrefixDto instance will be returned through
//       this parameter as a pointer to a new instance of
//       ErrPrefixDto.
//
func (ePrefDto *ErrPrefixDto) CopyPtr() *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return &newErrPrefixDto
}

// CopyIn - Copies the data fields from an incoming instance of
// ErrPrefixDto ('inComingErrPrefixDto') to the data fields of
// the current ErrPrefixDto instance ('ePrefDto').
//
// All of the data fields in current ErrPrefixDto instance
// ('ePrefDto') will be modified and overwritten.
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
//  inComingErrPrefixDto       *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrPrefixDto instance will be
//       copied to current ErrPrefixDto instance
//       ('ePrefDto').
//
//
//  eMsg                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'eMsg'. This parameter is optional.
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
//       chain and text passed by input parameter, 'eMsg'.
//
func (ePrefDto *ErrPrefixDto) CopyIn(
	inComingErrPrefixDto *ErrPrefixDto,
	eMsg string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	eMsg += "ErrPrefixDto.CopyIn()\n"

	return errPrefixDtoAtom{}.ptr().
		copyInErrPrefDto(
			ePrefDto,
			inComingErrPrefixDto,
			eMsg)
}

// CopyInFromIBuilder - Receives an object implementing the
// IBuilderErrorPrefix interface and proceeds copy its error prefix
// and error context information into the current instance of
// ErrPrefixDto.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
// This means that if IBuilderErrorPrefix parameter,
// 'inComingIBuilder' generates an empty set of error prefix
// error prefix information, the current ErrPrefixDto instance will
// likewise be configured with an identical empty set of error
// prefix information.
//
func (ePrefDto *ErrPrefixDto) CopyInFromIBuilder(
	inComingIBuilder IBuilderErrorPrefix,
	eMsg string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	eMsg += "ErrPrefixDto.CopyInFromIBuilder()\n"

	err := errPrefixDtoNanobot{}.ptr().
		setFromIBuilder(
			ePrefDto,
			inComingIBuilder,
			eMsg)

	return err
}

// CopyOutToIBuilder - Receives an object implementing the
// IBuilderErrorPrefix Interface and populates that object with the
// error prefix and error context information contained in the
// current ErrPrefixDto instance.
//
// This method will transfer error prefix and context information
// OUT to object implementing the IBuilderErrorPrefix interface.
//
func (ePrefDto *ErrPrefixDto) CopyOutToIBuilder(
	inComingIBuilder IBuilderErrorPrefix) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newTwoDSlice,
		_ := errPrefixDtoMechanics{}.ptr().
		getEPrefStrings(
			ePrefDto,
			"")

	inComingIBuilder.SetEPrefStrings(
		newTwoDSlice)
}

// CopyOut - Creates a deep copy of the data fields contained in
// the current ErrPrefixDto instance, and returns that data as a
// new instance of ErrPrefixDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  ErrPrefixDto
//     - If this method completes successfully, a deep copy of the
//       current ErrPrefixDto instance will be returned through
//       this parameter as a completely new instance of
//       ErrPrefixDto.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'eMsg'. The
//       'eMsg' text will be prefixed to the beginning of the returned
//       error message.
//
func (ePrefDto *ErrPrefixDto) CopyOut(
	eMsg string) (
	ErrPrefixDto,
	error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	eMsg += "ErrPrefixDto.CopyOut() "

	return errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			eMsg)
}

// Empty - Reinitializes all internal member variables for the
// current ErrPrefixDto instance to their zero values. This method
// will effectively delete all error prefix information contained
// in the current ErrPrefixDto instance.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
func (ePrefDto *ErrPrefixDto) Empty() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	_ = errPrefixDtoElectron{}.ptr().
		emptyErrorPrefixDto(
			ePrefDto,
			"")

	ePrefDto.lock.Unlock()

	ePrefDto.lock = nil
}

// EmptyEPrefCollection - All instances of ErrPrefixDto store
// error prefix information internally in an array of
// ErrorPrefixInfo objects. This method will delete or empty the
// ErrorPrefixInfo array for the current ErrPrefixDto instance and
// initialize it to nil.
//
//
// IMPORTANT
// Effectively, this will delete all existing error prefix
// information stored in the current ErrPrefixDto instance.
//
func (ePrefDto *ErrPrefixDto) EmptyEPrefCollection() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	_ = errPrefixDtoQuark{}.ptr().emptyErrPrefInfoCollection(
		ePrefDto,
		"")
}

// Equal - Returns a boolean flag signaling whether the data values
// contained in the current ErrPrefixDto instance are equal to
// those contained in input parameter, 'ePrefixDto2'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixDto2         *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. The data values
//       contained in this instance will be compared to those
//       contained in the current ErrPrefixDto instance (ePrefDto)
//       to determine equality.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - A boolean flag signaling whether the data values contained
//       in the current ErrPrefixDto instance are equal to those
//       contained in input parameter 'ePrefixDto2'. If the data
//       values are equal in all respects, this returned boolean
//       value will be set to 'true'.
//
func (ePrefDto *ErrPrefixDto) Equal(
	ePrefixDto2 *ErrPrefixDto) bool {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return errPrefixDtoAtom{}.ptr().areEqualErrPrefDtos(
		ePrefDto,
		ePrefixDto2)
}

// GetInputStringDelimiters - Returns the input string delimiters
// configured for the current instance of ErrPrefixDto.
//
// Input string delimiters are applied the current ErrPrefixDto
// instance receives and parses raw strings containing error prefix
// information. Such strings are parsed before extracting error
// prefix information for internal storage in the error prefix
// collection maintained by the current ErrPrefixDto instance.
//
// If input string delimiters were not directly configured by the
// user, the default input string delimiters will be applied.
//
// The default input string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// To set and control input string delimiters use method:
//    ErrPrefixDto.SetInputStringDelimiters()
//
func (ePrefDto *ErrPrefixDto) GetInputStringDelimiters() ErrPrefixDelimiters {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	inputDelims,
		_ :=
		ePrefDto.inputStrDelimiters.CopyOut("")

	return inputDelims
}

// GetIsLastLineTerminatedWithNewLine - Returns the boolean flag
// which determines whether the last line of error prefix strings
// returned by this ErrPrefixDto instance will be terminated with
// a new line character ('\n').
//
// By default, the last line of error prefix strings returned by
// the method 'ErrPrefixDto.String()' DO NOT end with a new line
// character ('\n'). In other words, the last line of returned
// error prefix strings ARE NOT, by default, terminated with new
// line characters ('\n').
//
// If terminating the last line of error prefix strings with a new
// line character ('\n') is a preferred option, use the method
// 'SetIsLastLineTermWithNewLine()' to configure this
// feature.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If this return parameter is set to 'true', it signals that
//       the last line of all error prefix strings returned by this
//       ErrPrefixDto instance WILL BE terminated with a new line
//       character ('\n').
//
func (ePrefDto *ErrPrefixDto) GetIsLastLineTerminatedWithNewLine() bool {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return ePrefDto.isLastLineTerminatedWithNewLine
}

// GetLeftMarginChar - Returns a rune or text character which will
// be used to populate the left margin of formatted error prefix
// strings returned by method ErrPrefixDto.String().
//
// The default for the Left Margin Character is the space character
// (' '). This means that a value of zero for the Left Margin
// character will cause a space character to be used in populating
// the left margin for error prefix strings.
//
// To set the Left Margin Character see method:
// ErrPrefixDto.SetLeftMarginChar().
//
// Finally, remember that the Left Margin will never be applied to
// error prefix strings unless the the Left Margin Length parameter
// is greater than zero. To set the Left Margin Length, see method:
// ErrPrefixDto.SetLeftMarginLength().
//
func (ePrefDto *ErrPrefixDto) GetLeftMarginChar() rune {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return ePrefDto.leftMarginChar
}

// GetLeftMarginLength - Returns an integer value defining the
// length of the left margin which will be applied to formatted
// error prefix strings returned by method ErrPrefixDto.String().
//
// If the length of the Left Margin is set to zero, no left margin
// will applied to error prefix strings  returned by method
//ErrPrefixDto.String().
//
func (ePrefDto *ErrPrefixDto) GetLeftMarginLength() int {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return ePrefDto.leftMarginLength
}

// GetEPrefCollection - Returns a deep copy of the current
// error prefix collection maintained by this ErrPrefixDto
// instance.
//
func (ePrefDto *ErrPrefixDto) GetEPrefCollection() []ErrorPrefixInfo {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	var newErrorPrefixCollection []ErrorPrefixInfo

	lenEPrefCol := len(ePrefDto.ePrefCol)

	if lenEPrefCol == 0 {
		newErrorPrefixCollection = make([]ErrorPrefixInfo, 0)
		return newErrorPrefixCollection
	}

	newErrorPrefixCollection =
		make([]ErrorPrefixInfo,
			lenEPrefCol)

	_ = copy(newErrorPrefixCollection, ePrefDto.ePrefCol)

	return newErrorPrefixCollection
}

// GetEPrefCollectionLen - Returns the number of elements in the
// error prefix array. Effectively this is a count of the number of
// error prefix elements maintained by this ErrPrefixDto instance.
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
//  int
//     - This integer value represents the number of separate error
//       prefix elements maintained by this ErrPrefixDto instance.
//
//
func (ePrefDto *ErrPrefixDto) GetEPrefCollectionLen() int {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol =
			make([]ErrorPrefixInfo, 0)
	}

	return len(ePrefDto.ePrefCol)
}

// GetEPrefStrings - Returns a two dimensional slice of Error
// Prefix and Context strings.
//
// The Error Prefix is always in the [x][0] position. The Error
// Context string is always in the [x][1] position. The Error
// Context string may be an empty string.
//
func (ePrefDto *ErrPrefixDto) GetEPrefStrings() [][2]string {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newTwoDSlice,
		_ := errPrefixDtoMechanics{}.ptr().
		getEPrefStrings(
			ePrefDto,
			"")

	return newTwoDSlice
}

// GetMaxTextLineLen - Returns the maximum limit on the
// number of characters allowed in an error prefix text line output
// for display purposes.
//
// This unsigned integer value controls the maximum character limit
// for this specific ErrPrefixDto instance. This maximum limit will
// remain in force for the life of this ErrPrefixDto instance or
// until a call is made to the method 'SetMaxTextLineLen()'
// which is used to change the value.
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
//  uint
//     - This unsigned integer specifies the maximum limit on the
//       number of characters allowed in an error prefix text line
//       output for display purposes.
//
//
func (ePrefDto *ErrPrefixDto) GetMaxTextLineLen() uint {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	return ePrefDto.maxErrPrefixTextLineLength
}

// GetMaxTextLineLenDefault - Returns the system default value for
// Maximum Text Line Length Limit.
//
// The Maximum Text Line Length Limit is the number of characters
// allowed in an error prefix text line output for display
// purposes.
//
// This system default value is provided for information purposes
// only. The system default value may or may not be applied to the
// current ErrPrefixDto instance. To learn the Maximum Text Line
// Length Limit applied to the current ErrPrefixDto instance, call
// method: ErrPrefixDto.GetMaxTextLineLen()
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
//  uint
//     - This unsigned integer specifies the system default value
//       for the Maximum Text Line Length Limit. This value is
//       provided for informational purposes only. The system
//       default value may or may not be applied to the current
//       ErrPrefixDto instance. To learn the Maximum Text Line
//       Length Limit applied to the current ErrPrefixDto instance,
//       call method:  ErrPrefixDto.GetMaxTextLineLen()
//
func (ePrefDto *ErrPrefixDto) GetMaxTextLineLenDefault() uint {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return errPrefQuark{}.ptr().
		getMasterErrPrefDisplayLineLength()
}

// GetOutputStringDelimiters - Returns the output string delimiters
// configured for the current instance of ErrPrefixDto.
//
// The output string delimiters are used to format error prefix and
// context information returned as strings by methods:
//   ErrPrefixDto.String()
//   ErrPrefixDto.StrMaxLineLen()
//
// Error prefix and context elements in these formatted, returned
// strings are separated by these output string delimiters.
//
// If output string delimiters were not directly configured by the
// user, the default output string delimiters will be applied.
//
// The default output string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
//
// To set and control output string delimiters use method:
//    ErrPrefixDto.SetOutputStringDelimiters()
//
func (ePrefDto *ErrPrefixDto) GetOutputStringDelimiters() ErrPrefixDelimiters {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	outputDelims,
		_ :=
		ePrefDto.outputStrDelimiters.CopyOut("")

	return outputDelims
}

// GetStrDelimiters - Returns two ErrPrefixDelimiters objects which
// containing the input string delimiters and the output string
// delimiters configured for the current instance of ErrPrefixDto.
//
// Input string delimiters are employed by the current ErrPrefixDto
// instance when receiving and parsing raw strings containing error
// prefix information. Such strings are parsed before extracting
// error prefix information to be stored in the error prefix
// collection maintained by the current ErrPrefixDto instance.
//
// Output string delimiters on the other hand are used by the
// current ErrPrefixDto instance to format error prefix strings
// for output. Formatted strings returned for output are generated
// by the following two methods:
//    ErrPrefixDto.String()
//    ErrPrefixDto.StrMaxLineLen()
//
//
// If the input and output string delimiters were not directly
// configured by the user, the system default string delimiters
// will be applied. The system default string delimiters are
// listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
func (ePrefDto *ErrPrefixDto) GetStrDelimiters() (
	inputStrDelimiters ErrPrefixDelimiters,
	outputStrDelimiters ErrPrefixDelimiters) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	inputStrDelimiters,
		_ = ePrefDto.inputStrDelimiters.
		CopyOut("")

	outputStrDelimiters,
		_ = ePrefDto.outputStrDelimiters.
		CopyOut("")

	return inputStrDelimiters, outputStrDelimiters
}

// GetTurnOffTextDisplay - Returns the current value of the
// "Turn Off Text Display" flag represented by internal member
// variable:
//    ErrPrefixDto.turnOffTextDisplay
//
// If this flag ('turnOffTextDisplay') is set to 'true', this
// instance of ErrPrefixDto WILL NOT format and return error prefix
// and context information through method, ErrPrefixDto.String().
// In this case, the method ErrPrefixDto.String() will instead
// return an empty string.
//
// Conversely, if this flag ('turnOffTextDisplay') is set to
// 'false' formatted error prefix and context information WILL BE
// formatted and returned through method ErrPrefixDto.String().
//
// To control the status of the "Turn Off Text Display" flag, see
// method ErrPrefixDto.SetTurnOffTextDisplay().
//
func (ePrefDto *ErrPrefixDto) GetTurnOffTextDisplay() bool {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return ePrefDto.turnOffTextDisplay
}

// IsValidInstance - Returns a boolean flag signalling whether the
// current ErrPrefixDto instance is valid, or not.
//
// If this method returns a boolean value of 'false', it signals
// that the current ErrPrefixDto instance is invalid.
//
// If this method returns a boolean value of 'true', it signals
// that the current ErrPrefixDto instance is valid in all
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
//       ErrPrefixDto instance is valid.
//
//       If this method returns a value of 'false', it signals that
//       the current ErrPrefixDto instance is invalid.
//
//       If this method returns a value of 'true', it signals that
//       the current ErrPrefixDto instance is valid in all
//       respects.
//
func (ePrefDto *ErrPrefixDto) IsValidInstance() bool {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	isValid,
		_ := errPrefixDtoQuark{}.ptr().
		testValidityOfErrPrefixDto(
			ePrefDto,
			"")

	return isValid
}

// IsValidInstanceError - Returns an error type signalling whether
// the current ErrPrefixDto instance is valid, or not.
//
// If this method returns an error value NOT equal to 'nil', it
// signals that the current ErrPrefixDto instance is invalid.
//
// If this method returns an error value which IS equal to 'nil',
// it signals that the current ErrPrefixDto instance is valid in
// all respects.
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
//       signals that the current ErrPrefixDto is valid in all
//       respects.
//
//       If this returned error type is NOT equal to 'nil', it
//       signals that the current ErrPrefixDto is invalid.
//
func (ePrefDto *ErrPrefixDto) IsValidInstanceError(
	ePrefix string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += "ErrPrefixDto.IsValidInstanceError()\n"

	_,
		err := errPrefixDtoQuark{}.ptr().
		testValidityOfErrPrefixDto(
			ePrefDto,
			"")

	return err
}

// MergeErrPrefixDto - Receives a pointer to another ErrPrefixDto
// object and proceeds to add the error prefix collection to that
// of the current ErrPrefixDto instance.
//
// This means that the error prefix collection for
// 'incomingErrPrefixDto' will be added to the end of the
// collection for the current ErrPrefixDto instance.
//
func (ePrefDto *ErrPrefixDto) MergeErrPrefixDto(
	incomingErrPrefixDto *ErrPrefixDto) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if incomingErrPrefixDto == nil {
		return
	}

	lenIncomingDto :=
		len(incomingErrPrefixDto.ePrefCol)

	if lenIncomingDto == 0 {
		return
	}

	ePrefDto.ePrefCol =
		append(ePrefDto.ePrefCol,
			incomingErrPrefixDto.ePrefCol...)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)
}

// New - Returns a new and properly initialized instance of
// ErrPrefixDto. String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a new, properly initialized
//       instance of ErrPrefixDto. The ErrorPrefixInfo collection
//       encapsulated by the new instance will be an empty
//       collection containing zero elements.
//
//
func (ePrefDto ErrPrefixDto) New() ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	return newErrPrefixDto
}

// NewEPrefCtx - Returns a new and properly initialized instance of
// ErrPrefixDto. The returned ErrPrefixDto instance will be
// configured with error prefix information extracted from input
// parameter, 'newErrPrefix'. Optional error context information
// is extracted from input parameter, 'newErrContext'.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix string typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This method is designed to process a single new error prefix
//       string. To process a collection of error prefix strings, see
//       method 'ErrPrefixDto.NewEPrefOld()'.
//
//
//  newErrContext       string
//     - This is the error context information associated with the
//       new error prefix string ('newErrPrefix'). This parameter
//       is optional and will accept an empty string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a new, properly initialized
//       instance of ErrPrefixDto containing error prefix
//       information extracted from input parameters,
//       'newErrPrefix' and 'newErrContext'.
//
func (ePrefDto ErrPrefixDto) NewEPrefCtx(
	newErrPrefix string,
	newErrContext string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&newErrPrefixDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		newErrPrefixDto.ePrefCol)

	return newErrPrefixDto
}

// NewEPrefOld - Returns a new and properly initialized instance of
// ErrPrefixDto. The returned ErrPrefixDto instance will be
// configured with error prefix information extracted from input
// parameter, 'oldErrPrefix'.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// This method is designed to process a series of error prefix
// strings passed through input parameter, 'oldErrPrefix'. If
// only one error prefix string is available, consider using
// method 'ErrPrefixDto.NewEPrefCtx()'.
//
// String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string which will encapsulate one or more error prefix
//       elements. This string will be parsed into error prefix and
//       error context components and stored in the returned
//       ErrPrefixDto instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a new, properly initialized
//       instance of ErrPrefixDto containing error prefix
//       information extracted from input parameter,
//       'oldErrPrefix'.
//
//
func (ePrefDto ErrPrefixDto) NewEPrefOld(
	oldErrPrefix string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		newErrPrefixDto.inputStrDelimiters,
		&newErrPrefixDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		newErrPrefixDto.ePrefCol)

	return newErrPrefixDto
}

// NewEPrefCollection - Returns a new and properly initialized
// instance of ErrPrefixDto. This method receives a string
// containing one or more error prefix and error context pairs.
// This error prefix information is parsed and stored internally
// as individual error prefix elements.
//
// Upon completion this method returns an integer value identifying
// the number of error prefix elements successfully parsed and
// stored for future use.
//
// String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefixCollection         string
//     - A string consisting of a series of error prefix strings.
//       Error prefixes should be delimited by either a new line
//       character ('\n') or the in-line delimiter string, " - ".
//
//       If the collection string contains error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  numberOfCollectionItemsParsed int
//     - This returned integer value contains the number of error
//       prefix elements parsed from input parameter
//       'errorPrefixCollection' and stored internally for future
//       use.
//
//
//  newErrPrefixDto               ErrPrefixDto
//     - This method will return a newly created instance of
//       ErrPrefixDto configured with the error prefix information
//       parsed from input parameter, 'errorPrefixCollection'.
//
//
func (ePrefDto ErrPrefixDto) NewEPrefCollection(
	errorPrefixCollection string) (
	numberOfCollectionItemsParsed int,
	newErrPrefixDto ErrPrefixDto) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	_ = errPrefixDtoNanobot{}.ptr().setFromString(
		&newErrPrefixDto,
		errorPrefixCollection,
		"")

	numberOfCollectionItemsParsed =
		len(newErrPrefixDto.ePrefCol)

	return numberOfCollectionItemsParsed, newErrPrefixDto
}

// NewFromIErrorPrefix - Receives an object which implements the
// IErrorPrefix interface and returns a new, populated instance of
// ErrPrefixDto.
//
// If the IErrorPrefix ErrorPrefixInfo collection is empty, this
// method will return an empty instance of ErrPrefixDto.
//
// String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  iEPref              IErrorPrefix
//     - An object which implements the IErrorPrefix interface.
//       Information extracted from this object will be used to
//       replicate error prefix information in a new instance of
//       ErrPrefixDto.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a instance of ErrPrefixDto
//       containing a duplicate of error prefix information
//       extracted from input parameter 'iEPref'.
//
func (ePrefDto ErrPrefixDto) NewFromIErrorPrefix(
	iEPref IErrorPrefix) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	var oldErrPrefStr string

	oldErrPrefStr = iEPref.String()

	if len(oldErrPrefStr) == 0 {
		return newErrPrefixDto
	}

	newErrPrefixDto.SetEPrefOld(oldErrPrefStr)

	return newErrPrefixDto
}

// NewFromStrings - Receives a series of string parameters and
// proceeds to generate an return a new, populated ErrPrefixDto
// instance. The string parameters include error prefix and
// associated error context information.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// This version of the 'New' method allows the user to specify both
// the input and output string delimiters. Input string delimiters
// are used to parse the Old Error Prefix string, the New Error
// Prefix string and the associated New Error Context string. The
// output string delimiters will be used to parse and format
// presentation error prefix information returned by methods:
//
//    ErrPrefixDto.String()
//    ErrPrefixDto.StrMaxLineLen()
//
// If all three of the input parameter strings are passed as empty
// strings, this method will simply return an unpopulated instance
// of ErrPrefixDto configured with the input and output string
// delimiters passed by the calling function.
//
// If either of the input or output string delimiter objects are
// judged as invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string which will encapsulate one or more error prefix
//       elements. This string will be parsed into error prefix and
//       error context components and stored in the returned
//       ErrPrefixDto instance.
//
//       This string will be parsed for error prefix and associated
//       error context information based on the input string
//       delimiters contained in input parameter,
//       'inputStrDelimiters'.
//
//       If this string contains associated error context strings
//       as well, they will also be parsed using the same input
//       string delimiters.
//
//       This parameter is optional and will accept an empty string.
//
//
//  newErrPrefix               string
//     - The new error prefix string typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This parameter is optional and will accept an empty string.
//
//
//  newErrContext              string
//     - This is the error context information associated with the
//       new error prefix string ('newErrPrefix').
//
//       This parameter is optional and will accept an empty string.
//
//
//  inputStrDelimiters         ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance when receiving and parsing raw
//       strings containing error prefix information. Such strings
//       are parsed in order to extract error prefix information
//       for storage in the error prefix collection maintained by
//       the current ErrPrefixDto instance.
//
//       If input string delimiters were not directly configured
//       by the user, the system default input string delimiters
//       will be applied.
//
//       The system default input string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//       To monitor input string delimiters, use method:
//             ErrPrefixDto.GetInputStringDelimiters()
//
//       If this input parameter is judged to be invalid, an
//       error will be returned.
//
//
//  outputStrDelimiters        ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance to format error prefix strings
//       returned by methods:
//              ErrPrefixDto.String()
//              ErrPrefixDto.StrMaxLineLen()
//
//       If the input parameter, 'outputStrDelimiters', is invalid,
//       this/ method will return an error.
//
//       If output string delimiters are not directly configured by
//       the user, the default output string delimiters will be
//       applied.
//
//       The default output string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//       To monitor output string delimiters, use method:
//         ErrPrefixDto.GetOutputStringDelimiters()
//
//       If this input parameter is judged to be invalid, an
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'. This parameter is optional.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a new, properly initialized
//       instance of ErrPrefixDto containing error prefix
//       information extracted from input parameter strings,
//       'oldErrPrefix', 'newErrPrefix' and 'newErrContext'.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (ePrefDto ErrPrefixDto) NewFromStrings(
	oldErrPrefix string,
	newErrPrefix string,
	newErrContext string,
	inputStrDelimiters ErrPrefixDelimiters,
	outputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) (ErrPrefixDto, error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += " ErrPrefixDto.NewFromStrings() "

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.ePrefCol = nil

	err :=
		newErrPrefixDto.inputStrDelimiters.
			CopyIn(
				&inputStrDelimiters,
				ePrefix+"newErrPrefixDto.inputStrDelimiters ")

	if err != nil {
		return newErrPrefixDto, err
	}

	err =
		newErrPrefixDto.outputStrDelimiters.
			CopyIn(
				&outputStrDelimiters,
				ePrefix+"newErrPrefixDto.outputStrDelimiters ")

	if err != nil {
		return newErrPrefixDto, err
	}

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		newErrPrefixDto.inputStrDelimiters,
		&newErrPrefixDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		newErrPrefixDto.ePrefCol)

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&newErrPrefixDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		newErrPrefixDto.ePrefCol)

	return newErrPrefixDto, err
}

// NewIBasicErrorPrefix - Receives an object which implements
// the IBasicErrorPrefix interface and returns a new, populated
// instance of ErrPrefixDto.
//
// String delimiters are set to the system defaults.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//       Information extracted from this object will be used to
//       generate error prefix information in a new instance of
//       ErrPrefixDto.
//
//       Be advised, if 'iEPref' is nil, this method will return
//       an error
//
//
//  newErrPrefix        string
//     - A new error prefix represents typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This error prefix information will be added to the
//       internal collection of error prefixes maintained by the
//       new instance of ErrPrefixDto returned by this method.
//
//       If this parameter is set to an empty string, no additional
//       error prefix information will be added to the returned
//       instance of ErrPrefixDto.
//
//
//  newErrContext       string
//     - This is error context information associated with the new
//       error prefix ('newErrPrefix') string described above.
//
//       This parameter is optional and will accept an empty
//       string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - If this method completes successfully, it will return a
//       pointer to a new instance of ErrPrefixDto containing a
//       duplicate of error prefix and error context information
//       extracted from input parameter 'iEPref' plus new error
//       prefix information supplied by parameters, 'newErrPrefix'
//       and 'newErrContext'.
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
//       'newErrPrefix' will be prefixed and attached to the
//       beginning of the error message
//
func (ePrefDto ErrPrefixDto) NewIBasicErrorPrefix(
	iEPref IBasicErrorPrefix,
	newErrPrefix string,
	newErrContext string) (
	*ErrPrefixDto,
	error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := new(ErrPrefixDto)

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	methodName := newErrPrefix +
		"\nErrPrefixDto.NewIBasicErrorPrefix()"

	err := errPrefixDtoNanobot{}.ptr().
		setFromIBasicErrorPrefix(
			newErrPrefixDto,
			iEPref,
			methodName)

	if err != nil {
		return newErrPrefixDto, err
	}

	if len(newErrPrefix) > 0 {

		errPrefNanobot{}.ptr().addEPrefInfo(
			newErrPrefix,
			newErrContext,
			&newErrPrefixDto.ePrefCol)

		errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
			newErrPrefixDto.ePrefCol)

	}

	return newErrPrefixDto, err
}

// NewIEmpty - Receives an empty interface. If that empty interface
// is convertible to one of the 10-valid types listed below, this
// method then proceeds to create and return a pointer to a new
// instance of ErrPrefixDto.
//
// String delimiters are set to the system defaults. To configure
// String delimiters use method:
//        ErrPrefixDto.NewIEmptyWithDelimiters()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  iEPref              interface{}
//     - An empty interface containing one of the 10-valid types
//       listed below. Any one of these valid types will generate
//       valid error prefix and error context information.
//
//       The error prefix and error context information for this
//       new ErrPrefixDto object will be extracted from input
//       parameter, 'iEPref'. 'iEPref' is an empty interface which
//       must be convertible to one of the following valid types:
//
//       1. nil               - A nil value is valid and generates an empty
//                              collection of error prefix and error context
//                              information.
//
//       2. Stringer          - The Stringer interface from the 'fmt' package.
//                              This interface has only one method:
//                                   type Stringer interface {
//                                      String() string
//                                   }
//
//       3. string            - A string containing error prefix information.
//
//       4. []string          - A one-dimensional slice of strings containing
//                              error prefix information.
//
//       5. [][2]string       - A two-dimensional slice of strings
//                              containing error prefix and error context
//                              information.
//
//       6. strings.Builder   - An instance of strings.Builder. Error prefix
//                              information will be imported into the new
//                              returned instance of ErrPrefixDto.
//
//       7  *strings.Builder  - A pointer to an instance of strings.Builder.
//                              Error prefix information will be imported into
//                              the new returned instance of ErrPrefixDto.
//
//       8. ErrPrefixDto      - An instance of ErrPrefixDto. The
//                              ErrorPrefixInfo from this object will be
//                              copied to the new returned instance of
//                              ErrPrefixDto.
//
//       9. *ErrPrefixDto     - A pointer to an instance of ErrPrefixDto.
//                              ErrorPrefixInfo from this object will be
//                              copied to 'errPrefDto'.
//
//      10. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       Any types not listed above will be considered invalid and
//       trigger the return of an error.
//
//
//  newErrPrefix        string
//     - A new error prefix represents typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This error prefix information will be added to the
//       internal collection of error prefixes maintained by the
//       new instance of ErrPrefixDto returned by this method.
//
//       If this parameter is set to an empty string, no additional
//       error prefix information will be added to the returned
//       instance of ErrPrefixDto.
//
//
//  newErrContext       string
//     - This is error context information associated with the new
//       error prefix ('newErrPrefix') string described above.
//
//       This parameter is optional and will accept an empty
//       string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - If this method completes successfully, it will return a
//       pointer to a new instance of ErrPrefixDto containing a
//       duplicate of error prefix and error context information
//       extracted from input parameter 'iEPref' plus new error
//       prefix information supplied by parameters, 'newErrPrefix'
//       and 'newErrContext'.
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
//       'newErrPrefix' will be prefixed and attached to the
//       beginning of the error message.
//
func (ePrefDto ErrPrefixDto) NewIEmpty(
	iEPref interface{},
	newErrPrefix string,
	newErrContext string) (
	*ErrPrefixDto,
	error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := new(ErrPrefixDto)

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	methodName := newErrPrefix +
		"\nErrPrefixDto.NewIEmpty()"

	err := errPrefixDtoMechanics{}.ptr().
		setFromEmptyInterface(
			newErrPrefixDto,
			iEPref,
			methodName)

	if err != nil {
		return newErrPrefixDto, err
	}

	if len(newErrPrefix) > 0 {

		errPrefNanobot{}.ptr().addEPrefInfo(
			newErrPrefix,
			newErrContext,
			&newErrPrefixDto.ePrefCol)

		errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
			newErrPrefixDto.ePrefCol)

	}

	return newErrPrefixDto, err
}

// NewIEmptyWithDelimiters - Receives an empty interface. If that
// empty interface is convertible to one of the 10-valid types
// listed below, this method then proceeds to create and return
// a pointer to a new, fully populated, instance of ErrPrefixDto.
//
// This method allows the user to set String Delimiters used to
// parse string values for incoming error prefix information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  iEPref              interface{}
//     - An empty interface containing one of the 10-valid types
//       listed below. Any one of these valid types will generate
//       valid error prefix and error context information.
//
//       The error prefix and error context information for this
//       new ErrPrefixDto object will be extracted from input
//       parameter, 'iEPref'. 'iEPref' is an empty interface which
//       must be convertible to one of the following valid types:
//
//       1. nil               - A nil value is valid and generates an empty
//                              collection of error prefix and error context
//                              information.
//
//       2. Stringer          - The Stringer interface from the 'fmt' package.
//                              This interface has only one method:
//                                   type Stringer interface {
//                                      String() string
//                                   }
//
//       3. string            - A string containing error prefix information.
//
//       4. []string          - A one-dimensional slice of strings containing
//                              error prefix information.
//
//       5. [][2]string       - A two-dimensional slice of strings
//                              containing error prefix and error context
//                              information.
//
//       6. strings.Builder   - An instance of strings.Builder. Error prefix
//                              information will be imported into the new
//                              returned instance of ErrPrefixDto.
//
//       7  *strings.Builder  - A pointer to an instance of strings.Builder.
//                              Error prefix information will be imported into
//                              the new returned instance of ErrPrefixDto.
//
//       8. ErrPrefixDto      - An instance of ErrPrefixDto. The
//                              ErrorPrefixInfo from this object will be
//                              copied to the new returned instance of
//                              ErrPrefixDto.
//
//       9. *ErrPrefixDto     - A pointer to an instance of ErrPrefixDto.
//                              ErrorPrefixInfo from this object will be
//                              copied to 'errPrefDto'.
//
//      10. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       Any types not listed above will be considered invalid and
//       trigger the return of an error.
//
//
//  newErrPrefix        string
//     - A new error prefix represents typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This error prefix information will be added to the
//       internal collection of error prefixes maintained by the
//       new instance of ErrPrefixDto returned by this method.
//
//       If this parameter is set to an empty string, no additional
//       error prefix information will be added to the returned
//       instance of ErrPrefixDto.
//
//
//  newErrContext       string
//     - This is error context information associated with the new
//       error prefix ('newErrPrefix') string described above.
//
//       This parameter is optional and will accept an empty
//       string.
//
//
//  inputStrDelimiters         ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance when receiving and parsing raw
//       strings containing error prefix information. Such strings
//       are parsed in order to extract error prefix information
//       for storage in the error prefix collection maintained by
//       the current ErrPrefixDto instance.
//
//       If input string delimiters were not directly configured
//       by the user, the system default input string delimiters
//       will be applied.
//
//       The system default input string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//         To monitor input string delimiters, use method:
//             ErrPrefixDto.GetInputStringDelimiters()
//
//       If this input parameter is judged to be invalid, an
//
//
//  outputStrDelimiters        ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance to format error prefix strings
//       returned by methods:
//              ErrPrefixDto.String()
//              ErrPrefixDto.StrMaxLineLen()
//
//       If the input parameter, 'outputStrDelimiters', is invalid,
//       this/ method will return an error.
//
//       If output string delimiters are not directly configured by
//       the user, the default output string delimiters will be
//       applied.
//
//       The default output string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//       To monitor output string delimiters, use method:
//         ErrPrefixDto.GetOutputStringDelimiters()
//
//       If this input parameter is judged to be invalid, an
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'. This parameter is optional.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - If this method completes successfully, it will return a
//       pointer to a new instance of ErrPrefixDto containing a
//       duplicate of error prefix and error context information
//       extracted from input parameter 'iEPref' plus new error
//       prefix information supplied by parameters, 'newErrPrefix'
//       and 'newErrContext'.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message.
//
//       In the event of an error, the value of parameter 'ePrefix'
//       will be prefixed and attached to the beginning of the
//       error message.
//
func (ePrefDto ErrPrefixDto) NewIEmptyWithDelimiters(
	iEPref interface{},
	newErrPrefix string,
	newErrContext string,
	inputStrDelimiters ErrPrefixDelimiters,
	outputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) (
	*ErrPrefixDto,
	error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += newErrPrefix +
		"\nErrPrefixDto.NewIEmptyWithDelimiters()"

	newErrPrefixDto := new(ErrPrefixDto)

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.ePrefCol = nil

	err :=
		newErrPrefixDto.inputStrDelimiters.
			CopyIn(
				&inputStrDelimiters,
				ePrefix+"newErrPrefixDto.inputStrDelimiters ")

	if err != nil {
		return newErrPrefixDto, err
	}

	err =
		newErrPrefixDto.outputStrDelimiters.
			CopyIn(
				&outputStrDelimiters,
				ePrefix+"newErrPrefixDto.outputStrDelimiters ")

	if err != nil {
		return newErrPrefixDto, err
	}

	err = errPrefixDtoMechanics{}.ptr().
		setFromEmptyInterface(
			newErrPrefixDto,
			iEPref,
			ePrefix)

	if err != nil {
		return newErrPrefixDto, err
	}

	if len(newErrPrefix) > 0 {

		errPrefNanobot{}.ptr().addEPrefInfo(
			newErrPrefix,
			newErrContext,
			&newErrPrefixDto.ePrefCol)

		errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
			newErrPrefixDto.ePrefCol)

	}

	return newErrPrefixDto, err
}

// Ptr - Returns a pointer to a new and properly initialized
// instance of ErrPrefixDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a pointer to a new and properly
//       initialized instance of ErrPrefixDto.
//
//
func (ePrefDto ErrPrefixDto) Ptr() *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := new(ErrPrefixDto)

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.ePrefCol = nil

	return newErrPrefixDto
}

// SetCtx - Sets or resets the error context for the last error
// prefix. This operation either adds, or replaces, the error
// context string associated with the last error prefix the
// current list of error prefixes maintained by this instance.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
// If the internal list of error prefixes is currently empty, this
// method will take no action and exit.
//
// If the input parameter string 'newErrContext' is 'empty' (zero
// length string), this method will delete the last error context
// associated with the last error prefix in the error prefix
// collection.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in internal storage already has an
//       associated error context, that context will be deleted and
//       replaced by 'newErrContext'. If, however, the last error
//       prefix does NOT have an associated error context, this
//       'newErrContext' string will be added and associated with
//       that last error prefix.
//
//       If this string is 'empty' (zero length string), this
//       method will delete the last error context associated with
//       the last error prefix in the error prefix collection.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetCtx(
	newErrContext string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	if len(ePrefDto.ePrefCol) == 0 {
		return
	}

	if len(newErrContext) == 0 {

		errPrefixDtoNanobot{}.ptr().
			deleteLastErrContext(ePrefDto)

		return
	}

	errPrefNanobot{}.ptr().
		setLastCtx(
			newErrContext,
			ePrefDto.ePrefCol)
}

// SetCtxEmpty - Deletes the last error context for the last error
// prefix in this instance of ErrPrefixDto.
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
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetCtxEmpty() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	collectionIdx := len(ePrefDto.ePrefCol)

	if collectionIdx == 0 {
		return
	}

	errPrefixDtoNanobot{}.ptr().
		deleteLastErrContext(
			ePrefDto)

	return
}

// SetEPref - Adds an error prefix string to the list of
// previous error prefix strings maintained by this instance of
// ErrPrefixDto.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
// This method is designed to process a single error prefix string
// passed in input parameter 'newErrPrefix'. If this string
// contains multiple error prefixes, use method
// 'ErrPrefixDto.SetEPrefOld()'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This new error prefix will be added to the internal list
//       of error prefixes maintained by this ErrPrefixDto
//       instance.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetEPref(
	newErrPrefix string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return
}

// SetEPrefCollection - Deletes the current error prefix collection
// and replaces it with a new collection passed as an input
// parameter to this method.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
// If the input array 'newEPrefCollection' has zero elements, then
// the current ErrorPrefixInfo instance will be configured with an
// array of zero elements thereby effectively deleting all previous
// existing error prefix information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newEPrefCollection  []ErrorPrefixInfo
//     - A collection of ErrorPrefixInfo objects. This collection
//       will replace the internal collection of error prefix
//       information objects maintained by the current ErrorPrefixInfo
//       instance.
//
//       If this array has zero elements, then the current ErrorPrefixInfo
//       instance will be configured with an array of zero elements.
//

func (ePrefDto *ErrPrefixDto) SetEPrefCollection(
	newEPrefCollection []ErrorPrefixInfo) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	lenNewEPrefCol := len(newEPrefCollection)

	ePrefDto.ePrefCol = make(
		[]ErrorPrefixInfo,
		lenNewEPrefCol)

	if lenNewEPrefCol == 0 {
		return
	}

	copy(ePrefDto.ePrefCol, newEPrefCollection)
}

// SetEPrefCtx - Adds an error prefix and an error context string
// to the list of previous error prefix/context information.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This method is designed to process a single new error prefix
//       string. To process a collection of error prefix strings, see
//       method 'ErrPrefixDto.SetEPrefOld()'.
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
func (ePrefDto *ErrPrefixDto) SetEPrefCtx(
	newErrPrefix string,
	newErrContext string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return
}

// SetEPrefOld - Deletes the current error prefix information and
// initializes this ErrPrefixDto instance with an old or preexisting
// error prefix string. This error prefix string input parameter
// typically includes one or more error prefix elements and may
// also include associated error context elements. This string will
// be parsed and into individual error prefix and error context
// components and stored internally in the current ErrPrefixDto
// instance.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components and stored in the current
//       ErrPrefixDto instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetEPrefOld(
	oldErrPrefix string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	_ = errPrefixDtoQuark{}.ptr().
		emptyErrPrefInfoCollection(
			ePrefDto,
			"")

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		ePrefDto.inputStrDelimiters,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return
}

// SetEPrefStrings - This method first deletes all pre-existing
// error prefix and error context information and then replaces
// that information with new data extracted from the
// two-dimensional string array passed as an input parameter
// ('twoDStrArray').
//
// This two-dimensional string array contains both error prefix and
// error context information. The Error Prefix string is always in
// the [x][0] position. The Error Context string is always in the
// [x][1] position. The Error Context string is optional and may be
// an empty string.
//
// If input parameter 'twoDStrArray' is 'nil' or a zero length
// array, this method will take no action and exit.
//
// To 'append' or 'add' a two-dimensional string array to existing
// error prefix information, see method:
//    ErrPrefixDto.AddEPrefStrings()
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
func (ePrefDto *ErrPrefixDto) SetEPrefStrings(
	twoDStrArray [][2]string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if len(twoDStrArray) == 0 {
		return
	}

	_ = errPrefixDtoQuark{}.ptr().emptyErrPrefInfoCollection(
		ePrefDto,
		"")

	_ = errPrefixDtoAtom{}.ptr().addTwoDimensionalStringArray(
		ePrefDto,
		twoDStrArray,
		"")
}

// SetIBasic - Deletes the current error prefix information and
// replaces it with new error prefix data passed to this method by
// an input parameter object implementing the IBasicErrorPrefix
// interface.
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
//  incomingIBasic      IBasicErrorPrefix
//     - An object implementing the IBasicErrorPrefix interface.
//       Error prefix and context information will be generated
//       from this object and used to overwrite and replace the
//       existing error prefix and context information contained in
//       the current ErrPrefixDto instance.
//
//
//  ePrefix             string
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
//       In the event of an error, the value of parameter 'ePrefix'
//       will be prefixed and attached to the beginning of the
//       error message.
//
func (ePrefDto *ErrPrefixDto) SetIBasic(
	incomingIBasic IBasicErrorPrefix,
	ePrefix string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix = ePrefix +
		"\nErrPrefixDto.SetIBasic()"

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	return errPrefixDtoNanobot{}.ptr().
		setFromIBasicErrorPrefix(
			ePrefDto,
			incomingIBasic,
			ePrefix)

}

// SetIBuilder - Deletes the current error prefix information and
// replaces it with new error prefix data passed to this method by
// an input parameter object implementing the IBuilderErrorPrefix
// interface.
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
//  inComingIBuilder    IBuilderErrorPrefix
//     - An object implementing the IBuilderErrorPrefix interface.
//       Error prefix and context information will be generated
//       from this object and used to overwrite and replace the
//       existing error prefix and context information contained in
//       the current ErrPrefixDto instance.
//
//
//  callingMethodName   string
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
//       'callingMethodName' will be prefixed and attached to the
//       beginning of the error message.
//
func (ePrefDto *ErrPrefixDto) SetIBuilder(
	inComingIBuilder IBuilderErrorPrefix,
	callingMethodName string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	callingMethodName = callingMethodName +
		"\nErrPrefixDto.SetIBuilder()"

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	return errPrefixDtoNanobot{}.ptr().
		setFromIBuilder(
			ePrefDto,
			inComingIBuilder,
			callingMethodName)
}

// SetIEmpty - Deletes the current error prefix information
// and replaces it with new error prefix data passed to this method
// through an empty interface.
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
//  iEPref              interface{}
//     - An empty interface containing one of the 10-valid types
//       listed below. Any one of these valid types will generate
//       valid error prefix and error context information. This
//       error prefix data will then overwrite and replace that
//       encapsulated in the current ErrPrefixDto instance.
//
//       The error prefix and error context information for this
//       new ErrPrefixDto object will be extracted from input
//       parameter, 'iEPref'. 'iEPref' is an empty interface which
//       must be convertible to one of the following valid types:
//
//       1. nil               - A nil value is valid and generates an empty
//                              collection of error prefix and error context
//                              information.
//
//       2. Stringer          - The Stringer interface from the 'fmt' package.
//                              This interface has only one method:
//                                   type Stringer interface {
//                                      String() string
//                                   }
//
//       3. string            - A string containing error prefix information.
//
//       4. []string          - A one-dimensional slice of strings containing
//                              error prefix information.
//
//       5. [][2]string       - A two-dimensional slice of strings
//                              containing error prefix and error context
//                              information.
//
//       6. strings.Builder   - An instance of strings.Builder. Error prefix
//                              information will be imported into the new
//                              returned instance of ErrPrefixDto.
//
//       7  *strings.Builder  - A pointer to an instance of strings.Builder.
//                              Error prefix information will be imported into
//                              the new returned instance of ErrPrefixDto.
//
//       8. ErrPrefixDto      - An instance of ErrPrefixDto. The
//                              ErrorPrefixInfo from this object will be
//                              copied to the new returned instance of
//                              ErrPrefixDto.
//
//       9. *ErrPrefixDto     - A pointer to an instance of ErrPrefixDto.
//                              ErrorPrefixInfo from this object will be
//                              copied to 'errPrefDto'.
//
//      10. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       Any types not listed above will be considered invalid and
//       trigger the return of an error.
//
//
//  callingMethodName   string
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
//       'callingMethodName' will be prefixed and attached to the
//       beginning of the error message.
//
func (ePrefDto *ErrPrefixDto) SetIEmpty(
	iEPref interface{},
	callingMethodName string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	callingMethodName = callingMethodName +
		"\nErrPrefixDto.SetIEmpty()"

	err := errPrefixDtoMechanics{}.ptr().
		setFromEmptyInterface(
			ePrefDto,
			iEPref,
			callingMethodName)

	return err
}

// SetInputStringDelimiters - Sets the string delimiters used by
// the current ErrPrefixDto instance when receiving and parsing raw
// strings containing error prefix information. Such strings are
// parsed in order to extract error prefix information for storage
// in the error prefix collection maintained by the current
// ErrPrefixDto instance.
//
// If input string delimiters were not directly configured by the
// user, the system default input string delimiters will be
// applied.
//
// The system default input string delimiters are listed as
// follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
//
// To monitor input string delimiters, use method:
//    ErrPrefixDto.GetInputStringDelimiters()
//
// To set and control output string delimiters, use method:
//    ErrPrefixDto.SetOutputStringDelimiters()
//
func (ePrefDto *ErrPrefixDto) SetInputStringDelimiters(
	inputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += " ErrPrefixDto.SetInputStringDelimiters() \n" +
		"inputStrDelimiters "

	err :=
		ePrefDto.inputStrDelimiters.
			CopyIn(
				&inputStrDelimiters,
				ePrefix)

	return err
}

// SetIsLastLineTermWithNewLine - By default, the last line of
// error prefix strings returned by the method
// ErrPrefixDto.String() ARE NOT terminated with a new line
// character ('\n'). In other words, by default, the last line of
// returned error prefix strings do not end with a new line
// character ('\n').
//
// If the user prefers to terminate the last line of error prefix
// strings returned by this instance of ErrPrefixDto with a new
// line character ('\n'), the input parameter,
// 'isLastLineTerminatedWithNewLine', should be set to 'true'.
//
// Error prefix strings are returned by calling method
// 'ErrPrefixDto.String()'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  isLastLineTerminatedWithNewLine    bool
//     - If this parameter is set to 'true', the last line of all
//       error prefix strings returned by this ErrPrefixDto
//       instance WILL BE terminated with a new line character
//       ('\n').
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetIsLastLineTermWithNewLine(
	isLastLineTerminatedWithNewLine bool) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.isLastLineTerminatedWithNewLine =
		isLastLineTerminatedWithNewLine
}

// SetLeftMarginChar - Sets the character used in creating the left
// margin applied to all new lines generated in the error prefix
// string returned by method: ErrPrefixDto.String().
//
// The default Left Margin Character is the empty space character
// (' '). Therefore, if the rune input parameter,
// 'leftMarginCharacter' is set to zero, it will default the Left
// Margin Character to an empty space value (' ').
//
// Be Advised: If the Left Margin Length is set to zero, this
// character will never be used and the no left margin will be
// added to error prefix strings generated by method:
// ErrPrefixDto.String().
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginCharacter rune
//     - This character will be used to configure the left margin
//       for error prefix strings returned by method
//       ErrPrefixDto.String(). If this value is set to zero, it
//       will engage the default Left Margin Character which is the
//       empty space character (' ').
//
//       If the Left Margin Length is set to zero, no left margin
//       will be generated for error prefix strings and this
//       character will never be used. Reference method
//       ErrPrefixDto.SetLeftMarginLength().
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
// ----------------------------------------------------------------
//
// Example Usage
//
// Example of Left Margin Character set to '*'.
// Left Margin Length set to 3.
// Maximum Line Length set to 40.
// Effective Maximum Line Length is 43.
//
//    "***Tx1.Something() - Tx2.SomethingElse()   "
//    "***Tx3.DoSomething() - Tx4() - Tx5()       "
//    "***Tx6.DoSomethingElse()                   "
//    "***Tx7.TrySomethingNew()                   "
//    "*** :  something->newSomething             "
//    "***Tx8.TryAnyCombination()                 "
//    "***Tx9.TryAHammer() : x->y - Tx10.X()      "
//    "***Tx11.TryAnything() - Tx12.TryASalad()   "
//    "***Tx13.SomeFabulousAndComplexStuff()      "
//    "***Tx14.MoreAwesomeGoodness                "
//    "*** :  A=7 B=8 C=9                         "
//
func (ePrefDto *ErrPrefixDto) SetLeftMarginChar(
	leftMarginCharacter rune) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.leftMarginChar = leftMarginCharacter
}

// SetLeftMarginLength - Sets the length of the left margin applied
// to all new lines generated in the error prefix string returned
// by method: ErrPrefixDto.String()
//
// The Left Margin Character can be set through method:
//     ErrPrefixDto.SetLeftMarginChar()
//
// If the Left Margin value is less than zero, this method will
// take no action and exit.
//
// BE CAREFUL
//
// This method will accept any 'leftMarginLength' greater than
// zero.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginLength    int
//     - This value will be used to determine the length of the
//       left margin configured in error prefix strings returned
//       by method ErrPrefixDto.String().
//
//       If this value is less than zero, this method will take no
//       action and exit.
//
//       Be advised, any left margin length greater than zer will
//       be accepted.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
// ----------------------------------------------------------------
//
// Example Usage
//
// Example of Left Margin Length set to 3.
// Left Margin Character set to empty space ' '.
// Maximum Line Length set to 40.
// Effective Maximum Line Length is 43.
//
//    "   Tx1.Something() - Tx2.SomethingElse()   "
//    "   Tx3.DoSomething() - Tx4() - Tx5()       "
//    "   Tx6.DoSomethingElse()                   "
//    "   Tx7.TrySomethingNew()                   "
//    "    :  something->newSomething             "
//    "   Tx8.TryAnyCombination()                 "
//    "   Tx9.TryAHammer() : x->y - Tx10.X()      "
//    "   Tx11.TryAnything() - Tx12.TryASalad()   "
//    "   Tx13.SomeFabulousAndComplexStuff()      "
//    "   Tx14.MoreAwesomeGoodness                "
//    "    :  A=7 B=8 C=9                         "
//
func (ePrefDto *ErrPrefixDto) SetLeftMarginLength(
	leftMarginLength int) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if leftMarginLength < 0 {
		return
	}

	ePrefDto.leftMarginLength = leftMarginLength
}

// SetMaxTextLineLen - Sets the maximum limit on the number of
// characters allowed in an error prefix text line output for
// display purposes.
//
// Setting this value will control the maximum character limit for
// this specific ErrPrefixDto instance. This maximum limit will
// remain in force for the life of this ErrPrefixDto instance or
// until another call is made to this method changing the value.
//
// When this instance of ErrPrefixDto was initially created, a
// default value of 40-characters was applied.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  maxErrPrefixTextLineLength      int
//     - This unsigned integer value will be used to set the
//       maximum number of characters allowed in a text display
//       line for error prefix information.
//
//       If 'maxErrPrefixTextLineLength' is set to a value less
//       than ten (10) or to a value greater than one-million
//       (1,000,000), the default value of 40-characters will be
//       applied.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
func (ePrefDto *ErrPrefixDto) SetMaxTextLineLen(
	maxErrPrefixTextLineLength int) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

		return
	}

	ePrefDto.maxErrPrefixTextLineLength =
		uint(maxErrPrefixTextLineLength)
}

// SetMaxTextLineLenToDefault - Maximum Error Prefix Line
// Length is the maximum limit on the number of characters allowed
// in a single error prefix text line.
//
// This method resets that maximum limit to its default value of
// 40-characters.
//
func (ePrefDto *ErrPrefixDto) SetMaxTextLineLenToDefault() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()
}

// SetOutputStringDelimiters - Sets the string delimiters used by
// the current ErrPrefixDto instance to format error prefix strings
// returned by methods:
//    ErrPrefixDto.String()
//    ErrPrefixDto.StrMaxLineLen()
//
// If the input parameter, 'outputStrDelimiters', is invalid, this
// method will return an error.
//
// If output string delimiters were not directly configured by the
// user, the default output string delimiters will be applied.
//
// The default output string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// To monitor output string delimiters, use method:
//    ErrPrefixDto.GetOutputStringDelimiters()
//
// To set and control input string delimiters, use method:
//    ErrPrefixDto.SetInputStringDelimiters()
//
func (ePrefDto *ErrPrefixDto) SetOutputStringDelimiters(
	outputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += " ErrPrefixDto.SetOutputStringDelimiters() "

	err :=
		ePrefDto.outputStrDelimiters.
			CopyIn(
				&outputStrDelimiters,
				ePrefix)

	return err
}

// SetStrDelimitersToDefault - Sets the Input String Delimiters and
// the Output String Delimiters to their system default values.
//
// The ErrPrefixDto type maintains two types of string delimiters,
// Input String Delimiters and Output String Delimiters.
//
// Input String Delimiters are used by ErrPrefixDto instances to
// parse raw string values containing error prefix and error
// context information received as input parameters from external
// sources. Methods performing this type of operation include but
// are not limited to:
//    ErrPrefixDto.NewIEmptyWithDelimiters()
//    ErrPrefixDto.NewFromStrings()
//    ErrPrefixDto.NewEPrefOld()
//    ErrPrefixDto.SetEPrefOld()
//
// Output String Delimiters are used by ErrPrefixDto instances to
// join or concatenate individual error prefix and error context
// components to form presentation text for output and use in
// preparation of error message strings. Methods performing this
// type of operation are:
//
//    ErrPrefixDto.String()
//    ErrPrefixDto.StrMaxLineLen()
//
// Initially, both the Input and Output System Delimiters for any
// given ErrPrefixDto instance are set to the system default
// values. This means that if the Input and Output String
// Delimiters were not directly configured by the user, the system
// default string delimiters are applied.
//
// The system defaults for both Input and Output String Delimiters
// are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// Users have the option of setting custom Input and Output String
// Delimiters using methods:
//
//    ErrPrefixDto.SetOutputStringDelimiters()
//    ErrPrefixDto.SetInputStringDelimiters()
//
// The current settings for String Delimiters can be monitored
// using methods:
//    ErrPrefixDto.GetStrDelimiters()
//    ErrPrefixDto.GetInputStringDelimiters()
//    ErrPrefixDto.GetOutputStringDelimiters()
//
// In the event it becomes necessary to reset the Input and Output
// String Delimiters to their original system default values, call
// this method.
//
func (ePrefDto *ErrPrefixDto) SetStrDelimitersToDefault() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()
	ePrefDto.inputStrDelimiters.SetToDefault()

	ePrefDto.outputStrDelimiters.SetToDefault()
}

// SetTurnOffTextDisplay - Controls the "Turn Off Text Display"
// flag represented by internal member variable:
//    ErrPrefixDto.turnOffTextDisplay
//
// If this flag (input parameter 'turnOffTextDisplay') is set to
// 'true', this instance of ErrPrefixDto WILL NOT format and return
// error prefix and context information through method,
// ErrPrefixDto.String(). In this case, the method
// ErrPrefixDto.String() will instead return an empty string.
//
// Conversely, if this flag (input parameter 'turnOffTextDisplay')
// is set to 'false' formatted error prefix and context information
// WILL BE formatted and returned through method
// ErrPrefixDto.String().
//
// To monitor the status of the "Turn Off Text Display" flag, see
// method ErrPrefixDto.GetTurnOffTextDisplay().
//
func (ePrefDto *ErrPrefixDto) SetTurnOffTextDisplay(
	turnOffTextDisplay bool) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.turnOffTextDisplay =
		turnOffTextDisplay
}

// String - Returns a formatted error prefix/context string
// incorporating all error prefixes previously added to this
// ErrPrefixDto instance.
//
// Error prefix information is stored internally in the
// 'ErrPrefixDto.ePrefCol' array.
//
// If the Left Margin has been set to a value greater than zero,
// that number of Left Margin Characters will be formatted in
// each new line of the returned error prefix string.
//
// The Left Margin value can be set through method:
//     ErrPrefixDto.SetLeftMarginLength()
//
// The Left Margin Character can be set through method:
//     ErrPrefixDto.SetLeftMarginChar()
//
// The default Left Margin value is zero. The default Left Margin
// Character is the empty space character (' ').
//
// Remember that configuring a left margin for the output error
// prefix string will extend the existing line length. So, if the
// maximum line length is set to 40 with a left margin length of 3,
// the effective maximum line length is 43.
//
// The error prefix and context information returned by this method
// will be separated and formatted using previously configured
// output string delimiters. If these output string delimiters were
// not directly configured by the user, the default output string
// delimiters will be used.
//
// The default output string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// To set and control output string delimiters, use method:
//    ErrPrefixDto.SetOutputStringDelimiters()
//
//
func (ePrefDto ErrPrefixDto) String() string {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.turnOffTextDisplay ||
		len(ePrefDto.ePrefCol) == 0 {
		return ""
	}

	if ePrefDto.maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	outPutStr := errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			ePrefDto.maxErrPrefixTextLineLength,
			ePrefDto.isLastLineTerminatedWithNewLine,
			ePrefDto.outputStrDelimiters,
			ePrefDto.ePrefCol)

	if ePrefDto.leftMarginLength > 0 {
		leftMarChar := ' '

		if ePrefDto.leftMarginChar > 0 {
			leftMarChar = ePrefDto.leftMarginChar
		}

		leftMarStr := strings.Repeat(
			string(leftMarChar),
			ePrefDto.leftMarginLength)

		outPutStr = leftMarStr + outPutStr

		leftMarStr = "\n" + leftMarStr

		outPutStr = strings.ReplaceAll(
			outPutStr, "\n", leftMarStr)

	}

	return outPutStr
}

// StrMaxLineLen - Returns a formatted error prefix/context string
// incorporating all error prefix and error context information
// previously added to this ErrPrefixDto instance.
//
// Error prefix information is stored internally in the 'ePrefCol'
// array.
//
// Input parameter 'maxLineLen' is used to set the Maximum Text
// Line length Limit (in characters) for text returned by this
// method. If the value of 'maxLineLen' is less than 10, it will
// be discarded and the current Maximum Text Line Length Limit
// value will be used. If input parameter 'maxLineLen' is equal to
// or greater than 10, the Maximum Text Line Length Limit for this
// current ErrPrefixDto instance will be reset and configured to
// this new value.
//
// If the "Turn Off Text Display Flag" is current set to 'true',
// this method will return an empty string. The default value for
// the "Turn Off Text Display Flag" is 'false'. To control the
// "Turn Off Text Display Flag" use method:
//    ErrPrefixDto.SetTurnOffTextDisplay()
//
// The error prefix string returned by this method will be
// formatted using output string delimiters previously configured
// for this instance of ErrPrefixDto.
//
// If output string delimiters are invalid or were not directly
// configured by the user, the default output string delimiters
// will be applied.
//
// The default output string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
// To set and control output string delimiters use method:
//    ErrPrefixDto.SetOutputStringDelimiters()
//
func (ePrefDto *ErrPrefixDto) StrMaxLineLen(
	maxLineLen int) string {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	maxErrPrefixTextLineLength := uint(maxLineLen)

	if maxErrPrefixTextLineLength < 10 {

		if ePrefDto.maxErrPrefixTextLineLength < 10 {

			ePrefDto.maxErrPrefixTextLineLength =
				errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

		}

		maxErrPrefixTextLineLength =
			ePrefDto.maxErrPrefixTextLineLength

	} else {

		ePrefDto.maxErrPrefixTextLineLength =
			maxErrPrefixTextLineLength
	}

	if ePrefDto.turnOffTextDisplay ||
		len(ePrefDto.ePrefCol) == 0 {
		return ""
	}

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			maxErrPrefixTextLineLength,
			ePrefDto.isLastLineTerminatedWithNewLine,
			ePrefDto.outputStrDelimiters,
			ePrefDto.ePrefCol)
}

// XCtx - Sets or resets the error context for the last error
// prefix. This operation either adds, or replaces, the error
// context string associated with the last error prefix the
// current list of error prefixes maintained by this instance.
//
// This method is identical in function to ErrPrefixDto.SetCtx().
// The only difference is that this method returns a pointer to the
// current ErrPrefixDto instance.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
// If the internal list of error prefixes is currently empty, this
// method will take no action and exit.
//
// If the input parameter string 'newErrContext' is 'empty' (zero
// length string), this method will delete the last error context
// associated with the last error prefix in the error prefix
// collection.
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in internal storage already has an
//       associated error context, that context will be deleted and
//       replaced by 'newErrContext'. If, however, the last error
//       prefix does NOT have an associated error context, this
//       'newErrContext' string will be added and associated with
//       that last error prefix.
//
//       If this string is 'empty' (zero length string), this
//       method will delete the last error context associated with
//       the last error prefix in the error prefix collection.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - Returns a pointer to the current ErrPrefixDto instance
//
func (ePrefDto *ErrPrefixDto) XCtx(
	newErrContext string) *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if len(ePrefDto.ePrefCol) == 0 {
		return ePrefDto
	}

	if len(newErrContext) == 0 {

		errPrefixDtoNanobot{}.ptr().
			deleteLastErrContext(ePrefDto)

		return ePrefDto
	}

	errPrefNanobot{}.ptr().setLastCtx(
		newErrContext,
		ePrefDto.ePrefCol)

	return ePrefDto
}

// XCtxEmpty - Deletes the last error context for the last error
// prefix in this instance of ErrPrefixDto
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
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
//  *ErrPrefixDto
//     - Returns a pointer to the current ErrPrefixDto instance
//
func (ePrefDto *ErrPrefixDto) XCtxEmpty() *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if len(ePrefDto.ePrefCol) == 0 {
		return ePrefDto
	}

	errPrefixDtoNanobot{}.ptr().
		deleteLastErrContext(
			ePrefDto)

	return ePrefDto
}

// XEPref - Adds an error prefix string to the list of previous
// error prefix strings maintained by this instance of ErrPrefixDto.
//
// This method is identical in function to ErrPrefixDto.SetEPref().
// The only difference is that this method returns a pointer to the
// current ErrPrefixDto instance.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
// This method is designed to process a single error prefix string
// passed in input parameter 'ErrPrefixDto'. If this string
// contains multiple error prefixes, use method
// 'ErrPrefixDto.SetEPrefOld()'.
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This new error prefix will be added to the internal list
//       of error prefixes maintained by this ErrPrefixDto
//       instance.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - Returns a pointer to the current ErrPrefixDto instance
//
func (ePrefDto *ErrPrefixDto) XEPref(
	newErrPrefix string) *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return ePrefDto
}

// XEPrefCtx - Adds an error prefix and an error context string
// to the list of previous error prefix/context information.
//
// This method is identical in function to ErrPrefixDto.SetEPrefCtx().
// The only difference is that this method returns a pointer to the
// current ErrPrefixDto instance.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This method is designed to process a single new error prefix
//       string. To process a collection of error prefix strings, see
//       method 'ErrPrefixDto.SetEPrefOld()'.
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
//  *ErrPrefixDto
//     - Returns a pointer to the current ErrPrefixDto instance.
//
func (ePrefDto *ErrPrefixDto) XEPrefCtx(
	newErrPrefix string,
	newErrContext string) *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return ePrefDto
}

// XEPrefOld - Deletes the current error prefix information and
// initializes this ErrPrefixDto instance with an old or preexisting
// error prefix string. This error prefix string input parameter
// typically includes one or more error prefix elements and may
// also include associated error context elements. This string will
// be parsed and into individual error prefix and error context
// components and stored internally in the current ErrPrefixDto
// instance.
//
// This method is identical in function to ErrPrefixDto.SetEPrefOld().
// The only difference is that this method returns a pointer to the
// current ErrPrefixDto instance.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
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
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components and stored in the current
//       ErrPrefixDto instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - Returns a pointer to the current ErrPrefixDto instance
//
func (ePrefDto *ErrPrefixDto) XEPrefOld(
	oldErrPrefix string) *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0)

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		ePrefDto.inputStrDelimiters,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return ePrefDto
}

// XSetFromStrings - Receives a series of string parameters and
// proceeds to delete and overwrite the existing error prefix
// information contained in the current ErrPrefixDto instance.
// The string parameters include error prefix and associated
// error context information.
//
// In addition, this method allows the user to specify both the
// input and output string delimiters. Input string delimiters
// are used to parse the Old Error Prefix string, the New Error
// Prefix string and the associated New Error Context string. The
// output string delimiters will be used to parse and format
// presentation error prefix information returned by methods:
//
//    ErrPrefixDto.String()
//    ErrPrefixDto.StrMaxLineLen()
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// If all three of the input parameter strings are passed as empty
// strings, this method will simply create an unpopulated or empty
// instance of ErrPrefixDto configured with the input and output
// string delimiters passed by the calling function.
//
// If either of the input or output string delimiter objects are
// judged as invalid, this method will return an error.
//
// Be advised: The 'X' methods always return a pointer to current,
// modified, ErrPrefixDto instance.
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
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string which will encapsulate one or more error prefix
//       elements. This string will be parsed into error prefix and
//       error context components and stored in the returned
//       ErrPrefixDto instance.
//
//       This string will be parsed for error prefix and associated
//       error context information based on the input string
//       delimiters contained in input parameter,
//       'inputStrDelimiters'.
//
//       If this string contains associated error context strings
//       as well, they will also be parsed using the same input
//       string delimiters.
//
//       This parameter is optional and will accept an empty string.
//
//
//  newErrPrefix               string
//     - The new error prefix string typically identifies the
//       function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This parameter is optional and will accept an empty string.
//
//
//  newErrContext              string
//     - This is the error context information associated with the
//       new error prefix string ('newErrPrefix').
//
//       This parameter is optional and will accept an empty string.
//
//
//  inputStrDelimiters         ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance when receiving and parsing raw
//       strings containing error prefix information. Such strings
//       are parsed in order to extract error prefix information
//       for storage in the error prefix collection maintained by
//       the current ErrPrefixDto instance.
//
//       If input string delimiters were not directly configured
//       by the user, the system default input string delimiters
//       will be applied.
//
//       The system default input string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//         To monitor input string delimiters, use method:
//             ErrPrefixDto.GetInputStringDelimiters()
//
//
//  outputStrDelimiters        ErrPrefixDelimiters
//     - Sets the string delimiters used by the current
//       ErrPrefixDto instance to format error prefix strings
//       returned by methods:
//              ErrPrefixDto.String()
//              ErrPrefixDto.StrMaxLineLen()
//
//       If the input parameter, 'outputStrDelimiters', is invalid,
//       this/ method will return an error.
//
//       If output string delimiters are not directly configured by
//       the user, the default output string delimiters will be
//       applied.
//
//       The default output string delimiters are listed as
//       follows:
//
//         New Line Error Prefix Delimiter = "\n"
//         In-Line Error Prefix Delimiter  = " - "
//         New Line Error Context Delimiter = "\n :  "
//         In-Line Error Context Delimiter = " : "
//
//       To monitor output string delimiters, use method:
//         ErrPrefixDto.GetOutputStringDelimiters()
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'. This parameter is optional.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *ErrPrefixDto
//     - Returns a pointer to the current, modified ErrPrefixDto
//       instance.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Usage Examples
//
//  ePDto2 := ErrPrefixDto{}
//  var err error
//
//   ePDto2,
//      err = ePDto2.
//        XSetFromStrings(
//          initialStr,
//          newErrPrefix,
//          newErrCtx,
//          outputDelimiters,
//          outputDelimiters,
//          funcName)
//
//  ---------------------------------------
//  var ePDto2 *ErrPrefixDto
//  var err error
//
//  ePDto2,
//  err = ErrPrefixDto{}.Ptr().
//        XSetFromStrings(
//          initialStr,
//          newErrPrefix,
//          newErrCtx,
//          outputDelimiters,
//          outputDelimiters,
//          funcName)
//
func (ePrefDto *ErrPrefixDto) XSetFromStrings(
	oldErrPrefix string,
	newErrPrefix string,
	newErrContext string,
	inputStrDelimiters ErrPrefixDelimiters,
	outputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) (*ErrPrefixDto, error) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefix += " ErrPrefixDto.XSetFromStrings() "

	err := inputStrDelimiters.IsValidInstanceError(
		ePrefix + "inputStrDelimiters ")

	if err != nil {
		return ePrefDto, err
	}

	err = outputStrDelimiters.IsValidInstanceError(
		ePrefix + "outputStrDelimiters ")

	if err != nil {
		return ePrefDto, err
	}

	err = errPrefixDtoQuark{}.ptr().
		emptyErrPrefInfoCollection(
			ePrefDto,
			ePrefix)

	if err != nil {
		return ePrefDto, err
	}

	err =
		ePrefDto.inputStrDelimiters.
			CopyIn(
				&inputStrDelimiters,
				ePrefix+"ePrefDto.inputStrDelimiters ")

	if err != nil {
		return ePrefDto, err
	}

	err =
		ePrefDto.outputStrDelimiters.
			CopyIn(
				&outputStrDelimiters,
				ePrefix+"ePrefDto.outputStrDelimiters ")

	if err != nil {
		return ePrefDto, err
	}

	ePrefAtom := errPrefixDtoAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		ePrefDto.inputStrDelimiters,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return ePrefDto, err
}

// ZCtx - Sets or resets the error context for the last error
// prefix. This operation either adds, or replaces, the error
// context string associated with the last error prefix the
// current list of error prefixes maintained by this instance.
//
// This method is identical in function to ErrPrefixDto.SetCtx().
// The only difference is that this method returns a deep copy of
// the current ErrPrefixDto instance.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
// If the internal list of error prefixes is currently empty, this
// method will take no action and exit.
//
// If the input parameter string 'newErrContext' is 'empty' (zero
// length string), this method will delete the last error context
// associated with the last error prefix in the error prefix
// collection.
//
// This method returns a deep copy of the current ErrPrefixDto
// instance after it is updated with new error context information.
// This copy of ErrPrefixDto returned by value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in internal storage already has an
//       associated error context, that context will be deleted and
//       replaced by 'newErrContext'. If, however, the last error
//       prefix does NOT have an associated error context, this
//       'newErrContext' string will be added and associated with
//       that last error prefix.
//
//       If this string is 'empty' (zero length string), this
//       method will delete the last error context associated with
//       the last error prefix in the error prefix collection.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method returns a deep copy of the current
//       ErrPrefixDto instance after it is updated with new error
//       context information. This copy is returned by value.
//
func (ePrefDto *ErrPrefixDto) ZCtx(
	newErrContext string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	var newErrPrefixDto ErrPrefixDto

	if len(ePrefDto.ePrefCol) == 0 {

	} else if len(newErrContext) == 0 {

		errPrefixDtoNanobot{}.ptr().
			deleteLastErrContext(ePrefDto)

	} else {

		errPrefNanobot{}.ptr().
			setLastCtx(
				newErrContext,
				ePrefDto.ePrefCol)

	}

	newErrPrefixDto,
		_ = errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}

// ZCtxEmpty - Deletes the last error context for the last error
// prefix in this instance of ErrPrefixDto.
//
// This method returns a deep copy of the current ErrPrefixDto
// instance after it is updated with the deletion of the last
// error context. This copy is returned by value.
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
//  ErrPrefixDto
//     - This method returns a deep copy of the current
//       ErrPrefixDto instance after it is updated with the
//       deletion of the last error context. This copy is returned
//       by value.
//
func (ePrefDto *ErrPrefixDto) ZCtxEmpty() ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if len(ePrefDto.ePrefCol) > 0 {

		errPrefixDtoNanobot{}.ptr().
			deleteLastErrContext(
				ePrefDto)

	}

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}

// ZEPref - Adds an error prefix string to the list of previous
// error prefix strings maintained by this instance of ErrPrefixDto.
//
// This method is identical in function to ErrPrefixDto.SetEPref().
// The only difference is that this method returns an ErrPrefixDto
// instance by value.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
// This method is designed to process a single error prefix string
// passed in input parameter 'ErrPrefixDto'. If this string
// contains multiple error prefixes, use method
// 'ErrPrefixDto.SetEPrefOld()'.
//
// This method returns a deep copy of the current ErrPrefixDto
// instance after it is updated with new error prefix information.
// This copy of ErrPrefixDto is returned by value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This new error prefix will be added to the internal list
//       of error prefixes maintained by this ErrPrefixDto
//       instance.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method returns a deep copy of the current
//       ErrPrefixDto instance after it is updated with
//       new error prefix information. This copy of
//       ErrPrefixDto is returned by value.
//
func (ePrefDto *ErrPrefixDto) ZEPref(
	newErrPrefix string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}

// ZEPrefCtx - Adds an error prefix and an error context string
// to the list of previous error prefix/context information.
//
// This method is identical in function to ErrPrefixDto.SetEPrefCtx().
// The only difference is that this method returns an ErrPrefixDto
// instance by value.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// This method returns a deep copy of the current ErrPrefixDto
// instance after it is updated with new error prefix and error
// context information. This copy of ErrPrefixDto is returned by
// value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This method is designed to process a single new error prefix
//       string. To process a collection of error prefix strings, see
//       method 'ErrPrefixDto.SetEPrefOld()'.
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
//  ErrPrefixDto
//     - This method returns a deep copy of the current
//       ErrPrefixDto instance after it is updated with new error
//       prefix and error context information. This copy of
//       ErrPrefixDto is returned by value.
//
func (ePrefDto *ErrPrefixDto) ZEPrefCtx(
	newErrPrefix string,
	newErrContext string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefixDtoAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}

// ZEPrefOld - Deletes the current error prefix information and
// initializes this ErrPrefixDto instance with an old or preexisting
// error prefix string. This error prefix string input parameter
// typically includes one or more error prefix elements and may
// also include associated error context elements. This string will
// be parsed and into individual error prefix and error context
// components and stored internally in the current ErrPrefixDto
// instance.
//
// This method is identical in function to ErrPrefixDto.SetEPrefOld().
// The only difference is that this method returns an ErrPrefixDto
// instance by value.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// This method returns a deep copy of the current ErrPrefixDto
// instance after it is updated with preexisting error prefix
// information. This copy of ErrPrefixDto is returned by value.
//
// IMPORTANT
// All existing error prefix and error context information in this
// ErrPrefixDto instance will be overwritten and deleted.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPrefix        string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components and stored in the current
//       ErrPrefixDto instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method returns a deep copy of the current
//       ErrPrefixDto instance after it is updated with preexisting
//       error prefix information. This copy of ErrPrefixDto is
//       returned by value.
//
func (ePrefDto *ErrPrefixDto) ZEPrefOld(
	oldErrPrefix string) ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	_ = errPrefixDtoQuark{}.ptr().emptyErrPrefInfoCollection(
		ePrefDto,
		"")

	ePrefAtom := errPrefixDtoAtom{}

	ePrefDto.inputStrDelimiters.SetToDefaultIfEmpty()

	ePrefDto.outputStrDelimiters.SetToDefaultIfEmpty()

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		ePrefDto.inputStrDelimiters,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefixDtoAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}
