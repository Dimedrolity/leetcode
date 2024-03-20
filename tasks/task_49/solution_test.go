package task_49

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	name     string
	strs     []string
	expected [][]string
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []TestCase{
		{"Группировка анаграмм строк одинаковой длины",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},

		{"Группировка анаграмм строк разной длины",
			[]string{"a", "ab", "ba"},
			[][]string{{"a"}, {"ab", "ba"}},
		},

		{"Пустая строка",
			[]string{""},
			[][]string{{""}},
		},

		{"Единственная строка",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			t.Log(testCase.name)

			actual := groupAnagrams(testCase.strs)
			assert.ElementsMatch(t, testCase.expected, actual)
		})
	}
}
