LOG=/data/dc/Log
PROC=/data/dc/DcCenter/DcCenter
FLAG='
-dc_center=127.0.0.1:9999
-log=/data/dc/Log
'
NRPROC=`ps ax | grep -v grep | grep -w $PROC | grep -w "$FLAG" | wc -l`
if [ $NRPROC -lt 1 ]
then
echo $(date +%Y-%m-%d) $(date +%H:%M:%S) $PROC >> $LOG/restart.log
$PROC $FLAG >> $LOG/$(basename $PROC).stderr 2>&1 &
fi
