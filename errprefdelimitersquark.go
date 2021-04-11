package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDelimitersQuark struct {
	lock *sync.Mutex
}

// testValidityOfErrorPrefixInfo - Performs a diagnostic review of
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
func (ePrefDelimsQuark *errPrefixDelimitersQuark) testValidityOfErrPrefixDelimiters(
	delimiters *ErrPrefixDelimiters,
	ePrefix string) (
	isValid bool,
	err error) {

	if ePrefDelimsQuark.lock == nil {
		ePrefDelimsQuark.lock = new(sync.Mutex)
	}

	ePrefDelimsQuark.lock.Lock()

	defer ePrefDelimsQuark.lock.Unlock()

	ePrefix += "errPrefixDelimitersQuark.testValidityOfErrPrefixDelimiters() "
	isValid = false

	if delimiters == nil {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'delimiters' is INVALID!\n"+
			"'delimiters' is a nil pointer!\n",
			ePrefix)

		return isValid, err
	}

	delimiters.lenInLinePrefixDelimiter =
		uint(len(delimiters.inLinePrefixDelimiter))

	if delimiters.lenInLinePrefixDelimiter == 0 {
		err =
			fmt.Errorf("%v\n"+
				"Error: The In Line Prefix Delimiter is an "+
				"empty string!\n",
				ePrefix)

		return isValid, err
	}

	delimiters.lenNewLinePrefixDelimiter =
		uint(len(delimiters.newLinePrefixDelimiter))

	if delimiters.lenNewLinePrefixDelimiter == 0 {
		err =
			fmt.Errorf("%v\n"+
				"Error: The In New Prefix Delimiter is an "+
				"empty string!\n",
				ePrefix)

		return isValid, err
	}

	delimiters.lenInLineContextDelimiter =
		uint(len(delimiters.inLineContextDelimiter))

	if delimiters.lenInLineContextDelimiter == 0 {
		err =
			fmt.Errorf("%v\n"+
				"Error: The In Line Context Delimiter is an "+
				"empty string!\n",
				ePrefix)

		return isValid, err
	}

	delimiters.lenNewLineContextDelimiter =
		uint(len(delimiters.newLineContextDelimiter))

	if delimiters.lenNewLineContextDelimiter == 0 {
		err =
			fmt.Errorf("%v\n"+
				"Error: The In New Prefix Delimiter is an "+
				"empty string!\n",
				ePrefix)

		return isValid, err
	}

	isValid = true

	return isValid, err
}
