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
