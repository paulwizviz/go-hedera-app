#!/bin/bash

if [ "$(basename $(realpath .))" != "go-hedera-app" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export SOLC_TOOL="ethereum/solc:0.8.28"
export DOCKER_DEFAULT_PLATFORM=linux/amd64

function compile_sol(){
    local sol="Hello.sol" # Change this to meet your script
    docker run -v $(PWD)/solidity/hello/$sol:/opt/solidity/$sol \
            -v $(PWD)/solidity/bin/hello/:/opt/bin \
            ${SOLC_TOOL} --bin /opt/solidity/$sol -o /opt/bin
}

COMMAND=$1

case $COMMAND in
    "clean")
        rm -rf $PWD/solidity/abi/hello
        rm -rf $PWD/internal/hello
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "compile")
        compile_sol
        ;;
    *)
        echo "Usage: $0 [abi | compile]"
        ;;
esac