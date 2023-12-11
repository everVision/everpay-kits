package schema

import (
	"strings"
)

const (
	TxVersionV1 = "v1"

	ChainTypeArweave    = "arweave"
	ChainTypeCrossArEth = "arweave,ethereum"
	ChainTypeEverpay    = "everpay"

	TxActionTransfer        = "transfer"
	TxActionMint            = "mint"
	TxActionBurn            = "burn"
	TxActionTransferOwner   = "transferOwner" // token owner
	TxActionAddWhiteList    = "addWhiteList"
	TxActionRemoveWhiteList = "removeWhiteList"
	TxActionPauseWhiteList  = "pauseWhiteList"
	TxActionAddBlackList    = "addBlackList"
	TxActionRemoveBlackList = "removeBlackList"
	TxActionPauseBlackList  = "pauseBlackList"
	TxActionPause           = "pause"

	ZeroAddress = "0x0000000000000000000000000000000000000000"
)

type Token struct {
	ID           string // On Native-Chain tokenId; Special AR token: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0xcc9141efa8c20c7df0778748255b1487957811be"
	Symbol       string
	ChainType    string                 // On everPay chainType; Special AR token: "arweave,ethereum"
	ChainID      string                 // On everPay chainId; Special AR token: "0,1"(mainnet) or "0,42"(testnet)
	Decimals     int                    // On everPay decimals
	TargetChains map[string]TargetChain // key: targetChainType
}

func (t *Token) Tag() string {
	return tag(t.ChainType, t.Symbol, t.ID)
}

type TargetChain struct {
	ChainID   string `json:"targetChainId"`
	ChainType string `json:"targetChainType"`         // e.g: "avalanche" "arweave" "ethereum","moon"
	Decimals  int    `json:"targetDecimals"`          // e.g: 18
	TokenId   string `json:"targetTokenId,omitempty"` // target chain token address
}

func tag(chainType, tokenSymbol, tokenID string) string {
	// process tokenId
	var id string
	switch chainType {
	case ChainTypeArweave:
		id = tokenID
	case ChainTypeCrossArEth: // now only AR token
		ids := strings.Split(tokenID, ",")
		if len(ids) != 2 {
			return "err_invalid_token"
		}

		ids[1] = strings.ToLower(ids[1])
		id = strings.Join(ids, ",")
	default: // "ethereum", "avalanche" and so on evm chain
		id = strings.ToLower(tokenID)
	}

	return strings.ToLower(chainType+"-"+tokenSymbol) + "-" + id
}
