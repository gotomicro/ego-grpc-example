etcdctl put "/ego/config/ego-grpc-server" '[server.grpc]
    port = 9002
[server.governor]
    port = 9003
[trace.jaeger]
[logger.default]
  level="info"
[etcd]
    addrs=["127.0.0.1:2379"]
[registry]
    connectTimeout = "1s"
    secure = false
'

etcdctl put "/ego/config/ego-grpc-server" '[server.grpc]
    port = 9002
[server.governor]
    port = 9003
[trace.jaeger]
[logger.default]
  level="debug"
[etcd]
    addrs=["127.0.0.1:2379"]
[registry]
    connectTimeout = "1s"
    secure = false
'