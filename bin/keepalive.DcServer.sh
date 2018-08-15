ROOT=$(cd "$(dirname "$0")"; cd ../; pwd)
LOG=${ROOT}/Log
PROC=${ROOT}/DcServer/DcServer
FLAG='
-dc_center=127.0.0.1:9999
-redis=127.0.0.1:6379
-udp_srv=127.0.0.1:11110
-log=/data/github/dc/Log
'
NRPROC=`ps ax | grep -v grep | grep -w $PROC | grep -w "$FLAG" | wc -l`
echo $NRPROC
if [ $NRPROC -lt 1 ]
then
echo $(date +%Y-%m-%d) $(date +%H:%M:%S) $PROC >> $LOG/restart.log
# echo "$PROC $FLAG >> $LOG/$(basename $PROC).stderr 2>&1 &"
$PROC $FLAG >> $LOG/$(basename $PROC).stderr 2>&1 &
fi
