package functions

import (
	"testing"
)

func TestConnectTwitterAPI(t *testing.T) {
	result := connectTwitterAPI()
	if result == nil {
		t.Fatal("\n認証テストに失敗")
	}

	t.Log("\n認証テストに成功")
}
