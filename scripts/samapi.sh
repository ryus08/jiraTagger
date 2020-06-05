#!/usr/bin/env bash
if [[ -z "${1}" ]]; then
    echo "Running local API!"
    if [[ -z "${2}" ]]; then
        sam.cmd local start-api --host 127.0.0.1
    else
        sam.cmd local start-api --docker-network ${2}
    fi

else
    echo "$(tput setaf 1) * Now Call API route from browser or POSTMAN, then open VisualStudio Code, put a breakpoint and start debugging!$(tput sgr0)"

    if [[ -z "${2}" ]]; then
         echo "No Network Passed, Applying default network!"
         sam.cmd local start-api -d 8997 --debugger-path ./scripts/linux --region us-east-1 --debug-args "-delveAPI=2"
    else
        echo "Running inside network ${2}"
        sam.cmd local start-api -d 8997 --debugger-path ./scripts/linux --region us-east-1 --docker-network ${2} --debug-args "-delveAPI=2"
    fi
fi