package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDtoElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of errPrefixDtoElectron.
//
func (ePrefDtoElectron errPrefixDtoElectron) ptr() *errPrefixDtoElectron {

	if ePrefDtoElectron.lock == nil {
		ePrefDtoElectron.lock = new(sync.Mutex)
	}

	ePrefDtoElectron.lock.Lock()

	defer ePrefDtoElectron.lock.Unlock()

	return &errPrefixDtoElectron{
		lock: new(sync.Mutex),
	}
}

// emptyErrorPrefixDto - Receives a pointer to an instance of
// ErrPrefixDto and proceeds to set all of the internal member
// variables to their zero value.
//
func (ePrefDtoElectron errPrefixDtoElectron) emptyErrorPrefixDto(
	ePrefixDto *ErrPrefixDto,
	errPrefStr string) error {

	if ePrefDtoElectron.lock == nil {
		ePrefDtoElectron.lock = new(sync.Mutex)
	}

	ePrefDtoElectron.lock.Lock()

	defer ePrefDtoElectron.lock.Unlock()

	errPrefStr = errPrefStr +
		"\nerrPrefixDtoElectron.emptyErrorPrefixDto()"

	if ePrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefixDto' is invalid!\n"+
			"'ePrefixDto' is a 'nil' pointer.\n",
			errPrefStr)
	}

	ePrefixDto.leftMarginLength = 0

	ePrefixDto.leftMarginChar = 0

	ePrefixDto.isLastLineTerminatedWithNewLine = false

	ePrefixDto.turnOffTextDisplay = false

	ePrefixDto.maxErrPrefixTextLineLength = 0

	ePrefixDto.inputStrDelimiters.Empty()

	ePrefixDto.outputStrDelimiters.Empty()

	ePrefixDto.leadingTextStr = ""

	ePrefixDto.trailingTextStr = ""

	return errPrefixDtoQuark{}.ptr().emptyErrPrefInfoCollection(
		ePrefixDto,
		errPrefStr)
}

// newZeroErrPrefixDto - Returns a pointer to a new instance of
// ErrPrefixDto with all internal member variables set to their
// initial or zero values.
//
func (ePrefDtoElectron *errPrefixDtoElectron) newPtrZeroErrPrefixDto() *ErrPrefixDto {

	if ePrefDtoElectron.lock == nil {
		ePrefDtoElectron.lock = new(sync.Mutex)
	}

	ePrefDtoElectron.lock.Lock()

	defer ePrefDtoElectron.lock.Unlock()

	newErrPrefixDto :=
		errPrefixDtoQuark{}.ptr().newZeroErrPrefixDto()

	return &newErrPrefixDto
}
