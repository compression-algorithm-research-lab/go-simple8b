package simple8b

import (
	"github.com/golang-infrastructure/go-gtypes"
)

type Mode int

// simple8b的8种模式
const (

	// Mode0BitPacking1 压缩1位整数，每个整数占1位。适用于整数范围为0到1的稀疏数组。
	Mode0BitPacking1 Mode = 0

	// Mode1BitPacking4 压缩4位整数，每个整数占4位。适用于整数范围为0到15的稀疏数组。
	Mode1BitPacking4 Mode = 1

	// Mode2BitPacking8 压缩8位整数，每个整数占8位。适用于整数范围为0到255的稀疏数组。
	Mode2BitPacking8 Mode = 2

	// Mode3BitPacking12 压缩12位整数，每个整数占12位。适用于整数范围为0到4095的稀疏数组。
	Mode3BitPacking12 Mode = 3

	// Mode4BitPacking16 压缩16位整数，每个整数占16位。适用于整数范围为0到65535的稀疏数组。
	Mode4BitPacking16 Mode = 4

	// Mode5BitPacking20 压缩20位整数，每个整数占20位。适用于整数范围为0到1048575的稀疏数组。
	Mode5BitPacking20 Mode = 5

	// Mode6BitPacking24 压缩24位整数，每个整数占24位。适用于整数范围为0到16777215的稀疏数组。
	Mode6BitPacking24 Mode = 6

	// Mode7DirectMode 直接存储整数，没有额外的压缩。适用于密集或者无法被前面7种模式有效压缩的整数数组。
	Mode7DirectMode Mode = 7
)

// WhichMode 判断给定的无符号整数适合用哪种压缩模式
func WhichMode[T gtypes.Unsigned](value T) Mode {
	if value <= 1 {
		return Mode0BitPacking1
	} else if value <= 15 {
		return Mode1BitPacking4
	} else if value <= 255 {
		return Mode2BitPacking8
	} else if value <= 4095 {
		return Mode3BitPacking12
	} else if value <= 65535 {
		return Mode4BitPacking16
	} else if value <= 1048575 {
		return Mode5BitPacking20
	} else if value <= 16777215 {
		return Mode6BitPacking24
	} else {
		return Mode7DirectMode
	}
}

// WhichModeForSlice 判断给定的无符号切片应该使用哪种压缩模式
func WhichModeForSlice[T gtypes.Unsigned](values []T) Mode {
	mode := Mode0BitPacking1
	for _, value := range values {
		m := WhichMode(value)
		// 要遵从切片中最大的那个数字的模式，要能够盛得下它
		if m > mode {
			mode = m
		}
	}
	return mode
}
