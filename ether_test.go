package etherUtils

import (
	"fmt"
	"math/big"
	"testing"
)

type testData struct {
	Str string
	Ui  int64
	Ok  bool
}

func TestStrToEther(t *testing.T) {
	testData := []testData{
		{Str: "1.0", Ui: 1000000000000000000, Ok: true},
		{Str: "0.000000000000000001", Ui: 1, Ok: true},
		{Str: "0.0000000000000000001", Ui: 0, Ok: false},
		{Str: "1.00", Ui: 1000000000000000000, Ok: true},
		{Str: "1.", Ui: 1000000000000000000, Ok: true},
		{Str: "1", Ui: 1000000000000000000, Ok: true},
	}
	for _, rec := range testData {
		fmt.Println(rec.Str)
		val, ok := StrToEther(rec.Str)
		if !ok && rec.Ok {
			t.Error("cannot convert ", rec.Str)
		}
		if !ok && !rec.Ok {
			continue
		}
		if !rec.Ok {
			t.Error("False conversion ", rec.Str, val.String())
		}
		num := big.NewInt(rec.Ui)
		if val.Cmp(num) != 0 {
			t.Errorf("StrToEther test %d,%d", val, num)
		}
	}
}

type testData2 struct {
	Str string
	Ui  int64
	Dec int
	Ok  bool
}

func TestStrToDecimals(t *testing.T) {
	testData := []testData2{
		{Str: "1.0", Ui: 1, Dec: 0, Ok: false},
		{Str: "1", Ui: 1, Dec: 0, Ok: true},
		{Str: "1", Ui: 10, Dec: 1, Ok: true},
		{Str: "1", Ui: 100, Dec: 2, Ok: true},
		{Str: "1", Ui: 1000, Dec: 3, Ok: true},
		{Str: "1", Ui: 10000, Dec: 4, Ok: true},
		{Str: "1", Ui: 100000, Dec: 5, Ok: true},
		{Str: "1000000000000000000", Ui: 1000000000000000000, Dec: 0, Ok: true},
		{Str: "1", Ui: 1000000000000000000, Dec: 18, Ok: true},
	}
	for _, rec := range testData {
		val, ok := StrToDecimals(rec.Str, rec.Dec)
		if !ok && rec.Ok {
			t.Error("cannot convert ", rec.Str)
		}
		if !ok && !rec.Ok {
			continue
		}
		if !rec.Ok {
			t.Error("False conversion ", rec.Str)
		}
		num := big.NewInt(rec.Ui)
		if val.Cmp(num) != 0 {
			t.Errorf("StrToEther test %d,%d", val, num)
		}
	}
}
