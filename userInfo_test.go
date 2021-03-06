package bbs

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/mmap"
)

func TestUserInforomTestData(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	m, err := mmap.Open("./testdata/dump.shm")
	require.NoError(err)
	defer m.Close()

	b := make([]byte, UserInfoSize)
	n, err := m.ReadAt(b, 0x42dd8)
	require.NoError(err)
	require.Equal(UserInfoSize, int64(n))

	u := &UserInfo{}
	_, err = u.ReadFrom(bytes.NewReader(b))
	require.NoError(err)
	require.EqualValues(1, u.UID)
	require.EqualValues(20436, u.PID)
	require.EqualValues(0, u.Socket)
	require.EqualValues(65535, u.Level)
	require.EqualValues("SYSOP", u.ID)
	require.EqualValues("1.164.105.65", u.From)

	t.Logf("%+v", u)
}

func TestUserInfoStruct(t *testing.T) {
	t.Parallel()

	t.Run("userData", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &userData{
			Level:      123456789,
			ID:         "ptt5566",
			Nickname:   "56 never die",
			From:       "127.0.0.1",
			FromIp:     456789,
			DarkWin:    1,
			DarkLose:   2,
			DarkTie:    3,
			Angelpause: 4,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(userDataSize, int64(n))

		actual := &userData{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(userDataSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("friendShip", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &friendShip{
			FriendCap: maxUserSize,
			FriendLen: 10,
		}

		for i := 0; i < int(except.FriendLen); i++ {
			except.Friends[i] = int32(i + 5566)
			except.FriendOnline[i] = uint32(i + 1234)
			except.Rejects[i] = int32(i + 456)
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(friendShipSize, int64(n))

		actual := &friendShip{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(friendShipSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("MsgQue", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &MsgQue{
			PID:        9408,
			ID:         "imyourfather",
			LastCallIn: "asdqwertewfxcvsdfasd",
			Mode:       777,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(msgQueSize, int64(n))

		actual := &MsgQue{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(msgQueSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("message", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &message{
			Count: 7,
		}
		for i := 0; i < int(except.Count); i++ {
			except.List[i] = MsgQue{
				PID: int32(i),
			}
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(messageSize, int64(n))

		actual := &message{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(messageSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("userStatus", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &userStatus{
			Birthday:  1,
			Active:    2,
			Invisible: 3,
			Mode:      4,
			Pager:     5,
			Conn6Win:  6,
			Lastact:   7,
			Alerts:    8,
			Conn6Lose: 9,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(userStatusSize, int64(n))

		actual := &userStatus{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(userStatusSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("chatRoom", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &chatRoom{
			Sig:        1,
			Conn6Tie:   2,
			Destuid:    3,
			Destuip:    4,
			Sockactive: 5,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(chatRoomSize, int64(n))

		actual := &chatRoom{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(chatRoomSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("chat", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &chat{
			InChat: 128,
			ChatID: "asdfghjklmn", // over length
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(chatSize, int64(n))

		actual := &chat{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(chatSize, int64(n))
		require.EqualValues("asdfghjklm", actual.ChatID)
		except.ChatID = "asdfghjklm"
		require.EqualValues(except, actual)
	})
	t.Run("game", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &game{
			LockMode: 3,
			Turn:     111,
			MateID:   "skywalker",
			Color:    11,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(gameSize, int64(n))

		actual := &game{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(gameSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("gameRecord", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &gameRecord{
			FiveWin:        123,
			FiveLose:       456,
			FiveTie:        789,
			ChcWin:         234,
			ChcLose:        567,
			ChcTie:         890,
			ChessEloRating: 135,
			GoWin:          246,
			GoLose:         357,
			GoTie:          468,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(gameRecordSize, int64(n))

		actual := &gameRecord{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(gameRecordSize, int64(n))
		require.EqualValues(except, actual)
	})
	t.Run("misc", func(t *testing.T) {
		t.Parallel()
		require := require.New(t)
		except := &misc{
			WithMe: 654,
			BrcID:  987,
		}

		b := bytes.NewBuffer(nil)
		n, err := except.WriteTo(b)
		require.NoError(err)
		require.EqualValues(miscSize, int64(n))

		actual := &misc{}
		n, err = actual.ReadFrom(bytes.NewReader(b.Bytes()))
		require.NoError(err)
		require.EqualValues(miscSize, int64(n))
		require.EqualValues(except, actual)
	})
}
