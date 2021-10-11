#!/bin/sh
# 检测进程是否在运行，并在发现进程消失/没有运行时重启进程。
while :
do
    site=$(ps -ef|grep "/www/wwwroot/wxapp.yidaogz.cn/cli/consume 3 15" | grep -v "grep")
    if [ "$site" ]
    then
        sleep 10
    else
        /www/wwwroot/wxapp.yidaogz.cn/cli/consume 3 15 &
    fi
    sleep 10
done