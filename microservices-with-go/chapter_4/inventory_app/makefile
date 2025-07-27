.PHONY: benchmark test gen-proto

test:
	go test -v ./...

benchmark:
	cd cmd/sizecompare && go test -bench=. -run=^$ -benchtime=20s -count=10 && cd ../..

gen-proto:
	protoc -I=api --go_out=. ./api/*.proto