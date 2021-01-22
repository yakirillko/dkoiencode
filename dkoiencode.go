package dkoiencode

/*var DKOItoUnicodefull = map[byte]rune{
	0x00: 0x0000, 0x01: 0x0001, 0x02: 0x0002, 0x03: 0x0003, 0x04: 0x009C, 0x05: 0x0009, 0x06: 0x0086, 0x07: 0x007F, 0x08: 0x0097, 0x09: 0x008D, 0x0A: 0x008E, 0x0B: 0x000B, 0x0C: 0x000C, 0x0D: 0x000D, 0x0E: 0x000E, 0x0F: 0x000F,
	0x10: 0x0010, 0x11: 0x0011, 0x12: 0x0012, 0x13: 0x0013, 0x14: 0x009D, 0x15: 0x0085, 0x16: 0x0008, 0x17: 0x0087, 0x18: 0x0018, 0x19: 0x0019, 0x1A: 0x0092, 0x1B: 0x008F, 0x1C: 0x001C, 0x1D: 0x001D, 0x1E: 0x001E, 0x1F: 0x001F,
	0x20: 0x0080, 0x21: 0x0081, 0x22: 0x0082, 0x23: 0x0000, 0x24: 0x0084, 0x25: 0x000A, 0x26: 0x0017, 0x27: 0x001B, 0x28: 0x0000, 0x29: 0x0000, 0x2A: 0x008A, 0x2B: 0x008B, 0x2C: 0x0000, 0x2D: 0x0005, 0x2E: 0x0006, 0x2F: 0x0007,
	0x30: 0x0000, 0x31: 0x0000, 0x32: 0x0016, 0x33: 0x0000, 0x34: 0x0094, 0x35: 0x0095, 0x36: 0x0096, 0x37: 0x0004, 0x38: 0x0000, 0x39: 0x0000, 0x3A: 0x0000, 0x3B: 0x009B, 0x3C: 0x0014, 0x3D: 0x0015, 0x3E: 0x0000, 0x3F: 0x001A,
	0x40: 0x0020, 0x41: 0x0000, 0x42: 0x0000, 0x43: 0x0000, 0x44: 0x0000, 0x45: 0x0000, 0x46: 0x0000, 0x47: 0x0000, 0x48: 0x0000, 0x49: 0x0000, 0x4A: 0x005B, 0x4B: 0x002E, 0x4C: 0x003C, 0x4D: 0x0028, 0x4E: 0x002B, 0x4F: 0x0021,
	0x50: 0x0026, 0x51: 0x0000, 0x52: 0x0000, 0x53: 0x0000, 0x54: 0x0000, 0x55: 0x0000, 0x56: 0x0000, 0x57: 0x0000, 0x58: 0x0000, 0x59: 0x0000, 0x5A: 0x005D, 0x5B: 0x00A4, 0x5C: 0x002A, 0x5D: 0x0029, 0x5E: 0x003B, 0x5F: 0x00AC,
	0x60: 0x002D, 0x61: 0x002F, 0x62: 0x0000, 0x63: 0x0000, 0x64: 0x0000, 0x65: 0x0000, 0x66: 0x0000, 0x67: 0x0000, 0x68: 0x0000, 0x69: 0x0000, 0x6A: 0x007C, 0x6B: 0x002C, 0x6C: 0x0025, 0x6D: 0x005F, 0x6E: 0x003E, 0x6F: 0x003F,
	0x70: 0x0000, 0x71: 0x0000, 0x72: 0x0000, 0x73: 0x0000, 0x74: 0x0000, 0x75: 0x0000, 0x76: 0x044E, 0x77: 0x0430, 0x78: 0x0431, 0x79: 0x0060, 0x7A: 0x003A, 0x7B: 0x0023, 0x7C: 0x0040, 0x7D: 0x0027, 0x7E: 0x003D, 0x7F: 0x0022,
	0x80: 0x0446, 0x81: 0x0061, 0x82: 0x0062, 0x83: 0x0063, 0x84: 0x0064, 0x85: 0x0065, 0x86: 0x0066, 0x87: 0x0067, 0x88: 0x0068, 0x89: 0x0069, 0x8A: 0x0434, 0x8B: 0x0435, 0x8C: 0x0444, 0x8D: 0x0433, 0x8E: 0x0445, 0x8F: 0x0438,
	0x90: 0x0439, 0x91: 0x006A, 0x92: 0x006B, 0x93: 0x006C, 0x94: 0x006D, 0x95: 0x006E, 0x96: 0x006F, 0x97: 0x0070, 0x98: 0x0071, 0x99: 0x0072, 0x9A: 0x043A, 0x9B: 0x043B, 0x9C: 0x043C, 0x9D: 0x043D, 0x9E: 0x043E, 0x9F: 0x043F,
	0xA0: 0x044F, 0xA1: 0x007E, 0xA2: 0x0073, 0xA3: 0x0074, 0xA4: 0x0075, 0xA5: 0x0076, 0xA6: 0x0077, 0xA7: 0x0078, 0xA8: 0x0079, 0xA9: 0x007A, 0xAA: 0x0440, 0xAB: 0x0441, 0xAC: 0x0442, 0xAD: 0x0443, 0xAE: 0x0436, 0xAF: 0x0432,
	0xB0: 0x044C, 0xB1: 0x044B, 0xB2: 0x0437, 0xB3: 0x0448, 0xB4: 0x044D, 0xB5: 0x0449, 0xB6: 0x0447, 0xB7: 0x044A, 0xB8: 0x042E, 0xB9: 0x0410, 0xBA: 0x0411, 0xBB: 0x0426, 0xBC: 0x0414, 0xBD: 0x0415, 0xBE: 0x0424, 0xBF: 0x0413,
	0xC0: 0x007B, 0xC1: 0x0041, 0xC2: 0x0042, 0xC3: 0x0043, 0xC4: 0x0044, 0xC5: 0x0045, 0xC6: 0x0046, 0xC7: 0x0047, 0xC8: 0x0048, 0xC9: 0x0049, 0xCA: 0x0425, 0xCB: 0x0418, 0xCC: 0x0419, 0xCD: 0x041A, 0xCE: 0x041B, 0xCF: 0x041C,
	0xD0: 0x007D, 0xD1: 0x004A, 0xD2: 0x004B, 0xD3: 0x004C, 0xD4: 0x004D, 0xD5: 0x004E, 0xD6: 0x004F, 0xD7: 0x0050, 0xD8: 0x0051, 0xD9: 0x0052, 0xDA: 0x041D, 0xDB: 0x041E, 0xDC: 0x041F, 0xDD: 0x042F, 0xDE: 0x0420, 0xDF: 0x0421,
	0xE0: 0x005C, 0xE1: 0x0000, 0xE2: 0x0053, 0xE3: 0x0054, 0xE4: 0x0055, 0xE5: 0x0056, 0xE6: 0x0057, 0xE7: 0x0058, 0xE8: 0x0059, 0xE9: 0x005A, 0xEA: 0x0422, 0xEB: 0x0423, 0xEC: 0x0416, 0xED: 0x0412, 0xEE: 0x042C, 0xEF: 0x042B,
	0xF0: 0x0030, 0xF1: 0x0031, 0xF2: 0x0032, 0xF3: 0x0033, 0xF4: 0x0034, 0xF5: 0x0035, 0xF6: 0x0036, 0xF7: 0x0037, 0xF8: 0x0038, 0xF9: 0x0039, 0xFA: 0x0417, 0xFB: 0x0428, 0xFC: 0x042D, 0xFD: 0x0429, 0xFE: 0x0427, 0xFF: 0x009F,
}*/

var UnicodetoDKOI = map[rune]byte{
	0x76: 0xA5, 0x437: 0xB2, 0x30: 0xF0, 0x64: 0x84, 0x44: 0xC4, 0x4B: 0xD2, 0x4: 0x37, 0x52: 0xD9, 0x42D: 0xFC, 0xB: 0xB, 0xF: 0xF, 0x27: 0x7D,
	0x426: 0xBB, 0x42C: 0xEE, 0x82: 0x22, 0x21: 0x4F, 0x44B: 0xB1, 0x4C: 0xD3, 0x53: 0xE2, 0x2A: 0x5C, 0x71: 0x98, 0x43: 0xC3, 0x1: 0x1,
	0x6B: 0x92, 0x72: 0x99, 0x58: 0xE7, 0x433: 0x8D, 0x432: 0xAF, 0x41C: 0xCF, 0x34: 0xF4, 0x92: 0x1A, 0x75: 0xA4, 0x56: 0xE5, 0xE: 0xE,
	0x28: 0x4D, 0x43E: 0x9E, 0x7D: 0xD0, 0x417: 0xFA, 0x12: 0x12, 0x69: 0x89, 0x4E: 0xD5, 0x41E: 0xDB, 0x9C: 0x4, 0x6A: 0x91, 0x413: 0xBF,
	0x84: 0x24, 0x96: 0x36, 0x46: 0xC6, 0x41A: 0xCD, 0x51: 0xD8, 0x435: 0x8B, 0x43B: 0x9B, 0x422: 0xEA, 0x47: 0xC7, 0x48: 0xC8, 0x421: 0xDF,
	0x5F: 0x6D, 0x6D: 0x94, 0x423: 0xEB, 0x10: 0x10, 0x20: 0x40, 0x23: 0x7B, 0x73: 0xA2, 0x59: 0xE8, 0x18: 0x18, 0x1D: 0x1D, 0x78: 0xA7,
	0x429: 0xFD, 0x16: 0x32, 0x440: 0xAA, 0x44C: 0xB0, 0x29: 0x5D, 0x40: 0x7C, 0x411: 0xBA, 0x50: 0xD7, 0x38: 0xF8, 0x3: 0x3, 0xD: 0xD, 0x8F: 0x1B,
	0x439: 0x90, 0xA4: 0x5B, 0x7C: 0x6A, 0x66: 0x86, 0x31: 0xF1, 0x6F: 0x96, 0x33: 0xF3, 0x1A: 0x3F, 0x5B: 0x4A, 0x415: 0xBD, 0xAC: 0x5F, 0x444: 0x8C,
	0x443: 0xAD, 0x8E: 0xA, 0x44E: 0x76, 0x419: 0xCC, 0x13: 0x13, 0x8B: 0x2B, 0x446: 0x80, 0x5A: 0xE9, 0x39: 0xF9, 0x9B: 0x3B, 0x448: 0xB3, 0x447: 0xB6,
	0x42: 0xC2, 0x85: 0x15, 0x8: 0x16, 0x3E: 0x6E, 0x6E: 0x95, 0x2: 0x2, 0x1E: 0x1E, 0x95: 0x35, 0x22: 0x7F, 0x57: 0xE6, 0x42B: 0xEF, 0x1B: 0x27,
	0x3D: 0x7E, 0x70: 0x97, 0x43F: 0x9F, 0x424: 0xBE, 0x7F: 0x7, 0x3C: 0x4C, 0x77: 0xA6, 0x449: 0xB5, 0x42E: 0xB8, 0x41B: 0xCE, 0x9D: 0x14, 0x8A: 0x2A,
	0x5D: 0x5A, 0x2D: 0x60, 0x79: 0xA8, 0x44D: 0xB4, 0x87: 0x17, 0x1C: 0x1C, 0x60: 0x79, 0x441: 0xAB, 0x414: 0xBC, 0x37: 0xF7, 0xA: 0x25, 0x2C: 0x6B,
	0x44F: 0xA0, 0x410: 0xB9, 0x49: 0xC9, 0x8D: 0x9, 0x6: 0x2E, 0x5: 0x2D, 0x7E: 0xA1, 0x45: 0xC5, 0x7: 0x2F, 0x25: 0x6C, 0x74: 0xA3, 0x54: 0xE3, 0x427: 0xFE,
	0x2E: 0x4B, 0x26: 0x50, 0x61: 0x81, 0x434: 0x8A, 0x6C: 0x93, 0x418: 0xCB, 0x7A: 0xA9, 0x36: 0xF6, 0x431: 0x78, 0x68: 0x88, 0x94: 0x34, 0x436: 0xAE,
	0x7B: 0xC0, 0x42F: 0xDD, 0x1F: 0x1F, 0x9: 0x5, 0x62: 0x82, 0x41D: 0xDA, 0x97: 0x8, 0x3B: 0x5E, 0x63: 0x83, 0x2B: 0x4E, 0x43C: 0x9C, 0x55: 0xE4, 0x86: 0x6,
	0x81: 0x21, 0x2F: 0x61, 0x43A: 0x9A, 0x4D: 0xD4, 0xC: 0xC, 0x3A: 0x7A, 0x9F: 0xFF, 0x67: 0x87, 0x4A: 0xD1, 0x41F: 0xDC, 0x420: 0xDE, 0x412: 0xED, 0x442: 0xAC,
	0x32: 0xF2, 0x35: 0xF5, 0x14: 0x3C, 0x15: 0x3D, 0x445: 0x8E, 0x425: 0xCA, 0x4F: 0xD6, 0x416: 0xEC, 0x19: 0x19, 0x17: 0x26, 0x3F: 0x6F, 0x41: 0xC1,
	0x428: 0xFB, 0x11: 0x11, 0x430: 0x77, 0x44A: 0xB7, 0x5C: 0xE0, 0x438: 0x8F, 0x43D: 0x9D, 0x80: 0x20, 0x65: 0x85, 0x0: 0x0,
}
var DKOItoUnicode = map[byte]rune{
	0x92: 0x6B, 0xEF: 0x42B, 0xFD: 0x429, 0x18: 0x18, 0x6F: 0x3F, 0xA6: 0x77, 0xDA: 0x41D, 0xE7: 0x58, 0x1: 0x1, 0xED: 0x412,
	0xEB: 0x423, 0x91: 0x6A, 0xAB: 0x441, 0xAC: 0x442, 0xD8: 0x51, 0xE8: 0x59, 0x19: 0x19, 0xE5: 0x56, 0x93: 0x6C, 0xA0: 0x44F,
	0xD9: 0x52, 0x81: 0x61, 0xA9: 0x7A, 0x61: 0x2F, 0x80: 0x446, 0x7A: 0x3A, 0x1F: 0x1F, 0x2A: 0x8A, 0x78: 0x431, 0x7F: 0x22,
	0x9B: 0x43B, 0x26: 0x17, 0xF4: 0x34, 0x5E: 0x3B, 0x9D: 0x43D, 0x4: 0x9C, 0x96: 0x6F, 0xCD: 0x41A, 0x32: 0x16, 0x9C: 0x43C,
	0xE2: 0x53, 0x97: 0x70, 0xF5: 0x35, 0x2E: 0x6, 0x21: 0x81, 0x6A: 0x7C, 0x90: 0x439, 0x94: 0x6D, 0xE0: 0x5C, 0x10: 0x10, 0x86: 0x66,
	0x8B: 0x435, 0xBF: 0x413, 0xF0: 0x30, 0x34: 0x94, 0x95: 0x6E, 0xB0: 0x44C, 0xFF: 0x9F, 0x6: 0x86, 0x5C: 0x2A, 0x84: 0x64,
	0xB4: 0x44D, 0xC8: 0x48, 0x27: 0x1B, 0x1C: 0x1C, 0x22: 0x82, 0x35: 0x95, 0x5F: 0xAC, 0xB2: 0x437, 0xF1: 0x31, 0x8E: 0x445,
	0xB5: 0x449, 0x3B: 0x9B, 0x40: 0x20, 0x4B: 0x2E, 0xC7: 0x47, 0x13: 0x13, 0xD5: 0x4E, 0xB7: 0x44A, 0xA2: 0x73, 0xF7: 0x37,
	0x7: 0x7F, 0x16: 0x8, 0x5: 0x9, 0x1E: 0x1E, 0x37: 0x4, 0x6B: 0x2C, 0x89: 0x69, 0xB3: 0x448, 0xEA: 0x422, 0x17: 0x87, 0x4E: 0x2B,
	0xD4: 0x4D, 0xDC: 0x41F, 0xDE: 0x420, 0x5A: 0x5D, 0xB6: 0x447, 0x20: 0x80, 0x14: 0x9D, 0x9F: 0x43F, 0xC6: 0x46, 0xFA: 0x417,
	0x8: 0x97, 0x88: 0x68, 0xAA: 0x440, 0xCA: 0x425, 0x9: 0x8D, 0x4D: 0x28, 0x8C: 0x444, 0xA4: 0x75, 0xFE: 0x427, 0xC: 0xC, 0x82: 0x62,
	0x99: 0x72, 0xAD: 0x443, 0xC1: 0x41, 0xD0: 0x7D, 0xE6: 0x57, 0xD: 0xD, 0xFC: 0x42D, 0xF6: 0x36, 0x4C: 0x3C, 0x6C: 0x25, 0xBC: 0x414,
	0x3C: 0x14, 0xCC: 0x419, 0xFB: 0x428, 0x85: 0x65, 0x7E: 0x3D, 0x87: 0x67, 0x9E: 0x43E, 0xD1: 0x4A, 0x50: 0x26, 0xB8: 0x42E,
	0x4F: 0x21, 0x79: 0x60, 0xC5: 0x45, 0xD6: 0x4F, 0x2: 0x2, 0xDF: 0x421, 0x1D: 0x1D, 0x6D: 0x5F, 0x9A: 0x43A, 0xF: 0xF, 0xE3: 0x54,
	0xEC: 0x416, 0xBA: 0x411, 0xA8: 0x79, 0xCF: 0x41C, 0x5D: 0x29, 0x11: 0x11, 0x7B: 0x23, 0x8F: 0x438, 0xA1: 0x7E, 0xE: 0xE,
	0xB1: 0x44B, 0xC4: 0x44, 0x24: 0x84, 0x4A: 0x5B, 0x8D: 0x433, 0xF9: 0x39, 0x3D: 0x15, 0x5B: 0xA4, 0xBD: 0x415, 0xC3: 0x43,
	0xD2: 0x4B, 0xD3: 0x4C, 0xE4: 0x55, 0x3: 0x3, 0x76: 0x44E, 0x2D: 0x5, 0xCB: 0x418, 0x6E: 0x3E, 0xA5: 0x76, 0xAF: 0x432, 0xEE: 0x42C,
	0xF3: 0x33, 0xF8: 0x38, 0xCE: 0x41B, 0x25: 0xA, 0x2B: 0x8B, 0x3F: 0x1A, 0x98: 0x71, 0xDB: 0x41E, 0xF2: 0x32, 0x1B: 0x8F, 0x77: 0x430,
	0x83: 0x63, 0xC0: 0x7B, 0xDD: 0x42F, 0xA: 0x8E, 0x36: 0x96, 0xA3: 0x74, 0xBB: 0x426, 0xC2: 0x42, 0xC9: 0x49, 0xD7: 0x50, 0x2F: 0x7,
	0x60: 0x2D, 0xE9: 0x5A, 0x15: 0x85, 0xBE: 0x424, 0x12: 0x12, 0xA7: 0x78, 0xB9: 0x410, 0xB: 0xB, 0x8A: 0x434, 0xAE: 0x436, 0x7C: 0x40, 0x7D: 0x27, 0x1A: 0x92, 0x0: 0x0,
}

func Decode(buf *[]byte) []rune {
	var out []rune
	var v byte

	for _, v = range *buf {

		if _, ok := DKOItoUnicode[v]; ok {
			out = append(out, DKOItoUnicode[v])
		} else {
			out = append(out, 0x00)
		}
	}
	return out
}
func Encode(buf *[]rune) []byte {
	var out []byte
	var v rune
	for _, v = range *buf {
		out = append(out, UnicodetoDKOI[v])
	}
	return out
}
