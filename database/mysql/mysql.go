package mysql

import (
	"fmt"

	"github.com/fazaalexander/montirku-be/config"
	"github.com/fazaalexander/montirku-be/database/seeds"
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
	re "github.com/fazaalexander/montirku-be/modules/entity/role"
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seeds.DBSeed(DB)
}

func InitDB() {
	var err error

	config := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		re.Role{},
		ue.User{},
		ue.UserDetail{},
		ue.UserRecovery{},
		bt.BengkelCategory{},
		bt.Bengkel{},
		bt.BengkelAddress{},
		bt.OperationalTime{},
		bt.BengkelServices{},
		bt.BengkelRating{},
		te.Transaction{},
		te.TransactionDetail{},
	)
	DB.Migrator().HasConstraint(&ue.User{}, "UserDetail")
	DB.Migrator().HasConstraint(&re.Role{}, "Users")
	DB.Migrator().HasConstraint(&bt.BengkelServices{}, "TransactionDetail")
}
