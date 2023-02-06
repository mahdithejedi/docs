function _RunByNetwork(){
  echo "Enter Network Name Options:$(docker network ls --format "{{.Name}}" | xargs )"
  echo
  read -r _Name
#  docker run "$(docker ps --filter network="$_Name" -q -a | xargs)"
  for container_id in `docker ps --filter network="$_Name" -q -a`; do docker start "$container_id";done
  docker ps --filter network="$_Name" --format "table {{.ID}} | {{.Status}}" -a
}

function DockerRun(){
  while true; do
    echo "Run by Network[N] or Quit[Q]"
    read -r option
    case $option in
      [Nn]* ) _RunByNetwork;break;;
      [Qq]* ) break ;;
      * ) echo "Network[N] or Quit[Q]";;
    esac
  done
}