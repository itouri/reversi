package db

type (
	App struct {
		Name  string
		Port  uint
		Debug bool
	}

	Db struct {
		Host     string
		Port     uint
		User     string
		Pwd      string
		Database string
	}
)

var (
	APPCONFIG *App
	DBCONFIG  *Db
)

func init() {
	APPCONFIG = &App{
		Name:  "amazon",
		Port:  8006,
		Debug: true,
	}

	DBCONFIG = &Db{
		Host:     "127.0.0.1",
		Port:     27017,
		Database: "test",
	}
}
