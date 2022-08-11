#!/usr/bin/env bash

set -e

host="localhost"
shift
user="test"
shift
password="test"
shift

echo "Waiting for mysql"
# until mysql -h"$host" -u"$user" -p"$password" -D testdb -P 5306 &> /dev/null
# do
#         >&2 echo -n "."
#         sleep 1
# done

>&2 echo "MySQL is up - executing command"
exec api