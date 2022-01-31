package guess

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"sync/atomic"
)

func (c *Client) generateAccount() (string, string) {
	prvKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	prvData := crypto.FromECDSA(prvKey)
	pubAddress := crypto.PubkeyToAddress(prvKey.PublicKey).Hex()
	return hexutil.Encode(prvData), pubAddress
}

func (c *Client) getBalance(address string) *big.Int {
	balance, err := c.conn.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		panic(err)
	}
	return balance
}

func (c *Client) parseEther(data *big.Int) *big.Float {
	fData := new(big.Float)
	fData.SetString(data.String())
	value := new(big.Float).Quo(fData, big.NewFloat(math.Pow(10, 18)))
	return value
}

func (c *Client) guessEther(finishChan chan<- bool) {
	for {
		privateKey, address := c.generateAccount()
		atomic.AddUint64(&c.total, 1)
		balance := c.getBalance(address)
		if balance.Cmp(big.NewInt(0)) == 1 {
			etherBalance := c.parseEther(balance)
			log.Infof("find account: %s, private key: %s, balance: %v\n", address, privateKey, etherBalance)
			finishChan <- true
			break
		} else {
			log.Infof("total: %d, account: %s, balance: %v\n", atomic.LoadUint64(&c.total), address, balance)
		}
	}
}
