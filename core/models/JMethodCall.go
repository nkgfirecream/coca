package models

type JMethodCall struct {
	Package           string
	Type              string
	Class             string
	MethodName        string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

func NewJMethodCall() JMethodCall {
	return *&JMethodCall{
		Package:           "",
		Type:              "",
		Class:             "",
		MethodName:        "",
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
	}
}