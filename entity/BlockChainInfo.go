package entity

import "math/big"

type BlockChainInfo struct {
	Chain string
	Blocks int64
	Headers int64
	BestBlockHash string
	Difficulty float64
	Mediantime float64
	Verificationprogress float64
	Initialblockdownload bool
	Chainblock string
	Size_on_desk big.Int
	Pruned bool

	Softfork SortFork
	Warnings string

	//Chain string`json:"chain"`
	//Blocks int`json:"blocks"`
	//Headers int`json:"headers"`
	//BestBlockHash string`json:"bestblockhash"`
	//Difficulty float64`json:"difficulty"`
	//Mediantime big.Int`json:"mediantime"`
	//Verificationprogress float64`json:"verificationprogress"`
	//Initialblockdownload bool`json:"initialblockdownload"`
	//Chainword string`json:"chainword"`
	//Size_on_disk big.Int`json:"size_on_disk"`
	//Pruned bool`json:"pruned"`
	//Pruneheight int`json:"pruneheight"`
	//Automatic_pruning bool`json:"automatic_pruning"`
	//Prune_target_size int64`json:"prune_target_size"`
	//
	//
	//Softfork BipSortFork`json:"softfork"`
	//Warnings string`json:"warnings"`

}
