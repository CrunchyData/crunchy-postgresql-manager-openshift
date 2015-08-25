
build:
		rm -rf bin
		mkdir bin
		cp ../../sbin/setenv.sh ./bin
		cp ../../sbin/initdb.sh ./bin
		cp ../../sbin/startpg.sh ./bin
		cp ../../sbin/start-cpmcontainerapi.sh ./bin
		cp ../../sbin/fail-over.sh ./bin
		cp ../../sbin/badger-generate.sh ./bin
		cp ../../../../../../bin/cpmcontainerapi ./bin
		cp ../../../../../../bin/dummyserver ./bin
	        sudo docker build -t cpm-node-os .
		sudo docker tag -f cpm-node-os:latest crunchydata/cpm-node-os

