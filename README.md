<p align="center">
<h1 align="center">fulu-gosdk</h1>
<p align="center">福禄API sdk golang实现</p>

## Installation

```bash
go get github.com/t2krew/fulu-gosdk
```

## Usage

```go
import (
    "context"
	"time"
	"os"
	"log"
    fulu"github.com/t2krew/fulu-gosdk"
)

var cfg = Config{
    Debug:     true,
    Endpoint:  "https://openapi.fulu.com/api/getway",
    AppKey:    os.Getenv("FULU_APPKEY"),
    AppSecret: os.Getenv("FULU_APPSECRET"),
}

client, err := fulu.New(cfg)
if err != nil {
	panic(err)
}

// 查询商品信息

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

var productID string
productInfo, err := client.GetProductInfo(ctx, productID)
if err != nil {
	panic(err)
}

log.Printf("product: %+v", productInfo)

```
