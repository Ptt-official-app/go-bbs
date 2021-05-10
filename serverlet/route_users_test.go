package main

import (
	"github.com/Ptt-official-app/go-bbs"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockUserRecord struct {
	userID string
}

func NewMockUserRecord(userID string) *MockUserRecord { return &MockUserRecord{userID: userID} }
func (u *MockUserRecord) UserID() string              { return u.userID }

// HashedPassword return user hashed password, it only for debug,
// If you want to check is user password correct, please use
// VerifyPassword insteaded.
func (u *MockUserRecord) HashedPassword() string { return "" }

// VerifyPassword will check user's password is OK. it will return null
// when OK and error when there are something wrong
func (u *MockUserRecord) VerifyPassword(password string) error { return nil }

// Nickname return a string for user's nickname, this string may change
// depend on user's mood, return empty string if this bbs system do not support
func (u *MockUserRecord) Nickname() string { return "" }

// RealName return a string for user's real name, this string may not be changed
// return empty string if this bbs system do not support
func (u *MockUserRecord) RealName() string { return "" }

// NumLoginDays return how many days this have been login since account created.
func (u *MockUserRecord) NumLoginDays() int { return 0 }

// NumPosts return how many posts this user has posted.
func (u *MockUserRecord) NumPosts() int { return 0 }

// Money return the money this user have.
func (u *MockUserRecord) Money() int { return 0 }

// LastLogin return last login time of user
func (u *MockUserRecord) LastLogin() time.Time { return time.Now() }

// LastHost return last login host of user, it is IPv4 address usually, but it
// could be domain name or IPv6 address.
func (u *MockUserRecord) LastHost() string { return "" }

// UserFlag return user setting.
// uint32, see https://github.com/ptt/pttbbs/blob/master/include/uflags.h
func (u *MockUserRecord) UserFlag() uint32 { return 0x00000001 }

func TestGetUserInformation(t *testing.T) {

	expected := NewMockUserRecord("SYSOP")

	userRecs = []bbs.UserRecord{
		expected,
	}

	req, err := http.NewRequest("GET", "/v1/users/SYSOP/information", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := newAccessTokenWithUsername(expected.UserID())
	t.Logf("testing token: %v", token)
	req.Header.Add("Authorization", "bearer "+token)

	rr := httptest.NewRecorder()
	r := http.NewServeMux()
	r.HandleFunc("/v1/users/", routeUsers)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	responsedMap := map[string]interface{}{}
	json.Unmarshal(rr.Body.Bytes(), &responsedMap)
	t.Logf("got response %v", rr.Body.String())
	responsedData := responsedMap["data"].(map[string]interface{})
	if responsedData["user_id"] != expected.UserID() {
		t.Errorf("handler returned unexpected body, user_id not match: got %v want userID %v",
			rr.Body.String(), expected.UserID())

	}

}

func TestParseUserPath(t *testing.T) {

	type TestCase struct {
		input         string
		expectdUserID string
		expectdItem   string
	}

	cases := []TestCase{
		{
			input:         "/v1/users/Pichu/information",
			expectdUserID: "Pichu",
			expectdItem:   "information",
		},
		{
			input:         "/v1/users/Pichu/",
			expectdUserID: "Pichu",
			expectdItem:   "",
		},
		{
			input:         "/v1/users/Pichu",
			expectdUserID: "Pichu",
			expectdItem:   "",
		},
	}

	for index, c := range cases {
		input := c.input
		expectdUserID := c.expectdUserID
		expectdItem := c.expectdItem
		actualUserID, actualItem, err := parseUserPath(input)
		if err != nil {
			t.Errorf("error on index %d, got: %v", index, err)

		}

		if actualUserID != expectdUserID {
			t.Errorf("userID not match on index %d, expected: %v, got: %v", index, expectdUserID, actualUserID)
		}

		if actualItem != expectdItem {
			t.Errorf("item not match on index %d, expected: %v, got: %v", index, expectdItem, actualItem)
		}

	}

}
