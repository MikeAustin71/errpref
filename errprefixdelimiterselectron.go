package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDelimitersElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance of
// ErrPrefixDelimiters ('incomingDelimiters') to the data fields of
// the target ErrPrefixDelimiters instance ('targetDelimiters')
//
// If 'incomingDelimiters' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in the 'targetDelimiters' instance will
// be overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetDelimiters           *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. This
//       method WILL CHANGE AND OVERWRITE all values of internal
//       member variables encapsulated in this object to achieve
//       the method's objectives.
//
//       This ErrPrefixDelimiters instance will receive all the
//       data values contained from input parameter
//       'incomingDelimiters'. When the copy operation is
//       completed, the data values in 'targetDelimiters' will be
//       identical to those contained in input parameter,
//       'incomingDelimiters'.
//
//
//  incomingDelimiters         *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrPrefixDelimiters instance will be
//       copied to input parameter 'targetDelimiters'.
//
//       If this ErrPrefixDelimiters instance proves to be invalid, an
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
func (ePrefDelimsElectron *errPrefixDelimitersElectron) copyIn(
	targetDelimiters *ErrPrefixDelimiters,
	incomingDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDelimsElectron.lock == nil {
		ePrefDelimsElectron.lock = new(sync.Mutex)
	}

	ePrefDelimsElectron.lock.Lock()

	defer ePrefDelimsElectron.lock.Unlock()

	ePrefix += "errorPrefixInfoElectron.copyIn() "

	if targetDelimiters == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetDelimiters' is INVALID!\n"+
			"'targetDelimiters' is a nil pointer!\n",
			ePrefix)

	}

	if incomingDelimiters == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'incomingDelimiters' is INVALID!\n"+
			"'incomingDelimiters' is a nil pointer!\n",
			ePrefix)
	}

	// 'incomingDelimiters' line lengths are re-calculated here.
	_,
		err := errPrefixDelimitersQuark{}.ptr().
		testValidityOfErrPrefixDelimiters(
			incomingDelimiters,
			ePrefix+
				"incomingDelimiters\n")

	if err != nil {
		return err
	}

	targetDelimiters.inLinePrefixDelimiter =
		incomingDelimiters.inLinePrefixDelimiter

	targetDelimiters.lenInLinePrefixDelimiter =
		incomingDelimiters.lenInLinePrefixDelimiter

	targetDelimiters.newLinePrefixDelimiter =
		incomingDelimiters.newLinePrefixDelimiter

	targetDelimiters.lenNewLinePrefixDelimiter =
		incomingDelimiters.lenNewLinePrefixDelimiter

	targetDelimiters.inLineContextDelimiter =
		incomingDelimiters.inLineContextDelimiter

	targetDelimiters.lenInLineContextDelimiter =
		incomingDelimiters.lenInLineContextDelimiter

	targetDelimiters.newLineContextDelimiter =
		incomingDelimiters.newLineContextDelimiter

	targetDelimiters.lenNewLineContextDelimiter =
		incomingDelimiters.lenNewLineContextDelimiter

	return nil
}

// copyOut - Creates a deep copy of the data fields contained in
// input parameter 'delimiters' and returns that data as a new
// instance ErrPrefixDelimiters.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  delimiters          *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. This method
//       will NOT change the values of internal member variables
//       contained in this object.
//
//       If this ErrPrefixDelimiters instance proves to be invalid, an
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
//  ErrPrefixDelimiters
//     - If this method completes successfully, a deep copy of
//       input parameter 'delimiters' will be returned.
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
func (ePrefDelimsElectron *errPrefixDelimitersElectron) copyOut(
	delimiters *ErrPrefixDelimiters,
	ePrefix string) (
	ErrPrefixDelimiters,
	error) {

	if ePrefDelimsElectron.lock == nil {
		ePrefDelimsElectron.lock = new(sync.Mutex)
	}

	ePrefDelimsElectron.lock.Lock()

	defer ePrefDelimsElectron.lock.Unlock()

	ePrefix += "errPrefixDelimitersElectron.copyOut() "

	newDelimiters := ErrPrefixDelimiters{}

	if delimiters == nil {
		return newDelimiters,
			fmt.Errorf("%v\n"+
				"\nInput parameter 'delimiters' is INVALID!\n"+
				"'delimiters' is a nil pointer!\n",
				ePrefix)

	}

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	// 'delimiters' line lengths are re-calculated here.
	_,
		err := ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix+
			"delimiters\n")

	if err != nil {
		return newDelimiters, err
	}

	newDelimiters.inLinePrefixDelimiter =
		delimiters.inLinePrefixDelimiter

	newDelimiters.lenInLinePrefixDelimiter =
		delimiters.lenInLinePrefixDelimiter

	newDelimiters.newLinePrefixDelimiter =
		delimiters.newLinePrefixDelimiter

	newDelimiters.lenNewLinePrefixDelimiter =
		delimiters.lenNewLinePrefixDelimiter

	newDelimiters.inLineContextDelimiter =
		delimiters.inLineContextDelimiter

	newDelimiters.lenInLineContextDelimiter =
		delimiters.lenInLineContextDelimiter

	newDelimiters.newLineContextDelimiter =
		delimiters.newLineContextDelimiter

	newDelimiters.lenNewLineContextDelimiter =
		delimiters.lenNewLineContextDelimiter

	newDelimiters.lock = new(sync.Mutex)

	return newDelimiters, nil
}

// equal - Receives pointers to two instances of
// ErrPrefixDelimiters and proceeds to determine whether the two
// instances have equal data values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  delimitersOne       *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. The
//       internal data values encapsulated by this object will be
//       compared to those contained in input parameter,
//       'delimitersTwo'. If the data values are equal, a boolean
//       flag of 'true' will be returned.
//
//
//  delimitersTwo       *ErrPrefixDelimiters
//     - A pointer to a second instance of ErrPrefixDelimiters. The
//       internal data values encapsulated by this object will be
//       compared to those contained in input parameter,
//       'delimitersOne'. If the data values are equal, a boolean
//       flag of 'true' will be returned.
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
//  areEqual            bool
//     - This method will compare the data values encapsulated by
//       input parameters, 'delimitersOne' and 'delimitersTwo'. If
//       the data values are equivalent, this boolean flag will be
//       set to 'true'.
//
//       If the data values are NOT equivalent, this parameter will
//       be set to 'false'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (ePrefDelimsElectron *errPrefixDelimitersElectron) equal(
	delimitersOne *ErrPrefixDelimiters,
	delimitersTwo *ErrPrefixDelimiters,
	ePrefix string) (
	areEqual bool,
	err error) {

	if ePrefDelimsElectron.lock == nil {
		ePrefDelimsElectron.lock = new(sync.Mutex)
	}

	ePrefDelimsElectron.lock.Lock()

	defer ePrefDelimsElectron.lock.Unlock()

	ePrefix += "errorPrefixInfoElectron.equal() "

	areEqual = false

	if delimitersOne == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'delimitersOne' is INVALID!\n"+
			"'delimitersOne' is a nil pointer!\n",
			ePrefix)

		return areEqual, err
	}

	if delimitersTwo == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'delimitersTwo' is INVALID!\n"+
			"'delimitersTwo' is a nil pointer!\n",
			ePrefix)

		return areEqual, err
	}

	if delimitersOne.inLinePrefixDelimiter !=
		delimitersTwo.inLinePrefixDelimiter {

		return areEqual, err
	}

	if delimitersOne.lenInLinePrefixDelimiter !=
		delimitersTwo.lenInLinePrefixDelimiter {

		return areEqual, err
	}

	if delimitersOne.newLinePrefixDelimiter !=
		delimitersTwo.newLinePrefixDelimiter {

		return areEqual, err
	}

	if delimitersOne.lenNewLinePrefixDelimiter !=
		delimitersTwo.lenNewLinePrefixDelimiter {

		return areEqual, err
	}

	if delimitersOne.inLineContextDelimiter !=
		delimitersTwo.inLineContextDelimiter {

		return areEqual, err
	}

	if delimitersOne.lenInLineContextDelimiter !=
		delimitersTwo.lenInLineContextDelimiter {

		return areEqual, err
	}

	if delimitersOne.newLineContextDelimiter !=
		delimitersTwo.newLineContextDelimiter {

		return areEqual, err
	}

	if delimitersOne.lenNewLineContextDelimiter !=
		delimitersTwo.lenNewLineContextDelimiter {

		return areEqual, err
	}

	areEqual = true

	return areEqual, err
}

// ptr - Returns a pointer to a new instance of errPrefixDelimitersElectron.
//
func (ePrefDelimsElectron errPrefixDelimitersElectron) ptr() *errPrefixDelimitersElectron {

	if ePrefDelimsElectron.lock == nil {
		ePrefDelimsElectron.lock = new(sync.Mutex)
	}

	ePrefDelimsElectron.lock.Lock()

	defer ePrefDelimsElectron.lock.Unlock()

	return &errPrefixDelimitersElectron{
		lock: new(sync.Mutex),
	}
}
