package errpref

type IBasicErrorPrefix interface {
	GetEPrefStrings() [][2]string
}

type IBuilderErrorPrefix interface {
	GetEPrefStrings() [][2]string

	SetEPrefStrings(twoDStrArray [][2]string)

	String() string
}

type IErrorPrefix interface {
	GetEPrefStrings() [][2]string

	GetEPrefCollectionLen() int

	GetIsLastLineTerminatedWithNewLine() bool

	SetCtx(newErrContext string)

	SetCtxEmpty()

	SetEPref(newErrPrefix string)

	SetEPrefCtx(newErrPrefix string, newErrContext string)

	SetEPrefOld(oldErrPrefix string)

	SetEPrefStrings(twoDStrArray [][2]string)

	SetMaxTextLineLen(maxErrPrefixTextLineLength int)

	SetIsLastLineTermWithNewLine(isLastLineTerminatedWithNewLine bool)

	String() string
}
