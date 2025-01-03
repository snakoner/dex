.PHONY: build
.default := build

CONTRACT_NAME_FACTORY = DEXFactory
CONTRACT_NAME_LP = LiquidityProviderERC20
CONTRACT_NAME_POOL = DEXPool

build:
	docker-compose build

generate_go_bindings:
	solc --base-path . --include-path node_modules --overwrite --abi --bin -o build contracts/$(CONTRACT_NAME_FACTORY).sol
	solc --base-path . --include-path node_modules --overwrite --abi --bin -o build contracts/$(CONTRACT_NAME_LP).sol
	solc --base-path . --include-path node_modules --overwrite --abi --bin -o build contracts/$(CONTRACT_NAME_POOL).sol

	abigen --bin=build/$(CONTRACT_NAME_FACTORY).bin --abi=build/$(CONTRACT_NAME_FACTORY).abi --pkg=factory --out=build/factory.go
	abigen --bin=build/$(CONTRACT_NAME_LP).bin --abi=build/$(CONTRACT_NAME_LP).abi --pkg=liquidityprovider --out=build/lp.go
	abigen --bin=build/$(CONTRACT_NAME_POOL).bin --abi=build/$(CONTRACT_NAME_POOL).abi --pkg=pool --out=build/pool.go

	cp ./build/factory.go ./backend/internal/bindings/factory
	cp ./build/lp.go ./backend/internal/bindings/liquidity-provider
	cp ./build/pool.go ./backend/internal/bindings/pool

deploy:
	rm -rf ./ignition/deployments
	npx hardhat ignition deploy ./ignition/modules/Lottery.ts --show-stack-traces --network sepolia

up:
	docker-compose up

stop:
	docker-compose down

rmi_last_image:
	docker rmi $(shell docker images -q | head -n 1)

rm_last_container:
	docker rm -f $(shell docker ps -a -q | head -n 1)