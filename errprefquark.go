package errpref

import (
	"fmt"
	"strings"
	"sync"
)

// constDefaultErrPrefLineLength - Do NOT access
// constDefaultErrPrefLineLength without first locking
// 'defaultErrPrefLineLenLock'
const constDefaultErrPrefLineLength = uint(40)

// defaultErrorPrefixDisplayLineLength - Do NOT access
// defaultErrorPrefixDisplayLineLength without first locking
// 'defaultErrPrefLineLenLock'
var defaultErrorPrefixDisplayLineLength uint

var defaultErrPrefLineLenLock *sync.Mutex

type errPrefQuark struct {
	lock *sync.Mutex
}

// convertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// Examples of non-printable characters are '\n', '\t' or 0x06
// (Acknowledge). These example characters would be translated into
// printable string characters as: "\\n", "\\t" and "[ACK]".
//
// Space characters are typically translated as " ". However, if
// the input parameter 'convertSpace' is set to 'true' then all
// spaces are converted to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters will be converted to
//       printable characters.
//
//  convertSpace        bool
//     - Space or white space characters (0x20) are by default
//       translated as " ". However, if this parameter is set to
//       'true', space characters will be converted to "[SPACE]".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  printableChars      string
//     - This returned string is identical to input parameter
//       'nonPrintableChars' with the exception that non-printable
//       characters are translated into printable characters.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello world!\n"
//  testRunes := []rune(testStr)
//  ePrefix := "theCallingFunction()"
//
//  ePrefQuark := errPrefQuark{}
//
//  actualStr :=
//    ePrefQuark.
//      convertNonPrintableChars(
//           testRunes,
//           true,
//           ePrefix)
//
//  ----------------------------------------------------
//  'actualStr' is now equal to:
//     "Hello[SPACE]world!\\n"
//
//
func (ePrefQuark *errPrefQuark) convertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool,
	ePrefix string) (
	printableChars string,
	err error) {

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "errPrefQuark.convertNonPrintableChars() "

	lenNonPrintableChars := len(nonPrintableChars)

	if lenNonPrintableChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nonPrintableChars' is invalid!\n"+
			"'nonPrintableChars' is an empty or zero lenght rune array.\n",
			ePrefix)

		return "[EMPTY] Array!", err
	}

	var b strings.Builder
	b.Grow(lenNonPrintableChars * 5)

	for i := 0; i < lenNonPrintableChars; i++ {
		cRune := nonPrintableChars[i]

		switch cRune {
		case 0:
			b.WriteString("[NULL]") // 0x00 NULL
		case 1:
			b.WriteString("[SOH]") // 0x01 State of Heading
		case 2:
			b.WriteString("[STX]") // 0x02 State of Text
		case 3:
			b.WriteString("[ETX]") // 0x03 End of Text
		case 4:
			b.WriteString("[EOT]") // 0X04 End of Transmission
		case 5:
			b.WriteString("[ENQ]") // 0x05 Enquiry
		case 6:
			b.WriteString("[ACK]") // 0x06 Acknowledge
		case '\a':
			b.WriteString("\\a") // U+0007 alert or bell
		case '\b':
			b.WriteString("\\b") // U+0008 backspace
		case '\t':
			b.WriteString("\\t") // U+0009 horizontal tab
		case '\n':
			b.WriteString("\\n") // U+000A line feed or newline
		case '\v':
			b.WriteString("\\v") // U+000B vertical tab
		case '\f':
			b.WriteString("\\f") // U+000C form feed
		case '\r':
			b.WriteString("\\r") // U+000D carriage return
		case 14:
			b.WriteString("[SO]") // U+000E Shift Out
		case 15:
			b.WriteString("[SI]") // U+000F Shift In
		case '\\':
			b.WriteString("\\") // U+005c backslash
		case ' ':
			// 0X20 Space character
			if convertSpace {
				b.WriteString("[SPACE]")
			} else {
				b.WriteRune(' ')
			}

		default:
			b.WriteRune(cRune)

		}

	}

	return b.String(), err
}

// emptyEPrefCollection - Receives a pointer to an ErrPrefixDto
// object an proceeds to delete all of the existing error prefix
// and error context information. When completed, the ErrPrefixDto
// object's collection of ErrorPrefixInfo elements will be set to
// 'nil'.
//
func (ePrefQuark *errPrefQuark) emptyErrPrefInfoCollection(
	ePrefixDto *ErrPrefixDto,
	errPrefStr string) error {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	errPrefStr = errPrefStr +
		"\nerrPrefQuark.emptyErrPrefInfoCollection()"

	if ePrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefixDto' is invalid!\n"+
			"'ePrefixDto' is a 'nil' pointer.\n",
			errPrefStr)
	}

	if ePrefixDto.ePrefCol == nil {
		return nil
	}

	lenEPrefCol := len(ePrefixDto.ePrefCol)

	if lenEPrefCol == 0 {
		return nil
	}

	for i := 0; i < lenEPrefCol; i++ {
		ePrefixDto.ePrefCol[i].Empty()
	}

	ePrefixDto.ePrefCol = nil

	return nil
}

// getErrPrefDisplayLineLength - Returns the current value for the
// maximum number of characters allowed in an error prefix text
// display line.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return an unsigned integer value
//       specifying the maximum number of characters allowed
//       in an error prefix text display line.
//
func (ePrefQuark *errPrefQuark) getErrPrefDisplayLineLength() uint {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	var maxErrPrefixStringLength uint

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	if defaultErrorPrefixDisplayLineLength == 0 {

		defaultErrorPrefixDisplayLineLength =
			constDefaultErrPrefLineLength

	}

	maxErrPrefixStringLength =
		defaultErrorPrefixDisplayLineLength

	return maxErrPrefixStringLength
}

// getMasterErrPrefDisplayLineLength - Always returns the master
// error prefix display line length. The default maximum length
// for error prefix display lines is currently 40-characters.
//
func (ePrefQuark *errPrefQuark) getMasterErrPrefDisplayLineLength() uint {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	var maxErrPrefixStringLength uint

	maxErrPrefixStringLength =
		constDefaultErrPrefLineLength

	return maxErrPrefixStringLength
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (ePrefQuark *errPrefQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}

// ptr() - Returns a pointer to a new instance of
// errPrefQuark.
//
func (ePrefQuark errPrefQuark) ptr() *errPrefQuark {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	return &errPrefQuark{}
}

// resetErrPrefDisplayLineLengthToDefault - Reset the maximum error
// prefix string line length to the default value.
func (ePrefQuark *errPrefQuark) resetErrPrefDisplayLineLengthToDefault() {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	defaultErrorPrefixDisplayLineLength =
		constDefaultErrPrefLineLength

}

// setErrPrefDisplayLineLength - Sets the value of the maximum
// error prefix line length. This maximum limit controls the length
// of text lines produced for display of error prefix information.
//
// This method will set a global variable to the line length
// specified by input parameter 'maxErrPrefixTextLineLength'. Once
// this maximum limit is set it will control the line lengths for
// all instances of the 'ErrPref' type unless or until this limit
// is reset to another value by calling this method.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  maxErrPrefixTextLineLength      uint
//     - This unsigned integer value will be used to set the
//       maximum number of characters allowed in a text display
//       line for error prefix information.
//
//       If 'maxErrPrefixTextLineLength' is set to a value of zero
//       (0), this method will take no action and return.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
func (ePrefQuark *errPrefQuark) setErrPrefDisplayLineLength(
	maxErrPrefixStringLength uint) {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	if maxErrPrefixStringLength == 0 {
		return
	}

	defaultErrorPrefixDisplayLineLength =
		maxErrPrefixStringLength

	return
}
