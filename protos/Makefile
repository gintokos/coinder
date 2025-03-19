.PHONY: coinupdate

coinupdate:
	@echo "Generating coinupdate protos..."
	@mkdir -p coinupdateprotos
	C:\Users\yasos\proto\bin\protoc.exe --go_out=./coinupdateprotos --go_opt=module=github.com/gintokos/coinder/coinupdateprotos --go-grpc_out=./coinupdateprotos --go-grpc_opt=module=github.com/gintokos/coinder/coinupdateprotos proto/coinupdate/coinupdate.proto
	@echo "Coinupdate protos generated successfully"