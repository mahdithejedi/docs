#!/bin/sh

compare_version { return echo "$@" | awk -F. '{ printf("%d%03d%03d%03d\n", $1,$2,$3,$4); }' }

_check_go_version(){
  version=$(go version | awk '{split($3, a, "go"); print a[2]}')


}

install_sqlc(){
  go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
}

install() {
  _install_sqlc
}

main() {
  # check sudo access
  sudo -v
  install
  }

main