package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefMolecule struct {
	lock *sync.Mutex
}

// assembleNewErrPref - Assembles, consolidates and formats a new
// error prefix string from two subsidiary components: the new
// error prefix string and the associated new error context
// description.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPref                    string
//     - This is the new error prefix string which will be combined
//       with the new error context description to create a
//       consolidated new error prefix text description. Typically,
//       a new error prefix string identifies the name of the
//       method which begin executing next.
//
//
//  newContext                    string
//     - An optional error context description. This is the error
//       context information associated with the new error prefix
//       ('newErrPref'). Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of this function.
//
//
//  maxErrPrefixTextLineLength            uint
//      - Specifies the maximum number of text characters which can
//        be on a single line for a new line character ('\n') is
//        inserted. If this value is zero, it will be set to the
//        default value of 40.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  consolidatedNewEPrefContext   string
//     - The text string which combines and format both the new
//       error prefix and the associated new error context
//       description.
//
//
//  lenConsolidatedNewEPrefContext int
//     - This parameter returns the length of the
//       'consolidatedNewEPrefContext' text string.
//
//  lenNewErrPrefCleanStr         int
//     - This parameter returns the length of the error prefix
//       string included in 'consolidatedNewEPrefContext'.
//
//  lenNewErrContextCleanStr      int
//     - This parameter returns the length of the error context
//       string included in 'consolidatedNewEPrefContext'.
//
func (ePrefMolecule *errPrefMolecule) assembleNewErrPref(
	newErrPref string,
	newContext string,
	maxErrStringLength uint) (
	consolidatedNewEPrefContext string,
	lenConsolidatedNewEPrefContext int,
	lenNewErrPrefCleanStr int,
	lenNewErrContextCleanStr int) {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength =
			ePrefQuark.getErrPrefDisplayLineLength()
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	newErrPref,
		lenNewErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(newErrPref)

	newContext,
		lenNewErrContextCleanStr =
		ePrefElectron.cleanErrorContextStr(newContext)

	if lenNewErrPrefCleanStr == 0 &&
		lenNewErrContextCleanStr == 0 {

		consolidatedNewEPrefContext = ""

	} else if lenNewErrPrefCleanStr > 0 &&
		lenNewErrContextCleanStr == 0 {

		consolidatedNewEPrefContext =
			newErrPref

	} else if lenNewErrPrefCleanStr == 0 &&
		lenNewErrContextCleanStr > 0 {

		consolidatedNewEPrefContext =
			newContext
	} else {
		//lenNewErrPrefCleanStr > 0 &&
		//lenNewErrContextCleanStr > 0
		if uint(lenNewErrPrefCleanStr+
			lenNewErrContextCleanStr+
			3) > maxErrStringLength {

			consolidatedNewEPrefContext =
				newErrPref +
					delimiters.GetNewLineContextDelimiter() +
					newContext

		} else {
			//uint(lenNewErrPrefCleanStr +
			//	lenNewErrContextCleanStr +
			//	3) <= maxErrPrefixTextLineLength
			consolidatedNewEPrefContext =
				newErrPref +
					delimiters.GetInLineContextDelimiter() +
					newContext
		}
	}

	lenConsolidatedNewEPrefContext =
		len(consolidatedNewEPrefContext)

	return consolidatedNewEPrefContext,
		lenConsolidatedNewEPrefContext,
		lenNewErrPrefCleanStr,
		lenNewErrContextCleanStr
}

// ptr() - Returns a pointer to a new instance of errPrefMolecule.
//
func (ePrefMolecule errPrefMolecule) ptr() *errPrefMolecule {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	newErrPrefMolecule := errPrefMolecule{}

	return &newErrPrefMolecule
}

// writeNewEPrefWithContext - Writes Error Prefix strings to a
// string builder. This algorithm is designed for Error Prefixes
// that DO have an associated error context string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder          *strings.Builder
//     - This method will write output characters to this string
//       builder.
//
//
//  lineLenCalc         *EPrefixLineLenCalc
//     - Pointer to an instance of EPrefixLineLenCalc. This object
//       contains all the data and line length calculations
//       necessary to format the error prefix string and associated
//       error context information. The formatted string will be
//       written to the string builder object passed through input
//       parameter 'strBuilder'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumStrFmtSpecCurrencyValueDto
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithContext(
	strBuilder *strings.Builder,
	lineLenCalc *EPrefixLineLenCalc) error {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	localErrPrefix := "errPrefMolecule.writeNewEPrefWithContext() "

	if strBuilder == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is a 'nil' pointer!\n",
			localErrPrefix)
	}

	if lineLenCalc == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'lineLenCalc' is a 'nil' pointer!\n",
			localErrPrefix)
	}

	err := lineLenCalc.IsValidInstanceError(
		localErrPrefix + "lineLenCalc\n")

	if err != nil {
		return err
	}

	ePrefNeutron := errPrefNeutron{}

	if lineLenCalc.CurrLineLenExceedsMaxLineLen() {
		// The lastStr is already longer than
		// than the maximum line length.

		ePrefNeutron.writeCurrentLineStr(
			strBuilder,
			lineLenCalc)
	}

	if lineLenCalc.EPrefixWithContextExceedsRemainLineLen() {

		if lineLenCalc.GetCurrLineStrLength() > 0 {

			ePrefNeutron.writeCurrentLineStr(
				strBuilder,
				lineLenCalc)

		}

		if lineLenCalc.EPrefixWithContextExceedsRemainLineLen() {

			if strBuilder.Len() > 0 {
				strBuilder.WriteString(
					lineLenCalc.GetDelimiterNewLineErrPrefix())
			}

			strBuilder.WriteString(
				lineLenCalc.GetErrorPrefixStr())

			strBuilder.WriteString(
				lineLenCalc.GetDelimiterNewLineErrContext())

			strBuilder.WriteString(
				lineLenCalc.GetErrorContextStr())

			//if !lineLenCalc.IsErrPrefixLastIndex() {
			//	strBuilder.WriteString(
			//		lineLenCalc.GetDelimiterNewLineErrPrefix())
			//}

			return nil
		}
		// End Of
		//newLenLastStr +
		//	lenEPrefWithContext > remainingLineLen
	}

	//newLenLastStr +
	//	lenEPrefWithContext <= remainingLineLen
	// The line length of the next write block
	// will fit on the end of the 'lastStr'
	newLastStr := ""

	if lineLenCalc.GetCurrLineStrLength() > 0 {

		newLastStr = lineLenCalc.GetCurrLineStr()

		newLastStr += lineLenCalc.GetDelimiterInLineErrPrefix()

	}

	newLastStr += lineLenCalc.GetErrorPrefixStr()
	newLastStr += lineLenCalc.GetDelimiterInLineErrContext()
	newLastStr += lineLenCalc.GetErrorContextStr()

	lineLenCalc.SetCurrentLineStr(newLastStr)

	return nil
}

// writeNewEPrefWithOutContext - Writes Error Prefix strings to a
// string builder. This algorithm is designed for Error Prefixes
// that do NOT have an associated error context string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder          *strings.Builder
//     - This method will write output characters to this string
//       builder.
//
//
//  lineLenCalc         *EPrefixLineLenCalc
//     - Pointer to an instance of EPrefixLineLenCalc. This object
//       contains all the data and line length calculations
//       necessary to format the error prefix string and associated
//       error context information. The formatted string will be
//       written to the string builder object passed through input
//       parameter 'strBuilder'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumStrFmtSpecCurrencyValueDto
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithOutContext(
	strBuilder *strings.Builder,
	lineLenCalc *EPrefixLineLenCalc) error {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	localErrPrefix := "errPrefMolecule.writeNewEPrefWithOutContext() "

	if strBuilder == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is a 'nil' pointer!\n",
			localErrPrefix)
	}

	if lineLenCalc == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'lineLenCalc' is a 'nil' pointer!\n",
			localErrPrefix)
	}

	err := lineLenCalc.IsValidInstanceError(
		localErrPrefix + "lineLenCalc\n")

	if err != nil {
		return err
	}

	ePrefNeutron := errPrefNeutron{}

	if lineLenCalc.CurrLineLenExceedsMaxLineLen() {
		// The lastStr is already longer than
		// than the maximum line length.

		ePrefNeutron.writeCurrentLineStr(
			strBuilder,
			lineLenCalc)
	}

	if lineLenCalc.EPrefWithoutContextExceedsRemainLineLen() {

		if lineLenCalc.GetCurrLineStrLength() > 0 {

			ePrefNeutron.writeCurrentLineStr(
				strBuilder,
				lineLenCalc)

		}

		if lineLenCalc.EPrefWithoutContextExceedsRemainLineLen() {

			if strBuilder.Len() > 0 {
				strBuilder.WriteString(
					lineLenCalc.GetDelimiterNewLineErrPrefix())
			}

			strBuilder.WriteString(
				lineLenCalc.GetErrorPrefixStr())

			return nil
			// End of if lenEPrefWithoutContext > remainingLineLen
		}
	}

	newLastStr := ""

	if lineLenCalc.GetCurrLineStrLength() > 0 {

		newLastStr = lineLenCalc.GetCurrLineStr()

		newLastStr += lineLenCalc.GetDelimiterInLineErrPrefix()

	}

	newLastStr += lineLenCalc.GetErrorPrefixStr()

	lineLenCalc.SetCurrentLineStr(newLastStr)

	return nil
}
