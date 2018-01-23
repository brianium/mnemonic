package entropy

import (
	"reflect"
	"testing"
)

func TestFromHex(t *testing.T) {
	ent, err := FromHex("00")
	if err != nil {
		t.Errorf("\n%s", err)
	}
	expected := "00000000"
	actual := string(bytesToBits(ent))
	if actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s", expected, actual)
	}
}

func TestCheckSum(t *testing.T) {
	ent, err := FromHex("00000000000000000000000000000000")
	if err != nil {
		t.Errorf("\n%s", err)
	}
	actual := string(CheckSum(ent))
	expected := string([]byte{'0', '0', '1', '1'})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %s\nActual: %s", expected, actual)
	}
}

func TestRandomTooLow(t *testing.T) {
	// length too low
	_, err := Random(127)
	if err == nil {
		t.Error("Expected length of 127 to return error")
	}
}

func TestRandomTooHigh(t *testing.T) {
	// length too low
	_, err := Random(257)
	if err == nil {
		t.Error("Expected length of 257 to return error")
	}
}

func TestRandomNotDivisibleBy32(t *testing.T) {
	// length too low
	_, err := Random(129)
	if err == nil {
		t.Error("Expected length of 129 to return error")
	}
}

func TestRandom(t *testing.T) {
	length := 128
	b, err := Random(length)
	if err != nil {
		t.Errorf("Random entropy failed %s", err)
	}
	if len(b) != length/32 {
		t.Errorf("Expeced byte length of %d", length/32)
	}
}
