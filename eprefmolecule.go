package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefMolecule struct {
	lock *sync.Mutex
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
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
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
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
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
