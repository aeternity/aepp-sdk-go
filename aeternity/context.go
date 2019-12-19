package aeternity

import (
	"github.com/aeternity/aepp-sdk-go/v7/account"
	"github.com/aeternity/aepp-sdk-go/v7/naet"
	"github.com/aeternity/aepp-sdk-go/v7/transactions"
)

type transactionSender interface {
	naet.GetStatuser                         // quickly get node version, networkID for VM/ABI
	naet.GetAccounter                        // for transactions.NewTTLNoncer()
	broadcastWaitTransactionNodeCapabilities // basic transaction broadcasting capabilities
}

type broadcastWaitTransactionNodeCapabilities interface {
	naet.PostTransactioner
	getTransactionByHashHeighter
}

type CompileEncoder interface {
	naet.CompileContracter
	naet.EncodeCalldataer
}

// ContextInterface describes the capabilities of Context, which provide
// information solely related to transaction creation/broadcasting. It allows
// for Context to be mocked out.
type ContextInterface interface {
	SenderAccount() string
	TTLNoncer() transactions.TTLNoncer
	Compiler() CompileEncoder
	NodeInfo() (networkID string, version string, err error)
	SignBroadcastWait(tx transactions.Transaction, blocks uint64) (r *TxReceipt, err error)
	SetCompiler(c CompileEncoder)
}

// Context holds information and node capabilities needed to create, sign and
// send transactions to a node. The node connection in Context does not need to
// be capable of other feature specific Swagger API endpoints.
type Context struct {
	SigningAccount *account.Account
	ttlNoncer      transactions.TTLNoncer
	compiler       CompileEncoder
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

// SenderAccount returns the address of the signing account, which should also
// be the sender address (for many transaction types)
func (c *Context) SenderAccount() string {
	return c.SigningAccount.Address
}

// TTLNoncer returns the TTLNoncer of Context.SigningAccount
func (c *Context) TTLNoncer() transactions.TTLNoncer {
	return c.ttlNoncer
}

// Compiler returns the compiler interface
func (c *Context) Compiler() CompileEncoder {
	return c.compiler
}

// NodeInfo returns the networkID and version of the currently connected node,
// needed for contract Tx creation
func (c *Context) NodeInfo() (networkID string, version string, err error) {
	s, err := c.txSender.GetStatus()
	if err != nil {
		return
	}
	return *s.NetworkID, *s.NodeVersion, err
}

// SignBroadcastWait signs, sends and waits for the transaction to be mined.
func (c *Context) SignBroadcastWait(tx transactions.Transaction, blocks uint64) (txReceipt *TxReceipt, err error) {
	networkID, _, err := c.NodeInfo()
	if err != nil {
		return
	}
	txReceipt, err = SignBroadcast(tx, c.SigningAccount, c.txSender, networkID)
	if err != nil {
		return
	}

	err = WaitSynchronous(txReceipt, blocks, c.txSender)
	return txReceipt, err
}

// SetCompiler changes the Context's compiler instance.
func (c *Context) SetCompiler(compiler CompileEncoder) {
	c.compiler = compiler
}