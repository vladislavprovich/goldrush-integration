package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//func main() {
//
//	url := "https://api.covalenthq.com/v1/eth-mainnet/address/0xd5B1D4e6626748968daee303133d2be5e2353f44/balances_v2/"
//	apiKey := "cqt_rQ6yfw3Y8fCMKqJKmKf494ffWqK3"
//
//	req, _ := http.NewRequest("GET", url, nil)
//
//	req.Header.Set("Content-Type", "application/json")
//	req.SetBasicAuth(apiKey, "")
//
//	res, _ := http.DefaultClient.Do(req)
//
//	defer res.Body.Close()
//	body, _ := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//
//}

func main() {

	url := "https://api.covalenthq.com/v1/allchains/transactions/"
	//url := "https://api.covalenthq.com/v1/allchains/address/0x742d35Cc6634C0532925a3b844Bc454e4438f44e/balances/"
	apiKey := "cqt_rQ6yfw3Y8fCMKqJKmKf494ffWqK3"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
