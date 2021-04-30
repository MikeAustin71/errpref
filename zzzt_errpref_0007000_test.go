package errpref

import "testing"

func TestErrorPrefixInfo_SetErrPrefixStr(t *testing.T) {

	ePInfo := ErrorPrefixInfo{}

	ePInfo.SetErrPrefixStr("")

	if ePInfo.IsValidInstance() {
		t.Error("ERROR:\n" +
			"Expected ePInfo.IsValidInstance()==false\n" +
			"because ePInfo is empty.\n" +
			"HOWEVER, ePInfo.IsValidInstance()==true\n")
	}

}
