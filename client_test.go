package wxworksdk

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := &Config{
		CallbackToken:             "todo",
		CallbackEncodingAESKeyRaw: "todo",
		CorpID:                    "todo",
		AppSecret:                 "todo",
		AppID:                     1,
	}
	if err := cfg.Check(); err != nil {
		t.Fatal(err)
	}
	c := NewClient(cfg, slog.Default())

	decode, err2 := c.Decode([]byte("todo"))
	if err2 != nil {
		t.Fatal(err2)
	}
	fmt.Println(decode)
}
