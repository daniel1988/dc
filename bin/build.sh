#!/usr/bin/env bash
ROOT="/data/dc/"
go=$(which "go")
build_list=("DcCenter" "DcServer")
for service in ${build_list[@]};do
    echo "cd ${ROOT}${service} && ${go} build"
    cd ${ROOT}${service} && ${go} build
done;