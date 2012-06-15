
package tools

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz"
var characters_len *big.Int

type BigInt big.Int
func (self *BigInt) BaseString(base int) string {
	return self.BaseStringBig(NewBigIntInt(base))
}
func (self *BigInt) BaseStringBig(base *BigInt) string {
	if self.MathInt().Cmp(base.MathInt()) < 0 {
		return string(characters[self.Int()])
	}
	rval := NewBigIntInt(0)
	rest := NewBigIntInt(0)
	rest.MathInt().SetBytes(self.MathInt().Bytes())
	rest.MathInt().DivMod(rest.MathInt(), base.MathInt(), rval.MathInt())
	return fmt.Sprintf("%s%s", rest.BaseStringBig(base), string(characters[rval.Int()]))
}
func (self *BigInt) MathInt() *big.Int {
	return (*big.Int)(self)
}
func (self *BigInt) Int() int {
	return int(self.MathInt().Int64())
}
func NewBigIntBytes(bytes []byte) *BigInt {
	rval := big.NewInt(int64(0))
	rval.SetBytes(bytes)
	return (*BigInt)(rval)
}
func NewBigIntInt(i int) *BigInt {
	return (*BigInt)(big.NewInt(int64(i)))
}

func init() {
	characters_len = big.NewInt(int64(len(characters)))
}

func String(i int64, base int) {
}

func Uuid() string {
	upper := time.Now().Unix()
	lower := int32(upper)
	return string([]int32{int32(upper >> 32), lower})
}

func RandomString(l int) string {
	buffer := bytes.NewBufferString("")
	for i := 0; i < l; i++ {
		x, err := rand.Int(rand.Reader, characters_len)
		if err != nil {
			panic(fmt.Sprint("Unable to create random string: ", err))
		}
		fmt.Fprintf(buffer, "%s", string(characters[int(x.Int64())]))
	}
	return string(buffer.Bytes())
}