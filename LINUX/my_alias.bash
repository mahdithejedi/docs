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