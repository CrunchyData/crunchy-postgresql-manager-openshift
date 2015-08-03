
build:
	        rm -rf bin
	        mkdir bin
	        cp ../../sql/clusteradmin.sql  ./bin/setup.sql
	        cp ../../sql/*  ./bin
	        cp ../../sbin/setenv.sh ./bin
	        cp ../../sbin/startadmin.sh ./bin
	        cp ../../sbin/start-pg-wrapper-admin.sh ./bin
	        cp ../../../../../../bin/adminapi  ./bin
	        cp ../../../../../../bin/dummyserver  ./bin
	        sudo docker build -t cpm-admin-os .
		sudo docker tag -f cpm-admin-os:latest crunchydata/cpm-admin-os

