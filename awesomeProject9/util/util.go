package util

import (
	"math/rand"
	"time"
)

func RandomStiring(n int) string{
	var letters=[]byte("asdadadfehhlkjllAEWENCIIPDNBC")
	result:=make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:=range result{
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string((result))
}