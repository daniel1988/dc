<?php
$client = new swoole_client(SWOOLE_SOCK_TCP, SWOOLE_SOCK_ASYNC);
//设置事件回调函数
$client->on("connect", function($cli) {
    $arr = [
        "DeviceKey" => md5(time()),
        "AppName"   => "Test",
        "AppVer"    => 0,
        "AppMarket" => 1,
        "OS"        => "10.0.1",
        "OsVer"     => 1,
        "AppVerCode"    => 0,
        "AppRefId"      => time(),
        "CliendAddr"    => '127.0.0.1',
        "UserID"        => rand(10000, 9999999),
    ];
    $message = json_encode($arr);
    $res=$cli->send( $message );

    var_dump( $res ) ;
    $cli->close();
});
$client->on("receive", function($cli, $data){
    echo "Received: ".$data."\n";
});
$client->on("error", function($cli){
    echo "Connect failed\n";
});
$client->on("close", function($cli){
    echo "Connection close\n";
});
//发起网络连接
//
$client->connect('127.0.0.1', 9999, 0.5);