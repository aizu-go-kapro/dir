#!/bin/sh
readonly HOST=${HOST}
readonly PASSWORD=${PASSWORD}
readonly SRC="PATH_WHAT_YOU_WANT"
readonly DEST="PATH_WEHERE_YOU_PLACE"

expect -c "
set timeout 4000
spawn scp -r $HOST:$SRC $DEST
expect \"Password:\" 
send -- \"${PASSWORD}\n\"
expect $
exit 0
"
