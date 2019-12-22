package unusedparam

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		expected []*Issue
		wantErr  bool
	}{
		{
			path:     "../../testdata/generated.go",
			expected: nil,
			wantErr:  true,
		},
		{
			path:     "../../testdata/empty_func.go",
			expected: []*Issue{},
		},
		{
			path: "../../testdata/assign_stmt.go",
			expected: []*Issue{
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/assign_stmt.go",
						Offset:   173,
						Line:     18,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/assign_stmt.go",
						Offset:   176,
						Line:     18,
						Column:   11,
					},
				},
			},
		},
		{
			path: "../../testdata/func_decl.go",
			expected: []*Issue{
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/func_decl.go",
						Offset:   103,
						Line:     10,
						Column:   11,
					},
				},
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/func_decl.go",
						Offset:   186,
						Line:     17,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/func_decl.go",
						Offset:   189,
						Line:     17,
						Column:   11,
					},
				},
			},
		},
		{
			path: "../../testdata/gen_decl.go",
			expected: []*Issue{
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/gen_decl.go",
						Offset:   90,
						Line:     9,
						Column:   11,
					},
				},
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/gen_decl.go",
						Offset:   149,
						Line:     14,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/gen_decl.go",
						Offset:   152,
						Line:     14,
						Column:   11,
					},
				},
			},
		},
		{
			path: "../../testdata/inc_dec_stmt.go",
			expected: []*Issue{
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/inc_dec_stmt.go",
						Offset:   74,
						Line:     9,
						Column:   11,
					},
				},
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/inc_dec_stmt.go",
						Offset:   127,
						Line:     14,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/inc_dec_stmt.go",
						Offset:   130,
						Line:     14,
						Column:   11,
					},
				},
			},
		},
		{
			path: "../../testdata/return_stmt.go",
			expected: []*Issue{
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/return_stmt.go",
						Offset:   65,
						Line:     7,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/return_stmt.go",
						Offset:   68,
						Line:     7,
						Column:   11,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/return_stmt.go",
						Offset:   125,
						Line:     11,
						Column:   11,
					},
				},
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/return_stmt.go",
						Offset:   122,
						Line:     11,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/return_stmt.go",
						Offset:   186,
						Line:     16,
						Column:   11,
					},
				},
			},
		},
		{
			path: "../../testdata/go_stmt.go",
			expected: []*Issue{
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/go_stmt.go",
						Offset:   101,
						Line:     10,
						Column:   11,
					},
				},
				&Issue{
					Text: "n is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/go_stmt.go",
						Offset:   188,
						Line:     17,
						Column:   8,
					},
				},
				&Issue{
					Text: "m is unused in _",
					Pos: token.Position{
						Filename: "../../testdata/go_stmt.go",
						Offset:   191,
						Line:     17,
						Column:   11,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			issues, err := Check(tc.path)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.ElementsMatch(t, tc.expected, issues)
		})
	}
}
