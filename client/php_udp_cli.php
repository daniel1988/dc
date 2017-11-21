<?php
$count = 0;
$timer = swoole_timer_tick(10, function() {

    global $count;
    global $timer;
    $count += 1;
    if ( $count >= 10 ) {
        swoole_timer_clear( $timer );
    }

    echo $count ;

    $sock = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);
    $arr = [
        "DeviceKey" => md5(time().$count),
        "AppName"   => "Test",
        "AppVer"    => 0,
        "AppMarket" => 1,
        "OS"        => "10.0.1",
        "OsVer"     => "1",
        "AppVerCode"    => 'xxy',
        "AppRefId"      => time(),
        "CliendAddr"    => '127.0.0.1',
        "UserID"        => rand(10000, 9999999),
    ];
    $message = json_encode($arr);
    $res = socket_sendto($sock, $message, strlen($message), 0, "127.0.0.1", "11110");

    echo sprintf("%s socket_sendto count:%d\n", date('YmdHis'), $count);
    socket_close($sock);
});

