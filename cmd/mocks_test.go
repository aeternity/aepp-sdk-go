package cmd

import (
	compilermodels "github.com/aeternity/aepp-sdk-go/v8/swagguard/compiler/models"
	models "github.com/aeternity/aepp-sdk-go/v8/swagguard/node/models"
)

type mockGetHeighter struct {
	h uint64
}

func (m *mockGetHeighter) GetHeight() (uint64, error) {
	return m.h, nil
}

type mockGetTopBlocker struct {
	msg string
}

func (m *mockGetTopBlocker) GetTopBlock() (*models.KeyBlockOrMicroBlockHeader, error) {
	kb := &models.KeyBlockOrMicroBlockHeader{}
	kb.UnmarshalBinary([]byte(m.msg))
	return kb, nil
}

type mockPostTransactioner struct{}

func (m *mockPostTransactioner) PostTransaction(signedEncodedTx, signedEncodedTxHash string) error {
	return nil
}

type mockGetStatuser struct {
	msg string
}

func (m *mockGetStatuser) GetStatus() (*models.Status, error) {
	kb := &models.Status{}
	kb.UnmarshalBinary([]byte(m.msg))
	return kb, nil
}

type mockGetAccounter struct {
	account string
}

func (m *mockGetAccounter) GetAccount(accountID string) (acc *models.Account, err error) {
	acc = &models.Account{}
	err = acc.UnmarshalBinary([]byte(m.account))
	return acc, err
}

type mockGetKeyBlockByHasher struct {
	keyBlock string
}

func (m *mockGetKeyBlockByHasher) GetKeyBlockByHash(keyBlockID string) (kb *models.KeyBlock, err error) {
	kb = &models.KeyBlock{}
	err = kb.UnmarshalBinary([]byte(m.keyBlock))
	return kb, err
}

type mockGetMicroBlockHeaderTransactions struct {
	mbHeader string
	mbTxs    string
}

func (m *mockGetMicroBlockHeaderTransactions) GetMicroBlockHeaderByHash(mbHash string) (mbHeader *models.MicroBlockHeader, err error) {
	mbHeader = &models.MicroBlockHeader{}
	err = mbHeader.UnmarshalBinary([]byte(m.mbHeader))
	return mbHeader, err
}
func (m *mockGetMicroBlockHeaderTransactions) GetMicroBlockTransactionsByHash(mbHash string) (txs *models.GenericTxs, err error) {
	txs = &models.GenericTxs{}
	err = txs.UnmarshalBinary([]byte(m.mbTxs))
	return txs, err
}

type mockGetTransactionByHasher struct {
	transaction string
}

func (m *mockGetTransactionByHasher) GetTransactionByHash(txHash string) (tx *models.GenericSignedTx, err error) {
	tx = &models.GenericSignedTx{}
	err = tx.UnmarshalBinary([]byte(m.transaction))
	return tx, err
}

type mockGetOracleByPubkeyer struct {
	oracle string
}

func (m *mockGetOracleByPubkeyer) GetOracleByPubkey(oracleID string) (o *models.RegisteredOracle, err error) {
	o = &models.RegisteredOracle{}
	err = o.UnmarshalBinary([]byte(m.oracle))
	return o, err
}

type mockGetNameEntryByNamer struct {
	nameEntry string
}

func (m *mockGetNameEntryByNamer) GetNameEntryByName(name string) (nameEntry *models.NameEntry, err error) {
	nameEntry = &models.NameEntry{}
	err = nameEntry.UnmarshalBinary([]byte(m.nameEntry))
	return nameEntry, err
}

type mockCompileContracter struct{}

func (m *mockCompileContracter) CompileContract(source string) (bytecode string, err error) {
	return "cb_+QYYRgKg+HOI9x+n5+MOEpnQ/zO+GoibqhQxGO4bgnvASx0vzB75BKX5AUmgOoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugeDc2V0uMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoP//////////////////////////////////////////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP///////////////////////////////////////////jJoEnsSQdsAgNxJqQzA+rc5DsuLDKUV7ETxQp+ItyJgJS3g2dldLhgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////////////////////////uEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+QKLoOIjHWzfyTkW3kyzqYV79lz0D8JW9KFJiz9+fJgMGZNEhGluaXS4wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACg//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEA//////////////////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYD//////////////////////////////////////////wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQFEYgAAj2IAAMKRgICAUX9J7EkHbAIDcSakMwPq3OQ7LiwylFexE8UKfiLciYCUtxRiAAE5V1CAgFF/4iMdbN/JORbeTLOphXv2XPQPwlb0oUmLP358mAwZk0QUYgAA0VdQgFF/OoWULXtHOgf10E7h2cFqXOqxa3kc6pKJYRpEw/nlugcUYgABG1dQYAEZUQBbYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUpBZYABRWVJgAFJgAPNbYACAUmAA81tgAFFRkFZbYCABUVGQUIOSUICRUFCAWZCBUllgIAGQgVJgIJADYAAZWWAgAZCBUmAgkANgAFmQgVKBUllgIAGQgVJgIJADYAOBUoFSkFCQVltgIAFRUVlQgJFQUGAAUYFZkIFSkFBgAFJZkFCQVltQUFlQUGIAAMpWhTMuMS4wHchc+w==", nil
}

type mockEncodeCalldataer struct{}

func (m *mockEncodeCalldataer) EncodeCalldata(source string, function string, args []string) (bytecode string, err error) {
	return "cb_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACDiIx1s38k5Ft5Ms6mFe/Zc9A/CVvShSYs/fnyYDBmTRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACo7j+li", nil
}

type mockdecodeCalldataer struct {
	decodedCalldata string
}

func (m *mockdecodeCalldataer) DecodeCalldataSource(source string, function string, callData string) (decodedCallData *compilermodels.DecodedCalldata, err error) {
	decodedCallData = &compilermodels.DecodedCalldata{}
	decodedCallData.UnmarshalBinary([]byte(m.decodedCalldata))
	return decodedCallData, nil
}
func (m *mockdecodeCalldataer) DecodeCalldataBytecode(bytecode string, calldata string) (decodedCallData *compilermodels.DecodedCalldata, err error) {
	decodedCallData = &compilermodels.DecodedCalldata{}
	decodedCallData.UnmarshalBinary([]byte(m.decodedCalldata))
	return decodedCallData, nil
}

type mockGenerateACIer struct {
	aci string
}

func (m *mockGenerateACIer) GenerateACI(source string) (aci *compilermodels.ACI, err error) {
	aci = &compilermodels.ACI{}
	err = aci.UnmarshalBinary([]byte(m.aci))
	return aci, err
}
