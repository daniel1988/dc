#!/usr/bin/env bash

ROOT="/data/dc/"
srv_list=("DcCenter" "DcServer" )
script_list=("keepalive.DcServer.sh" "keepalive.DcCenter.sh")

function listService() {
    for srv in ${srv_list[@]}; do
        # ps -ef | grep $srv|grep -v grep

        PID=`ps -ef | grep $srv | grep "$srv" | grep -v grep | awk '{print $2}'`
        if [ "$PID" ] ; then
            echo -e "$srv\t\033[33;49;2m[ 运行中 ]\033[39;49;0m"
        else
            echo -e "$srv\t\033[37;49;2m[ 未运行 ]\033[39;49;0m"
        fi
    done;
}

function shutdownService() {
    for srv in ${srv_list[@]}; do
        PID=`ps -ef | grep $srv | grep "$srv" | grep -v grep | awk '{print $2}'`
        if [ "$PID" ] ;then
            kill -9 ${PID}
            echo -e "$(date +'%Y-%m-%d %H:%M:%S') $srv\t\033[31;49;2m[ 已停止 ] \033[39;49;0m"
        else
            echo -e "$srv\t\033[37;49;2m[ 未运行 ]\033[39;49;0m"
        fi
    done;
}

function startService() {
    for script in ${script_list[@]}; do
        PNAME="${script:10:-3}"
        PID=`ps -ef | grep ${script:10:-3} | grep "$PNAME" | grep -v grep | awk '{print $2}'`
        if [ "$PID" ] ; then
            echo -e "$PNAME\t\033[33;49;2m[ 运行中 ]\033[39;49;0m"
        else
            cd ${ROOT}bin && bash ${script}
            echo -e "$(date +'%Y-%m-%d %H:%M:%S') $PNAME\t\033[32;49;2m[ 已启动 ]\033[39;49;0m"
        fi
    done;
}


if [ "$1" = "list" ];then
    listService
elif [ "$1" = "kill" ];then
    shutdownService
elif [ "$1" = "start" ];then
    startService
elif [ "$1" = "restart" ];then
    shutdownService
    startService
else
    listService
fi
