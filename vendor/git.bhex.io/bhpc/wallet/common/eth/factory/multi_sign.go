/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: multi_sign.go
 * @Date: 2018/06/04
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// MultiSigWalletABI is the input ABI used to generate the binding from.
const MultiSigWalletABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"}]"

// MultiSigWalletBin is the compiled bytecode used for deploying new contracts.
const MultiSigWalletBin = `0x60806040523480156200001157600080fd5b506040516200115a3803806200115a83398101604052805160208201519101805190919060009082603282118015906200004b5750818111155b80156200005757508015155b80156200006357508115155b15156200006f57600080fd5b600092505b845183101562000147576002600086858151811015156200009157fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000e757508483815181101515620000cf57fe5b90602001906020020151600160a060020a0316600014155b1515620000f357600080fd5b60016002600087868151811015156200010857fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000074565b84516200015c9060039060208801906200016e565b50505060049190915550620002029050565b828054828255906000526020600020908101928215620001c6579160200282015b82811115620001c65782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906200018f565b50620001d4929150620001d8565b5090565b620001ff91905b80821115620001d4578054600160a060020a0319168155600101620001df565b90565b610f4880620002126000396000f3006080604052600436106100f05763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461013257806320ea8d86146101665780632f54bf6e1461017e5780633411c81c146101b357806354741525146101d7578063784547a7146102085780638b51d13f146102205780639ace38c214610238578063a0e67e2b146102f3578063a8abe69a14610358578063b5dc40c31461037d578063b77bf60014610395578063c01a8c84146103aa578063c6427474146103c2578063d74f8edd1461042b578063dc8452cd14610440578063ee22610b14610455575b60003411156101305760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b34801561013e57600080fd5b5061014a60043561046d565b60408051600160a060020a039092168252519081900360200190f35b34801561017257600080fd5b50610130600435610495565b34801561018a57600080fd5b5061019f600160a060020a036004351661054f565b604080519115158252519081900360200190f35b3480156101bf57600080fd5b5061019f600435600160a060020a0360243516610564565b3480156101e357600080fd5b506101f660043515156024351515610584565b60408051918252519081900360200190f35b34801561021457600080fd5b5061019f6004356105f0565b34801561022c57600080fd5b506101f6600435610674565b34801561024457600080fd5b506102506004356106e3565b6040518085600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156102b557818101518382015260200161029d565b50505050905090810190601f1680156102e25780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156102ff57600080fd5b506103086107a1565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561034457818101518382015260200161032c565b505050509050019250505060405180910390f35b34801561036457600080fd5b5061030860043560243560443515156064351515610804565b34801561038957600080fd5b5061030860043561093d565b3480156103a157600080fd5b506101f6610ab6565b3480156103b657600080fd5b50610130600435610abc565b3480156103ce57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f6948235600160a060020a0316946024803595369594606494920191908190840183828082843750949750610b879650505050505050565b34801561043757600080fd5b506101f6610ba6565b34801561044c57600080fd5b506101f6610bab565b34801561046157600080fd5b50610130600435610bb1565b600380548290811061047b57fe5b600091825260209091200154600160a060020a0316905081565b3360008181526002602052604090205460ff1615156104b357600080fd5b60008281526001602090815260408083203380855292529091205483919060ff1615156104df57600080fd5b600084815260208190526040902060030154849060ff161561050057600080fd5b6000858152600160209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b6005548110156105e9578380156105b1575060008181526020819052604090206003015460ff16155b806105d557508280156105d5575060008181526020819052604090206003015460ff165b156105e1576001820191505b600101610588565b5092915050565b600080805b60035481101561066d576000848152600160205260408120600380549192918490811061061e57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610652576001820191505b600454821415610665576001925061066d565b6001016105f5565b5050919050565b6000805b6003548110156106dd57600083815260016020526040812060038054919291849081106106a157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16156106d5576001820191505b600101610678565b50919050565b6000602081815291815260409081902080546001808301546002808501805487516101009582161595909502600019011691909104601f8101889004880284018801909652858352600160a060020a039093169590949192919083018282801561078e5780601f106107635761010080835404028352916020019161078e565b820191906000526020600020905b81548152906001019060200180831161077157829003601f168201915b5050506003909301549192505060ff1684565b606060038054806020026020016040519081016040528092919081815260200182805480156107f957602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107db575b505050505090505b90565b606080600080600554604051908082528060200260200182016040528015610836578160200160208202803883390190505b50925060009150600090505b6005548110156108bd5785801561086b575060008181526020819052604090206003015460ff16155b8061088f575084801561088f575060008181526020819052604090206003015460ff165b156108b5578083838151811015156108a357fe5b60209081029091010152600191909101905b600101610842565b8787036040519080825280602002602001820160405280156108e9578160200160208202803883390190505b5093508790505b8681101561093257828181518110151561090657fe5b906020019060200201518489830381518110151561092057fe5b602090810290910101526001016108f0565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610972578160200160208202803883390190505b50925060009150600090505b600354811015610a2f57600085815260016020526040812060038054919291849081106109a757fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610a275760038054829081106109e257fe5b6000918252602090912001548351600160a060020a0390911690849084908110610a0857fe5b600160a060020a03909216602092830290910190910152600191909101905b60010161097e565b81604051908082528060200260200182016040528015610a59578160200160208202803883390190505b509350600090505b81811015610aae578281815181101515610a7757fe5b906020019060200201518482815181101515610a8f57fe5b600160a060020a03909216602092830290910190910152600101610a61565b505050919050565b60055481565b3360008181526002602052604090205460ff161515610ada57600080fd5b6000828152602081905260409020548290600160a060020a03161515610aff57600080fd5b60008381526001602090815260408083203380855292529091205484919060ff1615610b2a57600080fd5b6000858152600160208181526040808420338086529252808420805460ff1916909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a3610b8085610bb1565b5050505050565b6000610b94848484610d71565b9050610b9f81610abc565b9392505050565b603281565b60045481565b3360008181526002602052604081205490919060ff161515610bd257600080fd5b60008381526001602090815260408083203380855292529091205484919060ff161515610bfe57600080fd5b600085815260208190526040902060030154859060ff1615610c1f57600080fd5b610c28866105f0565b15610d69576000868152602081815260409182902060038101805460ff19166001908117909155815481830154600280850180548851601f60001997831615610100029790970190911692909204948501879004870282018701909752838152939a50610cfc95600160a060020a0390921694909391908390830182828015610cf25780601f10610cc757610100808354040283529160200191610cf2565b820191906000526020600020905b815481529060010190602001808311610cd557829003601f168201915b5050505050610e61565b15610d315760405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a2610d69565b60405186907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a260038501805460ff191690555b505050505050565b600083600160a060020a0381161515610d8957600080fd5b60055460408051608081018252600160a060020a0388811682526020808301898152838501898152600060608601819052878152808452959095208451815473ffffffffffffffffffffffffffffffffffffffff191694169390931783555160018301559251805194965091939092610e09926002850192910190610e84565b50606091909101516003909101805460ff191691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b6000806040516020840160008287838a8c6187965a03f198975050505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ec557805160ff1916838001178555610ef2565b82800160010185558215610ef2579182015b82811115610ef2578251825591602001919060010190610ed7565b50610efe929150610f02565b5090565b61080191905b80821115610efe5760008155600101610f085600a165627a7a7230582010952d0336df6f52898893a6f7f04db0fdeeeb16ac52169323f3c89879ea87be0029`

// DeployMultiSigWallet deploys a new Ethereum contract, binding an instance of MultiSigWallet to it.
func DeployMultiSigWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int) (common.Address, *types.Transaction, *MultiSigWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletBin), backend, _owners, _required)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// MultiSigWallet is an auto generated Go binding around an Ethereum contract.
type MultiSigWallet struct {
	MultiSigWalletCaller     // Read-only binding to the contract
	MultiSigWalletTransactor // Write-only binding to the contract
	MultiSigWalletFilterer   // Log filterer for contract events
}

// MultiSigWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletSession struct {
	Contract     *MultiSigWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSigWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletCallerSession struct {
	Contract *MultiSigWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultiSigWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletTransactorSession struct {
	Contract     *MultiSigWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultiSigWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletRaw struct {
	Contract *MultiSigWallet // Generic contract binding to access the raw methods on
}

// MultiSigWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletCallerRaw struct {
	Contract *MultiSigWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactorRaw struct {
	Contract *MultiSigWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWallet creates a new instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWallet(address common.Address, backend bind.ContractBackend) (*MultiSigWallet, error) {
	contract, err := bindMultiSigWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// NewMultiSigWalletCaller creates a new read-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletCaller, error) {
	contract, err := bindMultiSigWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletCaller{contract: contract}, nil
}

// NewMultiSigWalletTransactor creates a new write-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletTransactor, error) {
	contract, err := bindMultiSigWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletTransactor{contract: contract}, nil
}

// NewMultiSigWalletFilterer creates a new log filterer instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletFilterer, error) {
	contract, err := bindMultiSigWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFilterer{contract: contract}, nil
}

// bindMultiSigWallet binds a generic wrapper to an already deployed contract.
func bindMultiSigWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.MultiSigWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWallet *MultiSigWalletCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiSigWallet *MultiSigWalletCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiSigWallet *MultiSigWalletSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiSigWallet *MultiSigWalletCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWallet.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MultiSigWallet *MultiSigWalletCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MultiSigWallet.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MultiSigWallet *MultiSigWalletSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MultiSigWallet *MultiSigWalletTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MultiSigWallet *MultiSigWalletSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MultiSigWallet *MultiSigWalletTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// MultiSigWalletConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWallet contract.
type MultiSigWalletConfirmationIterator struct {
	Event *MultiSigWalletConfirmation // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletConfirmation)
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
		it.Event = new(MultiSigWalletConfirmation)
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
func (it *MultiSigWalletConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletConfirmation represents a Confirmation event raised by the MultiSigWallet contract.
type MultiSigWalletConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletConfirmationIterator{contract: _MultiSigWallet.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletConfirmation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// MultiSigWalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWallet contract.
type MultiSigWalletDepositIterator struct {
	Event *MultiSigWalletDeposit // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletDeposit)
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
		it.Event = new(MultiSigWalletDeposit)
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
func (it *MultiSigWalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletDeposit represents a Deposit event raised by the MultiSigWallet contract.
type MultiSigWalletDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(sender indexed address, value uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletDepositIterator{contract: _MultiSigWallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(sender indexed address, value uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletDeposit)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// MultiSigWalletExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionIterator struct {
	Event *MultiSigWalletExecution // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecution)
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
		it.Event = new(MultiSigWalletExecution)
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
func (it *MultiSigWalletExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecution represents a Execution event raised by the MultiSigWallet contract.
type MultiSigWalletExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionIterator{contract: _MultiSigWallet.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecution)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Execution", log); err != nil {
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

// MultiSigWalletExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailureIterator struct {
	Event *MultiSigWalletExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecutionFailure)
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
		it.Event = new(MultiSigWalletExecutionFailure)
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
func (it *MultiSigWalletExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecutionFailure represents a ExecutionFailure event raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionFailureIterator{contract: _MultiSigWallet.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecutionFailure)
				if err := _MultiSigWallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// MultiSigWalletRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWallet contract.
type MultiSigWalletRevocationIterator struct {
	Event *MultiSigWalletRevocation // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletRevocation)
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
		it.Event = new(MultiSigWalletRevocation)
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
func (it *MultiSigWalletRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletRevocation represents a Revocation event raised by the MultiSigWallet contract.
type MultiSigWalletRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletRevocationIterator{contract: _MultiSigWallet.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletRevocation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// MultiSigWalletSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWallet contract.
type MultiSigWalletSubmissionIterator struct {
	Event *MultiSigWalletSubmission // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletSubmission)
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
		it.Event = new(MultiSigWalletSubmission)
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
func (it *MultiSigWalletSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletSubmission represents a Submission event raised by the MultiSigWallet contract.
type MultiSigWalletSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletSubmissionIterator{contract: _MultiSigWallet.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletSubmission)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Submission", log); err != nil {
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
