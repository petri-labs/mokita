package stableswap

import (
	"math/rand"
	"testing"

	"github.com/petri-labs/mokita/mokimath"
)

func BenchmarkCFMM(b *testing.B) {
	// Uses solveCfmm
	for i := 0; i < b.N; i++ {
		runCalcCFMM(solveCfmm)
	}
}

func BenchmarkBinarySearchMultiAsset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runCalcMultiAsset(solveCFMMBinarySearchMulti)
	}
}

func runCalcCFMM(solve func(mokimath.BigDec, mokimath.BigDec, []mokimath.BigDec, mokimath.BigDec) mokimath.BigDec) {
	xReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	yIn := mokimath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, []mokimath.BigDec{}, yIn)
}

func runCalcTwoAsset(solve func(mokimath.BigDec, mokimath.BigDec, mokimath.BigDec) mokimath.BigDec) {
	xReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	yIn := mokimath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, yIn)
}

func runCalcMultiAsset(solve func(mokimath.BigDec, mokimath.BigDec, mokimath.BigDec, mokimath.BigDec) mokimath.BigDec) {
	xReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	yReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	mReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	nReserve := mokimath.NewBigDec(rand.Int63n(100000) + 50000)
	w := mReserve.Mul(mReserve).Add(nReserve.Mul(nReserve))
	yIn := mokimath.NewBigDec(rand.Int63n(100000))
	solve(xReserve, yReserve, w, yIn)
}
