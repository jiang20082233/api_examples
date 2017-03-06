package main

//依赖包引入
import (
	"github.com/nntaoli/crypto_coin_api/huobi"
	"log"
	"net/http"
	//"github.com/nntaoli/crypto_coin_api/okcoin"
	. "github.com/nntaoli/crypto_coin_api"
	//"github.com/nntaoli/crypto_coin_api/builder"
)

func main() {
	//accessKey,secretKey 需要去交易所网站申请，用于调用API接口，授权使用
	accessKey := ""
	secretKey := ""

	//New Huobi API
	api := huobi.New(http.DefaultClient, accessKey, secretKey)

	//_api := builder.NewAPIBuilder().APIKey(accessKey).APISecretkey(secretKey).Build("huobi.com") //用Builder构建对应交易所的API接口实例

	//api := okcoin.New(http.DefaultClient , accessKey , secretKey) //New OKCoin.cn API
	//api := okcoin.NewFuture(http.DefaultClient , accessKey , secretKey) //New OKCoin Future API

	//调用方法，GetAccount 获取账户资产信息
	account, err := api.GetAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("总资产:%.4f,净资产:%.4f", account.Asset, account.NetAsset)
	log.Printf("[BTC资产]可用:%.4f,冻结:%.4f,借贷:%.4f", account.SubAccounts[BTC].Amount,
		account.SubAccounts[BTC].ForzenAmount, account.SubAccounts[BTC].LoanAmount)
	log.Printf("[LTC资产]可用:%.4f,冻结:%.4f,借贷:%.4f", account.SubAccounts[LTC].Amount,
		account.SubAccounts[LTC].ForzenAmount, account.SubAccounts[LTC].LoanAmount)
	log.Printf("CNY资产:可用:%.4f,冻结:%.4f,借贷:%.4f", account.SubAccounts[CNY].Amount,
		account.SubAccounts[CNY].ForzenAmount, account.SubAccounts[CNY].LoanAmount)
}
