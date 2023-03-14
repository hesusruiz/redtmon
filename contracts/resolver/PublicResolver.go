// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolver

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ResolverABI is the input ABI used to generate the binding from.
const ResolverABI = "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"document\",\"type\":\"string\"}],\"name\":\"AlaDIDDocumentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"}],\"name\":\"AlaDIDPubkeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"}],\"name\":\"AlaDIDPubkeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"}],\"name\":\"AlaDIDPubkeyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuthorised\",\"type\":\"bool\"}],\"name\":\"AuthorisationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"AlaDIDPublicEntity\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_DIDHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_domain_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_DIDDocument\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"AlaTSP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"AlaTSPNumberServices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlaTSPService\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"X509SKI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"X509Certificate\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"X509SKI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"X509Certificate\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"addAlaTSPService\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"partDIDHash\",\"type\":\"bytes32\"}],\"name\":\"addParticipant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyValue\",\"type\":\"string\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_DIDHash\",\"type\":\"bytes32\"}],\"name\":\"addressFromDID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorisations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"partDIDHash\",\"type\":\"bytes32\"}],\"name\":\"confirmCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"credential\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"credentialHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numParticipants\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"credentialParticipant\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"didHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"signed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_DIDHash\",\"type\":\"bytes32\"}],\"name\":\"nodeFromDID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_node\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"}],\"name\":\"publicKey\",\"outputs\":[{\"internalType\":\"enumAlaDIDPubkeyResolver.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"keyValue\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"publicKeyAtIndex\",\"outputs\":[{\"internalType\":\"enumAlaDIDPubkeyResolver.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"keyValue\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"keyID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_DIDHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_DIDDocument\",\"type\":\"string\"}],\"name\":\"setAlaDIDDocument\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_DIDHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_domain_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_DIDDocument\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setAlaDIDPublicEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"org\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setAlaTSP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuthorised\",\"type\":\"bool\"}],\"name\":\"setAuthorisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"credentialHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"part1DIDhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part2DIDhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part3DIDhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part4DIDhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part5DIDhash\",\"type\":\"bytes32\"}],\"name\":\"setCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ResolverParsedABI, _ = abi.JSON(strings.NewReader(ResolverABI))

// Resolver is an auto generated Go binding around an Ethereum contract.
type Resolver struct {
	ResolverCaller     // Read-only binding to the contract
	ResolverTransactor // Write-only binding to the contract
	ResolverFilterer   // Log filterer for contract events
}

// ResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverSession struct {
	Contract     *Resolver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverCallerSession struct {
	Contract *ResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverTransactorSession struct {
	Contract     *ResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRaw struct {
	Contract *Resolver // Generic contract binding to access the raw methods on
}

// ResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverCallerRaw struct {
	Contract *ResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverTransactorRaw struct {
	Contract *ResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolver creates a new instance of Resolver, bound to a specific deployed contract.
func NewResolver(address common.Address, backend bind.ContractBackend) (*Resolver, error) {
	contract, err := bindResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// NewResolverCaller creates a new read-only instance of Resolver, bound to a specific deployed contract.
func NewResolverCaller(address common.Address, caller bind.ContractCaller) (*ResolverCaller, error) {
	contract, err := bindResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverCaller{contract: contract}, nil
}

// NewResolverTransactor creates a new write-only instance of Resolver, bound to a specific deployed contract.
func NewResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverTransactor, error) {
	contract, err := bindResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverTransactor{contract: contract}, nil
}

// NewResolverFilterer creates a new log filterer instance of Resolver, bound to a specific deployed contract.
func NewResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverFilterer, error) {
	contract, err := bindResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverFilterer{contract: contract}, nil
}

// bindResolver binds a generic wrapper to an already deployed contract.
func bindResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.ResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transact(opts, method, params...)
}

// AlaDIDPublicEntity is a free data retrieval call binding the contract method 0xfe171107.
//
// Solidity: function AlaDIDPublicEntity(bytes32 node) view returns(bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active)
func (_Resolver *ResolverCaller) AlaDIDPublicEntity(opts *bind.CallOpts, node [32]byte) (struct {
	DIDHash     [32]byte
	DomainName  string
	DIDDocument string
	Active      bool
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "AlaDIDPublicEntity", node)

	outstruct := new(struct {
		DIDHash     [32]byte
		DomainName  string
		DIDDocument string
		Active      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DIDHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DomainName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.DIDDocument = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// AlaDIDPublicEntity is a free data retrieval call binding the contract method 0xfe171107.
//
// Solidity: function AlaDIDPublicEntity(bytes32 node) view returns(bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active)
func (_Resolver *ResolverSession) AlaDIDPublicEntity(node [32]byte) (struct {
	DIDHash     [32]byte
	DomainName  string
	DIDDocument string
	Active      bool
}, error) {
	return _Resolver.Contract.AlaDIDPublicEntity(&_Resolver.CallOpts, node)
}

// AlaDIDPublicEntity is a free data retrieval call binding the contract method 0xfe171107.
//
// Solidity: function AlaDIDPublicEntity(bytes32 node) view returns(bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active)
func (_Resolver *ResolverCallerSession) AlaDIDPublicEntity(node [32]byte) (struct {
	DIDHash     [32]byte
	DomainName  string
	DIDDocument string
	Active      bool
}, error) {
	return _Resolver.Contract.AlaDIDPublicEntity(&_Resolver.CallOpts, node)
}

// AlaTSP is a free data retrieval call binding the contract method 0x6782868c.
//
// Solidity: function AlaTSP(bytes32 node) view returns(string URI, string org, bool active)
func (_Resolver *ResolverCaller) AlaTSP(opts *bind.CallOpts, node [32]byte) (struct {
	URI    string
	Org    string
	Active bool
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "AlaTSP", node)

	outstruct := new(struct {
		URI    string
		Org    string
		Active bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.URI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Org = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// AlaTSP is a free data retrieval call binding the contract method 0x6782868c.
//
// Solidity: function AlaTSP(bytes32 node) view returns(string URI, string org, bool active)
func (_Resolver *ResolverSession) AlaTSP(node [32]byte) (struct {
	URI    string
	Org    string
	Active bool
}, error) {
	return _Resolver.Contract.AlaTSP(&_Resolver.CallOpts, node)
}

// AlaTSP is a free data retrieval call binding the contract method 0x6782868c.
//
// Solidity: function AlaTSP(bytes32 node) view returns(string URI, string org, bool active)
func (_Resolver *ResolverCallerSession) AlaTSP(node [32]byte) (struct {
	URI    string
	Org    string
	Active bool
}, error) {
	return _Resolver.Contract.AlaTSP(&_Resolver.CallOpts, node)
}

// AlaTSPNumberServices is a free data retrieval call binding the contract method 0x411ae317.
//
// Solidity: function AlaTSPNumberServices(bytes32 node) view returns(uint256)
func (_Resolver *ResolverCaller) AlaTSPNumberServices(opts *bind.CallOpts, node [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "AlaTSPNumberServices", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlaTSPNumberServices is a free data retrieval call binding the contract method 0x411ae317.
//
// Solidity: function AlaTSPNumberServices(bytes32 node) view returns(uint256)
func (_Resolver *ResolverSession) AlaTSPNumberServices(node [32]byte) (*big.Int, error) {
	return _Resolver.Contract.AlaTSPNumberServices(&_Resolver.CallOpts, node)
}

// AlaTSPNumberServices is a free data retrieval call binding the contract method 0x411ae317.
//
// Solidity: function AlaTSPNumberServices(bytes32 node) view returns(uint256)
func (_Resolver *ResolverCallerSession) AlaTSPNumberServices(node [32]byte) (*big.Int, error) {
	return _Resolver.Contract.AlaTSPNumberServices(&_Resolver.CallOpts, node)
}

// AlaTSPService is a free data retrieval call binding the contract method 0xf1c8c052.
//
// Solidity: function AlaTSPService(bytes32 node, uint256 index) view returns(string X509SKI, string serviceName, bytes X509Certificate, bool active)
func (_Resolver *ResolverCaller) AlaTSPService(opts *bind.CallOpts, node [32]byte, index *big.Int) (struct {
	X509SKI         string
	ServiceName     string
	X509Certificate []byte
	Active          bool
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "AlaTSPService", node, index)

	outstruct := new(struct {
		X509SKI         string
		ServiceName     string
		X509Certificate []byte
		Active          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X509SKI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ServiceName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.X509Certificate = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// AlaTSPService is a free data retrieval call binding the contract method 0xf1c8c052.
//
// Solidity: function AlaTSPService(bytes32 node, uint256 index) view returns(string X509SKI, string serviceName, bytes X509Certificate, bool active)
func (_Resolver *ResolverSession) AlaTSPService(node [32]byte, index *big.Int) (struct {
	X509SKI         string
	ServiceName     string
	X509Certificate []byte
	Active          bool
}, error) {
	return _Resolver.Contract.AlaTSPService(&_Resolver.CallOpts, node, index)
}

// AlaTSPService is a free data retrieval call binding the contract method 0xf1c8c052.
//
// Solidity: function AlaTSPService(bytes32 node, uint256 index) view returns(string X509SKI, string serviceName, bytes X509Certificate, bool active)
func (_Resolver *ResolverCallerSession) AlaTSPService(node [32]byte, index *big.Int) (struct {
	X509SKI         string
	ServiceName     string
	X509Certificate []byte
	Active          bool
}, error) {
	return _Resolver.Contract.AlaTSPService(&_Resolver.CallOpts, node, index)
}

// AddressFromDID is a free data retrieval call binding the contract method 0x32783614.
//
// Solidity: function addressFromDID(bytes32 _DIDHash) view returns(address _owner)
func (_Resolver *ResolverCaller) AddressFromDID(opts *bind.CallOpts, _DIDHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "addressFromDID", _DIDHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressFromDID is a free data retrieval call binding the contract method 0x32783614.
//
// Solidity: function addressFromDID(bytes32 _DIDHash) view returns(address _owner)
func (_Resolver *ResolverSession) AddressFromDID(_DIDHash [32]byte) (common.Address, error) {
	return _Resolver.Contract.AddressFromDID(&_Resolver.CallOpts, _DIDHash)
}

// AddressFromDID is a free data retrieval call binding the contract method 0x32783614.
//
// Solidity: function addressFromDID(bytes32 _DIDHash) view returns(address _owner)
func (_Resolver *ResolverCallerSession) AddressFromDID(_DIDHash [32]byte) (common.Address, error) {
	return _Resolver.Contract.AddressFromDID(&_Resolver.CallOpts, _DIDHash)
}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_Resolver *ResolverCaller) Authorisations(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "authorisations", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_Resolver *ResolverSession) Authorisations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _Resolver.Contract.Authorisations(&_Resolver.CallOpts, arg0, arg1, arg2)
}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_Resolver *ResolverCallerSession) Authorisations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _Resolver.Contract.Authorisations(&_Resolver.CallOpts, arg0, arg1, arg2)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Resolver *ResolverCaller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Resolver *ResolverSession) Contenthash(node [32]byte) ([]byte, error) {
	return _Resolver.Contract.Contenthash(&_Resolver.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Resolver *ResolverCallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _Resolver.Contract.Contenthash(&_Resolver.CallOpts, node)
}

// Credential is a free data retrieval call binding the contract method 0x79f5d409.
//
// Solidity: function credential(bytes32 node, string key) view returns(bytes credentialHash, uint256 numParticipants)
func (_Resolver *ResolverCaller) Credential(opts *bind.CallOpts, node [32]byte, key string) (struct {
	CredentialHash  []byte
	NumParticipants *big.Int
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "credential", node, key)

	outstruct := new(struct {
		CredentialHash  []byte
		NumParticipants *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CredentialHash = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.NumParticipants = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Credential is a free data retrieval call binding the contract method 0x79f5d409.
//
// Solidity: function credential(bytes32 node, string key) view returns(bytes credentialHash, uint256 numParticipants)
func (_Resolver *ResolverSession) Credential(node [32]byte, key string) (struct {
	CredentialHash  []byte
	NumParticipants *big.Int
}, error) {
	return _Resolver.Contract.Credential(&_Resolver.CallOpts, node, key)
}

// Credential is a free data retrieval call binding the contract method 0x79f5d409.
//
// Solidity: function credential(bytes32 node, string key) view returns(bytes credentialHash, uint256 numParticipants)
func (_Resolver *ResolverCallerSession) Credential(node [32]byte, key string) (struct {
	CredentialHash  []byte
	NumParticipants *big.Int
}, error) {
	return _Resolver.Contract.Credential(&_Resolver.CallOpts, node, key)
}

// CredentialParticipant is a free data retrieval call binding the contract method 0x473e2a04.
//
// Solidity: function credentialParticipant(bytes32 node, string key, uint256 index) view returns(bytes32 didHash, bool signed)
func (_Resolver *ResolverCaller) CredentialParticipant(opts *bind.CallOpts, node [32]byte, key string, index *big.Int) (struct {
	DidHash [32]byte
	Signed  bool
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "credentialParticipant", node, key, index)

	outstruct := new(struct {
		DidHash [32]byte
		Signed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DidHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Signed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CredentialParticipant is a free data retrieval call binding the contract method 0x473e2a04.
//
// Solidity: function credentialParticipant(bytes32 node, string key, uint256 index) view returns(bytes32 didHash, bool signed)
func (_Resolver *ResolverSession) CredentialParticipant(node [32]byte, key string, index *big.Int) (struct {
	DidHash [32]byte
	Signed  bool
}, error) {
	return _Resolver.Contract.CredentialParticipant(&_Resolver.CallOpts, node, key, index)
}

// CredentialParticipant is a free data retrieval call binding the contract method 0x473e2a04.
//
// Solidity: function credentialParticipant(bytes32 node, string key, uint256 index) view returns(bytes32 didHash, bool signed)
func (_Resolver *ResolverCallerSession) CredentialParticipant(node [32]byte, key string, index *big.Int) (struct {
	DidHash [32]byte
	Signed  bool
}, error) {
	return _Resolver.Contract.CredentialParticipant(&_Resolver.CallOpts, node, key, index)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string data) pure returns(bytes32)
func (_Resolver *ResolverCaller) Hash(opts *bind.CallOpts, data string) ([32]byte, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "hash", data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string data) pure returns(bytes32)
func (_Resolver *ResolverSession) Hash(data string) ([32]byte, error) {
	return _Resolver.Contract.Hash(&_Resolver.CallOpts, data)
}

// Hash is a free data retrieval call binding the contract method 0xb411ee94.
//
// Solidity: function hash(string data) pure returns(bytes32)
func (_Resolver *ResolverCallerSession) Hash(data string) ([32]byte, error) {
	return _Resolver.Contract.Hash(&_Resolver.CallOpts, data)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Resolver *ResolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Resolver *ResolverSession) Name(node [32]byte) (string, error) {
	return _Resolver.Contract.Name(&_Resolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Resolver *ResolverCallerSession) Name(node [32]byte) (string, error) {
	return _Resolver.Contract.Name(&_Resolver.CallOpts, node)
}

// NodeFromDID is a free data retrieval call binding the contract method 0x91b2fb8f.
//
// Solidity: function nodeFromDID(bytes32 _DIDHash) view returns(bytes32 _node)
func (_Resolver *ResolverCaller) NodeFromDID(opts *bind.CallOpts, _DIDHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "nodeFromDID", _DIDHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NodeFromDID is a free data retrieval call binding the contract method 0x91b2fb8f.
//
// Solidity: function nodeFromDID(bytes32 _DIDHash) view returns(bytes32 _node)
func (_Resolver *ResolverSession) NodeFromDID(_DIDHash [32]byte) ([32]byte, error) {
	return _Resolver.Contract.NodeFromDID(&_Resolver.CallOpts, _DIDHash)
}

// NodeFromDID is a free data retrieval call binding the contract method 0x91b2fb8f.
//
// Solidity: function nodeFromDID(bytes32 _DIDHash) view returns(bytes32 _node)
func (_Resolver *ResolverCallerSession) NodeFromDID(_DIDHash [32]byte) ([32]byte, error) {
	return _Resolver.Contract.NodeFromDID(&_Resolver.CallOpts, _DIDHash)
}

// PublicKey is a free data retrieval call binding the contract method 0xa00efac0.
//
// Solidity: function publicKey(bytes32 node, string keyID) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue, uint256 keyIndex)
func (_Resolver *ResolverCaller) PublicKey(opts *bind.CallOpts, node [32]byte, keyID string) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
	KeyIndex  *big.Int
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "publicKey", node, keyID)

	outstruct := new(struct {
		Status    uint8
		StartDate *big.Int
		EndDate   *big.Int
		KeyValue  string
		KeyIndex  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.KeyValue = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.KeyIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicKey is a free data retrieval call binding the contract method 0xa00efac0.
//
// Solidity: function publicKey(bytes32 node, string keyID) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue, uint256 keyIndex)
func (_Resolver *ResolverSession) PublicKey(node [32]byte, keyID string) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
	KeyIndex  *big.Int
}, error) {
	return _Resolver.Contract.PublicKey(&_Resolver.CallOpts, node, keyID)
}

// PublicKey is a free data retrieval call binding the contract method 0xa00efac0.
//
// Solidity: function publicKey(bytes32 node, string keyID) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue, uint256 keyIndex)
func (_Resolver *ResolverCallerSession) PublicKey(node [32]byte, keyID string) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
	KeyIndex  *big.Int
}, error) {
	return _Resolver.Contract.PublicKey(&_Resolver.CallOpts, node, keyID)
}

// PublicKeyAtIndex is a free data retrieval call binding the contract method 0x8c0083ea.
//
// Solidity: function publicKeyAtIndex(bytes32 node, string keyID, uint256 index) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue)
func (_Resolver *ResolverCaller) PublicKeyAtIndex(opts *bind.CallOpts, node [32]byte, keyID string, index *big.Int) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
}, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "publicKeyAtIndex", node, keyID, index)

	outstruct := new(struct {
		Status    uint8
		StartDate *big.Int
		EndDate   *big.Int
		KeyValue  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.KeyValue = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// PublicKeyAtIndex is a free data retrieval call binding the contract method 0x8c0083ea.
//
// Solidity: function publicKeyAtIndex(bytes32 node, string keyID, uint256 index) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue)
func (_Resolver *ResolverSession) PublicKeyAtIndex(node [32]byte, keyID string, index *big.Int) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
}, error) {
	return _Resolver.Contract.PublicKeyAtIndex(&_Resolver.CallOpts, node, keyID, index)
}

// PublicKeyAtIndex is a free data retrieval call binding the contract method 0x8c0083ea.
//
// Solidity: function publicKeyAtIndex(bytes32 node, string keyID, uint256 index) view returns(uint8 status, uint256 startDate, uint256 endDate, string keyValue)
func (_Resolver *ResolverCallerSession) PublicKeyAtIndex(node [32]byte, keyID string, index *big.Int) (struct {
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
	KeyValue  string
}, error) {
	return _Resolver.Contract.PublicKeyAtIndex(&_Resolver.CallOpts, node, keyID, index)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Resolver *ResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Resolver.Contract.SupportsInterface(&_Resolver.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _Resolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverSession) Text(node [32]byte, key string) (string, error) {
	return _Resolver.Contract.Text(&_Resolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Resolver *ResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _Resolver.Contract.Text(&_Resolver.CallOpts, node, key)
}

// AddAlaTSPService is a paid mutator transaction binding the contract method 0x212f2953.
//
// Solidity: function addAlaTSPService(bytes32 node, bytes32 label, string X509SKI, string serviceName, bytes X509Certificate, bool active) returns()
func (_Resolver *ResolverTransactor) AddAlaTSPService(opts *bind.TransactOpts, node [32]byte, label [32]byte, X509SKI string, serviceName string, X509Certificate []byte, active bool) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "addAlaTSPService", node, label, X509SKI, serviceName, X509Certificate, active)
}

// AddAlaTSPService is a paid mutator transaction binding the contract method 0x212f2953.
//
// Solidity: function addAlaTSPService(bytes32 node, bytes32 label, string X509SKI, string serviceName, bytes X509Certificate, bool active) returns()
func (_Resolver *ResolverSession) AddAlaTSPService(node [32]byte, label [32]byte, X509SKI string, serviceName string, X509Certificate []byte, active bool) (*types.Transaction, error) {
	return _Resolver.Contract.AddAlaTSPService(&_Resolver.TransactOpts, node, label, X509SKI, serviceName, X509Certificate, active)
}

// AddAlaTSPService is a paid mutator transaction binding the contract method 0x212f2953.
//
// Solidity: function addAlaTSPService(bytes32 node, bytes32 label, string X509SKI, string serviceName, bytes X509Certificate, bool active) returns()
func (_Resolver *ResolverTransactorSession) AddAlaTSPService(node [32]byte, label [32]byte, X509SKI string, serviceName string, X509Certificate []byte, active bool) (*types.Transaction, error) {
	return _Resolver.Contract.AddAlaTSPService(&_Resolver.TransactOpts, node, label, X509SKI, serviceName, X509Certificate, active)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xb1446a84.
//
// Solidity: function addParticipant(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverTransactor) AddParticipant(opts *bind.TransactOpts, node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "addParticipant", node, key, partDIDHash)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xb1446a84.
//
// Solidity: function addParticipant(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverSession) AddParticipant(node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.AddParticipant(&_Resolver.TransactOpts, node, key, partDIDHash)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xb1446a84.
//
// Solidity: function addParticipant(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverTransactorSession) AddParticipant(node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.AddParticipant(&_Resolver.TransactOpts, node, key, partDIDHash)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x73c3ce6c.
//
// Solidity: function addPublicKey(bytes32 node, string keyID, string keyValue) returns()
func (_Resolver *ResolverTransactor) AddPublicKey(opts *bind.TransactOpts, node [32]byte, keyID string, keyValue string) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "addPublicKey", node, keyID, keyValue)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x73c3ce6c.
//
// Solidity: function addPublicKey(bytes32 node, string keyID, string keyValue) returns()
func (_Resolver *ResolverSession) AddPublicKey(node [32]byte, keyID string, keyValue string) (*types.Transaction, error) {
	return _Resolver.Contract.AddPublicKey(&_Resolver.TransactOpts, node, keyID, keyValue)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x73c3ce6c.
//
// Solidity: function addPublicKey(bytes32 node, string keyID, string keyValue) returns()
func (_Resolver *ResolverTransactorSession) AddPublicKey(node [32]byte, keyID string, keyValue string) (*types.Transaction, error) {
	return _Resolver.Contract.AddPublicKey(&_Resolver.TransactOpts, node, keyID, keyValue)
}

// ConfirmCredential is a paid mutator transaction binding the contract method 0x9b985c3c.
//
// Solidity: function confirmCredential(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverTransactor) ConfirmCredential(opts *bind.TransactOpts, node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "confirmCredential", node, key, partDIDHash)
}

// ConfirmCredential is a paid mutator transaction binding the contract method 0x9b985c3c.
//
// Solidity: function confirmCredential(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverSession) ConfirmCredential(node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.ConfirmCredential(&_Resolver.TransactOpts, node, key, partDIDHash)
}

// ConfirmCredential is a paid mutator transaction binding the contract method 0x9b985c3c.
//
// Solidity: function confirmCredential(bytes32 node, string key, bytes32 partDIDHash) returns()
func (_Resolver *ResolverTransactorSession) ConfirmCredential(node [32]byte, key string, partDIDHash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.ConfirmCredential(&_Resolver.TransactOpts, node, key, partDIDHash)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0xe3f16cb5.
//
// Solidity: function deletePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverTransactor) DeletePublicKey(opts *bind.TransactOpts, node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "deletePublicKey", node, keyID, index)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0xe3f16cb5.
//
// Solidity: function deletePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverSession) DeletePublicKey(node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.Contract.DeletePublicKey(&_Resolver.TransactOpts, node, keyID, index)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0xe3f16cb5.
//
// Solidity: function deletePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverTransactorSession) DeletePublicKey(node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.Contract.DeletePublicKey(&_Resolver.TransactOpts, node, keyID, index)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x30a03491.
//
// Solidity: function revokePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverTransactor) RevokePublicKey(opts *bind.TransactOpts, node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "revokePublicKey", node, keyID, index)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x30a03491.
//
// Solidity: function revokePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverSession) RevokePublicKey(node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.Contract.RevokePublicKey(&_Resolver.TransactOpts, node, keyID, index)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x30a03491.
//
// Solidity: function revokePublicKey(bytes32 node, string keyID, uint256 index) returns()
func (_Resolver *ResolverTransactorSession) RevokePublicKey(node [32]byte, keyID string, index *big.Int) (*types.Transaction, error) {
	return _Resolver.Contract.RevokePublicKey(&_Resolver.TransactOpts, node, keyID, index)
}

// SetAlaDIDDocument is a paid mutator transaction binding the contract method 0xbe7dbe64.
//
// Solidity: function setAlaDIDDocument(bytes32 _DIDHash, string _DIDDocument) returns()
func (_Resolver *ResolverTransactor) SetAlaDIDDocument(opts *bind.TransactOpts, _DIDHash [32]byte, _DIDDocument string) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setAlaDIDDocument", _DIDHash, _DIDDocument)
}

// SetAlaDIDDocument is a paid mutator transaction binding the contract method 0xbe7dbe64.
//
// Solidity: function setAlaDIDDocument(bytes32 _DIDHash, string _DIDDocument) returns()
func (_Resolver *ResolverSession) SetAlaDIDDocument(_DIDHash [32]byte, _DIDDocument string) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaDIDDocument(&_Resolver.TransactOpts, _DIDHash, _DIDDocument)
}

// SetAlaDIDDocument is a paid mutator transaction binding the contract method 0xbe7dbe64.
//
// Solidity: function setAlaDIDDocument(bytes32 _DIDHash, string _DIDDocument) returns()
func (_Resolver *ResolverTransactorSession) SetAlaDIDDocument(_DIDHash [32]byte, _DIDDocument string) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaDIDDocument(&_Resolver.TransactOpts, _DIDHash, _DIDDocument)
}

// SetAlaDIDPublicEntity is a paid mutator transaction binding the contract method 0xb48322d0.
//
// Solidity: function setAlaDIDPublicEntity(bytes32 node, bytes32 label, bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active, address _owner) returns()
func (_Resolver *ResolverTransactor) SetAlaDIDPublicEntity(opts *bind.TransactOpts, node [32]byte, label [32]byte, _DIDHash [32]byte, _domain_name string, _DIDDocument string, _active bool, _owner common.Address) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setAlaDIDPublicEntity", node, label, _DIDHash, _domain_name, _DIDDocument, _active, _owner)
}

// SetAlaDIDPublicEntity is a paid mutator transaction binding the contract method 0xb48322d0.
//
// Solidity: function setAlaDIDPublicEntity(bytes32 node, bytes32 label, bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active, address _owner) returns()
func (_Resolver *ResolverSession) SetAlaDIDPublicEntity(node [32]byte, label [32]byte, _DIDHash [32]byte, _domain_name string, _DIDDocument string, _active bool, _owner common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaDIDPublicEntity(&_Resolver.TransactOpts, node, label, _DIDHash, _domain_name, _DIDDocument, _active, _owner)
}

// SetAlaDIDPublicEntity is a paid mutator transaction binding the contract method 0xb48322d0.
//
// Solidity: function setAlaDIDPublicEntity(bytes32 node, bytes32 label, bytes32 _DIDHash, string _domain_name, string _DIDDocument, bool _active, address _owner) returns()
func (_Resolver *ResolverTransactorSession) SetAlaDIDPublicEntity(node [32]byte, label [32]byte, _DIDHash [32]byte, _domain_name string, _DIDDocument string, _active bool, _owner common.Address) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaDIDPublicEntity(&_Resolver.TransactOpts, node, label, _DIDHash, _domain_name, _DIDDocument, _active, _owner)
}

// SetAlaTSP is a paid mutator transaction binding the contract method 0x5543eebd.
//
// Solidity: function setAlaTSP(bytes32 node, bytes32 label, string URI, string org, bool active) returns()
func (_Resolver *ResolverTransactor) SetAlaTSP(opts *bind.TransactOpts, node [32]byte, label [32]byte, URI string, org string, active bool) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setAlaTSP", node, label, URI, org, active)
}

// SetAlaTSP is a paid mutator transaction binding the contract method 0x5543eebd.
//
// Solidity: function setAlaTSP(bytes32 node, bytes32 label, string URI, string org, bool active) returns()
func (_Resolver *ResolverSession) SetAlaTSP(node [32]byte, label [32]byte, URI string, org string, active bool) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaTSP(&_Resolver.TransactOpts, node, label, URI, org, active)
}

// SetAlaTSP is a paid mutator transaction binding the contract method 0x5543eebd.
//
// Solidity: function setAlaTSP(bytes32 node, bytes32 label, string URI, string org, bool active) returns()
func (_Resolver *ResolverTransactorSession) SetAlaTSP(node [32]byte, label [32]byte, URI string, org string, active bool) (*types.Transaction, error) {
	return _Resolver.Contract.SetAlaTSP(&_Resolver.TransactOpts, node, label, URI, org, active)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_Resolver *ResolverTransactor) SetAuthorisation(opts *bind.TransactOpts, node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setAuthorisation", node, target, isAuthorised)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_Resolver *ResolverSession) SetAuthorisation(node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _Resolver.Contract.SetAuthorisation(&_Resolver.TransactOpts, node, target, isAuthorised)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_Resolver *ResolverTransactorSession) SetAuthorisation(node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _Resolver.Contract.SetAuthorisation(&_Resolver.TransactOpts, node, target, isAuthorised)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Resolver *ResolverTransactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Resolver *ResolverSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Resolver.Contract.SetContenthash(&_Resolver.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Resolver *ResolverTransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Resolver.Contract.SetContenthash(&_Resolver.TransactOpts, node, hash)
}

// SetCredential is a paid mutator transaction binding the contract method 0xb58dc321.
//
// Solidity: function setCredential(bytes32 node, string key, bytes credentialHash, bytes32 part1DIDhash, bytes32 part2DIDhash, bytes32 part3DIDhash, bytes32 part4DIDhash, bytes32 part5DIDhash) returns()
func (_Resolver *ResolverTransactor) SetCredential(opts *bind.TransactOpts, node [32]byte, key string, credentialHash []byte, part1DIDhash [32]byte, part2DIDhash [32]byte, part3DIDhash [32]byte, part4DIDhash [32]byte, part5DIDhash [32]byte) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setCredential", node, key, credentialHash, part1DIDhash, part2DIDhash, part3DIDhash, part4DIDhash, part5DIDhash)
}

// SetCredential is a paid mutator transaction binding the contract method 0xb58dc321.
//
// Solidity: function setCredential(bytes32 node, string key, bytes credentialHash, bytes32 part1DIDhash, bytes32 part2DIDhash, bytes32 part3DIDhash, bytes32 part4DIDhash, bytes32 part5DIDhash) returns()
func (_Resolver *ResolverSession) SetCredential(node [32]byte, key string, credentialHash []byte, part1DIDhash [32]byte, part2DIDhash [32]byte, part3DIDhash [32]byte, part4DIDhash [32]byte, part5DIDhash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.SetCredential(&_Resolver.TransactOpts, node, key, credentialHash, part1DIDhash, part2DIDhash, part3DIDhash, part4DIDhash, part5DIDhash)
}

// SetCredential is a paid mutator transaction binding the contract method 0xb58dc321.
//
// Solidity: function setCredential(bytes32 node, string key, bytes credentialHash, bytes32 part1DIDhash, bytes32 part2DIDhash, bytes32 part3DIDhash, bytes32 part4DIDhash, bytes32 part5DIDhash) returns()
func (_Resolver *ResolverTransactorSession) SetCredential(node [32]byte, key string, credentialHash []byte, part1DIDhash [32]byte, part2DIDhash [32]byte, part3DIDhash [32]byte, part4DIDhash [32]byte, part5DIDhash [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.SetCredential(&_Resolver.TransactOpts, node, key, credentialHash, part1DIDhash, part2DIDhash, part3DIDhash, part4DIDhash, part5DIDhash)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Resolver *ResolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Resolver *ResolverSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Resolver.Contract.SetName(&_Resolver.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Resolver *ResolverTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Resolver.Contract.SetName(&_Resolver.TransactOpts, node, name)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Resolver *ResolverTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Resolver *ResolverSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Resolver.Contract.SetText(&_Resolver.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Resolver *ResolverTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Resolver.Contract.SetText(&_Resolver.TransactOpts, node, key, value)
}

// ResolverAlaDIDDocumentChangedIterator is returned from FilterAlaDIDDocumentChanged and is used to iterate over the raw logs and unpacked data for AlaDIDDocumentChanged events raised by the Resolver contract.
type ResolverAlaDIDDocumentChangedIterator struct {
	Event *ResolverAlaDIDDocumentChanged // Event containing the contract specifics and raw log

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
func (it *ResolverAlaDIDDocumentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAlaDIDDocumentChanged)
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
		it.Event = new(ResolverAlaDIDDocumentChanged)
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
func (it *ResolverAlaDIDDocumentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAlaDIDDocumentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAlaDIDDocumentChanged represents a AlaDIDDocumentChanged event raised by the Resolver contract.
type ResolverAlaDIDDocumentChanged struct {
	Node     [32]byte
	Document string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAlaDIDDocumentChanged is a free log retrieval operation binding the contract event 0x72e481a357cf16af888285758647697856c4d59a91025de1c0fc85cf42547dbe.
//
// Solidity: event AlaDIDDocumentChanged(bytes32 indexed node, string document)
func (_Resolver *ResolverFilterer) FilterAlaDIDDocumentChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverAlaDIDDocumentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AlaDIDDocumentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAlaDIDDocumentChangedIterator{contract: _Resolver.contract, event: "AlaDIDDocumentChanged", logs: logs, sub: sub}, nil
}

var AlaDIDDocumentChangedTopicHash = "0x72e481a357cf16af888285758647697856c4d59a91025de1c0fc85cf42547dbe"

// WatchAlaDIDDocumentChanged is a free log subscription operation binding the contract event 0x72e481a357cf16af888285758647697856c4d59a91025de1c0fc85cf42547dbe.
//
// Solidity: event AlaDIDDocumentChanged(bytes32 indexed node, string document)
func (_Resolver *ResolverFilterer) WatchAlaDIDDocumentChanged(opts *bind.WatchOpts, sink chan<- *ResolverAlaDIDDocumentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AlaDIDDocumentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAlaDIDDocumentChanged)
				if err := _Resolver.contract.UnpackLog(event, "AlaDIDDocumentChanged", log); err != nil {
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

// ParseAlaDIDDocumentChanged is a log parse operation binding the contract event 0x72e481a357cf16af888285758647697856c4d59a91025de1c0fc85cf42547dbe.
//
// Solidity: event AlaDIDDocumentChanged(bytes32 indexed node, string document)
func (_Resolver *ResolverFilterer) ParseAlaDIDDocumentChanged(log types.Log) (*ResolverAlaDIDDocumentChanged, error) {
	event := new(ResolverAlaDIDDocumentChanged)
	if err := _Resolver.contract.UnpackLog(event, "AlaDIDDocumentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverAlaDIDPubkeyAddedIterator is returned from FilterAlaDIDPubkeyAdded and is used to iterate over the raw logs and unpacked data for AlaDIDPubkeyAdded events raised by the Resolver contract.
type ResolverAlaDIDPubkeyAddedIterator struct {
	Event *ResolverAlaDIDPubkeyAdded // Event containing the contract specifics and raw log

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
func (it *ResolverAlaDIDPubkeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAlaDIDPubkeyAdded)
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
		it.Event = new(ResolverAlaDIDPubkeyAdded)
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
func (it *ResolverAlaDIDPubkeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAlaDIDPubkeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAlaDIDPubkeyAdded represents a AlaDIDPubkeyAdded event raised by the Resolver contract.
type ResolverAlaDIDPubkeyAdded struct {
	Node       [32]byte
	IndexedKey common.Hash
	KeyID      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlaDIDPubkeyAdded is a free log retrieval operation binding the contract event 0x9e1efcecf97b4c9c969926e4bc9aede41e0610058cd84797715bce7ae5795f07.
//
// Solidity: event AlaDIDPubkeyAdded(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) FilterAlaDIDPubkeyAdded(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverAlaDIDPubkeyAddedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AlaDIDPubkeyAdded", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAlaDIDPubkeyAddedIterator{contract: _Resolver.contract, event: "AlaDIDPubkeyAdded", logs: logs, sub: sub}, nil
}

var AlaDIDPubkeyAddedTopicHash = "0x9e1efcecf97b4c9c969926e4bc9aede41e0610058cd84797715bce7ae5795f07"

// WatchAlaDIDPubkeyAdded is a free log subscription operation binding the contract event 0x9e1efcecf97b4c9c969926e4bc9aede41e0610058cd84797715bce7ae5795f07.
//
// Solidity: event AlaDIDPubkeyAdded(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) WatchAlaDIDPubkeyAdded(opts *bind.WatchOpts, sink chan<- *ResolverAlaDIDPubkeyAdded, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AlaDIDPubkeyAdded", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAlaDIDPubkeyAdded)
				if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyAdded", log); err != nil {
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

// ParseAlaDIDPubkeyAdded is a log parse operation binding the contract event 0x9e1efcecf97b4c9c969926e4bc9aede41e0610058cd84797715bce7ae5795f07.
//
// Solidity: event AlaDIDPubkeyAdded(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) ParseAlaDIDPubkeyAdded(log types.Log) (*ResolverAlaDIDPubkeyAdded, error) {
	event := new(ResolverAlaDIDPubkeyAdded)
	if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverAlaDIDPubkeyDeletedIterator is returned from FilterAlaDIDPubkeyDeleted and is used to iterate over the raw logs and unpacked data for AlaDIDPubkeyDeleted events raised by the Resolver contract.
type ResolverAlaDIDPubkeyDeletedIterator struct {
	Event *ResolverAlaDIDPubkeyDeleted // Event containing the contract specifics and raw log

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
func (it *ResolverAlaDIDPubkeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAlaDIDPubkeyDeleted)
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
		it.Event = new(ResolverAlaDIDPubkeyDeleted)
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
func (it *ResolverAlaDIDPubkeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAlaDIDPubkeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAlaDIDPubkeyDeleted represents a AlaDIDPubkeyDeleted event raised by the Resolver contract.
type ResolverAlaDIDPubkeyDeleted struct {
	Node       [32]byte
	IndexedKey common.Hash
	KeyID      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlaDIDPubkeyDeleted is a free log retrieval operation binding the contract event 0x1cc51114a966dfe16ba8f2c0ef40d5305de77c4dc2e0f3cdacb7f21e54c6206a.
//
// Solidity: event AlaDIDPubkeyDeleted(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) FilterAlaDIDPubkeyDeleted(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverAlaDIDPubkeyDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AlaDIDPubkeyDeleted", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAlaDIDPubkeyDeletedIterator{contract: _Resolver.contract, event: "AlaDIDPubkeyDeleted", logs: logs, sub: sub}, nil
}

var AlaDIDPubkeyDeletedTopicHash = "0x1cc51114a966dfe16ba8f2c0ef40d5305de77c4dc2e0f3cdacb7f21e54c6206a"

// WatchAlaDIDPubkeyDeleted is a free log subscription operation binding the contract event 0x1cc51114a966dfe16ba8f2c0ef40d5305de77c4dc2e0f3cdacb7f21e54c6206a.
//
// Solidity: event AlaDIDPubkeyDeleted(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) WatchAlaDIDPubkeyDeleted(opts *bind.WatchOpts, sink chan<- *ResolverAlaDIDPubkeyDeleted, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AlaDIDPubkeyDeleted", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAlaDIDPubkeyDeleted)
				if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyDeleted", log); err != nil {
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

// ParseAlaDIDPubkeyDeleted is a log parse operation binding the contract event 0x1cc51114a966dfe16ba8f2c0ef40d5305de77c4dc2e0f3cdacb7f21e54c6206a.
//
// Solidity: event AlaDIDPubkeyDeleted(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) ParseAlaDIDPubkeyDeleted(log types.Log) (*ResolverAlaDIDPubkeyDeleted, error) {
	event := new(ResolverAlaDIDPubkeyDeleted)
	if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverAlaDIDPubkeyRevokedIterator is returned from FilterAlaDIDPubkeyRevoked and is used to iterate over the raw logs and unpacked data for AlaDIDPubkeyRevoked events raised by the Resolver contract.
type ResolverAlaDIDPubkeyRevokedIterator struct {
	Event *ResolverAlaDIDPubkeyRevoked // Event containing the contract specifics and raw log

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
func (it *ResolverAlaDIDPubkeyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAlaDIDPubkeyRevoked)
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
		it.Event = new(ResolverAlaDIDPubkeyRevoked)
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
func (it *ResolverAlaDIDPubkeyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAlaDIDPubkeyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAlaDIDPubkeyRevoked represents a AlaDIDPubkeyRevoked event raised by the Resolver contract.
type ResolverAlaDIDPubkeyRevoked struct {
	Node       [32]byte
	IndexedKey common.Hash
	KeyID      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlaDIDPubkeyRevoked is a free log retrieval operation binding the contract event 0xede032e08b2e68ee190432bff5ae18127b098c506b59eb1cba23429f922f9ee6.
//
// Solidity: event AlaDIDPubkeyRevoked(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) FilterAlaDIDPubkeyRevoked(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverAlaDIDPubkeyRevokedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AlaDIDPubkeyRevoked", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAlaDIDPubkeyRevokedIterator{contract: _Resolver.contract, event: "AlaDIDPubkeyRevoked", logs: logs, sub: sub}, nil
}

var AlaDIDPubkeyRevokedTopicHash = "0xede032e08b2e68ee190432bff5ae18127b098c506b59eb1cba23429f922f9ee6"

// WatchAlaDIDPubkeyRevoked is a free log subscription operation binding the contract event 0xede032e08b2e68ee190432bff5ae18127b098c506b59eb1cba23429f922f9ee6.
//
// Solidity: event AlaDIDPubkeyRevoked(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) WatchAlaDIDPubkeyRevoked(opts *bind.WatchOpts, sink chan<- *ResolverAlaDIDPubkeyRevoked, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AlaDIDPubkeyRevoked", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAlaDIDPubkeyRevoked)
				if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyRevoked", log); err != nil {
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

// ParseAlaDIDPubkeyRevoked is a log parse operation binding the contract event 0xede032e08b2e68ee190432bff5ae18127b098c506b59eb1cba23429f922f9ee6.
//
// Solidity: event AlaDIDPubkeyRevoked(bytes32 indexed node, string indexed indexedKey, string keyID)
func (_Resolver *ResolverFilterer) ParseAlaDIDPubkeyRevoked(log types.Log) (*ResolverAlaDIDPubkeyRevoked, error) {
	event := new(ResolverAlaDIDPubkeyRevoked)
	if err := _Resolver.contract.UnpackLog(event, "AlaDIDPubkeyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverAuthorisationChangedIterator is returned from FilterAuthorisationChanged and is used to iterate over the raw logs and unpacked data for AuthorisationChanged events raised by the Resolver contract.
type ResolverAuthorisationChangedIterator struct {
	Event *ResolverAuthorisationChanged // Event containing the contract specifics and raw log

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
func (it *ResolverAuthorisationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAuthorisationChanged)
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
		it.Event = new(ResolverAuthorisationChanged)
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
func (it *ResolverAuthorisationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAuthorisationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAuthorisationChanged represents a AuthorisationChanged event raised by the Resolver contract.
type ResolverAuthorisationChanged struct {
	Node         [32]byte
	Owner        common.Address
	Target       common.Address
	IsAuthorised bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorisationChanged is a free log retrieval operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_Resolver *ResolverFilterer) FilterAuthorisationChanged(opts *bind.FilterOpts, node [][32]byte, owner []common.Address, target []common.Address) (*ResolverAuthorisationChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AuthorisationChanged", nodeRule, ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAuthorisationChangedIterator{contract: _Resolver.contract, event: "AuthorisationChanged", logs: logs, sub: sub}, nil
}

var AuthorisationChangedTopicHash = "0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df"

// WatchAuthorisationChanged is a free log subscription operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_Resolver *ResolverFilterer) WatchAuthorisationChanged(opts *bind.WatchOpts, sink chan<- *ResolverAuthorisationChanged, node [][32]byte, owner []common.Address, target []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AuthorisationChanged", nodeRule, ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAuthorisationChanged)
				if err := _Resolver.contract.UnpackLog(event, "AuthorisationChanged", log); err != nil {
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

// ParseAuthorisationChanged is a log parse operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_Resolver *ResolverFilterer) ParseAuthorisationChanged(log types.Log) (*ResolverAuthorisationChanged, error) {
	event := new(ResolverAuthorisationChanged)
	if err := _Resolver.contract.UnpackLog(event, "AuthorisationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the Resolver contract.
type ResolverContenthashChangedIterator struct {
	Event *ResolverContenthashChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContenthashChanged)
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
		it.Event = new(ResolverContenthashChanged)
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
func (it *ResolverContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContenthashChanged represents a ContenthashChanged event raised by the Resolver contract.
type ResolverContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Resolver *ResolverFilterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContenthashChangedIterator{contract: _Resolver.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

var ContenthashChangedTopicHash = "0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578"

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Resolver *ResolverFilterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *ResolverContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContenthashChanged)
				if err := _Resolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Resolver *ResolverFilterer) ParseContenthashChanged(log types.Log) (*ResolverContenthashChanged, error) {
	event := new(ResolverContenthashChanged)
	if err := _Resolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Resolver contract.
type ResolverNameChangedIterator struct {
	Event *ResolverNameChanged // Event containing the contract specifics and raw log

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
func (it *ResolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverNameChanged)
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
		it.Event = new(ResolverNameChanged)
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
func (it *ResolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverNameChanged represents a NameChanged event raised by the Resolver contract.
type ResolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Resolver *ResolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverNameChangedIterator{contract: _Resolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

var NameChangedTopicHash = "0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7"

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Resolver *ResolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ResolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverNameChanged)
				if err := _Resolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Resolver *ResolverFilterer) ParseNameChanged(log types.Log) (*ResolverNameChanged, error) {
	event := new(ResolverNameChanged)
	if err := _Resolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Resolver contract.
type ResolverTextChangedIterator struct {
	Event *ResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *ResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverTextChanged)
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
		it.Event = new(ResolverTextChanged)
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
func (it *ResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverTextChanged represents a TextChanged event raised by the Resolver contract.
type ResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverTextChangedIterator{contract: _Resolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

var TextChangedTopicHash = "0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverTextChanged)
				if err := _Resolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Resolver *ResolverFilterer) ParseTextChanged(log types.Log) (*ResolverTextChanged, error) {
	event := new(ResolverTextChanged)
	if err := _Resolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
