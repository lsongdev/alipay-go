package main

import (
	"log"

	"github.com/song940/alipay-go/alipay"
)

func main() {
	ali := alipay.NewClient(&alipay.Config{
		Gateway: "https://openapi-sandbox.dl.alipaydev.com/gateway.do",
		AppID:   "2019032563747119",
		AppKey: `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAjmogezo9O01wxLcTxFOOOZJpHsWYtEUURKerObFY5/rSVoBcPAkX5EIp+kaC5f8EAYfUnJmoyNnrk9w/0OGLRODdVELnbQXN5jpa6gBlzLgcTrV6hpM52SoZ0EkxgqA/IanIIYqkLqTk7xDhLnQ+0IZUbwBweSH8/dNQ/a0Yqyr2AAKKPnuUesOxWWkktoIPmX6/tmZwKoR7JmgOECixK+hjY/DWjJsfJpls/Xz7mxPOJus9E7L2BbWUAmud1IBau/mpXGs/L74lpryNpz2cxINgaOIkVyqbOPhtisL1SgaSBXfj0GrvxwVBvN99mpP1yew90Yj3jrnPLTF+NxRB8QIDAQABAoIBAHHlRFbpC+FqnJ5mgIOKA3vdsP5wLyE1AfGqWpEIWb8lZKTTgXBuIVJm6+WCENvaKJ0EwbAAX/FJ/1LAWbU3PEd4wunJnAYgFzxiPSVZ7hBcyh7UmsoX4rLgLcbbUqJSgEru2uSgWZAIRiM/z6d0RmxEkjA4HLtzyD5Di0ll7w2sWqXtXVkwVAot7h0/H6TKY5G3VlyuTqWqGpbUHMn1S5IVGgRMm7HCVR5YSjJZRFDupKQisfdUXtEveBFIk4IqSAyn9yPQPT520KrlypG7JBskFFXg5A39SYfPQmKxLfWrPLfOL6YtJQeJV2Oivoz1O9OxEYEBM+qYGgtEMCas9TECgYEAzyeTIkUHXY4BHx2eplDhxFyDnS/+RkzHXzp0jT36KOt09jDnkwEtYXImeHpnNTJutL7p+VG0WfobEdrJDthFvp0wVfIGOIqJhXzPCLq42CWRvcrfUtuLUDBxRLaywStkOjDzXrNctemRp3MUIvG0nHTauZPFltCAouEEegntsZ0CgYEAr/6qFuhCCsS4BQ1tnp54I3DHf3GNd9WKCaogjEY6ITxIA7nl4xLG5jx6sG6I3DTT5wzlkL5nKW2Q9VoU5VD2aNcoiAP6pp3q/VF5Px67AQJAEJKL5qRngxR7GCdWa/jh70Cle1aZLFPORSBAu3t23YpcgkhFTh2KQQFQM152O2UCgYB4AlZD4TeuJElUDGXPtkXE60+4LYiik3JhIc1J0iDtudKNmbFewazXqjjNTSQjdm8aOQv1SzcvdSxfgJ9AAV0OW6QX9llSQjf/ZFnQldPmLIWtLS2Jo/SmZRoJk8olDI1JBPjI4SIRpRmjp5B/2gUnKq9YGVq7z1jmg3ODe+L/JQKBgQCuaQ7Ags0oBMlk4GjY/6yJWrOpvat3rVwNtdZpjRMAas/nOWvzu2D3O8pOXEwvBf9VgvdhmP99E8LLEsmQc8quHUNif6b/RZJiFkK05cxm9IbupXwVRqn6Qeq3BgzkFZI52vPjpe9H+Yl6AbuE0Jb8d6izx9E+15FyWE3VinBa7QKBgH2iAzNJpg5kAezyXHtOxCHhgeSg5ZdYrpXm/HM01ytPBayZyh/xw7qDIHTDabNrYXzs29oEM4OO1LxqGqTCyxstWP40xExoMX7DIWt66UXKjNe7da+g8Cvo59zdPt1u/hscSj+qb2Z2OJ8DHW2dqxAfCN4GWDFAT7Zx81MvAb08
-----END RSA PRIVATE KEY-----`,
		AlipayPublicKey: `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqlnZyUOxHwCm1ZVMzTEgu3uf6+/9UsDSrZIqFqkrz4Gs3R2QJMPHykrzwIByHrEb0AiA6L2iy1aY6H28lGNXwsuAAS1iAPJyiB84uleRhRe+jg/NNbFAGM3TP6cqmpUXOtSbPD69d/B77UNbGDey/cFHpkQCzAunHqvirIq50PoGyQhpEMZjqEfGUQvZh1hlF5Tjnga8c0Ob7aZJBCptkP5F2N/k4CPYpxNWo0wTz0w/V5lcOtcREsect1TRmZiJ2tyKcEFeZmVXgwn8q5aLkl13wrZUYBYT91Hcuk7Kvq5KFtiQeA1Pd22h/HrwErqrpHFBwO1gjCEW8nIQU7bG1QIDAQAB
-----END PUBLIC KEY-----`,
	})
	res, err := ali.TradePreCreate(alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:     "test",
			OutTradeNo:  "20150320010101001",
			TotalAmount: "0.01",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res["qr_code"])
}
