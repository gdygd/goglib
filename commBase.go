package goglib

const (
	ED_BIG    = 1
	ED_LITTLE = 2
)

var crc32_table []uint32

func GetNumber(src []byte, pos int, length int, endian int) int {
	var value int = 0

	switch endian {
	case ED_BIG:
		for idx := 0; idx < length; idx++ {
			value = (value * 256) + int(src[pos+idx])
		}
		break
	case ED_LITTLE:
		for idx := length - 1; idx >= 0; idx-- {
			value = (value * 256) + int(src[pos+idx])
		}
		break
	}

	return value
}

func SetNumber(buf []byte, pos int, value int, length int, endian int) {
	switch endian {
	case ED_BIG:
		for idx := pos + length - 1; idx >= pos; idx-- {
			buf[idx] = byte(value % 256)
			value /= 256
		}
		break
	case ED_LITTLE:
		for idx := pos; idx < pos+length; idx++ {
			buf[idx] = byte(value % 256)
			value /= 256
		}
		break
	}
}

func SetNumberU(buf []byte, pos int, value uint32, length int, endian int) {
	switch endian {
	case ED_BIG:
		for idx := pos + length - 1; idx >= pos; idx-- {
			buf[idx] = byte(value % 256)
			value /= 256
		}
		break
	case ED_LITTLE:
		for idx := pos; idx < pos+length; idx++ {
			buf[idx] = byte(value % 256)
			value /= 256
		}
		break
	}
}

func GenLRC(buf []byte, pos int, lastIdx int) byte {
	var lrc byte = 0
	for idx := pos; idx < lastIdx; idx++ {
		lrc ^= buf[idx]
	}

	return lrc
}

func InitCrcTb(buf []uint32) {

	crc32_table = make([]uint32, len(buf))

	for idx, v := range buf {
		crc32_table[idx] = v
	}
}

func SetCrc32InitValue(start uint32) uint32 {
	var init uint32 = 0
	var bt uint32
	var val int
	if start > 0 {
		val = 32
		for val > 0 {
			val -= 8
			bt = (start >> val) & 0xFF
			init = crc32_table[((init>>24)^bt)&0xFF] ^ (init << 8)
		}
	}

	return init
}

func UpdateCrc(crc_accum uint32, pData []byte, len int) uint32 {
	var idx int = 0

	for j := 0; j < len; j++ {
		//i = ((int)(crc_accum >> 24) ^ *data_blk_ptr++) & 0xff

		i := ((int)(crc_accum>>24) ^ int(pData[idx])) & 0xff
		idx++
		crc_accum = (crc_accum << 8) ^ crc32_table[i]

	}

	return crc_accum
}

func CheckCrc(data []byte, len int, rcvcrc uint32) bool {
	crc := UpdateCrc(SetCrc32InitValue(0xFFFFFFFF), data, len)

	if crc == rcvcrc {
		return true
	}
	return false
}

// UpdateCrc(SetCrc32InitValue(0xFFFFFFFF), vuf, len)
