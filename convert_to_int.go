package u

import (
	"regexp"
	"strconv"
	"strings"
)

// ToInt to int
func ToInt(val interface{}) int {
	return int(ToInt64(val))
}

// ToInt8 to int8
func ToInt8(val interface{}) int8 {
	return int8(ToInt64(val))
}

// ToInt16 to int16
func ToInt16(val interface{}) int16 {
	return int16(ToInt64(val))
}

// ToInt32 to int32
func ToInt32(val interface{}) int32 {
	return int32(ToInt64(val))
}

// ToInt64 to int64
func ToInt64(val interface{}) int64 {
	switch val.(type) {
	case string:
		str := val.(string)
		matched, _ := regexp.MatchString(`^[0-9.]+$`, str)
		if !matched {
			return 0
		}

		if strings.Count(val.(string), ".") <= 1 {
			floatVal, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0
			}
			return int64(floatVal)
		}
		return 0
	case int:
		return int64(val.(int))
	case int8:
		return int64(val.(int8))
	case int16:
		return int64(val.(int16))
	case int32:
		return int64(val.(int32))
	case int64:
		return val.(int64)
	case uint:
		return int64(val.(uint))
	case uint8:
		return int64(val.(uint8))
	case uint16:
		return int64(val.(uint16))
	case uint32:
		return int64(val.(uint32))
	case uint64:
		return int64(val.(int64))
	case float32:
		return int64(val.(float32))
	case float64:
		return int64(val.(float64))
	default:
		return 0
	}
}