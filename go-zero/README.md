# Docker
```
docker run --name etcd -d -p 2379:2379 -p 2388:2380 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:latest etcd 
```