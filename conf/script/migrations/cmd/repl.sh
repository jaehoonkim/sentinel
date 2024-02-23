#!/usr/bin/env bash

# echo "MORPHEUS_DB_SERVER_USERNAME=$MORPHEUS_DB_SERVER_USERNAME"
# echo "MORPHEUS_DB_ROOT_PASSWORD  =$MORPHEUS_DB_ROOT_PASSWORD"
# echo "MORPHEUS_DB_HOST           =$MORPHEUS_DB_HOST"
# echo "MORPHEUS_DB_PORT           =$MORPHEUS_DB_PORT"
# echo "MORPHEUS_DB_SCHEME         =$MORPHEUS_DB_SCHEME"

DIR="."
printf -v SOURCE "file://%s" "$DIR"
printf -v DATABASE "mysql://%s:%s@tcp(%s:%s)/%s" ${MORPHEUS_DB_SERVER_USERNAME} ${MORPHEUS_DB_ROOT_PASSWORD} ${MORPHEUS_DB_HOST} ${MORPHEUS_DB_PORT} ${MORPHEUS_DB_SCHEME}

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

