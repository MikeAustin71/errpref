package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefixDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of errPrefixDtoMechanics.
//
func (ePrefDtoMech errPrefixDtoMechanics) ptr() *errPrefixDtoMechanics {

	if ePrefDtoMech.lock == nil {
		ePrefDtoMech.lock = new(sync.Mutex)
	}

	ePrefDtoMech.lock.Lock()

	defer ePrefDtoMech.lock.Unlock()

	return &errPrefixDtoMechanics{
		lock: new(sync.Mutex),
	}
}

// get2dEPrefStrings - Receives a pointer to an instance of
// ErrPrefixDto and proceeds to convert the internal collection of
// ErrorPrefixInfo objects to a two dimensional string array
// which is returned to the caller.
//
// The two-dimensional string array contains both error prefix and
// error context information. The Error Prefix string is always in
// the [x][0] position. The Error Context string is always in the
// [x][1] position. The Error Context string is optional and may be
// an empty string.
//
func (ePrefDtoMech *errPrefixDtoMechanics) get2dEPrefStrings(
	errPrefDto *ErrPrefixDto,
	errorPrefStr string) (
	twoDStrArray [][2]string,
	err error) {

	if ePrefDtoMech.lock == nil {
		ePrefDtoMech.lock = new(sync.Mutex)
	}

	ePrefDtoMech.lock.Lock()

	defer ePrefDtoMech.lock.Unlock()

	methodNames := errorPrefStr + "\n" +
		"errPrefixDtoMechanics.get2dEPrefStrings()"

	twoDStrArray = nil

	if errPrefDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodNames)

		return twoDStrArray, err
	}

	colLen := len(errPrefDto.ePrefCol)

	if colLen == 0 {
		return twoDStrArray, err
	}

	twoDStrArray = make([][2]string, colLen)

	for i := 0; i < colLen; i++ {

		twoDStrArray[i][0] =
			errPrefDto.ePrefCol[i].errorPrefixStr

		twoDStrArray[i][1] =
			errPrefDto.ePrefCol[i].errorContextStr
	}

	return twoDStrArray, err
}

// setFromIBasicErrorPrefix - Receives an ErrPrefixDto object and
// an empty interface. If that empty interface is convertible to
// one of the 10-valid types listed below, this method then proceeds
// to overwrite and reset the ErrPrefixDto object's ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from input
// parameter, 'iEPref'. 'iEPref' is an empty interface which must
// be convertible to one of the following valid types:
//
//
//   1. nil               - A nil value is valid and generates an empty
//                          collection of error prefix and error context
//                          information.
//
//   2. Stringer          - The Stringer interface from the 'fmt' package.
//                          This interface has only one method:
//                            type Stringer interface {
//                               String() string
//                            }
//
//   3. string            - A string containing error prefix information.
//
//   4. []string          - A one-dimensional slice of strings containing
//                          error prefix information
//
//   5. [][2]string       - A two-dimensional slice of strings
//                          containing error prefix and error context
//                          information.
//
//   6. strings.Builder   - An instance of strings.Builder. Error prefix
//                          information will be imported into the new
//                          returned instance of ErrPrefixDto.
//
//   7  *strings.Builder  - A pointer to an instance of strings.Builder.
//                          Error prefix information will be imported into
//                          the new returned instance of ErrPrefixDto.
//
//   8. ErrPrefixDto      - An instance of ErrPrefixDto. The
//                          ErrorPrefixInfo from this object will be
//                          copied to the new returned instance of
//                          ErrPrefixDto.
//
//   9. *ErrPrefixDto     - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                          copied to 'errPrefDto'.
//
//  10. IBasicErrorPrefix - An interface to a method generating
//                          a two-dimensional slice of strings
//                          containing error prefix and error
//                          context information.
//
// Any types not listed above will be considered invalid and
// trigger the return of an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              interface{}
//     - An empty interface containing one of the 7-valid type
//       listed above. Any one of these valid types will generate
//       valid error prefix and error context information.
//
//       If 'iEPref' is NOT convertible to one of the 7-valid types
//       listed above, this method will return an error.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to an empty string: "".
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
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoMech *errPrefixDtoMechanics) setFromEmptyInterface(
	errPrefDto *ErrPrefixDto,
	iEPref interface{},
	errorPrefStr string) error {

	if ePrefDtoMech.lock == nil {
		ePrefDtoMech.lock = new(sync.Mutex)
	}

	ePrefDtoMech.lock.Lock()

	defer ePrefDtoMech.lock.Unlock()

	methodNames := errorPrefStr + "\n" +
		"errPrefixDtoMechanics.setFromEmptyInterface()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodNames)
	}

	ePrefDtoQuark := errPrefixDtoQuark{}

	ePrefDtoQuark.normalizeErrPrefixDto(errPrefDto)

	errNanobot := errPrefixDtoNanobot{}

	ePrfAtom := errPrefixDtoAtom{}

	var backup ErrPrefixDto
	var err2 error
	originalDataIsGood := true

	backup,
		err2 = ePrfAtom.copyOutErrPrefDto(
		errPrefDto,
		"")

	if err2 != nil {
		originalDataIsGood = false
	}

	_ = ePrefDtoQuark.
		emptyErrPrefInfoCollection(
			errPrefDto,
			"")

	if iEPref == nil {

		return nil
	}

	var ok bool

	var str string

	// string
	str,
		ok = iEPref.(string)

	if ok {

		return errNanobot.setFromString(
			errPrefDto,
			str,
			methodNames)
	}

	var strs []string

	// string array
	strs,
		ok = iEPref.([]string)

	if ok {

		return errNanobot.setFromStringArray(
			errPrefDto,
			strs,
			methodNames)
	}

	var twoDStrArray [][2]string

	// Two Dimensional String Array
	twoDStrArray,
		ok = iEPref.([][2]string)

	if ok {

		return errNanobot.setFromTwoDStrArray(
			errPrefDto,
			twoDStrArray,
			methodNames)
	}

	var dto ErrPrefixDto

	// ErrPrefixDto
	dto,
		ok = iEPref.(ErrPrefixDto)

	if ok {

		return ePrfAtom.
			copyInErrPrefDto(
				errPrefDto,
				&dto,
				methodNames)
	}

	var dtoPtr *ErrPrefixDto

	// *ErrPrefixDto
	dtoPtr,
		ok = iEPref.(*ErrPrefixDto)

	if ok {

		return ePrfAtom.
			copyInErrPrefDto(
				errPrefDto,
				dtoPtr,
				methodNames)
	}

	var strBuildr strings.Builder

	// strings.Builder
	strBuildr,
		ok = iEPref.(strings.Builder)

	if ok {

		return errNanobot.setFromStringBuilder(
			errPrefDto,
			&strBuildr,
			methodNames)

	}

	var strBuildrPtr *strings.Builder

	// *strings.Builder
	strBuildrPtr,
		ok = iEPref.(*strings.Builder)

	if ok {

		return errNanobot.setFromStringBuilder(
			errPrefDto,
			strBuildrPtr,
			methodNames)

	}

	var iBasicEPref IBasicErrorPrefix

	// IBasicErrorPrefix
	iBasicEPref,
		ok = iEPref.(IBasicErrorPrefix)

	if ok {

		return errNanobot.setFromIBasicErrorPrefix(
			errPrefDto,
			iBasicEPref,
			methodNames)
	}

	var iStr fmt.Stringer

	// fmt.Stringer
	iStr,
		ok = iEPref.(fmt.Stringer)

	if ok {

		if iStr == nil {
			return nil
		}

		return errNanobot.setFromString(
			errPrefDto,
			iStr.String(),
			methodNames)
	}

	// An Error Occurred!
	// An Error Condition NOW EXISTS!
	// Attempt to recover original data
	//
	var errMsg string

	errMsg =
		fmt.Sprintf("%v\n"+
			"Error: Could NOT extract error prefix information\n"+
			"from input parameter 'iEPref'.\n"+
			"'iEPref' IS NOT convertible to one of the seven\n"+
			"supported types listed as follows:\n"+
			"    nil\n"+
			"    fmt.Stringer\n"+
			"    string\n"+
			"    []string\n"+
			"    [][2]string\n"+
			"    strings.Builder"+
			"    *strings.Builder"+
			"    ErrPrefixDto\n"+
			"    *ErrPrefixDto\n"+
			"    IBasicErrorPrefix\n",
			methodNames)

	if originalDataIsGood {
		err2 = ePrfAtom.copyInErrPrefDto(
			errPrefDto,
			&backup,
			"")

		if err2 != nil {
			errMsg +=
				fmt.Sprintf("An additional error occurred while attempting to\n"+
					"recover original data for the 'errPrefDto' parameter. Therefore,\n"+
					"'errPrefDto' now contains corrupted data.\n"+
					"2ndError=\n%v\n",
					err2.Error())
		}

	} else {

		errMsg +=
			"An additional error occurred while attempting to\n" +
				"recover original data for the 'errPrefDto' parameter. Therefore,\n" +
				"'errPrefDto' now contains corrupted data.\n"
	}

	return fmt.Errorf("%v",
		errMsg)
}
