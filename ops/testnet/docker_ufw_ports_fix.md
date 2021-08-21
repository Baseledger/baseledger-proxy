1. sudo wget -O /usr/local/bin/ufw-docker \
  https://github.com/chaifeng/ufw-docker/raw/master/ufw-docker


2. sudo chmod +x /usr/local/bin/ufw-docker

3. ufw-docker install

4. sudo systemctl restart ufw

8. cleanup old UFW rules

5. reboot the machine if ports accesible

6. if reboot, docker start both proxyapp and blockchainapp containers

7. docker exec baseledger-node_blockchain_app_1 baseledgerd start

8. ufw route allow proto tcp from any to any port 4222

9. ufw route allow proto tcp from any to any port 26656