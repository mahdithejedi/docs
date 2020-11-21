#/bin/bash

get_interface_ip_address {
	return ifconfig "$1" | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b" | head -1
}
