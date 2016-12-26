#/usr/bin/env bash

###############################################################################
#
# This script is for exporting heroku config vars to your local environment.
#
# Use wisely...
#
###############################################################################

# -------------------------------------
# Variables
# -------------------------------------
TMPFILE="$2-config.sh"
APP=
RESET_FLAG=
EXPORT_FLAG=
VERBOSE_FLAG=

# -------------------------------------
# Helper Functions
# -------------------------------------
function cleanup {
    rm -rf $TMPFILE
}

function print_help() {
    echo "This script is used for export heroku config variables to your"\
         "local environment. \nUse wisely..."
    echo ""
    echo "options:"
    echo "-h, --help       show help menu"
    echo "--app            specify heroku app"
    echo "--export         export heroku config vars to a file"
    echo "--file           specify file to export to (default=$TMPFILE)"
}

# Run this in a subshell so we can capture the return code and cleanup
function main {
    # -------------------------------------
    # Get Arguments
    # -------------------------------------

    while [[ $# -gt 0 ]]
    do
    key="$1"
    case $key in
        -h|--help)
            print_help
            return 0
            ;;
        --app)
            APP=$2
            shift
            ;;
        --reset)
            RESET_FLAG="true"
            ;;
        --export)
            EXPORT_FLAG="true"
            ;;
        --file)
            TMPFILE=$2
            shift
            ;;
        *)
            # unknown option
            echo "Error: Unknown option " $1 $1
            echo ""
            print_help
            return 1
        ;;
    esac
    shift # past argument or value
    done

    # If no app is specified just exit
    if [ -z "$APP" ]; then
        echo "ERROR: app was not specified"
        echo ""
        print_help
        return 1
    fi

    # -------------------------------------
    # Export Config Vars
    # -------------------------------------
    if [ -n "$EXPORT_FLAG" ]; then
        echo "Exporting heroku app $APP to config vars to: "$TMPFILE
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | sed -e "/^[A-Z_]*=/s/^/export /" > $TMPFILE
    else
        echo "Exporting heroku app $APP to config vars to: local environment"
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | sed -e "/^[A-Z_]*=/s/^/export /" > $TMPFILE
        source $TMPFILE
        cleanup
    fi
} || cleanup
