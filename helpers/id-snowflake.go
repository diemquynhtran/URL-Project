package helpers

import (
	"errors"
	"learning-go/snowflake"
	"math"
	"strings"
)

var (
	Base         uint64 = 62
	CharacterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func IdBase62(id uint64) string {
	return Encode(id)
}
func IdBase10(id string) uint64 {
	a, _ := Decode(id)
	return uint64(a)
}

//convert base10 to base62
func Encode(num uint64) string {
	b := make([]byte, 0)

	// loop as long the num is bigger than zero
	for num > 0 {
		// receive the rest
		r := math.Mod(float64(num), float64(Base))

		num /= Base

		// append chars
		b = append([]byte{CharacterSet[int(r)]}, b...)
	}

	return string(b)
}

//convert base62 to base10
func Decode(s string) (int, error) {
	var r, pow int

	// loop through the input
	for i, v := range s {
		// convert position to power
		pow = len(s) - (i + 1)

		// IndexRune returns -1 if v is not part of CharacterSet.
		pos := strings.IndexRune(CharacterSet, v)

		if pos == -1 {
			return pos, errors.New("invalid character: " + string(v))
		}

		// calculate
		r += pos * int(math.Pow(float64(Base), float64(pow)))
	}

	return int(r), nil
}

func CreateID() uint64 {
	snowflake.SetEpoch(1577836800) // timestamp start from 01/01/2020 @ 12:00am (UTC)
	snowflake.SetNodeBits(10)      // reserve 2^10 machine numbers
	snowflake.SetSeqBits(12)       // 2^12 unique ID per second
	_ = snowflake.Init()           // must be called to instantiate new config
	nodeID := uint64(1)

	node, _ := snowflake.CreateNode(nodeID)
	id := node.GenerateID()
	return id
}

/*func CreateID() (id snowflake.ID, e error) {
	node, err := snowflake.NewNode(1)
	e = err
	if e != nil {
		return 0, err
	}

	id = node.Generate()

	return id, nil
}*/

/*func IdToString() string {
	id, _ := CreateID()
	return id.String()
}*/
