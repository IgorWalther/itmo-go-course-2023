CURDIR=$(shell pwd)
GENERATEDDIR=${CURDIR}/pkg/generated

generate-proto:
	rm -rf ${GENERATEDDIR}/proto
	mkdir -p ${GENERATEDDIR}/proto
	protoc --proto_path=${CURDIR}/internal/api/proto \
	--go_out=${GENERATEDDIR}/proto \
	--go-grpc_out=${GENERATEDDIR}/proto \
	${CURDIR}/internal/api/proto/note.proto
	go mod tidy
