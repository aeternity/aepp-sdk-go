package account

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDerivePathFromSeedAeAccount(t *testing.T) {
	tests := []struct {
		name        string
		seed        string
		path        string
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "Standard Seed, m/0'",
			seed:        "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			path:        "m/0'",
			wantAddress: "ak_2o3zmfG3hMFu4oveTs4VcsmdimnsUBX7sEp3LwXhj9kutRm8mN",
			wantErr:     false,
		},
		{
			name:        "Standard Seed, m/44'/457'/0'/0'/0'",
			seed:        "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			path:        "m/44'/457'/0'/0'/0'",
			wantAddress: "ak_2Z74Jhbo3xqF47k2h6NoUpr5gTfc9EQFX7wPH2Vf7Q5PCVcZSW",
			wantErr:     false,
		},
		{
			name:        "Standard Seed, m/44'/457'/0'/0'/3'",
			seed:        "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			path:        "m/44'/457'/0'/0'/3'",
			wantAddress: "ak_2wPpjbxDhdn8PURqLPsunqBTWYSbe9iac1gjfJcQVzY4aZYEzq",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		seedBytes, err := hex.DecodeString(tt.seed)
		if err != nil {
			t.Fatal(err)
		}
		key, err := DerivePathFromSeed(seedBytes, tt.path)
		if err != nil {
			t.Fatal(err)
		}
		acc, err := BIP32KeyToAeKey(key)
		if err != nil {
			t.Fatal(err)
		}

		if acc.Address != tt.wantAddress {
			t.Errorf("%s should have been %s but got %s", tt.seed, tt.wantAddress, acc.Address)
		}
	}
}

func TestDerivePathFromSeed_SLIP0010_TestVectors(t *testing.T) {
	// BIP32 adds an extra 0 byte of padding in the beginning of the public key.
	// We do not use the padded public key. Instead we just feed the private key
	// into the official ed25519 scheme and use its public key. Nevertheless it
	// may be useful to verify the implementation with the test vectors in
	// SLIP0010.
	tests := []struct {
		name          string
		seed          string
		path          string
		wantChainCode string
		wantPrivate   string
		wantPublic    string
		wantErr       bool
	}{
		{
			name:          "Test Vector 1 m",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m",
			wantChainCode: "90046a93de5380a72b5e45010748567d5ea02bbf6522f979e05c0d8d8ca9fffb",
			wantPrivate:   "2b4be7f19ee27bbf30c667b642d5f4aa69fd169872f8fc3059c08ebae2eb19e7",
			wantPublic:    "00a4b2856bfec510abab89753fac1ac0e1112364e7d250545963f135f2a33188ed",
			wantErr:       false,
		},
		{
			name:          "Test Vector 1 m/0'",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m/0'",
			wantChainCode: "8b59aa11380b624e81507a27fedda59fea6d0b779a778918a2fd3590e16e9c69",
			wantPrivate:   "68e0fe46dfb67e368c75379acec591dad19df3cde26e63b93a8e704f1dade7a3",
			wantPublic:    "008c8a13df77a28f3445213a0f432fde644acaa215fc72dcdf300d5efaa85d350c",
			wantErr:       false,
		},
		{
			name:          "Test Vector 1 m/0'/1'",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m/0'/1'",
			wantChainCode: "a320425f77d1b5c2505a6b1b27382b37368ee640e3557c315416801243552f14",
			wantPrivate:   "b1d0bad404bf35da785a64ca1ac54b2617211d2777696fbffaf208f746ae84f2",
			wantPublic:    "001932a5270f335bed617d5b935c80aedb1a35bd9fc1e31acafd5372c30f5c1187",
			wantErr:       false,
		},
		{
			name:          "Test Vector 1 m/0'/1'/2'",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m/0'/1'/2'",
			wantChainCode: "2e69929e00b5ab250f49c3fb1c12f252de4fed2c1db88387094a0f8c4c9ccd6c",
			wantPrivate:   "92a5b23c0b8a99e37d07df3fb9966917f5d06e02ddbd909c7e184371463e9fc9",
			wantPublic:    "00ae98736566d30ed0e9d2f4486a64bc95740d89c7db33f52121f8ea8f76ff0fc1",
			wantErr:       false,
		},
		{
			name:          "Test Vector 1 m/0'/1'/2'/2'",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m/0'/1'/2'/2'",
			wantChainCode: "8f6d87f93d750e0efccda017d662a1b31a266e4a6f5993b15f5c1f07f74dd5cc",
			wantPrivate:   "30d1dc7e5fc04c31219ab25a27ae00b50f6fd66622f6e9c913253d6511d1e662",
			wantPublic:    "008abae2d66361c879b900d204ad2cc4984fa2aa344dd7ddc46007329ac76c429c",
			wantErr:       false,
		},
		{
			name:          "Test Vector 1 m/0'/1'/2'/2'/1000000000'",
			seed:          "000102030405060708090a0b0c0d0e0f",
			path:          "m/0'/1'/2'/2'/1000000000'",
			wantChainCode: "68789923a0cac2cd5a29172a475fe9e0fb14cd6adb5ad98a3fa70333e7afa230",
			wantPrivate:   "8f94d394a8e8fd6b1bc2f3f49f5c47e385281d5c17e65324b0f62483e37e8793",
			wantPublic:    "003c24da049451555d51a7014a37337aa4e12d41e485abccfa46b47dfb2af54b7a",
			wantErr:       false,
		}, {
			name:          "Test Vector 2 m",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m",
			wantChainCode: "ef70a74db9c3a5af931b5fe73ed8e1a53464133654fd55e7a66f8570b8e33c3b",
			wantPrivate:   "171cb88b1b3c1db25add599712e36245d75bc65a1a5c9e18d76f9f2b1eab4012",
			wantPublic:    "008fe9693f8fa62a4305a140b9764c5ee01e455963744fe18204b4fb948249308a",
			wantErr:       false,
		},
		{
			name:          "Test Vector 2 m/0'",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m/0'",
			wantChainCode: "0b78a3226f915c082bf118f83618a618ab6dec793752624cbeb622acb562862d",
			wantPrivate:   "1559eb2bbec5790b0c65d8693e4d0875b1747f4970ae8b650486ed7470845635",
			wantPublic:    "0086fab68dcb57aa196c77c5f264f215a112c22a912c10d123b0d03c3c28ef1037",
			wantErr:       false,
		},
		{
			name:          "Test Vector 2 m/0'/2147483647'",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m/0'/2147483647'",
			wantChainCode: "138f0b2551bcafeca6ff2aa88ba8ed0ed8de070841f0c4ef0165df8181eaad7f",
			wantPrivate:   "ea4f5bfe8694d8bb74b7b59404632fd5968b774ed545e810de9c32a4fb4192f4",
			wantPublic:    "005ba3b9ac6e90e83effcd25ac4e58a1365a9e35a3d3ae5eb07b9e4d90bcf7506d",
			wantErr:       false,
		},
		{
			name:          "Test Vector 2 m/0'/2147483647'/1'",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m/0'/2147483647'/1'",
			wantChainCode: "73bd9fff1cfbde33a1b846c27085f711c0fe2d66fd32e139d3ebc28e5a4a6b90",
			wantPrivate:   "3757c7577170179c7868353ada796c839135b3d30554bbb74a4b1e4a5a58505c",
			wantPublic:    "002e66aa57069c86cc18249aecf5cb5a9cebbfd6fadeab056254763874a9352b45",
			wantErr:       false,
		},
		{
			name:          "Test Vector 2 m/0'/2147483647'/1'/2147483646'",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m/0'/2147483647'/1'/2147483646'",
			wantChainCode: "0902fe8a29f9140480a00ef244bd183e8a13288e4412d8389d140aac1794825a",
			wantPrivate:   "5837736c89570de861ebc173b1086da4f505d4adb387c6a1b1342d5e4ac9ec72",
			wantPublic:    "00e33c0f7d81d843c572275f287498e8d408654fdf0d1e065b84e2e6f157aab09b",
			wantErr:       false,
		},
		{
			name:          "Test Vector 2 m/0'/2147483647'/1'/2147483646'/2'",
			seed:          "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			path:          "m/0'/2147483647'/1'/2147483646'/2'",
			wantChainCode: "5d70af781f3a37b829f0d060924d5e960bdc02e85423494afc0b1a41bbe196d4",
			wantPrivate:   "551d333177df541ad876a60ea71f00447931c0a9da16f227c11ea080d7391b8d",
			wantPublic:    "0047150c75db263559a70d5778bf36abbab30fb061ad69f69ece61a72b0cfa4fc0",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		seedBytes, err := hex.DecodeString(tt.seed)
		if err != nil {
			t.Fatal(err)
		}
		key, err := DerivePathFromSeed(seedBytes, tt.path)
		if err != nil {
			t.Fatal(err)
		}

		if hexify(key.Key) != tt.wantPrivate {
			t.Errorf("%s private key should have been %s but got %s", tt.name, tt.wantPrivate, hexify(key.Key))
		}
		if hexify(key.ChainCode) != tt.wantChainCode {
			t.Errorf("%s chaincode should have been %s but got %s", tt.name, tt.wantChainCode, hexify(key.ChainCode))
		}

		k := ed25519.NewKeyFromSeed(key.Key)
		kPub := []byte(k.Public().(ed25519.PublicKey))
		kPubHex := fmt.Sprintf("00%s", hexify(kPub))
		if kPubHex != tt.wantPublic {
			t.Errorf("%s public key should have been %s but got %s", tt.name, tt.wantPublic, kPubHex)
		}
	}
}

func TestBIP32KeyToAEKey(t *testing.T) {
	keySeedHex := "bb220e6d0c1c8fdeb523dc258b30be85def7aa03c0c499fb3639bae2b5a342ed"
	expectedAddress := "ak_GL2DaDKGMV6K9QUKM9s1NheKvGjsmqRSCN4e8HRcJ5uwvzxH3"
	keySeed, _ := hex.DecodeString(keySeedHex)
	key, err := NewMasterKey(keySeed)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(key.Key))

	acc, err := BIP32KeyToAeKey(key)
	if err != nil {
		t.Fatal(err)
	}

	if acc.Address != expectedAddress {
		t.Errorf("BIP32 Key Seed %s should have become the aeternity account %s, got %s", keySeedHex, expectedAddress, acc.Address)
	}
}

func TestParseMnemonic(t *testing.T) {
	type args struct {
		mnemonic string
	}
	tests := []struct {
		name           string
		args           args
		wantMasterSeed string
		wantErr        bool
	}{
		{
			name: "ring defense obey exhaust boat popular surround supreme front lemon monster number",
			args: args{
				mnemonic: "ring defense obey exhaust boat popular surround supreme front lemon monster number",
			},
			wantMasterSeed: "60812c7c93d6f9cb346bbcf799957b6ec776aea84b01bdd9f9b7916522cc52c6ea5d07960b68668cd37b0a77f6c4fe283f146bd916153c426df126a8b8707b39",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMasterSeed, err := ParseMnemonic(tt.args.mnemonic)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotMasterSeedHex := hexify(gotMasterSeed)
			if gotMasterSeedHex != tt.wantMasterSeed {
				t.Errorf("ParseMnemonic() = %v, want %v", gotMasterSeedHex, tt.wantMasterSeed)
			}
		})
	}
}
