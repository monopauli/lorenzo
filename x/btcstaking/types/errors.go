package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/btcstaking module sentinel errors
var (
	ErrParseBTCTx                  = errorsmod.Register(ModuleName, 1100, "cank't parse btc transaction")
	ErrDupBTCTx                    = errorsmod.Register(ModuleName, 1101, "duplicate btc transaction")
	ErrBlkHdrNotFound              = errorsmod.Register(ModuleName, 1102, "btc block header not found")
	ErrBlkHdrNotConfirmed          = errorsmod.Register(ModuleName, 1103, "btc block header not confirmed yet")
	ErrBTCTxNotIncluded            = errorsmod.Register(ModuleName, 1104, "btc transaction not included in the Bitcoin chain")
	ErrInvalidReceivingAddr        = errorsmod.Register(ModuleName, 1105, "invalid btc receiving address")
	ErrInvalidTransaction          = errorsmod.Register(ModuleName, 1106, "invalid transaction")
	ErrMintToModule                = errorsmod.Register(ModuleName, 1107, "fail to mint to module")
	ErrMintToAddr                  = errorsmod.Register(ModuleName, 1108, "invalid mint-to address")
	ErrTransferToAddr              = errorsmod.Register(ModuleName, 1109, "fail to transfer to mint-to address")
	ErrRecordStaking               = errorsmod.Register(ModuleName, 1110, "fail to record staking transaction")
	ErrInvalidBurnBtcTargetAddress = errorsmod.Register(ModuleName, 1111, "invalid btc target address in MsgBurnRequest")
	ErrBurn                        = errorsmod.Register(ModuleName, 1112, "fail to burn")
	ErrEmitEvent                   = errorsmod.Register(ModuleName, 1113, "fail to emit a event")
	ErrBurnAmountLeDust            = errorsmod.Register(ModuleName, 1114, "amount less or equal btc dust threshold")
	ErrBurnInsufficientBalance     = errorsmod.Register(ModuleName, 1115, "insufficient account balance")
	ErrBurnInvalidSigner           = errorsmod.Register(ModuleName, 1116, "invalid signer")
)
