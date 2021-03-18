package language

type Language int32

const (
	CPP  Language = 1
	JAVA Language = 2
)

var (
	CodeToTextMap = map[Language]string{
		CPP:  "C++",
		JAVA: "JAVA",
	}
)
