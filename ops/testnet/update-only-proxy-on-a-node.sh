#local

#navigate to repo root/proxy_app
docker build -f ops/Dockerfile -t baseledger/proxy_app .
docker push baseledger/proxy_app

# node 1	
docker stop <proxy_container>
docker rm <proxy_container>
docker rmi <proxy_image>
docker run --name=baseledger-node_proxy_app_1 --restart=always --network=baseledger-node_baseledger -p 8081:8080 --add-host=host.docker.internal:host-gateway -e DB_HOST=postgres-local-node -e ORGANIZATION_ID=d9c102eb-b173-45ac-b640-24d28a3c9f0c -e DB_UB_USER=<user> -e DB_UB_PWD=<pass> -e API_UB_USER=<user> -e API_UB_PWD=<pass> -e ETHEREUM_PRIVATE_KEY=<pvtkey> -e ETHEREUM_API_URL=<infura> -e SWAGGER_HOST=137.184.72.13:8081 -e JWT_SECRET=<secret> -d baseledger/proxy_app

# node 2
docker stop <proxy_container>
docker rm <proxy_container>
docker rmi <proxy_image>
docker run --name=baseledger-node_proxy_app_1 --restart=always --network=baseledger-node_baseledger -p 8081:8080 --add-host=host.docker.internal:host-gateway -e DB_HOST=postgres-local-node -e ORGANIZATION_ID=0bd0319b-faea-468f-b571-7c6de389f050 -e DB_UB_USER=<user> -e DB_UB_PWD=<pass> -e API_UB_USER=<user> -e API_UB_PWD=<pass> -e ETHEREUM_PRIVATE_KEY=<pvtkey> -e ETHEREUM_API_URL=<infura> -e SWAGGER_HOST=137.184.25.137:8081 -e JWT_SECRET=<secret> -d baseledger/proxy_app