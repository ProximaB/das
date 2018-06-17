package util

import "time"

func init() {
	JsonRequestDecoder.RegisterConverter(time.Time{}, decodeDate)
	JsonRequestDecoder.RegisterConverter([]int{}, decodeIntArray)
}
