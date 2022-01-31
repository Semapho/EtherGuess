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

func (c *Client) generateAccount() (string, string, error) {
	prvKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	prvData := crypto.FromECDSA(prvKey)
	pubAddress := crypto.PubkeyToAddress(prvKey.PublicKey).Hex()
	return hexutil.Encode(prvData), pubAddress, nil
}

func (c *Client) getBalance(address string) (*big.Int, error) {
	return c.conn.BalanceAt(context.Background(), common.HexToAddress(address), nil)
}

func (c *Client) parseEther(data *big.Int) *big.Float {
	fData := new(big.Float)
	fData.SetString(data.String())
	value := new(big.Float).Quo(fData, big.NewFloat(math.Pow(10, 18)))
	return value
}

func (c *Client) guessEther(finishChan chan<- bool) {
	for {
		privateKey, address, err := c.generateAccount()
		if err != nil {
			log.Errorf("generate account error: %v", err)
			continue
		}
		balance, err := c.getBalance(address)
		if err != nil {
			log.Errorf("get balance error: %v", err)
			continue
		}
		atomic.AddUint64(&c.total, 1)
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
