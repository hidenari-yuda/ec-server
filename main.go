package main

import "github.com/hidenari-yuda/ec-server/app/controllers"

func main() {

	// controller
	//http.HandleFunc("/", echoHello)
	// port
	//http.ListenAndServe(":8000", nil)

	//ローカルサーバーを立ち上げる
	controllers.StartMainServer()
}

//func echoHello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "<h1>Hello World</h1>")

//}
