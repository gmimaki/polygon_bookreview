# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.js
```

# メモ

-  デプロイ: npx hardhat run scripts/deploy.js --network mumbai
- GoとContactの接続: https://goethereumbook.org/smart-contract-compile/

```go
solc --abi BookReview.sol
solc --bin BookReview.sol
abigen --bin=BookReview_sol_BookReview.bin --abi=BookReview_sol_BookReview.abi --pkg=store --out=BookReview.go
```