#!/bin/bash

## source https://stackoverflow.com/questions/7069682/how-to-get-arguments-with-flags-in-bash


check_detached() {
	if [ -z "$(git symbolic-ref -q HEAD)"]
	then
		return 1
	else
		return 0
	fi
}


branch_check(){
	# echo "ARGS is ****$args1 ******"
	if [[ $(git branch --list $arg1)=$arg1 > /dev/null ]]
	then
		return 1
	fi
	return 0
}

temp_set(){

	TEMP='__temp__'
	if  branch_check "${TEMP}"
	then
		echo "defualt branch __temp__ is exists"
		echo "please specify a temperory branch name with -b"
		exit 0
	fi

}

#TODO check if git is installed
#if ! [ -x "$(command -v git)"]
#then
#	echo "git is not installed"
#	exit 0
# TODO proper way to chech if git is intialized
# elif [ "$(git branch 1>/dev/null)" == '\n']
# then
#	echo "git it not initialied"
# 	exit 0
# fi


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
elif branch_check "${BRANCH}"
then
	echo "$BRANCH branch is not exists"
	exit 0
fi
if [ -z "${TEMP}" ]
then
	temp_set
	
elif ! branch_check "${TEMP}"
then
	echo "$TEMP branch is exist"
       	echo "defualt will be set"
	temp_set
fi

if check_detached
then
	echo "You do not face any detached head problem"
	exit 0  
fi

git checkout -b $TEMP 1>/dev/null
git branch -f $BRANCH $TEMP 1>/dev/null
git checkout $BRANCH 1>/dev/null
git checkout -B $BRANCH $TEMP 1>/dev/null
git branch -d $TEMP 1>/dev/null

