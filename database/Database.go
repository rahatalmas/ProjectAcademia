package database
import(
	"fmt"
	"database/sql"
	"strings"
    "sync"
    "reflect"
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

func GetData(table string, obj any) {
    result := []any{}
    typ := reflect.TypeOf(obj)
    val := reflect.ValueOf(obj)

    var attrs []string
    for i := 0; i < val.NumField(); i++ {
        tag := typ.Field(i).Tag.Get("json")
        if tag != "" {
            attrs = append(attrs, tag)
        }
    }
    if len(attrs) == 0 {
        return 
    }

    query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(attrs, ","), table)
    rows, err := conn.Query(query)
    if err != nil {
        return
    }
    defer rows.Close()
    fmt.Println(result)
    for rows.Next() {
        var temp []any
        for i:=0;i<len(attrs);i++{
            var str string 
            fmt.Println(&str)
            temp = append(temp,&str)
        }
        fmt.Println(temp)
        err := rows.Scan(temp...)
        if err!=nil{
            fmt.Println(err)
        }
        for i := range temp {
            if value, ok := temp[i].(*string); ok {
                fmt.Println(*value)
            } else {
                fmt.Println("Type assertion failed")
            }
        }
        
    }
}

func join(elems []string,ch string) string{
	result := ""
	for i,e:= range elems{
         if i>0{
			result+=ch
		 }
		 result+=e
	}
	return result
}
func searchChar(str string) bool{
	for _,c := range str{
		if c == '-' || c == ',' {
			return false
		}
	}
	return true
}