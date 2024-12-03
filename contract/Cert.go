// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Issued\",\"inputs\":[{\"name\":\"course\",\"type\":\"string\",\"components\":null,\"internalType\":null,\"indexed\":false},{\"name\":\"id\",\"type\":\"uint256\",\"components\":null,\"internalType\":null,\"indexed\":false},{\"name\":\"date\",\"type\":\"string\",\"components\":null,\"internalType\":null,\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"name\":\"issue_certificate\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"components\":null,\"internalType\":null},{\"name\":\"_name\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"_course\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"_grade\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"_date\",\"type\":\"string\",\"components\":null,\"internalType\":null}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"fetch_certificate\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"components\":null,\"internalType\":null}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"course\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"grade\",\"type\":\"string\",\"components\":null,\"internalType\":null},{\"name\":\"date\",\"type\":\"string\",\"components\":null,\"internalType\":null}],\"internalType\":null}]},{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]}]",
	Bin: "0x3461001a573360015561037461001e61000039610374610000f35b5f80fd5f3560e01c60026001821660011b61037001601e395f51565b632b9c2b7781186103685760a43610341761036c5760243560040180356018811161036c575060208135018082604037505060443560040180356012811161036c575060208135018082608037505060643560040180356001811161036c57506020813501808260c037505060843560040180356008811161036c57506020813501808261010037505060015433181561012d576020806101a052600d610140527f4163636573732044656e6965640000000000000000000000000000000000000061016052610140816101a00181518152602082015160208201528051806020830101601f825f03163682375050601f19601f8251602001011690509050810190506308c379a0610180528060040161019cfd5b60026004356020525f5260405f2060405181556060516001820155608051600282015560a051600160028301015560c051600482015560e0516001600483010155610100516006820155610120516001600683010155507fc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c606080610140528061014001608051815260a05160208201528051806020830101601f825f03163682375050601f19601f8251602001011690508101905060043561016052806101805280610140016101005181526101205160208201528051806020830101601f825f03163682375050601f19601f82516020010116905081019050610140a1005b63078aea0481186103685760243610341761036c5760208060405260026004356020525f5260405f2081604001608080825280820183548152600184015460208201528051806020830101601f825f03163682375050601f19601f825160200101169050810190508060208301526002830181830181548152600182015460208201528051806020830101601f825f03163682375050601f19601f8251602001011690509050810190508060408301526004830181830181548152600182015460208201528051806020830101601f825f03163682375050601f19601f8251602001011690509050810190508060608301526006830181830181548152600182015460208201528051806020830101601f825f03163682375050601f19601f82516020010116905090508101905090509050810190506040f35b5f5ffd5b5f80fd022e001884190374810400a1657679706572830004000014",
}

// CertABI is the input ABI used to generate the binding from.
// Deprecated: Use CertMetaData.ABI instead.
var CertABI = CertMetaData.ABI

// CertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertMetaData.Bin instead.
var CertBin = CertMetaData.Bin

// DeployCert deploys a new Ethereum contract, binding an instance of Cert to it.
func DeployCert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cert, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	CertCaller     // Read-only binding to the contract
	CertTransactor // Write-only binding to the contract
	CertFilterer   // Log filterer for contract events
}

// CertCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertSession struct {
	Contract     *Cert             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertCallerSession struct {
	Contract *CertCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertTransactorSession struct {
	Contract     *CertTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRaw struct {
	Contract *Cert // Generic contract binding to access the raw methods on
}

// CertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertCallerRaw struct {
	Contract *CertCaller // Generic read-only contract binding to access the raw methods on
}

// CertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertTransactorRaw struct {
	Contract *CertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCert creates a new instance of Cert, bound to a specific deployed contract.
func NewCert(address common.Address, backend bind.ContractBackend) (*Cert, error) {
	contract, err := bindCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// NewCertCaller creates a new read-only instance of Cert, bound to a specific deployed contract.
func NewCertCaller(address common.Address, caller bind.ContractCaller) (*CertCaller, error) {
	contract, err := bindCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertCaller{contract: contract}, nil
}

// NewCertTransactor creates a new write-only instance of Cert, bound to a specific deployed contract.
func NewCertTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTransactor, error) {
	contract, err := bindCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTransactor{contract: contract}, nil
}

// NewCertFilterer creates a new log filterer instance of Cert, bound to a specific deployed contract.
func NewCertFilterer(address common.Address, filterer bind.ContractFilterer) (*CertFilterer, error) {
	contract, err := bindCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertFilterer{contract: contract}, nil
}

// bindCert binds a generic wrapper to an already deployed contract.
func bindCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.CertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertCaller) FetchCertificate(opts *bind.CallOpts, _id *big.Int) (Struct0, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "fetch_certificate", _id)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertSession) FetchCertificate(_id *big.Int) (Struct0, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, _id)
}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertCallerSession) FetchCertificate(_id *big.Int) (Struct0, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, _id)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactor) IssueCertificate(opts *bind.TransactOpts, _id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue_certificate", _id, _name, _course, _grade, _date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertSession) IssueCertificate(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactorSession) IssueCertificate(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// CertIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Cert contract.
type CertIssuedIterator struct {
	Event *CertIssued // Event containing the contract specifics and raw log

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
func (it *CertIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertIssued)
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
		it.Event = new(CertIssued)
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
func (it *CertIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Course string
	Id     *big.Int
	Date   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) FilterIssued(opts *bind.FilterOpts) (*CertIssuedIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return &CertIssuedIterator{contract: _Cert.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *CertIssued) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertIssued)
				if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) ParseIssued(log types.Log) (*CertIssued, error) {
	event := new(CertIssued)
	if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
