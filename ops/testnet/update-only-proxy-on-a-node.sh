node 1	
docker run --name=baseledger-node_proxy_app_1 --restart=always --network=baseledger-node_baseledger -p 8081:8080 --add-host=host.docker.internal:host-gateway -e DB_HOST=postgres-local-node -e ORGANIZATION_ID=d9c102eb-b173-45ac-b640-24d28a3c9f0c -e API_CONCIRCLE_URL=s4h.rp.concircle.com -d baseledger/proxy_app

node 2
docker run --name=baseledger-node_proxy_app_1 --restart=always --network=baseledger-node_baseledger -p 8081:8080 --add-host=host.docker.internal:host-gateway -e DB_HOST=postgres-local-node -e ORGANIZATION_ID=0bd0319b-faea-468f-b571-7c6de389f050 -e API_CONCIRCLE_URL=s4p.rp.concircle.com -d baseledger/proxy_app