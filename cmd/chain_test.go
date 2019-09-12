package cmd

import (
	"flag"
	"fmt"
	"testing"

	"github.com/aeternity/aepp-sdk-go/v5/aeternity"
)

// Prefixing each test with Example makes go-test check the stdout
// For now, just verify that none of the commands segfault.

func init() {
	flag.BoolVar(&online, "online", false, "Run tests that need a running node on localhost:3013, Network ID ae_docker")
	flag.Parse()
	setPrivateNetParams()
}

func Test_topFunc(t *testing.T) {
	type args struct {
		conn aeternity.GetTopBlocker
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Normal KeyBlockOrMicroBlockHeader",
			args: args{
				conn: &mockGetTopBlocker{
					msg: `{"key_block":{"beneficiary":"ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi","hash":"kh_2H9nMAH8nLFtWeP6YEb4w19d48h6nhUZ2NubUg9Lo5KHLDHiem","height":119,"info":"cb_AAAAAfy4hFE=","miner":"ak_SC7dBmcXyG4i37aKjZM6whqHUU58bsiNths7ivknrG2Z5iF3g","nonce":13382994076605909652,"pow":[1625,2569,3697,6403,6886,7037,7992,9276,9408,10308,10497,10968,12664,12986,13110,13664,13691,14455,14515,18467,19918,20032,20108,20318,20965,22501,23617,23701,23813,24091,24835,24948,24961,26396,26900,27956,28857,29077,30615,30766,31045,32088],"prev_hash":"kh_2HkDM6Bbc3kkPPKWw1psepDH95JFE8LNKKoP9waeh4wWSE1sKh","prev_key_hash":"kh_2HkDM6Bbc3kkPPKWw1psepDH95JFE8LNKKoP9waeh4wWSE1sKh","state_hash":"bs_3sa2GSmv8RGaZfcnaQmjyjFshoBrs68VuUf874RzqCNob1huX","target":539127532,"time":1562686597558,"version":3}}`,
				},
				args: []string{},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Online Test",
			args: args{
				conn: newAeNode(),
				args: []string{},
			},
			wantErr: false,
			online:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !online && tt.online {
				t.Skip("Skipping online test")
			}
			if err := topFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("topFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_broadcastFunc(t *testing.T) {
	type args struct {
		conn aeternity.PostTransactioner
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal PostTransaction without error", // this looks like it tests nothing, but actually exercises hashing/cross checking code. the mock pretends that the node checked the hash and found that it matched.
			args: args{
				conn: &mockPostTransactioner{},
				args: []string{"tx_+KgLAfhCuEAPX1l3BdFOcLeduH3PPwPV25mETXZE8IBDe6PGuasSEKJeB/cDDm+kW05Cdp38+mpvVSTTPMx7trL/7qxfUr8IuGD4XhYBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+8wGTcXVlcnkgU3BlY2lmaWNhdGlvbpZyZXNwb25zZSBTcGVjaWZpY2F0aW9uAABkhrXmIPSAAIIB9AHdGxXf"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := broadcastFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("broadcastFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_statusFunc(t *testing.T) {
	type args struct {
		conn aeternity.GetStatuser
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Normal Status",
			args: args{
				conn: &mockGetStatuser{msg: `{"difficulty":21749349,"genesis_key_block_hash":"kh_2v3mQTSSiyTrhPZjtYcwfbwzt4d6SbFZ5LS9Q7qCosu2QR1beh","listening":true,"network_id":"ae_docker","node_revision":"93c2bd73ae273e3068fb16893024030cf49817b5","node_version":"3.1.0","peer_count":0,"pending_transactions_count":0,"protocols":[{"effective_at_height":1,"version":3},{"effective_at_height":0,"version":1}],"solutions":0,"sync_progress":100,"syncing":false}`},
				args: []string{},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Online Status",
			args: args{
				conn: newAeNode(),
				args: []string{},
			},
			wantErr: false,
			online:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !online && tt.online {
				t.Skip("Skipping online test")
			}
			if err := statusFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("statusFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ttlFunc(t *testing.T) {
	type args struct {
		conn aeternity.GetHeighter
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Height is 1337",
			args: args{
				conn: &mockGetHeighter{h: 1337},
				args: []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ttlFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ttlFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_networkIDFunc(t *testing.T) {
	type args struct {
		conn aeternity.GetStatuser
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		online  bool
	}{
		{
			name: "Status from ae_docker node",
			args: args{
				conn: &mockGetStatuser{msg: `{"difficulty":21749349,"genesis_key_block_hash":"kh_2v3mQTSSiyTrhPZjtYcwfbwzt4d6SbFZ5LS9Q7qCosu2QR1beh","listening":true,"network_id":"ae_docker","node_revision":"93c2bd73ae273e3068fb16893024030cf49817b5","node_version":"3.1.0","peer_count":0,"pending_transactions_count":0,"protocols":[{"effective_at_height":1,"version":3},{"effective_at_height":0,"version":1}],"solutions":0,"sync_progress":100,"syncing":false}`},
				args: []string{},
			},
			wantErr: false,
			online:  false,
		},
		{
			name: "Online Test",
			args: args{
				conn: newAeNode(),
				args: []string{},
			},
			wantErr: false,
			online:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(online, tt.online)
			if !online && tt.online {
				t.Skip("Skipping online test")
			}
			if err := networkIDFunc(tt.args.conn, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("networkIDFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
