### start
```bash
docker run --rm ttbb/bookkeeper:mate
```
### start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 --rm -d ttbb/bookkeeper:mate
```
### localhost
```bash
docker run --rm -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" ttbb/bookkeeper:mate
```
### localhost, start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 --rm -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" -e BOOKIE_ROOT_LOGGER="DEBUG,INFO,CONSOLE" ttbb/bookkeeper:mate
```
### cluster tls
```bash
docker run --rm -d -e ZOOKEEPER_TLS_ENABLE=true -e BOOKKEEPER_TLS_ENABLE=true ttbb/bookkeeper:mate
```
### cluster
```bash
docker run --rm -e ZK_ADDR=172.16.0.56:2181,172.16.0.55:2181,172.16.0.54:2181 -e HOSTNAME=test-bookkeeper-0 -d ttbb/bookkeeper:cluster
```