#rm -rf channel-artifacts
#rm -rf crypto-config
docker kill $(docker ps -q)
docker rm $(docker ps -a -q)
docker volume prune

