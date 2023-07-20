package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	var key = "iwrupvqb"

	var i = 1

	for {
		var hash = md5.Sum([]byte(key + strconv.Itoa(i)))
		var hex = hex.EncodeToString(hash[:])
		if hex[:6] == "000000" {
			fmt.Print(i)
			break
		}
		i++
	}
}
