package neo

import (
	"fmt"
	"os"

	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/polynetwork/poly-io-test/config"
	"golang.org/x/crypto/ssh/terminal"
)

func LoadAccount() *wallet.Account {
	// open the NEO wallet
	//neoAccount, err := wallet.NewAccountFromWIF(config.DefConfig.NeoWalletWIF)
	w, err := wallet.NewWalletFromFile(config.DefConfig.NeoWalletFile)
	if err != nil {
		panic("[NEO Relayer] Failed to open NEO wallet")
		return nil
	}

	fmt.Printf("Neo Wallet Password:")
	pwd, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic("[NEO Wallet] Invalid password entered")

	}
	neoPwd := string(pwd)

	err = w.DecryptAll(neoPwd)
	if err != nil {
		panic("[NEO Wallet] Failed to decrypt NEO account")
	}
	return w.Accounts[0]
}
