// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookreview

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

// BookReviewBook is an auto generated low-level Go binding around an user-defined struct.
type BookReviewBook struct {
	Name   string
	Author string
}

// BookReviewReview is an auto generated low-level Go binding around an user-defined struct.
type BookReviewReview struct {
	BookId  *big.Int
	Content string
}

// BookreviewMetaData contains all meta data concerning the Bookreview contract.
var BookreviewMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"}],\"name\":\"NewBook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reviewId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contenxt\",\"type\":\"string\"}],\"name\":\"NewReview\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"books\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"findBook\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"}],\"internalType\":\"structBookReview.Book[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reviewId\",\"type\":\"uint256\"}],\"name\":\"findReview\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"internalType\":\"structBookReview.Review\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_author\",\"type\":\"string\"}],\"name\":\"registerBook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bookId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"registerReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reviews\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bookId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506110ce806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806306a83e0914610067578063195a011914610083578063610241bb1461009f57806368744046146100cf578063a873f82814610100578063e83ddcea1461011e575b600080fd5b610081600480360381019061007c91906108cf565b61014f565b005b61009d6004803603810190610098919061092b565b610206565b005b6100b960048036038101906100b491906109a3565b6102ca565b6040516100c69190610a9b565b60405180910390f35b6100e960048036038101906100e491906109a3565b6103a2565b6040516100f7929190610b07565b60405180910390f35b6101086104e6565b6040516101159190610c44565b60405180910390f35b610138600480360381019061013391906109a3565b610669565b604051610146929190610c75565b60405180910390f35b600160405180604001604052808481526020018381525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010190816101ae9190610eb1565b5050506000600180805490506101c49190610fb2565b90507f664e6e86d3202d6c94bfaba88bc4520e137f0c5c3bad1b6b668db2ad1b8ca38d8184846040516101f993929190610fe6565b60405180910390a1505050565b60006040518060400160405280848152602001838152509080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001908161025b9190610eb1565b5060208201518160010190816102719190610eb1565b505050600060016000805490506102889190610fb2565b90507f89f52bd49e0e4531790e61d745759c87e09ad49b706dcbc7815e80b1e565a0258184846040516102bd93929190611024565b60405180910390a1505050565b6102d2610725565b600182815481106102e6576102e5611069565b5b90600052602060002090600202016040518060400160405290816000820154815260200160018201805461031990610cd4565b80601f016020809104026020016040519081016040528092919081815260200182805461034590610cd4565b80156103925780601f1061036757610100808354040283529160200191610392565b820191906000526020600020905b81548152906001019060200180831161037557829003601f168201915b5050505050815250509050919050565b600081815481106103b257600080fd5b90600052602060002090600202016000915090508060000180546103d590610cd4565b80601f016020809104026020016040519081016040528092919081815260200182805461040190610cd4565b801561044e5780601f106104235761010080835404028352916020019161044e565b820191906000526020600020905b81548152906001019060200180831161043157829003601f168201915b50505050509080600101805461046390610cd4565b80601f016020809104026020016040519081016040528092919081815260200182805461048f90610cd4565b80156104dc5780601f106104b1576101008083540402835291602001916104dc565b820191906000526020600020905b8154815290600101906020018083116104bf57829003601f168201915b5050505050905082565b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610660578382906000526020600020906002020160405180604001604052908160008201805461053d90610cd4565b80601f016020809104026020016040519081016040528092919081815260200182805461056990610cd4565b80156105b65780601f1061058b576101008083540402835291602001916105b6565b820191906000526020600020905b81548152906001019060200180831161059957829003601f168201915b505050505081526020016001820180546105cf90610cd4565b80601f01602080910402602001604051908101604052809291908181526020018280546105fb90610cd4565b80156106485780601f1061061d57610100808354040283529160200191610648565b820191906000526020600020905b81548152906001019060200180831161062b57829003601f168201915b5050505050815250508152602001906001019061050a565b50505050905090565b6001818154811061067957600080fd5b90600052602060002090600202016000915090508060000154908060010180546106a290610cd4565b80601f01602080910402602001604051908101604052809291908181526020018280546106ce90610cd4565b801561071b5780601f106106f05761010080835404028352916020019161071b565b820191906000526020600020905b8154815290600101906020018083116106fe57829003601f168201915b5050505050905082565b604051806040016040528060008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61076681610753565b811461077157600080fd5b50565b6000813590506107838161075d565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107dc82610793565b810181811067ffffffffffffffff821117156107fb576107fa6107a4565b5b80604052505050565b600061080e61073f565b905061081a82826107d3565b919050565b600067ffffffffffffffff82111561083a576108396107a4565b5b61084382610793565b9050602081019050919050565b82818337600083830152505050565b600061087261086d8461081f565b610804565b90508281526020810184848401111561088e5761088d61078e565b5b610899848285610850565b509392505050565b600082601f8301126108b6576108b5610789565b5b81356108c684826020860161085f565b91505092915050565b600080604083850312156108e6576108e5610749565b5b60006108f485828601610774565b925050602083013567ffffffffffffffff8111156109155761091461074e565b5b610921858286016108a1565b9150509250929050565b6000806040838503121561094257610941610749565b5b600083013567ffffffffffffffff8111156109605761095f61074e565b5b61096c858286016108a1565b925050602083013567ffffffffffffffff81111561098d5761098c61074e565b5b610999858286016108a1565b9150509250929050565b6000602082840312156109b9576109b8610749565b5b60006109c784828501610774565b91505092915050565b6109d981610753565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a195780820151818401526020810190506109fe565b60008484015250505050565b6000610a30826109df565b610a3a81856109ea565b9350610a4a8185602086016109fb565b610a5381610793565b840191505092915050565b6000604083016000830151610a7660008601826109d0565b5060208301518482036020860152610a8e8282610a25565b9150508091505092915050565b60006020820190508181036000830152610ab58184610a5e565b905092915050565b600082825260208201905092915050565b6000610ad9826109df565b610ae38185610abd565b9350610af38185602086016109fb565b610afc81610793565b840191505092915050565b60006040820190508181036000830152610b218185610ace565b90508181036020830152610b358184610ace565b90509392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006040830160008301518482036000860152610b878282610a25565b91505060208301518482036020860152610ba18282610a25565b9150508091505092915050565b6000610bba8383610b6a565b905092915050565b6000602082019050919050565b6000610bda82610b3e565b610be48185610b49565b935083602082028501610bf685610b5a565b8060005b85811015610c325784840389528151610c138582610bae565b9450610c1e83610bc2565b925060208a01995050600181019050610bfa565b50829750879550505050505092915050565b60006020820190508181036000830152610c5e8184610bcf565b905092915050565b610c6f81610753565b82525050565b6000604082019050610c8a6000830185610c66565b8181036020830152610c9c8184610ace565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610cec57607f821691505b602082108103610cff57610cfe610ca5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610d677fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d2a565b610d718683610d2a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610dae610da9610da484610753565b610d89565b610753565b9050919050565b6000819050919050565b610dc883610d93565b610ddc610dd482610db5565b848454610d37565b825550505050565b600090565b610df1610de4565b610dfc818484610dbf565b505050565b5b81811015610e2057610e15600082610de9565b600181019050610e02565b5050565b601f821115610e6557610e3681610d05565b610e3f84610d1a565b81016020851015610e4e578190505b610e62610e5a85610d1a565b830182610e01565b50505b505050565b600082821c905092915050565b6000610e8860001984600802610e6a565b1980831691505092915050565b6000610ea18383610e77565b9150826002028217905092915050565b610eba826109df565b67ffffffffffffffff811115610ed357610ed26107a4565b5b610edd8254610cd4565b610ee8828285610e24565b600060209050601f831160018114610f1b5760008415610f09578287015190505b610f138582610e95565b865550610f7b565b601f198416610f2986610d05565b60005b82811015610f5157848901518255600182019150602085019450602081019050610f2c565b86831015610f6e5784890151610f6a601f891682610e77565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610fbd82610753565b9150610fc883610753565b9250828203905081811115610fe057610fdf610f83565b5b92915050565b6000606082019050610ffb6000830186610c66565b6110086020830185610c66565b818103604083015261101a8184610ace565b9050949350505050565b60006060820190506110396000830186610c66565b818103602083015261104b8185610ace565b9050818103604083015261105f8184610ace565b9050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220816a8b71a07eb9a55753e155214c0dcd8ed6c94bf5e97f8b4aa9b4b5657ffcba64736f6c63430008130033",
}

// BookreviewABI is the input ABI used to generate the binding from.
// Deprecated: Use BookreviewMetaData.ABI instead.
var BookreviewABI = BookreviewMetaData.ABI

// BookreviewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BookreviewMetaData.Bin instead.
var BookreviewBin = BookreviewMetaData.Bin

// DeployBookreview deploys a new Ethereum contract, binding an instance of Bookreview to it.
func DeployBookreview(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookreview, error) {
	parsed, err := BookreviewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BookreviewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookreview{BookreviewCaller: BookreviewCaller{contract: contract}, BookreviewTransactor: BookreviewTransactor{contract: contract}, BookreviewFilterer: BookreviewFilterer{contract: contract}}, nil
}

// Bookreview is an auto generated Go binding around an Ethereum contract.
type Bookreview struct {
	BookreviewCaller     // Read-only binding to the contract
	BookreviewTransactor // Write-only binding to the contract
	BookreviewFilterer   // Log filterer for contract events
}

// BookreviewCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookreviewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookreviewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookreviewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookreviewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookreviewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookreviewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookreviewSession struct {
	Contract     *Bookreview       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookreviewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookreviewCallerSession struct {
	Contract *BookreviewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BookreviewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookreviewTransactorSession struct {
	Contract     *BookreviewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BookreviewRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookreviewRaw struct {
	Contract *Bookreview // Generic contract binding to access the raw methods on
}

// BookreviewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookreviewCallerRaw struct {
	Contract *BookreviewCaller // Generic read-only contract binding to access the raw methods on
}

// BookreviewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookreviewTransactorRaw struct {
	Contract *BookreviewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookreview creates a new instance of Bookreview, bound to a specific deployed contract.
func NewBookreview(address common.Address, backend bind.ContractBackend) (*Bookreview, error) {
	contract, err := bindBookreview(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookreview{BookreviewCaller: BookreviewCaller{contract: contract}, BookreviewTransactor: BookreviewTransactor{contract: contract}, BookreviewFilterer: BookreviewFilterer{contract: contract}}, nil
}

// NewBookreviewCaller creates a new read-only instance of Bookreview, bound to a specific deployed contract.
func NewBookreviewCaller(address common.Address, caller bind.ContractCaller) (*BookreviewCaller, error) {
	contract, err := bindBookreview(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookreviewCaller{contract: contract}, nil
}

// NewBookreviewTransactor creates a new write-only instance of Bookreview, bound to a specific deployed contract.
func NewBookreviewTransactor(address common.Address, transactor bind.ContractTransactor) (*BookreviewTransactor, error) {
	contract, err := bindBookreview(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookreviewTransactor{contract: contract}, nil
}

// NewBookreviewFilterer creates a new log filterer instance of Bookreview, bound to a specific deployed contract.
func NewBookreviewFilterer(address common.Address, filterer bind.ContractFilterer) (*BookreviewFilterer, error) {
	contract, err := bindBookreview(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookreviewFilterer{contract: contract}, nil
}

// bindBookreview binds a generic wrapper to an already deployed contract.
func bindBookreview(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookreviewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookreview *BookreviewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookreview.Contract.BookreviewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookreview *BookreviewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookreview.Contract.BookreviewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookreview *BookreviewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookreview.Contract.BookreviewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookreview *BookreviewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookreview.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookreview *BookreviewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookreview.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookreview *BookreviewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookreview.Contract.contract.Transact(opts, method, params...)
}

// Books is a free data retrieval call binding the contract method 0x68744046.
//
// Solidity: function books(uint256 ) view returns(string name, string author)
func (_Bookreview *BookreviewCaller) Books(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Author string
}, error) {
	var out []interface{}
	err := _Bookreview.contract.Call(opts, &out, "books", arg0)

	outstruct := new(struct {
		Name   string
		Author string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Author = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Books is a free data retrieval call binding the contract method 0x68744046.
//
// Solidity: function books(uint256 ) view returns(string name, string author)
func (_Bookreview *BookreviewSession) Books(arg0 *big.Int) (struct {
	Name   string
	Author string
}, error) {
	return _Bookreview.Contract.Books(&_Bookreview.CallOpts, arg0)
}

// Books is a free data retrieval call binding the contract method 0x68744046.
//
// Solidity: function books(uint256 ) view returns(string name, string author)
func (_Bookreview *BookreviewCallerSession) Books(arg0 *big.Int) (struct {
	Name   string
	Author string
}, error) {
	return _Bookreview.Contract.Books(&_Bookreview.CallOpts, arg0)
}

// FindBook is a free data retrieval call binding the contract method 0xa873f828.
//
// Solidity: function findBook() view returns((string,string)[])
func (_Bookreview *BookreviewCaller) FindBook(opts *bind.CallOpts) ([]BookReviewBook, error) {
	var out []interface{}
	err := _Bookreview.contract.Call(opts, &out, "findBook")

	if err != nil {
		return *new([]BookReviewBook), err
	}

	out0 := *abi.ConvertType(out[0], new([]BookReviewBook)).(*[]BookReviewBook)

	return out0, err

}

// FindBook is a free data retrieval call binding the contract method 0xa873f828.
//
// Solidity: function findBook() view returns((string,string)[])
func (_Bookreview *BookreviewSession) FindBook() ([]BookReviewBook, error) {
	return _Bookreview.Contract.FindBook(&_Bookreview.CallOpts)
}

// FindBook is a free data retrieval call binding the contract method 0xa873f828.
//
// Solidity: function findBook() view returns((string,string)[])
func (_Bookreview *BookreviewCallerSession) FindBook() ([]BookReviewBook, error) {
	return _Bookreview.Contract.FindBook(&_Bookreview.CallOpts)
}

// FindReview is a free data retrieval call binding the contract method 0x610241bb.
//
// Solidity: function findReview(uint256 _reviewId) view returns((uint256,string))
func (_Bookreview *BookreviewCaller) FindReview(opts *bind.CallOpts, _reviewId *big.Int) (BookReviewReview, error) {
	var out []interface{}
	err := _Bookreview.contract.Call(opts, &out, "findReview", _reviewId)

	if err != nil {
		return *new(BookReviewReview), err
	}

	out0 := *abi.ConvertType(out[0], new(BookReviewReview)).(*BookReviewReview)

	return out0, err

}

// FindReview is a free data retrieval call binding the contract method 0x610241bb.
//
// Solidity: function findReview(uint256 _reviewId) view returns((uint256,string))
func (_Bookreview *BookreviewSession) FindReview(_reviewId *big.Int) (BookReviewReview, error) {
	return _Bookreview.Contract.FindReview(&_Bookreview.CallOpts, _reviewId)
}

// FindReview is a free data retrieval call binding the contract method 0x610241bb.
//
// Solidity: function findReview(uint256 _reviewId) view returns((uint256,string))
func (_Bookreview *BookreviewCallerSession) FindReview(_reviewId *big.Int) (BookReviewReview, error) {
	return _Bookreview.Contract.FindReview(&_Bookreview.CallOpts, _reviewId)
}

// Reviews is a free data retrieval call binding the contract method 0xe83ddcea.
//
// Solidity: function reviews(uint256 ) view returns(uint256 bookId, string content)
func (_Bookreview *BookreviewCaller) Reviews(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BookId  *big.Int
	Content string
}, error) {
	var out []interface{}
	err := _Bookreview.contract.Call(opts, &out, "reviews", arg0)

	outstruct := new(struct {
		BookId  *big.Int
		Content string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BookId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Reviews is a free data retrieval call binding the contract method 0xe83ddcea.
//
// Solidity: function reviews(uint256 ) view returns(uint256 bookId, string content)
func (_Bookreview *BookreviewSession) Reviews(arg0 *big.Int) (struct {
	BookId  *big.Int
	Content string
}, error) {
	return _Bookreview.Contract.Reviews(&_Bookreview.CallOpts, arg0)
}

// Reviews is a free data retrieval call binding the contract method 0xe83ddcea.
//
// Solidity: function reviews(uint256 ) view returns(uint256 bookId, string content)
func (_Bookreview *BookreviewCallerSession) Reviews(arg0 *big.Int) (struct {
	BookId  *big.Int
	Content string
}, error) {
	return _Bookreview.Contract.Reviews(&_Bookreview.CallOpts, arg0)
}

// RegisterBook is a paid mutator transaction binding the contract method 0x195a0119.
//
// Solidity: function registerBook(string _name, string _author) returns()
func (_Bookreview *BookreviewTransactor) RegisterBook(opts *bind.TransactOpts, _name string, _author string) (*types.Transaction, error) {
	return _Bookreview.contract.Transact(opts, "registerBook", _name, _author)
}

// RegisterBook is a paid mutator transaction binding the contract method 0x195a0119.
//
// Solidity: function registerBook(string _name, string _author) returns()
func (_Bookreview *BookreviewSession) RegisterBook(_name string, _author string) (*types.Transaction, error) {
	return _Bookreview.Contract.RegisterBook(&_Bookreview.TransactOpts, _name, _author)
}

// RegisterBook is a paid mutator transaction binding the contract method 0x195a0119.
//
// Solidity: function registerBook(string _name, string _author) returns()
func (_Bookreview *BookreviewTransactorSession) RegisterBook(_name string, _author string) (*types.Transaction, error) {
	return _Bookreview.Contract.RegisterBook(&_Bookreview.TransactOpts, _name, _author)
}

// RegisterReview is a paid mutator transaction binding the contract method 0x06a83e09.
//
// Solidity: function registerReview(uint256 _bookId, string _content) returns()
func (_Bookreview *BookreviewTransactor) RegisterReview(opts *bind.TransactOpts, _bookId *big.Int, _content string) (*types.Transaction, error) {
	return _Bookreview.contract.Transact(opts, "registerReview", _bookId, _content)
}

// RegisterReview is a paid mutator transaction binding the contract method 0x06a83e09.
//
// Solidity: function registerReview(uint256 _bookId, string _content) returns()
func (_Bookreview *BookreviewSession) RegisterReview(_bookId *big.Int, _content string) (*types.Transaction, error) {
	return _Bookreview.Contract.RegisterReview(&_Bookreview.TransactOpts, _bookId, _content)
}

// RegisterReview is a paid mutator transaction binding the contract method 0x06a83e09.
//
// Solidity: function registerReview(uint256 _bookId, string _content) returns()
func (_Bookreview *BookreviewTransactorSession) RegisterReview(_bookId *big.Int, _content string) (*types.Transaction, error) {
	return _Bookreview.Contract.RegisterReview(&_Bookreview.TransactOpts, _bookId, _content)
}

// BookreviewNewBookIterator is returned from FilterNewBook and is used to iterate over the raw logs and unpacked data for NewBook events raised by the Bookreview contract.
type BookreviewNewBookIterator struct {
	Event *BookreviewNewBook // Event containing the contract specifics and raw log

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
func (it *BookreviewNewBookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookreviewNewBook)
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
		it.Event = new(BookreviewNewBook)
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
func (it *BookreviewNewBookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookreviewNewBookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookreviewNewBook represents a NewBook event raised by the Bookreview contract.
type BookreviewNewBook struct {
	BookId *big.Int
	Name   string
	Author string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewBook is a free log retrieval operation binding the contract event 0x89f52bd49e0e4531790e61d745759c87e09ad49b706dcbc7815e80b1e565a025.
//
// Solidity: event NewBook(uint256 bookId, string name, string author)
func (_Bookreview *BookreviewFilterer) FilterNewBook(opts *bind.FilterOpts) (*BookreviewNewBookIterator, error) {

	logs, sub, err := _Bookreview.contract.FilterLogs(opts, "NewBook")
	if err != nil {
		return nil, err
	}
	return &BookreviewNewBookIterator{contract: _Bookreview.contract, event: "NewBook", logs: logs, sub: sub}, nil
}

// WatchNewBook is a free log subscription operation binding the contract event 0x89f52bd49e0e4531790e61d745759c87e09ad49b706dcbc7815e80b1e565a025.
//
// Solidity: event NewBook(uint256 bookId, string name, string author)
func (_Bookreview *BookreviewFilterer) WatchNewBook(opts *bind.WatchOpts, sink chan<- *BookreviewNewBook) (event.Subscription, error) {

	logs, sub, err := _Bookreview.contract.WatchLogs(opts, "NewBook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookreviewNewBook)
				if err := _Bookreview.contract.UnpackLog(event, "NewBook", log); err != nil {
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

// ParseNewBook is a log parse operation binding the contract event 0x89f52bd49e0e4531790e61d745759c87e09ad49b706dcbc7815e80b1e565a025.
//
// Solidity: event NewBook(uint256 bookId, string name, string author)
func (_Bookreview *BookreviewFilterer) ParseNewBook(log types.Log) (*BookreviewNewBook, error) {
	event := new(BookreviewNewBook)
	if err := _Bookreview.contract.UnpackLog(event, "NewBook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookreviewNewReviewIterator is returned from FilterNewReview and is used to iterate over the raw logs and unpacked data for NewReview events raised by the Bookreview contract.
type BookreviewNewReviewIterator struct {
	Event *BookreviewNewReview // Event containing the contract specifics and raw log

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
func (it *BookreviewNewReviewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookreviewNewReview)
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
		it.Event = new(BookreviewNewReview)
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
func (it *BookreviewNewReviewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookreviewNewReviewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookreviewNewReview represents a NewReview event raised by the Bookreview contract.
type BookreviewNewReview struct {
	ReviewId *big.Int
	BookId   *big.Int
	Contenxt string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewReview is a free log retrieval operation binding the contract event 0x664e6e86d3202d6c94bfaba88bc4520e137f0c5c3bad1b6b668db2ad1b8ca38d.
//
// Solidity: event NewReview(uint256 reviewId, uint256 bookId, string contenxt)
func (_Bookreview *BookreviewFilterer) FilterNewReview(opts *bind.FilterOpts) (*BookreviewNewReviewIterator, error) {

	logs, sub, err := _Bookreview.contract.FilterLogs(opts, "NewReview")
	if err != nil {
		return nil, err
	}
	return &BookreviewNewReviewIterator{contract: _Bookreview.contract, event: "NewReview", logs: logs, sub: sub}, nil
}

// WatchNewReview is a free log subscription operation binding the contract event 0x664e6e86d3202d6c94bfaba88bc4520e137f0c5c3bad1b6b668db2ad1b8ca38d.
//
// Solidity: event NewReview(uint256 reviewId, uint256 bookId, string contenxt)
func (_Bookreview *BookreviewFilterer) WatchNewReview(opts *bind.WatchOpts, sink chan<- *BookreviewNewReview) (event.Subscription, error) {

	logs, sub, err := _Bookreview.contract.WatchLogs(opts, "NewReview")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookreviewNewReview)
				if err := _Bookreview.contract.UnpackLog(event, "NewReview", log); err != nil {
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

// ParseNewReview is a log parse operation binding the contract event 0x664e6e86d3202d6c94bfaba88bc4520e137f0c5c3bad1b6b668db2ad1b8ca38d.
//
// Solidity: event NewReview(uint256 reviewId, uint256 bookId, string contenxt)
func (_Bookreview *BookreviewFilterer) ParseNewReview(log types.Log) (*BookreviewNewReview, error) {
	event := new(BookreviewNewReview)
	if err := _Bookreview.contract.UnpackLog(event, "NewReview", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
