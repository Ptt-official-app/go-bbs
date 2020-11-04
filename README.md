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

## 授權

Apache 2.0 (TBD)
