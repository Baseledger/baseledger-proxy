

First node Windows:
SET POSTGRES_EXPOSED_PORT=5432& SET NATS_EXPOSED_PORT=4222& SET STARPORT_API_PORT=1317& SET TENDERMINT_NODE_PORT=26657
docker-compose -p first_node up


First node Linux: 
export POSTGRES_EXPOSED_PORT=5432
export NATS_EXPOSED_PORT=4222
export STARPORT_API_PORT=1317
export TENDERMINT_NODE_PORT=26657
docker-compose -p first_node up


Second node Windows:
SET POSTGRES_EXPOSED_PORT=5433& SET NATS_EXPOSED_PORT=4223& SET STARPORT_API_PORT=1318& SET TENDERMINT_NODE_PORT=26658
export POSTGRES_EXPOSED_PORT=5433& export NATS_EXPOSED_PORT=4223& export STARPORT_API_PORT=1318& export TENDERMINT_NODE_PORT=26658
docker-compose -p second_node up

Second node Linux:
export POSTGRES_EXPOSED_PORT=5433
export NATS_EXPOSED_PORT=4223
export STARPORT_API_PORT=1318
export TENDERMINT_NODE_PORT=26658
docker-compose -p first_node up

TODO https://tutorials.cosmos.network/nameservice/tutorial/20-build-run.html