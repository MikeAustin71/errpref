package errpref

import (
	"fmt"
	"sync"
)

type errorPrefixInfoQuark struct {
	lock *sync.Mutex
}

// testValidityOfErrorPrefixInfo - Performs a diagnostic review of
// the input parameter 'errPrefixInfo', an instance of
// ErrorPrefixInfo. The purpose of this diagnostic review is to
// determine whether this ErrorPrefixInfo instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixInfo       *ErrorPrefixInfo
//     - A pointer to an instance of ErrorPrefixInfo. This object
//       will be evaluated to determine whether or not it is a
//       valid instance.
//
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
//  isValid             bool
//     - This returned boolean value will signal whether the input
//       parameter, 'errPrefixInfo', is valid, or not. If the
//       'errPrefixInfo' object contains valid data,  this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
//  err                 error
//     - If the input parameter object, 'errPrefixInfo', contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the input parameter object, 'errPrefixInfo', is valid,
//       this error parameter will be set to 'nil'.
//
func (ePrefDtoQuark *errorPrefixInfoQuark) testValidityOfErrorPrefixInfo(
	errPrefixInfo *ErrorPrefixInfo,
	ePrefix string) (
	isValid bool,
	err error) {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	ePrefix += "errorPrefixInfoQuark.testValidityOfErrorPrefixInfo() "
	isValid = false

	if errPrefixInfo == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'errPrefixInfo' is INVALID!\n"+
			"'errPrefixInfo' is a nil pointer!\n",
			ePrefix)

		return isValid, err
	}

	errPrefixInfo.lenErrorPrefixStr =
		uint(len(errPrefixInfo.errorPrefixStr))

	if errPrefixInfo.lenErrorPrefixStr == 0 {

		err =
			fmt.Errorf("%v\n"+
				"Error: Error Prefix is an empty string!\n",
				ePrefix)
		return isValid, err
	}

	errPrefixInfo.lenErrorContextStr =
		uint(len(errPrefixInfo.errorContextStr))

	if errPrefixInfo.lenErrorContextStr == 0 {
		errPrefixInfo.errPrefixHasContextStr = false
	} else {
		errPrefixInfo.errPrefixHasContextStr = true
	}

	isValid = true

	return isValid, err
}
