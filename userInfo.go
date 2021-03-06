package bbs

import (
	"encoding/binary"
	"io"
	"strings"
	"unicode"
)

func clangStringToGolangString(b []byte) string {
	// enhancement: maybe using regex instead of character checking
	return strings.TrimFunc(string(b), func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r) && !unicode.IsSpace(r)
	})
}

func golangStringToClangString(s string, size int) []byte {
	b := make([]byte, size)
	if len(s) >= size {
		copy(b, s[:size-1])
	} else {
		copy(b, s)
	}

	return b
}

var order = binary.LittleEndian

const (
	userDataSize   int64 = 80
	friendShipSize int64 = 2196
	msgQueSize     int64 = 97
	messageSize    int64 = (maxMsgQueNum * msgQueSize) + 4
	userStatusSize int64 = 21
	chatRoomSize   int64 = 12
	chatSize       int64 = 12
	gameSize       int64 = 16
	gameRecordSize int64 = 20
	miscSize       int64 = 8
	UserInfoSize   int64 = userDataSize + friendShipSize + msgQueSize + messageSize +
		userStatusSize + chatRoomSize + chatSize + gameSize + gameRecordSize + miscSize + 20 // should be 3456 byte
)

type userData struct {
	Level    uint32
	ID       string
	Nickname string
	From     string
	FromIp   uint32

	DarkWin    uint16
	DarkLose   uint16
	DarkTie    uint16
	Angelpause uint8
}

func (t *userData) ReadFrom(r io.Reader) (_ int64, err error) {
	var gap byte
	if err = binary.Read(r, order, &t.Level); err != nil {
		return
	}
	id := make([]byte, 13)
	if err = binary.Read(r, order, id); err != nil {
		return
	}
	t.ID = clangStringToGolangString(id)
	nickname := make([]byte, 24)
	if err = binary.Read(r, order, nickname); err != nil {
		return
	}
	t.Nickname = clangStringToGolangString(nickname)
	from := make([]byte, 27)
	if err = binary.Read(r, order, from); err != nil {
		return
	}
	t.From = clangStringToGolangString(from)
	if err = binary.Read(r, order, &t.FromIp); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.DarkWin); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.DarkLose); err != nil {
		return
	}
	if err = binary.Read(r, order, &gap); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Angelpause); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.DarkTie); err != nil {
		return
	}

	return userDataSize, nil
}

func (t *userData) WriteTo(w io.Writer) (_ int64, err error) {
	var gap byte
	if err = binary.Write(w, order, t.Level); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.ID, 13)); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.Nickname, 24)); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.From, 27)); err != nil {
		return
	}
	if err = binary.Write(w, order, t.FromIp); err != nil {
		return
	}
	if err = binary.Write(w, order, t.DarkWin); err != nil {
		return
	}
	if err = binary.Write(w, order, t.DarkLose); err != nil {
		return
	}
	if err = binary.Write(w, order, gap); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Angelpause); err != nil {
		return
	}
	if err = binary.Write(w, order, t.DarkTie); err != nil {
		return
	}

	return userDataSize, nil
}

const (
	maxUserSize   = 256
	maxRejectSize = 32
)

type friendShip struct {
	FriendCap    int32
	FriendLen    int16
	Friends      [maxUserSize]int32
	FriendOnline [maxUserSize]uint32
	Rejects      [maxRejectSize]int32
}

func (t *friendShip) ReadFrom(r io.Reader) (_ int64, err error) {
	var gap2 uint16
	var gap4 uint32
	if err = binary.Read(r, order, &t.FriendCap); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.FriendLen); err != nil {
		return
	}
	if err = binary.Read(r, order, &gap2); err != nil {
		return
	}
	for i := 0; i < maxUserSize; i++ {
		var friend int32
		if err = binary.Read(r, order, &friend); err != nil {
			return
		}
		t.Friends[i] = friend

	}
	if err = binary.Read(r, order, &gap4); err != nil {
		return
	}
	for i := 0; i < maxUserSize; i++ {
		var friend uint32
		if err = binary.Read(r, order, &friend); err != nil {
			return
		}
		t.FriendOnline[i] = friend
	}
	if err = binary.Read(r, order, &gap4); err != nil {
		return
	}
	for i := 0; i < maxRejectSize; i++ {
		var reject int32
		if err = binary.Read(r, order, &reject); err != nil {
			return
		}
		t.Rejects[i] = reject
	}

	return friendShipSize, nil
}

func (t *friendShip) WriteTo(w io.Writer) (_ int64, err error) {
	var gap2 uint16
	var gap4 uint32
	if err = binary.Write(w, order, t.FriendCap); err != nil {
		return
	}
	if err = binary.Write(w, order, t.FriendLen); err != nil {
		return
	}
	if err = binary.Write(w, order, gap2); err != nil {
		return
	}
	for _, friend := range t.Friends {
		if err = binary.Write(w, order, friend); err != nil {
			return
		}
	}
	if err = binary.Write(w, order, gap4); err != nil {
		return
	}
	for _, friend := range t.FriendOnline {
		if err = binary.Write(w, order, friend); err != nil {
			return
		}
	}
	if err = binary.Write(w, order, gap4); err != nil {
		return
	}
	for _, reject := range t.Rejects {
		if err = binary.Write(w, order, reject); err != nil {
			return
		}
	}
	if err = binary.Write(w, order, gap4); err != nil {
		return
	}

	return friendShipSize, nil
}

type MsgQue struct {
	PID        int32
	ID         string
	LastCallIn string
	Mode       int32
}

func (t *MsgQue) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.PID); err != nil {
		return
	}
	id := make([]byte, 13)
	if err = binary.Read(r, order, id[:]); err != nil {
		return
	}
	t.ID = clangStringToGolangString(id)
	callIn := make([]byte, 76)
	if err = binary.Read(r, order, callIn[:]); err != nil {
		return
	}
	t.LastCallIn = clangStringToGolangString(callIn)
	if err = binary.Read(r, order, &t.Mode); err != nil {
		return
	}

	return msgQueSize, nil
}

func (t *MsgQue) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.PID); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.ID, 13)); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.LastCallIn, 76)); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Mode); err != nil {
		return
	}

	return msgQueSize, nil
}

const maxMsgQueNum = 10

type message struct {
	Count int8
	List  [maxMsgQueNum]MsgQue
}

func (t *message) ReadFrom(r io.Reader) (_ int64, err error) {
	var gap3 [3]byte
	var gap MsgQue
	if err = binary.Read(r, order, &t.Count); err != nil {
		return
	}
	if err = binary.Read(r, order, gap3[:]); err != nil {
		return
	}
	for i := 0; i < maxMsgQueNum; i++ {
		if _, err = t.List[i].ReadFrom(r); err != nil {
			return
		}
	}
	if _, err = gap.ReadFrom(r); err != nil {
		return
	}

	return messageSize, nil
}

func (t *message) WriteTo(w io.Writer) (_ int64, err error) {
	var gap3 [3]byte
	var gap MsgQue
	if err = binary.Write(w, order, t.Count); err != nil {
		return
	}
	if err = binary.Write(w, order, gap3[:]); err != nil {
		return
	}
	for _, msg := range t.List {
		if _, err = msg.WriteTo(w); err != nil {
			return
		}
	}
	if _, err = gap.WriteTo(w); err != nil {
		return
	}

	return messageSize, nil
}

type userStatus struct {
	Birthday  int8
	Active    uint8
	Invisible uint8
	Mode      uint8
	Pager     uint8
	Conn6Win  uint16
	Lastact   int32
	Alerts    int8
	Conn6Lose uint16
}

func (t *userStatus) ReadFrom(r io.Reader) (_ int64, err error) {
	var gap1 byte
	if err = binary.Read(r, order, &t.Birthday); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Active); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Invisible); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Mode); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Pager); err != nil {
		return
	}
	if err = binary.Read(r, order, &gap1); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Conn6Win); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Lastact); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Alerts); err != nil {
		return
	}
	if err = binary.Read(r, order, &gap1); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Conn6Lose); err != nil {
		return
	}
	if err = binary.Read(r, order, &gap1); err != nil {
		return
	}

	return userStatusSize, nil
}

func (t *userStatus) WriteTo(w io.Writer) (_ int64, err error) {
	var gap1 byte
	if err = binary.Write(w, order, t.Birthday); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Active); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Invisible); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Mode); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Pager); err != nil {
		return
	}
	if err = binary.Write(w, order, gap1); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Conn6Win); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Lastact); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Alerts); err != nil {
		return
	}
	if err = binary.Write(w, order, gap1); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Conn6Lose); err != nil {
		return
	}
	if err = binary.Write(w, order, gap1); err != nil {
		return
	}

	return userStatusSize, nil
}

type chatRoom struct {
	Sig        uint8
	Conn6Tie   uint16
	Destuid    int32
	Destuip    int32
	Sockactive uint8
}

func (t *chatRoom) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.Sig); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Conn6Tie); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Destuid); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Destuip); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Sockactive); err != nil {
		return
	}

	return chatRoomSize, nil
}

func (t *chatRoom) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.Sig); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Conn6Tie); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Destuid); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Destuip); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Sockactive); err != nil {
		return
	}

	return chatRoomSize, nil
}

type chat struct {
	InChat uint8
	ChatID string
}

func (t *chat) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.InChat); err != nil {
		return
	}
	id := make([]byte, 11)
	if err = binary.Read(r, order, &id); err != nil {
		return
	}
	t.ChatID = clangStringToGolangString(id)

	return chatSize, nil
}

func (t *chat) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.InChat); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.ChatID, 11)); err != nil {
		return
	}

	return chatSize, nil
}

type game struct {
	LockMode uint8
	Turn     int8
	MateID   string
	Color    int8
}

func (t *game) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.LockMode); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Turn); err != nil {
		return
	}
	id := make([]byte, 13)
	if err = binary.Read(r, order, &id); err != nil {
		return
	}
	t.MateID = clangStringToGolangString(id)
	if err = binary.Read(r, order, &t.Color); err != nil {
		return
	}

	return gameSize, nil
}

func (t *game) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.LockMode); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Turn); err != nil {
		return
	}
	if err = binary.Write(w, order, golangStringToClangString(t.MateID, 13)); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Color); err != nil {
		return
	}

	return gameSize, nil
}

type gameRecord struct {
	FiveWin        uint16
	FiveLose       uint16
	FiveTie        uint16
	ChcWin         uint16
	ChcLose        uint16
	ChcTie         uint16
	ChessEloRating uint16
	GoWin          uint16
	GoLose         uint16
	GoTie          uint16
}

func (t *gameRecord) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.FiveWin); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.FiveLose); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.FiveTie); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.ChcWin); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.ChcLose); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.ChcTie); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.ChessEloRating); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.GoWin); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.GoLose); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.GoTie); err != nil {
		return
	}

	return gameRecordSize, nil
}

func (t *gameRecord) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.FiveWin); err != nil {
		return
	}
	if err = binary.Write(w, order, t.FiveLose); err != nil {
		return
	}
	if err = binary.Write(w, order, t.FiveTie); err != nil {
		return
	}
	if err = binary.Write(w, order, t.ChcWin); err != nil {
		return
	}
	if err = binary.Write(w, order, t.ChcLose); err != nil {
		return
	}
	if err = binary.Write(w, order, t.ChcTie); err != nil {
		return
	}
	if err = binary.Write(w, order, t.ChessEloRating); err != nil {
		return
	}
	if err = binary.Write(w, order, t.GoWin); err != nil {
		return
	}
	if err = binary.Write(w, order, t.GoLose); err != nil {
		return
	}
	if err = binary.Write(w, order, t.GoTie); err != nil {
		return
	}

	return gameRecordSize, nil
}

type misc struct {
	WithMe uint32
	BrcID  uint32
}

func (t *misc) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.WithMe); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.BrcID); err != nil {
		return
	}
	return miscSize, nil
}

func (t *misc) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.WithMe); err != nil {
		return
	}
	if err = binary.Write(w, order, t.BrcID); err != nil {
		return
	}

	return miscSize, nil
}

type UserInfo struct {
	UID    int32
	PID    int32
	Socket int32
	userData
	friendShip
	message
	userStatus
	chatRoom
	chat
	game
	gameRecord
	misc
	WbTime int64
}

func (t *UserInfo) ReadFrom(r io.Reader) (_ int64, err error) {
	if err = binary.Read(r, order, &t.UID); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.PID); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.Socket); err != nil {
		return
	}
	if _, err = t.userData.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.friendShip.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.message.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.userStatus.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.chatRoom.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.chat.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.game.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.gameRecord.ReadFrom(r); err != nil {
		return
	}
	if _, err = t.misc.ReadFrom(r); err != nil {
		return
	}
	if err = binary.Read(r, order, &t.WbTime); err != nil {
		return
	}

	return UserInfoSize, nil
}

func (t *UserInfo) WriteTo(w io.Writer) (_ int64, err error) {
	if err = binary.Write(w, order, t.UID); err != nil {
		return
	}
	if err = binary.Write(w, order, t.PID); err != nil {
		return
	}
	if err = binary.Write(w, order, t.Socket); err != nil {
		return
	}
	if _, err = t.userData.WriteTo(w); err != nil {
		return
	}
	if _, err = t.friendShip.WriteTo(w); err != nil {
		return
	}
	if _, err = t.message.WriteTo(w); err != nil {
		return
	}
	if _, err = t.userStatus.WriteTo(w); err != nil {
		return
	}
	if _, err = t.chatRoom.WriteTo(w); err != nil {
		return
	}
	if _, err = t.chat.WriteTo(w); err != nil {
		return
	}
	if _, err = t.game.WriteTo(w); err != nil {
		return
	}
	if _, err = t.gameRecord.WriteTo(w); err != nil {
		return
	}
	if _, err = t.misc.WriteTo(w); err != nil {
		return
	}
	if err = binary.Write(w, order, t.WbTime); err != nil {
		return
	}

	return UserInfoSize, nil
}
