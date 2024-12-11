#!/bin/bash

# CREATED:
# yan.stalinsky.domrks@gmail.com
#
# Yan Stalinskiy
#
git pull origin devops-k8s
git pull origin master 
git push origin devops-k8s
docker build . -t gcr.io/pstracker-305808/carrick-api
docker push gcr.io/pstracker-305808/carrick-api