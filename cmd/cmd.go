package main

import (
	"eth/internal/guess"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var parallel int

func main() {
	prepare()
	err := runCmd()
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd() error {
	// 设置命令行
	var rootCmd = &cobra.Command{
		Use:     "ether-guess",
		Short:   "Ethereum Guess",
		Example: "ether-guess -p 10 ETHER_NODE_URL",
		Args:    cobra.ExactArgs(1),
		Run:     cmdRunFunc,
	}
	rootCmd.Flags().IntVarP(&parallel, "parallel", "p", 10, "parallel")
	return rootCmd.Execute()
}

func cmdRunFunc(cmd *cobra.Command, args []string) {
	rpcURL := args[0]
	if rpcURL == "" {
		log.Fatal("rpc url empty")
	}
	guessClient, err := guess.NewClient(rpcURL)
	if err != nil {
		log.Fatalf("connect to rpc %s error: %s", rpcURL, err)
	}
	defer guessClient.Close()
	guessClient.Run(parallel)
}
