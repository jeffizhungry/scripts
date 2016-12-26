#/usr/bin/env bash

TMPFILE="$2-config.sh"
function cleanup {
    rm -rf $TMPFILE    
}

while test $# -gt 0; do
    case "$1" in
        -h|--help)
            echo "$0 [options]"
            echo " "
            echo "options:"
            echo "-h, --help       show brief help"
            echo "--app            specify app to export config from"
            echo "--reset          clear config vars exported by this app"
            exit 0
            ;;
        --app)
            touch $TMPFILE
            heroku config --app $2 > $TMPFILE
            exit 0
            ;;
    esac
done

