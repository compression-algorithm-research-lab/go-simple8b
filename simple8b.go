package simple8b

import (
	"github.com/compression-algorithm-research-lab/go-zigzag"
	"github.com/golang-infrastructure/go-gtypes"
)

// CanZipToBits 判断可以压缩到几个bit
func CanZipToBits[T gtypes.Integer](slice []T) int {
	var sum T
	for _, value := range slice {
		sum |= zigzag.ToZigZag(value)
	}
	bitCount := 0
	for sum > 0 {
		bitCount++
		sum >>= 1
	}
	return bitCount
}

// Encode 编码
func Encode[T gtypes.Integer](slice []T) []byte {

	result := make([]byte, 0)

	// 先计算一下最少可以压缩到几位
	bits := CanZipToBits(slice)
	blockSize := (bits + 7) / 8
	result = append(result, uint8(blockSize))

	// 然后就按照这个来存储了
	for _, value := range slice {
		// 结合zigzag算法，可以用来压缩负数
		result = append(result, IntToBytes(zigzag.ToZigZag(value), blockSize)...)
	}

	return result
}

// DecodeE 解码
func DecodeE[T gtypes.Integer](bytes []byte) ([]T, error) {

	// 先读取是一个块
	if len(bytes) == 0 {
		return nil, ErrFormatNotOk
	}
	blockSize := BytesToInt[int]([]byte{bytes[0]})
	// 进行长度校验，康康块是不是都是合法的
	if (len(bytes)-1)%blockSize != 0 {
		return nil, ErrFormatNotOk
	}

	// 然后就一块一块的读取
	result := make([]T, 0)
	count := (len(bytes) - 1) / blockSize
	for i := 0; i < count; i++ {
		valueIndex := 1 + i*blockSize
		valueBytes := bytes[valueIndex : valueIndex+blockSize]
		result = append(result, zigzag.FromToZigZag(BytesToInt[T](valueBytes)))
	}
	return result, nil
}

//func EncodeToBytes[T gtypes.Integer](slice []T) []byte {
//
//}
//
//func DecodeFromBytes[T gtypes.Integer](bytes []byte) ([]T, error) {
//
//}
