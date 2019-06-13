# SMTP

這是用golang製作的基本httpServer的程式，歡迎參閱。
即可下載成功。

## 運行測試

可以直接在程式中```import ( "github.com/PietaTony/APILib/SMTP" )```來引入Lib。

需要的參數分別為

* "from":寄信端(Email string)
* "to": 收信端(Email string)
* "subj": 主題(string)
* "body": 內容(string)
* "SMTPServer": SMTP的伺服器
* "SMTPMail": SMTP帳號(Email string)
* "SMTPPassword": SMTP密碼(Email password string)

### GMail SMTP使用方法

#### 設定 IMAP

步驟 1：檢查 IMAP 是否已啟用

1. 在電腦上開啟 Gmail。

2. 按一下右上方的「設定」圖示 設定。

3. 按一下 [設定]。

4. 按一下 [轉寄和 POP/IMAP] 分頁標籤。

5. 在「IMAP 存取」部分中，選取 [啟用 IMAP]。

6. 按一下 [儲存變更]。

步驟 2：變更電子郵件用戶端中的 SMTP 和其他設定

請根據以下表格中的資訊來設定您的郵件程式。如需相關協助，請搜尋您的電子郵件程式的說明中心，查詢設定 IMAP 的操作說明。

| 資訊                       | 設定            | 
| -------------------------- | :-----------------------------------------------:  | 
| 內送郵件 (IMAP) 伺服器      | imap.gmail.com <br> 需要安全資料傳輸層 (SSL)：是 <br>  通訊埠：993                                         | 
| 外寄郵件 (SMTP) 伺服器      | smtp.gmail.com <br> 需要安全資料傳輸層 (SSL)：是 <br> 需要傳輸層安全性 (TLS)：是 (如果可用) <br> 需要驗證：是 <br> 安全資料傳輸層 (SSL) 通訊埠：465 <br> 傳輸層安全性 (TLS)/STARTTLS 通訊埠：587     | 
| 姓名或顯示名稱      | 您的姓名     | 
| 帳戶名稱、使用者名稱或電子郵件地址      | 您的完整電子郵件地址     | 
| 密碼      | 您的 Gmail 密碼     | 

# 資料來源
[Google SMTP相關設定](https://support.google.com/mail/answer/7126229?hl=zh-Hant)
