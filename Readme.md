To trigger a request synchronization request:

1. Make sure docker installed
2. Navigate to repo root/ops folder
3. Run ./run_blockchain.sh (WSL )or sudo sh run_blockchain.sh (Non root user on Linux)
4. get node1 address from logs
5. get workgroup id and second organization id from proxy_app/ops/000001_initial.up.sql
6. import proxy_app/misc/Baseledger Proxy v2.postman_collection.json to postman - 
7. create sync request
8. copy node1 address to *from*, workgroup id to *workgroup_id* and second organization id to *recipient*