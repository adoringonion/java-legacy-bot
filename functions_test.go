package functions

import (
	"testing"
)

func TestConnectTwitterAPI(t *testing.T) {
}

func TestGetTweetsCount(t *testing.T) {
	result, err := GetTweetsCount("テスト")
	if err != nil {
		t.Fatal("\nツイート取得失敗")
	}

	t.Log("\nツイート取得成功", result)
}