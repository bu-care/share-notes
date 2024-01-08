
# kill and rm agent docker instances
docker ps|grep xbu_api |awk '{print $1}'|xargs docker kill  > /dev/null 2>&1
docker ps -a|grep xbu_api |awk '{print $1}'|xargs docker rm > /dev/null 2>&1

docker-compose down &

