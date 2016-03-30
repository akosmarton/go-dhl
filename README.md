# Unofficial Go client for DHL API

## Install

`go get -u github.com/shipwallet/go-dhl/...`

## Usage

```
package main

import (
    "fmt"
    "github.com/shipwallet/dhl"
)

func main() {
    c := dhl.NewDHLClient("DServiceVal", "testServVal")

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
  			PieceType{PieceID: "1", Height: 30, Depth: 20, Width: 10, Weight: 1.0},
  		},
  	}
  	bdr.IsDutiable = "N"
  	bdr.NetworkTypeCode = "AL"
  	bdr.InsuredValue = 400.0
  	bdr.InsuredCurrency = "IDR"

  	du := &DCTDutiable{}
  	du.DeclaredCurrency = "EUR"
  	du.DeclaredValue = 9.0

  	res, err := client.GetQuote(from, to, bdr, du)
    if err != nil {
      fmt.Println(err)
    }

    fmt.Printf("%+v\n", res)
}
```

## Pull requests / Contributions

Contributions are very welcome. Please follow these guidelines:

- Fork the master branch and issue pull requests targeting the master branch
- If you are adding an enhancement, please open an issue first with your proposed change.


## Contributors

- Christoffer Åhrling - @nilpath
- Ilia Mikhailov - @iljoo
- Erik Johansson - @grillbiff

## License

The MIT License (MIT)

Copyright (c) 2016 Shipwallet AB

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
