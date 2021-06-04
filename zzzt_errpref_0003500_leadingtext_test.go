package errpref

import (
	"fmt"
	"strings"
	"testing"
)

func TestLeadingText_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	outputStr := fmt.Sprintf("%v"+
		"Error: Divide by Zero!",
		ePDto.String())

	expectedStr :=
		"\\n--------------------------------------------------\\n" +
			"Tx1.Something()[SPACE]-[SPACE]Tx2.SomethingElse()" +
			"\\n--------------------------------------------------\\n" +
			"Error:[SPACE]Divide[SPACE]by[SPACE]Zero!"

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(outputStr),
		true)

	if expectedStr != actualStr {
		t.Errorf("ERROR Expected String DID NOT MATCH Actual String!\n"+
			"Expected Str: %v\n"+
			"  Actual Str: %v\n",
			expectedStr, actualStr)
	}
}

func TestErrPrefixDto_GetLeadingTextStr_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	actualLeadingTxt := ePDto.GetLeadingTextStr()

	if leadingText != actualLeadingTxt {
		t.Errorf("ERROR: Expected Leading Text String DOES "+
			"NOT MATCH\nActual Leading Text String!\n"+
			"Expected Leading Text String= '%v'\n"+
			"  Actual Leading Text String= '%v'\n",
			leadingText,
			actualLeadingTxt)
	}

}

func TestErrPrefixDto_GetTrailingTextStr_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	actualTrailingTxt := ePDto.GetTrailingTextStr()

	if trailingText != actualTrailingTxt {
		t.Errorf("ERROR: Expected Trailing Text String DOES "+
			"NOT MATCH\nActual Trailing Text String!\n"+
			"Expected Trailing Text String= '%v'\n"+
			"  Actual Trailing Text String= '%v'\n",
			trailingText,
			actualTrailingTxt)
	}

}

func TestErrPrefixDto_ClearLeadingTextStr_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	actualLeadingTxt := ePDto.GetLeadingTextStr()

	if leadingText != actualLeadingTxt {
		t.Errorf("ERROR: Expected Leading Text String DOES "+
			"NOT MATCH\nActual Leading Text String!\n"+
			"Expected Leading Text String= '%v'\n"+
			"  Actual Leading Text String= '%v'\n",
			leadingText,
			actualLeadingTxt)
		return
	}

	ePDto.ClearLeadingTextStr()

	actualLeadingTxt = ePDto.GetLeadingTextStr()

	if len(actualLeadingTxt) != 0 {
		t.Errorf("ERROR: After calling ePDto.ClearLeadingTextStr(),"+
			"the leading text string NOT EMPTY!\n"+
			"Actual Leading Text String= '%v'\n",
			actualLeadingTxt)
	}

}

func TestErrPrefixDto_ClearTrailingTextStr_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	actualTrailingTxt := ePDto.GetTrailingTextStr()

	if trailingText != actualTrailingTxt {
		t.Errorf("ERROR: Expected Trailing Text String DOES "+
			"NOT MATCH\nActual Trailing Text String!\n"+
			"Expected Trailing Text String= '%v'\n"+
			"  Actual Trailing Text String= '%v'\n",
			trailingText,
			actualTrailingTxt)
		return
	}

	ePDto.ClearTrailingTextStr()

	actualTrailingTxt = ePDto.GetTrailingTextStr()

	if len(actualTrailingTxt) != 0 {
		t.Errorf("ERROR: After calling ePDto.ClearTrailingTextStr(),"+
			"the trailing text string NOT EMPTY!\n"+
			"Actual Trailing Text String= '%v'\n",
			actualTrailingTxt)
	}
}
