#!/bin/bash
# CREATED:
#
# Yan Stalinskiy
#
file=hash.txt
cp .env.example .env
while IFS= read -r line
    do
 	    mkdir -p $pwd/hash/"$line"
        sed -i "s/MIX_PUBLISHER_ID=/MIX_PUBLISHER_ID=$line/" .env
        npm install && npm run prod
        cp ./dist/tracking.js $pwd/hash/"$line"/tracking.js
        sed -i "s/MIX_PUBLISHER_ID=$line/MIX_PUBLISHER_ID=/" .env
        echo "$line"
    done <"$file"
