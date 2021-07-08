export POSTGRES_EXPOSED_PORT=5432 && export NATS_EXPOSED_PORT=4222 && export STARPORT_API_PORT=1317 && export TENDERMINT_NODE_PORT=26657
docker-compose -p first_node down -v
docker-compose -p second_node down -v