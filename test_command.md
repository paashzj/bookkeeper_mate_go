### start
```bash
docker run ttbb/bookkeeper:stand-alone
```
### start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 -d ttbb/bookkeeper:stand-alone
```
### localhost
```bash
docker run -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" ttbb/bookkeeper:stand-alone
```
### localhost, start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" -e BOOKIE_ROOT_LOGGER="DEBUG,INFO,CONSOLE" ttbb/bookkeeper:stand-alone
```
### cluster
```bash
docker run -e ZK_ADDR=172.16.0.56:2181,172.16.0.55:2181,172.16.0.54:2181 -e HOSTNAME=test-bookkeeper-0 -d ttbb/bookkeeper:cluster
```