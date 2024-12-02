package simple8b

import "github.com/golang-infrastructure/go-gtypes"

// IntToBytes 把给定的整数的低n位转换为字节数组
func IntToBytes[T gtypes.Integer](value T, blockSize int) []byte {
	result := make([]byte, blockSize)
	for index := range result {
		byteValue := (uint64(0xFF) << index) & uint64(value)
		result[index] = uint8(byteValue)
	}
	return result
}

// BytesToInt 把字节转为整数
func BytesToInt[T gtypes.Integer](bytes []byte) T {
	var r uint64
	weight := 0
	for _, x := range bytes {
		r = r | (uint64(x) << weight)
		weight += 8
	}
	return T(r)
}
