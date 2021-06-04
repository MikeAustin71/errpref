package errpref

import "sync"

type errPrefPreon struct {
	lock *sync.Mutex
}

// ptr() - Returns a pointer to a new instance of
// errPrefPreon.
//
func (ePrefPreon errPrefPreon) ptr() *errPrefPreon {

	if ePrefPreon.lock == nil {
		ePrefPreon.lock = new(sync.Mutex)
	}

	ePrefPreon.lock.Lock()

	defer ePrefPreon.lock.Unlock()

	return &errPrefPreon{}
}

// getMinErrPrefLineLength - Returns the Minimum Error Prefix Line
// Length in number of characters.
//
// If users attempt to set the value of the Error Prefix Line
// Length to a value less than this minimum, parent functions will
// either return an error or reset the value to the default Error
// Prefix Line Length.
//
func (ePrefPreon *errPrefPreon) getMinErrPrefLineLength() uint {

	if ePrefPreon.lock == nil {
		ePrefPreon.lock = new(sync.Mutex)
	}

	ePrefPreon.lock.Lock()

	defer ePrefPreon.lock.Unlock()

	return 10
}

func (ePrefPreon *errPrefPreon) getDefaultErrPrefLineLength() uint {

	if ePrefPreon.lock == nil {
		ePrefPreon.lock = new(sync.Mutex)
	}

	ePrefPreon.lock.Lock()

	defer ePrefPreon.lock.Unlock()

	return 40
}
