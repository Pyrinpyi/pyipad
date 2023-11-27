// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util_test

import (
	"math"
	"testing"

	"github.com/Pyrinpyi/pyipad/domain/consensus/utils/constants"
)

func TestAmountCreation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		valid    bool
		expected Amount
	}{
		// Positive tests.
		{
			name:     "zero",
			amount:   0,
			valid:    true,
			expected: 0,
		},
		{
			name:     "max producible",
			amount:   29e9,
			valid:    true,
			expected: Amount(constants.MaxLeor),
		},
		{
			name:     "one hundred",
			amount:   100,
			valid:    true,
			expected: 100 * constants.LeorPerPyrin,
		},
		{
			name:     "fraction",
			amount:   0.01234567,
			valid:    true,
			expected: 1234567,
		},
		{
			name:     "rounding up",
			amount:   54.999999999999943157,
			valid:    true,
			expected: 55 * constants.LeorPerPyrin,
		},
		{
			name:     "rounding down",
			amount:   55.000000000000056843,
			valid:    true,
			expected: 55 * constants.LeorPerPyrin,
		},

		// Negative tests.
		{
			name:   "not-a-number",
			amount: math.NaN(),
			valid:  false,
		},
		{
			name:   "-infinity",
			amount: math.Inf(-1),
			valid:  false,
		},
		{
			name:   "+infinity",
			amount: math.Inf(1),
			valid:  false,
		},
	}

	for _, test := range tests {
		a, err := NewAmount(test.amount)
		switch {
		case test.valid && err != nil:
			t.Errorf("%v: Positive test Amount creation failed with: %v", test.name, err)
			continue
		case !test.valid && err == nil:
			t.Errorf("%v: Negative test Amount creation succeeded (value %v) when should fail", test.name, a)
			continue
		}

		if a != test.expected {
			t.Errorf("%v: Created amount %v does not match expected %v", test.name, a, test.expected)
			continue
		}
	}
}

func TestAmountUnitConversions(t *testing.T) {
	tests := []struct {
		name      string
		amount    Amount
		unit      AmountUnit
		converted float64
		s         string
	}{
		{
			name:      "MPYI",
			amount:    Amount(constants.MaxLeor),
			unit:      AmountMegaPYI,
			converted: 29000,
			s:         "29000 MPYI",
		},
		{
			name:      "kPYI",
			amount:    44433322211100,
			unit:      AmountKiloPYI,
			converted: 444.33322211100,
			s:         "444.333222111 kPYI",
		},
		{
			name:      "PYI",
			amount:    44433322211100,
			unit:      AmountPYI,
			converted: 444333.22211100,
			s:         "444333.222111 PYI",
		},
		{
			name:      "mPYI",
			amount:    44433322211100,
			unit:      AmountMilliPYI,
			converted: 444333222.11100,
			s:         "444333222.111 mPYI",
		},
		{

			name:      "μPYI",
			amount:    44433322211100,
			unit:      AmountMicroPYI,
			converted: 444333222111.00,
			s:         "444333222111 μPYI",
		},
		{

			name:      "leor",
			amount:    44433322211100,
			unit:      AmountLeor,
			converted: 44433322211100,
			s:         "44433322211100 Leor",
		},
		{

			name:      "non-standard unit",
			amount:    44433322211100,
			unit:      AmountUnit(-1),
			converted: 4443332.2211100,
			s:         "4443332.22111 1e-1 PYI",
		},
	}

	for _, test := range tests {
		f := test.amount.ToUnit(test.unit)
		if f != test.converted {
			t.Errorf("%v: converted value %v does not match expected %v", test.name, f, test.converted)
			continue
		}

		s := test.amount.Format(test.unit)
		if s != test.s {
			t.Errorf("%v: format '%v' does not match expected '%v'", test.name, s, test.s)
			continue
		}

		// Verify that Amount.ToPYI works as advertised.
		f1 := test.amount.ToUnit(AmountPYI)
		f2 := test.amount.ToPYI()
		if f1 != f2 {
			t.Errorf("%v: ToPYI does not match ToUnit(AmountPYI): %v != %v", test.name, f1, f2)
		}

		// Verify that Amount.String works as advertised.
		s1 := test.amount.Format(AmountPYI)
		s2 := test.amount.String()
		if s1 != s2 {
			t.Errorf("%v: String does not match Format(AmountPYI): %v != %v", test.name, s1, s2)
		}
	}
}

func TestAmountMulF64(t *testing.T) {
	tests := []struct {
		name string
		amt  Amount
		mul  float64
		res  Amount
	}{
		{
			name: "Multiply 0.1 PYI by 2",
			amt:  100e5, // 0.1 PYI
			mul:  2,
			res:  200e5, // 0.2 PYI
		},
		{
			name: "Multiply 0.2 PYI by 0.02",
			amt:  200e5, // 0.2 PYI
			mul:  1.02,
			res:  204e5, // 0.204 PYI
		},
		{
			name: "Round down",
			amt:  49, // 49 Leor
			mul:  0.01,
			res:  0,
		},
		{
			name: "Round up",
			amt:  50, // 50 Leor
			mul:  0.01,
			res:  1, // 1 Leor
		},
		{
			name: "Multiply by 0.",
			amt:  1e8, // 1 PYI
			mul:  0,
			res:  0, // 0 PYI
		},
		{
			name: "Multiply 1 by 0.5.",
			amt:  1, // 1 Leor
			mul:  0.5,
			res:  1, // 1 Leor
		},
		{
			name: "Multiply 100 by 66%.",
			amt:  100, // 100 Leor
			mul:  0.66,
			res:  66, // 66 Leor
		},
		{
			name: "Multiply 100 by 66.6%.",
			amt:  100, // 100 Leor
			mul:  0.666,
			res:  67, // 67 Leor
		},
		{
			name: "Multiply 100 by 2/3.",
			amt:  100, // 100 Leor
			mul:  2.0 / 3,
			res:  67, // 67 Leor
		},
	}

	for _, test := range tests {
		a := test.amt.MulF64(test.mul)
		if a != test.res {
			t.Errorf("%v: expected %v got %v", test.name, test.res, a)
		}
	}
}
