/*
 Copyright 2015 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package backup

import (
	"github.com/crunchydata/crunchy-postgresql-manager-openshift/logit"
	"github.com/crunchydata/crunchy-postgresql-manager-openshift/util"
)

type DefaultJob struct {
	request BackupRequest
}

//this is the func that implements the cron Job interface
func (t DefaultJob) Run() {
	logit.Info.Println("running ScheduleID:" + t.request.ScheduleID)
	dbConn, err := util.GetConnection(CLUSTERADMIN_DB)
	if err != nil {
		logit.Error.Println("BackupNow: error " + err.Error())
	}
	defer dbConn.Close()

	ProvisionBackupJob(dbConn, &t.request)
}
