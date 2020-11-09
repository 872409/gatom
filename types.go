package gatom

import "strconv"

type SInt string

func (receiver SInt) Int64() int64 {
	_val, err := strconv.ParseInt(string(receiver), 10, 64)
	if err != nil {
		return 0
	}
	return _val
}

func (receiver SInt) String() string {
	return string(receiver)
}

func ParseSInt(sint int64) SInt {
	return SInt(strconv.FormatInt(sint, 10))
}
