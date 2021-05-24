// 實作幾個 method 傳入 userid 回傳該使用者的 BBS 設定
// 目前已經有可以取得使用者userec的method可用了
//
// 請見 userec_t 中的 uflag 和 userlevel
// 每個選項對應的 flag id 請參見 user.c desc1 和 masks1
// user.c - https://github.com/ptt/pttbbs/blob/5715b35f510f48eb5092d32882f1aa09181dc3a1/mbbsd/user.c#L438
// uflags.h - https://github.com/ptt/pttbbs/blob/4d56e77f264960e43e060b77e442e166e5706417/include/uflags.h

package pttbbs

import (
	"fmt"

	"github.com/Ptt-official-app/go-bbs"
)

var userRecs []bbs.UserRecord

// (Question) I'm not sure how to access this function from serverlet/route_token.go
func findUserecByID(userid string) (bbs.UserRecord, error) {
	for _, it := range userRecs {
		if userid == it.UserID() {
			return it, nil
		}
	}
	return nil, fmt.Errorf("user record not found")
}

// get ALL UserFlag at once:
type uFlags struct {
	dbanner          bool
	dbanner_usong    bool
	rej_outtamail    bool
	defbackup        bool
	secure_login     bool
	fav_addnew       bool
	fav_nohilight    bool
	no_modmark       bool
	colored_modmark  bool
	dbcs_aware       bool
	dbcs_drop_repeat bool
	dbcs_nointresc   bool
	cursor_ascii     bool
	menu_lightbar    bool
	new_angel_pager  bool
}

func getUserFlagAllByID(userid string) (*uFlags, error) {
	userrec, err := findUserecByID(userid)
	if err != nil {
		return nil, fmt.Errorf("user record not found")
	}
	dbanner := getAdbannerByUserrec(userrec)
	dbannerUsong := getAdbannerUsongByUserrec(userrec)
	rejOuttamail := getRejOuttamailByUserrec(userrec)
	defBackup := getDefBackupByUserrec(userrec)
	secureLogin := getSecureLoginByUserrec(userrec)
	favAddnew := getFavAddNewByUserrec(userrec)
	favNohilight := getFavNoHilightByUserrec(userrec)
	noModmark := getFavNoModMarkByUserrec(userrec)
	coloredModmark := getColoredModMarkByUserrec(userrec)
	// TO DO...#ifdef DBCSAWARE
	dbcsAware := getDbcsAwareByUserrec(userrec)
	dbcsDropRepeat := getDbcsDropRepeatByUserrec(userrec)
	dbcsNointresc := getDbcsNointrescByUserrec(userrec)
	// #endif
	cursorASCII := getCursorASCIIByUserrec(userrec)
	// TO DO...#ifdef USE_PFTERM
	menuLightbar := getMenuLightBarByUserrec(userrec)
	// #endif
	// TO DO...#ifdef PLAY_ANGEL
	newAngelPager := getNewAngelPagerByUserrec(userrec)
	// #endif

	return &uFlags{
		dbanner,
		dbannerUsong,
		rejOuttamail,
		defBackup,
		secureLogin,
		favAddnew,
		favNohilight,
		noModmark,
		coloredModmark,
		dbcsAware,
		dbcsDropRepeat,
		dbcsNointresc,
		cursorASCII,
		menuLightbar,
		newAngelPager}, nil
}

func getAdbannerByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfAdbanner != 0
}

func getAdbannerUsongByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfAdbannerUsong != 0
}

func getRejOuttamailByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfRejOuttamail != 0
}

func getDefBackupByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfDefbackup != 0
}

func getSecureLoginByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfSecureLogin != 0
}

func getFavAddNewByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfFavAddnew != 0
}

func getFavNoHilightByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfFavNohilight != 0
}

func getFavNoModMarkByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfNoModmark != 0
}

func getColoredModMarkByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfColoredModmark != 0
}

// #ifdef DBCSAWARE (Not handling yet...)
func getDbcsAwareByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfDbcsAware != 0
}

// #ifdef DBCSAWARE (Not handling yet...)
func getDbcsDropRepeatByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfDbscDropRepeat != 0
}

// #ifdef DBCSAWARE (Not handling yet...)
func getDbcsNointrescByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfDbcsNointresc != 0
}

func getCursorASCIIByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfCursorASCII != 0
}

// #ifdef USE_PFTERM (Not handling yet...)
func getMenuLightBarByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfMenuLightbar != 0
}

// #ifdef PLAY_ANGEL (Not handling yet...)
func getNewAngelPagerByUserrec(userrec bbs.UserRecord) bool {
	return userrec.UserFlag()&UfNewAngelPager != 0
}
