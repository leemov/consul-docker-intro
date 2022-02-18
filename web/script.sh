#!/bin/bash

echo "nameserver $(grep 'consul-server' /etc/hosts | cut -f1)" > /etc/resolv.conf
./main