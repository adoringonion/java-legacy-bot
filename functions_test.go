package functions

import (
	"testing"
)

func TestGetTweetsCount(t *testing.T) {
	result, respCode, err := GetTweetsCount("テスト")
	t.Log("レスポンスコード:", respCode)
	if err != nil {
		t.Fatal("\nツイート取得失敗")
	}

	t.Log("\nツイート取得成功", result)
}
