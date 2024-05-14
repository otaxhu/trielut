package trielut

import "testing"

func TestInsert(t *testing.T) {
	testCases := []struct {
		Name           string
		Buffers        [][]byte
		ExpectedReturn bool
	}{
		{
			Name: "Success_OneEntry",
			Buffers: [][]byte{
				[]byte("cat"),
			},
			ExpectedReturn: true,
		},
		{
			Name: "Success_MultipleEntries_NotOverlapping",
			Buffers: [][]byte{
				[]byte("cat"),
				[]byte("bob"),
				[]byte("airplane"),
			},
			ExpectedReturn: true,
		},
		{
			Name: "Success_MultipleEntries_Overlapping",
			Buffers: [][]byte{
				[]byte("car"),
				[]byte("carl"),
				[]byte("carlson"),
			},
			ExpectedReturn: true,
		},
		{
			// Support null value as entry for the trie, below we test that the trie doesn't allow
			// multiple null values at the same time
			Name: "Success_NilEntry",
			Buffers: [][]byte{
				nil,
			},
			ExpectedReturn: true,
		},
		{
			Name: "Failure_TwoEntriesEquals",
			Buffers: [][]byte{
				[]byte("cat"),
				[]byte("cat"),
			},
			ExpectedReturn: false,
		},
		{
			Name: "Failure_TwoNilEntries",
			Buffers: [][]byte{
				nil, nil,
			},
			ExpectedReturn: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			trie := &Trie{}
			for i, b := range tc.Buffers {
				ret := trie.Insert(b)
				if ret != tc.ExpectedReturn && i == (len(tc.Buffers)-1) {
					t.Errorf("unexpected return value, expected %v got %v", tc.ExpectedReturn, ret)
				}
			}
		})
	}
}

func TestHas(t *testing.T) {
	// TODO
}

func TestDelete(t *testing.T) {
	// TODO
}
