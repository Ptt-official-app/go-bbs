# BBS Parser


這個專案是提供 Golang 開發者存取現有在台灣的 BBS 資料結構的函式庫。

目前主要支援的 BBS 結構以 CurrentPTT 為主

未來可能會支援 FormosaBBS


## 目前支援的檔案

- [ ] 看板文章目錄檔
- [ ] 使用者檔案

## 測試方式

```
go test
```

這樣。


## 參考文件

系統檔案列表
```
   1 歡迎畫面                             etc/Welcome
    2 出站畫面                             etc/Logout
    3 錯誤登入訊息                         etc/goodbye
    4 站長名單                             etc/sysop
    5 帳號站長信箱收信列表                 etc/mail_account_sysop
    6 帳號站長信箱收信說明                 etc/mail_account_sysop_desc
    7 拒絕連線IP列表 (BANIP)               etc/banip.conf
    8 進站畫面0                            etc/Welcome_login.0
    9 進站畫面1                            etc/Welcome_login.1
   10 進站畫面2                            etc/Welcome_login.2
   11 進站畫面3                            etc/Welcome_login.3
   12 進站畫面4                            etc/Welcome_login.4
   13 過度轉錄開的罰單通知信               etc/crosspost.txt
   14 我的最愛預設列表                     etc/myfav_defaults
   15 發文注意事項                         etc/post.note
   16 看板期限                             etc/expire2.conf
   17 節日                                 etc/feast
   18 故鄉                                 etc/domain_name_query.cidr
   19 註冊 email 白名單                    etc/whitemail
   20 註冊 email 未在白名單的通知訊息      etc/whitemail.notice
   21 註冊 email 黑名單                    etc/banemail
   22 註冊範例                             etc/register
   23 註冊通過通知                         etc/registered
   24 新使用者需知                         etc/newuser
   25 退註通知附加說明                     etc/reg_reject.notes
   26 註冊單填寫說明                       etc/regnotes/front
   27 註冊細項說明[是否現住台灣]           etc/regnotes/foreign
   28 註冊細項說明[姓名]                   etc/regnotes/name
   29 註冊細項說明[職業]                   etc/regnotes/career
   30 註冊細項說明[住址]                   etc/regnotes/address
   31 註冊細項說明[電話]                   etc/regnotes/phone
   32 註冊細項說明[手機]                   etc/regnotes/mobile
   33 註冊細項說明[生日]                   etc/regnotes/birthday
   34 註冊細項說明[性別]                   etc/regnotes/sex
  35 看板列表說明                         etc/boardlist.help
   36 文章列表說明                         etc/board.help
   37 小天使認證通知                       etc/angel_notify
   38 小天使功能說明                       etc/angel_usage
   39 小天使功能說明(有留言)               etc/angel_usage2
   40 小天使離線訊息(有留言)               etc/angel_offline2
   41 外籍使用者認證通知                   etc/foreign_welcome
   42 外籍使用者過期警告通知               etc/foreign_expired_warn
  ```
  
  建立新看板的設定值
```
  【 建立新板 】



A. (無作用)             Ｘ              Q. 不可噓               Ｘ
B. 不列入統計           Ｘ              R. (無作用)             Ｘ
C. (無作用)             Ｘ              S. 限看板會員發文       Ｘ
D. 群組板               Ｘ              T. Guest可以發表        Ｘ
E. 隱藏板               Ｘ              U. 冷靜                 Ｘ
F. 限制(不需設定)       Ｘ              V. 自動留轉錄記錄       ˇ
G. 匿名板               Ｘ              W. 禁止快速推文         Ｘ
H. 預設匿名板           Ｘ              X. 推文記錄 IP          Ｘ
I. 發文無獎勵           Ｘ              Y. 十八禁               Ｘ
J. 連署專用看板         Ｘ              Z. 對齊式推文           Ｘ
K. 已警告要廢除         Ｘ              0. 不可自刪             Ｘ
L. 熱門看板群組         Ｘ              1. 板主可刪特定文字     Ｘ
M. 不可推薦             Ｘ              2. 沒想到               Ｘ
N. 小天使可匿名         Ｘ              3. 沒想到               Ｘ
O. 板主設定列入記錄     Ｘ              4. 沒想到               Ｘ
P. 連結看板             Ｘ              5. 沒想到               Ｘ

```

  發表權限
```
  設定 [test] 看板之(發表)權限：

A. 基本權力             Ｘ              Q. 不列入排行榜         Ｘ
B. 進入聊天室           Ｘ              R. 違法通緝中           Ｘ
C. 找人聊天             Ｘ              S. 小天使(本站無效)     Ｘ
D. 發表文章             Ｘ              T. 不允許認證碼註冊     Ｘ
E. 註冊程序認證         Ｘ              U. 視覺站長             Ｘ
F. 信件無上限           Ｘ              V. 觀察使用者行蹤       Ｘ
G. 隱身術               Ｘ              W. 禠奪公權             Ｘ
H. 看見忍者             Ｘ              X. 群組長               Ｘ
I. 永久保留帳號         Ｘ              Y. 帳號審核組           Ｘ
J. 站長隱身術           Ｘ              Z. 程式組               Ｘ
K. 板主                 Ｘ              0. 活動組               Ｘ
L. 帳號總管             Ｘ              1. 美工組               Ｘ
M. 聊天室總管           Ｘ              2. 警察總管             Ｘ
N. 看板總管             Ｘ              3. 小組長               Ｘ
O. 站長                 Ｘ              4. 退休站長             Ｘ
P. BBSADM               Ｘ              5. 警察                 Ｘ
```

## 專有名詞部分

有些名詞因為太常出現，可能會考慮直接在程式碼裡面以縮寫表示而不寫出全名：

* BM: Board Moderator 版主的意思


## PosOf的命名慣例以及解析方式

雖然 `binary.Read` 確實是方便好用，但實際上各站的 BBS 維護者可能不小心在該站的檔案結構中因為增加欄位長度沒考慮到 Padding 問題，
舉例來說 Ptt 的程式碼大多數沒有另外宣告 `__attribute__((__packed__))` 因此在不同架構的環境下編譯出的 struct 長度可能會不同。

因此為了方便其他人維護，專案目前採用PosOf手動算出欄位記憶體相對位置的方式，效率上可能較低，但是有可能被編譯器最佳化因此未來也不確定低多少。
未來如果需要加速同時很確定貴站的資料結構時，推薦在testing以及benchmark都確定過之後切換成以 `binary.Read` 實作的版本。

目前因為為了要支援多套BBS版本的緣故，因此命名結構會是以下形式：

`PosOfPttFileHeaderFilename`

* `PosOf` 開頭
* `Ptt` 站點本身的 Code Name
* `FileHeader` 哪個結構體名稱
* `Filename` 欄位名稱

預設省略版本號，如果需要新增版本號的範例如下


`PosOfPttR5939FileHeaderFilename`

* `R5939` 就是版本號

## 授權

Apache 2.0 (TBD)
