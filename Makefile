.PHONY : all
all : proto-gen-go

proto-gen-go:
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f block.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f blockState.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f candles.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f customEvent.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f dexeEvents.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f directSwap.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f holder.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f holdersBlock.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f limitOrder.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f limitTrade.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f liquidityEvant.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f newPair.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f pair.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f pairSwap.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f priceAlert.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f transferEvent.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f trade.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f tradesBlock.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f priceUpdates.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f cexPriceUpdate.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f apiTrades.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f transaction.proto -l go -o .
	@docker run --rm -v `pwd`:/defs namely/protoc-all:1.51_1 -i proto -f smartSwap.proto -l go -o .
