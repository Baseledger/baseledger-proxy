To trigger a request synchronization request:

1. Make sure docker installed
2. Clone https://github.com/Baseledger/baseledger-lakewood.git
3. Copy the blockchain_app folder from lakewood repo to root of this repo
4. Navigate to repo root/ops/local folder
5. Run ./run_blockchain.sh (WSL) or sudo sh run_blockchain.sh (Non root user on Linux)
6. Import 'proxy_app/misc/Baseledger Proxy v2.postman_collection.json' to postman
7. Open 'create initial suggestion' request
8. Fire away


---
**Possible Problems when running on MacOS:**
1. exhausted ports:
  - when getting the <span style="color:red">*connect: cannot assign requested address* text</span>
 error while running starport
  - try to extend the ports available on the mac by running the following two commands:
    - `sudo sysctl -w net.inet.ip.portrange.first=32768`
    - `sudo sysctl -w net.inet.ip.portrange.hifirst=32768`
2. No such Container:
   - when getting a <span style="color:red">*no such Container*</span> error due to the name convention of docker-compose on macOS
   - downgrade docker-compose from v2 to v1
   - When using the Docker Desktop application it is possible to simply uncheck the "Use Docker Compose V2" in the preferences under general.
3. .baseledger/config/config.toml no such file or directory
  - when getting the  <span style="color:red">*.baseledger/config/config.toml no such file or directory*</span> error
  - replace "~" with "/root" in the  ops/local/run-blockchain.sh script
