package common

import (
	"errors"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cosmos/cosmos-sdk/types"
	dogchaincfg "github.com/eager7/dogd/chaincfg"
	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	whivechaincfg "github.com/nicholas-mwaura/twhd/chaincfg"
	btypes "gitlab.com/thorchain/binance-sdk/common/types"
)

var (
	EmptyChain = Chain("")
	BNBChain   = Chain("BNB")
	ETHChain   = Chain("ETH")
	BTCChain   = Chain("BTC")
	LTCChain   = Chain("LTC")
	BCHChain   = Chain("BCH")
	DOGEChain  = Chain("DOGE")
	THORChain  = Chain("THOR")
	WHIVEChain = Chain("WHIVE")

	SigningAlgoSecp256k1 = SigninAlgo("secp256k1")
	SigningAlgoEd25519   = SigninAlgo("ed25519")
)

type SigninAlgo string

// Chain is an alias of string , represent a block chain
type Chain string

// Chains represent a slice of Chain
type Chains []Chain

// Validate validates chain format, should consist only of uppercase letters
func (c Chain) Validate() error {
	if len(c) < 3 {
		return errors.New("chain id len is less than 3")
	}
	if len(c) > 10 {
		return errors.New("chain id len is more than 10")
	}
	for _, ch := range string(c) {
		if ch < 'A' || ch > 'Z' {
			return errors.New("chain id can consist only of uppercase letters")
		}
	}
	return nil
}

// NewChain create a new Chain and default the siging_algo to Secp256k1
func NewChain(chainID string) (Chain, error) {
	chain := Chain(strings.ToUpper(chainID))
	if err := chain.Validate(); err != nil {
		return chain, err
	}
	return chain, nil
}

// Equals compare two chain to see whether they represent the same chain
func (c Chain) Equals(c2 Chain) bool {
	return strings.EqualFold(c.String(), c2.String())
}

func (c Chain) IsTHORChain() bool {
	return c.Equals(THORChain)
}

// IsEmpty is to determinate whether the chain is empty
func (c Chain) IsEmpty() bool {
	return strings.TrimSpace(c.String()) == ""
}

// String implement fmt.Stringer
func (c Chain) String() string {
	// convert it to upper case again just in case someone created a ticker via Chain("rune")
	return strings.ToUpper(string(c))
}

// IsBNB determinate whether it is BNBChain
func (c Chain) IsBNB() bool {
	return c.Equals(BNBChain)
}

// GetSigningAlgo get the signing algorithm for the given chain
func (c Chain) GetSigningAlgo() SigninAlgo {
	switch c {
	case BNBChain, ETHChain, BTCChain, THORChain:
		return SigningAlgoSecp256k1
	default:
		return SigningAlgoSecp256k1
	}
}

// GetGasAsset chain's base asset
func (c Chain) GetGasAsset() Asset {
	switch c {
	case THORChain:
		return RuneNative
	case BNBChain:
		return BNBAsset
	case BTCChain:
		return BTCAsset
	case LTCChain:
		return LTCAsset
	case BCHChain:
		return BCHAsset
	case DOGEChain:
		return DOGEAsset
	case ETHChain:
		return ETHAsset
	case WHIVEChain:
		return WHIVEAsset
	default:
		return EmptyAsset
	}
}

// IsValidAddress make sure the address is correct for the chain
// And this also make sure testnet doesn't use mainnet address vice versa
func (c Chain) IsValidAddress(addr Address) bool {
	network := GetCurrentChainNetwork()
	prefix := c.AddressPrefix(network)
	return strings.HasPrefix(addr.String(), prefix)
}

// AddressPrefix return the address prefix used by the given network (testnet/mainnet)
func (c Chain) AddressPrefix(cn ChainNetwork) string {
	switch cn {
	case MockNet:
		switch c {
		case BNBChain:
			return btypes.TestNetwork.Bech32Prefixes()
		case ETHChain:
			return "0x"
		case THORChain:
			// TODO update this to use testnet address prefix
			return types.GetConfig().GetBech32AccountAddrPrefix()
		case BTCChain:
			return chaincfg.RegressionNetParams.Bech32HRPSegwit
		case LTCChain:
			return ltcchaincfg.RegressionNetParams.Bech32HRPSegwit
		case DOGEChain:
			return dogchaincfg.RegressionNetParams.Bech32HRPSegwit
		case WHIVEChain:
			return whivechaincfg.RegressionNetParams.Bech32HRPSegwit
		}
	case TestNet:
		switch c {
		case BNBChain:
			return btypes.TestNetwork.Bech32Prefixes()
		case ETHChain:
			return "0x"
		case THORChain:
			// TODO update this to use testnet address prefix
			return types.GetConfig().GetBech32AccountAddrPrefix()
		case BTCChain:
			return chaincfg.TestNet3Params.Bech32HRPSegwit
		case LTCChain:
			return ltcchaincfg.TestNet4Params.Bech32HRPSegwit
		case DOGEChain:
			return dogchaincfg.TestNet3Params.Bech32HRPSegwit
		case WHIVEChain:
			return whivechaincfg.TestNet3Params.Bech32HRPSegwit
		}
	case MainNet:
		switch c {
		case BNBChain:
			return btypes.ProdNetwork.Bech32Prefixes()
		case ETHChain:
			return "0x"
		case THORChain:
			return types.GetConfig().GetBech32AccountAddrPrefix()
		case BTCChain:
			return chaincfg.MainNetParams.Bech32HRPSegwit
		case LTCChain:
			return ltcchaincfg.MainNetParams.Bech32HRPSegwit
		case DOGEChain:
			return dogchaincfg.MainNetParams.Bech32HRPSegwit
		case WHIVEChain:
			return whivechaincfg.MainNetParams.Bech32HRPSegwit
		}
	}
	return ""
}

// Has check whether chain c is in the list
func (chains Chains) Has(c Chain) bool {
	for _, ch := range chains {
		if ch.Equals(c) {
			return true
		}
	}
	return false
}

// Distinct return a distinct set of chains, no duplicates
func (chains Chains) Distinct() Chains {
	var newChains Chains
	for _, chain := range chains {
		if !newChains.Has(chain) {
			newChains = append(newChains, chain)
		}
	}
	return newChains
}

func (chains Chains) Strings() []string {
	strings := make([]string, len(chains))
	for i, c := range chains {
		strings[i] = c.String()
	}
	return strings
}
