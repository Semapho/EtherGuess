# EtherGuess

![EtherGuess](https://i.imgur.com/XNGXjPT_d.png?maxwidth=1520)

This program generates random Ethereum private keys and check for account balance, if the account have ether, it will
print the address and private key.

Use it on your own risk, it is just for fun.

### Run

You need to prepare your Ethereum node, or register [INFURA](https://infura.io/) account first.

```shell
ether-guess -p 10 ETHER_NODE_URL
```

### Build

Install golang and run `make` command to build binary file.
