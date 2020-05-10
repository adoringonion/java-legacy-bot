package functions

import (
	"testing"
)

func TestConnectTwitterAPI(t *testing.T) {
	api := ConnectTwitterAPI()

	if api == nil {
		t.Fatal("\n認証テストに失敗")
	}

	t.Log("\n認証テストに成功")
}

func TestGetTweetsCount(t *testing.T) {
	result, err := GetTweetsCount("テスト")
	if err != nil {
		t.Fatal("\nツイート取得失敗")
	}

	t.Log("\nツイート取得成功", result)
}
