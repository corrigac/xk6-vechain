package toolchain

//go:generate docker run -v ./:/sources ethereum/solc:0.8.19 -o /sources --abi --bin /sources/Toolchain.sol --overwrite
//go:generate go run github.com/darrenvechain/thorgo/cmd/thorgen@v1.0.0 --abi Toolchain.abi --bin Toolchain.bin --pkg toolchain --out contract.go
//go:generate rm -f Toolchain.abi Toolchain.bin
