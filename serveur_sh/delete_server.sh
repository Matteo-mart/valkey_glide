#!/bin/bash

echo "extinction du serveur"
sudo fuser -k 6379/tcp
echo "serveur éteint"