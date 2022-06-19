package database

import (
	"log"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	migrate := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn(&models.User{}, "username")
			},
		},
	})

	if migUp {
		if err := migrate.Migrate(); err != nil {
			return err
		}
		log.Fatal("migrated successfully")
		return nil
	}
	if migDown {
		if err := migrate.RollbackLast(); err != nil {
			return err
		}
		log.Fatal("Rollback successfully")
		return nil
	}

	migrate.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.Users{},
			&models.Vehicles{},
		)
		if err != nil {
			return err
		}
		return nil
	})

	if err := migrate.Migrate(); err != nil {
		return err
	}
	log.Fatal("init schema successfully")
	return nil
}
