package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	name := query.Get("name") // "名前"から"name"へ修正

	// レスポンス用のマップを作成
	response := map[string]string{
		"message": "Hello " + name, // “message”： "Hello " + name、から修正
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {

	var categories = []string{"categories1", "categories2", "categories3"}
	// レスポンス用のマップを作成
	response := map[string]interface{}{
		"categories": categories,
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータを解析する
	query := r.URL.Query()
	operator := query.Get("o")

	//URL上"+"が空白として認識されるため、コード上は空白をプラスに変換
	if operator == " " {
		operator = "+"
	}
	//x,yを定義する
	x, err := strconv.ParseFloat(query.Get("x"), 64)
	if err != nil {
		http.Error(w, "x is number or float", http.StatusBadGateway)
	}
	y, err := strconv.ParseFloat(query.Get("y"), 64)
	if err != nil {
		http.Error(w, "y is number or float", http.StatusBadGateway)
	}
	//演算方法の決定
	var result float64
	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			http.Error(w, "Division by zero is not allowed.", http.StatusBadRequest)
			return
		}
		result = x / y
	default:
		http.Error(w, "Uncensoured operator.", http.StatusBadRequest)
	}

	// レスポンス用のマップを作成
	response := map[string]float64{
		"result": result,
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/categories", categoriesHandler)
	http.HandleFunc("/api/calculator", calculatorHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
