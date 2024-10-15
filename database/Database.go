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

func GetData(table string, obj any) ([]any, error) {
    result := []any{}
    typ := reflect.TypeOf(obj)
    val := reflect.ValueOf(obj)

    // Extract JSON tags from the struct
    var attrs []string
    for i := 0; i < val.NumField(); i++ {
        tag := typ.Field(i).Tag.Get("json")
        if tag != "" {
            attrs = append(attrs, tag)
        }
    }
    if len(attrs) == 0 {
        return nil, nil
    }

    // Prepare SQL query
    query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(attrs, ","), table)
    rows, err := conn.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate the result
    for rows.Next() {
        // Create a new instance of the object
        newObj := reflect.New(typ).Elem()

        // Use reflection to scan the row into the object
        values := make([]any, len(attrs))
        for i := range values {
            values[i] = newObj.Field(i).Addr().Interface()
        }
        
        if err := rows.Scan(values...); err != nil {
            return nil, err
        }

        // Append the populated object to the result
        result = append(result, newObj.Interface())
    }

    // Check for any error encountered during iteration
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return result, nil
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