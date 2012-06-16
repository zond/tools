
package tools

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var TimingMap = make(map[string]int64)

func TimeIn(s string) {
	TimingMap[s] = TimingMap[s] - time.Now().Unix()
}
func TimeOut(s string) {
	TimingMap[s] = TimingMap[s] + time.Now().Unix()
}

const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz"
const MAX_BASE = 61

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
func NewBigIntString(s string, base int) *BigInt {
	rval := big.NewInt(int64(0))
	rval.SetString(s, base)
	return (*BigInt)(rval)
}
func NewBigIntBytes(bytes []byte) *BigInt {
	rval := big.NewInt(int64(0))
	rval.SetBytes(bytes)
	return (*BigInt)(rval)
}
func NewBigIntInt(i int) *BigInt {
	return (*BigInt)(big.NewInt(int64(i)))
}
func NewBigIntInt64(i int64) *BigInt {
	return (*BigInt)(big.NewInt(i))
}

func Uuid() string {
	timePart := NewBigIntInt64(time.Now().Unix())
	return fmt.Sprint(timePart.BaseString(MAX_BASE), RandomString(10))
}

func RandomString(l int) string {
	buffer := bytes.NewBufferString("")
	for i := 0; i < l; i++ {
		x, err := rand.Int(rand.Reader, big.NewInt(MAX_BASE))
		if err != nil {
			panic(fmt.Sprint("Unable to create random string: ", err))
		}
		fmt.Fprintf(buffer, "%s", string(characters[int(x.Int64())]))
	}
	return string(buffer.Bytes())
}