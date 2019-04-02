package main

import (
	"log"

	"github.com/sclevine/agouti"
)

var pageURL = "http://abehiroshi.la.coocan.jp/"
var screenShotFileName = "abe_hiroshi.jpg"

var setPageURL = func(s string) string { return s }
var setFileName = func(s string) string { return s }

func main() {
	driver := agouti.ChromeDriver()
	// driver := agouti.ChromeDriver(
	// 	// headlessの場合
	// 	agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
	// )

	if err := driver.Start(); err != nil {
		log.Fatalf("driverの起動に失敗しました : %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("セッション作成に失敗しました : %v", err)
	}

	// 阿部寛のウェブページに遷移する
	if err := page.Navigate(setPageURL(pageURL)); err != nil {
		log.Fatalf("阿部寛になにかあったかもしれません : %v", err)
	}

	// 写真集のリンクを検索する
	// framesetの中の要素を検索するには一旦該当のフレームにフォーカスしなければならない
	// 左側のフレームにフォーカスする
	if err := page.FindByXPath("/html/frameset/frame[1]").SwitchToFrame(); err != nil {
		log.Fatalf("阿部寛の左側frameにフォーカスできませんでした : %v", err)
	}
	// 「写真集」をクリック
	if err := page.FindByXPath("html/body/table/tbody/tr[10]/td[3]/p/a").Click(); err != nil {
		log.Fatalf("阿部寛の写真集が見つかりませんでした : %v", err)
	}

	// フレームのフォーカス外すためrootにもどる
	if err := page.SwitchToRootFrame(); err != nil {
		log.Fatalf("しょうも無いエラーが発生しました : %v", err)
	}

	// 右側のフレームにフォーカスする
	if err := page.FindByXPath("/html/frameset/frame[2]").SwitchToFrame(); err != nil {
		log.Fatalf("阿部寛の右側frameにフォーカスできませんでした : %v", err)
	}

	// 「メンズノンノ阿部寛」をクリック
	if err := page.FindByXPath("/html/body/table/tbody/tr[8]/td[2]/strong/a").Click(); err != nil {
		log.Fatalf("阿部寛のメンズノンノがみつかりませんでした : %v", err)
	}

	// スクショとる
	if err := page.Screenshot(setFileName(screenShotFileName)); err != nil {
		log.Fatalf("スクショ取れまへん : %v", err)
	}
}
