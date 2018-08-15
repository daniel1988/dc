#!/usr/bin/env bash
ROOT=$(cd "$(dirname "$0")"; cd ../; pwd)
go=$(which "go")
build_list=("DcCenter" "DcServer")
for service in ${build_list[@]};do
    echo "cd ${ROOT}/${service} && ${go} build"
    cd ${ROOT}/${service} && ${go} build
done;