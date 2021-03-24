package strings

import "strings"

type pad struct{}

var Pad = pad{}

func (p pad) pad(str, padString string, length int) string {
	shortage := length - len(str)
	if shortage <= 0 {
		return ""
	}
	return strings.Repeat(padString, shortage/len(padString)) + padString[:shortage%len(padString)]
}

func (p pad) Left(str, padString string, length int) string {
	return p.pad(str, padString, length) + str
}

func (p pad) Right(str, padString string, length int) string {
	return str + p.pad(str, padString, length)
}
