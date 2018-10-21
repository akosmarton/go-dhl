# Unofficial Go client for DHL API

## Install

`go get -u github.com/akosmarton/go-dhl/...`

## Usage

```go
package main

import (
    "fmt"
    "github.com/akosmarton/go-dhl"
)

func main() {
	config := ClientConfig{}
	client, _ := NewDHLExpressClient("DServiceVal", "testServVal", config)

	from := &DCTFrom{}
	from.CountryCode = "ID"
	from.PostalCode = "31251"

	to := &DCTTo{}
	to.CountryCode = "JP"
	to.PostalCode = "9811513"

	t := time.Now()
	bdr := &BkgDetailsRequest{}
	bdr.PaymentCountryCode = "ID"
	bdr.Date = t.Format("2006-01-02")
	bdr.ReadyTime = t.Format("PT15H04M")
	bdr.ReadyTimeGMTOffset = "+01:00"
	bdr.DimensionUnit = "CM"
	bdr.WeightUnit = "KG"
	bdr.Pieces = &Pieces{
		Piece: []PieceType{
			{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
		},
	}
	bdr.IsDutiable = "N"
	bdr.NetworkTypeCode = "AL"
	bdr.InsuredValue = 400.0
	bdr.InsuredCurrency = "IDR"

	du := &DCTDutiable{}
	du.DeclaredCurrency = "EUR"
	du.DeclaredValue = 9.0

	resp, err := client.GetQuote(from, to, bdr, du)
    if err != nil {
      fmt.Println(err)
    }

    fmt.Printf("%+v\n", resp)
}
```