package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDelimitersMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of errPrefixDelimitersMechanics.
//
func (ePrefDelimitersMech errPrefixDelimitersMechanics) ptr() *errPrefixDelimitersMechanics {

	if ePrefDelimitersMech.lock == nil {
		ePrefDelimitersMech.lock = new(sync.Mutex)
	}

	ePrefDelimitersMech.lock.Lock()

	defer ePrefDelimitersMech.lock.Unlock()

	return &errPrefixDelimitersMechanics{
		lock: new(sync.Mutex),
	}
}

// setErrPrefDelimiters - Receives a pointer to an instance of
// ErrPrefixDelimiters and proceeds to overwrite and reset the data
// values for all internal member variables in that instance.
//
// The new values are generated from string values submitted as
// input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefDelimiters            *ErrPrefixDelimiters
//     - A pointer to an instance of ErrPrefixDelimiters. All of
//       the data value encapsulated by this instance will be
//       deleted, overwritten and replaced with new data values
//       generated from the following input parameters.
//
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
func (ePrefDelimitersMech *errPrefixDelimitersMechanics) setErrPrefDelimiters(
	ePrefDelimiters *ErrPrefixDelimiters,
	newLinePrefixDelimiters string,
	inLinePrefixDelimiters string,
	newLineContextDelimiters string,
	inLineContextDelimiters string,
	ePrefix string) (
	err error) {

	if ePrefDelimitersMech.lock == nil {
		ePrefDelimitersMech.lock = new(sync.Mutex)
	}

	ePrefDelimitersMech.lock.Lock()

	defer ePrefDelimitersMech.lock.Unlock()

	ePrefix += "\nerrPrefixDelimitersMechanics.setErrPrefDelimiters() "

	if ePrefDelimiters == nil {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'ePrefDelimiters' is INVALID!\n"+
			"'ePrefDelimiters' is a nil pointer!\n",
			ePrefix)

		return err
	}

	if len(newLinePrefixDelimiters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'newLinePrefixDelimiter' is INVALID!\n"+
			"'newLinePrefixDelimiter' is an empty string!\n",
			ePrefix)

		return err
	}

	if len(inLinePrefixDelimiters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'inLinePrefixDelimiters' is INVALID!\n"+
			"'inLinePrefixDelimiters' is an empty string!\n",
			ePrefix)

		return err
	}

	if len(newLineContextDelimiters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'newLineContextDelimiter' is INVALID!\n"+
			"'newLineContextDelimiter' is an empty string!\n",
			ePrefix)

		return err
	}

	if len(inLineContextDelimiters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'inLineContextDelimiter' is INVALID!\n"+
			"'inLineContextDelimiter' is an empty string!\n",
			ePrefix)

		return err
	}

	ePrefDelimiters.inLinePrefixDelimiter =
		inLinePrefixDelimiters

	ePrefDelimiters.lenInLinePrefixDelimiter =
		uint(len(ePrefDelimiters.inLinePrefixDelimiter))

	ePrefDelimiters.newLinePrefixDelimiter =
		newLinePrefixDelimiters

	ePrefDelimiters.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelimiters.newLinePrefixDelimiter))

	ePrefDelimiters.inLineContextDelimiter =
		inLineContextDelimiters

	ePrefDelimiters.lenInLineContextDelimiter =
		uint(len(ePrefDelimiters.inLineContextDelimiter))

	ePrefDelimiters.newLineContextDelimiter =
		newLineContextDelimiters

	ePrefDelimiters.lenNewLineContextDelimiter =
		uint(len(ePrefDelimiters.newLineContextDelimiter))

	return err
}

// setToDefault - Receives a pointer to an instance of
// ErrPrefixDelimiters and proceeds to set the data values for this
// instance to those of the system default.
//
func (ePrefDelimitersMech *errPrefixDelimitersMechanics) setToDefault(
	ePrefDelimiters *ErrPrefixDelimiters,
	ePrefix string) (
	err error) {

	if ePrefDelimitersMech.lock == nil {
		ePrefDelimitersMech.lock = new(sync.Mutex)
	}

	ePrefDelimitersMech.lock.Lock()

	defer ePrefDelimitersMech.lock.Unlock()

	ePrefix += "\nerrPrefixDelimitersMechanics.setToDefault() "

	if ePrefDelimiters == nil {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'ePrefDelimiters' is INVALID!\n"+
			"'ePrefDelimiters' is a nil pointer!\n",
			ePrefix)

		return err
	}

	delimiters :=
		errPrefElectron{}.ptr().getDelimiters()

	ePrefDelimiters.inLinePrefixDelimiter =
		delimiters.inLinePrefixDelimiter

	ePrefDelimiters.lenInLinePrefixDelimiter =
		uint(len(ePrefDelimiters.inLinePrefixDelimiter))

	ePrefDelimiters.newLinePrefixDelimiter =
		delimiters.newLinePrefixDelimiter

	ePrefDelimiters.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelimiters.newLinePrefixDelimiter))

	ePrefDelimiters.inLineContextDelimiter =
		delimiters.inLineContextDelimiter

	ePrefDelimiters.lenInLineContextDelimiter =
		uint(len(ePrefDelimiters.inLineContextDelimiter))

	ePrefDelimiters.newLineContextDelimiter =
		delimiters.newLineContextDelimiter

	ePrefDelimiters.lenNewLineContextDelimiter =
		uint(len(ePrefDelimiters.newLineContextDelimiter))

	return err
}
