package mokimath_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/petri-labs/mokita/mokimath"
)

type decimalTestSuite struct {
	suite.Suite
}

var (
	zeroAdditiveErrTolerance = mokimath.ErrTolerance{
		AdditiveTolerance: sdk.ZeroDec(),
	}
)

func TestDecimalTestSuite(t *testing.T) {
	suite.Run(t, new(decimalTestSuite))
}

// assertMutResult given expected value after applying a math operation, a start value,
// mutative and non mutative results with start values, asserts that mutation are only applied
// to the mutative versions. Also, asserts that both results match the expected value.
func (s *decimalTestSuite) assertMutResult(expectedResult, startValue, mutativeResult, nonMutativeResult, mutativeStartValue, nonMutativeStartValue mokimath.BigDec) {
	// assert both results are as expected.
	s.Require().Equal(expectedResult, mutativeResult)
	s.Require().Equal(expectedResult, nonMutativeResult)

	// assert that mutative method mutated the receiver
	s.Require().Equal(mutativeStartValue, expectedResult)
	// assert that non-mutative method did not mutate the receiver
	s.Require().Equal(nonMutativeStartValue, startValue)
}

func (s *decimalTestSuite) TestAddMut() {
	toAdd := mokimath.MustNewDecFromStr("10")
	tests := map[string]struct {
		startValue        mokimath.BigDec
		expectedMutResult mokimath.BigDec
	}{
		"0":  {mokimath.NewBigDec(0), mokimath.NewBigDec(10)},
		"1":  {mokimath.NewBigDec(1), mokimath.NewBigDec(11)},
		"10": {mokimath.NewBigDec(10), mokimath.NewBigDec(20)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.AddMut(toAdd)
			resultNonMut := startNonMut.Add(toAdd)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestQuoMut() {
	quoBy := mokimath.MustNewDecFromStr("2")
	tests := map[string]struct {
		startValue        mokimath.BigDec
		expectedMutResult mokimath.BigDec
	}{
		"0":  {mokimath.NewBigDec(0), mokimath.NewBigDec(0)},
		"1":  {mokimath.NewBigDec(1), mokimath.MustNewDecFromStr("0.5")},
		"10": {mokimath.NewBigDec(10), mokimath.NewBigDec(5)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.QuoMut(quoBy)
			resultNonMut := startNonMut.Quo(quoBy)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}
func TestDecApproxEq(t *testing.T) {
	// d1 = 0.55, d2 = 0.6, tol = 0.1
	d1 := mokimath.NewDecWithPrec(55, 2)
	d2 := mokimath.NewDecWithPrec(6, 1)
	tol := mokimath.NewDecWithPrec(1, 1)

	require.True(mokimath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.55, d2 = 0.6, tol = 1E-5
	d1 = mokimath.NewDecWithPrec(55, 2)
	d2 = mokimath.NewDecWithPrec(6, 1)
	tol = mokimath.NewDecWithPrec(1, 5)

	require.False(mokimath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.6, d2 = 0.61, tol = 0.01
	d1 = mokimath.NewDecWithPrec(6, 1)
	d2 = mokimath.NewDecWithPrec(61, 2)
	tol = mokimath.NewDecWithPrec(1, 2)

	require.True(mokimath.DecApproxEq(t, d1, d2, tol))
}

// create a decimal from a decimal string (ex. "1234.5678")
func (s *decimalTestSuite) MustNewDecFromStr(str string) (d mokimath.BigDec) {
	d, err := mokimath.NewDecFromStr(str)
	s.Require().NoError(err)

	return d
}

func (s *decimalTestSuite) TestNewDecFromStr() {
	largeBigInt, success := new(big.Int).SetString("3144605511029693144278234343371835", 10)
	s.Require().True(success)

	tests := []struct {
		decimalStr string
		expErr     bool
		exp        mokimath.BigDec
	}{
		{"", true, mokimath.BigDec{}},
		{"0.-75", true, mokimath.BigDec{}},
		{"0", false, mokimath.NewBigDec(0)},
		{"1", false, mokimath.NewBigDec(1)},
		{"1.1", false, mokimath.NewDecWithPrec(11, 1)},
		{"0.75", false, mokimath.NewDecWithPrec(75, 2)},
		{"0.8", false, mokimath.NewDecWithPrec(8, 1)},
		{"0.11111", false, mokimath.NewDecWithPrec(11111, 5)},
		{"314460551102969.31442782343433718353144278234343371835", true, mokimath.NewBigDec(3141203149163817869)},
		{
			"314460551102969314427823434337.18357180924882313501835718092488231350",
			true, mokimath.NewDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{
			"314460551102969314427823434337.1835",
			false, mokimath.NewDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{".", true, mokimath.BigDec{}},
		{".0", true, mokimath.NewBigDec(0)},
		{"1.", true, mokimath.NewBigDec(1)},
		{"foobar", true, mokimath.BigDec{}},
		{"0.foobar", true, mokimath.BigDec{}},
		{"0.foobar.", true, mokimath.BigDec{}},
		{"179769313486231590772930519078902473361797697894230657273430081157732675805500963132708477322407536021120113879871393357658789768814416622492847430639474124377767893424865485276302219601246094119453082952085005768838150682342462881473913110540827237163350510684586298239947245938479716304835356329624224137216", true, mokimath.BigDec{}},
	}

	for tcIndex, tc := range tests {
		res, err := mokimath.NewDecFromStr(tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			s.Require().True(res.Equal(tc.exp), "equality was incorrect, res %v, exp %v, tc %v", res, tc.exp, tcIndex)
		}

		// negative tc
		res, err = mokimath.NewDecFromStr("-" + tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			exp := tc.exp.Mul(mokimath.NewBigDec(-1))
			s.Require().True(res.Equal(exp), "equality was incorrect, res %v, exp %v, tc %v", res, exp, tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestDecString() {
	tests := []struct {
		d    mokimath.BigDec
		want string
	}{
		{mokimath.NewBigDec(0), "0.000000000000000000000000000000000000"},
		{mokimath.NewBigDec(1), "1.000000000000000000000000000000000000"},
		{mokimath.NewBigDec(10), "10.000000000000000000000000000000000000"},
		{mokimath.NewBigDec(12340), "12340.000000000000000000000000000000000000"},
		{mokimath.NewDecWithPrec(12340, 4), "1.234000000000000000000000000000000000"},
		{mokimath.NewDecWithPrec(12340, 5), "0.123400000000000000000000000000000000"},
		{mokimath.NewDecWithPrec(12340, 8), "0.000123400000000000000000000000000000"},
		{mokimath.NewDecWithPrec(1009009009009009009, 17), "10.090090090090090090000000000000000000"},
		{mokimath.MustNewDecFromStr("10.090090090090090090090090090090090090"), "10.090090090090090090090090090090090090"},
	}
	for tcIndex, tc := range tests {
		s.Require().Equal(tc.want, tc.d.String(), "bad String(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecFloat64() {
	tests := []struct {
		d    mokimath.BigDec
		want float64
	}{
		{mokimath.NewBigDec(0), 0.000000000000000000},
		{mokimath.NewBigDec(1), 1.000000000000000000},
		{mokimath.NewBigDec(10), 10.000000000000000000},
		{mokimath.NewBigDec(12340), 12340.000000000000000000},
		{mokimath.NewDecWithPrec(12340, 4), 1.234000000000000000},
		{mokimath.NewDecWithPrec(12340, 5), 0.123400000000000000},
		{mokimath.NewDecWithPrec(12340, 8), 0.000123400000000000},
		{mokimath.NewDecWithPrec(1009009009009009009, 17), 10.090090090090090090},
	}
	for tcIndex, tc := range tests {
		value, err := tc.d.Float64()
		s.Require().Nil(err, "error getting Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, value, "bad Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, tc.d.MustFloat64(), "bad MustFloat64(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestSdkDec() {
	tests := []struct {
		d        mokimath.BigDec
		want     sdk.Dec
		expPanic bool
	}{
		{mokimath.NewBigDec(0), sdk.MustNewDecFromStr("0.000000000000000000"), false},
		{mokimath.NewBigDec(1), sdk.MustNewDecFromStr("1.000000000000000000"), false},
		{mokimath.NewBigDec(10), sdk.MustNewDecFromStr("10.000000000000000000"), false},
		{mokimath.NewBigDec(12340), sdk.MustNewDecFromStr("12340.000000000000000000"), false},
		{mokimath.NewDecWithPrec(12340, 4), sdk.MustNewDecFromStr("1.234000000000000000"), false},
		{mokimath.NewDecWithPrec(12340, 5), sdk.MustNewDecFromStr("0.123400000000000000"), false},
		{mokimath.NewDecWithPrec(12340, 8), sdk.MustNewDecFromStr("0.000123400000000000"), false},
		{mokimath.NewDecWithPrec(1009009009009009009, 17), sdk.MustNewDecFromStr("10.090090090090090090"), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { tc.d.SDKDec() })
		} else {
			value := tc.d.SDKDec()
			s.Require().Equal(tc.want, value, "bad SdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDec() {
	tests := []struct {
		d        sdk.Dec
		want     mokimath.BigDec
		expPanic bool
	}{
		{sdk.MustNewDecFromStr("0.000000000000000000"), mokimath.NewBigDec(0), false},
		{sdk.MustNewDecFromStr("1.000000000000000000"), mokimath.NewBigDec(1), false},
		{sdk.MustNewDecFromStr("10.000000000000000000"), mokimath.NewBigDec(10), false},
		{sdk.MustNewDecFromStr("12340.000000000000000000"), mokimath.NewBigDec(12340), false},
		{sdk.MustNewDecFromStr("1.234000000000000000"), mokimath.NewDecWithPrec(12340, 4), false},
		{sdk.MustNewDecFromStr("0.123400000000000000"), mokimath.NewDecWithPrec(12340, 5), false},
		{sdk.MustNewDecFromStr("0.000123400000000000"), mokimath.NewDecWithPrec(12340, 8), false},
		{sdk.MustNewDecFromStr("10.090090090090090090"), mokimath.NewDecWithPrec(1009009009009009009, 17), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { mokimath.BigDecFromSDKDec(tc.d) })
		} else {
			value := mokimath.BigDecFromSDKDec(tc.d)
			s.Require().Equal(tc.want, value, "bad mokimath.BigDecFromSdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDecSlice() {
	tests := []struct {
		d        []sdk.Dec
		want     []mokimath.BigDec
		expPanic bool
	}{
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000000000000000000")}, []mokimath.BigDec{mokimath.NewBigDec(0)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000000000000000000"), sdk.MustNewDecFromStr("1.000000000000000000")}, []mokimath.BigDec{mokimath.NewBigDec(0), mokimath.NewBigDec(1)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("1.000000000000000000"), sdk.MustNewDecFromStr("0.000000000000000000"), sdk.MustNewDecFromStr("0.000123400000000000")}, []mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(0), mokimath.NewDecWithPrec(12340, 8)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("10.000000000000000000")}, []mokimath.BigDec{mokimath.NewBigDec(10)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("12340.000000000000000000")}, []mokimath.BigDec{mokimath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("1.234000000000000000"), sdk.MustNewDecFromStr("12340.000000000000000000")}, []mokimath.BigDec{mokimath.NewDecWithPrec(12340, 4), mokimath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.123400000000000000"), sdk.MustNewDecFromStr("12340.000000000000000000")}, []mokimath.BigDec{mokimath.NewDecWithPrec(12340, 5), mokimath.NewBigDec(12340)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("0.000123400000000000"), sdk.MustNewDecFromStr("10.090090090090090090")}, []mokimath.BigDec{mokimath.NewDecWithPrec(12340, 8), mokimath.NewDecWithPrec(1009009009009009009, 17)}, false},
		{[]sdk.Dec{sdk.MustNewDecFromStr("10.090090090090090090"), sdk.MustNewDecFromStr("10.090090090090090090")}, []mokimath.BigDec{mokimath.NewDecWithPrec(1009009009009009009, 17), mokimath.NewDecWithPrec(1009009009009009009, 17)}, false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { mokimath.BigDecFromSDKDecSlice(tc.d) })
		} else {
			value := mokimath.BigDecFromSDKDecSlice(tc.d)
			s.Require().Equal(tc.want, value, "bad mokimath.BigDecFromSdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestEqualities() {
	tests := []struct {
		d1, d2     mokimath.BigDec
		gt, lt, eq bool
	}{
		{mokimath.NewBigDec(0), mokimath.NewBigDec(0), false, false, true},
		{mokimath.NewDecWithPrec(0, 2), mokimath.NewDecWithPrec(0, 4), false, false, true},
		{mokimath.NewDecWithPrec(100, 0), mokimath.NewDecWithPrec(100, 0), false, false, true},
		{mokimath.NewDecWithPrec(-100, 0), mokimath.NewDecWithPrec(-100, 0), false, false, true},
		{mokimath.NewDecWithPrec(-1, 1), mokimath.NewDecWithPrec(-1, 1), false, false, true},
		{mokimath.NewDecWithPrec(3333, 3), mokimath.NewDecWithPrec(3333, 3), false, false, true},

		{mokimath.NewDecWithPrec(0, 0), mokimath.NewDecWithPrec(3333, 3), false, true, false},
		{mokimath.NewDecWithPrec(0, 0), mokimath.NewDecWithPrec(100, 0), false, true, false},
		{mokimath.NewDecWithPrec(-1, 0), mokimath.NewDecWithPrec(3333, 3), false, true, false},
		{mokimath.NewDecWithPrec(-1, 0), mokimath.NewDecWithPrec(100, 0), false, true, false},
		{mokimath.NewDecWithPrec(1111, 3), mokimath.NewDecWithPrec(100, 0), false, true, false},
		{mokimath.NewDecWithPrec(1111, 3), mokimath.NewDecWithPrec(3333, 3), false, true, false},
		{mokimath.NewDecWithPrec(-3333, 3), mokimath.NewDecWithPrec(-1111, 3), false, true, false},

		{mokimath.NewDecWithPrec(3333, 3), mokimath.NewDecWithPrec(0, 0), true, false, false},
		{mokimath.NewDecWithPrec(100, 0), mokimath.NewDecWithPrec(0, 0), true, false, false},
		{mokimath.NewDecWithPrec(3333, 3), mokimath.NewDecWithPrec(-1, 0), true, false, false},
		{mokimath.NewDecWithPrec(100, 0), mokimath.NewDecWithPrec(-1, 0), true, false, false},
		{mokimath.NewDecWithPrec(100, 0), mokimath.NewDecWithPrec(1111, 3), true, false, false},
		{mokimath.NewDecWithPrec(3333, 3), mokimath.NewDecWithPrec(1111, 3), true, false, false},
		{mokimath.NewDecWithPrec(-1111, 3), mokimath.NewDecWithPrec(-3333, 3), true, false, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.gt, tc.d1.GT(tc.d2), "GT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.lt, tc.d1.LT(tc.d2), "LT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, tc.d1.Equal(tc.d2), "equality result is incorrect, tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecsEqual() {
	tests := []struct {
		d1s, d2s []mokimath.BigDec
		eq       bool
	}{
		{[]mokimath.BigDec{mokimath.NewBigDec(0)}, []mokimath.BigDec{mokimath.NewBigDec(0)}, true},
		{[]mokimath.BigDec{mokimath.NewBigDec(0)}, []mokimath.BigDec{mokimath.NewBigDec(1)}, false},
		{[]mokimath.BigDec{mokimath.NewBigDec(0)}, []mokimath.BigDec{}, false},
		{[]mokimath.BigDec{mokimath.NewBigDec(0), mokimath.NewBigDec(1)}, []mokimath.BigDec{mokimath.NewBigDec(0), mokimath.NewBigDec(1)}, true},
		{[]mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(0)}, []mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(0)}, true},
		{[]mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(0)}, []mokimath.BigDec{mokimath.NewBigDec(0), mokimath.NewBigDec(1)}, false},
		{[]mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(0)}, []mokimath.BigDec{mokimath.NewBigDec(1)}, false},
		{[]mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(2)}, []mokimath.BigDec{mokimath.NewBigDec(2), mokimath.NewBigDec(4)}, false},
		{[]mokimath.BigDec{mokimath.NewBigDec(3), mokimath.NewBigDec(18)}, []mokimath.BigDec{mokimath.NewBigDec(1), mokimath.NewBigDec(6)}, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.eq, mokimath.DecsEqual(tc.d1s, tc.d2s), "equality of decional arrays is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, mokimath.DecsEqual(tc.d2s, tc.d1s), "equality of decional arrays is incorrect (converse), tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestArithmetic() {
	tests := []struct {
		d1, d2                                mokimath.BigDec
		expMul, expMulTruncate                mokimath.BigDec
		expQuo, expQuoRoundUp, expQuoTruncate mokimath.BigDec
		expAdd, expSub                        mokimath.BigDec
	}{
		//  d1         d2         MUL    MulTruncate    QUO    QUORoundUp QUOTrunctate  ADD         SUB
		{mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0)},
		{mokimath.NewBigDec(1), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(1), mokimath.NewBigDec(1)},
		{mokimath.NewBigDec(0), mokimath.NewBigDec(1), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(1), mokimath.NewBigDec(-1)},
		{mokimath.NewBigDec(0), mokimath.NewBigDec(-1), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(-1), mokimath.NewBigDec(1)},
		{mokimath.NewBigDec(-1), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(0), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1)},

		{mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(2), mokimath.NewBigDec(0)},
		{mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(-2), mokimath.NewBigDec(0)},
		{mokimath.NewBigDec(1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(0), mokimath.NewBigDec(2)},
		{mokimath.NewBigDec(-1), mokimath.NewBigDec(1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(-1), mokimath.NewBigDec(0), mokimath.NewBigDec(-2)},

		{
			mokimath.NewBigDec(3), mokimath.NewBigDec(7), mokimath.NewBigDec(21), mokimath.NewBigDec(21),
			mokimath.MustNewDecFromStr("0.428571428571428571428571428571428571"), mokimath.MustNewDecFromStr("0.428571428571428571428571428571428572"), mokimath.MustNewDecFromStr("0.428571428571428571428571428571428571"),
			mokimath.NewBigDec(10), mokimath.NewBigDec(-4),
		},
		{
			mokimath.NewBigDec(2), mokimath.NewBigDec(4), mokimath.NewBigDec(8), mokimath.NewBigDec(8), mokimath.NewDecWithPrec(5, 1), mokimath.NewDecWithPrec(5, 1), mokimath.NewDecWithPrec(5, 1),
			mokimath.NewBigDec(6), mokimath.NewBigDec(-2),
		},

		{mokimath.NewBigDec(100), mokimath.NewBigDec(100), mokimath.NewBigDec(10000), mokimath.NewBigDec(10000), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(200), mokimath.NewBigDec(0)},

		{
			mokimath.NewDecWithPrec(15, 1), mokimath.NewDecWithPrec(15, 1), mokimath.NewDecWithPrec(225, 2), mokimath.NewDecWithPrec(225, 2),
			mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(1), mokimath.NewBigDec(3), mokimath.NewBigDec(0),
		},
		{
			mokimath.NewDecWithPrec(3333, 4), mokimath.NewDecWithPrec(333, 4), mokimath.NewDecWithPrec(1109889, 8), mokimath.NewDecWithPrec(1109889, 8),
			mokimath.MustNewDecFromStr("10.009009009009009009009009009009009009"), mokimath.MustNewDecFromStr("10.009009009009009009009009009009009010"), mokimath.MustNewDecFromStr("10.009009009009009009009009009009009009"),
			mokimath.NewDecWithPrec(3666, 4), mokimath.NewDecWithPrec(3, 1),
		},
	}

	for tcIndex, tc := range tests {
		tc := tc
		resAdd := tc.d1.Add(tc.d2)
		resSub := tc.d1.Sub(tc.d2)
		resMul := tc.d1.Mul(tc.d2)
		resMulTruncate := tc.d1.MulTruncate(tc.d2)
		s.Require().True(tc.expAdd.Equal(resAdd), "exp %v, res %v, tc %d", tc.expAdd, resAdd, tcIndex)
		s.Require().True(tc.expSub.Equal(resSub), "exp %v, res %v, tc %d", tc.expSub, resSub, tcIndex)
		s.Require().True(tc.expMul.Equal(resMul), "exp %v, res %v, tc %d", tc.expMul, resMul, tcIndex)
		s.Require().True(tc.expMulTruncate.Equal(resMulTruncate), "exp %v, res %v, tc %d", tc.expMulTruncate, resMulTruncate, tcIndex)

		if tc.d2.IsZero() { // panic for divide by zero
			s.Require().Panics(func() { tc.d1.Quo(tc.d2) })
		} else {
			resQuo := tc.d1.Quo(tc.d2)
			s.Require().True(tc.expQuo.Equal(resQuo), "exp %v, res %v, tc %d", tc.expQuo.String(), resQuo.String(), tcIndex)

			resQuoRoundUp := tc.d1.QuoRoundUp(tc.d2)
			s.Require().True(tc.expQuoRoundUp.Equal(resQuoRoundUp), "exp %v, res %v, tc %d",
				tc.expQuoRoundUp.String(), resQuoRoundUp.String(), tcIndex)

			resQuoTruncate := tc.d1.QuoTruncate(tc.d2)
			s.Require().True(tc.expQuoTruncate.Equal(resQuoTruncate), "exp %v, res %v, tc %d",
				tc.expQuoTruncate.String(), resQuoTruncate.String(), tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBankerRoundChop() {
	tests := []struct {
		d1  mokimath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("0.75"), 1},
		{s.MustNewDecFromStr("0.5"), 0},
		{s.MustNewDecFromStr("7.5"), 8},
		{s.MustNewDecFromStr("1.5"), 2},
		{s.MustNewDecFromStr("2.5"), 2},
		{s.MustNewDecFromStr("0.545"), 1}, // 0.545-> 1 even though 5 is first decimal and 1 not even
		{s.MustNewDecFromStr("1.545"), 2},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().RoundInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.RoundInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestTruncate() {
	tests := []struct {
		d1  mokimath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0.75"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("1.5"), 1},
		{s.MustNewDecFromStr("7.5"), 7},
		{s.MustNewDecFromStr("7.6"), 7},
		{s.MustNewDecFromStr("7.4"), 7},
		{s.MustNewDecFromStr("100.1"), 100},
		{s.MustNewDecFromStr("1000.1"), 1000},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().TruncateInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.TruncateInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestStringOverflow() {
	// two random 64 bit primes
	dec1, err := mokimath.NewDecFromStr("51643150036226787134389711697696177267")
	s.Require().NoError(err)
	dec2, err := mokimath.NewDecFromStr("-31798496660535729618459429845579852627")
	s.Require().NoError(err)
	dec3 := dec1.Add(dec2)
	s.Require().Equal(
		"19844653375691057515930281852116324640.000000000000000000000000000000000000",
		dec3.String(),
	)
}

func (s *decimalTestSuite) TestDecMulInt() {
	tests := []struct {
		sdkDec mokimath.BigDec
		sdkInt mokimath.BigInt
		want   mokimath.BigDec
	}{
		{mokimath.NewBigDec(10), mokimath.NewInt(2), mokimath.NewBigDec(20)},
		{mokimath.NewBigDec(1000000), mokimath.NewInt(100), mokimath.NewBigDec(100000000)},
		{mokimath.NewDecWithPrec(1, 1), mokimath.NewInt(10), mokimath.NewBigDec(1)},
		{mokimath.NewDecWithPrec(1, 5), mokimath.NewInt(20), mokimath.NewDecWithPrec(2, 4)},
	}
	for i, tc := range tests {
		got := tc.sdkDec.MulInt(tc.sdkInt)
		s.Require().Equal(tc.want, got, "Incorrect result on test case %d", i)
	}
}

func (s *decimalTestSuite) TestDecCeil() {
	testCases := []struct {
		input    mokimath.BigDec
		expected mokimath.BigDec
	}{
		{mokimath.MustNewDecFromStr("0.001"), mokimath.NewBigDec(1)},   // 0.001 => 1.0
		{mokimath.MustNewDecFromStr("-0.001"), mokimath.ZeroDec()},     // -0.001 => 0.0
		{mokimath.ZeroDec(), mokimath.ZeroDec()},                       // 0.0 => 0.0
		{mokimath.MustNewDecFromStr("0.9"), mokimath.NewBigDec(1)},     // 0.9 => 1.0
		{mokimath.MustNewDecFromStr("4.001"), mokimath.NewBigDec(5)},   // 4.001 => 5.0
		{mokimath.MustNewDecFromStr("-4.001"), mokimath.NewBigDec(-4)}, // -4.001 => -4.0
		{mokimath.MustNewDecFromStr("4.7"), mokimath.NewBigDec(5)},     // 4.7 => 5.0
		{mokimath.MustNewDecFromStr("-4.7"), mokimath.NewBigDec(-4)},   // -4.7 => -4.0
	}

	for i, tc := range testCases {
		res := tc.input.Ceil()
		s.Require().Equal(tc.expected, res, "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestApproxRoot() {
	testCases := []struct {
		input    mokimath.BigDec
		root     uint64
		expected mokimath.BigDec
	}{
		{mokimath.OneDec(), 10, mokimath.OneDec()},                                                                            // 1.0 ^ (0.1) => 1.0
		{mokimath.NewDecWithPrec(25, 2), 2, mokimath.NewDecWithPrec(5, 1)},                                                    // 0.25 ^ (0.5) => 0.5
		{mokimath.NewDecWithPrec(4, 2), 2, mokimath.NewDecWithPrec(2, 1)},                                                     // 0.04 ^ (0.5) => 0.2
		{mokimath.NewDecFromInt(mokimath.NewInt(27)), 3, mokimath.NewDecFromInt(mokimath.NewInt(3))},                          // 27 ^ (1/3) => 3
		{mokimath.NewDecFromInt(mokimath.NewInt(-81)), 4, mokimath.NewDecFromInt(mokimath.NewInt(-3))},                        // -81 ^ (0.25) => -3
		{mokimath.NewDecFromInt(mokimath.NewInt(2)), 2, mokimath.MustNewDecFromStr("1.414213562373095048801688724209698079")}, // 2 ^ (0.5) => 1.414213562373095048801688724209698079
		{mokimath.NewDecWithPrec(1005, 3), 31536000, mokimath.MustNewDecFromStr("1.000000000158153903837946258002096839")},    // 1.005 ^ (1/31536000) ≈ 1.000000000158153903837946258002096839
		{mokimath.SmallestDec(), 2, mokimath.NewDecWithPrec(1, 18)},                                                           // 1e-36 ^ (0.5) => 1e-18
		{mokimath.SmallestDec(), 3, mokimath.MustNewDecFromStr("0.000000000001000000000000000002431786")},                     // 1e-36 ^ (1/3) => 1e-12
		{mokimath.NewDecWithPrec(1, 8), 3, mokimath.MustNewDecFromStr("0.002154434690031883721759293566519280")},              // 1e-8 ^ (1/3) ≈ 0.002154434690031883721759293566519
	}

	// In the case of 1e-8 ^ (1/3), the result repeats every 5 iterations starting from iteration 24
	// (i.e. 24, 29, 34, ... give the same result) and never converges enough. The maximum number of
	// iterations (100) causes the result at iteration 100 to be returned, regardless of convergence.

	for i, tc := range testCases {
		res, err := tc.input.ApproxRoot(tc.root)
		s.Require().NoError(err)
		s.Require().True(tc.expected.Sub(res).Abs().LTE(mokimath.SmallestDec()), "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestApproxSqrt() {
	testCases := []struct {
		input    mokimath.BigDec
		expected mokimath.BigDec
	}{
		{mokimath.OneDec(), mokimath.OneDec()},                                                                             // 1.0 => 1.0
		{mokimath.NewDecWithPrec(25, 2), mokimath.NewDecWithPrec(5, 1)},                                                    // 0.25 => 0.5
		{mokimath.NewDecWithPrec(4, 2), mokimath.NewDecWithPrec(2, 1)},                                                     // 0.09 => 0.3
		{mokimath.NewDecFromInt(mokimath.NewInt(9)), mokimath.NewDecFromInt(mokimath.NewInt(3))},                           // 9 => 3
		{mokimath.NewDecFromInt(mokimath.NewInt(-9)), mokimath.NewDecFromInt(mokimath.NewInt(-3))},                         // -9 => -3
		{mokimath.NewDecFromInt(mokimath.NewInt(2)), mokimath.MustNewDecFromStr("1.414213562373095048801688724209698079")}, // 2 => 1.414213562373095048801688724209698079
	}

	for i, tc := range testCases {
		res, err := tc.input.ApproxSqrt()
		s.Require().NoError(err)
		s.Require().Equal(tc.expected, res, "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestDecSortableBytes() {
	tests := []struct {
		d    mokimath.BigDec
		want []byte
	}{
		{mokimath.NewBigDec(0), []byte("000000000000000000000000000000000000.000000000000000000000000000000000000")},
		{mokimath.NewBigDec(1), []byte("000000000000000000000000000000000001.000000000000000000000000000000000000")},
		{mokimath.NewBigDec(10), []byte("000000000000000000000000000000000010.000000000000000000000000000000000000")},
		{mokimath.NewBigDec(12340), []byte("000000000000000000000000000000012340.000000000000000000000000000000000000")},
		{mokimath.NewDecWithPrec(12340, 4), []byte("000000000000000000000000000000000001.234000000000000000000000000000000000")},
		{mokimath.NewDecWithPrec(12340, 5), []byte("000000000000000000000000000000000000.123400000000000000000000000000000000")},
		{mokimath.NewDecWithPrec(12340, 8), []byte("000000000000000000000000000000000000.000123400000000000000000000000000000")},
		{mokimath.NewDecWithPrec(1009009009009009009, 17), []byte("000000000000000000000000000000000010.090090090090090090000000000000000000")},
		{mokimath.NewDecWithPrec(-1009009009009009009, 17), []byte("-000000000000000000000000000000000010.090090090090090090000000000000000000")},
		{mokimath.MustNewDecFromStr("1000000000000000000000000000000000000"), []byte("max")},
		{mokimath.MustNewDecFromStr("-1000000000000000000000000000000000000"), []byte("--")},
	}
	for tcIndex, tc := range tests {
		s.Require().Equal(tc.want, mokimath.SortableDecBytes(tc.d), "bad String(), index: %v", tcIndex)
	}

	s.Require().Panics(func() { mokimath.SortableDecBytes(mokimath.MustNewDecFromStr("1000000000000000000000000000000000001")) })
	s.Require().Panics(func() {
		mokimath.SortableDecBytes(mokimath.MustNewDecFromStr("-1000000000000000000000000000000000001"))
	})
}

func (s *decimalTestSuite) TestDecEncoding() {
	testCases := []struct {
		input   mokimath.BigDec
		rawBz   string
		jsonStr string
		yamlStr string
	}{
		{
			mokimath.NewBigDec(0), "30",
			"\"0.000000000000000000000000000000000000\"",
			"\"0.000000000000000000000000000000000000\"\n",
		},
		{
			mokimath.NewDecWithPrec(4, 2),
			"3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"0.040000000000000000000000000000000000\"",
			"\"0.040000000000000000000000000000000000\"\n",
		},
		{
			mokimath.NewDecWithPrec(-4, 2),
			"2D3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"-0.040000000000000000000000000000000000\"",
			"\"-0.040000000000000000000000000000000000\"\n",
		},
		{
			mokimath.MustNewDecFromStr("1.414213562373095048801688724209698079"),
			"31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"1.414213562373095048801688724209698079\"",
			"\"1.414213562373095048801688724209698079\"\n",
		},
		{
			mokimath.MustNewDecFromStr("-1.414213562373095048801688724209698079"),
			"2D31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"-1.414213562373095048801688724209698079\"",
			"\"-1.414213562373095048801688724209698079\"\n",
		},
	}

	for _, tc := range testCases {
		bz, err := tc.input.Marshal()
		s.Require().NoError(err)
		s.Require().Equal(tc.rawBz, fmt.Sprintf("%X", bz))

		var other mokimath.BigDec
		s.Require().NoError((&other).Unmarshal(bz))
		s.Require().True(tc.input.Equal(other))

		bz, err = json.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.jsonStr, string(bz))
		s.Require().NoError(json.Unmarshal(bz, &other))
		s.Require().True(tc.input.Equal(other))

		bz, err = yaml.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.yamlStr, string(bz))
	}
}

// Showcase that different orders of operations causes different results.
func (s *decimalTestSuite) TestOperationOrders() {
	n1 := mokimath.NewBigDec(10)
	n2 := mokimath.NewBigDec(1000000010)
	s.Require().Equal(n1.Mul(n2).Quo(n2), mokimath.NewBigDec(10))
	s.Require().NotEqual(n1.Mul(n2).Quo(n2), n1.Quo(n2).Mul(n2))
}

func BenchmarkMarshalTo(b *testing.B) {
	b.ReportAllocs()
	bis := []struct {
		in   mokimath.BigDec
		want []byte
	}{
		{
			mokimath.NewBigDec(1e8), []byte{
				0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
			},
		},
		{mokimath.NewBigDec(0), []byte{0x30}},
	}
	data := make([]byte, 100)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, bi := range bis {
			if n, err := bi.in.MarshalTo(data); err != nil {
				b.Fatal(err)
			} else {
				if !bytes.Equal(data[:n], bi.want) {
					b.Fatalf("Mismatch\nGot:  % x\nWant: % x\n", data[:n], bi.want)
				}
			}
		}
	}
}

func (s *decimalTestSuite) TestLog2() {
	var expectedErrTolerance = mokimath.MustNewDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue mokimath.BigDec
		expected     mokimath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}; invalid; panic": {
			initialValue:  mokimath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_2{0}; invalid; panic": {
			initialValue:  mokimath.ZeroDec(),
			expectedPanic: true,
		},
		"log_2{0.001} = -9.965784284662087043610958288468170528": {
			initialValue: mokimath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+33+digits
			expected: mokimath.MustNewDecFromStr("-9.965784284662087043610958288468170528"),
		},
		"log_2{0.56171821941421412902170941} = -0.832081497183140708984033250637831402": {
			initialValue: mokimath.MustNewDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.56171821941421412902170941+with+36+digits
			expected: mokimath.MustNewDecFromStr("-0.832081497183140708984033250637831402"),
		},
		"log_2{0.999912345} = -0.000126464976533858080645902722235833": {
			initialValue: mokimath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+37+digits
			expected: mokimath.MustNewDecFromStr("-0.000126464976533858080645902722235833"),
		},
		"log_2{1} = 0": {
			initialValue: mokimath.NewBigDec(1),
			expected:     mokimath.NewBigDec(0),
		},
		"log_2{2} = 1": {
			initialValue: mokimath.NewBigDec(2),
			expected:     mokimath.NewBigDec(1),
		},
		"log_2{7} = 2.807354922057604107441969317231830809": {
			initialValue: mokimath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+7+37+digits
			expected: mokimath.MustNewDecFromStr("2.807354922057604107441969317231830809"),
		},
		"log_2{512} = 9": {
			initialValue: mokimath.NewBigDec(512),
			expected:     mokimath.NewBigDec(9),
		},
		"log_2{580} = 9.179909090014934468590092754117374938": {
			initialValue: mokimath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+600+37+digits
			expected: mokimath.MustNewDecFromStr("9.179909090014934468590092754117374938"),
		},
		"log_2{1024} = 10": {
			initialValue: mokimath.NewBigDec(1024),
			expected:     mokimath.NewBigDec(10),
		},
		"log_2{1024.987654321} = 10.001390817654141324352719749259888355": {
			initialValue: mokimath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+1024.987654321+38+digits
			expected: mokimath.MustNewDecFromStr("10.001390817654141324352719749259888355"),
		},
		"log_2{912648174127941279170121098210.92821920190204131121} = 99.525973560175362367047484597337715868": {
			initialValue: mokimath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+912648174127941279170121098210.92821920190204131121+38+digits
			expected: mokimath.MustNewDecFromStr("99.525973560175362367047484597337715868"),
		},
		"log_2{Max Spot Price} = 128": {
			initialValue: mokimath.BigDecFromSDKDec(mokimath.MaxSpotPrice), // 2^128 - 1
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%282%5E128%29+-+1%29+38+digits
			expected: mokimath.MustNewDecFromStr("128"),
		},
		// The value tested below is: gammtypes.MaxSpotPrice * 0.99 = (2^128 - 1) * 0.99
		"log_2{336879543251729078828740861357450529340.45} = 127.98550043030488492336620207564264562": {
			initialValue: mokimath.MustNewDecFromStr("336879543251729078828740861357450529340.45"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%28%282%5E128%29+-+1%29*0.99%29++38+digits
			expected: mokimath.MustNewDecFromStr("127.98550043030488492336620207564264562"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			mokimath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that LogbBase2() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.LogBase2()
				require.True(mokimath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestLn() {
	var expectedErrTolerance = mokimath.MustNewDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue mokimath.BigDec
		expected     mokimath.BigDec

		expectedPanic bool
	}{
		"log_e{-1}; invalid; panic": {
			initialValue:  mokimath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_e{0}; invalid; panic": {
			initialValue:  mokimath.ZeroDec(),
			expectedPanic: true,
		},
		"log_e{0.001} = -6.90775527898213705205397436405309262": {
			initialValue: mokimath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log0.001+to+36+digits+with+36+decimals
			expected: mokimath.MustNewDecFromStr("-6.90775527898213705205397436405309262"),
		},
		"log_e{0.56171821941421412902170941} = -0.576754943768592057376050794884207180": {
			initialValue: mokimath.MustNewDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log0.56171821941421412902170941+to+36+digits
			expected: mokimath.MustNewDecFromStr("-0.576754943768592057376050794884207180"),
		},
		"log_e{0.999912345} = -0.000087658841924023373535614212850888": {
			initialValue: mokimath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log0.999912345+to+32+digits
			expected: mokimath.MustNewDecFromStr("-0.000087658841924023373535614212850888"),
		},
		"log_e{1} = 0": {
			initialValue: mokimath.NewBigDec(1),
			expected:     mokimath.NewBigDec(0),
		},
		"log_e{e} = 1": {
			initialValue: mokimath.MustNewDecFromStr("2.718281828459045235360287471352662498"),
			// From: https://www.wolframalpha.com/input?i=e+with+36+decimals
			expected: mokimath.NewBigDec(1),
		},
		"log_e{7} = 1.945910149055313305105352743443179730": {
			initialValue: mokimath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log7+up+to+36+decimals
			expected: mokimath.MustNewDecFromStr("1.945910149055313305105352743443179730"),
		},
		"log_e{512} = 6.238324625039507784755089093123589113": {
			initialValue: mokimath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log512+up+to+36+decimals
			expected: mokimath.MustNewDecFromStr("6.238324625039507784755089093123589113"),
		},
		"log_e{580} = 6.36302810354046502061849560850445238": {
			initialValue: mokimath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log580+up+to+36+decimals
			expected: mokimath.MustNewDecFromStr("6.36302810354046502061849560850445238"),
		},
		"log_e{1024.987654321} = 6.93243584693509415029056534690631614": {
			initialValue: mokimath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log1024.987654321+to+36+digits
			expected: mokimath.MustNewDecFromStr("6.93243584693509415029056534690631614"),
		},
		"log_e{912648174127941279170121098210.92821920190204131121} = 68.986147965719214790400745338243805015": {
			initialValue: mokimath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log912648174127941279170121098210.92821920190204131121+to+38+digits
			expected: mokimath.MustNewDecFromStr("68.986147965719214790400745338243805015"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			mokimath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.Ln()
				require.True(mokimath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestTickLog() {
	tests := map[string]struct {
		initialValue mokimath.BigDec
		expected     mokimath.BigDec

		expectedErrTolerance mokimath.BigDec
		expectedPanic        bool
	}{
		"log_1.0001{-1}; invalid; panic": {
			initialValue:  mokimath.OneDec().Neg(),
			expectedPanic: true,
		},
		"log_1.0001{0}; invalid; panic": {
			initialValue:  mokimath.ZeroDec(),
			expectedPanic: true,
		},
		"log_1.0001{0.001} = -69081.006609899112313305835611219486392199": {
			initialValue: mokimath.MustNewDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.001%29+to+41+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000143031879"),
			expected:             mokimath.MustNewDecFromStr("-69081.006609899112313305835611219486392199"),
		},
		"log_1.0001{0.999912345} = -0.876632247930741919880461740717176538": {
			initialValue: mokimath.MustNewDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.999912345%29+to+36+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000138702"),
			expected:             mokimath.MustNewDecFromStr("-0.876632247930741919880461740717176538"),
		},
		"log_1.0001{1} = 0": {
			initialValue: mokimath.NewBigDec(1),

			expectedErrTolerance: mokimath.ZeroDec(),
			expected:             mokimath.NewBigDec(0),
		},
		"log_1.0001{1.0001} = 1": {
			initialValue: mokimath.MustNewDecFromStr("1.0001"),

			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000152500"),
			expected:             mokimath.OneDec(),
		},
		"log_1.0001{512} = 62386.365360724158196763710649998441051753": {
			initialValue: mokimath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28512%29+to+41+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000129292137"),
			expected:             mokimath.MustNewDecFromStr("62386.365360724158196763710649998441051753"),
		},
		"log_1.0001{1024.987654321} = 69327.824629506998657531621822514042777198": {
			initialValue: mokimath.NewDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%281024.987654321%29+to+41+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000143836264"),
			expected:             mokimath.MustNewDecFromStr("69327.824629506998657531621822514042777198"),
		},
		"log_1.0001{912648174127941279170121098210.92821920190204131121} = 689895.972156319183538389792485913311778672": {
			initialValue: mokimath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28912648174127941279170121098210.92821920190204131121%29+to+42+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000001429936067"),
			expected:             mokimath.MustNewDecFromStr("689895.972156319183538389792485913311778672"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			mokimath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.TickLog()
				fmt.Println(name, res.Sub(tc.expected).Abs())
				require.True(mokimath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestCustomBaseLog() {
	tests := map[string]struct {
		initialValue mokimath.BigDec
		base         mokimath.BigDec

		expected             mokimath.BigDec
		expectedErrTolerance mokimath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}: normal base, invalid argument - panics": {
			initialValue:  mokimath.NewBigDec(-1),
			base:          mokimath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_2{0}: normal base, invalid argument - panics": {
			initialValue:  mokimath.NewBigDec(0),
			base:          mokimath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_(-1)(2): invalid base, normal argument - panics": {
			initialValue:  mokimath.NewBigDec(2),
			base:          mokimath.NewBigDec(-1),
			expectedPanic: true,
		},
		"log_1(2): base cannot equal to 1 - panics": {
			initialValue:  mokimath.NewBigDec(2),
			base:          mokimath.NewBigDec(1),
			expectedPanic: true,
		},
		"log_30(100) = 1.353984985057691049642502891262784015": {
			initialValue: mokimath.NewBigDec(100),
			base:         mokimath.NewBigDec(30),
			// From: https://www.wolframalpha.com/input?i=log_30%28100%29+to+37+digits
			expectedErrTolerance: mokimath.ZeroDec(),
			expected:             mokimath.MustNewDecFromStr("1.353984985057691049642502891262784015"),
		},
		"log_0.2(0.99) = 0.006244624769837438271878639001855450": {
			initialValue: mokimath.MustNewDecFromStr("0.99"),
			base:         mokimath.MustNewDecFromStr("0.2"),
			// From: https://www.wolframalpha.com/input?i=log_0.2%280.99%29+to+34+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000000013"),
			expected:             mokimath.MustNewDecFromStr("0.006244624769837438271878639001855450"),
		},

		"log_0.0001(500000) = -1.424742501084004701196565276318876743": {
			initialValue: mokimath.NewBigDec(500000),
			base:         mokimath.NewDecWithPrec(1, 4),
			// From: https://www.wolframalpha.com/input?i=log_0.0001%28500000%29+to+37+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000000003"),
			expected:             mokimath.MustNewDecFromStr("-1.424742501084004701196565276318876743"),
		},

		"log_500000(0.0001) = -0.701881216598197542030218906945601429": {
			initialValue: mokimath.NewDecWithPrec(1, 4),
			base:         mokimath.NewBigDec(500000),
			// From: https://www.wolframalpha.com/input?i=log_500000%280.0001%29+to+36+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000000001"),
			expected:             mokimath.MustNewDecFromStr("-0.701881216598197542030218906945601429"),
		},

		"log_10000(5000000) = 1.674742501084004701196565276318876743": {
			initialValue: mokimath.NewBigDec(5000000),
			base:         mokimath.NewBigDec(10000),
			// From: https://www.wolframalpha.com/input?i=log_10000%285000000%29+to+37+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000000002"),
			expected:             mokimath.MustNewDecFromStr("1.674742501084004701196565276318876743"),
		},
		"log_0.123456789(1) = 0": {
			initialValue: mokimath.OneDec(),
			base:         mokimath.MustNewDecFromStr("0.123456789"),

			expectedErrTolerance: mokimath.ZeroDec(),
			expected:             mokimath.ZeroDec(),
		},
		"log_1111(1111) = 1": {
			initialValue: mokimath.NewBigDec(1111),
			base:         mokimath.NewBigDec(1111),

			expectedErrTolerance: mokimath.ZeroDec(),
			expected:             mokimath.OneDec(),
		},

		"log_1.123{1024.987654321} = 59.760484327223888489694630378785099461": {
			initialValue: mokimath.NewDecWithPrec(1024987654321, 9),
			base:         mokimath.NewDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%281024.987654321%29+to+38+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000007686"),
			expected:             mokimath.MustNewDecFromStr("59.760484327223888489694630378785099461"),
		},

		"log_1.123{912648174127941279170121098210.92821920190204131121} = 594.689327867863079177915648832621538986": {
			initialValue: mokimath.MustNewDecFromStr("912648174127941279170121098210.92821920190204131121"),
			base:         mokimath.NewDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%28912648174127941279170121098210.92821920190204131121%29+to+39+digits
			expectedErrTolerance: mokimath.MustNewDecFromStr("0.000000000000000000000000000000077705"),
			expected:             mokimath.MustNewDecFromStr("594.689327867863079177915648832621538986"),
		},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			mokimath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()
				res := tc.initialValue.CustomBaseLog(tc.base)
				require.True(mokimath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestPowerInteger() {
	var expectedErrTolerance = mokimath.MustNewDecFromStr("0.000000000000000000000000000000100000")

	tests := map[string]struct {
		base           mokimath.BigDec
		exponent       uint64
		expectedResult mokimath.BigDec

		expectedToleranceOverwrite mokimath.BigDec
	}{
		"0^2": {
			base:     mokimath.ZeroDec(),
			exponent: 2,

			expectedResult: mokimath.ZeroDec(),
		},
		"1^2": {
			base:     mokimath.OneDec(),
			exponent: 2,

			expectedResult: mokimath.OneDec(),
		},
		"4^4": {
			base:     mokimath.MustNewDecFromStr("4"),
			exponent: 4,

			expectedResult: mokimath.MustNewDecFromStr("256"),
		},
		"5^3": {
			base:     mokimath.MustNewDecFromStr("5"),
			exponent: 4,

			expectedResult: mokimath.MustNewDecFromStr("625"),
		},
		"e^10": {
			base:     mokimath.EulersNumber,
			exponent: 10,

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: mokimath.MustNewDecFromStr("22026.465794806716516957900645284244366354"),
		},
		"geom twap overflow: 2^log_2{max spot price + 1}": {
			base: mokimath.TwoBigDec,
			// add 1 for simplicity of calculation to isolate overflow.
			exponent: uint64(mokimath.BigDecFromSDKDec(mokimath.MaxSpotPrice).Add(mokimath.OneDec()).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128%29%29%29+++39+digits
			expectedResult: mokimath.MustNewDecFromStr("340282366920938463463374607431768211456"),
		},
		"geom twap overflow: 2^log_2{max spot price}": {
			base:     mokimath.TwoBigDec,
			exponent: uint64(mokimath.BigDecFromSDKDec(mokimath.MaxSpotPrice).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128+-+1%29%29%29+++39+digits
			expectedResult: mokimath.MustNewDecFromStr("170141183460469231731687303715884105728"),
		},
		"geom twap overflow: 2^log_2{max spot price / 2 - 2017}": { // 2017 is prime.
			base:     mokimath.TwoBigDec,
			exponent: uint64(mokimath.BigDecFromSDKDec(mokimath.MaxSpotPrice.Quo(sdk.NewDec(2)).Sub(sdk.NewDec(2017))).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: mokimath.MustNewDecFromStr("85070591730234615865843651857942052864"),
		},

		// sdk.Dec test vectors copied from osmosis-labs/cosmos-sdk:

		"1.0 ^ (10) => 1.0": {
			base:     mokimath.OneDec(),
			exponent: 10,

			expectedResult: mokimath.OneDec(),
		},
		"0.5 ^ 2 => 0.25": {
			base:     mokimath.NewDecWithPrec(5, 1),
			exponent: 2,

			expectedResult: mokimath.NewDecWithPrec(25, 2),
		},
		"0.2 ^ 2 => 0.04": {
			base:     mokimath.NewDecWithPrec(2, 1),
			exponent: 2,

			expectedResult: mokimath.NewDecWithPrec(4, 2),
		},
		"3 ^ 3 => 27": {
			base:     mokimath.NewBigDec(3),
			exponent: 3,

			expectedResult: mokimath.NewBigDec(27),
		},
		"-3 ^ 4 = 81": {
			base:     mokimath.NewBigDec(-3),
			exponent: 4,

			expectedResult: mokimath.NewBigDec(81),
		},
		"-3 ^ 50 = 717897987691852588770249": {
			base:     mokimath.NewBigDec(-3),
			exponent: 50,

			expectedResult: mokimath.MustNewDecFromStr("717897987691852588770249"),
		},
		"-3 ^ 51 = -2153693963075557766310747": {
			base:     mokimath.NewBigDec(-3),
			exponent: 51,

			expectedResult: mokimath.MustNewDecFromStr("-2153693963075557766310747"),
		},
		"1.414213562373095049 ^ 2 = 2": {
			base:     mokimath.NewDecWithPrec(1414213562373095049, 18),
			exponent: 2,

			expectedResult:             mokimath.NewBigDec(2),
			expectedToleranceOverwrite: mokimath.MustNewDecFromStr("0.0000000000000000006"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			tolerance := expectedErrTolerance
			if !tc.expectedToleranceOverwrite.IsNil() {
				tolerance = tc.expectedToleranceOverwrite
			}

			// Main system under test
			actualResult := tc.base.PowerInteger(tc.exponent)
			require.True(mokimath.DecApproxEq(s.T(), tc.expectedResult, actualResult, tolerance))

			// Secondary system under test.
			// To reduce boilerplate from the same test cases when exponent is a
			// positive integer, we also test Power().
			// Negative exponent and base are not supported for Power()
			if tc.exponent >= 0 && !tc.base.IsNegative() {
				actualResult2 := tc.base.Power(mokimath.NewDecFromInt(mokimath.NewIntFromUint64(tc.exponent)))
				require.True(mokimath.DecApproxEq(s.T(), tc.expectedResult, actualResult2, tolerance))
			}
		})
	}
}

func (s *decimalTestSuite) TestClone() {
	tests := map[string]struct {
		startValue mokimath.BigDec
	}{
		"1.1": {
			startValue: mokimath.MustNewDecFromStr("1.1"),
		},
		"-3": {
			startValue: mokimath.MustNewDecFromStr("-3"),
		},
		"0": {
			startValue: mokimath.MustNewDecFromStr("-3"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			copy := tc.startValue.Clone()

			s.Require().Equal(tc.startValue, copy)

			copy.MulMut(mokimath.NewBigDec(2))
			// copy and startValue do not share internals.
			s.Require().NotEqual(tc.startValue, copy)
		})
	}
}

// TestMul_Mutation tests that MulMut mutates the receiver
// while Mut is not.
func (s *decimalTestSuite) TestMul_Mutation() {

	mulBy := mokimath.MustNewDecFromStr("2")

	tests := map[string]struct {
		startValue        mokimath.BigDec
		expectedMulResult mokimath.BigDec
	}{
		"1.1": {
			startValue:        mokimath.MustNewDecFromStr("1.1"),
			expectedMulResult: mokimath.MustNewDecFromStr("2.2"),
		},
		"-3": {
			startValue:        mokimath.MustNewDecFromStr("-3"),
			expectedMulResult: mokimath.MustNewDecFromStr("-6"),
		},
		"0": {
			startValue:        mokimath.ZeroDec(),
			expectedMulResult: mokimath.ZeroDec(),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.MulMut(mulBy)
			resultNonMut := startNonMut.Mul(mulBy)

			s.assertMutResult(tc.expectedMulResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

// TestPowerInteger_Mutation tests that PowerIntegerMut mutates the receiver
// while PowerInteger is not.
func (s *decimalTestSuite) TestPowerInteger_Mutation() {

	exponent := uint64(2)

	tests := map[string]struct {
		startValue     mokimath.BigDec
		expectedResult mokimath.BigDec
	}{
		"1": {
			startValue:     mokimath.OneDec(),
			expectedResult: mokimath.OneDec(),
		},
		"-3": {
			startValue:     mokimath.MustNewDecFromStr("-3"),
			expectedResult: mokimath.MustNewDecFromStr("9"),
		},
		"0": {
			startValue:     mokimath.ZeroDec(),
			expectedResult: mokimath.ZeroDec(),
		},
		"4": {
			startValue:     mokimath.MustNewDecFromStr("4.5"),
			expectedResult: mokimath.MustNewDecFromStr("20.25"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {

			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.PowerIntegerMut(exponent)
			resultNonMut := startNonMut.PowerInteger(exponent)

			s.assertMutResult(tc.expectedResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestPower() {
	tests := map[string]struct {
		base           mokimath.BigDec
		exponent       mokimath.BigDec
		expectedResult mokimath.BigDec
		expectPanic    bool
		errTolerance   mokimath.ErrTolerance
	}{
		// N.B.: integer exponents are tested under TestPowerInteger.

		"3 ^ 2 = 9 (integer base and integer exponent)": {
			base:     mokimath.NewBigDec(3),
			exponent: mokimath.NewBigDec(2),

			expectedResult: mokimath.NewBigDec(9),

			errTolerance: zeroAdditiveErrTolerance,
		},
		"2^0.5 (base of 2 and non-integer exponent)": {
			base:     mokimath.MustNewDecFromStr("2"),
			exponent: mokimath.MustNewDecFromStr("0.5"),

			// https://www.wolframalpha.com/input?i=2%5E0.5+37+digits
			expectedResult: mokimath.MustNewDecFromStr("1.414213562373095048801688724209698079"),

			errTolerance: mokimath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       mokimath.RoundDown,
			},
		},
		"3^0.33 (integer base other than 2 and non-integer exponent)": {
			base:     mokimath.MustNewDecFromStr("3"),
			exponent: mokimath.MustNewDecFromStr("0.33"),

			// https://www.wolframalpha.com/input?i=3%5E0.33+37+digits
			expectedResult: mokimath.MustNewDecFromStr("1.436977652184851654252692986409357265"),

			errTolerance: mokimath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       mokimath.RoundDown,
			},
		},
		"e^0.98999 (non-integer base and non-integer exponent)": {
			base:     mokimath.EulersNumber,
			exponent: mokimath.MustNewDecFromStr("0.9899"),

			// https://www.wolframalpha.com/input?i=e%5E0.9899+37+digits
			expectedResult: mokimath.MustNewDecFromStr("2.690965362357751196751808686902156603"),

			errTolerance: mokimath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       mokimath.RoundUnconstrained,
			},
		},
		"10^0.001 (small non-integer exponent)": {
			base:     mokimath.NewBigDec(10),
			exponent: mokimath.MustNewDecFromStr("0.001"),

			// https://www.wolframalpha.com/input?i=10%5E0.001+37+digits
			expectedResult: mokimath.MustNewDecFromStr("1.002305238077899671915404889328110554"),

			errTolerance: mokimath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       mokimath.RoundUnconstrained,
			},
		},
		"13^100.7777 (large non-integer exponent)": {
			base:     mokimath.NewBigDec(13),
			exponent: mokimath.MustNewDecFromStr("100.7777"),

			// https://www.wolframalpha.com/input?i=13%5E100.7777+37+digits
			expectedResult: mokimath.MustNewDecFromStr("1.822422110233759706998600329118969132").Mul(mokimath.NewBigDec(10).PowerInteger(112)),

			errTolerance: mokimath.ErrTolerance{
				MultiplicativeTolerance: minDecTolerance,
				RoundingDir:             mokimath.RoundDown,
			},
		},
		"large non-integer exponent with large non-integer base - panics": {
			base:     mokimath.MustNewDecFromStr("169.137"),
			exponent: mokimath.MustNewDecFromStr("100.7777"),

			expectPanic: true,
		},
		"negative base - panic": {
			base:     mokimath.NewBigDec(-3),
			exponent: mokimath.MustNewDecFromStr("4"),

			expectPanic: true,
		},
		"negative exponent - panic": {
			base:     mokimath.NewBigDec(1),
			exponent: mokimath.MustNewDecFromStr("-4"),

			expectPanic: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			mokimath.ConditionalPanic(s.T(), tc.expectPanic, func() {
				actualResult := tc.base.Power(tc.exponent)
				s.Require().Equal(0, tc.errTolerance.CompareBigDec(tc.expectedResult, actualResult))
			})
		})
	}
}
