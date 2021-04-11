package errpref

import (
	"fmt"
	"sync"
)

type errorPrefixInfoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance of
// ErrorPrefixInfo ('inComingErrPrefixDto') to the data fields of
// the target ErrorPrefixInfo instance ('targetErrPrefixDto')
//
// If 'inComingErrPrefixDto' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in the 'targetErrPrefixDto' instance will
// be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetErrPrefixInfo        *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. This method
//       WILL CHANGE AND OVERWRITE all values of internal member
//       variables encapsulated in this object to achieve the
//       method's objectives.
//
//       This ErrorPrefixInfo instance will receive all the data
//       values contained from input parameter
//       'inComingErrPrefixInfo'. When the copy operation is
//       completed, the data values in 'targetErrPrefixInfo' will be
//       identical to those contained in input parameter,
//       'inComingErrPrefixInfo'.
//
//
//  inComingErrPrefixInfo      *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrorPrefixInfo instance will be
//       copied to input parameter 'targetErrPrefixInfo'.
//
//       If this ErrorPrefixInfo instance proves to be invalid, an
//       error will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (ePrefInfoElectron *errorPrefixInfoElectron) copyIn(
	targetErrPrefixInfo *ErrorPrefixInfo,
	inComingErrPrefixInfo *ErrorPrefixInfo,
	ePrefix string) error {

	if ePrefInfoElectron.lock == nil {
		ePrefInfoElectron.lock = new(sync.Mutex)
	}

	ePrefInfoElectron.lock.Lock()

	defer ePrefInfoElectron.lock.Unlock()

	ePrefix += "errorPrefixInfoElectron.copyIn() "

	if targetErrPrefixInfo == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetErrPrefixInfo' is INVALID!\n"+
			"'targetErrPrefixInfo' is a nil pointer!\n",
			ePrefix)

	}

	if inComingErrPrefixInfo == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'inComingErrPrefixInfo' is INVALID!\n"+
			"'inComingErrPrefixInfo' is a nil pointer!\n",
			ePrefix)
	}

	ePrefDtoQuark := errorPrefixInfoQuark{}

	// Re-calculates line lengths here!
	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixInfo(
		inComingErrPrefixInfo,
		ePrefix+
			"inComingErrPrefixInfo\n")

	if err != nil {
		return err
	}

	targetErrPrefixInfo.isFirstIdx =
		inComingErrPrefixInfo.isFirstIdx

	targetErrPrefixInfo.isLastIdx =
		inComingErrPrefixInfo.isLastIdx

	targetErrPrefixInfo.isPopulated =
		inComingErrPrefixInfo.isPopulated

	targetErrPrefixInfo.errorPrefixStr =
		inComingErrPrefixInfo.errorPrefixStr

	targetErrPrefixInfo.lenErrorPrefixStr =
		inComingErrPrefixInfo.lenErrorPrefixStr

	targetErrPrefixInfo.errPrefixHasContextStr =
		inComingErrPrefixInfo.errPrefixHasContextStr

	targetErrPrefixInfo.errorContextStr =
		inComingErrPrefixInfo.errorContextStr

	targetErrPrefixInfo.lenErrorContextStr =
		inComingErrPrefixInfo.lenErrorContextStr

	return nil
}

// copyOut - Creates a deep copy of the data fields contained in
// input parameter 'errPrefixInfo' and returns that data as a new
// instance ErrorPrefixInfo.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixInfo       *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. This method
//       will NOT change the values of internal member variables
//       contained in this object.
//
//       If this ErrorPrefixInfo instance proves to be invalid, an
//       error will be returned.
//
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
//     - If this method completes successfully, a deep copy of
//       input parameter 'errPrefixInfo' will be returned.
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
func (ePrefInfoElectron *errorPrefixInfoElectron) copyOut(
	errPrefixInfo *ErrorPrefixInfo,
	ePrefix string) (
	ErrorPrefixInfo,
	error) {

	if ePrefInfoElectron.lock == nil {
		ePrefInfoElectron.lock = new(sync.Mutex)
	}

	ePrefInfoElectron.lock.Lock()

	defer ePrefInfoElectron.lock.Unlock()

	ePrefix += "errorPrefixInfoElectron.copyOut() "

	newErrPrefixInfo := ErrorPrefixInfo{}

	if errPrefixInfo == nil {
		return newErrPrefixInfo,
			fmt.Errorf("%v\n"+
				"\nInput parameter 'errPrefixInfo' is INVALID!\n"+
				"'errPrefixInfo' is a nil pointer!\n",
				ePrefix)

	}

	ePrefDtoQuark := errorPrefixInfoQuark{}

	// Re-calculates line lengths here!
	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixInfo(
		errPrefixInfo,
		ePrefix+
			"errPrefixInfo\n")

	if err != nil {
		return newErrPrefixInfo, err
	}

	newErrPrefixInfo.lock = new(sync.Mutex)

	newErrPrefixInfo.isFirstIdx =
		errPrefixInfo.isFirstIdx

	newErrPrefixInfo.isLastIdx =
		errPrefixInfo.isLastIdx

	newErrPrefixInfo.isPopulated =
		errPrefixInfo.isPopulated

	newErrPrefixInfo.errorPrefixStr =
		errPrefixInfo.errorPrefixStr

	newErrPrefixInfo.lenErrorPrefixStr =
		errPrefixInfo.lenErrorPrefixStr

	newErrPrefixInfo.errPrefixHasContextStr =
		errPrefixInfo.errPrefixHasContextStr

	newErrPrefixInfo.errorContextStr =
		errPrefixInfo.errorContextStr

	newErrPrefixInfo.lenErrorContextStr =
		errPrefixInfo.lenErrorContextStr

	newErrPrefixInfo.lock = new(sync.Mutex)

	return newErrPrefixInfo, nil
}

// equal - Returns a boolean flag signaling whether the
// data values of two ErrorPrefixInfo objects are equal
// in all respects.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixInfo01       *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. The data
//       values contained in this instance will be compared to
//       those contained in 'ePrefixInfo02' to determine equality.
//
//
//  ePrefixInfo02       *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. The data
//       values contained in this instance will be compared to
//       those contained in 'ePrefixInfo01' to determine equality.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - A boolean flag signaling whether the data values contained
//       in 'ePrefixInfo01' and 'ePrefixInfo02' are equal in all
//       respects. If the data values are equal, this returned
//       boolean value will be set to 'true'.
//
func (ePrefInfoElectron *errorPrefixInfoElectron) equal(
	ePrefixInfo01 *ErrorPrefixInfo,
	ePrefixInfo02 *ErrorPrefixInfo) bool {

	if ePrefInfoElectron.lock == nil {
		ePrefInfoElectron.lock = new(sync.Mutex)
	}

	ePrefInfoElectron.lock.Lock()

	defer ePrefInfoElectron.lock.Unlock()

	if ePrefixInfo01 == nil ||
		ePrefixInfo02 == nil {
		return false
	}

	if ePrefixInfo01.isFirstIdx !=
		ePrefixInfo02.isFirstIdx {
		return false
	}

	if ePrefixInfo01.isLastIdx !=
		ePrefixInfo02.isLastIdx {
		return false
	}

	if ePrefixInfo01.isPopulated !=
		ePrefixInfo02.isPopulated {
		return false
	}

	if ePrefixInfo01.errorPrefixStr !=
		ePrefixInfo02.errorPrefixStr {
		return false
	}

	if ePrefixInfo01.lenErrorPrefixStr !=
		ePrefixInfo02.lenErrorPrefixStr {
		return false
	}

	if ePrefixInfo01.errorContextStr !=
		ePrefixInfo02.errorContextStr {
		return false
	}

	if ePrefixInfo01.lenErrorContextStr !=
		ePrefixInfo02.lenErrorContextStr {
		return false
	}

	if ePrefixInfo01.errPrefixHasContextStr !=
		ePrefixInfo02.errPrefixHasContextStr {
		return false
	}

	return true
}

// ptr() - Returns a pointer to a new instance of
// ePrefixLineLenCalcQuark.
//
func (ePrefInfoElectron errorPrefixInfoElectron) ptr() *errorPrefixInfoElectron {

	if ePrefInfoElectron.lock == nil {
		ePrefInfoElectron.lock = new(sync.Mutex)
	}

	ePrefInfoElectron.lock.Lock()

	defer ePrefInfoElectron.lock.Unlock()

	return &errorPrefixInfoElectron{}
}
