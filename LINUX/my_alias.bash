#   *********
#     GIT
#   *********

alias p='git pull'
alias s='git status'
alias f='git fetch'
alias fa='git fetch --all'
function push() {
    while true; do
      echo "Do you want to push to $(git branch --show-current)??"
        read -r yn
        case $yn in
            [Yy]* ) git push; break;;
            [Nn]* ) break ;;
            * ) echo "Please answer yes or no.";;
        esac
    done
}

alias m='git merge'
alias ma='git merge --abort'

alias ba='git branch --all'

alias c='git checkout'

alias sts='git stash save'
alias st='git stash'

function pycache() {
    find . | grep -E "(/__pycache__$|\.pyc$|\.pyo$)" | xargs rm -rf
}



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