
build:
		rm -rf bin
		mkdir bin
		cp ../../sbin/start-backupserver.sh ./bin
		cp ../../../../../../bin/backupserver ./bin
		cp ../../../../../../bin/dummyserver ./bin
	        sudo docker build -t cpm-backup-os .
		sudo docker tag -f cpm-backup-os:latest crunchydata/cpm-backup-os

