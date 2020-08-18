package id

import (
	"github.com/wenj91/util/num"
	"github.com/wenj91/util/str"
)

func GetClubID() string {
	i :=  num.RandomFromTo(10000000, 99999999)
	return str.String(i)
}