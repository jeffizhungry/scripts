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
SHOW_FLAG=
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
    echo ""
    echo "--show           show horoku app config vars in local env"
    echo "--reset          clear heroku app config vars in local env"
    echo "--export         export heroku app config vars to a file"
    echo "--file           specify file to export to (default=$TMPFILE)"
    echo ""
    echo "examples:"
    echo "  source $0 --app <myapp>          # export config vars to local env"
    echo "  source $0 --app <myapp> --reset  # clear config vars from local env"
    echo "  $0 --app <myapp> --show          # show config vars in local env"
    echo "  $0 --app <myapp> --export        # export config vars to file"
    echo ""
}

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
        --show)
            SHOW_FLAG="true"
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
            echo "Error: Unknown option " $1 $2
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
    if [ -n "$SHOW_FLAG" ]; then
        echo "-- Show heroku app $APP config vars found in local environment"
        echo ""
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | python -c "import sys; import re; [sys.stdout.write('echo ' + re.match('^[A-Z_]*', line).group() + '=$' + re.match('^[A-Z_]*', line).group()+'\n') for line in sys.stdin if re.match('^[A-Z_]*=', line)]" > $TMPFILE
        source $TMPFILE
        cleanup
        return 0
    elif [ -n "$RESET_FLAG" ]; then
        echo "-- Reseting heroku app $APP config vars to empty strings"
        echo ""
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | sed -e "/^[A-Z_]*=/s/^/export /" | python -c "import sys; import re; [sys.stdout.write(re.match('^export [A-Z_]*=', line).group()+'\n') for line in sys.stdin if re.match('^export', line)]" > $TMPFILE
        source $TMPFILE
        cleanup
        return 0
    elif [ -n "$EXPORT_FLAG" ]; then
        echo "-- Exporting heroku app $APP config vars to: "$TMPFILE
        echo ""
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | sed -e "/^[A-Z_]*=/s/^/export /" > $TMPFILE
        return 0
    else
        echo "-- Exporting heroku app $APP config vars to: local environment"
        echo ""
        cleanup
        touch $TMPFILE
        heroku config -s --app $APP | sed -e "/^[A-Z_]*=/s/^/export /" > $TMPFILE
        source $TMPFILE
        cleanup
        return 0
    fi
} 

main "$@" || cleanup
