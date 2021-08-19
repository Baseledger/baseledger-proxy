# Sets up environment for a node
export POSTGRES_EXPOSED_PORT=5432 && 
export NATS_EXPOSED_PORT=4222 && 
export BLOCKCHAIN_APP_API_PORT=1317 && 
export TENDERMINT_NODE_GRPC_PORT=26657 && 
export TENDERMINT_NODE_PORT=26655 && 
export PROXY_APP_PORT=8081 &&
export ORGANIZATION_ID=d45c9b93-3eef-4993-add6-aa1c84d17eea # unique identifier of the organization, generate a uuid and later add it as org through postman

docker-compose -p baseledger-node up -d