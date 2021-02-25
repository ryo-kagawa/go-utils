package logger

type Type uint

const (
	TypeUnknown Type = iota
	TypeNone
	TypeStandardOut
	TypeStandardError
	TypeFile
)

func ParseType(v string) Type {
	switch v {
	case "NONE":
		return TypeNone
	case "STDOUT":
		return TypeStandardOut
	case "STDERR":
		return TypeStandardError
	case "FILE":
		return TypeFile
	}
	return TypeUnknown
}
