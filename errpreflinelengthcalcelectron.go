package errpref

import (
	"fmt"
	"sync"
)

type ePrefixLineLenCalcElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance of
// EPrefixLineLenCalc ('incomingLineLenCalc') to the data fields of
// the target EPrefixLineLenCalc instance ('targetLineLenCalc')
//
// If 'incomingLineLenCalc' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in the 'targetLineLenCalc' instance will
// be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetLineLenCalc          *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. This
//       method WILL CHANGE AND OVERWRITE all values of internal
//       member variables encapsulated in this object to achieve
//       the method's objectives.
//
//       This EPrefixLineLenCalc instance will receive all the
//       data values contained from input parameter
//       'incomingLineLenCalc'. When the copy operation is
//       completed, the data values in 'targetLineLenCalc' will be
//       identical to those contained in input parameter,
//       'incomingLineLenCalc'.
//
//
//  incomingLineLenCalc        *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this EPrefixLineLenCalc instance will be
//       copied to input parameter 'targetLineLenCalc'.
//
//       If this EPrefixLineLenCalc instance proves to be invalid, an
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
func (ePrefLineLenCalcElectron *ePrefixLineLenCalcElectron) copyIn(
	targetLineLenCalc *EPrefixLineLenCalc,
	incomingLineLenCalc *EPrefixLineLenCalc,
	ePrefix string) error {

	if ePrefLineLenCalcElectron.lock == nil {
		ePrefLineLenCalcElectron.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcElectron.lock.Lock()

	defer ePrefLineLenCalcElectron.lock.Unlock()

	ePrefix += "ePrefixLineLenCalcElectron.copyIn() "

	if targetLineLenCalc == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetLineLenCalc' is INVALID!\n"+
			"'targetLineLenCalc' is a nil pointer!\n",
			ePrefix)

	}

	if incomingLineLenCalc == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'incomingLineLenCalc' is INVALID!\n"+
			"'incomingLineLenCalc' is a nil pointer!\n",
			ePrefix)
	}

	_,
		err := ePrefixLineLenCalcQuark{}.ptr().
		testValidityOfEPrefixLineLenCalc(
			incomingLineLenCalc,
			ePrefix+
				"Testing validity of 'incomingLineLenCalc'\n")

	if err != nil {
		return err
	}

	err =
		targetLineLenCalc.ePrefDelimiters.CopyIn(
			&incomingLineLenCalc.ePrefDelimiters,
			ePrefix+
				"incomingLineLenCalc->targetLineLenCalc\n")

	if err != nil {
		return err
	}

	err =
		targetLineLenCalc.errorPrefixInfo.CopyIn(
			&incomingLineLenCalc.errorPrefixInfo,
			ePrefix+
				"incomingLineLenCalc->targetLineLenCalc\n")

	if err != nil {
		return err
	}

	targetLineLenCalc.currentLineStr =
		incomingLineLenCalc.currentLineStr

	targetLineLenCalc.maxErrStringLength =
		incomingLineLenCalc.maxErrStringLength

	return nil
}

// copyOut - Creates a deep copy of the data fields contained in
// input parameter 'lineLenCalc' and returns that data as a new
// instance EPrefixLineLenCalc.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  lineLenCalc         *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. This method
//       will NOT change the values of internal member variables
//       contained in this object.
//
//       If this EPrefixLineLenCalc instance proves to be invalid, an
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
//  EPrefixLineLenCalc
//     - If this method completes successfully, a deep copy of
//       input parameter 'lineLenCalc' will be returned.
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
func (ePrefLineLenCalcElectron *ePrefixLineLenCalcElectron) copyOut(
	lineLenCalc *EPrefixLineLenCalc,
	ePrefix string) (
	EPrefixLineLenCalc,
	error) {

	if ePrefLineLenCalcElectron.lock == nil {
		ePrefLineLenCalcElectron.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcElectron.lock.Lock()

	defer ePrefLineLenCalcElectron.lock.Unlock()

	ePrefix += "ePrefixLineLenCalcElectron.copyOut() "

	newLineLenCalc := EPrefixLineLenCalc{}

	if lineLenCalc == nil {
		return newLineLenCalc,
			fmt.Errorf("%v\n"+
				"\nInput parameter 'lineLenCalc' is INVALID!\n"+
				"'lineLenCalc' is a nil pointer!\n",
				ePrefix)

	}

	_,
		err := ePrefixLineLenCalcQuark{}.ptr().
		testValidityOfEPrefixLineLenCalc(
			lineLenCalc,
			ePrefix+
				"Testing validity of 'lineLenCalc'\n")

	if err != nil {
		return newLineLenCalc, err
	}

	err =
		newLineLenCalc.ePrefDelimiters.CopyIn(
			&lineLenCalc.ePrefDelimiters,
			ePrefix+
				"lineLenCalc->newLineLenCalc\n")

	if err != nil {
		return newLineLenCalc, err
	}

	err =
		newLineLenCalc.errorPrefixInfo.CopyIn(
			&lineLenCalc.errorPrefixInfo,
			ePrefix+
				"lineLenCalc->newLineLenCalc\n")

	if err != nil {
		return newLineLenCalc, err
	}

	newLineLenCalc.currentLineStr =
		lineLenCalc.currentLineStr

	newLineLenCalc.maxErrStringLength =
		lineLenCalc.maxErrStringLength

	return newLineLenCalc, nil
}

// equal - Receives pointers to two EPrefixLineLenCalc objects and
// proceeds to compare the internal data values. If the
// EPrefixLineLenCalc objects contain data values which ARE EQUAL
// in all respects, this method returns 'true'.
//
// If the data values for the two EPrefixLineLenCalc objects ARE
// NOT EQUAL, this method will return 'false'.
//
func (ePrefLineLenCalcElectron *ePrefixLineLenCalcElectron) equal(
	lineLenCalcOne *EPrefixLineLenCalc,
	lineLenCalcTwo *EPrefixLineLenCalc,
	ePrefix string) (
	areEqual bool,
	err error) {

	if ePrefLineLenCalcElectron.lock == nil {
		ePrefLineLenCalcElectron.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcElectron.lock.Lock()

	defer ePrefLineLenCalcElectron.lock.Unlock()

	ePrefix += "ePrefixLineLenCalcElectron.equal() "

	areEqual = false

	if lineLenCalcOne == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'lineLenCalcOne' is INVALID!\n"+
			"'lineLenCalcOne' is a nil pointer!\n",
			ePrefix)

		return areEqual, err
	}

	if lineLenCalcTwo == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'lineLenCalcTwo' is INVALID!\n"+
			"'lineLenCalcTwo' is a nil pointer!\n",
			ePrefix)

		return areEqual, err
	}

	if !lineLenCalcOne.ePrefDelimiters.Equal(
		&lineLenCalcTwo.ePrefDelimiters) {
		return areEqual, err
	}

	if !lineLenCalcOne.errorPrefixInfo.Equal(
		&lineLenCalcTwo.errorPrefixInfo) {
		return areEqual, err
	}

	if lineLenCalcOne.currentLineStr !=
		lineLenCalcTwo.currentLineStr {
		return areEqual, err
	}

	if lineLenCalcOne.maxErrStringLength !=
		lineLenCalcTwo.maxErrStringLength {
		return areEqual, err
	}

	areEqual = true

	return areEqual, err
}

// ptr() - Returns a pointer to a new instance of
// ePrefixLineLenCalcElectron.
//
func (ePrefLineLenCalcElectron ePrefixLineLenCalcElectron) ptr() *ePrefixLineLenCalcElectron {

	if ePrefLineLenCalcElectron.lock == nil {
		ePrefLineLenCalcElectron.lock = new(sync.Mutex)
	}

	ePrefLineLenCalcElectron.lock.Lock()

	defer ePrefLineLenCalcElectron.lock.Unlock()

	return &ePrefixLineLenCalcElectron{
		lock: new(sync.Mutex),
	}
}
