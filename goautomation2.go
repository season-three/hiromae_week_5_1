package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

func main() {

	//毎回ランダムな数を出す -> 待機時間に使う
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)

	//ドライバの起動＋クロームの起動＋スタート
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	err := driver.Start()
	if err != nil {
		fmt.Println("Failed to start driver")
	}
	//defer driver.Stop()

	//新しいページを開く
	page, err := driver.NewPage()
	if err != nil {
		fmt.Println("Failed to open new page")
	}

	//ページ開く
	//err = page.Navigate(os.Getenv("url1"))
	err = page.Navigate("https://www.linkedin.com/search/results/people/?facetGeoRegion=%5B%22jp%3A0%22%5D&keywords=engineer&origin=GLOBAL_SEARCH_HEADER")

	if err != nil {
		fmt.Println("Failed to open login page")

	}
	time.Sleep(time.Duration(n) * time.Second)

	//アカウント登録の画面 -> ログインページに
	page.FindByClass("main__sign-in-link").Click()
	time.Sleep(time.Duration(n) * time.Second)

	//usernameとpasswordのIDを取得
	username := page.FindByID("username")
	password := page.FindByID("password")

	//usernameとpasswordを入力
	username.Click()
	time.Sleep(time.Duration(n) * time.Second)

	username.Fill(os.Getenv("username"))
	time.Sleep(time.Duration(n) * time.Second)

	username.Click()
	time.Sleep(time.Duration(n) * time.Second)

	password.Fill(os.Getenv("password"))
	time.Sleep(time.Duration(n) * time.Second)

	//サインインボタンを押す
	err = page.FindByButton("サインイン").Submit()
	if err != nil {
		fmt.Println("Failed to login")
	}
	time.Sleep(5 * time.Second)

	//次へというボタンを押すコードを書きたい
	//FindbyClass
	//page.FirstByClass("artdeco-pagination__button artdeco-pagination__button--next artdeco-button artdeco-button--muted artdeco-button--icon-right artdeco-button--1 artdeco-button--tertiary ember-view")
	//page.FirstByClass("artdeco-button__icon")
	//page.AllByButton("次へ").Click()
	//page.FindByButton("次へ").Click()
	//page.AllByID("ember551").Click()
	//page.FindByID("ember551").Click()
	//page.AllByClass("artdeco-button__icon").Click()
	//page.FindByClass("artdeco-button__icon").Click()
	//page.AllByName("次へ")

	next := page.FindByClass("artdeco-pagination__button artdeco-pagination__button--next artdeco-button artdeco-button--muted artdeco-button--icon-right artdeco-button--1 artdeco-button--tertiary ember-view")
	next.MouseToElement()
	next.Click()
	time.Sleep(time.Duration(n) * time.Second)

	next2 := page.FindByID("ember387")
	next2.MouseToElement()
	next2.Click()
	time.Sleep(time.Duration(n) * time.Second)
	/*
		next3 := page.FirstByID("ember387")
		next3.MouseToElement()
		next3.Click()
		time.Sleep(time.Duration(n) * time.Second)
	*/

	/*

		//申請の上限までつながり申請をする
		count := 0

		for i := 0; i < 5; i++ {
			err = page.FirstByButton("つながりを申請").Click()
			time.Sleep(time.Duration(n) * time.Second)

			page.FindByButton("完了").Click()
			time.Sleep(time.Duration(n) * time.Second)

			if err == nil {
				count++
			}

			if err != nil {

				//ここでスクロールダウンしたいが、方法が分からない
				//もしくは、そもそもスクロールダウンしなくても反映されるのか？？

				time.Sleep(time.Duration(n) * time.Second)

				//	} else {
				//		page.FirstByButton("次へ").Click()
				//		time.Sleep(time.Duration(n) * time.Second)

			}
		}
		fmt.Printf("Sent %d requests\n", count)

	*/

	/*
		//別の単語で検索 or 新しいURLで検索
		url12 := "https://www.linkedin.com/search/results/people/?facetGeoRegion=%5B%22jp%3A0%22%5D&keywords=&origin=GLOBAL_SEARCH_HEADER"

		err = page.Navigate(url12)
		if err != nil {
			fmt.Println("Failed to open search page")
		}

		searchwords := page.FindByID("ember44")
		searchwords.Click()
		time.Sleep(time.Duration(n) * time.Second)

		searchwords.Fill("searchword")
		time.Sleep(time.Duration(n) * time.Second)

		//Clickは反応するのにClearとFillが反応しないのはなぜ、、？？
		//serchworespop := page.FindByClass("search-global-typeahead search-global-typeahead--focused ember-view")
		//serchworespop := page.FirstByButton("検索")
		//serchworespop.Click()
		//time.Sleep(time.Duration(n) * time.Second)

		//searchwords.Fill(os.Getenv("searchword"))
		//serchworespop.Fill(os.Getenv("searchword"))
		//time.Sleep(time.Duration(n) * time.Second)

		page.FindByClass("search-global-typeahead__button").Click()
		time.Sleep(time.Duration(n) * time.Second)


	*/

	/*
		//プロフィール見てくれた人につながり申請(スクロールの方法が分かったら)
		url9 := https://www.linkedin.com/me/profile-views/urn:li:wvmp:summary/

		err = page.Navigate(url9)
		if err != nil {
		fmt.Println("Failed to open viewer page")
		}
	*/

	/*

		//いいねを1日5個の投稿に押す
		url10 := "https://www.linkedin.com/feed/"

		err = page.Navigate(url10)
		if err != nil {
			fmt.Println("Failed to open feed page")
		}

		for i := 0; i < 4; i++ {
			err = page.FindByClass("artdeco-button__text react-button__text").Click()
			time.Sleep(time.Duration(n) * time.Second)

			if err != nil {
				fmt.Println("Failed to like posts")
			}
			//スクロール
		}
	*/

	/*
		//新しいタブを開くにはどうすれば良いか？？
		page2, err := driver.NextWindow()
		if err != nil {
			fmt.Println("Failed to open new page")
		}
	*/

	//繋がり申請があれば承認押する
	url11 := "https://www.linkedin.com/mynetwork/"

	err = page.Navigate(url11)
	if err != nil {
		fmt.Println("Failed to open accept page")
	}
	time.Sleep(5 * time.Second)
	page.FindByClass("msg-overlay-bubble-header").Click()
	time.Sleep(time.Duration(n) * time.Second)

	count2 := 0
	for {
		err = page.FirstByButton("承認").Click()
		if err == nil {
			count2++
		}

		time.Sleep(time.Duration(n) * time.Second)

		if err != nil {
			break
		}

	}
	fmt.Printf("Accepted %d requests\n", count2)

}
