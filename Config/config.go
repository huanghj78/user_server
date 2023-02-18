package Config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	Server   *server
	LogConf  *logConf
	Database *database
)

type server struct {
	Address  string
	Port     int
	Frontend string
}

type logConf struct {
	FilePath string
	FileName string
}

type database struct {
	Address  string
	Port     int
	Database string
	Username string
	Password string
}

func init() {
	var err error
	conf, err := loadConfig("./Config/config.ini")
	if err != nil {
		log.Fatal("Fail to load config", err)
	}
	Server, err = (&server{}).initServerConfig(conf)
	if err != nil {
		log.Fatal("FAILED TO INIT SERVER CONFIG")
	}
	LogConf, err = (&logConf{}).initLogConfig(conf)
	if err != nil {
		log.Fatal("FAILED TO INIT LOG CONFIG")
	}
	Database, err = (&database{}).initDatabaseConfig(conf)
	if err != nil {
		log.Fatal("FAILED TO INIT DATABASE CONFIG")
	}

}

func (s *server) initServerConfig(conf *ini.File) (*server, error) {
	if conf == nil {
		return nil, errors.New("EMPTY CONFIG")
	}
	s.Address = conf.Section("server").Key("address").MustString("0.0.0.0")
	s.Port = conf.Section("server").Key("port").MustInt(8080)
	s.Frontend = conf.Section("server").Key("frontend").MustString("./Frontend/dist")
	return s, nil
}

func (s *logConf) initLogConfig(conf *ini.File) (*logConf, error) {
	if conf == nil {
		return nil, errors.New("EMPTY CONFIG")
	}
	s.FilePath = conf.Section("log").Key("filepath").MustString("./logs/")
	s.FileName = conf.Section("log").Key("filename").MustString("user_server")
	return s, nil
}

func (s *database) initDatabaseConfig(conf *ini.File) (*database, error) {
	if conf == nil {
		return nil, errors.New("EMPTY CONFIG")
	}
	s.Address = conf.Section("database").Key("address").MustString("127.0.0.1")
	s.Port = conf.Section("database").Key("port").MustInt(3306)
	s.Database = conf.Section("database").Key("database").MustString("")
	s.Username = conf.Section("database").Key("username").MustString("")
	s.Password = conf.Section("database").Key("password").MustString("")
	return s, nil
}

func loadConfig(path string) (*ini.File, error) {
	if exists, err := pathExists(path); !exists {
		return nil, err
	}
	conf, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
