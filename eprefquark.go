package errpref

import (
	"fmt"
	"strings"
	"sync"
)

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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'. This parameter is optional.
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
//  err                 error
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

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

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

// convertPrintableChars - Converts selected printable characters
// to their non-printable or native equivalent. For example,
// instances of '\\n' in a string will be converted to '\n'.
//
// Additional examples of converted printable string characters
// are: "\\n", "\\t" and "[ACK]". These printable characters be
// converted into their native, non-printable state: '\n', '\t' or
// 0x06 (Acknowledge).
//
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  printableChars      string
//     - A string which may contain non-printable characters converted
//       to their printable equivalents. These printable characters will
//       be converted back to their native, non-printable values.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'. This parameter is optional.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters were be converted from the
//       printable characters contained in input parameter
//       'printableChars'.
//
//
//  err                 error
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
//  testStr := "Hello[SPACE]world!\\n"
//  ePrefix := "theCallingFunction()"
//
//  ePrefQuark := errPrefQuark{}
//
//  actualRuneArray :=
//    ePrefQuark.
//      convertPrintableChars(
//           testStr,
//           ePrefix)
//
//  ----------------------------------------------------
//  'actualRuneArray' is now equal to:
//     "Hello world!\n"
//
func (ePrefQuark *errPrefQuark) convertPrintableChars(
	printableChars string,
	ePrefix string) (
	nonPrintableChars []rune,
	err error) {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "errPrefQuark.convertPrintableChars() "

	nonPrintableChars = make([]rune, 0)

	lenPrintableChars := len(printableChars)

	if lenPrintableChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'printableChars' is invalid!\n"+
			"'printableChars' is an empty or zero lenght string.\n",
			ePrefix)

		return nonPrintableChars, err
	}

	printableChars = strings.Replace(
		printableChars, "[SPACE]", " ", -1)

	printableChars = strings.Replace(
		printableChars, "[NULL]", string(rune(0x00)), -1) // 0x00 NULL

	printableChars = strings.Replace(
		printableChars, "[SOH]", string(rune(0x01)), -1) // 0x01 State of Heading

	printableChars = strings.Replace(
		printableChars, "[STX]", string(rune(0x02)), -1) // 0x02 State of Text

	printableChars = strings.Replace(
		printableChars, "[ETX]", string(rune(0x03)), -1) // 0x03 End of Text

	printableChars = strings.Replace(
		printableChars, "[EOT]", string(rune(0x04)), -1) // 0X04 End of Transmission

	printableChars = strings.Replace(
		printableChars, "[ENQ]", string(rune(0x05)), -1) // 0x05 Enquiry

	printableChars = strings.Replace(
		printableChars, "[ACK]", string(rune(0x06)), -1) // 0x06 Acknowledge

	printableChars = strings.Replace(
		printableChars, "\\a", string(rune(0x07)), -1) // U+0007 alert or bell

	printableChars = strings.Replace(
		printableChars, "\\b", string(rune(0x08)), -1) // U+0008 backspace

	printableChars = strings.Replace(
		printableChars, "\\t", string(rune(0x09)), -1) // U+0009 horizontal tab

	printableChars = strings.Replace(
		printableChars, "\\n", string(rune(0x0A)), -1) // U+000A line feed or newline

	printableChars = strings.Replace(
		printableChars, "\\v", string(rune(0x0B)), -1) // U+000B vertical tab

	printableChars = strings.Replace(
		printableChars, "\\f", string(rune(0x0C)), -1) // U+000C form feed

	printableChars = strings.Replace(
		printableChars, "\\r", string(rune(0x0D)), -1) // U+000D carriage return

	printableChars = strings.Replace(
		printableChars, "[SO]", string(rune(0x0E)), -1) // U+000E Shift Out

	printableChars = strings.Replace(
		printableChars, "[SI]", string(rune(0x0F)), -1) // U+000F Shift In

	printableChars = strings.Replace(
		printableChars, "\\", string(rune(0x5c)), -1) // U+005c backslash

	nonPrintableChars = []rune(printableChars)

	return nonPrintableChars, err
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

	erPrfPreon := errPrefPreon{}

	minimumErrPrefLineLen := erPrfPreon.getMinErrPrefLineLength()

	if defaultErrorPrefixDisplayLineLength <
		minimumErrPrefLineLen {

		defaultErrorPrefixDisplayLineLength =
			erPrfPreon.getDefaultErrPrefLineLength()

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

	return errPrefPreon{}.ptr().getDefaultErrPrefLineLength()
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
		errPrefPreon{}.ptr().getDefaultErrPrefLineLength()

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
//       If 'maxErrPrefixTextLineLength' is set to a value less
//       than the minimum Error Prefix Line Length, this method
//       will take no action and return.
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

	minimum := errPrefPreon{}.ptr().getMinErrPrefLineLength()

	if maxErrPrefixStringLength < minimum {
		return
	}

	defaultErrorPrefixDisplayLineLength =
		maxErrPrefixStringLength

	return
}
