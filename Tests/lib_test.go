package libtest

import (
	"testing"
	// TODO set this back to "github.com/user/project/Tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts-Template-Rust/Tests/testutils"
	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/packages/solo"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestLib(t *testing.T) {
	// Contract name - Defined in SmartContract/Cargo.toml > package > name
	const contractName = "my_iota_sc"
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, contractName)

	// Name of the SC function to be requested - Defined in lib.rs > add_call > my_sc_request
	functionName := "my_sc_request"

	env := solo.New(t, false, false)
	chainName := contractName + "Chain"
	chain := env.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Defines which contract and function will be called by chain.PostRequest
	req := solo.NewCallParams(contractName, functionName)

	// Calls contract my_iota_sc, function my_sc_request
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}
