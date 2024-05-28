// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOracleTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IOracleTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IOracleTaskManagerTask struct {
	DolarDatetime             *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IOracleTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOracleTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	DolarDatetime      *big.Int
}

// IOracleTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOracleTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOracleTaskManagerMetaData contains all meta data concerning the ContractOracleTaskManager contract.
var ContractOracleTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"dolarDatetime\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOracleTaskManager.Task\",\"components\":[{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOracleTaskManager.Task\",\"components\":[{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOracleTaskManager.Task\",\"components\":[{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOracleTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dolarDatetime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOracleTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005b1f38038062005b1f8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615828620002f7600039600081816102720152818161057b015261189901526000818161054401526127b00152600081816103f301528181611fc0015261299a01526000818161041a01528181612b700152612d3201526000818161045401528181610dea015281816124ad01528181612626015261285401526158286000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636b532e9e116101255780639d3a0f2d116100ad578063f2fde38b1161007c578063f2fde38b14610566578063f5c9899d14610579578063f63c5bab1461059f578063f8c8765e146105a7578063fabc1cbc146105ba57600080fd5b80639d3a0f2d146104fe578063b98d090814610511578063cefdc1d41461051e578063df5cf7231461053f57600080fd5b806372d18e8d116100f457806372d18e8d1461049f5780637afa1eed146104b2578063886f1195146104c55780638b00ce7c146104d85780638da5cb5b146104ed57600080fd5b80636b532e9e1461043c5780636d14a9871461044f5780636efb463614610476578063715018a61461049757600080fd5b8063416c7e5e116101a85780635baec9a0116101775780635baec9a0146103b05780635c975abb146103c35780635decc3f5146103cb5780635df45946146103ee578063683048351461041557600080fd5b8063416c7e5e146103425780634f739f7414610355578063595c6a67146103755780635ac86ab71461037d57600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b6102296102243660046143fc565b6105cd565b005b610229610239366004614419565b610689565b61025161024c366004614597565b6107c8565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b609b546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e2366004614605565b60996020526000908152604090205481565b604051908152602001610264565b6102f4610310366004614605565b60986020526000908152604090205481565b610335610330366004614622565b610952565b604051610264919061477d565b6102296103503660046147a5565b610de8565b61036861036336600461480a565b610f5d565b604051610264919061490e565b6102296115e1565b6103a061038b3660046149d8565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6102296103be366004614cc3565b6116a8565b6066546102f4565b6103a06103d9366004614605565b609a6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61022961044a366004614d37565b611b27565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b610489610484366004614dbd565b6120f9565b604051610264929190614e7d565b610229612fe7565b609754610100900463ffffffff16610294565b609c546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60975461029490610100900463ffffffff1681565b6033546001600160a01b03166102bc565b61022961050c366004614ec6565b612ffb565b6097546103a09060ff1681565b61053161052c366004614f2a565b6131b3565b604051610264929190614f6c565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102296105743660046143fc565b613345565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b5366004614f8d565b6133bb565b6102296105c8366004614419565b61350c565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610620573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106449190614fe9565b6001600160a01b0316336001600160a01b03161461067d5760405162461bcd60e51b815260040161067490615006565b60405180910390fd5b61068681613668565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f59190615050565b6107115760405162461bcd60e51b81526004016106749061506d565b6066548181161461078a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610810576108106150b5565b60200201518951600160200201518a60200151600060028110610835576108356150b5565b60200201518b60200151600160028110610851576108516150b5565b602090810291909101518c518d8301516040516108ae9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d191906150cb565b90506109446108ea6108e3888461375f565b86906137f6565b6108f261388a565b61093a61092b85610925604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061375f565b6109348c61394a565b906137f6565b886201d4c06139da565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610994573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b89190614fe9565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1e9190614fe9565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a849190614fe9565b9050600086516001600160401b03811115610aa157610aa1614432565b604051908082528060200260200182016040528015610ad457816020015b6060815260200190600190039081610abf5790505b50905060005b8751811015610ddc576000888281518110610af757610af76150b5565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b8091908101906150ed565b905080516001600160401b03811115610b9b57610b9b614432565b604051908082528060200260200182016040528015610be657816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bb95790505b50848481518110610bf957610bf96150b5565b602002602001018190525060005b8151811015610dc6576040518060600160405280876001600160a01b03166347b314e8858581518110610c3c57610c3c6150b5565b60200260200101516040518263ffffffff1660e01b8152600401610c6291815260200190565b602060405180830381865afa158015610c7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca39190614fe9565b6001600160a01b03168152602001838381518110610cc357610cc36150b5565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cf157610cf16150b5565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d71919061517d565b6001600160601b0316815250858581518110610d8f57610d8f6150b5565b60200260200101518281518110610da857610da86150b5565b60200260200101819052508080610dbe906151bc565b915050610c07565b5050508080610dd4906151bc565b915050610ada565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6a9190614fe9565b6001600160a01b0316336001600160a01b031614610f165760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610674565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610f886040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fec9190614fe9565b90506110196040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611049908b90899089906004016151d7565b600060405180830381865afa158015611066573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261108e9190810190615221565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110c0908b908b908b906004016152d8565b600060405180830381865afa1580156110dd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111059190810190615221565b6040820152856001600160401b0381111561112257611122614432565b60405190808252806020026020018201604052801561115557816020015b60608152602001906001900390816111405790505b50606082015260005b60ff81168711156114f2576000856001600160401b0381111561118357611183614432565b6040519080825280602002602001820160405280156111ac578160200160208202803683370190505b5083606001518360ff16815181106111c6576111c66150b5565b602002602001018190525060005b868110156113f25760008c6001600160a01b03166304ec63518a8a858181106111ff576111ff6150b5565b905060200201358e8860000151868151811061121d5761121d6150b5565b60200260200101516040518463ffffffff1660e01b815260040161125a9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611277573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129b9190615301565b90508a8a8560ff168181106112b2576112b26150b5565b6001600160c01b03841692013560f81c9190911c6001908116141590506113df57856001600160a01b031663dd9846b98a8a858181106112f4576112f46150b5565b905060200201358d8d8860ff16818110611310576113106150b5565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611366573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138a919061532a565b85606001518560ff16815181106113a3576113a36150b5565b602002602001015184815181106113bc576113bc6150b5565b63ffffffff90921660209283029190910190910152826113db816151bc565b9350505b50806113ea816151bc565b9150506111d4565b506000816001600160401b0381111561140d5761140d614432565b604051908082528060200260200182016040528015611436578160200160208202803683370190505b50905060005b828110156114b75784606001518460ff168151811061145d5761145d6150b5565b60200260200101518181518110611476576114766150b5565b6020026020010151828281518110611490576114906150b5565b63ffffffff90921660209283029190910190910152806114af816151bc565b91505061143c565b508084606001518460ff16815181106114d2576114d26150b5565b6020026020010181905250505080806114ea90615347565b91505061115e565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611533573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115579190614fe9565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061158a908b908b908e90600401615367565b600060405180830381865afa1580156115a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115cf9190810190615221565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611629573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164d9190615050565b6116695760405162461bcd60e51b81526004016106749061506d565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b609b546001600160a01b031633146117025760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610674565b60006117146040850160208601614605565b90503660006117266040870187615391565b9092509050600061173d6080880160608901614605565b9050609860006117506020890189614605565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161177c91906153d7565b60405160208183030381529060405280519060200120146118055760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b600060998161181760208a018a614605565b63ffffffff1663ffffffff16815260200190815260200160002054146118945760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610674565b6118be7f000000000000000000000000000000000000000000000000000000000000000085615478565b63ffffffff164363ffffffff16111561192f5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610674565b60008660405160200161194291906154be565b60405160208183030381529060405280519060200120905060008061196a8387878a8c6120f9565b9150915060005b85811015611a69578460ff1683602001518281518110611993576119936150b5565b60200260200101516119a591906154cc565b6001600160601b03166064846000015183815181106119c6576119c66150b5565b60200260200101516001600160601b03166119e191906154fb565b1015611a57576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610674565b80611a61816151bc565b915050611971565b5060408051808201825263ffffffff43168152602080820184905291519091611a96918c9184910161551a565b60405160208183030381529060405280519060200120609960008c6000016020810190611ac39190614605565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a82604051611b1292919061551a565b60405180910390a15050505050505050505050565b6000611b366020850185614605565b63ffffffff8116600090815260996020526040902054909150853590611ba85760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610674565b8484604051602001611bbb929190615546565b60408051601f19818403018152918152815160209283012063ffffffff85166000908152609990935291205414611c5a5760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b63ffffffff82166000908152609a602052604090205460ff1615611cf25760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610674565b6064611d016020860186614605565b611d0b9190615478565b63ffffffff164363ffffffff161115611d8c5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610674565b6000611d9882806154fb565b9050602086013581146001811415611de657604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a3505050506120f3565b600085516001600160401b03811115611e0157611e01614432565b604051908082528060200260200182016040528015611e2a578160200160208202803683370190505b50905060005b8651811015611e9c57611e6d878281518110611e4e57611e4e6150b5565b6020026020010151805160009081526020918201519091526040902090565b828281518110611e7f57611e7f6150b5565b602090810291909101015280611e94816151bc565b915050611e30565b506000611eaf60408b0160208c01614605565b82604051602001611ec192919061557c565b60405160208183030381529060405280519060200120905087602001358114611f6b5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610674565b600087516001600160401b03811115611f8657611f86614432565b604051908082528060200260200182016040528015611faf578160200160208202803683370190505b50905060005b88518110156120a2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611fff57611fff6150b5565b60200260200101516040518263ffffffff1660e01b815260040161202591815260200190565b602060405180830381865afa158015612042573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120669190614fe9565b828281518110612078576120786150b5565b6001600160a01b03909216602092830291909101909101528061209a816151bc565b915050611fb5565b5063ffffffff87166000818152609a6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505b50505050565b60408051808201909152606080825260208201526000846121705760405162461bcd60e51b815260206004820152603760248201526000805160206157d383398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610674565b60408301515185148015612188575060a08301515185145b8015612198575060c08301515185145b80156121a8575060e08301515185145b6122125760405162461bcd60e51b815260206004820152604160248201526000805160206157d383398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610674565b8251516020840151511461228a5760405162461bcd60e51b8152602060048201526044602482018190526000805160206157d3833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610674565b4363ffffffff168463ffffffff1611156122fa5760405162461bcd60e51b815260206004820152603c60248201526000805160206157d383398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610674565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561233b5761233b614432565b604051908082528060200260200182016040528015612364578160200160208202803683370190505b506020820152866001600160401b0381111561238257612382614432565b6040519080825280602002602001820160405280156123ab578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156123df576123df614432565b604051908082528060200260200182016040528015612408578160200160208202803683370190505b5081526020860151516001600160401b0381111561242857612428614432565b604051908082528060200260200182016040528015612451578160200160208202803683370190505b50816020018190525060006125238a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156124fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061251e91906155c4565b613bfe565b905060005b87602001515181101561279f5761254e88602001518281518110611e4e57611e4e6150b5565b83602001518281518110612564576125646150b5565b602090810291909101015280156126245760208301516125856001836155e1565b81518110612595576125956150b5565b602002602001015160001c836020015182815181106125b6576125b66150b5565b602002602001015160001c11612624576040805162461bcd60e51b81526020600482015260248101919091526000805160206157d383398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612669576126696150b5565b60200260200101518b8b600001518581518110612688576126886150b5565b60200260200101516040518463ffffffff1660e01b81526004016126c59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156126e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127069190615301565b6001600160c01b031683600001518281518110612725576127256150b5565b60200260200101818152505061278b6108e361275f8486600001518581518110612751576127516150b5565b602002602001015116613cb9565b8a602001518481518110612775576127756150b5565b6020026020010151613ce490919063ffffffff16565b945080612797816151bc565b915050612528565b50506127aa83613dc8565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561280c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061283091906155f8565b60975490915060ff1660005b8a811015612eb6578115612998578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612893576128936150b5565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156128d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f791906155f8565b6129019190615611565b10156129985760405162461bcd60e51b815260206004820152606660248201526000805160206157d383398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106129d9576129d96150b5565b9050013560f81c60f81b60f81c8c8c60a0015185815181106129fd576129fd6150b5565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612a59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a7d9190615629565b6001600160401b031916612aa08a604001518381518110611e4e57611e4e6150b5565b67ffffffffffffffff191614612b3c5760405162461bcd60e51b815260206004820152606160248201526000805160206157d383398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610674565b612b6c89604001518281518110612b5557612b556150b5565b6020026020010151876137f690919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612baf57612baf6150b5565b9050013560f81c60f81b60f81c8c8c60c001518581518110612bd357612bd36150b5565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612c2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c53919061517d565b85602001518281518110612c6957612c696150b5565b6001600160601b03909216602092830291909101820152850151805182908110612c9557612c956150b5565b602002602001015185600001518281518110612cb357612cb36150b5565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612ea157612d2b86600001518281518110612cfd57612cfd6150b5565b60200260200101518f8f86818110612d1757612d176150b5565b600192013560f81c9290921c811614919050565b15612e8f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612d7157612d716150b5565b9050013560f81c60f81b60f81c8e89602001518581518110612d9557612d956150b5565b60200260200101518f60e001518881518110612db357612db36150b5565b60200260200101518781518110612dcc57612dcc6150b5565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612e30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e54919061517d565b8751805185908110612e6857612e686150b5565b60200260200101818151612e7c9190615654565b6001600160601b03169052506001909101905b80612e99816151bc565b915050612cd7565b50508080612eae906151bc565b91505061283c565b505050600080612ed08c868a606001518b608001516107c8565b9150915081612f415760405162461bcd60e51b815260206004820152604360248201526000805160206157d383398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610674565b80612fa25760405162461bcd60e51b815260206004820152603960248201526000805160206157d383398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610674565b50506000878260200151604051602001612fbd92919061557c565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612fef613e63565b612ff96000613ebd565b565b609c546001600160a01b0316331461305f5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610674565b613096604051806080016040528060008152602001600063ffffffff16815260200160608152602001600063ffffffff1681525090565b63ffffffff8086168252438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516130fe9082906020016156a8565b60408051601f1981840301815282825280516020918201206097805463ffffffff610100918290048116600090815260989095529490932091909155540416907f1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c489061316b9084906156a8565b60405180910390a260975461318c90610100900463ffffffff166001615478565b609760016101000a81548163ffffffff021916908363ffffffff1602179055505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106131ee576131ee6150b5565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061322a908890869060040161570d565b600060405180830381865afa158015613247573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261326f9190810190615221565b600081518110613281576132816150b5565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156132ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133119190615301565b6001600160c01b03169050600061332782613f0f565b9050816133358a838a610952565b9550955050505050935093915050565b61334d613e63565b6001600160a01b0381166133b25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610674565b61068681613ebd565b600054610100900460ff16158080156133db5750600054600160ff909116105b806133f55750303b1580156133f5575060005460ff166001145b6134585760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610674565b6000805460ff19166001179055801561347b576000805461ff0019166101001790555b613486856000613f6c565b61348f84613ebd565b609b80546001600160a01b038086166001600160a01b031992831617909255609c8054928516929091169190911790558015613505576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561355f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135839190614fe9565b6001600160a01b0316336001600160a01b0316146135b35760405162461bcd60e51b815260040161067490615006565b6066541981196066541916146136315760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107bd565b6001600160a01b0381166136f65760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610674565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261377b61430d565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156137ae576137b0565bfe5b50806137ee5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610674565b505092915050565b604080518082019091526000808252602082015261381261432b565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156137ae5750806137ee5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610674565b613892614349565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061397a6000805160206157b3833981519152866150cb565b90505b61398681614056565b90935091506000805160206157b38339815191528283098314156139c0576040805180820190915290815260208101919091529392505050565b6000805160206157b383398151915260018208905061397d565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613a0c61436e565b60005b6002811015613bd1576000613a258260066154fb565b9050848260028110613a3957613a396150b5565b60200201515183613a4b836000615611565b600c8110613a5b57613a5b6150b5565b6020020152848260028110613a7257613a726150b5565b60200201516020015183826001613a899190615611565b600c8110613a9957613a996150b5565b6020020152838260028110613ab057613ab06150b5565b6020020151515183613ac3836002615611565b600c8110613ad357613ad36150b5565b6020020152838260028110613aea57613aea6150b5565b6020020151516001602002015183613b03836003615611565b600c8110613b1357613b136150b5565b6020020152838260028110613b2a57613b2a6150b5565b602002015160200151600060028110613b4557613b456150b5565b602002015183613b56836004615611565b600c8110613b6657613b666150b5565b6020020152838260028110613b7d57613b7d6150b5565b602002015160200151600160028110613b9857613b986150b5565b602002015183613ba9836005615611565b600c8110613bb957613bb96150b5565b60200201525080613bc9816151bc565b915050613a0f565b50613bda61438d565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613c0a846140d8565b90508015613cb0578260ff168460018651613c2591906155e1565b81518110613c3557613c356150b5565b016020015160f81c10613cb05760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610674565b90505b92915050565b6000805b8215613cb357613cce6001846155e1565b9092169180613cdc81615761565b915050613cbd565b60408051808201909152600080825260208201526102008261ffff1610613d405760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610674565b8161ffff1660011415613d54575081613cb3565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613dbd57600161ffff871660ff83161c81161415613da057613d9d84846137f6565b93505b613daa83846137f6565b92506201fffe600192831b169101613d70565b509195945050505050565b60408051808201909152600080825260208201528151158015613ded57506020820151155b15613e0b575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206157b38339815191528460200151613e3e91906150cb565b613e56906000805160206157b38339815191526155e1565b905292915050565b919050565b6033546001600160a01b03163314612ff95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610674565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613f65576001811b915083821615613f5557828160f81b604051602001613f43929190615783565b60405160208183030381529060405292505b613f5e816151bc565b9050613f15565b5050919050565b6065546001600160a01b0316158015613f8d57506001600160a01b03821615155b61400f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261405282613668565b5050565b600080806000805160206157b383398151915260036000805160206157b3833981519152866000805160206157b38339815191528889090908905060006140cc827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206157b3833981519152614265565b91959194509092505050565b6000610100825111156141615760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610674565b815161416f57506000919050565b60008083600081518110614185576141856150b5565b0160200151600160f89190911c81901b92505b845181101561425c578481815181106141b3576141b36150b5565b0160200151600160f89190911c1b91508282116142485760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610674565b91811791614255816151bc565b9050614198565b50909392505050565b60008061427061438d565b6142786143ab565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156137ae5750826143025760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610674565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061435c6143c9565b81526020016143696143c9565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068657600080fd5b60006020828403121561440e57600080fd5b8135613cb0816143e7565b60006020828403121561442b57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561446a5761446a614432565b60405290565b60405161010081016001600160401b038111828210171561446a5761446a614432565b604051601f8201601f191681016001600160401b03811182821017156144bb576144bb614432565b604052919050565b6000604082840312156144d557600080fd5b6144dd614448565b9050813581526020820135602082015292915050565b600082601f83011261450457600080fd5b604051604081018181106001600160401b038211171561452657614526614432565b806040525080604084018581111561453d57600080fd5b845b81811015613dbd57803583526020928301920161453f565b60006080828403121561456957600080fd5b614571614448565b905061457d83836144f3565b815261458c83604084016144f3565b602082015292915050565b60008060008061012085870312156145ae57600080fd5b843593506145bf86602087016144c3565b92506145ce8660608701614557565b91506145dd8660e087016144c3565b905092959194509250565b63ffffffff8116811461068657600080fd5b8035613e5e816145e8565b60006020828403121561461757600080fd5b8135613cb0816145e8565b60008060006060848603121561463757600080fd5b8335614642816143e7565b92506020848101356001600160401b038082111561465f57600080fd5b818701915087601f83011261467357600080fd5b81358181111561468557614685614432565b614697601f8201601f19168501614493565b915080825288848285010111156146ad57600080fd5b80848401858401376000848284010152508094505050506146d0604085016145fa565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561476f578385038a52825180518087529087019087870190845b8181101561475a57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614716565b50509a87019a955050918501916001016146f8565b509298975050505050505050565b60208152600061479060208301846146d9565b9392505050565b801515811461068657600080fd5b6000602082840312156147b757600080fd5b8135613cb081614797565b60008083601f8401126147d457600080fd5b5081356001600160401b038111156147eb57600080fd5b60208301915083602082850101111561480357600080fd5b9250929050565b6000806000806000806080878903121561482357600080fd5b863561482e816143e7565b9550602087013561483e816145e8565b945060408701356001600160401b038082111561485a57600080fd5b6148668a838b016147c2565b9096509450606089013591508082111561487f57600080fd5b818901915089601f83011261489357600080fd5b8135818111156148a257600080fd5b8a60208260051b85010111156148b757600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561490357815163ffffffff16875295820195908201906001016148e1565b509495945050505050565b60006020808352835160808285015261492a60a08501826148cd565b905081850151601f198086840301604087015261494783836148cd565b9250604087015191508086840301606087015261496483836148cd565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156149bb57848783030184526149a98287516148cd565b9588019593880193915060010161498f565b509998505050505050505050565b60ff8116811461068657600080fd5b6000602082840312156149ea57600080fd5b8135613cb0816149c9565b600060808284031215614a0757600080fd5b50919050565b600060408284031215614a0757600080fd5b60006001600160401b03821115614a3857614a38614432565b5060051b60200190565b600082601f830112614a5357600080fd5b81356020614a68614a6383614a1f565b614493565b82815260059290921b84018101918181019086841115614a8757600080fd5b8286015b84811015614aab578035614a9e816145e8565b8352918301918301614a8b565b509695505050505050565b600082601f830112614ac757600080fd5b81356020614ad7614a6383614a1f565b82815260069290921b84018101918181019086841115614af657600080fd5b8286015b84811015614aab57614b0c88826144c3565b835291830191604001614afa565b600082601f830112614b2b57600080fd5b81356020614b3b614a6383614a1f565b82815260059290921b84018101918181019086841115614b5a57600080fd5b8286015b84811015614aab5780356001600160401b03811115614b7d5760008081fd5b614b8b8986838b0101614a42565b845250918301918301614b5e565b60006101808284031215614bac57600080fd5b614bb4614470565b905081356001600160401b0380821115614bcd57600080fd5b614bd985838601614a42565b83526020840135915080821115614bef57600080fd5b614bfb85838601614ab6565b60208401526040840135915080821115614c1457600080fd5b614c2085838601614ab6565b6040840152614c328560608601614557565b6060840152614c448560e086016144c3565b6080840152610120840135915080821115614c5e57600080fd5b614c6a85838601614a42565b60a0840152610140840135915080821115614c8457600080fd5b614c9085838601614a42565b60c0840152610160840135915080821115614caa57600080fd5b50614cb784828501614b1a565b60e08301525092915050565b600080600060808486031215614cd857600080fd5b83356001600160401b0380821115614cef57600080fd5b614cfb878388016149f5565b9450614d0a8760208801614a0d565b93506060860135915080821115614d2057600080fd5b50614d2d86828701614b99565b9150509250925092565b60008060008060c08587031215614d4d57600080fd5b84356001600160401b0380821115614d6457600080fd5b614d70888389016149f5565b9550614d7f8860208901614a0d565b9450614d8e8860608901614a0d565b935060a0870135915080821115614da457600080fd5b50614db187828801614ab6565b91505092959194509250565b600080600080600060808688031215614dd557600080fd5b8535945060208601356001600160401b0380821115614df357600080fd5b614dff89838a016147c2565b909650945060408801359150614e14826145e8565b90925060608701359080821115614e2a57600080fd5b50614e3788828901614b99565b9150509295509295909350565b600081518084526020808501945080840160005b838110156149035781516001600160601b031687529582019590820190600101614e58565b6040815260008351604080840152614e986080840182614e44565b90506020850151603f19848303016060850152614eb58282614e44565b925050508260208301529392505050565b60008060008060608587031215614edc57600080fd5b8435614ee7816145e8565b93506020850135614ef7816145e8565b925060408501356001600160401b03811115614f1257600080fd5b614f1e878288016147c2565b95989497509550505050565b600080600060608486031215614f3f57600080fd5b8335614f4a816143e7565b9250602084013591506040840135614f61816145e8565b809150509250925092565b828152604060208201526000614f8560408301846146d9565b949350505050565b60008060008060808587031215614fa357600080fd5b8435614fae816143e7565b93506020850135614fbe816143e7565b92506040850135614fce816143e7565b91506060850135614fde816143e7565b939692955090935050565b600060208284031215614ffb57600080fd5b8151613cb0816143e7565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561506257600080fd5b8151613cb081614797565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826150e857634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561510057600080fd5b82516001600160401b0381111561511657600080fd5b8301601f8101851361512757600080fd5b8051615135614a6382614a1f565b81815260059190911b8201830190838101908783111561515457600080fd5b928401925b8284101561517257835182529284019290840190615159565b979650505050505050565b60006020828403121561518f57600080fd5b81516001600160601b0381168114613cb057600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156151d0576151d06151a6565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561520457600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561523457600080fd5b82516001600160401b0381111561524a57600080fd5b8301601f8101851361525b57600080fd5b8051615269614a6382614a1f565b81815260059190911b8201830190838101908783111561528857600080fd5b928401925b828410156151725783516152a0816145e8565b8252928401929084019061528d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006152f86040830184866152af565b95945050505050565b60006020828403121561531357600080fd5b81516001600160c01b0381168114613cb057600080fd5b60006020828403121561533c57600080fd5b8151613cb0816145e8565b600060ff821660ff81141561535e5761535e6151a6565b60010192915050565b60408152600061537b6040830185876152af565b905063ffffffff83166020830152949350505050565b6000808335601e198436030181126153a857600080fd5b8301803591506001600160401b038211156153c257600080fd5b60200191503681900382131561480357600080fd5b6020815281356020820152600060208301356153f2816145e8565b63ffffffff81166040840152506040830135601e1984360301811261541657600080fd5b830180356001600160401b0381111561542e57600080fd5b80360385131561543d57600080fd5b6080606085015261545560a0850182602085016152af565b915050615464606085016145fa565b63ffffffff81166080850152509392505050565b600063ffffffff808316818516808303821115615497576154976151a6565b01949350505050565b80356154ab816145e8565b63ffffffff168252602090810135910152565b60408101613cb382846154a0565b60006001600160601b03808316818516818304811182151516156154f2576154f26151a6565b02949350505050565b6000816000190483118215151615615515576155156151a6565b500290565b6080810161552882856154a0565b63ffffffff8351166040830152602083015160608301529392505050565b6080810161555482856154a0565b823561555f816145e8565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156155b75781518552938201939082019060010161559b565b5092979650505050505050565b6000602082840312156155d657600080fd5b8151613cb0816149c9565b6000828210156155f3576155f36151a6565b500390565b60006020828403121561560a57600080fd5b5051919050565b60008219821115615624576156246151a6565b500190565b60006020828403121561563b57600080fd5b815167ffffffffffffffff1981168114613cb057600080fd5b60006001600160601b0383811690831681811015615674576156746151a6565b039392505050565b60005b8381101561569757818101518382015260200161567f565b838111156120f35750506000910152565b60208152815160208201526000602083015163ffffffff8082166040850152604085015191506080606085015281518060a08601526156ee8160c087016020860161567c565b60609590950151166080840152505060c0601f909201601f1916010190565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561575457845183529383019391830191600101615738565b5090979650505050505050565b600061ffff80831681811415615779576157796151a6565b6001019392505050565b6000835161579581846020880161567c565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212201ea30e5210fc34b3cebcbb12c7b0a5fdd61316c29839108f909dd17ce7b8dddd64736f6c634300080c0033",
}

// ContractOracleTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOracleTaskManagerMetaData.ABI instead.
var ContractOracleTaskManagerABI = ContractOracleTaskManagerMetaData.ABI

// ContractOracleTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOracleTaskManagerMetaData.Bin instead.
var ContractOracleTaskManagerBin = ContractOracleTaskManagerMetaData.Bin

// DeployContractOracleTaskManager deploys a new Ethereum contract, binding an instance of ContractOracleTaskManager to it.
func DeployContractOracleTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractOracleTaskManager, error) {
	parsed, err := ContractOracleTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOracleTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOracleTaskManager{ContractOracleTaskManagerCaller: ContractOracleTaskManagerCaller{contract: contract}, ContractOracleTaskManagerTransactor: ContractOracleTaskManagerTransactor{contract: contract}, ContractOracleTaskManagerFilterer: ContractOracleTaskManagerFilterer{contract: contract}}, nil
}

// ContractOracleTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractOracleTaskManager struct {
	ContractOracleTaskManagerCaller     // Read-only binding to the contract
	ContractOracleTaskManagerTransactor // Write-only binding to the contract
	ContractOracleTaskManagerFilterer   // Log filterer for contract events
}

// ContractOracleTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOracleTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOracleTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOracleTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOracleTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOracleTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOracleTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOracleTaskManagerSession struct {
	Contract     *ContractOracleTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractOracleTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOracleTaskManagerCallerSession struct {
	Contract *ContractOracleTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractOracleTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOracleTaskManagerTransactorSession struct {
	Contract     *ContractOracleTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractOracleTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOracleTaskManagerRaw struct {
	Contract *ContractOracleTaskManager // Generic contract binding to access the raw methods on
}

// ContractOracleTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOracleTaskManagerCallerRaw struct {
	Contract *ContractOracleTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOracleTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOracleTaskManagerTransactorRaw struct {
	Contract *ContractOracleTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOracleTaskManager creates a new instance of ContractOracleTaskManager, bound to a specific deployed contract.
func NewContractOracleTaskManager(address common.Address, backend bind.ContractBackend) (*ContractOracleTaskManager, error) {
	contract, err := bindContractOracleTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManager{ContractOracleTaskManagerCaller: ContractOracleTaskManagerCaller{contract: contract}, ContractOracleTaskManagerTransactor: ContractOracleTaskManagerTransactor{contract: contract}, ContractOracleTaskManagerFilterer: ContractOracleTaskManagerFilterer{contract: contract}}, nil
}

// NewContractOracleTaskManagerCaller creates a new read-only instance of ContractOracleTaskManager, bound to a specific deployed contract.
func NewContractOracleTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOracleTaskManagerCaller, error) {
	contract, err := bindContractOracleTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerCaller{contract: contract}, nil
}

// NewContractOracleTaskManagerTransactor creates a new write-only instance of ContractOracleTaskManager, bound to a specific deployed contract.
func NewContractOracleTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOracleTaskManagerTransactor, error) {
	contract, err := bindContractOracleTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerTransactor{contract: contract}, nil
}

// NewContractOracleTaskManagerFilterer creates a new log filterer instance of ContractOracleTaskManager, bound to a specific deployed contract.
func NewContractOracleTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOracleTaskManagerFilterer, error) {
	contract, err := bindContractOracleTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerFilterer{contract: contract}, nil
}

// bindContractOracleTaskManager binds a generic wrapper to an already deployed contract.
func bindContractOracleTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOracleTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOracleTaskManager *ContractOracleTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOracleTaskManager.Contract.ContractOracleTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOracleTaskManager *ContractOracleTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.ContractOracleTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOracleTaskManager *ContractOracleTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.ContractOracleTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOracleTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractOracleTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractOracleTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOracleTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOracleTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Aggregator(&_ContractOracleTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Aggregator(&_ContractOracleTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOracleTaskManager.Contract.AllTaskHashes(&_ContractOracleTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOracleTaskManager.Contract.AllTaskHashes(&_ContractOracleTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOracleTaskManager.Contract.AllTaskResponses(&_ContractOracleTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOracleTaskManager.Contract.AllTaskResponses(&_ContractOracleTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.BlsApkRegistry(&_ContractOracleTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.BlsApkRegistry(&_ContractOracleTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOracleTaskManager.Contract.CheckSignatures(&_ContractOracleTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOracleTaskManager.Contract.CheckSignatures(&_ContractOracleTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Delegation(&_ContractOracleTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Delegation(&_ContractOracleTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Generator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Generator(&_ContractOracleTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Generator(&_ContractOracleTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOracleTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOracleTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOracleTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOracleTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOracleTaskManager.Contract.GetOperatorState(&_ContractOracleTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOracleTaskManager.Contract.GetOperatorState(&_ContractOracleTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOracleTaskManager.Contract.GetOperatorState0(&_ContractOracleTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOracleTaskManager.Contract.GetOperatorState0(&_ContractOracleTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOracleTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOracleTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOracleTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOracleTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractOracleTaskManager.Contract.LatestTaskNum(&_ContractOracleTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractOracleTaskManager.Contract.LatestTaskNum(&_ContractOracleTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Owner() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Owner(&_ContractOracleTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.Owner(&_ContractOracleTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractOracleTaskManager.Contract.Paused(&_ContractOracleTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractOracleTaskManager.Contract.Paused(&_ContractOracleTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractOracleTaskManager.Contract.Paused0(&_ContractOracleTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractOracleTaskManager.Contract.Paused0(&_ContractOracleTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.PauserRegistry(&_ContractOracleTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.PauserRegistry(&_ContractOracleTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.RegistryCoordinator(&_ContractOracleTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.RegistryCoordinator(&_ContractOracleTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.StakeRegistry(&_ContractOracleTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractOracleTaskManager.Contract.StakeRegistry(&_ContractOracleTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOracleTaskManager.Contract.StaleStakesForbidden(&_ContractOracleTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOracleTaskManager.Contract.StaleStakesForbidden(&_ContractOracleTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TaskNumber(&_ContractOracleTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractOracleTaskManager.Contract.TaskNumber(&_ContractOracleTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractOracleTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractOracleTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractOracleTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractOracleTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractOracleTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOracleTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOracleTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOracleTaskManager *ContractOracleTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOracleTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOracleTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 dolarDatetime, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, dolarDatetime uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "createNewTask", dolarDatetime, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 dolarDatetime, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) CreateNewTask(dolarDatetime uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.CreateNewTask(&_ContractOracleTaskManager.TransactOpts, dolarDatetime, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x9d3a0f2d.
//
// Solidity: function createNewTask(uint32 dolarDatetime, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) CreateNewTask(dolarDatetime uint32, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.CreateNewTask(&_ContractOracleTaskManager.TransactOpts, dolarDatetime, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Initialize(&_ContractOracleTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Initialize(&_ContractOracleTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Pause(&_ContractOracleTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Pause(&_ContractOracleTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.PauseAll(&_ContractOracleTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.PauseAll(&_ContractOracleTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, taskResponseMetadata IOracleTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) RaiseAndResolveChallenge(task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, taskResponseMetadata IOracleTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RaiseAndResolveChallenge(&_ContractOracleTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) RaiseAndResolveChallenge(task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, taskResponseMetadata IOracleTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RaiseAndResolveChallenge(&_ContractOracleTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RenounceOwnership(&_ContractOracleTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RenounceOwnership(&_ContractOracleTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) RespondToTask(task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RespondToTask(&_ContractOracleTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) RespondToTask(task IOracleTaskManagerTask, taskResponse IOracleTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.RespondToTask(&_ContractOracleTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.SetPauserRegistry(&_ContractOracleTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.SetPauserRegistry(&_ContractOracleTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.SetStaleStakesForbidden(&_ContractOracleTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.SetStaleStakesForbidden(&_ContractOracleTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.TransferOwnership(&_ContractOracleTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.TransferOwnership(&_ContractOracleTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Unpause(&_ContractOracleTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOracleTaskManager *ContractOracleTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOracleTaskManager.Contract.Unpause(&_ContractOracleTaskManager.TransactOpts, newPausedStatus)
}

// ContractOracleTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerInitializedIterator struct {
	Event *ContractOracleTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerInitialized represents a Initialized event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOracleTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerInitializedIterator{contract: _ContractOracleTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerInitialized)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractOracleTaskManagerInitialized, error) {
	event := new(ContractOracleTaskManagerInitialized)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerNewTaskCreatedIterator struct {
	Event *ContractOracleTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IOracleTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOracleTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerNewTaskCreatedIterator{contract: _ContractOracleTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerNewTaskCreated)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractOracleTaskManagerNewTaskCreated, error) {
	event := new(ContractOracleTaskManagerNewTaskCreated)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerOwnershipTransferredIterator struct {
	Event *ContractOracleTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOracleTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerOwnershipTransferredIterator{contract: _ContractOracleTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerOwnershipTransferred)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOracleTaskManagerOwnershipTransferred, error) {
	event := new(ContractOracleTaskManagerOwnershipTransferred)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerPausedIterator struct {
	Event *ContractOracleTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerPaused represents a Paused event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractOracleTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerPausedIterator{contract: _ContractOracleTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerPaused)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParsePaused(log types.Log) (*ContractOracleTaskManagerPaused, error) {
	event := new(ContractOracleTaskManagerPaused)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerPauserRegistrySetIterator struct {
	Event *ContractOracleTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractOracleTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerPauserRegistrySetIterator{contract: _ContractOracleTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerPauserRegistrySet)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractOracleTaskManagerPauserRegistrySet, error) {
	event := new(ContractOracleTaskManagerPauserRegistrySet)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractOracleTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractOracleTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractOracleTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractOracleTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractOracleTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractOracleTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractOracleTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerTaskChallengedSuccessfully)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractOracleTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractOracleTaskManagerTaskChallengedSuccessfully)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractOracleTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractOracleTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractOracleTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractOracleTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskCompletedIterator struct {
	Event *ContractOracleTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOracleTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerTaskCompletedIterator{contract: _ContractOracleTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerTaskCompleted)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractOracleTaskManagerTaskCompleted, error) {
	event := new(ContractOracleTaskManagerTaskCompleted)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskRespondedIterator struct {
	Event *ContractOracleTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerTaskResponded represents a TaskResponded event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerTaskResponded struct {
	TaskResponse         IOracleTaskManagerTaskResponse
	TaskResponseMetadata IOracleTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOracleTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerTaskRespondedIterator{contract: _ContractOracleTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerTaskResponded)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractOracleTaskManagerTaskResponded, error) {
	event := new(ContractOracleTaskManagerTaskResponded)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerUnpausedIterator struct {
	Event *ContractOracleTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOracleTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOracleTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOracleTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleTaskManagerUnpaused represents a Unpaused event raised by the ContractOracleTaskManager contract.
type ContractOracleTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractOracleTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleTaskManagerUnpausedIterator{contract: _ContractOracleTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractOracleTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOracleTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleTaskManagerUnpaused)
				if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOracleTaskManager *ContractOracleTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractOracleTaskManagerUnpaused, error) {
	event := new(ContractOracleTaskManagerUnpaused)
	if err := _ContractOracleTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
