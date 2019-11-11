package cmd

import (
	"testing"

	"github.com/aeternity/aepp-sdk-go/v7/naet"
)

func Test_printNameEntry(t *testing.T) {
	type args struct {
		conn naet.GetNameEntryByNamer
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Name from testnet",
			args: args{
				conn: &mockGetNameEntryByNamer{
					nameEntry: `{"id":"nm_WBWNnGHa2snFGFyPZAQh2hsqw4g1oXsueoxGYfB3vi3d6YdhB","pointers":[{"id":"ak_2pvi7arAQTS71XTUapv2PwShKomcPhWUp6Hxu66BDQpeFumkJP","key":"account_pubkey"}],"ttl":157928}`,
				},
				name: "lolilalus.test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printNameEntry(tt.args.conn, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("printNameEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printAccount(t *testing.T) {
	type args struct {
		conn      naet.GetAccounter
		accountID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal account",
			args: args{
				conn:      &mockGetAccounter{account: `{"balance":1600000000000000077131306000000000000000,"id":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","kind":"basic","nonce":0}`},
				accountID: "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printAccount(tt.args.conn, tt.args.accountID); (err != nil) != tt.wantErr {
				t.Errorf("printAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printMicroBlockAndTransactions(t *testing.T) {
	type args struct {
		conn   getMicroBlockHeaderTransactions
		mbHash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Microblock with 5 txs from testnet",
			args: args{
				conn: &mockGetMicroBlockHeaderTransactions{
					mbHeader: `{"hash":"mh_28qJfYLgxL1pESJwnke1AzCrHu4LAHNdW4qb4rAv9k28TwH3Fk","height":108438,"pof_hash":"no_fraud","prev_hash":"mh_2FkCx11HiP6z2qfaVDbEkpAwJVmVWgCJj5PVjFRDVa7uru2Kid","prev_key_hash":"kh_2Ke7nqWNm6ADC5hnCnsxwAVrpXsNYDwzF1fkEFG7AS8DDMc8U8","signature":"sg_KmAAvxM3iwALyMGRjo6Ddh7ohoj9BUVgrYoGeVCDp4sP54DmmgbJnz96wLv4g9rDkEfWEgih14cL7rsnyCxEDK3nxParZ","state_hash":"bs_2tFVszM5mCCLcNpiVYKsTTjZLXaKjz9AHxnTNryjmumW7JDk7z","time":1562935271123,"txs_hash":"bx_24S2Snxot252rbcAJNpwuK87jGjdp6ynr4Ghcs3TK8AFNYiifp","version":3}`,
					mbTxs:    `{"transactions":[{"block_hash":"mh_YR2XivA1vNopkxewAwuF8VbNJwtKgioV9TUFyRHynMc77YQBM","block_height":108457,"hash":"th_2auAXYLdLXXtRsgujVsnCQtb3QrJD5HcQtdcuFwt69Jtw2DVnD","signatures":["sg_JKcXEWDXrLf9TTffeHGkuKj7JCi8ev4o7uLWSbmoegmJB6CWNpX5uHsfz4CLGqp9iwktQLsNow2FbEb5SGBKpgfHRgHET"],"tx":{"abi_version":1,"amount":0,"call_data":"cb_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","caller_id":"ak_2MDUVSTt1ANKo2oL92L3uXe1WLKtpemKvmAydUdiQTvTpp85Nb","contract_id":"ct_t4sf8hJttG2s9qXX68q1baTFDKrkrXxpe76aoLkn6xgfsyDZo","fee":2062980000000000,"gas":1579000,"gas_price":1000000000,"nonce":834,"type":"ContractCallTx","version":1}}]}`,
				},
				mbHash: "mh_28qJfYLgxL1pESJwnke1AzCrHu4LAHNdW4qb4rAv9k28TwH3Fk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printMicroBlockAndTransactions(tt.args.conn, tt.args.mbHash); (err != nil) != tt.wantErr {
				t.Errorf("printMicroBlockAndTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printKeyBlockByHash(t *testing.T) {
	type args struct {
		conn   naet.GetKeyBlockByHasher
		kbHash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Key block from testnet",
			args: args{
				conn: &mockGetKeyBlockByHasher{
					keyBlock: `{"beneficiary":"ak_2iBPH7HUz3cSDVEUWiHg76MZJ6tZooVNBmmxcgVK6VV8KAE688","hash":"kh_oDqm8GN52dHmPr14QEZMJvfVP66tm8eBmWxLQmgZRJXgT7JFy","height":108441,"info":"cb_AAAAAfy4hFE=","miner":"ak_aE7aDeuV9pXYigeiLQHWdFK1ikNcCYMx1hQaMMuxuVc8SSgL1","nonce":2722962700732600136,"pow":[7237248,17339428,19311923,24469735,41960858,51470860,65768197,67950693,89283243,92572228,121706880,172095449,211686159,216835594,233435022,262747378,265967366,267111342,267230747,268929332,272493684,272942478,274270131,288317690,311217919,318687310,330374834,356441059,376376355,381000412,387400363,394749719,395982845,422309957,436094149,437906840,495692399,503123017,509591917,513558670,514890693,534392063],"prev_hash":"kh_2JYYqcBCtt2DGrifj8Z72aDkuxm53Jie4bD5pxn5uNe9ATieZn","prev_key_hash":"kh_2JYYqcBCtt2DGrifj8Z72aDkuxm53Jie4bD5pxn5uNe9ATieZn","state_hash":"bs_241bfFjn54ymFEeQYHUDbqUjTxcRk8kbY3mLqSxTKy3qoZeapz","target":538266778,"time":1562935866886,"version":3}`,
				},
				kbHash: `kh_oDqm8GN52dHmPr14QEZMJvfVP66tm8eBmWxLQmgZRJXgT7JFy`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printKeyBlockByHash(tt.args.conn, tt.args.kbHash); (err != nil) != tt.wantErr {
				t.Errorf("printKeyBlockByHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printTransactionByHash(t *testing.T) {
	type args struct {
		conn   naet.GetTransactionByHasher
		txHash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Spend Tx from testnet",
			args: args{
				conn: &mockGetTransactionByHasher{
					transaction: `{
						"block_hash": "mh_2LHbqc3iT3kJJgsvrDUawnUwMmaoDa4ak8vyH6sdRU8N4FDBmT",
						"block_height": 108313,
						"hash": "th_2QnsbqNA7CXXikzJy1kNGaEt98vU5woYMMMGQq7ieXKRAWRzqA",
						"signatures": [
						  "sg_Peh639PMQixTm63SP3Tebq2qc6PbNf2JY4kuBktXrLHE9NKyU2p6LKUDgL2NfwgAGNAuPeGr6TMvEFzPdUbi8HnwEmDHN"
						],
						"tx": {
						  "amount": 1598765432100000000,
						  "fee": 16920000000000,
						  "nonce": 157,
						  "payload": "ba_Xfbg4g==",
						  "recipient_id": "ak_9HCjPndyrL7vZBjKxacSdLiR64ChrxUYjY1V6fHxzfpmKphPb",
						  "sender_id": "ak_QxNXZsZnDEePC7XzmB3vUeUYAowDsMRXeFww1oPqPg2nEqfve",
						  "ttl": 108323,
						  "type": "SpendTx",
						  "version": 1
						}
					  }`,
				},
				txHash: "th_2QnsbqNA7CXXikzJy1kNGaEt98vU5woYMMMGQq7ieXKRAWRzqA",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printTransactionByHash(tt.args.conn, tt.args.txHash); (err != nil) != tt.wantErr {
				t.Errorf("printTransactionByHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_printOracleByPubkey(t *testing.T) {
	type args struct {
		conn     naet.GetOracleByPubkeyer
		oracleID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Oracle from mainnet",
			args: args{
				conn: &mockGetOracleByPubkeyer{
					oracle: `{
						"abi_version": 0,
						"id": "ok_28QDg7fkF5qiKueSdUvUBtCYPJdmMEoS73CztzXCRAwMGKHKZh",
						"query_fee": 1000000000000000000,
						"query_format": "string",
						"response_format": "string",
						"ttl": 108596
					  }`,
				},
				oracleID: "ok_28QDg7fkF5qiKueSdUvUBtCYPJdmMEoS73CztzXCRAwMGKHKZh",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printOracleByPubkey(tt.args.conn, tt.args.oracleID); (err != nil) != tt.wantErr {
				t.Errorf("printOracleByPubkey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
