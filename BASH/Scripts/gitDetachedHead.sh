#! /usr/bin/sh

## source https://stackoverflow.com/questions/7069682/how-to-get-arguments-with-flags-in-bash

branch_check(){
	if [[ $(git branch --liist $arg1)=$arg1 > /dev/null ]]
	then
		return 1
	fi
	return 0
}




while test $# -gt 0
do
	case "$1" in
		-h| --help)
			echo "This script will resolve git DETACHED haed problem"
			echo " "
			echo "-h, --help	brief help"
			echo "-b, --branch 	the branch that have detached head"
			echo "-t, --temp	temperory branch for ???"
			exit 0
		;;
		-b| --branch)
			shift
			if test $# -gt 0
			then
				export BRANCH=$1
			fi
			shift
			;;
		-t| --temp)
			shift
			if test $# -gt 0
			then
				export TEMP=$1
			fi
			shift
			;;
		*)
		break
		;;
	esac
done

if [ -z "${BRANCH}" ]
then
	echo " no branch specified!"
	echo "specify a branch with '-b' of '--branch'"
	exit 0
elif [branch_check $BRANCH]:
then
	echo "branch is not exists"
	exit 0
fi
