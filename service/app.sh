#!/bin/sh

mode=$1
dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )

addToService()
{
    command=$1
    dir=$2
    opt=$3
    sed 's|###COMMAND###|'"$command"'|g; s|###DIR###|'"$dir"'|g; s|###OPT###|'"$opt"'|g;' init.script > /etc/init.d/${command}
    chmod 755 /etc/init.d/${command}
    chkconfig $command on
}


case "$mode" in
    'install')
        addToService ipas-mcs ${dir}
        ;;

    'start')
        service ipas-mcs start
        ;;

    'stop')
        service ipas-mcs stop
        ;;
esac


