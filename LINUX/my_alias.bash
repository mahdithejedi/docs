#   *********
#     GIT
#   *********

alias p='git pull'
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

alias ma='git merge --abort'

alias ba='git branch --all'

alias c='git checkout'

function pycache() {
    find . | grep -E "(/__pycache__$|\.pyc$|\.pyo$)" | xargs rm -rf
}

# ***** HOW TO RUN
#if [ -f ~/.my_alias.sh ]; then
#        source ~/.my_alias.sh
#else
#        print "Linked my_alias not found"
#fi