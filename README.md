# What is this
This is learning repository of running consul inside container environment locally

The topology of the cluster can be found below : 

Basically there is a cluster with 1 Consul server and 3 Consul client + 3 web app

![alt text](https://github.com/leemov/consul-docker-intro/blob/main/blob/topology.png?raw=true)

## Pre-requisite 

Docker & Docker Compose version 3.8

## Docker Installation

You can install docker by following this manual 

```bash
https://docs.docker.com/get-docker/
```

## Usage
- Clone the repository. 
- On current main directory, open terminal
- Run :
```
> docker compose build
```
```
> docker compose up
```
to shut down the cluster. run this command 
```
> docker compose down
```

## Contributing
Pull requests are welcome.

## License
[MIT](https://choosealicense.com/licenses/mit/)

