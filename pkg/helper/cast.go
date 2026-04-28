package helper

import (
	"strconv"
	"time"

	"github.com/spf13/cast"
)

func ToInt(v any) int {
	return cast.ToInt(v)
}

func ToUint(v any) uint {
	if val, ok := v.(int8); ok {
		return uint(val)
	}
	return cast.ToUint(v)
}

func ToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	}
	return ""
}

func ToStringPtr(v any) *string {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case string:
		return &val
	case []byte:
		s := string(val)
		return &s
	}
	return nil
}

func ToIntPtr(v any) *int {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case int32:
		i := int(val)
		return &i
	case uint32:
		i := int(val)
		return &i
	case int64:
		i := int(val)
		return &i
	case float64:
		i := int(val)
		return &i
	case []byte:
		n, err := strconv.Atoi(string(val))
		if err != nil {
			return nil
		}
		return &n
	case string:
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		return &n
	}
	return nil
}

func ToTimePtr(v any) *time.Time {
	if v == nil {
		return nil
	}
	t, ok := v.(time.Time)
	if !ok {
		return nil
	}
	return &t
}

func ToBool(v any) bool {
	return cast.ToBool(v)
}
