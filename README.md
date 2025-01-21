# wxworksdk

企业微信SDK


## 初始化客户端

```go
package main

import "gitlab.minum.cloud/BackendTeam/pkg/wxmpsdk"

func foo() {
	wxmpsdk.NewClient( /*&wxmpsdk.Config{
			CallbackToken:             "abceyr359Ml",
			CallbackEncodingAESKeyRaw: "UMGKXiH811AipU7Pf1JNUOB24tSqcqwzzVAzJvHJK8z",
			CorpID:                    "ww21a88888bd0a8a888",
			AppSecret:                 "lv6T-2-pFV5FTwevI_P09wxdn9PC1WCfxabcdegh1",
			AppID:                     1000001,
		}*/nil, logger.Logger)
}
```

## 接收企业微信回调

```go
package main

func foo() {
	// get 回调，验证可用性
	srv.Route("/hook").GET("", func(context http.Context) error {
		msg_signature := context.Query().Get("msg_signature")
		timestamp := context.Query().Get("timestamp")
		nonce := context.Query().Get("nonce")
		echostr := context.Query().Get("echostr")
		logger.Info("hook get", zap.String("msg_signature", msg_signature), zap.String("timestamp", timestamp), zap.String("nonce", nonce), zap.String("echostr", echostr))
		content, err := wxClient.VerifyCallback(echostr, msg_signature, nonce, timestamp)
		if err != nil {
			return err
		}
		logger.Info("hook get", zap.String("content", string(content)))
		return context.String(200, string(content))
	})

	// post 回调，处理事件
	srv.Route("/hook").POST("", func(ctx http.Context) error {
		body := ctx.Request().Body
		echostr, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		content, err := wxClient.Decode(echostr)
		if err != nil {
			return err
		}
		logger.Info("hook post", zap.String("content", string(content)))
		return ctx.String(200, "success")
	})

}
```