package guess

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn  *ethclient.Client
	total uint64
}

func NewClient(rpcURL string) (*Client, error) {
	log.Infof("connecting to ethereum node %s ...", rpcURL)
	client, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: client,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
