#!/bin/bash -x
#
# $1 - openshift URL
# $2 - username
# $3 - password
# $4 - object name
# $5 - object type
#

OPENSHIFT_URL=$1
OPENSHIFT_USERNAME=$2
OPENSHIFT_PASSWORD=$3
OBJECT_NAME=$4
OBJECT_TYPE=$5

/var/cpm/bin/oc login $OPENSHIFT_URL --insecure-skip-tls-verify=true \
	--username=$2 --password=$3

/var/cpm/bin/oc delete $OBJECT_TYPE $OBJECT_NAME
