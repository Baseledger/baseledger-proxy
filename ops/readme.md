SET POSTGRES_EXPOSED_PORT=5432 - Postgres
SET NATS_EXPOSED_PORT=4222 - Nats
SET STARPORT_API_PORT=1317 - Starport backend API exposed port
SET TENDERMINT_NODE_PORT=26657 - Tendermint node exposed port
docker-compose.exe -p first_node up

SET POSTGRES_EXPOSED_PORT=5433 - Postgres
SET NATS_EXPOSED_PORT=4223 - Nats
SET STARPORT_API_PORT=1318 - Starport backend API exposed port
SET TENDERMINT_NODE_PORT=26658 - Tendermint node exposed port
docker-compose.exe -p second_node up

TODO https://tutorials.cosmos.network/nameservice/tutorial/20-build-run.html