package errpref

import "sync"

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
	isLastLineTerminatedWithNewLine bool
	maxErrPrefixTextLineLength      uint
	lock                            *sync.Mutex
}

// AddEPrefCollectionStr - Receives a string containing one or more
// error prefix and error context pairs. This error prefix
// information is parsed and added to the internal store of
// error prefix elements.
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	if ePrefDto.maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	previousCollectionLen := len(ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().getEPrefContextArray(
		errorPrefixCollectionStr,
		&ePrefDto.ePrefCol)

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
		_ := errPrefAtom{}.ptr().
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
		_ := errPrefAtom{}.ptr().
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
//       of 'eMsg'.
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

	return errPrefAtom{}.ptr().
		copyInErrPrefDto(
			ePrefDto,
			inComingErrPrefixDto,
			eMsg)
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

	return errPrefAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			eMsg)
}

// EmptyEPrefCollection - All instances of ErrPrefixDto store
// error prefix information internally in an array of
// ErrorPrefixInfo objects. This method will delete or empty the
// current ErrorPrefixInfo array and initialize it to a zero length
// array.
//
// Effectively, this will delete all existing error prefix
// information stored in the current ErrPrefixDto instance.
//
//
func (ePrefDto *ErrPrefixDto) EmptyEPrefCollection() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 256)
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

	return errPrefAtom{}.ptr().areEqualErrPrefDtos(
		ePrefDto,
		ePrefixDto2)
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 256)
		newErrorPrefixCollection = make([]ErrorPrefixInfo, 0, 256)
		return newErrorPrefixCollection
	}

	lenEPrefCol := len(ePrefDto.ePrefCol)

	if lenEPrefCol == 0 {
		newErrorPrefixCollection = make([]ErrorPrefixInfo, 0, 256)
		return newErrorPrefixCollection
	}

	newErrorPrefixCollection =
		make([]ErrorPrefixInfo,
			lenEPrefCol,
			lenEPrefCol+256)

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
			make([]ErrorPrefixInfo, 0, 256)
	}

	return len(ePrefDto.ePrefCol)
}

// GetMaxErrPrefTextLineLength - Returns the maximum limit on the
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
func (ePrefDto *ErrPrefixDto) GetMaxErrPrefTextLineLength() uint {

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

// MergeErrPrefixDto - Receives a pointer to another ErrPrefixDto
// object and proceeds to add the error prefix collection to that
// of the current ErrPrefixDto instance.
//
// This means that the error prefix collection for
// 'incomingErrPrefixDto' will be added to the end of the
// ePrefDto collection.
//
func (ePrefDto *ErrPrefixDto) MergeErrPrefixDto(
	incomingErrPrefixDto *ErrPrefixDto) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol =
			make([]ErrorPrefixInfo, 0, 256)
	}

	if incomingErrPrefixDto.ePrefCol == nil {
		incomingErrPrefixDto.ePrefCol =
			make([]ErrorPrefixInfo, 0, 256)
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

}

// New - Returns a new and properly initialized instance of
// ErrPrefixDto.
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
//       instance of ErrPrefixDto.
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

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	return newErrPrefixDto
}

// NewEPrefOld - Returns a new and properly initialized instance of
// ErrPrefixDto. The returned ErrPrefixDto instance will be
// configured with error prefix information extracted from input
// parameter, 'oldErrPrefix'.
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

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	ePrefAtom := errPrefAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
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

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	previousCollectionLen := len(ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().getEPrefContextArray(
		errorPrefixCollection,
		&ePrefDto.ePrefCol)

	numberOfCollectionItemsParsed =
		len(ePrefDto.ePrefCol) -
			previousCollectionLen

	return numberOfCollectionItemsParsed, newErrPrefixDto
}

// NewFromIErrorPrefix - Receives an object which implements the
// IErrorPrefix interface and returns a new, populated instance of
// ErrPrefixDto.
//
// If the IErrorPrefix ErrorPrefixInfo collection is empty, this
// method will return an empty instance of ErrPrefixDto.
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

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	var oldErrPrefStr string

	oldErrPrefStr = iEPref.String()

	if len(oldErrPrefStr) == 0 {
		return newErrPrefixDto
	}

	newErrPrefixDto.SetEPrefOld(oldErrPrefStr)

	return newErrPrefixDto
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

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
		return
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return
	}

	ePrefNanobot := errPrefNanobot{}

	if len(newErrContext) == 0 {

		ePrefNanobot.deleteLastErrContext(ePrefDto)

		return
	}

	ePrefNanobot.setLastCtx(
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 10)
		return
	}

	collectionIdx := len(ePrefDto.ePrefCol)

	if collectionIdx == 0 {
		return
	}

	errPrefNanobot{}.ptr().deleteLastErrContext(
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
// passed in input parameter 'ErrPrefixDto'. If this string
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return
}

// SetEPrefCollection - Deletes the current error prefix collection
// and replaces with a new collection passed as an input parameter
// to this method.
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

	if newEPrefCollection == nil {
		newEPrefCollection = make([]ErrorPrefixInfo, 0, 256)
		return
	}

	lenNewEPrefCol := len(newEPrefCollection)

	ePrefDto.ePrefCol = make(
		[]ErrorPrefixInfo,
		lenNewEPrefCol,
		lenNewEPrefCol+256)

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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
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

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	ePrefAtom := errPrefAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return
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
//       (1,000,000), this method will take no action and exit.
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

	if maxErrPrefixTextLineLength < 10 ||
		maxErrPrefixTextLineLength > 1000000 {
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

// String - Returns a formatted error prefix/context string
// incorporating all error prefixes previously added to this
// instance.
//
// Error prefix information is stored internally in the 'ePrefCol'
// array.
//
func (ePrefDto ErrPrefixDto) String() string {

	if ePrefDto.lock == nil {
		return ""
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		return ""
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return ""
	}

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			ePrefDto.maxErrPrefixTextLineLength,
			ePrefDto.isLastLineTerminatedWithNewLine,
			ePrefDto.ePrefCol)

}

// StrMaxLineLen - Returns a formatted error prefix/context string
// incorporating all error prefix and error context information
// previously added to this ErrPrefixDto instance.
//
// Input parameter 'maxLineLen' is used to set the maximum line
// length for text returned by this method. It does not alter the
// maximum line length default value for this ErrPrefixDto instance.
//
// Error prefix information is stored internally in the 'ePrefCol'
// array.
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
		maxErrPrefixTextLineLength =
			ePrefDto.maxErrPrefixTextLineLength
	}

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return ""
	}

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			maxErrPrefixTextLineLength,
			ePrefDto.isLastLineTerminatedWithNewLine,
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
		return ePrefDto
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return ePrefDto
	}

	ePrefNanobot := errPrefNanobot{}

	if len(newErrContext) == 0 {

		ePrefNanobot.deleteLastErrContext(ePrefDto)

		return ePrefDto
	}

	ePrefNanobot.setLastCtx(
		newErrContext,
		ePrefDto.ePrefCol)

	return ePrefDto
}

// XCtxEmpty - Deletes the last error context for the last error
// prefix in this instance of ErrPrefixDto
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 10)
		return ePrefDto
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return ePrefDto
	}

	errPrefNanobot{}.ptr().deleteLastErrContext(
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
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
//     - Returns a pointer to the current ErrPrefixDto instance
//
func (ePrefDto *ErrPrefixDto) XEPrefCtx(
	newErrPrefix string,
	newErrContext string) *ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
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

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	ePrefAtom := errPrefAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	return ePrefDto
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	ePrefNanobot := errPrefNanobot{}

	if len(ePrefDto.ePrefCol) == 0 {

	} else if len(newErrContext) == 0 {

		ePrefNanobot.deleteLastErrContext(ePrefDto)

	} else {

		ePrefNanobot.setLastCtx(
			newErrContext,
			ePrefDto.ePrefCol)

	}

	newErrPrefixDto,
		_ = errPrefAtom{}.ptr().
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 10)
	}

	if len(ePrefDto.ePrefCol) > 0 {

		errPrefNanobot{}.ptr().deleteLastErrContext(
			ePrefDto)

	}

	newErrPrefixDto,
		_ := errPrefAtom{}.ptr().
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefAtom{}.ptr().
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

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		&ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefAtom{}.ptr().
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

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	ePrefAtom := errPrefAtom{}

	ePrefAtom.getEPrefContextArray(
		oldErrPrefix,
		&ePrefDto.ePrefCol)

	ePrefAtom.setFlagsErrorPrefixInfoArray(
		ePrefDto.ePrefCol)

	newErrPrefixDto,
		_ := errPrefAtom{}.ptr().
		copyOutErrPrefDto(
			ePrefDto,
			"")

	return newErrPrefixDto
}
