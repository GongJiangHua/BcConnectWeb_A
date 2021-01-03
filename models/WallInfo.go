package models

type WallInfo struct {
	WalletName              string
	WalletVersion           int
	Balance                 float64
	Unconfirmed_balance     float64
	Immature_balance        float64
	Txcoun                  int
	Keypoololdest           int
	Keypoolsize             int
	Hdseedid                string
	Keypoolsize_hd_internal int
	Paytxfee                float64
	Private_keys_enabled    bool
	Avoid_reuse             bool
	Scanning                bool
}
