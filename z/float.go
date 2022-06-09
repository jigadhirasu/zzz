package z

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NewFloat(data []byte) (Float, error) {
	data = bytes.ReplaceAll(data, []byte(`\"`), []byte(``))
	data = bytes.ReplaceAll(data, []byte(`"`), []byte(``))
	in := string(data)
	if !floatMatcher.MatchString(in) {
		return 0, fmt.Errorf(`str must match : "%s"`, regexpFloat)
	}

	ln := 9
	if idx := strings.Index(in, "."); idx > -1 {
		ln -= len(in[idx+1:])
	}

	for i := 0; i < ln; i++ {
		in = in + "0"
	}

	r := strings.ReplaceAll(in, ".", "")
	out, err := strconv.ParseInt(r, 10, 64)
	return Float(out), err
}

// 所有跟金額有關的整數需使用該倍率
const FloatUnit Float = 1000000000

const regexpFloat = "^-?[0-9]+.?[0-9]{0,9}$"

var floatMatcher = regexp.MustCompile(regexpFloat)

type Float int64

func (coin Float) MarshalJSON() ([]byte, error) {
	float := coin.String()
	if float == "" {
		float = "0"
	}
	return json.Marshal(float)
}

func (coin *Float) UnmarshalJSON(data []byte) error {
	float, err := NewFloat(data)
	*coin = float
	return err
}

func (coin Float) String() string {
	dash := ""
	if coin < 0 {
		dash = "-"
		coin *= -1
	}
	s := fmt.Sprintf("%010d", coin)
	ln := len(s)
	dot := ln - 9

	if s[dot:] == "000000000" {
		return dash + s[:dot]
	}

	rs := strings.Split(s[dot:], "")

	for i := 9; i > 0; i-- {
		if rs[i-1] == "0" {
			continue
		}

		return dash + s[:dot] + "." + s[dot:dot+i]
	}

	return dash + s[:dot] + "." + s[dot:]
}

func (coin Float) Int64() int64 {
	return int64(coin)
}
