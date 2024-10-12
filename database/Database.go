package database
import(
	"fmt"
	"database/sql"
    "sync"

	_ "github.com/go-sql-driver/mysql"
)
var (
	conn  *sql.DB
	once  sync.Once
)

func DBInit() error{
	var err error
    once.Do(func(){
		conn, err = sql.Open("mysql","root:@/academia_db_main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.Ping()
		if err != nil{
			fmt.Println(err)
		}
	})
	return err
}

func GetDB() *sql.DB{
	return conn
}