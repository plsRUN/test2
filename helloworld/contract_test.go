package helloworld

import (
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
const testGreeting = "test"

var tests = map[string]testutils.PrecompileTest{
	"calling sayHello from NoRole should succeed": {
		Caller:     allowlist.TestNoRoleAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: func() []byte {
			// by default we don't Configure initial state for
			// the module since Config is empty.
			// This means we don't apply default greeting to the state.
			res, err := PackSayHelloOutput("")
			if err != nil {
				panic(err)
			}
			return res
		}(),
		SuppliedGas: SayHelloGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
	},
	"calling sayHello from Enabled should succeed": {
		Caller:     allowlist.TestEnabledAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: func() []byte {
			// by default we don't Configure initial state for
			// the module since Config is empty.
			// This means we don't apply default greeting to the state.
			res, err := PackSayHelloOutput("")
			if err != nil {
				panic(err)
			}
			return res
		}(),
		SuppliedGas: SayHelloGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
	},
	"calling sayHello from Admin should succeed": {
		Caller:     allowlist.TestAdminAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: func() []byte {
			// by default we don't Configure initial state for
			// the module since Config is empty.
			// This means we don't apply default greeting to the state.
			res, err := PackSayHelloOutput("")
			if err != nil {
				panic(err)
			}
			return res
		}(),
		SuppliedGas: SayHelloGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
	},
	"calling sayHello from NoRole with a config should return default greeting": {
		Caller:     allowlist.TestNoRoleAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		Config:     NewConfig(common.Big0, nil, nil),
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: func() []byte {
			res, err := PackSayHelloOutput(defaultGreeting)
			if err != nil {
				panic(err)
			}
			return res
		}(),
		SuppliedGas: SayHelloGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
	},
	"insufficient gas for sayHello should fail": {
		Caller: common.Address{1},
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		SuppliedGas: SayHelloGasCost - 1,
		ReadOnly:    false,
		ExpectedErr: vmerrs.ErrOutOfGas.Error(),
	},
	"calling setGreeting from NoRole should fail": {
		Caller:     allowlist.TestNoRoleAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			// CUSTOM CODE STARTS HERE
			// set test input to a value here
			input, err := PackSetGreeting(testGreeting)
			require.NoError(t, err)
			return input
		},
		SuppliedGas: SetGreetingGasCost,
		ReadOnly:    false,
		ExpectedErr: ErrCannotSetGreeting.Error(),
	},
	"calling setGreeting from Enabled should succeed": {
		Caller:     allowlist.TestEnabledAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			// CUSTOM CODE STARTS HERE
			// set test input to a value here
			input, err := PackSetGreeting(testGreeting)
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: []byte{},
		SuppliedGas: SetGreetingGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
		AfterHook: func(t testing.TB, state contract.StateDB) {
			greeting := GetGreeting(state)
			require.Equal(t, greeting, testGreeting)
		},
	},
	"calling setGreeting from Admin should succeed": {
		Caller:     allowlist.TestAdminAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			// CUSTOM CODE STARTS HERE
			// set test input to a value here
			input, err := PackSetGreeting(testGreeting)
			require.NoError(t, err)
			return input
		},
		// This test is for a successful call. You can set the expected output here.
		// CUSTOM CODE STARTS HERE
		ExpectedRes: []byte{},
		SuppliedGas: SetGreetingGasCost,
		ReadOnly:    false,
		ExpectedErr: "",
		AfterHook: func(t testing.TB, state contract.StateDB) {
			greeting := GetGreeting(state)
			require.Equal(t, greeting, testGreeting)
		},
	},
	"readOnly setGreeting should fail": {
		Caller: common.Address{1},
		InputFn: func(t testing.TB) []byte {
			// CUSTOM CODE STARTS HERE
			// set test input to a value here
			var testInput string
			input, err := PackSetGreeting(testInput)
			require.NoError(t, err)
			return input
		},
		SuppliedGas: SetGreetingGasCost,
		ReadOnly:    true,
		ExpectedErr: vmerrs.ErrWriteProtection.Error(),
	},
	"insufficient gas for setGreeting should fail": {
		Caller: common.Address{1},
		InputFn: func(t testing.TB) []byte {
			// CUSTOM CODE STARTS HERE
			// set test input to a value here
			var testInput string
			input, err := PackSetGreeting(testInput)
			require.NoError(t, err)
			return input
		},
		SuppliedGas: SetGreetingGasCost - 1,
		ReadOnly:    false,
		ExpectedErr: vmerrs.ErrOutOfGas.Error(),
	},
	// more custom tests
	"store greeting then say hello from non-enabled address": {
		Caller: allowlist.TestNoRoleAddr,
		BeforeHook: func(t testing.TB, state contract.StateDB) {
			allowlist.SetDefaultRoles(Module.Address)(t, state)
			StoreGreeting(state, testGreeting)
		},
		InputFn: func(t testing.TB) []byte {
			input, err := PackSayHello()
			require.NoError(t, err)
			return input
		},
		SuppliedGas: SayHelloGasCost,
		ReadOnly:    true,
		ExpectedRes: func() []byte {
			res, err := PackSayHelloOutput(testGreeting)
			if err != nil {
				panic(err)
			}
			return res
		}(),
	},
	"set a very long greeting from enabled address": {
		Caller:     allowlist.TestEnabledAddr,
		BeforeHook: allowlist.SetDefaultRoles(Module.Address),
		InputFn: func(t testing.TB) []byte {
			longString := "a very long string that is longer than 32 bytes and will cause an error"
			input, err := PackSetGreeting(longString)
			require.NoError(t, err)

			return input
		},
		SuppliedGas: SetGreetingGasCost,
		ReadOnly:    false,
		ExpectedErr: ErrInputExceedsLimit.Error(),
	},
}

// TestHelloWorldRun tests the Run function of the precompile contract.
func TestHelloWorldRun(t *testing.T) {
	// Run tests with allowlist tests.
	// This adds allowlist run tests to your custom tests
	// and runs them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.RunPrecompileWithAllowListTests(t, Module, state.NewTestStateDB, tests)
}

func BenchmarkHelloWorld(b *testing.B) {
	// Benchmark tests with allowlist tests.
	// This adds allowlist run tests to your custom tests
	// and benchmarks them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.BenchPrecompileWithAllowList(b, Module, state.NewTestStateDB, tests)
}
