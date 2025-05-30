package myWallet

import (
	"encoding/json"
	"go_example/util"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	priKeyG         = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDZUFJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0FtSXdnZ0plQWdFQUFvR0JBTVdTTmtKZUJ2WjRESGJkClZMWENqTS9mM0RZbzdNYlRzVU9xMm1LUzhXVnE0ZVVXeGJoNE5lUXZnNGdqdUszM3hDWkZSTEwrTXVjbVZodDcKRnl2OUxRK1VYTkRvS1NMUHJBK0dPbXJZMnduWWczMzZWSVc2NEwxLzRmcitsY3hzMDVmckF5cFpjT3RrMlFxdwpZMzY4aVQzbjQxZDRRVzJ0dW5zYmVSOURVTjFEQWdNQkFBRUNnWUVBcmNPeXhSdzlzM2hTUGhqYjhDQjBDRUF6ClJjOG9zSlp2U2J4eTVrK20wRFA1Q2F2RnFrRFA1U21FM2EvTk5mUzhKNHkwcDFpN0hHR3pTZWd3c1JlekVIemQKUUZ0MkZqdTdUNk1kN3FWdHg4SkRRS3RZcGgrRkgwaWtCa1RwblRhL1NVRFEvWnJyekRsZW9qVHpENXpCLzVERgpHc1NpTUZkSnZCWVhhK3NkQjRFQ1FRRFlSQnFjZXQ3MThUb0FwUnMzWjU0Q2VtT29rOUlmUGtMeWJEUmFCQUhOCjZYTHZJTjVuNThPOVQ3eVNkVnJUVllKUURkem1XdENLWEtaaDFkdTd0MzJiQWtFQTZkN1BJT2lRVGcwU3RvVWkKL1N3RVEydjQ1Qms4N3g0cXdwbDlSSFYvM0xmblgzZVJvMEZIbS9uaStiTnNTY1NTcjB5UUJkb0RpQlpkQmJxMAppc3J0ZVFKQVR4bEpxbVgrV25IcVJ3WWNXRDFidTRoTUh3SkIzZytGcU9rT2xNWFdheHV4WCtqanI1bERMR0NYCmxmRzZVSVY3N2cvRDliVE5oVzJ4cFNMMUJVbHRkUUpCQUp6SWVQVW4wNjN2aDY4Q25BMDdsL1FYQ3pqblRubEQKTVFsYVdxY3RXalYrdUViQTRzcWVOY0owQ0Z5N2t3bjFGUjBkNTZOMG4wOTVKbzF6dUJzZnBZa0NRUURVbFpMawpPUEI1cTdVNzdiQ3RTaUIxaWh6OVhSQ1lKYWdDWEJzK2xyYUNSWkhGWFdBdnAzWnQ2UFYwUVA4cUFGN09vUE5CCnAvOFVPeXp4ZDBsaTFNUnQKLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"
	merchantCodeG   = "R1Hh7"
	platformPubKeyG = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERXRLWWdBWDVXS2h6R2ljTVNSdVl2VlB2QgpNeGpvN1JVbzROWk9CU2ZIRkt6ZnppUVFKcUtRK3dUd2p0UmdhTndyQzd5bEtGcGNxamlqTTBMU2VscHpoWmwrCnFxTmVacE4yaDJkMW5wQ0wzbVBPYjJZYjJyUHc4T01oZmZ0WTc1dWxXYVY1cXBWTW1WSitSZ3VKOVlDN2tIaW8KQU9tUVVkUWhqOStPTmdCZFV3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	priBaseUrlG     = "https://walapi.wearelucky2025.top/a001/api"
)

func TestGetTokenList(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}
	tokenList, err := GetTokenList()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(tokenList)
}

func TestPostCreateAddress(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}
	netWork := "Polygon"
	netWork = "TRON"
	netWork = "BSC"
	netWork = "Ethereum"
	netWork = "Nano"
	netWork = "Ravencoin"
	netWork = "Dogecoin"
	netWork = "Litecoin"
	//netWork = "Solana"

	address, err := PostCreateAddress(netWork, "http://92.118.228.90:19900")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(address)
}

func TestPostChangeAddress(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}
	address, err := PostChangeAddress("0x18ef4bdc0472d55460ad86ed4c04304de0bbe576", "http://92.118.228.90:19900")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(address)
}

func TestGetTradeConfirm(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	//res, err := GetTradeConfirm("aUVN72025052589133128")
	//res, err := GetTradeConfirm("aUVN7202505258913312a")

	//res, err := GetTradeConfirm("CAbA72025052584301669")
	//res, err := GetTradeConfirm("dFpy72025052541932302")
	code, res, err := GetTradeConfirm("nioJ72025052508154677")

	t.Log(code)
	if nil != err {
		t.Fatal(err)
	}

	buf, _ := json.Marshal(res)
	t.Log(string(buf))
}

func TestGetTrade(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	//trade, err := GetTrade("Polygon", "0xdf5d751d90b79c1a5f9df08b5d7165b225f4dcc46093c8985c369ce3a041b9c8", "0x83ad34bab18f6b7553b2772ca254a0ed13b66331")
	//trade, err := GetTrade("Polygon", "0xdf5d751d90b79c1a5f9df08b5d7165b225f4dcc46093c8985c369ce3a041b9c1", "0x83ad34bab18f6b7553b2772ca254a0ed13b66331")

	//trade, err := GetTrade("TRON", "0x43f23b8b7fe5505bff3611eabfa3cc4fc8eaee6894f5ce03c1b5b03eb03fdddd", "TRDvJ5bnjsJ77pSA8Tpba9d297PxgK7H7y")
	//trade, err := GetTrade("TRON", "43f23b8b7fe5505bff3611eabfa3cc4fc8eaee6894f5ce03c1b5b03eb03fdddd", "TRDvJ5bnjsJ77pSA8Tpba9d297PxgK7H7y")

	//trade, err := GetTrade("BSC", "0xb5d7ca794df401234bc490cf9f8e46121126b75eaa865dc6391e476fd80db8b1", "0xbab8580b9641bc98eaede2ec894d8cae0bda11b6")

	//trade, err := GetTrade("Ethereum", "0x7c8524d9523ef8bf326b07c3fe24131b0195900d5fb34eee6e1473306110fa9f", "0xad0243d3e9cf3fc6ce180cf3a5d23e06bccf4bc8")
	//trade, err := GetTrade("Ethereum", "0x7c8524d9523ef8bf326b07c3fe24131b0195900d5fb34eee6e1473306110fa91", "0xad0243d3e9cf3fc6ce180cf3a5d23e06bccf4bc8")

	//trade, err := GetTrade("Solana", "3acgfRLi3FaSzMPeBuGMrkKiNpHbCAJPuK1G9r8pnFvsEv5mN4w2oQxgunJE6P75LZ9ZZgyvKm6q4NXRV5VZXjDs", "4xtxLJpJ6Fqgoo7u8QEqyoc8d6f19ztU93uCbjjgukwq")
	//trade, err := GetTrade("Solana", "3acgfRLi3FaSzMPeBuGMrkKiNpHbCAJPuK1G9r8pnFvsEv5mN4w2oQxgunJE6P75LZ9ZZgyvKm6q4NXRV5VZXjDs", "4xtxLJpJ6Fqgoo7u8QEqyoc8d6f19ztU93uCbjjgukwq")
	//trade, err := GetTrade("Solana", "52iR4gmq7S1vBwj5yLsWkMWBJZtM5dT7tP5AJqKvU1hXrwWb8cxy9Eiu9s1QpRGviobQSssQ9jQNP8Vf1bXWy1dE", "4xtxLJpJ6Fqgoo7u8QEqyoc8d6f19ztU93uCbjjgukwq")
	//trade, err := GetTrade("Solana", "4Au7NrJ27YPk8VaBvnRGJEh3rrBZsFJy1jEYRGMk9d44ZayPWichRkLAPbHZiYRFjL1jELnP7H8wHxpKE7NHZEcX", "4xtxLJpJ6Fqgoo7u8QEqyoc8d6f19ztU93uCbjjgukwq")
	//trade, err := GetTrade("Solana", "4xEKKchQBJ3LroYR552Ks4ovKoxtSLf1HGBuG5fB8obX1j3VAicRb9MdZod34yj8NAHmcTr32JA3id9KwGmBSxCB", "4xtxLJpJ6Fqgoo7u8QEqyoc8d6f19ztU93uCbjjgukwq")

	//trade, err := GetTrade("Nano", "13DCFD2600736B0CE46E0AFFF0E537BB41357634EE12043C60D06CE5EDB1E563", "nano_1hphkwo88exyrccjmeiwjohpk1cgxi5w6gmowodxagmpiczh81bg447dutfj")
	//trade, err := GetTrade("Nano", "13DCFD2600736B0CE46E0AFFF0E537BB41357634EE12043C60D06CE5EDB1E513", "nano_1hphkwo88exyrccjmeiwjohpk1cgxi5w6gmowodxagmpiczh81bg447dutfj")
	code, trade, err := GetTrade("Nano", "FF16F6C651044FDC591C43D695986CF617FB32036FE7E312F36346D0DC5EDAD7", "nano_338yr6inx5tbg8ii6f6reffw3x7dj557xh9itsquk4z4cw6bg7zy7ichqq16")

	//trade, err := GetTrade("Ravencoin", "9b6f315e3892ac1ccb6813852877b19931f763119b6183925c40b0d6caf225fc", "RJEbayw2FGvTBFQRGdqMPcyGDhAhKVrYNU")
	//trade, err := GetTrade("Ravencoin", "bace2a4d86b1d67acee0dc1291779b9ef92ba9a8d64e6da9190a8cd91fa59e2c", "RJEbayw2FGvTBFQRGdqMPcyGDhAhKVrYNU")

	if nil != err {
		t.Log(code)
		t.Fatal(err)
	}
	t.Log(trade)
}

func TestPostWithdrawLTC(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	wg := sync.WaitGroup{}
	addr := []string{"ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklfz"}
	amount := []string{"19330"}
	wg.Add(len(addr))

	for i := 0; i < len(addr); i++ {
		go func() {
			defer func() {
				if e := recover(); nil != e {

				}
				wg.Done()
			}()
			tradeId := util.GetInstanceByRandomUtil().RandCharacterString(32)
			t.Log(tradeId)
			// rvn RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a
			res, err := PostWithdraw(tradeId,
				addr[i], "LTC-Litecoin",
				amount[i], "")
			t.Log(res)
			t.Log(err)
		}()

	}
	wg.Wait()
	t.Log("finish")
}

func TestPostWithdrawDoge(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	wg := sync.WaitGroup{}
	//addr := []string{"DTKjLFcLAHpH1UzB6k3tbq89UusqnnA2vR", "DRoxxxUPVTeC2dvmM1VA7nDEqiMhkkwfYJ"}
	//amount := []string{"66085001", "2000900"}

	addr := []string{"DTKjLFcLAHpH1UzB6k3tbq89UusqnnA2vR"}
	amount := []string{"21706100"}

	wg.Add(len(addr))

	for i := 0; i < len(addr); i++ {
		go func() {
			defer func() {
				if e := recover(); nil != e {

				}
				wg.Done()
			}()
			tradeId := util.GetInstanceByRandomUtil().RandCharacterString(32)
			t.Log(tradeId)
			// rvn RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a
			res, err := PostWithdraw(tradeId,
				addr[i], "DOGE-Dogecoin",
				amount[i], "")
			t.Log(res)
			t.Log(err)
		}()

	}
	wg.Wait()
	t.Log("finish")
}

func TestPostWithdrawRVN(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	addr := []string{"RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a", "RFKKNiPGGxBb4dt8ghY8RG6ygyM45YJjoE", "RJEbayw2FGvTBFQRGdqMPcyGDhAhKVrYNU"}
	//amount := []string{"180577500", "120695000", "133777500"}
	amount := []string{"5157500", "5500", "6157500"}

	//20695000
	//89577500
	//33777500

	for i := 0; i < len(addr); i++ {
		go func() {
			defer func() {
				if e := recover(); nil != e {

				}
				wg.Done()
			}()
			tradeId := util.GetInstanceByRandomUtil().RandCharacterString(32)
			t.Log(tradeId)
			// rvn RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a
			res, err := PostWithdraw(tradeId,
				addr[i], "RVN-Ravencoin",
				amount[i], "")
			t.Log(res)
			t.Log(err)
		}()

	}
	wg.Wait()
	t.Log("finish")

}

func TestPostWithdraw(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	tradeId := util.GetInstanceByRandomUtil().RandCharacterString(32)
	t.Log(tradeId)
	// rvn RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a
	res, err := PostWithdraw(tradeId,
		"RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a", "RVN-Ravencoin",
		"1010000", "")

	//res, err := PostWithdraw(tradeId,
	//	"nano_1nk7kf11j4uhwqmry818xxhxqq9ickwrmm3xtwpzzp3ybjz75yuoeuc7chz8", "XNO-Nano",
	//	"145000000000000000000000000000", "")

	//res, err := PostWithdraw(tradeId,
	//	"0x9b5836EdC7647C83628e12098a81B3f2b1800340", "USDT-ERC20",
	//	"20000000", "")

	//res, err := PostWithdraw(tradeId,
	//	"2eRJS1WHDEZjFK4333moyMW8H2d3k9D3k8ZaSsCnnAh1", "SOL-Solana",
	//	"10000", "") // code=2014,tip=最小转账金额 3500000

	//res, err := PostWithdraw(tradeId,
	//	"2eRJS1WHDEZjFK4333moyMW8H2d3k9D3k8ZaSsCnnAh1", "SOL-Solana",
	//	"3500000", "")

	//res, err := PostWithdraw(tradeId,
	//	"2eRJS1WHDEZjFK4333moyMW8H2d3k9D3k8ZaSsCnnAh1", "USDC-Solana",
	//	"2000000", "")

	//res, err := PostWithdraw(tradeId,
	//	"TDwUgqkqJ8BTnDUwx9fWQPLteUXftexr2t", "TRX-TRON",
	//	"10000", "")

	if nil != err {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetWithdrawInfo(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	//res, err := GetWithdrawInfo("SuD1Shtd3qvDp2cyunkqfHjaR8BKGZ0M")
	//res, err := GetWithdrawInfo("1dw3zBpvOs6uZD0Ky62xgqbZ8KbIOa7R")
	res, err := GetWithdrawInfo("1dw3zBpvOs6uZD0Ky62xgqbZ8KbIOa71")

	if nil != err {
		t.Fatal(err)
	}
	buf, _ := json.Marshal(res)
	t.Log(string(buf))

}

func TestReqLimit(t *testing.T) {
	if e := Init(priKeyG, priBaseUrlG, merchantCodeG, platformPubKeyG); nil != e {
		t.Fatal(e)
	}

	wg := sync.WaitGroup{}
	wg.Add(5)

	counter := atomic.Int64{}

	for i := 0; i < 5; i++ {
		go func(t *testing.T) {
			defer func() { wg.Done() }()
			for j := 0; j < 100; j++ {
				counter.Add(1)
				res, err := GetWithdrawInfo("SuD1Shtd3qvDp2cyunkqfHjaR8BKGZ0M")

				if nil != err {
					t.Log(counter.Load(), "-", err)
					return
				}
				buf, _ := json.Marshal(res)
				t.Log(counter.Load(), "-", string(buf))
			}
		}(t)
	}

	wg.Wait()

}
