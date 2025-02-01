package shouchantypes

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// BinaryKey is a byte slice
type BinaryKey []byte

// 0x for hex, 0s for base64, no prefix or less than 3 bytes copy the text as byte slice
func (bk *BinaryKey) UnmarshalText(text []byte) error {

	var err error
	var buflen int
	var buf []byte
	if len(text) == 0 {
		bk = nil
		return nil
	}
	switch strings.ToLower(string(text[:2])) {
	case "0x":
		if len(text) < 4 {
			return fmt.Errorf("%v is not a valid hex encoded string", string(text))
		}
		buflen = (len(text) - 2) / 2
		buf = make([]byte, buflen)
		_, err = hex.Decode(buf, text[2:])
	case "0s":
		if len(text) < 3 {
			return fmt.Errorf("%v is not a valid base64 encoded string", string(text))
		}
		buf, err = base64.StdEncoding.DecodeString(string(text[2:]))
		buflen = len(buf)
	default:
		*bk = make([]byte, len(text))
		copy(*bk, text)
		return nil

	}
	if err == nil {
		*bk = make([]byte, buflen)
		copy(*bk, buf)
	}
	return err
}

// if fmt is "0s", return bk encoded as base64;
// if fmt is "0x", return bk encoded as hex;
// otherwise, return string of bk
func (bk *BinaryKey) ToStr(fmt string) string {
	switch strings.ToLower(fmt) {
	case "0s":
		return "0s" + base64.StdEncoding.EncodeToString(*bk)
	case "0x":
		return "0x" + hex.EncodeToString(*bk)

	}
	return string(*bk)
}

func (bk BinaryKey) MarshalText() (text []byte, err error) {
	return []byte(bk.String()), nil

}

// if bk only contains non-space, printable ASCII chars, then output is plan text,
// otherwise output is base64 encoded string
func (bk BinaryKey) String() string {
	for _, ch := range bk {
		if ch < 33 || ch > 126 {
			return bk.ToStr("0s")
		}
	}
	return bk.ToStr("")
}
