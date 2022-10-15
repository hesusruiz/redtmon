package redt

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type StatisticsRedT struct {
	allValidators      map[common.Address]*NodeInfo
	valSet             []common.Address
	asProposer         map[common.Address]int
	asSigner           map[common.Address]int
	lastBlockProcessed int64
	previousTimestamp  uint64
	elapsed            uint64
}

func NewStatistics(allValidators map[common.Address]*NodeInfo, valSet []common.Address) *StatisticsRedT {
	st := &StatisticsRedT{}
	st.allValidators = allValidators
	st.valSet = valSet

	// Initialise the counters for validators/signers
	st.asProposer = map[common.Address]int{}
	st.asSigner = map[common.Address]int{}

	return st
}

func (st *StatisticsRedT) ValidatorSet() []common.Address {
	return st.valSet
}

func (st *StatisticsRedT) ValidatorInfo(validator common.Address) *NodeInfo {
	return st.allValidators[validator]
}

// UpdateStatisticsForBlock uses the info of the block to update the statistics
// It receives a Header (not a full Block) and returns the address of the author and an array with the signers of the block
func (st *StatisticsRedT) UpdateStatisticsForBlock(header *ethertypes.Header) (author common.Address, signers []common.Address, err error) {

	// Get the author (proposer) and signers for this block
	author, signers, err = SignersFromBlock(header)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Do nothing if the block was already processed
	thisBlockNumber := header.Number.Int64()
	if thisBlockNumber <= st.lastBlockProcessed {
		return author, signers, nil
	}

	// Update the last block processed
	st.lastBlockProcessed = thisBlockNumber

	// Calculate elapsed time since last block processed
	st.elapsed = header.Time - st.previousTimestamp

	// And update the last timestamp
	st.previousTimestamp = header.Time

	// Increment the counter for authors
	st.asProposer[author] += 1

	// Increment counters for signers
	for _, seal := range signers {
		// Increment the counter of signatures
		st.asSigner[seal] += 1
	}

	return author, signers, nil

}

// StatisticsForHeader returns a map with statistical data prepared for HTML templates
func (st *StatisticsRedT) StatisticsForHeader(header *ethertypes.Header, latestTimestamp uint64) (map[string]any, uint64) {

	data := make(map[string]any)

	currentTimestamp := header.Time

	// Calculate the elapsed time with respect to the latest one we received
	elapsed := currentTimestamp - latestTimestamp

	// Update the statistics counters
	author, signers, err := st.UpdateStatisticsForBlock(header)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Get the info of the node operator
	oper := st.ValidatorInfo(author)

	// Determine the next node that should be proposer, according to the round-robin
	// selection algorithm
	var nextProposer common.Address
	for i := 0; i < len(st.valSet); i++ {
		if author == st.valSet[i] {
			nextProposer = st.valSet[(i+1)%len(st.valSet)]
			break
		}
	}

	// Get the name of the node operator
	nextProposerOperator := st.ValidatorInfo(nextProposer).Operator

	// Round the timestamp to seconds
	t := time.Unix(int64(currentTimestamp), 0)

	data["number"] = header.Number
	data["elapsed"] = elapsed
	data["timestamp"] = t
	data["operator"] = oper.Operator
	data["authorCount"] = st.asProposer[author]
	data["authorAddress"] = author
	data["nextProposerOperator"] = nextProposerOperator

	data["gasLimit"] = header.GasLimit
	data["gasUsed"] = header.GasUsed

	// Create the map with signers of this block
	var currentSigners = map[common.Address]bool{}
	for _, seal := range signers {
		currentSigners[seal] = true
	}

	sigt := make([]map[string]any, len(st.ValidatorSet()))

	for i, val := range st.ValidatorSet() {

		d := make(map[string]any)

		item := st.ValidatorInfo(val)

		authorCount := st.asProposer[item.Address]

		authorCountStr := fmt.Sprintf("%v", authorCount)
		if authorCount == 0 {
			authorCountStr = fmt.Sprintf("<span class='w3-badge w3-red'>%v</span>", authorCount)
		}
		if author.Hex() == item.Address.Hex() {
			authorCountStr = fmt.Sprintf("<span class='w3-badge'>%v</span>", authorCount)
		}

		d["authorCount"] = authorCountStr

		signerCount := st.asSigner[item.Address]

		signerCountStr := fmt.Sprintf("%v", signerCount)
		if signerCount == 0 {
			signerCountStr = fmt.Sprintf("<span class='w3-badge w3-red'>%v</span>", signerCount)
		}
		if currentSigners[item.Address] {
			signerCountStr = fmt.Sprintf("<span class='w3-badge'>%v</span>", signerCount)
		}

		d["signerCount"] = signerCountStr

		d["operator"] = item.Operator
		d["address"] = item.Address

		sigt[i] = d

	}

	data["signers"] = sigt

	return data, currentTimestamp

}

type StatisticsSummary struct {
	BlockNumber      int64
	Elapsed          uint64
	Timestamp        time.Time
	ProposerName     string
	ProposerCount    int
	ProposerAddress  common.Address
	NextProposerName string
	GasLimit         uint64
	GasUsed          uint64
	Signers          []SignersTable
}

type SignersTable struct {
	NextProposer    bool
	CurrentProposer bool
	AsProposer      int
	CurrentSigner   bool
	AsSigner        int
	Name            string
	Address         common.Address
}

// StatisticsForHeader returns a map with statistical data prepared for HTML templates
func (st *StatisticsRedT) StatisticsForHeaderNew(header *ethertypes.Header) *StatisticsSummary {
	data := &StatisticsSummary{}

	// Update the statistics counters
	author, signers, err := st.UpdateStatisticsForBlock(header)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Get the info of the node operator
	oper := st.ValidatorInfo(author)

	// Determine the next node that should be proposer, according to the round-robin
	// selection algorithm
	var nextProposer common.Address
	for i := 0; i < len(st.valSet); i++ {
		if author == st.valSet[i] {
			nextProposer = st.valSet[(i+1)%len(st.valSet)]
			break
		}
	}

	// Calculate values for the header
	data.BlockNumber = header.Number.Int64()
	data.Elapsed = st.elapsed
	data.Timestamp = time.Unix(int64(header.Time), 0)
	data.ProposerName = oper.Operator
	data.ProposerCount = st.asProposer[author]
	data.ProposerAddress = author
	data.NextProposerName = st.ValidatorInfo(nextProposer).Operator
	data.GasLimit = header.GasLimit
	data.GasUsed = header.GasUsed

	// Create the map with signers of this block
	var currentSigners = map[common.Address]bool{}
	for _, seal := range signers {
		currentSigners[seal] = true
	}

	// Calculate the table of validators
	validatorTable := make([]SignersTable, len(st.ValidatorSet()))

	for i, addr := range st.ValidatorSet() {

		// Mark if this is the next proposer
		if nextProposer.Hex() == addr.Hex() {
			validatorTable[i].NextProposer = true
		}

		// Mark if this is the proposer of this block
		if author.Hex() == addr.Hex() {
			validatorTable[i].CurrentProposer = true
		}

		// Number of times this has been a proposer
		validatorTable[i].AsProposer = st.asProposer[addr]

		// Mark is this is a signer of current block
		if currentSigners[addr] {
			validatorTable[i].CurrentSigner = true
		}

		// Number of times this has been a signer
		validatorTable[i].AsSigner = st.asSigner[addr]

		// Name of the node
		validatorTable[i].Name = st.ValidatorInfo(addr).Operator

		// Address of the node
		validatorTable[i].Address = addr
	}

	data.Signers = validatorTable

	return data

}
