package errpref

import (
	"sync"
)

// ErrPref - This type is provides methods useful in formatting
// error prefix and error context strings.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to
// document the thread of code execution by listing the calling
// sequence for specific functions and methods.
//
// The error context string is designed to provide additional
// error context information associated with the currently
// executing function or method. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// Note that there are no 'pointer' methods provided for this
// type. This is because the type is not designed to store
// information. Its only function is to receive process and
// return strings of error prefix information.
//
// ------------------------------------------------------------------------
//
// - IMPORTANT -
// None of the error prefix strings returned by the methods on this
// type are terminated with a new line character ('\n'). That means
// that none of the strings end with a new line character.
//
// If you prefer that error prefix strings be terminated with a new
// line character, you have two options:
//
//   1. Add the terminating new line character in your code.
//
//                   OR
//
//   2. Use the Error Prefix Data Transfer Object type
//      'ErrPrefixDto'.
//
// ------------------------------------------------------------------------
//
// Recommended Usage Examples
//
//  ePrefix = ErrPref{}.EPrefCtx(
//    ePrefix,
//    "Tx7.TrySomethingNew()",
//     "")
//
//  ePrefix = ErrPref{}.EPref(
//     ePrefix,
//     "Tx14.SomeFabulousAndComplexStuff()")
//
//      actualStr := ErrPref{}.SetCtxt(
//       initialStr,
//        "A!=B")
//
//  Example Error Prefix String with Context information.
//
//       "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
//       " :  A->B\\n" +
//       "Tx2.SomethingElse() : A==B\\n" +
//       "Tx3.DoSomething() : A==10\\n" +
//       "Tx4() : A/10==4 - Tx5() : A!=B"
//
//
type ErrPref struct {
	maxErrPrefixTextLineLength uint
	lock                       *sync.Mutex
}

// ConvertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// This method is primarily used for testing an evaluation.
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
// This method is useful for verifying error prefix strings which
// are routinely populated with non-printable characters.
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
func (ePref ErrPref) ConvertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool) (
	printableChars string) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	printableChars,
		_ = ePrefQuark.convertNonPrintableChars(
		nonPrintableChars,
		convertSpace,
		"")

	return printableChars
}

// ConvertPrintableChars - Converts printable characters to their
// non-printable or native equivalent. For example, instances of
// '\\n' in a string will be converted to '\n'.
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
func (ePref ErrPref) ConvertPrintableChars(
	printableChars string,
	ePrefix string) (
	nonPrintableChars []rune,
	err error) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "ErrPref.ConvertPrintableChars() "

	nonPrintableChars,
		err = errPrefQuark{}.ptr().convertPrintableChars(
		printableChars,
		ePrefix)

	return nonPrintableChars, err
}

// FmtStr - Returns a formatted text representation of all error
// prefix and error context information contained in the input
// parameter string, 'errPref'.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context descriptive text provides additional
// information about the function or method identified in the
// associated error prefix string. Typical context information
// might include variable names, variable values and further
// details on function execution.
//
// The returned formatted text is generated using system default
// string delimiters.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errPref             string
//     - This string holds a series of error prefix and error
//       context text elements. Each element must be separated by
//       either a new line character '\n' or the string " - ". This
//       will format these elements and return them in a properly
//       configured error prefix string for text presentation.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - A string containing properly formatted error prefix and
//       error context pairs configured for configuration in an
//       error message.
//
func (ePref ErrPref) FmtStr(
	errPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePref.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getErrPrefDisplayLineLength()

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return errPrefMechanics{}.ptr().formatErrPrefix(
		ePref.maxErrPrefixTextLineLength,
		delimiters,
		errPref)
}

// GetDelimiters - Returns an ErrPrefixDelimiters object containing
// the string delimiters used to delimit error prefix and error
// context elements with strings.
//
// The ErrPref type uses the system default string delimiters for
// parsing both input and output error prefix and context strings.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
func (ePref ErrPref) GetDelimiters() ErrPrefixDelimiters {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return delimiters
}

// GetLastEPref - Returns the last error prefix, error context pair
// from a string consisting of a series of error prefix, error
// context pairs.
//
// The input parameter 'oldErrPrefix' will be parsed and the last
// error prefix and error context pair will be identified using
// system default string delimiters.
//
// The system default string delimiters are listed as follows:
//
//    New Line Error Prefix Delimiter = "\n"
//    In-Line Error Prefix Delimiter  = " - "
//    New Line Error Context Delimiter = "\n :  "
//    In-Line Error Context Delimiter = " : "
//
func (ePref ErrPref) GetLastEPref(
	oldErrPrefix string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePref.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getErrPrefDisplayLineLength()

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return errPrefMechanics{}.ptr().
		extractLastErrPrefCtxPair(
			ePref.maxErrPrefixTextLineLength,
			oldErrPrefix,
			delimiters)
}

// GetMaxErrPrefTextLineLength - Returns the current maximum number
// of characters allowed in an error prefix text line output
// display.
//
// To change or reset this maximum limit value, see method:
//     ErrPref.SetMaxErrPrefTextLineLength().
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  maxErrPrefixStringLength      uint
//     - This method will return an unsigned integer value
//       specifying the maximum number of characters allowed
//       in an error prefix text display line.
//
func (ePref ErrPref) GetMaxErrPrefTextLineLength() (
	maxErrPrefixStringLength uint) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	maxErrPrefixStringLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	return maxErrPrefixStringLength
}

// EPref - Returns a string concatenating the old error prefix the
// new custom, user-defined error prefix. The new error prefix is
// typically used to document method or function chains in error
// messages.
//
// The old error prefix contains the function chain or series which
// led to the function next in line for execution.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
// - IMPORTANT -
//
// The last line of error prefix strings returned by the methods on
// this type are NOT terminated with a new line character ('\n').
// That means that the last line of error prefix strings will never
// end with a new line character ('\n').
//
// If you prefer that the last line of error prefix strings be
// terminated with a new line character ('\n'), you have two
// options:
//
//   1. Add the terminating new line character in your code.
//
//                   OR
//
//   2. Use the Error Prefix Data Transfer Object type
//      'ErrPrefixDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components before being converted into
//       a single, formatted string containing error prefix and
//       error context information.
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited using system
//       default string delimiters.
//
//       The system default string delimiters for error prefix
//       elements are listed as follows:
//          New Line Error Prefix Delimiter = "\n"
//          In-Line Error Prefix Delimiter  = " - "
//
//       If this string contains associated error context strings
//       as well, they should be delimited using the same system
//       default string delimiters for error context elements:
//          New Line Error Context Delimiter = "\n :  "
//          In-Line Error Context Delimiter = " : "
//
//
//  newErrPref          string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       If 'newErrPref' equates to an empty string, this method will
//       return the formatted version of 'oldErrPref' and no new error
//       prefix information will be added.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return the consolidated error prefix text.
//
//       The error prefix text is designed to be configured at the
//       beginning of error messages and is most often used to
//       document the thread of code execution by listing the calling
//       sequence for specific functions and methods.
//
//
func (ePref ErrPref) EPref(
	oldErrPref string,
	newErrPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		"",
		ePref.maxErrPrefixTextLineLength,
		delimiters)
}

// EPrefCtx - Receives an old error prefix, new error prefix and
// a new context string which are concatenated and returned as a
// combined string.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// - IMPORTANT -
//
// The last line of error prefix strings returned by the methods on
// this type are NOT terminated with a new line character ('\n').
// That means that the last line of error prefix strings will never
// end with a new line character ('\n').
//
// If you prefer that the last line of error prefix strings be
// terminated with a new line character ('\n'), you have two
// options:
//
//   1. Add the terminating new line character in your code.
//
//                   OR
//
//   2. Use the Error Prefix Data Transfer Object type
//      'ErrPrefixDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components before being converted into
//       a single, formatted string containing error prefix and
//       error context information.
//
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited using system
//       default string delimiters.
//
//       The system default string delimiters for error prefix
//       elements are listed as follows:
//          New Line Error Prefix Delimiter = "\n"
//          In-Line Error Prefix Delimiter  = " - "
//
//       If this string contains associated error context strings
//       as well, they should be delimited using the same system
//       default string delimiters for error context elements:
//          New Line Error Context Delimiter = "\n :  "
//          In-Line Error Context Delimiter = " : "
//
//
//  newErrPref          string
//     - The new error prefix represents the error prefix string
//       associated with the function or method which is currently
//       executing. This parameter is optional and will accept an
//       empty string, but there isn't much point in calling this
//       method without a substantive value for 'newErrPref'.
//
//
//  newContext          string
//     - This is the error context information associated with the
//       new error prefix ('newErrPref'). This parameter is
//       optional and will accept an empty string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return the consolidated error prefix text.
//
//       The error prefix text is designed to be configured at the
//       beginning of error messages and is most often used to
//       document the thread of code execution by listing the calling
//       sequence for specific functions and methods.
//
//
//
// -----------------------------------------------------------------
//
// Usage Examples
//
//  errorPrefix = ErrPref{}.EPrefCtx(
//                           oldErrPref, // Assuming this is the old
//                                        // error prefix
//                           newErrPref,
//                           newContext)
//
//
func (ePref ErrPref) EPrefCtx(
	oldErrPref string,
	newErrPref string,
	newContext string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	ePrefMech := errPrefMechanics{}

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		newContext,
		ePref.maxErrPrefixTextLineLength,
		delimiters)
}

// EPrefOld - Receives an old or preexisting error prefix string
// which is parsed into error prefix and error context components
// and returned as a properly formatted error prefix string.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
// - IMPORTANT -
//
// The last line of error prefix strings returned by the methods on
// this type are NOT terminated with a new line character ('\n').
// That means that the last line of error prefix strings will never
// end with a new line character ('\n').
//
// If you prefer that the last line of error prefix strings be
// terminated with a new line character ('\n'), you have two
// options:
//
//   1. Add the terminating new line character in your code.
//
//                   OR
//
//   2. Use the Error Prefix Data Transfer Object type
//      'ErrPrefixDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components before being converted into
//       a single, formatted string containing error prefix and
//       error context information.
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited using system
//       default string delimiters.
//
//       The system default string delimiters for error prefix
//       elements are listed as follows:
//          New Line Error Prefix Delimiter = "\n"
//          In-Line Error Prefix Delimiter  = " - "
//
//       If this string contains associated error context strings
//       as well, they should be delimited using the same system
//       default string delimiters for error context elements:
//          New Line Error Context Delimiter = "\n :  "
//          In-Line Error Context Delimiter = " : "
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return the consolidated error prefix text.
//
//       The error prefix text is designed to be configured at the
//       beginning of error messages and is most often used to
//       document the thread of code execution by listing the calling
//       sequence for specific functions and methods.
//
//
//
// -----------------------------------------------------------------
//
// Usage Examples
//
//  errorPrefix = ErrPref{}.EPrefOld(
//                           oldErrPref) // Assuming this is the old
//                                       // or preexisting error
//                                       // prefix
//
//
func (ePref ErrPref) EPrefOld(
	oldErrPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	ePrefMech := errPrefMechanics{}

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		"",
		"",
		ePref.maxErrPrefixTextLineLength,
		delimiters)
}

// SetCtxt - Sets or resets the error context for the last error
// prefix. This operation either adds, or replaces, the error
// context string associated with the last error prefix in input
// parameter, 'oldErrPref'.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
// - IMPORTANT -
//
// The last line of error prefix strings returned by the methods on
// this type are NOT terminated with a new line character ('\n').
// That means that the last line of error prefix strings will never
// end with a new line character ('\n').
//
// If you prefer that the last line of error prefix strings be
// terminated with a new line character ('\n'), you have two
// options:
//
//   1. Add the terminating new line character in your code.
//
//                   OR
//
//   2. Use the Error Prefix Data Transfer Object type
//      'ErrPrefixDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous or preexisting error prefix
//       string. This string will be parsed into error prefix
//       and error context components before being converted into
//       a single, formatted string containing error prefix and
//       error context information.
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited using system
//       default string delimiters.
//
//       The system default string delimiters for error prefix
//       elements are listed as follows:
//          New Line Error Prefix Delimiter = "\n"
//          In-Line Error Prefix Delimiter  = " - "
//
//       If this string contains associated error context strings
//       as well, they should be delimited using the same system
//       default string delimiters for error context elements:
//          New Line Error Context Delimiter = "\n :  "
//          In-Line Error Context Delimiter = " : "
//
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in 'oldErrPref' already has an
//       associated error context, that context will be deleted and
//       replaced by 'newErrContext'. If, however, the last error
//       prefix in 'oldErrPref' does NOT have an associated error
//       context, this 'newErrContext' string will be added and
//       associated with the last error prefix in 'oldErrPref'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return the consolidated error prefix text.
//
//       The error prefix text is designed to be configured at the
//       beginning of error messages and is most often used to
//       document the thread of code execution by listing the calling
//       sequence for specific functions and methods.
//
func (ePref ErrPref) SetCtxt(
	oldErrPref string,
	newErrContext string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePref.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().
			getErrPrefDisplayLineLength()

	// Must have at least one error prefix
	// before you can add an error context

	if len(oldErrPref) == 0 {
		return oldErrPref
	}

	if len(newErrContext) == 0 {
		return oldErrPref
	}

	delimiters := ErrPrefixDelimiters{}.NewDefaults()

	return errPrefMechanics{}.ptr().setErrorContext(
		oldErrPref,
		newErrContext,
		ePref.maxErrPrefixTextLineLength,
		delimiters)
}

// SetMaxErrPrefTextLineLength - Sets the maximum limit on the
// number of characters allowed in an error prefix text line output
// display.
//
// -IMPORTANT -
// Setting this value will control the maximum character limit not
// only for this ErrPref instance, but will also control all that
// limit for all other instances of ErrPref created in this session.
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
func (ePref ErrPref) SetMaxErrPrefTextLineLength(
	maxErrPrefixTextLineLength uint) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	errPrefQuark{}.ptr().setErrPrefDisplayLineLength(
		maxErrPrefixTextLineLength)

	return
}

// SetMaxErrPrefTextLineLengthToDefault - Maximum Error Prefix Line
// Length is the maximum limit on the number of characters allowed
// in a single error prefix text line.
//
// This method resets that maximum limit to its default value of
// 40-characters.
//
func (ePref ErrPref) SetMaxErrPrefTextLineLengthToDefault() {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	errPrefQuark{}.ptr().resetErrPrefDisplayLineLengthToDefault()

}
