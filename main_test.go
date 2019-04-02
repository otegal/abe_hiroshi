package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// ローカルに保存したHTMLをMockとしてテストする
	// テスト用に値を変更
	pageURL = "file:///Users/h_hiroki/go/src/abe_hiroshi/MockHtml/阿部寛のホームページ.html"
	screenShotFileName = "test_screen_shot.jpg"

	// 実際の処理を実行
	main()
}
