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