package bbs

type UserDraft interface {
	Raw() []byte
}

type userDraft struct {
	raw []byte
}

func NewUserDraft(raw []byte) UserDraft {
	return &userDraft{raw}
}

func (u userDraft) Raw() []byte {
	return u.raw
}