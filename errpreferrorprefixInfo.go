package errpref

import (
	"sync"
)

type ErrorPrefixInfo struct {
	isFirstIdx             bool // Signals the first index in the array
	isLastIdx              bool // Signals the last index in the array
	isPopulated            bool // Signals whether this object contains valid data
	errorPrefixStr         string
	lenErrorPrefixStr      uint
	errPrefixHasContextStr bool
	errorContextStr        string
	lenErrorContextStr     uint
	lock                   *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// ErrorPrefixInfo ('inComingErrPrefixInfo') to the data fields of
// the current ErrorPrefixInfo instance ('errorPrefixInfo').
//
// If 'inComingErrPrefixInfo' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in current ErrorPrefixInfo instance
// ('errorPrefixInfo') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  inComingErrPrefixInfo      *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrorPrefixInfo instance will be
//       copied to current ErrorPrefixInfo instance
//       ('errorPrefixInfo').
//
//       If this ErrorPrefixInfo instance proves to be invalid, an
//       error will be returned.
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
func (errorPrefixInfo *ErrorPrefixInfo) CopyIn(
	inComingErrPrefixInfo *ErrorPrefixInfo,
	ePrefix string) error {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()
	ePrefix += "ErrorPrefixInfo.CopyIn()\n"

	ePrefDtoElectron := errorPrefixInfoElectron{}

	return ePrefDtoElectron.copyIn(
		errorPrefixInfo,
		inComingErrPrefixInfo,
		ePrefix)
}

// CopyOut - Creates a deep copy of the data fields contained in
// the current ErrorPrefixInfo instance, and returns that data as a
// new instance of ErrorPrefixInfo.
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
//  ErrorPrefixInfo
//     - If this method completes successfully, a deep copy of the
//       current ErrorPrefixInfo instance will be returned through
//       this parameter as a completely new instance of
//       ErrorPrefixInfo.
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
func (errorPrefixInfo *ErrorPrefixInfo) CopyOut(
	ePrefix string) (
	ErrorPrefixInfo,
	error) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	ePrefix += "ErrorPrefixInfo.CopyOut() "

	ePrefDtoElectron := errorPrefixInfoElectron{}

	return ePrefDtoElectron.copyOut(
		errorPrefixInfo,
		ePrefix)
}

// Equal - Returns a boolean flag signaling whether the data values
// contained in the current ErrorPrefixInfo instance are equal to
// to those contained in input parameter, 'ePrefixInfo02'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixInfo02       *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. The data
//       values contained in this instance will be compared to
//       those contained in the current ErrorPrefixInfo instance
//       (errorPrefixInfo) to determine equality.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - A boolean flag signaling whether the data values contained
//       in the current ErrorPrefixInfo instance are equal to those
//       contained in input parameter 'ePrefixInfo02'. If the data
//       values are equal in all respects, this returned boolean
//       value will be set to 'true'.
//
func (errorPrefixInfo *ErrorPrefixInfo) Equal(
	ePrefixInfo02 *ErrorPrefixInfo) bool {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfoElectron{}.ptr().equal(
		errorPrefixInfo,
		ePrefixInfo02)
}

// GetErrContextStr - Returns the Error Context String.
func (errorPrefixInfo *ErrorPrefixInfo) GetErrContextStr() string {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.errorContextStr
}

// GetErrPrefixStr - Returns the Error Prefix String.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetErrPrefixStr() string {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.errorPrefixStr
}

// GetErrPrefixHasContextStr - Returns a boolean flag signaling
// whether the Error Prefix String defined in this ErrorPrefixInfo
// instance has an associated Error Context string.
//
// If the returned boolean value is 'true' is signals that this
// Error Prefix has an associated Error Context String.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetErrPrefixHasContextStr() bool {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.errPrefixHasContextStr
}

// GetIsFirstIndex - Returns a boolean flag signaling whether the
// current ErrorPrefixInfo instance is the First element in an array.
//
// If the returned boolean value is 'true', it signals that this
// ErrorPrefixInfo object is the First element in an array.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetIsFirstIndex() bool {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.isFirstIdx
}

// GetIsLastIndex - Returns a boolean flag signaling whether the
// current ErrorPrefixInfo instance is the last element in an array.
//
// If the returned boolean value is 'true', it signals that this
// ErrorPrefixInfo object is the last element in an array.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetIsLastIndex() bool {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.isLastIdx
}

// GetIsPopulated - Returns a boolean flag signaling whether the
// current ErrorPrefixInfo instance is populated with valid data.
//
// If the returned boolean value is 'true', it signals that this
// ErrorPrefixInfo object IS POPULATED with valid data.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetIsPopulated() bool {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.lenErrorPrefixStr =
		uint(len(errorPrefixInfo.errorPrefixStr))

	if errorPrefixInfo.lenErrorPrefixStr > 0 {
		errorPrefixInfo.isPopulated = true
	} else {
		errorPrefixInfo.isPopulated = false
	}

	return errorPrefixInfo.isPopulated
}

// IsValidInstance - Returns a boolean flag signaling whether the
// current ErrorPrefixInfo instance is invalid.
//
// If the returned boolean value is 'true', it signals that the
// current ErrorPrefixInfo instance is valid and populated with data.
//
// If the boolean value is 'false', it signals that the current
// ErrorPrefixInfo instance is invalid.
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
//  isValid             bool
//     - If the current ErrorPrefixInfo instance, 'errorPrefixInfo',
//       contains invalid data, this boolean value is set to 'false'.
//
//       If the current ErrorPrefixInfo instance, 'errorPrefixInfo',
//       is valid, this boolean value is set to 'true'.
//
func (errorPrefixInfo *ErrorPrefixInfo) IsValidInstance() (
	isValid bool) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	ePrefDtoQuark := errorPrefixInfoQuark{}

	isValid,
		_ = ePrefDtoQuark.testValidityOfErrorPrefixInfo(
		errorPrefixInfo,
		"")

	return isValid
}

// IsValidInstanceError - Returns an error object signaling whether
// the current ErrorPrefixInfo instance is invalid.
//
// If the returned error value is 'nil', it signals that the
// current ErrorPrefixInfo instance is valid and populated with
// data.
//
// If the returned error value is NOT 'nil', it signals that the
// current ErrorPrefixInfo is invalid. In this case, the error
// object will contain an appropriate error message.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current ErrorPrefixInfo instance, 'errorPrefixInfo',
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current ErrorPrefixInfo instance, 'errorPrefixInfo',
//       is valid, this error parameter will be set to 'nil'.
//
func (errorPrefixInfo *ErrorPrefixInfo) IsValidInstanceError(
	ePrefix string) error {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	ePrefix += "ErrorPrefixInfo.IsValidInstanceError() "

	ePrefDtoQuark := errorPrefixInfoQuark{}

	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixInfo(
		errorPrefixInfo,
		ePrefix)

	return err
}

// GetLengthErrContextStr - Returns the number of characters
// in the Error Context String. This is also known as the string
// length and it is returned as a type unsigned integer.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetLengthErrContextStr() uint {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.lenErrorContextStr
}

// GetLengthErrPrefixStr - Returns the number of characters in
// the Error Prefix String. This also known as the string length
// and it is returned as a type unsigned integer.
//
func (errorPrefixInfo *ErrorPrefixInfo) GetLengthErrPrefixStr() uint {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return errorPrefixInfo.lenErrorPrefixStr
}

// New - Creates and returns a new instance of type,
// 'ErrorPrefixInfo'.
//
func (errorPrefixInfo ErrorPrefixInfo) New() ErrorPrefixInfo {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	return ErrorPrefixInfo{}
}

// Ptr - Returns a pointer to a new instance of ErrorPrefixInfo.
//
func (errorPrefixInfo ErrorPrefixInfo) Ptr() *ErrorPrefixInfo {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	newErrPrefixInfo := ErrorPrefixInfo{}

	return &newErrPrefixInfo
}

// SetErrContextStr - Sets the Error Context String value.
//
// In addition, this method also calculates and set the value of
// length or number of characters contained in the Error Context
// String.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetErrContextStr(
	errContext string) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.errorContextStr = errContext

	errorPrefixInfo.lenErrorContextStr =
		uint(len(errorPrefixInfo.errorContextStr))

	if errorPrefixInfo.lenErrorContextStr > 0 {
		errorPrefixInfo.errPrefixHasContextStr = true
	} else {
		errorPrefixInfo.errPrefixHasContextStr = false
	}

}

// SetErrPrefixHasContext - Sets an internal boolean flag
// specifying whether the Error Prefix String for this
// ErrorPrefixInfo instance has an associated Error Context String.
//
// If this boolean flag is set to 'true' is signals that this
// Error Prefix String has an associated Error Context String.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetErrPrefixHasContext(
	errPrefixHasContextStr bool) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.errPrefixHasContextStr = errPrefixHasContextStr
}

// SetErrPrefixStr - Sets the value of the error prefix string.
//
// In addition, this method also calculates and set the value of
// length or number of characters contained in the error prefix
// string.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetErrPrefixStr(
	errPrefix string) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.errorPrefixStr = errPrefix

	errorPrefixInfo.lenErrorPrefixStr =
		uint(len(errorPrefixInfo.errorPrefixStr))

	if errorPrefixInfo.lenErrorPrefixStr > 0 {
		errorPrefixInfo.isPopulated = true
	} else {
		errorPrefixInfo.isPopulated = false
	}

}

// SetIsFirstIndex - Sets an internal flag designating whether the
// current ErrorPrefixInfo instance is the first element in an array
// of ErrorPrefixInfo objects.
//
// If the input parameter 'isFirstIndex' is 'true', it signals that
// this ErrorPrefixInfo object is the FIRST element in an array.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetIsFirstIndex(isFirstIndex bool) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.isFirstIdx = isFirstIndex

}

// SetIsLastIndex - Sets an internal flag designating whether the
// current ErrorPrefixInfo instance is the last element in an array
// of ErrorPrefixInfo objects.
//
// If the input parameter 'isLastIdx' is 'true', it signals that
// this ErrorPrefixInfo object is the LAST element in an array.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetIsLastIndex(isLastIdx bool) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.isLastIdx = isLastIdx
}

// SetIsPopulated - Sets an internal flag indicating whether the
// current ErrorPrefixInfo instance is, or is NOT, populated with
// valid data.
//
// If input parameter 'isPopulated' is set to 'true' it signals
// this instance of ErrorPrefixInfo is populated with valid data.
//
func (errorPrefixInfo *ErrorPrefixInfo) SetIsPopulated(
	isPopulated bool) {

	if errorPrefixInfo.lock == nil {
		errorPrefixInfo.lock = new(sync.Mutex)
	}

	errorPrefixInfo.lock.Lock()

	defer errorPrefixInfo.lock.Unlock()

	errorPrefixInfo.isPopulated = isPopulated

}
