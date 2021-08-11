// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package override

import (
	"bytes"
	"testing"
)

type marshalTest struct {
	input interface{}
	want  string
}

type unmarshalTest struct {
	input        string
	want         interface{}
	wantErr      error // if set, decoding must fail on any platform
	wantErr32bit error // if set, decoding must fail on 32bit platforms (used for Uint tests)
}

var (
	encodeBytesTests = []marshalTest{
		{[]byte{}, "Ex"},
		{[]byte{0}, "Ex00"},
		{[]byte{0, 0, 1, 2}, "Ex00000102"},
	}

	decodeBytesTests = []unmarshalTest{
		// invalid
		{input: ``, wantErr: ErrEmptyString},
		{input: `0`, wantErr: ErrMissingPrefix},
		{input: `Ex0`, wantErr: ErrOddLength},
		{input: `Ex023`, wantErr: ErrOddLength},
		{input: `Exxx`, wantErr: ErrSyntax},
		{input: `Ex01zz01`, wantErr: ErrSyntax},
		// valid
		{input: `Ex`, want: []byte{}},
		{input: `EX`, want: []byte{}},
		{input: `Ex02`, want: []byte{0x02}},
		{input: `EX02`, want: []byte{0x02}},
		{input: `Exffffffffff`, want: []byte{0xff, 0xff, 0xff, 0xff, 0xff}},
		{
			input: `Exffffffffffffffffffffffffffffffffffff`,
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}
)

func TestEncode(t *testing.T) {
	for _, test := range encodeBytesTests {
		enc := Encode(test.input.([]byte))
		if enc != test.want {
			t.Errorf("input %x: wrong encoding %s", test.input, enc)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range decodeBytesTests {
		dec, err := Decode(test.input)
		if !checkError(t, test.input, err, test.wantErr) {
			continue
		}
		if !bytes.Equal(test.want.([]byte), dec) {
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, dec, test.want)
			continue
		}
	}
}
