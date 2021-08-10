# Docker prereqs
export COMPOSE_DOCKER_CLI_BUILD=1 && export DOCKER_BUILDKIT=1

first_node_tendermint_p2p_port=26655
tendermint_internal_grpc_port=26657
node1_id=$(docker exec first_node_starport_1 baseledgerd tendermint show-node-id)
internal_host_ip=$(docker exec first_node_starport_1 getent hosts host.docker.internal | awk '{print $1}')

docker exec third_node_starport_1 baseledgerd init node3 --chain-id baseledger
docker exec third_node_starport_1 baseledgerd keys add node3_validator --keyring-backend test
docker exec third_node_starport_1 sed -i 's/persistent_peers = ""/persistent_peers = "'${node1_id}'@'${internal_host_ip}':'${first_node_tendermint_p2p_port}'"/' ~/.baseledger/config/config.toml
docker exec third_node_starport_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml
docker exec third_node_starport_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml
docker exec third_node_starport_1 sed -i 's@laddr = "tcp://127.0.0.1:'${tendermint_internal_grpc_port}'"@laddr = "tcp://0.0.0.0:'${tendermint_internal_grpc_port}'"@' ~/.baseledger/config/config.toml
docker exec third_node_starport_1 sed -i 's/enable = false/enable = true/' ~/.baseledger/config/app.toml

docker cp ./genesis.json third_node_starport_1:/root/.baseledger/config/genesis.json

# docker exec third_node_starport_1 baseledgerd start

