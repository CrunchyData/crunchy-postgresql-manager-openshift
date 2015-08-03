
OSFLAVOR=centos7

all:
		rm -rf bin
		mkdir bin
		cp ../../config/.bash_profile conf/
		cp ../../sql/node-setup.sql bin/setup.sql
		cp ../../sbin/* bin/
		cp ../../config/postgresql.service bin/
	        sudo docker build -t cpm-base-os -f Dockerfile.$(OSFLAVOR) .
		sudo docker tag -f cpm-base-os:latest crunchydata/cpm-base-os

