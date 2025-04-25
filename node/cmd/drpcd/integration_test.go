package main

import (
	"math/big"
	"testing"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// deployContracts is a placeholder: implement your deploy logic using ABI/bin
func deployContracts(t *testing.T, sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, common.Address) {
	t.Skip("integration test scaffold: implement Deploy with your ABI/bin files")
	return common.Address{}, common.Address{}
}

func TestRegisterProviderIntegration(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1e18)}
	sim := backends.NewSimulatedBackend(alloc, 8000000)

	chanAddr, regAddr := deployContracts(t, sim, auth)
	sim.Commit()

	pc, err := contracts.NewPaymentChannel(chanAddr, sim)
	if err != nil {
		t.Fatal(err)
	}
	pr, err := contracts.NewProviderRegistry(regAddr, sim)
	if err != nil {
		t.Fatal(err)
	}

	srv := &RPCServer{
		registry:       pr,
		paymentChannel: pc,
		providerReg:    pr,
		privateKey:     key,
		config:         Config{},
	}

	if err := srv.registerProvider(); err != nil {
		t.Fatal(err)
	}

	info, err := pr.Providers(&bind.CallOpts{}, auth.From)
	if err != nil {
		t.Fatal(err)
	}
	if !info.Registered {
		t.Error("expected provider to be registered on chain")
	}
}
