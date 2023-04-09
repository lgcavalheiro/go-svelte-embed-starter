ENTRY=main.go
FLAGS=-ldflags="-s -w"
OUTDIR=output
PROJECT=gses
FRONTMANAGER=cd frontend && npm run
COVER=coverage

# https://stackoverflow.com/questions/28459102/golang-compile-environment-variable-into-binary

present:
	@echo "   ___________ ___________"
	@echo "  / ____/ ___// ____/ ___/"
	@echo " / / __ \__ \/ __/  \__ \ "
	@echo "/ /_/ /___/ / /___ ___/ /"
	@echo "\____//____/_____//____/"

run: present front-build build
	go run ${ENTRY}

build: present front-build
	go build -o ${OUTDIR}/${PROJECT} ${FLAGS} ${ENTRY}

test:
	go test -failfast -vet=off -race -timeout=1m ./...

cover:
	$(shell [ ! -e ${COVER} ] && mkdir ${COVER})
	go test ./... -race -covermode=atomic -coverprofile=${COVER}/coverage.out
	go tool cover -html=${COVER}/${COVER}.out -o ${COVER}/${COVER}.html

front-run: present
	${FRONTMANAGER} dev

front-build: present
	${FRONTMANAGER} check 
	${FRONTMANAGER} build

front-preview: present
	${FRONTMANAGER} preview
