package fuzzing

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzParseData(f *testing.F) {
	testcases := [][]byte{
		[]byte("3\nhello\ngoodbye\ngreetings\n"),
		[]byte("0\n"),
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		r := bytes.NewReader(in)
		out, err := ParseData(r)
		if err != nil {
			t.Skip("handled error")
		}
		roundTrip := ToData(out)
		rtr := bytes.NewReader(roundTrip)
		out2, _ := ParseData(rtr)
		if diff := cmp.Diff(out, out2); diff != "" {
			t.Error(diff)
		}
	})
}
