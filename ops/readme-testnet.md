# Running the test net blockchain in a node-per-server setup

## Setup a node infrastructure on a server

1. Copy docker-compose.yml to the server
2. Copy *setup-node-infrastructure-for-testnet.sh* to the same folder on the server 
3. Run *setup-node-infrastructure-for-testnet.sh*
      This one sets up all components necessary for a node to function.
      Run it on as many servers as nodes are needed.
4. Open all ports in the firewall that are necessary for external communication


## Setup a blokchain with two validator nodes

1. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd init node1 --chain-id baseledger

2. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd keys add node1_validator --keyring-backend test

### Add first node validator account as genesis
3. NODE 1: node1_validator_address=$(docker exec baseledger_node_blockchain_app_1 baseledgerd keys show node1_validator -a --keyring-backend test)
           docker exec baseledger_node_blockchain_app_1 baseledgerd add-genesis-account ${node1_validator_address} 100000000000stake

### Initialize second node validator account
4. NODE 2: docker exec baseledger_node_blockchain_app_1 baseledgerd keys add node2_validator --keyring-backend test

### Add second node validator account as genesis
5. NODE 2: docker exec baseledger_node_blockchain_app_1 baseledgerd keys show node2_validator -a --keyring-backend test
6. copy address to clipboard
7. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd add-genesis-account <copied node 2 address> 100000000000stake

### Copy genesis from first to second node via host machine for gentx generation
8. NODE 1: docker cp baseledger_node_blockchain_app_1:/root/.baseledger/config/genesis.json .
9. NODE 1: Copy genesis.json to clipboard
10. NODE 2: copy cliboard to genesis.json
11. NODE 2: docker cp ./genesis.json baseledger_node_blockchain_app_1:/root/.baseledger/config/genesis.json

### Generate genensis transaction on the second node
12. NODE 2: docker exec baseledger_node_blockchain_app_1 baseledgerd gentx node2_validator 100000000stake --chain-id baseledger --keyring-backend test

### Copy gentx from second to first node via host machine for genesis collection
13. NODE 2: docker cp baseledger_node_blockchain_app_1:/root/.baseledger/config/gentx .
14. NODE 2: copy gentx to clipboard
15. NODE 1: copy clipboard to gentx
16. NODE 1: docker cp ./gentx baseledger_node_blockchain_app_1:/root/.baseledger/config

### Generate genensis transaction on the first node
17. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd gentx node1_validator 100000000stake --chain-id baseledger --keyring-backend test

### Collect both genesis transactions
18. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd collect-gentxs

### Copy fully formed genesis from first to second node via host machine
19. repeat 8, 9, 10 and 11

### Setup config toml
20. NODE 1: node1_id=$(docker exec baseledger_node_blockchain_app_1 baseledgerd tendermint show-node-id)

### this adds peers to each other
22. NODE 1: docker exec baseledger_node_blockchain_app_1 sed -i 's/persistent_peers = ".*/persistent_peers = "'${node2_id}'@'<static_ip_of_node_2>':'26655'"/' ~/.baseledger/config/config.toml

### this enables rest api, it is only enable = false entry, maybe we can make it a bit more precise?
23. NODE 1: docker exec baseledger_node_blockchain_app_1 sed -i 's/enable = false/enable = true/' ~/.baseledger/config/app.toml

### this enables grpc
23. NODE 1:docker exec baseledger_node_blockchain_app_1 sed -i 's@laddr = "tcp://127.0.0.1:'26657'"@laddr = "tcp://0.0.0.0:'26657'"@' ~/.baseledger/config/config.toml

### this allows connecting peers not in the address book
24. NODE 1: docker exec baseledger_node_blockchain_app_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml

### This allows connections from localhost to tendermint API
25. NODE 1: docker exec baseledger_node_blockchain_app_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml

### This increases the timeout between blocks to 30s
26. NODE 1: docker exec baseledger_node_blockchain_app_1 sed -i 's/timeout_commit = "5s"/timeout_commit = "30s"/' ~/.baseledger/config/config.toml

### Node 2 setup
27. NODE 2: repeat 20-26 
node2_id=$(docker exec baseledger_node_blockchain_app_1 baseledgerd tendermint show-node-id)
docker exec baseledger_node_blockchain_app_1 sed -i 's/persistent_peers = ""/persistent_peers = "'${node1_id}'@'<static_ip_of_node1>':'26655'"/' ~/.baseledger/config/config.toml


### Run both nodes
28. NODE 1: docker exec baseledger_node_blockchain_app_1 baseledgerd start
28. NODE 2: docker exec baseledger_node_blockchain_app_1 baseledgerd start


## Setup a third 'general' node
1. NODE 3: docker exec baseledger_node_blockchain_app_1 baseledgerd init node3 --chain-id baseledger

2. NODE 3: docker exec baseledger_node_blockchain_app_1 baseledgerd keys add node3_validator --keyring-backend test

3. NODE 3: docker exec baseledger_node_blockchain_app_1 sed -i 's/persistent_peers = ""/persistent_peers = "'${node1_id}'@'<static_ip_of_node1>':'26655'"/' ~/.baseledger/config/config.toml

4. NODE 3: docker exec baseledger_node_blockchain_app_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml

5. NODE 3: docker exec baseledger_node_blockchain_app_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml

6. step 19 but from node 1 to node 3

7. NODE 3: docker exec baseledger_node_blockchain_app_1 baseledgerd start

## Add third node as validator TODO - format in the same manner as above

1. Send a minimal amount of STAKE tokens from Node1 to the Node_to_become_validator:
sudo docker exec first_node_blockchain_app_1 baseledgerd tx bank send node1_validator baseledger1xax2e85vqn4n26wxk0qfcy9jgmwlgvnw750hzm 1stake --yes
2. Node_To_become_Validator now take the minimal amount of STAKE tokens received and stakes them to make him a validator:
sudo docker exec third_node_blockchain_app_1 baseledgerd tx staking create-validator  --amount=100stake  --pubkey=baseledgervalconspub1zcjduepqarfuxqq6nrp9fdt6mrdhp8kxdpmvhurk2suds3qd37xyusrywmdqav5vfq --moniker="node3"  --commission-rate="0.10" --commission-max-rate="0.20" --commission-max-change-rate="0.01" --min-self-delegation="1" --from=node3_validator --yes 
In the command above i removed (-gas="200000" --gas-prices="0.025stake" ) as we assume to have 0 gas costs that way
--from = <name of the node to become validator>
--pubkey <output of tendermint show-validator on node_to_become_validator>
--moniker= <unique name for the validator>
3. Now the new validator should be in the validator set in status UNBONDED (he has to feww tokens staked to participate). We stake the right amount from Node1 (our token controlling node):
sudo docker exec first_node_blockchain_app_1 baseledgerd tx staking delegate baseledgervaloper1xax2e85vqn4n26wxk0qfcy9jgmwlgvnwpq7z6g 100000000stake --from=node1_validator --yes 
--baseledgervaloper-address from the new validator node, can be seen in "sudo docker exec first_node_blockchain_app_1 baseledgerd query staking validators"
--from=<our token controlling node1>