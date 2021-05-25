# info

# help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Compile the binary. Copy binary product to current directory
build:
	go build -o ./toad_ocr_rpc_client

## clean: Clean output
clean:
	rm -f toad_ocr_rpc_client

## generate: generate idl code
generate:
	@sh toad_ocr_engine_idl_generate.sh
	@sh toad_ocr_preprocessor_idl_generate.sh
