
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
1. SET POSTGRES_EXPOSED_PORT=5432& SET NATS_EXPOSED_PORT=4222& SET STARPORT_API_PORT=1317& SET TENDERMINT_NODE_PORT=26657
2. Navigate to repo root/ops folder
3. docker-compose -p first_node up
4. starport serve --verbose

Linux: 
1. export POSTGRES_EXPOSED_PORT=5432 && export NATS_EXPOSED_PORT=4222 && export STARPORT_API_PORT=1317 && export TENDERMINT_NODE_PORT=26657
2. Navigate to repo root/ops folder
3. export COMPOSE_DOCKER_CLI_BUILD=1 && export DOCKER_BUILDKIT=1
4. sudo -E docker-compose -p first_node up
5. starport serve --verbose
