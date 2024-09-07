/*
	mrg63k3a.go - an implementation of the 63bit L'Ecuyer PRNG
	Copyright (C) 2018 Mason Marchildon <mason@riffle.ca>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package mrg63k3a

import "math/rand"

const (
	m1   uint64 = 9223372036854769163 // 2^63-6645
	m2   uint64 = 9223372036854754679 // 2^63-21129
	a12  uint64 = 1754669720
	a13n uint64 = 3182104042
	a21  uint64 = 31387477935
	a23n uint64 = 6199136374
)

// Please note that design of this package follows the work completed
// by Jochen Voss on his 64bit Mersenne Twister PRNG, see
// <https://github.com/seehuhn/mt19937>
//
// With his kind permission, I offer you this alternative 63bit PRNG.
//
// (Pseudo-)random number generator (or combined multiple recursive
// generator - CMRG) developed by: L'Ecuyer, P. Good parameters and
// implementations for combined multiple recursive random number
// generators. Operations Research, 47(1):159-164, 1999.
//
// This PRNG has a period of 2^377 (compared to the 2^19937-1 period
// of the Mersenne Twister PRNG). The following code is based on the
// c code is given in L'Ecuyer (1999), but follows the modified
// version made in Lemieux, C. (2009) Monte Carlo and Quasi-Monte
// Carlo Sampling. Springer Science. 373pp. [page 63.]

// MRG63k3a is the structure to hold the state of one instance of
// L'Ecuyer's PRNG.  New instances can be allocated using the
// mrg63k3a.New() function.  MRG63k3a implements the rand.Source
// interface and rand.New() from the math/rand package can be used to
// generate different distributions from a MRG63k3a PRNG.
//
// This class is not safe for concurrent accesss by different
// goroutines.  If more than one goroutine accesses the PRNG, the
// callers must synchronise access using sync.Mutex or similar.
type MRG63k3a struct {
	s [2][3]uint64
}

// New allocates a new instance of a 63bit L'Ecuyer's PRNG.
// A seed can be set using the .Seed() method.
func New() *MRG63k3a {
	return &MRG63k3a{}
}

// Seed for s00, s01, s02, s10, s11, and s12 are derrived using Go's
// internal PRNG, see L'Ecuyer (1999) for rules on choosing these seeds.
func (mrg *MRG63k3a) Seed(seed int64) {
	var r = rand.New(rand.NewSource(seed))
	for j := 0; j < 3; j++ {
	again0:
		f := r.Intn(int(m1))
		if f == 0 {
			goto again0
		}
		mrg.s[0][j] = uint64(f)
	}
	for j := 0; j < 3; j++ {
	again1:
		f := r.Intn(int(m2))
		if f == 0 {
			goto again1
		}
		mrg.s[1][j] = uint64(f)
	}
}

func (mrg *MRG63k3a) Uint64() uint64 {
	return uint64(mrg.Int63())>>31 | uint64(mrg.Int63())<<32
}

// Int63 generates a (pseudo-)random 63bit value.  The output can be
// used as a replacement for a sequence of independent, uniformly
// distributed samples in the range 0, 1, ..., 2^63-6645 (m1).  This method is
// part of the rand.Source interface.
func (mrg *MRG63k3a) Int63() int64 {
	mrg.s[0][2] = (a12*mrg.s[0][1] - a13n*mrg.s[0][0]) % m1
	mrg.s[1][2] = (a21*mrg.s[1][1] - a23n*mrg.s[1][0]) % m2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			mrg.s[i][j] = mrg.s[i][j+1]
		}
	}
	return int64((mrg.s[0][2] - mrg.s[1][2]) % m1)
}
