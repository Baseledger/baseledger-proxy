# Docker prereqs
export COMPOSE_DOCKER_CLI_BUILD=1 && export DOCKER_BUILDKIT=1

# Sets up environment for first node and runs it
export POSTGRES_EXPOSED_PORT=5432 && export NATS_EXPOSED_PORT=4222 && export STARPORT_API_PORT=1317 && export TENDERMINT_NODE_PORT=26657
docker-compose -p first_node up -d

# Sets up environment for second node and runs it
export POSTGRES_EXPOSED_PORT=5433  && export NATS_EXPOSED_PORT=4223  && export STARPORT_API_PORT=1318  && export TENDERMINT_NODE_PORT=26658
docker-compose -p second_node up -d


# If needed, waiting mechanism
# ./await_tcp.sh -h localhost -p 1317
# ./await_tcp.sh -h localhost -p 1318

# Initialize first node
docker exec first_node_starport_1 baseledgerd init node1 --chain-id baseledger

# Initialize first node validator account
docker exec first_node_starport_1 baseledgerd keys add node1_validator --keyring-backend test

# Add first node validator account as genesis
node1_validator_address=$(docker exec first_node_starport_1 baseledgerd keys show node1_validator -a --keyring-backend test)
docker exec first_node_starport_1 baseledgerd add-genesis-account ${node1_validator_address} 100000000000stake

# Initialize second node validator account
docker exec second_node_starport_1 baseledgerd keys add node2_validator --keyring-backend test

# Add second node validator account as genesis
node2_validator_address=$(docker exec second_node_starport_1 baseledgerd keys show node2_validator -a --keyring-backend test)
docker exec first_node_starport_1 baseledgerd add-genesis-account ${node2_validator_address} 100000000000stake

# Copy genesis from first to second node via host machine for gentx generation
docker cp first_node_starport_1:/root/.baseledger/config/genesis.json .
docker cp ./genesis.json second_node_starport_1:/root/.baseledger/config/genesis.json

# Generate genensis transaction on the second node
docker exec second_node_starport_1 baseledgerd gentx node2_validator 100000000stake --chain-id baseledger --keyring-backend test

# Copy gentx from second to first node via host machine for genesis collection
docker cp second_node_starport_1:/root/.baseledger/config/gentx .
docker cp ./gentx first_node_starport_1:/root/.baseledger/config

# Generate genensis transaction on the first node
docker exec first_node_starport_1 baseledgerd gentx node1_validator 100000000stake --chain-id baseledger --keyring-backend test

# Collect both genesis transactions
docker exec first_node_starport_1 baseledgerd collect-gentxs

# Copy fully formed genesis from first to second node via host machine
docker cp first_node_starport_1:/root/.baseledger/config/genesis.json . 
docker cp ./genesis.json second_node_starport_1:/root/.baseledger/config/genesis.json

# Setup config toml
node1_id=$(docker exec first_node_starport_1 baseledgerd tendermint show-node-id)
node2_id=$(docker exec second_node_starport_1 baseledgerd tendermint show-node-id)

internal_host_ip=$(docker exec first_node_starport_1 getent hosts host.docker.internal | awk '{print $1}')

# this adds peers to each other
docker exec first_node_starport_1 sed -i 's/persistent_peers = ".*/persistent_peers = "'${node2_id}'@'${internal_host_ip}':26658"/' ~/.baseledger/config/config.toml
docker exec second_node_starport_1 sed -i 's/persistent_peers = ""/persistent_peers = "'${node1_id}'@'${internal_host_ip}':26657"/' ~/.baseledger/config/config.toml

# this enables grpc
docker exec first_node_starport_1 sed -i 's@laddr = "tcp://127.0.0.1:26657"@laddr = "tcp://0.0.0.0:26657"@' ~/.baseledger/config/config.toml
docker exec second_node_starport_1 sed -i 's@laddr = "tcp://127.0.0.1:26657"@laddr = "tcp://0.0.0.0:26657"@' ~/.baseledger/config/config.toml

# this enables rest api, it is only enable = false entry, maybe we can make it a bit more precise?
docker exec first_node_starport_1 sed -i 's/enable = false/enable = true/' ~/.baseledger/config/app.toml
docker exec second_node_starport_1 sed -i 's/enable = false/enable = true/' ~/.baseledger/config/app.toml

# this allows connecting peers not in the address book
docker exec first_node_starport_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml
docker exec second_node_starport_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml

# This allows connections from localhost to tendermint API
docker exec first_node_starport_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml
docker exec second_node_starport_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml

# This increases the timeout between blocks to 30s
docker exec first_node_starport_1 sed -i 's/timeout_commit = "5s"/timeout_commit = "30s"/' ~/.baseledger/config/config.toml
docker exec second_node_starport_1 sed -i 's/timeout_commit = "5s"/timeout_commit = "30s"/' ~/.baseledger/config/config.toml


# start first node - TODO: Has to  be executed in a separate window after running this script in order to have logs
# docker exec first_node_starport_1 baseledgerd start

# start second node - TODO: Has to  be executed in a separate window after running this script in order to have logs
# docker exec second_node_starport_1 baseledgerd start

# cleanup

rm ./genesis.json
rm -rf ./gentx

