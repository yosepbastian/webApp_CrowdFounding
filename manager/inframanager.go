package manager

import (
	"fmt"
	"os"
	"web-app-crowdfounding/config"

	_ "github.com/lib/pq"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type infraManager struct {
	db *gorm.DB
	config config.Config

}
type InfraManager interface {
	SqlDb() *gorm.DB
}

func (i *infraManager) InitDb() {
	db, err := gorm.Open(mysql.Open(i.config.DataSourceName), &gorm.Config{})

	if err != nil {
			fmt.Println("Database connection error:", err)
			os.Exit(1)
	}

	i.db = db // Simpan koneksi database ke dalam variabel i.db
	
	fmt.Println("Database Connected")
}


func (i *infraManager) SqlDb() *gorm.DB {
	return  i.db
}

func NewInfraManager(config config.Config) *infraManager {
	infra := infraManager{
		config: config,
	}
	infra.InitDb()
	return &infra
}