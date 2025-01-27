package storage

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbConfig struct {
	host     string
	port     int64
	database string
	username string
	password string
}

type DBs struct {
	ReaderDB gorm.DB
	WriterDB gorm.DB
}

func (d DBs) Migrate(model interface{}) {
	d.WriterDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model)
}
func initializeDb(config dbConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.username, config.password, config.host, config.port, config.database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		errMsg := "Failed to initialize db"
		log.Fatalf("%s", errMsg)
		panic(errMsg)
	}

	return db
}

func getDBConfigs(myEnv map[string]string) (dbConfig, dbConfig) {
	var (
		READER_HOST     string = myEnv["READER_DB_HOST"]
		READER_PORT     string = myEnv["READER_DB_PORT"]
		READER_USERNAME string = myEnv["READER_DB_USERNAME"]
		READER_PASSWORD string = myEnv["READER_DB_PASSWORD"]
		READER_DB       string = myEnv["READER_DB_NAME"]
	)

	reader_port, err := strconv.ParseInt(READER_PORT, 10, 16)
	if err != nil {
		panic("invalid value of READER_DB_PORT")
	}
	var readerConfigs dbConfig = dbConfig{
		host:     READER_HOST,
		port:     reader_port,
		username: READER_USERNAME,
		password: READER_PASSWORD,
		database: READER_DB,
	}

	var (
		WRITER_HOST     string = myEnv["WRITER_DB_HOST"]
		WRITER_PORT     string = myEnv["WRITER_DB_PORT"]
		WRITER_USERNAME string = myEnv["WRITER_DB_USERNAME"]
		WRITER_PASSWORD string = myEnv["WRITER_DB_PASSWORD"]
		WRITER_DB       string = myEnv["WRITER_DB_NAME"]
	)

	writer_port, err := strconv.ParseInt(WRITER_PORT, 10, 16)
	if err != nil {
		panic("invalid value of READER_DB_PORT")
	}
	var writerConfigs dbConfig = dbConfig{
		host:     WRITER_HOST,
		port:     writer_port,
		username: WRITER_USERNAME,
		password: WRITER_PASSWORD,
		database: WRITER_DB,
	}
	return readerConfigs, writerConfigs

}
func New(myEnv map[string]string) DBs {

	readerConfig, writerConfig := getDBConfigs(myEnv)

	readerGrom := initializeDb(readerConfig)
	writeRGrom := initializeDb(writerConfig)

	dbs := DBs{
		ReaderDB: *readerGrom,
		WriterDB: *writeRGrom,
	}

	return dbs

}
