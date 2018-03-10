package checker

import (
	"github.com/sepuka/cryptoledger/structs"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const (
	apiurl = "https://api.etherscan.io"
	pathTemplate = "/api?module=account&action=balance&address=%v&tag=latest&apikey=%v"
	weiFactor = 1e18
)

var answer map[string]string

func Ethereum(wallets []structs.WatchEntity, secret string) {
	for _, wallet := range wallets {
		url := buildUrl(wallet, secret)
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Cannot get ethereum balance of %v wallet, error: %v", wallet.Address, err)
			continue
		}

		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)

		if decodeError := json.Unmarshal(body, &answer); decodeError != nil {
			log.Println("etherscan decode error ", decodeError)
		} else {
			log.Printf("Ethereum wallet %v contents %v", wallet.Address, answer)
			actualAMount := fetchAmount(answer)
			if wallet.MinAlertValue >= actualAMount {
				log.Printf("Balance of %v wallet too small (%v)!", wallet.Address, actualAMount)
			}
		}
	}
}

func buildUrl(entity structs.WatchEntity, secret string) string {
	path := fmt.Sprintf(pathTemplate, entity.Address, secret)
	return fmt.Sprint(apiurl, path)
}

func fetchAmount(answer map[string]string) int64 {
	amount, err := strconv.ParseInt(answer["result"], 10, 64)
	if err != nil {
		log.Println("Ethereum balance convert failure: ", err)
		return 0
	}
	return amount / weiFactor
}