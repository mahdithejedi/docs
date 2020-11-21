#!/bin/bash

get_app_name(){
  echo 'dualVPN'
}

get_root_privilege(){
  echo ERROR: You have to run this command as root
  exit 1
}

check_root(){
  if [ "$EUID" -ne 0 ]
  then get_root_privilege
fi
}

get_interface_ip_address() {
  interface_name=$1
	ifconfig "$interface_name" | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b" | head -1
}

add_ip_route_rule(){
  ip route add "$1" via "$2" dev "$3"
}

get_setting_path(){
  _home_path=$(xdg-user-dir)
  _app_name=$(get_app_name)
  echo "($_home_path)/.($_app_name)"
}