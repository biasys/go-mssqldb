package mssql

import (
	"testing"
)

func TestParseParams(t *testing.T) {
	values := []struct {
		s string
		d string
		n int
	}{
		{"select ?", "select @p1", 1},
		{"select ?, ?", "select @p1, @p2", 2},
		{"select ? -- ?", "select @p1 -- ?", 1},
		{"select ? - ?", "select @p1 - @p2", 2},
		{"select ? /* ? */, ?", "select @p1 /* ? */, @p2", 2},
		{"select ? /* ? * ? */, ?", "select @p1 /* ? * ? */, @p2", 2},
		{"select \"foo?\", [foo?], 'foo?', ?", "select \"foo?\", [foo?], 'foo?', @p1", 1},
		{"select \"x\"\"y\", [x]]y], 'x''y', ?", "select \"x\"\"y\", [x]]y], 'x''y', @p1", 1},
	}

	for _, v := range values {
		d, n := parseParams(v.s)
		if d != v.d {
			t.Error("Parse params don't match ", d, v.d)
		}
		if n != v.n {
			t.Error("Parse number of params don't match", n, v.n)
		}
	}
}
