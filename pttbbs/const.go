// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Note: variable in there will be unexported or removed soon.

package pttbbs

const (
	//////////
	//pttstruch.h
	//////////
	// IDLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L13
	IDLength = 12 /* Length of bid/uid */
	// IPV4Length https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L14
	IPV4Length = 15 /* a.b.c.d form */
	// PasswordInputLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L49
	PasswordInputLength = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	// PasswordLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L51
	PasswordLength = 14 /* Length of encrypted passwd field */
	// RegistrationLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L52
	RegistrationLength = 38 /* Length of registration data */
	// RealNameSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L62
	RealNameSize = 20 /* Size of real-name field */
	// NicknameSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L63
	NicknameSize = 24 /* SIze of nick-name field */
	// EmailSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L53
	EmailSize = 50 /* Size of email field */
	// AddressSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L79
	AddressSize = 50 /* Size of address field */
	// CareerSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L91
	CareerSize = 40 /* Size of career field */
	// PhoneSize https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L92
	PhoneSize = 20 /* Size of phone field */
	// PasswdVersion https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L55
	PasswdVersion = 4194 /* passwd version */
	// TitleLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L253
	TitleLength = 64 /* Length of title */
	// FileNameLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L254
	FileNameLength = 28 /* Length of filename */
)

const (
	//////////
	//pttstruch.h: 292
	//////////
	// FileLocal https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L292
	FileLocal = 0x01 /* local saved,  non-mail */
	// FileRead https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L293
	FileRead = 0x01 /* already read, mail only */
	// FileMarked https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L294
	FileMarked = 0x02 /* non-mail + mail */
	// FileDigest https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L295
	FileDigest = 0x04 /* digest,       non-mail */
	// FileReplied https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L296
	FileReplied = 0x04 /* replied,      mail only */
	// FileBottom https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L297
	FileBottom = 0x08 /* push_bottom,  non-mail */
	// FileMulti https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L298
	FileMulti = 0x08 /* multi send,   mail only */
	// FileSolved https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L299
	FileSolved = 0x10 /* problem solved, sysop/BM non-mail only */
	// FileHide https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L300
	FileHide = 0x20 /* hide,	in announce */
	// FileBoardID https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L301
	FileBoardID = 0x20 /* bid,		in non-announce */
	// FileBoardMaster https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L302
	FileBoardMaster = 0x40 /* BM only,	in announce */
	// FileVote https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L303
	FileVote = 0x40 /* for vote,	in non-announce */
	// FileAnonymous https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L304
	FileAnonymous = 0x80 /* anonymous file */
)

const (
	////////////
	// uflags.h
	// (https://github.com/ptt/pttbbs/blob/master/include/uflags.h)
	////////////
	UF_FAV_NOHILIGHT uint32 = 0x00000001 /* false if hilight favorite */
	UF_FAV_ADDNEW    uint32 = 0x00000002 /* true to add new board into one's fav */
	// UF_PAGER uint32 = 0x00000004 /* deprecated by cuser.pager: true if pager was OFF last session */
	// UF_CLOAK uint32 = 0x00000008 /* deprecated by cuser.invisible: true if cloak was ON last session */
	UF_FRIEND         uint32 = 0x00000010 /* true if show friends only */
	UF_BRDSORT        uint32 = 0x00000020 /* true if the boards sorted alphabetical */
	UF_ADBANNER       uint32 = 0x00000040 /* (was: MOVIE_FLAG, true if show advertisement banner */
	UF_ADBANNER_USONG uint32 = 0x00000080 /* true if show user songs in banner */
	// UF_MIND uint32 = 0x00000100 /* deprecated: true if mind search mode open <-Heat */
	UF_DBCS_AWARE       uint32 = 0x00000200 /* true if DBCS-aware enabled */
	UF_DBCS_NOINTRESC   uint32 = 0x00000400 /* no Escapes interupting DBCS characters */
	UF_DBCS_DROP_REPEAT uint32 = 0x00000800 /* detect and drop repeated input from evil clients */
	UF_NO_MODMARK       uint32 = 0x00001000 /* true if modified files are NOT marked */
	UF_COLORED_MODMARK  uint32 = 0x00002000 /* true if mod-mark is coloured */
	// UF_MODMARK_??? uint32 = 0x00004000 /* reserved */
	// UF_MODMARK_??? uint32 = 0x00008000 /* reserved */
	UF_DEFBACKUP       uint32 = 0x00010000 /* true if user defaults to backup */
	UF_NEW_ANGEL_PAGER uint32 = 0x00020000 /* true if user (angel) wants the new pager */
	UF_REJ_OUTTAMAIL   uint32 = 0x00040000 /* true if don't accept outside mails */
	UF_SECURE_LOGIN    uint32 = 0x00080000 /* true if login from insecure (ex, telnet) connection will be rejected */
	UF_FOREIGN         uint32 = 0x00100000 /* true if a foreign */
	UF_LIVERIGHT       uint32 = 0x00200000 /* true if get "liveright" already */
	// UF_COUNTRY_??? uint32 = 0x00400000 /* reserved */
	// UF_COUNTRY_??? uint32 = 0x00800000 /* reserved */
	UF_MENU_LIGHTBAR uint32 = 0x01000000 /* true to use lightbar-based menu */
	UF_CURSOR_ASCII  uint32 = 0x02000000 /* true to enable ASCII-safe cursor */
	// UF_??? uint32 = 0x04000000 /* reserved */
	// UF_??? uint32 = 0x08000000 /* reserved */
	// UF_??? uint32 = 0x10000000 /* reserved */
	// UF_??? uint32 = 0x20000000 /* reserved */
	// UF_??? uint32 = 0x40000000 /* reserved */
	// UF_??? uint32 = 0x80000000 /* reserved */
)
