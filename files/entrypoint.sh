#!/bin/bash

#zerotier-one
supervisord -c /etc/supervisor/supervisord.conf

# run the management program (sets up ZT, routes, etc)
manage