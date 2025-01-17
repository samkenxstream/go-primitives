// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2022-04-02 15:53:00.363318 &#43;0200 CEST m=&#43;0.001941418
// using data from coins.yml
package coin

import (
	"fmt"
)

const (
	coinPrefix  = "c"
	tokenPrefix = "t"
)

// Coin is the native currency of a blockchain
type Coin struct {
	ID               uint
	Handle           string
	Symbol           string
	Name             string
	Decimals         uint
	BlockTime        int
	MinConfirmations int64
}

type AssetID string

func (c *Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Name, c.ID)
}

func (c Coin) AssetID() AssetID {
	return AssetID(coinPrefix + fmt.Sprint(c.ID))
}

func (c Coin) TokenAssetID(t string) AssetID {
	result := c.AssetID()
	if len(t) > 0 {
		result += AssetID("_" + tokenPrefix + t)
	}

	return result
}

const (
	ETHEREUM     = 60
	CLASSIC      = 61
	ICON         = 74
	COSMOS       = 118
	RIPPLE       = 144
	STELLAR      = 148
	POA          = 178
	TRON         = 195
	FIO          = 235
	NIMIQ        = 242
	IOTEX        = 304
	ZILLIQA      = 313
	AION         = 425
	AETERNITY    = 457
	KAVA         = 459
	THETA        = 500
	BINANCE      = 714
	VECHAIN      = 818
	CALLISTO     = 820
	TOMOCHAIN    = 889
	THUNDERTOKEN = 1001
	ONTOLOGY     = 1024
	TEZOS        = 1729
	KIN          = 2017
	NEBULAS      = 2718
	GOCHAIN      = 6060
	WANCHAIN     = 5718350
	WAVES        = 5741564
	BITCOIN      = 0
	LITECOIN     = 2
	DOGE         = 3
	DASH         = 5
	VIACOIN      = 14
	GROESTLCOIN  = 17
	ZCASH        = 133
	FIRO         = 136
	BITCOINCASH  = 145
	RAVENCOIN    = 175
	QTUM         = 2301
	ZELCASH      = 19167
	DECRED       = 42
	ALGORAND     = 283
	NANO         = 165
	DIGIBYTE     = 20
	HARMONY      = 1023
	KUSAMA       = 434
	POLKADOT     = 354
	SOLANA       = 501
	NEAR         = 397
	ELROND       = 508
	SMARTCHAIN   = 20000714
	FILECOIN     = 461
	OASIS        = 474
	MONACOIN     = 22
	BITCOINGOLD  = 156
	EOS          = 194
	TERRA        = 330
	BAND         = 494
	NEO          = 888
	CARDANO      = 1815
	NULS         = 8964
	POLYGON      = 966
	THORCHAIN    = 931
	OPTIMISM     = 10000070
	XDAI         = 10000100
	AVALANCHEC   = 10009000
	HECO         = 10000553
	FANTOM       = 10000250
	ARBITRUM     = 10042221
	CELO         = 52752
	RONIN        = 10002020
	OSMOSIS      = 10000118
	CRONOS       = 10000025
)

var Coins = map[uint]Coin{
	ETHEREUM: {
		ID:               60,
		Handle:           "ethereum",
		Symbol:           "ETH",
		Name:             "Ethereum",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	CLASSIC: {
		ID:               61,
		Handle:           "classic",
		Symbol:           "ETC",
		Name:             "Ethereum Classic",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	ICON: {
		ID:               74,
		Handle:           "icon",
		Symbol:           "ICX",
		Name:             "ICON",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	COSMOS: {
		ID:               118,
		Handle:           "cosmos",
		Symbol:           "ATOM",
		Name:             "Cosmos",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	RIPPLE: {
		ID:               144,
		Handle:           "ripple",
		Symbol:           "XRP",
		Name:             "Ripple",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	STELLAR: {
		ID:               148,
		Handle:           "stellar",
		Symbol:           "XLM",
		Name:             "Stellar",
		Decimals:         7,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	POA: {
		ID:               178,
		Handle:           "poa",
		Symbol:           "POA",
		Name:             "Poa",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	TRON: {
		ID:               195,
		Handle:           "tron",
		Symbol:           "TRX",
		Name:             "Tron",
		Decimals:         6,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	FIO: {
		ID:               235,
		Handle:           "fio",
		Symbol:           "FIO",
		Name:             "FIO",
		Decimals:         9,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	NIMIQ: {
		ID:               242,
		Handle:           "nimiq",
		Symbol:           "NIM",
		Name:             "Nimiq",
		Decimals:         5,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	IOTEX: {
		ID:               304,
		Handle:           "iotex",
		Symbol:           "IOTX",
		Name:             "IoTeX",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	ZILLIQA: {
		ID:               313,
		Handle:           "zilliqa",
		Symbol:           "ZIL",
		Name:             "Zilliqa",
		Decimals:         12,
		BlockTime:        30000,
		MinConfirmations: 1,
	},
	AION: {
		ID:               425,
		Handle:           "aion",
		Symbol:           "AION",
		Name:             "Aion",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	AETERNITY: {
		ID:               457,
		Handle:           "aeternity",
		Symbol:           "AE",
		Name:             "Aeternity",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	KAVA: {
		ID:               459,
		Handle:           "kava",
		Symbol:           "KAVA",
		Name:             "Kava",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	THETA: {
		ID:               500,
		Handle:           "theta",
		Symbol:           "THETA",
		Name:             "Theta",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	BINANCE: {
		ID:               714,
		Handle:           "binance",
		Symbol:           "BNB",
		Name:             "BNB",
		Decimals:         8,
		BlockTime:        1000,
		MinConfirmations: 2,
	},
	VECHAIN: {
		ID:               818,
		Handle:           "vechain",
		Symbol:           "VET",
		Name:             "VeChain Token",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	CALLISTO: {
		ID:               820,
		Handle:           "callisto",
		Symbol:           "CLO",
		Name:             "Callisto",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	TOMOCHAIN: {
		ID:               889,
		Handle:           "tomochain",
		Symbol:           "TOMO",
		Name:             "TOMO",
		Decimals:         18,
		BlockTime:        4000,
		MinConfirmations: 0,
	},
	THUNDERTOKEN: {
		ID:               1001,
		Handle:           "thundertoken",
		Symbol:           "TT",
		Name:             "ThunderCore",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	ONTOLOGY: {
		ID:               1024,
		Handle:           "ontology",
		Symbol:           "ONT",
		Name:             "Ontology",
		Decimals:         0,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	TEZOS: {
		ID:               1729,
		Handle:           "tezos",
		Symbol:           "XTZ",
		Name:             "Tezos",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	KIN: {
		ID:               2017,
		Handle:           "kin",
		Symbol:           "KIN",
		Name:             "Kin",
		Decimals:         5,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	NEBULAS: {
		ID:               2718,
		Handle:           "nebulas",
		Symbol:           "NAS",
		Name:             "Nebulas",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	GOCHAIN: {
		ID:               6060,
		Handle:           "gochain",
		Symbol:           "GO",
		Name:             "GoChain GO",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	WANCHAIN: {
		ID:               5718350,
		Handle:           "wanchain",
		Symbol:           "WAN",
		Name:             "Wanchain",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	WAVES: {
		ID:               5741564,
		Handle:           "waves",
		Symbol:           "WAVES",
		Name:             "WAVES",
		Decimals:         8,
		BlockTime:        30000,
		MinConfirmations: 1,
	},
	BITCOIN: {
		ID:               0,
		Handle:           "bitcoin",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	LITECOIN: {
		ID:               2,
		Handle:           "litecoin",
		Symbol:           "LTC",
		Name:             "Litecoin",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
	},
	DOGE: {
		ID:               3,
		Handle:           "doge",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	DASH: {
		ID:               5,
		Handle:           "dash",
		Symbol:           "DASH",
		Name:             "Dash",
		Decimals:         8,
		BlockTime:        180000,
		MinConfirmations: 0,
	},
	VIACOIN: {
		ID:               14,
		Handle:           "viacoin",
		Symbol:           "VIA",
		Name:             "Viacoin",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
	},
	GROESTLCOIN: {
		ID:               17,
		Handle:           "groestlcoin",
		Symbol:           "GRS",
		Name:             "Groestlcoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	ZCASH: {
		ID:               133,
		Handle:           "zcash",
		Symbol:           "ZEC",
		Name:             "Zcash",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
	},
	FIRO: {
		ID:               136,
		Handle:           "firo",
		Symbol:           "FIRO",
		Name:             "Firo",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
	},
	BITCOINCASH: {
		ID:               145,
		Handle:           "bitcoincash",
		Symbol:           "BCH",
		Name:             "Bitcoin Cash",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	RAVENCOIN: {
		ID:               175,
		Handle:           "ravencoin",
		Symbol:           "RVN",
		Name:             "Raven",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	QTUM: {
		ID:               2301,
		Handle:           "qtum",
		Symbol:           "QTUM",
		Name:             "Qtum",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	ZELCASH: {
		ID:               19167,
		Handle:           "zelcash",
		Symbol:           "ZEL",
		Name:             "Zelcash",
		Decimals:         8,
		BlockTime:        120000,
		MinConfirmations: 0,
	},
	DECRED: {
		ID:               42,
		Handle:           "decred",
		Symbol:           "DCR",
		Name:             "Decred",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
	},
	ALGORAND: {
		ID:               283,
		Handle:           "algorand",
		Symbol:           "ALGO",
		Name:             "Algorand",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	NANO: {
		ID:               165,
		Handle:           "nano",
		Symbol:           "XNO",
		Name:             "Nano",
		Decimals:         30,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	DIGIBYTE: {
		ID:               20,
		Handle:           "digibyte",
		Symbol:           "DGB",
		Name:             "DigiByte",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
	},
	HARMONY: {
		ID:               1023,
		Handle:           "harmony",
		Symbol:           "ONE",
		Name:             "Harmony",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	KUSAMA: {
		ID:               434,
		Handle:           "kusama",
		Symbol:           "KSM",
		Name:             "Kusama",
		Decimals:         12,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	POLKADOT: {
		ID:               354,
		Handle:           "polkadot",
		Symbol:           "DOT",
		Name:             "Polkadot",
		Decimals:         10,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	SOLANA: {
		ID:               501,
		Handle:           "solana",
		Symbol:           "SOL",
		Name:             "Solana",
		Decimals:         9,
		BlockTime:        500,
		MinConfirmations: 0,
	},
	NEAR: {
		ID:               397,
		Handle:           "near",
		Symbol:           "NEAR",
		Name:             "NEAR",
		Decimals:         24,
		BlockTime:        2000,
		MinConfirmations: 0,
	},
	ELROND: {
		ID:               508,
		Handle:           "elrond",
		Symbol:           "eGLD",
		Name:             "Elrond",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	SMARTCHAIN: {
		ID:               20000714,
		Handle:           "smartchain",
		Symbol:           "BNB",
		Name:             "Smart Chain",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
	},
	FILECOIN: {
		ID:               461,
		Handle:           "filecoin",
		Symbol:           "FIL",
		Name:             "Filecoin",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
	},
	OASIS: {
		ID:               474,
		Handle:           "oasis",
		Symbol:           "ROSE",
		Name:             "Oasis",
		Decimals:         9,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	MONACOIN: {
		ID:               22,
		Handle:           "monacoin",
		Symbol:           "MONA",
		Name:             "Monacoin",
		Decimals:         8,
		BlockTime:        90000,
		MinConfirmations: 0,
	},
	BITCOINGOLD: {
		ID:               156,
		Handle:           "bitcoingold",
		Symbol:           "BTG",
		Name:             "Bitcoin Gold",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	EOS: {
		ID:               194,
		Handle:           "eos",
		Symbol:           "EOS",
		Name:             "EOS",
		Decimals:         4,
		BlockTime:        500,
		MinConfirmations: 0,
	},
	TERRA: {
		ID:               330,
		Handle:           "terra",
		Symbol:           "LUNA",
		Name:             "Terra",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	BAND: {
		ID:               494,
		Handle:           "band",
		Symbol:           "BAND",
		Name:             "BandChain",
		Decimals:         6,
		BlockTime:        2000,
		MinConfirmations: 0,
	},
	NEO: {
		ID:               888,
		Handle:           "neo",
		Symbol:           "NEO",
		Name:             "NEO",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	CARDANO: {
		ID:               1815,
		Handle:           "cardano",
		Symbol:           "ADA",
		Name:             "Cardano",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	NULS: {
		ID:               8964,
		Handle:           "nuls",
		Symbol:           "NULS",
		Name:             "NULS",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	POLYGON: {
		ID:               966,
		Handle:           "polygon",
		Symbol:           "MATIC",
		Name:             "Polygon",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	THORCHAIN: {
		ID:               931,
		Handle:           "thorchain",
		Symbol:           "RUNE",
		Name:             "THORChain",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	OPTIMISM: {
		ID:               10000070,
		Handle:           "optimism",
		Symbol:           "OETH",
		Name:             "Optimism Ethereum",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	XDAI: {
		ID:               10000100,
		Handle:           "xdai",
		Symbol:           "xDAI",
		Name:             "xDai",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	AVALANCHEC: {
		ID:               10009000,
		Handle:           "avalanchec",
		Symbol:           "AVAX",
		Name:             "Avalanche C-Chain",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	HECO: {
		ID:               10000553,
		Handle:           "heco",
		Symbol:           "HT",
		Name:             "Huobi ECO Chain",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	FANTOM: {
		ID:               10000250,
		Handle:           "fantom",
		Symbol:           "FTM",
		Name:             "Fantom",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	ARBITRUM: {
		ID:               10042221,
		Handle:           "arbitrum",
		Symbol:           "ARETH",
		Name:             "Arbitrum",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	CELO: {
		ID:               52752,
		Handle:           "celo",
		Symbol:           "CELO",
		Name:             "Celo",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	RONIN: {
		ID:               10002020,
		Handle:           "ronin",
		Symbol:           "RON",
		Name:             "Ronin",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	OSMOSIS: {
		ID:               10000118,
		Handle:           "osmosis",
		Symbol:           "OSMO",
		Name:             "Osmosis",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	CRONOS: {
		ID:               10000025,
		Handle:           "cronos",
		Symbol:           "CRO",
		Name:             "Cronos",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
}

var Chains = map[string]Coin{
	Ethereum().Handle: {
		ID:               60,
		Handle:           "ethereum",
		Symbol:           "ETH",
		Name:             "Ethereum",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Classic().Handle: {
		ID:               61,
		Handle:           "classic",
		Symbol:           "ETC",
		Name:             "Ethereum Classic",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	Icon().Handle: {
		ID:               74,
		Handle:           "icon",
		Symbol:           "ICX",
		Name:             "ICON",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Cosmos().Handle: {
		ID:               118,
		Handle:           "cosmos",
		Symbol:           "ATOM",
		Name:             "Cosmos",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Ripple().Handle: {
		ID:               144,
		Handle:           "ripple",
		Symbol:           "XRP",
		Name:             "Ripple",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Stellar().Handle: {
		ID:               148,
		Handle:           "stellar",
		Symbol:           "XLM",
		Name:             "Stellar",
		Decimals:         7,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Poa().Handle: {
		ID:               178,
		Handle:           "poa",
		Symbol:           "POA",
		Name:             "Poa",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Tron().Handle: {
		ID:               195,
		Handle:           "tron",
		Symbol:           "TRX",
		Name:             "Tron",
		Decimals:         6,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Fio().Handle: {
		ID:               235,
		Handle:           "fio",
		Symbol:           "FIO",
		Name:             "FIO",
		Decimals:         9,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Nimiq().Handle: {
		ID:               242,
		Handle:           "nimiq",
		Symbol:           "NIM",
		Name:             "Nimiq",
		Decimals:         5,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	Iotex().Handle: {
		ID:               304,
		Handle:           "iotex",
		Symbol:           "IOTX",
		Name:             "IoTeX",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Zilliqa().Handle: {
		ID:               313,
		Handle:           "zilliqa",
		Symbol:           "ZIL",
		Name:             "Zilliqa",
		Decimals:         12,
		BlockTime:        30000,
		MinConfirmations: 1,
	},
	Aion().Handle: {
		ID:               425,
		Handle:           "aion",
		Symbol:           "AION",
		Name:             "Aion",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Aeternity().Handle: {
		ID:               457,
		Handle:           "aeternity",
		Symbol:           "AE",
		Name:             "Aeternity",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	Kava().Handle: {
		ID:               459,
		Handle:           "kava",
		Symbol:           "KAVA",
		Name:             "Kava",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Theta().Handle: {
		ID:               500,
		Handle:           "theta",
		Symbol:           "THETA",
		Name:             "Theta",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Binance().Handle: {
		ID:               714,
		Handle:           "binance",
		Symbol:           "BNB",
		Name:             "BNB",
		Decimals:         8,
		BlockTime:        1000,
		MinConfirmations: 2,
	},
	Vechain().Handle: {
		ID:               818,
		Handle:           "vechain",
		Symbol:           "VET",
		Name:             "VeChain Token",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	Callisto().Handle: {
		ID:               820,
		Handle:           "callisto",
		Symbol:           "CLO",
		Name:             "Callisto",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Tomochain().Handle: {
		ID:               889,
		Handle:           "tomochain",
		Symbol:           "TOMO",
		Name:             "TOMO",
		Decimals:         18,
		BlockTime:        4000,
		MinConfirmations: 0,
	},
	Thundertoken().Handle: {
		ID:               1001,
		Handle:           "thundertoken",
		Symbol:           "TT",
		Name:             "ThunderCore",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Ontology().Handle: {
		ID:               1024,
		Handle:           "ontology",
		Symbol:           "ONT",
		Name:             "Ontology",
		Decimals:         0,
		BlockTime:        10000,
		MinConfirmations: 0,
	},
	Tezos().Handle: {
		ID:               1729,
		Handle:           "tezos",
		Symbol:           "XTZ",
		Name:             "Tezos",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	Kin().Handle: {
		ID:               2017,
		Handle:           "kin",
		Symbol:           "KIN",
		Name:             "Kin",
		Decimals:         5,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Nebulas().Handle: {
		ID:               2718,
		Handle:           "nebulas",
		Symbol:           "NAS",
		Name:             "Nebulas",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	Gochain().Handle: {
		ID:               6060,
		Handle:           "gochain",
		Symbol:           "GO",
		Name:             "GoChain GO",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	Wanchain().Handle: {
		ID:               5718350,
		Handle:           "wanchain",
		Symbol:           "WAN",
		Name:             "Wanchain",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
	},
	Waves().Handle: {
		ID:               5741564,
		Handle:           "waves",
		Symbol:           "WAVES",
		Name:             "WAVES",
		Decimals:         8,
		BlockTime:        30000,
		MinConfirmations: 1,
	},
	Bitcoin().Handle: {
		ID:               0,
		Handle:           "bitcoin",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	Litecoin().Handle: {
		ID:               2,
		Handle:           "litecoin",
		Symbol:           "LTC",
		Name:             "Litecoin",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
	},
	Doge().Handle: {
		ID:               3,
		Handle:           "doge",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	Dash().Handle: {
		ID:               5,
		Handle:           "dash",
		Symbol:           "DASH",
		Name:             "Dash",
		Decimals:         8,
		BlockTime:        180000,
		MinConfirmations: 0,
	},
	Viacoin().Handle: {
		ID:               14,
		Handle:           "viacoin",
		Symbol:           "VIA",
		Name:             "Viacoin",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
	},
	Groestlcoin().Handle: {
		ID:               17,
		Handle:           "groestlcoin",
		Symbol:           "GRS",
		Name:             "Groestlcoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	Zcash().Handle: {
		ID:               133,
		Handle:           "zcash",
		Symbol:           "ZEC",
		Name:             "Zcash",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
	},
	Firo().Handle: {
		ID:               136,
		Handle:           "firo",
		Symbol:           "FIRO",
		Name:             "Firo",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
	},
	Bitcoincash().Handle: {
		ID:               145,
		Handle:           "bitcoincash",
		Symbol:           "BCH",
		Name:             "Bitcoin Cash",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	Ravencoin().Handle: {
		ID:               175,
		Handle:           "ravencoin",
		Symbol:           "RVN",
		Name:             "Raven",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	Qtum().Handle: {
		ID:               2301,
		Handle:           "qtum",
		Symbol:           "QTUM",
		Name:             "Qtum",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
	},
	Zelcash().Handle: {
		ID:               19167,
		Handle:           "zelcash",
		Symbol:           "ZEL",
		Name:             "Zelcash",
		Decimals:         8,
		BlockTime:        120000,
		MinConfirmations: 0,
	},
	Decred().Handle: {
		ID:               42,
		Handle:           "decred",
		Symbol:           "DCR",
		Name:             "Decred",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
	},
	Algorand().Handle: {
		ID:               283,
		Handle:           "algorand",
		Symbol:           "ALGO",
		Name:             "Algorand",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
	},
	Nano().Handle: {
		ID:               165,
		Handle:           "nano",
		Symbol:           "XNO",
		Name:             "Nano",
		Decimals:         30,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Digibyte().Handle: {
		ID:               20,
		Handle:           "digibyte",
		Symbol:           "DGB",
		Name:             "DigiByte",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
	},
	Harmony().Handle: {
		ID:               1023,
		Handle:           "harmony",
		Symbol:           "ONE",
		Name:             "Harmony",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
	},
	Kusama().Handle: {
		ID:               434,
		Handle:           "kusama",
		Symbol:           "KSM",
		Name:             "Kusama",
		Decimals:         12,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	Polkadot().Handle: {
		ID:               354,
		Handle:           "polkadot",
		Symbol:           "DOT",
		Name:             "Polkadot",
		Decimals:         10,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	Solana().Handle: {
		ID:               501,
		Handle:           "solana",
		Symbol:           "SOL",
		Name:             "Solana",
		Decimals:         9,
		BlockTime:        500,
		MinConfirmations: 0,
	},
	Near().Handle: {
		ID:               397,
		Handle:           "near",
		Symbol:           "NEAR",
		Name:             "NEAR",
		Decimals:         24,
		BlockTime:        2000,
		MinConfirmations: 0,
	},
	Elrond().Handle: {
		ID:               508,
		Handle:           "elrond",
		Symbol:           "eGLD",
		Name:             "Elrond",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	Smartchain().Handle: {
		ID:               20000714,
		Handle:           "smartchain",
		Symbol:           "BNB",
		Name:             "Smart Chain",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
	},
	Filecoin().Handle: {
		ID:               461,
		Handle:           "filecoin",
		Symbol:           "FIL",
		Name:             "Filecoin",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
	},
	Oasis().Handle: {
		ID:               474,
		Handle:           "oasis",
		Symbol:           "ROSE",
		Name:             "Oasis",
		Decimals:         9,
		BlockTime:        6000,
		MinConfirmations: 0,
	},
	Monacoin().Handle: {
		ID:               22,
		Handle:           "monacoin",
		Symbol:           "MONA",
		Name:             "Monacoin",
		Decimals:         8,
		BlockTime:        90000,
		MinConfirmations: 0,
	},
	Bitcoingold().Handle: {
		ID:               156,
		Handle:           "bitcoingold",
		Symbol:           "BTG",
		Name:             "Bitcoin Gold",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
	},
	Eos().Handle: {
		ID:               194,
		Handle:           "eos",
		Symbol:           "EOS",
		Name:             "EOS",
		Decimals:         4,
		BlockTime:        500,
		MinConfirmations: 0,
	},
	Terra().Handle: {
		ID:               330,
		Handle:           "terra",
		Symbol:           "LUNA",
		Name:             "Terra",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Band().Handle: {
		ID:               494,
		Handle:           "band",
		Symbol:           "BAND",
		Name:             "BandChain",
		Decimals:         6,
		BlockTime:        2000,
		MinConfirmations: 0,
	},
	Neo().Handle: {
		ID:               888,
		Handle:           "neo",
		Symbol:           "NEO",
		Name:             "NEO",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Cardano().Handle: {
		ID:               1815,
		Handle:           "cardano",
		Symbol:           "ADA",
		Name:             "Cardano",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Nuls().Handle: {
		ID:               8964,
		Handle:           "nuls",
		Symbol:           "NULS",
		Name:             "NULS",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Polygon().Handle: {
		ID:               966,
		Handle:           "polygon",
		Symbol:           "MATIC",
		Name:             "Polygon",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Thorchain().Handle: {
		ID:               931,
		Handle:           "thorchain",
		Symbol:           "RUNE",
		Name:             "THORChain",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Optimism().Handle: {
		ID:               10000070,
		Handle:           "optimism",
		Symbol:           "OETH",
		Name:             "Optimism Ethereum",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Xdai().Handle: {
		ID:               10000100,
		Handle:           "xdai",
		Symbol:           "xDAI",
		Name:             "xDai",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Avalanchec().Handle: {
		ID:               10009000,
		Handle:           "avalanchec",
		Symbol:           "AVAX",
		Name:             "Avalanche C-Chain",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Heco().Handle: {
		ID:               10000553,
		Handle:           "heco",
		Symbol:           "HT",
		Name:             "Huobi ECO Chain",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Fantom().Handle: {
		ID:               10000250,
		Handle:           "fantom",
		Symbol:           "FTM",
		Name:             "Fantom",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Arbitrum().Handle: {
		ID:               10042221,
		Handle:           "arbitrum",
		Symbol:           "ARETH",
		Name:             "Arbitrum",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Celo().Handle: {
		ID:               52752,
		Handle:           "celo",
		Symbol:           "CELO",
		Name:             "Celo",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Ronin().Handle: {
		ID:               10002020,
		Handle:           "ronin",
		Symbol:           "RON",
		Name:             "Ronin",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Osmosis().Handle: {
		ID:               10000118,
		Handle:           "osmosis",
		Symbol:           "OSMO",
		Name:             "Osmosis",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
	},
	Cronos().Handle: {
		ID:               10000025,
		Handle:           "cronos",
		Symbol:           "CRO",
		Name:             "Cronos",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
	},
}

func Ethereum() Coin {
	return Coins[ETHEREUM]
}

func Classic() Coin {
	return Coins[CLASSIC]
}

func Icon() Coin {
	return Coins[ICON]
}

func Cosmos() Coin {
	return Coins[COSMOS]
}

func Ripple() Coin {
	return Coins[RIPPLE]
}

func Stellar() Coin {
	return Coins[STELLAR]
}

func Poa() Coin {
	return Coins[POA]
}

func Tron() Coin {
	return Coins[TRON]
}

func Fio() Coin {
	return Coins[FIO]
}

func Nimiq() Coin {
	return Coins[NIMIQ]
}

func Iotex() Coin {
	return Coins[IOTEX]
}

func Zilliqa() Coin {
	return Coins[ZILLIQA]
}

func Aion() Coin {
	return Coins[AION]
}

func Aeternity() Coin {
	return Coins[AETERNITY]
}

func Kava() Coin {
	return Coins[KAVA]
}

func Theta() Coin {
	return Coins[THETA]
}

func Binance() Coin {
	return Coins[BINANCE]
}

func Vechain() Coin {
	return Coins[VECHAIN]
}

func Callisto() Coin {
	return Coins[CALLISTO]
}

func Tomochain() Coin {
	return Coins[TOMOCHAIN]
}

func Thundertoken() Coin {
	return Coins[THUNDERTOKEN]
}

func Ontology() Coin {
	return Coins[ONTOLOGY]
}

func Tezos() Coin {
	return Coins[TEZOS]
}

func Kin() Coin {
	return Coins[KIN]
}

func Nebulas() Coin {
	return Coins[NEBULAS]
}

func Gochain() Coin {
	return Coins[GOCHAIN]
}

func Wanchain() Coin {
	return Coins[WANCHAIN]
}

func Waves() Coin {
	return Coins[WAVES]
}

func Bitcoin() Coin {
	return Coins[BITCOIN]
}

func Litecoin() Coin {
	return Coins[LITECOIN]
}

func Doge() Coin {
	return Coins[DOGE]
}

func Dash() Coin {
	return Coins[DASH]
}

func Viacoin() Coin {
	return Coins[VIACOIN]
}

func Groestlcoin() Coin {
	return Coins[GROESTLCOIN]
}

func Zcash() Coin {
	return Coins[ZCASH]
}

func Firo() Coin {
	return Coins[FIRO]
}

func Bitcoincash() Coin {
	return Coins[BITCOINCASH]
}

func Ravencoin() Coin {
	return Coins[RAVENCOIN]
}

func Qtum() Coin {
	return Coins[QTUM]
}

func Zelcash() Coin {
	return Coins[ZELCASH]
}

func Decred() Coin {
	return Coins[DECRED]
}

func Algorand() Coin {
	return Coins[ALGORAND]
}

func Nano() Coin {
	return Coins[NANO]
}

func Digibyte() Coin {
	return Coins[DIGIBYTE]
}

func Harmony() Coin {
	return Coins[HARMONY]
}

func Kusama() Coin {
	return Coins[KUSAMA]
}

func Polkadot() Coin {
	return Coins[POLKADOT]
}

func Solana() Coin {
	return Coins[SOLANA]
}

func Near() Coin {
	return Coins[NEAR]
}

func Elrond() Coin {
	return Coins[ELROND]
}

func Smartchain() Coin {
	return Coins[SMARTCHAIN]
}

func Filecoin() Coin {
	return Coins[FILECOIN]
}

func Oasis() Coin {
	return Coins[OASIS]
}

func Monacoin() Coin {
	return Coins[MONACOIN]
}

func Bitcoingold() Coin {
	return Coins[BITCOINGOLD]
}

func Eos() Coin {
	return Coins[EOS]
}

func Terra() Coin {
	return Coins[TERRA]
}

func Band() Coin {
	return Coins[BAND]
}

func Neo() Coin {
	return Coins[NEO]
}

func Cardano() Coin {
	return Coins[CARDANO]
}

func Nuls() Coin {
	return Coins[NULS]
}

func Polygon() Coin {
	return Coins[POLYGON]
}

func Thorchain() Coin {
	return Coins[THORCHAIN]
}

func Optimism() Coin {
	return Coins[OPTIMISM]
}

func Xdai() Coin {
	return Coins[XDAI]
}

func Avalanchec() Coin {
	return Coins[AVALANCHEC]
}

func Heco() Coin {
	return Coins[HECO]
}

func Fantom() Coin {
	return Coins[FANTOM]
}

func Arbitrum() Coin {
	return Coins[ARBITRUM]
}

func Celo() Coin {
	return Coins[CELO]
}

func Ronin() Coin {
	return Coins[RONIN]
}

func Osmosis() Coin {
	return Coins[OSMOSIS]
}

func Cronos() Coin {
	return Coins[CRONOS]
}
