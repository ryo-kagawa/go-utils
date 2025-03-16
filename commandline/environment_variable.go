//go:build unix || plan9 || windows

package commandline

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"unsafe"
)

var (
	ErrorEnvironmentVariableInvalidDefault  error = errors.New("environment variable is invalid default")
	ErrorEnvironmentVariableInvalidValue    error = errors.New("environment variable is invalid value")
	ErrorEnvironmentVariableInvalidType     error = errors.New("environment variable is invalid type")
	ErrorEnvironmentVariableInvalidValidate error = errors.New("environment variable is invalid validate")
)

type validator interface {
	Validate() error
}

// Environment variables behave differently depending on the OS:
//
//	Unix: case sensitive.
//	Plan9: case sensitive.
//	Windows: case insensitive.
func EnvironmentVariableParse[T any]() (T, error) {
	setKeyList := make([]string, 0)
	var result T
	resultType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(&result)

	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		fieldKey := fieldType.Tag.Get("key")
		fieldValue := resultValue.Elem().Field(i)
		environmentValue, environmentExists := os.LookupEnv(fieldKey)
		if !environmentExists {
			continue
		}
		switch fieldValue.Kind() {
		case reflect.Bool:
			if environmentValue != "" {
				return result, errors.Join(ErrorEnvironmentVariableInvalidValue, fmt.Errorf("invalid environment variable fieldName: %s, fieldType: %s, environmentKey: %s, environmentValue: %s", fieldType.Name, fieldType.Type.Name(), fieldKey, environmentValue))
			}
			valuePointer := (*bool)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			*valuePointer = true
			setKeyList = append(setKeyList, fieldKey)
		case reflect.Int:
			valuePointer := (*int)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			intValue, err := strconv.ParseInt(environmentValue, 10, 64)
			if err != nil {
				return result, errors.Join(ErrorEnvironmentVariableInvalidValue, fmt.Errorf("invalid environment variable fieldName: %s, fieldType: %s, environmentKey: %s, environmentValue: %s", fieldType.Name, fieldType.Type.Name(), fieldKey, environmentValue))
			}
			*valuePointer = int(intValue)
			setKeyList = append(setKeyList, fieldKey)
		case reflect.String:
			valuePointer := (*string)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			*valuePointer = environmentValue
			setKeyList = append(setKeyList, fieldKey)
		default:
			return result, errors.Join(ErrorEnvironmentVariableInvalidType, fmt.Errorf("invalid environment variable type fieldName: %s, fieldType: %s", fieldType.Name, fieldType.Type.Name()))
		}
	}

	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		fieldKey := fieldType.Tag.Get("key")
		defaultValue, defaultExists := fieldType.Tag.Lookup("default")
		if !defaultExists {
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
				return result, errors.Join(ErrorEnvironmentVariableInvalidDefault, fmt.Errorf("invalid environment variable default fieldName: %s, fieldType: %s, default: %s", fieldType.Name, fieldType.Type.Name(), defaultValue))
			}
			*valuePointer = int(intValue)
		case reflect.String:
			valuePointer := (*string)(unsafe.Pointer(fieldValue.UnsafeAddr()))
			*valuePointer = defaultValue
		default:
			return result, errors.Join(ErrorEnvironmentVariableInvalidType, fmt.Errorf("invalid environment variable type fieldName: %s, fieldType: %s", fieldType.Name, fieldType.Type.Name()))
		}
	}

	if validator, ok := any(&result).(validator); ok {
		if err := validator.Validate(); err != nil {
			return result, errors.Join(ErrorEnvironmentVariableInvalidValidate, err)
		}
	}

	return result, nil
}
