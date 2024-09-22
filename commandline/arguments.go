package commandline

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"unsafe"
)

var (
	ErrorInvalidArgument error = errors.New("invalid argument")
	ErrorInvalidDefault  error = errors.New("invalid default")
	ErrorInvalidType     error = errors.New("invalid type")
)

func ArgumentsParse[T any](arguments []string) (T, error) {
	setKeyList := make([]string, 0, len(arguments))
	var result T
	resultType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(&result)

	for _, arg := range arguments {
		argKey, argValue, argCut := strings.Cut(arg, "=")
		argParseOk := false
		for i := 0; i < resultType.NumField(); i++ {
			fieldType := resultType.Field(i)
			fieldKey := fieldType.Tag.Get("key")
			if argKey == fieldKey {
				fieldValue := resultValue.Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					if argCut {
						return result, errors.Join(ErrorInvalidArgument, fmt.Errorf("invalid arguments fieldName: %s, fieldType: %s, argument: %s", fieldType.Name, fieldType.Type.Name(), arg))
					}
					valuePointer := (*bool)(unsafe.Pointer(fieldValue.UnsafeAddr()))
					*valuePointer = true
					setKeyList = append(setKeyList, fieldKey)
				case reflect.Int:
					valuePointer := (*int)(unsafe.Pointer(fieldValue.UnsafeAddr()))
					intValue, err := strconv.ParseInt(argValue, 10, 64)
					if err != nil {
						return result, errors.Join(ErrorInvalidArgument, fmt.Errorf("invalid arguments fieldName: %s, fieldType: %s, argument: %s", fieldType.Name, fieldType.Type.Name(), arg))
					}
					*valuePointer = int(intValue)
					setKeyList = append(setKeyList, fieldKey)
				case reflect.Slice:
					switch fieldType.Type.Elem().Kind() {
					case reflect.Int:
						valuePointer := (*[]int)(unsafe.Pointer(fieldValue.UnsafeAddr()))
						intValue, err := strconv.ParseInt(argValue, 10, 64)
						if err != nil {
							return result, errors.Join(ErrorInvalidArgument, fmt.Errorf("invalid arguments fieldName: %s, fieldType: []%s, argument: %s", fieldType.Name, fieldType.Type.Elem().Kind(), arg))
						}
						*valuePointer = append(*valuePointer, int(intValue))
						setKeyList = append(setKeyList, fieldKey)
					case reflect.String:
						valuePointer := (*[]string)(unsafe.Pointer(fieldValue.UnsafeAddr()))
						*valuePointer = append(*valuePointer, argValue)
						setKeyList = append(setKeyList, fieldKey)
					default:
						return result, errors.Join(ErrorInvalidType, fmt.Errorf("invalid type fieldName: %s, fieldType: []%s", fieldType.Name, fieldType.Type.Elem().Kind()))
					}
				case reflect.String:
					valuePointer := (*string)(unsafe.Pointer(fieldValue.UnsafeAddr()))
					*valuePointer = argValue
					setKeyList = append(setKeyList, fieldKey)
				default:
					return result, errors.Join(ErrorInvalidType, fmt.Errorf("invalid type fieldName: %s, fieldType: %s", fieldType.Name, fieldType.Type.Name()))
				}
				argParseOk = true
			}
		}
		if !argParseOk {
			return result, errors.Join(ErrorInvalidArgument, fmt.Errorf("invalid argument not found target field argument: %s", arg))
		}
	}

	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		fieldKey := fieldType.Tag.Get("key")
		defaultValue, ok := fieldType.Tag.Lookup("default")
		if !ok {
			continue
		}
		if slices.Contains(setKeyList, fieldKey) {
			continue
		}
		fieldValue := resultValue.Elem().Field(i)
		switch fieldValue.Kind() {
		case reflect.Int:
			valuePointer := (*int)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			intValue, err := strconv.ParseInt(defaultValue, 10, 64)
			if err != nil {
				return result, errors.Join(ErrorInvalidDefault, fmt.Errorf("invalid default fieldName: %s, fieldType: %s, default: %s", fieldType.Name, fieldType.Type.Name(), defaultValue))
			}
			*valuePointer = int(intValue)
		case reflect.String:
			valuePointer := (*string)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			*valuePointer = defaultValue
		default:
			return result, errors.Join(ErrorInvalidType, fmt.Errorf("invalid type fieldName: %s, fieldType: %s", fieldType.Name, fieldType.Type.Name()))
		}
	}

	return result, nil
}
