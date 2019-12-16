package aeternity

import (
	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type transactionSender interface {
	naet.Infoer                              // quickly get node version, networkID for VM/ABI
	naet.GetAccounter                        // for transactions.NewTTLNoncer()
	broadcastWaitTransactionNodeCapabilities // basic transaction broadcasting capabilities
}

type compileencoder interface {
	naet.CompileContracter
	naet.EncodeCalldataer
}

type ContextInterface interface {
	SenderAccount() string
	TTLNoncer() transactions.TTLNoncer
	Compiler() compileencoder
	NodeInfo() (networkID string, version string)
	SignBroadcastWait(tx transactions.Transaction, blocks uint64) (*TxReceipt, error)
	SetCompiler(c compileencoder)
}

// Context holds information and node capabilities needed to create, sign and
// send transactions to a node. The node connection in Context does not need to
// be capable of other feature specific Swagger API endpoints.
type Context struct {
	SigningAccount *account.Account
	ttlNoncer      transactions.TTLNoncer
	compiler       compileencoder
	txSender       transactionSender
}

// NewContext creates a new Context, but does not force one to provide a
// compiler (which can be set via SetCompiler)
func NewContext(signingAccount *account.Account, node transactionSender) (b *Context, err error) {
	return &Context{
		SigningAccount: signingAccount,
		ttlNoncer:      transactions.NewTTLNoncer(node),
		txSender:       node,
	}, nil
}

func (c *Context) SenderAccount() string {
	return c.SigningAccount.Address
}

func (c *Context) TTLNoncer() transactions.TTLNoncer {
	return c.ttlNoncer
}

func (c *Context) Compiler() compileencoder {
	return c.compiler
}

func (c *Context) NodeInfo() (networkID string, version string) {
	return c.txSender.Info()
}

func (c *Context) SignBroadcastWait(tx transactions.Transaction, blocks uint64) (txReceipt *TxReceipt, err error) {
	networkID, _ := c.txSender.Info()
	return SignBroadcastWaitTransaction(tx, c.SigningAccount, c.txSender, networkID, blocks)
}

func (c *Context) SetCompiler(compiler compileencoder) {
	c.compiler = compiler
}
