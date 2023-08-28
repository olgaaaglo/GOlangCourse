package main 

import ( 
	"testing" 
	)

func Test_split(t *testing.T) {
	res, err := split("aaa/bbb/ccc", "/")
	expected := []string{"aaa", "bbb", "ccc"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function split, res = ", res,
				"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function split - error = ", err)
	}

	res, err = split("aaa bbb ccc")
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function split, res = ", res,
			"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function split - error = ", err)
	}

	res, err = split("aaa bbb ccc", "", " ")
	
	if err == nil {
		t.Error("Error in function split - no error")
	}
}

func Test_splitAndSort(t *testing.T) {
	res, err := splitAndSort("bbb ccc aaa")
	expected := []string{"aaa", "bbb", "ccc"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function splitAndSort, res = ", res,
				"expected = ", expected)
		}
	}
	if err != nil {
		t.Error("Error in function splitAndSort - error = ", err)
	}
}

func Test_getNumberOfNonEmptyLines(t *testing.T) {
	res := getNumberOfNonEmptyLines("aaa\n\nbbb\nccc\n\nd")
	expected := 4
	if res != expected {
		t.Error("Error in function getNumberOfNonEmptyLines, res = ", res,
			"expected = ", expected)
	}
}

func Test_getNumOfWords(t *testing.T) {
	res := getNumOfWords("aaa bbb\n ccc ddd")
	expected := 4
	if res != expected {
		t.Error("Error in function getNumOfWords, res = ", res,
			"expected = ", expected)
	}
}

func Test_getNumOfNonWhiteChars(t *testing.T) {
	res := getNumOfNonWhiteChars("aaa bbb\n ccc d\te")
	expected := 11
	if res != expected {
		t.Error("Error in function getNumOfNonWhiteChars, res = ", res,
			"expected = ", expected)
	}
}

func Test_getFrequencyForWords(t *testing.T) {
	res := getFrequencyForWords("aaa bbb aaa ccc bbb aaa")
	expected := map[string] int{"aaa":3, "bbb":2, "ccc":1}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
	}
	for key, val := range res {
		if val != expected[key] {
			t.Error("Error in function getFrequencyForWords, res = ", res,
				"expected = ", expected)
		}
	}
}

func Test_getSortedPalindroms(t *testing.T) {
	res := getSortedPalindroms("bccb cda aaa aba dcdc")
	expected := []string{"aaa", "aba", "bccb"}
	if len(res) != len(expected) {
		t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
	}
	for i, val := range res {
		if val != expected[i] {
			t.Error("Error in function getSortedPalindroms, res = ", res,
				"expected = ", expected)
		}
	}
}