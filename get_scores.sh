#!/bin/sh

if [ -z "${AOC_SESSION}" ]; then
  echo "Must set AOC_SESSION environment variable to the session string."
  exit 1
fi

cat > session.json <<EOF
{
    "__meta__": {
        "about": "HTTPie session file",
        "help": "https://httpie.org/doc#sessions",
        "httpie": "2.3.0"
    },
    "auth": {
        "password": null,
        "type": null,
        "username": null
    },
    "cookies": {
        "session": {
            "expires": null,
            "path": "/",
            "secure": false,
            "value": "${AOC_SESSION}"
        }
    },
    "headers": {}
}
EOF

http --session=./session.json https://adventofcode.com/2021/leaderboard/private/view/534400.json > 534400.json

rm session.json
