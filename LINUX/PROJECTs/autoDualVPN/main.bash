#!/bin/bash

. ./lib.bash

load_settings(){
  settings_path=$get_settings_path
  echo "${settings_path}"
  . "${setting_path}"
}


check_root

echo 'load settings...'
load_settings



