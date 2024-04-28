package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/YieldAccruingToken.json
	YieldAccruingTokenJSON []byte //nolint: golint

	YieldAccruingTokenContract evmtypes.CompiledContract
)

func init() {

	// contract code
	// https://github.com/Lorenzo-Protocol/builtin-contracts/blob/main/contracts/YieldAccruingToken.sol

	err := json.Unmarshal(YieldAccruingTokenJSON, &YieldAccruingTokenContract)
	if err != nil {
		panic(err)
	}

	if len(YieldAccruingTokenContract.Bin) == 0 {
		panic("load contract failed")
	}
}
