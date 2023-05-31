After Installing redis -> use this command in terminal "redis-server"

Docker volume create :

> docker volume create pgaesthetic
> Docker create network :
> docker network create database-net
> Docker psql run :
> docker exec -it 37b797098943 psql -U postgres
> Check Database by command :
> \l

config sementara postgresql
POSTGRES_USER=postgres
POSTGRES_PASSWORD=g3G79n5pAWqPDhojyntH
POSTGRES_DB=railway
POSTGRES_HOST=containers-us-west-79.railway.app
POSTGRES_PORT=7730

config local postgresql
POSTGRES_USER=postgres
POSTGRES_PASSWORD=123456
POSTGRES_DB=aesthetic
POSTGRES_HOST=localhost
POSTGRES_PORT=5433

config docker postgresql
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=aesthetic
POSTGRES_HOST=host.docker.internal
POSTGRES_PORT=5434
