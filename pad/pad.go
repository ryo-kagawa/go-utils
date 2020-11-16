package pad

import "strings"

func pad(str, padString string, length int) string {
	shortage := length - len(str)
	if shortage <= 0 {
		return ""
	}
	return strings.Repeat(padString, shortage/len(padString)) + padString[:shortage%len(padString)]
}

func Left(str, padString string, length int) string {
	return pad(str, padString, length) + str
}

func Right(str, padString string, length int) string {
	return str + pad(str, padString, length)
}
