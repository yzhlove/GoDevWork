package dh

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

var (
	randEngine  = rand.New(rand.NewSource(time.Now().UnixNano()))
	DH1BASE     = big.NewInt(3)
	DH1PRIME, _ = big.NewInt(0).SetString("0x7FFFFFc3", 0)
	MAXINT64    = big.NewInt(math.MaxInt64)
)

func DHExchange() (*big.Int, *big.Int) {
	SECRET := big.NewInt(0).Rand(randEngine, MAXINT64)
	MODPOWER := big.NewInt(0).Exp(DH1BASE, SECRET, DH1PRIME)
	return SECRET, MODPOWER
}

func DHKey(SECRET, MODPOWER *big.Int) *big.Int {
	KEY := big.NewInt(0).Exp(MODPOWER, SECRET, DH1PRIME)
	return KEY
}
