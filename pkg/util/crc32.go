package util

import "hash/crc32"

func Crc32(str string) uint32 {
	return crc32.Checksum([]byte(str), crc32.IEEETable)
}
