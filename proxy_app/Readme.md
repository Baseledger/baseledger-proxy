To test

1. ops/run_blockchain
2. get node1 address from logs
3. get workgroup id and second organization id from migration script
4. open postman - create sync request
5. copy node1 address to *from*, workgroup id to *workgroup_id* and second organization id to *recipient*

To generate new swagger changes (make sure you have https://github.com/swaggo/swag installed)

1. cd httpd
2. swag init --parseDependency --parseInternal --parseDepth 1