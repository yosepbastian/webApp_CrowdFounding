package testing

import (
	"testing"
	"web-app-crowdfounding/config"
	"web-app-crowdfounding/manager"
)


func TestGetAllUser(t *testing.T){
	dbConfig := config.Config{
		DataSourceName: "root:@tcp(localhost:3306)/web_crowdfunding",
	}
	infraManager := manager.NewInfraManager(dbConfig)
	infraManager.SqlDb()

	
}