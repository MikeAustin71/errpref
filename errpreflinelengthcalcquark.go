package errpref

import (
	"fmt"
	"sync"
)

type ePrefixLineLenCalcQuark struct {
	lock *sync.Mutex
}

// ptr() - Returns a pointer to a new instance of
// ePrefixLineLenCalcQuark.
//
func (ePrefLineLenCalcQuark ePrefixLineLenCalcQuark) ptr() *ePrefixLineLenCalcQuark {

	if ePrefLineLenCalcQuark.lock == nil {
		ePrefLineLenCalcQuark.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcQuark.lock.Lock()

	defer ePrefLineLenCalcQuark.lock.Unlock()

	return &ePrefixLineLenCalcQuark{
		lock: new(sync.Mutex),
	}
}

// testValidityOfEPrefixLineLenCalc - Performs a diagnostic review of
// the input parameter 'delimiters', an instance of
// ErrPrefixDelimiters. The purpose of this diagnostic review is to
// determine whether this ErrPrefixDelimiters instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  delimiters          *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. This object
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
//       parameter, 'delimiters', is valid, or not. If the
//       'errPrefixDto' object contains valid data,  this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
//  err                 error
//     - If the input parameter object, 'errPrefixDto', contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the input parameter object, 'errPrefixDto', is valid,
//       this error parameter will be set to 'nil'.
//
func (ePrefLineLenCalcQuark *ePrefixLineLenCalcQuark) testValidityOfEPrefixLineLenCalc(
	ePrefLineLenCalc *EPrefixLineLenCalc,
	ePrefix string) (
	isValid bool,
	err error) {

	if ePrefLineLenCalcQuark.lock == nil {
		ePrefLineLenCalcQuark.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcQuark.lock.Lock()

	defer ePrefLineLenCalcQuark.lock.Unlock()

	ePrefix += "ePrefixLineLenCalcQuark.testValidityOfEPrefixLineLenCalc() "
	isValid = false

	if ePrefLineLenCalc == nil {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'ePrefLineLenCalc' is INVALID!\n"+
			"'ePrefLineLenCalc' is a nil pointer!\n",
			ePrefix)

		return isValid, err
	}

	err = ePrefLineLenCalc.ePrefDelimiters.IsValidInstanceError(
		ePrefix +
			"ePrefLineLenCalc.ePrefDelimiters validity check\n")

	if err != nil {
		return isValid, err
	}

	err = ePrefLineLenCalc.errorPrefixInfo.IsValidInstanceError(
		ePrefix +
			"Testing validity of 'ePrefLineLenCalc.errorPrefixInfo'\n")

	if err != nil {
		return isValid, err
	}

	if ePrefLineLenCalc.maxErrStringLength == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member variable 'ePrefLineLenCalc.maxErrStringLength'\n"+
			"is invalid!\n"+
			"ePrefLineLenCalc.maxErrStringLength == 0\n",
			ePrefix)

		return isValid, err
	}

	isValid = true

	return isValid, err
}
