# DC

数据收集，可用于日志上报与相应的统计数据上报。


## DcClient

业务上报客户端，暂时只提供golang客户端。以json格式上报

## DcServer

收集业务机器上报的数据，每台业务机器必须运行一个DcServer,因为UDP上报数据时会有丢包，本机则没问题。UDPServer收到数据
Rpush到Redis里面，由DcSync服务同步到DcCenter

## DcCenter

DcCenter接收各DcSync服务同步过来的数据，存储到ElasticSearch里面

>为啥不直接由DcSync存储到ElasticSearch？DcCenter接收数据时可以过滤做相应的处理，然后再存储，放在DcSync处理过滤逻辑，
很明显会增加业务机器的开销。


## Build服务

```
bash /data/dc/bin/build.sh
```

## 运行服务

本地运行时，可以直接用`dc.manager.sh start`

```
cd /data/dc/bin && bash dc.manager.sh start
```


## 关闭服务

```
bash /data/dc/bin/dc.manager.sh kill
```

## 重启服务

```
bash /data/dc/bin/dc.manager.sh restart
```



## ElasticSearch

[官方](https://www.elastic.co/downloads/elasticsearch)

[Ubuntu Install](https://www.cnblogs.com/pigzhu/p/4705870.html)