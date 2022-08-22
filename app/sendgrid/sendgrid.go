package sendgrid

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
)

// requert.Bodyに格納するJSONの元となるメールを表す構造体
type Mail struct {
	Subject          string             `json:"subject"`
	Personalizations []Personalizations `json:"personalizations"`
	From             MailUser           `json:"from"`
	Content          []Contents         `json:"content"`
}

// 封筒のようなもの
// メールのメタデータを表す構造体
type Personalizations struct {
	To []MailUser `json:"to"`
}

// メールのユーザーを表す構造体
type MailUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// メールの中身を表す構造体
type Contents struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// .envファイルを読み込んで、ロードする
func Env_load() {
	// .envファイルを読み込んで、ロード
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

// メールの中身を作成して、メールを送信する
func SendMail(subject, contents, Email, Name string) {

	Env_load()

	// .envファイルに格納したAPI KEYを取得
	apiKey := os.Getenv("SENDGRID_API_KEY")
	// ホスト
	host := "https://api.sendgrid.com"
	// エンドポイント
	endpoint := "/v3/mail/send"

	// API KEYとエンドポイント、ホストからrestパッケージのRequestを生成
	request := sendgrid.GetRequest(apiKey, endpoint, host)
	// requestのMethodをPostに
	request.Method = "POST"

	// メールの内容をJSONで作成する
	mail := Mail{
		Subject: subject,
		Personalizations: []Personalizations{
			{To: []MailUser{{
				Email: Email,
				Name:  Name,
			},
			}}},
		From: MailUser{
			Email: os.Getenv("SENDER_ADDRESS"),
			Name:  os.Getenv("SENDER_NAME"),
		},
		Content: []Contents{{
			Type:  "text/plain",
			Value: contents,
		}},
	}

	// GoのコードをJSON化
	data, err := json.Marshal(mail)

	log.Println(string(data))

	if err != nil {
		log.Println(err)
	}

	// JSON化したmailの内容をrequest.Bodyに代入
	request.Body = data

	// sendGridのAPIにリクエストをセット
	// 戻り値でresponseが返ってくる
	_, err = sendgrid.API(request)
	//response, err := client.Send(request)
	if err != nil {
		log.Println(err)
	}
}
