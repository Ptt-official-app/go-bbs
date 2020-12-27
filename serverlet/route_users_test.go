package main

import (
	"testing"
)

func TestParseUserPath(t *testing.T) {

	type TestCase struct {
		input         string
		expectdUserId string
		expectdItem   string
	}

	cases := []TestCase{
		{
			input:         "/v1/users/Pichu/information",
			expectdUserId: "Pichu",
			expectdItem:   "information",
		},
		{
			input:         "/v1/users/Pichu/",
			expectdUserId: "Pichu",
			expectdItem:   "",
		},
		{
			input:         "/v1/users/Pichu",
			expectdUserId: "Pichu",
			expectdItem:   "",
		},
	}

	for index, c := range cases {
		input := c.input
		expectdUserId := c.expectdUserId
		expectdItem := c.expectdItem
		actualUserId, actualItem, err := parseUserPath(input)
		if err != nil {
			t.Errorf("error on index %d, got: %v", index, err)

		}

		if actualUserId != expectdUserId {
			t.Errorf("userId not match on index %d, expected: %v, got: %v", index, expectdUserId, actualUserId)
		}

		if actualItem != expectdItem {
			t.Errorf("item not match on index %d, expected: %v, got: %v", index, expectdItem, actualItem)
		}

	}

}
