```shell
docker pull elasticsearch:6.5.4

docker network create elastic-net

docker run -d --name elastic --net elastic-net -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:6.5.4
```
