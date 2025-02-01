package shouchantypes_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/hujun-open/shouchantypes/v2"
)

func doTest(bk shouchantypes.BinaryKey, outB64 bool) error {

	plainStr := bk.ToStr("")
	fmt.Println(plainStr)
	rplainbk := new(shouchantypes.BinaryKey)
	err := rplainbk.UnmarshalText([]byte(plainStr))
	if err != nil {
		return err
	}
	if !bytes.Equal(*rplainbk, bk) {
		return fmt.Errorf("result %v does equal to orig %v", *rplainbk, bk)
	}

	hexStr := bk.ToStr("0x")
	rbk := new(shouchantypes.BinaryKey)
	err = rbk.UnmarshalText([]byte(hexStr))
	if err != nil {
		return err
	}
	if !bytes.Equal(*rbk, bk) {
		return fmt.Errorf("result %v does equal to orig %v", *rbk, bk)
	}

	b64Str := bk.ToStr("0s")
	r64bk := new(shouchantypes.BinaryKey)
	err = r64bk.UnmarshalText([]byte(b64Str))
	if err != nil {
		return err
	}
	if !bytes.Equal(*r64bk, bk) {
		return fmt.Errorf("result %v does equal to orig %v", *r64bk, bk)
	}
	rstr := bk.String()
	if outB64 != strings.HasPrefix(rstr, "0s") {
		return fmt.Errorf("output format use base64 %v is different from expected %v", strings.HasPrefix(rstr, "0s"), outB64)
	}
	return nil
}

func TestBinaryKey(t *testing.T) {
	type inputStruct struct {
		bk     shouchantypes.BinaryKey
		outb64 bool
	}

	inputList := []inputStruct{

		{
			bk:     shouchantypes.BinaryKey([]byte("abc")),
			outb64: false,
		},

		{
			bk:     shouchantypes.BinaryKey([]byte{1, 2, 3, 4}),
			outb64: true,
		},
	}

	for i, c := range inputList {
		err := doTest(c.bk, c.outb64)
		if err != nil {
			t.Fatalf("case %d failed, %v", i, err)
		}
	}

}
