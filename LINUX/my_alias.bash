#__DIR__="$(realpath ./)"
__DIR__=$(realpath "$0")
__DIR__=${__DIR__%/*}

#   *********
#     GIT
#   *********
source "$__DIR__"/BashAlias/git.bash
echo "Git commands imported"


#   *********
#     DOCKER
# TODO: Integrate with golang
# https://github.com/spf13/cobra
#   *********
source "$__DIR__"/BashAlias/Docker.bash
echo "Docker commands imported"



# Arsenal
alias a='arsenal'

# ***** HOW TO RUN
#if [ -f ~/.my_alias.sh ]; then
#        source ~/.my_alias.sh
#else
#        print "Linked my_alias not found"
#fi


# Unset an Env file
function UnsetEnv() {
  unset $(more "$1" | awk -F= '{print $1}' | xargs)
}

function SetEnv(){
  export $(xargs < "$1" )
}

function ResetEnv(){
  UnsetEnv "$@"
  SetEnv "$@"
}

function _openPorts() {
    egrep -e '[[:digit:]]*\/' /etc/services | awk '{print $2}' | cut -d "/" -f 1
}

function ports(){
  _openPorts | xargs
}
#
#function portServiceName(){
#  port=_openPorts | grep $1
#  if [[ port ]]
#  then
#    grep ''
#}
#}