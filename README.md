# alipay-go

> 💰 Alipay SDK in Golang

## Example

```go
import (
  "github.com/song940/alipay-go/alipay"
)

func main() {
  ali := alipay.NewClient(&alipay.Config{
    AppID:   "2016073100133470",
    AppKey:  "",
    Gateway: "https://openapi.alipaydev.com/gateway.do",
  })
  res, err := ali.TradePreCreate(alipay.TradePreCreateParams{
    Subject:     "test",
    OutTradeNo:  "test",
    TotalAmount: "0.01",
  })
  log.Println(res, err)
}
```

## Node.js Implementation

> 💰 alipay sdk for node.js
>
> https://github.com/song940/node-alipay

## License

This project is licensed under the MIT License.