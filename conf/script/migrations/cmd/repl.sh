#!/usr/bin/env bash

# echo "SYNAPSE_DB_SERVER_USERNAME=$SYNAPSE_DB_SERVER_USERNAME"
# echo "SYNAPSE_DB_ROOT_PASSWORD  =$SYNAPSE_DB_ROOT_PASSWORD"
# echo "SYNAPSE_DB_HOST           =$SYNAPSE_DB_HOST"
# echo "SYNAPSE_DB_PORT           =$SYNAPSE_DB_PORT"
# echo "SYNAPSE_DB_SCHEME         =$SYNAPSE_DB_SCHEME"

DIR="."
printf -v SOURCE "file://%s" "$DIR"
printf -v DATABASE "mysql://%s:%s@tcp(%s:%s)/%s" ${SYNAPSE_DB_SERVER_USERNAME} ${SYNAPSE_DB_ROOT_PASSWORD} ${SYNAPSE_DB_HOST} ${SYNAPSE_DB_PORT} ${SYNAPSE_DB_SCHEME}

echo "source  =\"$SOURCE\""
echo "database=\"$DATABASE\""

until false ; do 
    echo "type 'quit' or 'q' to quit"
    read -p "migrate > " CMD

    if [[ $CMD == "quit" ]] ; then 
        break
    fi

    if [[ $CMD == "q" ]] ; then 
         break
    fi

    IFS=' '
    read -ra CREATE_CMD <<< "$CMD"
    if [[ $CREATE_CMD == "create" ]] ; then 
        CREATE_ARG=$(sed -E 's|([^ ]+) ([^ ]+)|\2|' <<< $CMD)
        
        migrate create -ext sql -dir "$DIR" -seq "$CREATE_ARG"
    else
        migrate -source "$SOURCE" -database "$DATABASE" -lock-timeout 60 $up
    fi
    
    echo ""
done 


# read -p "Press Enter to Continue"

