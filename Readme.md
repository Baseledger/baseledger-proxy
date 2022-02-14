To trigger a request synchronization request:

1. Make sure docker installed
2. Clone https://github.com/Baseledger/baseledger-lakewood.git
3. Copy the blockchain_app folder from lakewood repo to root of this repo
4. Navigate to repo root/ops folder
5. Run ./run_blockchain.sh (WSL) or sudo sh run_blockchain.sh (Non root user on Linux)
6. Import 'proxy_app/misc/Baseledger Proxy v2.postman_collection.json' to postman
7. Open 'create initial suggestion' request
8. Fire away