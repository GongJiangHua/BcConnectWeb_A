package models

type MiningInfo struct {
	Blocks        int
	Difficulty    int
	Networkhashps int
	Pooledtx      int
	Chain         string
	Warnings      string
}
