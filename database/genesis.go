package database

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

var genesisJson = `{
  "genesis_time": "2020-06-01T00:00:00.000000000Z",
  "chain_id": "the-blockchain-bar-ledger",
  "symbol": "TBB",
  "balances": {
    "0xCDde88A0E6E4598952aEf7a1340b93C9033e2458": 1000000
  },
  "fork_tip_1": 35
}`

type Genesis struct {
	Balances map[common.Address]uint `json:"balances"`
	Symbol   string                  `json:"symbol"`

	ForkTIP1 uint64 `json:"fork_tip_1"`
}

func loadGenesis(path string) (Genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}

func writeGenesisToDisk(path string, genesis []byte) error {
	return ioutil.WriteFile(path, genesis, 0644)
}
