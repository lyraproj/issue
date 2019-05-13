package issue_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lyraproj/issue/issue"
)

func TestParseLocation_full(t *testing.T) {
	loc := issue.ParseLocation(`(file: /foo/fee, line: 23, column: 4)`)
	require.Equal(t, `/foo/fee`, loc.File())
	require.Equal(t, 23, loc.Line())
	require.Equal(t, 4, loc.Pos())
}

func TestParseLocation_file_line(t *testing.T) {
	loc := issue.ParseLocation(`(file: /foo/fee, line: 23)`)
	require.Equal(t, `/foo/fee`, loc.File())
	require.Equal(t, 23, loc.Line())
	require.Equal(t, 0, loc.Pos())
}

func TestParseLocation_file(t *testing.T) {
	loc := issue.ParseLocation(`(file: /foo/fee)`)
	require.Equal(t, `/foo/fee`, loc.File())
	require.Equal(t, 0, loc.Line())
	require.Equal(t, 0, loc.Pos())
}

func TestParseLocation_line_pos(t *testing.T) {
	loc := issue.ParseLocation(`(line: 23, column: 4)`)
	require.Equal(t, ``, loc.File())
	require.Equal(t, 23, loc.Line())
	require.Equal(t, 4, loc.Pos())
}

func TestParseLocation_line(t *testing.T) {
	loc := issue.ParseLocation(`(line: 23)`)
	require.Equal(t, ``, loc.File())
	require.Equal(t, 23, loc.Line())
	require.Equal(t, 0, loc.Pos())
}
