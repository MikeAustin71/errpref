package errpref

import "fmt"

func getValidErrPrefixDelimiters() ErrPrefixDelimiters {
	return ErrPrefixDelimiters{}.NewDefaults()
}

func getValidErrorPrefixDto() ErrPrefixDto {

	ePDto := ErrPrefixDto{}.New()

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	return ePDto
}

func getValidErrorPrefixCollectionStr() string {
	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness() : A=7 B=8 C=9"

	return initialStr
}

func getValidErrorPrefixInfoArray() []ErrorPrefixInfo {

	ePDto := ErrPrefixDto{}.New()

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	return ePDto.GetEPrefCollection()
}

func getValidErrorPrefixInfo() ErrorPrefixInfo {

	ePrefInfoOne := ErrorPrefixInfo{}.New()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	return ePrefInfoOne
}

func getValidEPrefixLineLenCalc(errPref string) (
	EPrefixLineLenCalc,
	error) {

	errPref += "getValidEPrefixLineLenCalc()\n"

	ePrefLineLenOne := EPrefixLineLenCalc{}

	ePrefInfoOne := ErrorPrefixInfo{}.Ptr()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	err :=
		ePrefLineLenOne.SetErrPrefixInfo(
			ePrefInfoOne,
			errPref)

	if err != nil {

		return EPrefixLineLenCalc{},
			fmt.Errorf("%v\n"+
				"ERROR returned by ePrefLineLenOne.SetErrPrefixInfo()\n"+
				"%v\n",
				errPref,
				err.Error())
	}

	err =
		ePrefLineLenOne.SetEPrefDelimiters(
			ErrPrefixDelimiters{}.NewDefaults(),
			errPref)

	if err != nil {

		return EPrefixLineLenCalc{},
			fmt.Errorf("%v\n"+
				"ERROR returned by ePrefLineLenOne.SetEPrefDelimiters()\n"+
				"%v\n",
				errPref,
				err.Error())
	}

	return ePrefLineLenOne, nil
}

func getValidTwoDStrArray() [][2]string {

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	return twoDSlice
}
