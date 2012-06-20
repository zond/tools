
package tools

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"time"
)

var timeMap = make(map[string]int64)
var countMap = make(map[string]int64)

func Pad(s interface{}, p interface{}, min int) string {
	buffer := bytes.NewBufferString(fmt.Sprint(s))
	for buffer.Len() < min {
		buffer.Write([]byte(fmt.Sprint(p)))
	}
	return string(buffer.Bytes())
}

type ProfData struct {
	name string
	count int64
	spent int64
}
func NewProfData(s string) *ProfData {
	return &ProfData{s, countMap[s], timeMap[s]}
}
func (self *ProfData) String(p int) string {
	return fmt.Sprint(Pad(self.name, " ", p), Pad(self.Avg(), " ", p), Pad(self.count, " ", p))
}
func (self *ProfData) Avg() int64 {
	return self.spent / self.count
}

type ProfDataSlice []*ProfData
func (self ProfDataSlice) Len() int { return len(self) }
func (self ProfDataSlice) Less(i, j int) bool { return self[i].Avg() < self[j].Avg() }
func (self ProfDataSlice) Swap(i, j int) { self[i], self[j] = self[j], self[i] }

func TimeClear() {
	timeMap = make(map[string]int64)
	countMap = make(map[string]int64)
}
func TimeIn(s string) {
	timeMap[s] = timeMap[s] - time.Now().UnixNano()
	countMap[s] = countMap[s] + 1
}
func TimeOut(s string) {
	timeMap[s] = timeMap[s] + time.Now().UnixNano()
}
func Prof(pad int) []*ProfData {
	var rval ProfDataSlice
	for s, _ := range timeMap {
		rval = append(rval, NewProfData(s))
	}
	sort.Sort(rval)
	return rval
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
func (self *BigInt) Bytes() []byte {
	return self.MathInt().Bytes()
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