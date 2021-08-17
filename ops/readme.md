
Linux requirements Docker: BuildKit (https://docs.docker.com/develop/develop-images/build_enhancements/#to-enable-buildkit-builds)
COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_BUILDKIT=1

# Running the full blockchain

1. Navigate to repo root/ops folder
2. Run ./run_blockchain.sh (WSL )or sudo sh run_blockchain.sh (Non root user on Linux)

# Cleanup

1. Navigate to repo root/ops folder
2. Run ./clear_blockchain.sh (WSL )or sudo sh clear_blockchain.sh (Non root user on Linux)

# Running a single node locally for development purposes
Windows:
1. SET POSTGRES_EXPOSED_PORT=5432& SET NATS_EXPOSED_PORT=4222& SET BLOCKCHAIN_APP_API_PORT=1317& SET TENDERMINT_NODE_GRPC_PORT=26657& SET TENDERMINT_NODE_PORT=26656& SET PROXY_API_PORT=8081
2. Navigate to repo root/ops folder
3. docker-compose -p first_node up
4. starport serve --verbose

Linux: 
1. export POSTGRES_EXPOSED_PORT=5432 && export NATS_EXPOSED_PORT=4222 && export BLOCKCHAIN_APP_API_PORT=1317 && export TENDERMINT_NODE_GRPC_PORT=26657 && export TENDERMINT_NODE_PORT=26656 && export PROXY_API_PORT=8082
2. Navigate to repo root/ops folder
3. export COMPOSE_DOCKER_CLI_BUILD=1 && export DOCKER_BUILDKIT=1
4. sudo -E docker-compose -p first_node up
5. starport serve --verbose


# Setup to get to the state where a third node can be added as a validator
Linux: 
1. make sure blockchain_app and proxy_app images are up to date (delete if neccessary before running for the first time)
2. navigate to ops folder and *./run_blockchain.sh* (has to be wsl if on windows)
    this:
    * builds the latest version of the images
    * creates two validator nodes and relevant infrastucture for them
    * creates relevant infrastructure for a third node to be initialized later
3. open separate cmd and run: *docker exec first_node_blockchain_app_1 baseledgerd start*
    this:
    * runs the first validator node
4. open separate cmd and run: *docker exec second_node_blockchain_app_1 baseledgerd start*
    this:
    * runs the second validator node
5. navigate to ops folder and *./add-node-to-running-blokchain.sh* (has to be wsl if on windows)
    this:
    * creates a regular node on the infrastructure created in step 2 by using the genesis file from that step

You can now a open separate cmd and run: *docker exec third_node_blockchain_app_1 baseledgerd start* if you want to start the third node

To reset the state, just run *./clear-blokchain.sh* and repeat the process. 
