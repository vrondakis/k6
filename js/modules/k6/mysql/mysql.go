package mysql
	
import (
	"context"
	"database/sql"
	gomysql "github.com/go-sql-driver/mysql"
)

type MYSQL struct{}

type MySQLConnection struct {
	connection *sql.DB
}

type MySQLResponse struct {
	Status  bool             
	Error   string          
	Db      *MySQLConnection
	Data    string     		
	ctx 	context.Context
}

func New() *MYSQL {
	return &MYSQL{}
}



func (db *MySQLConnection) Close(ctx context.Context) (bool){
	if db.connection != nil {
		db.connection.Close()
	}

	return true
}


func (db *MySQLConnection) Query(ctx context.Context, query string) (*MySQLResponse, error){
	res, err := db.connection.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()
	
	return nil, nil
}

func (*MYSQL) Connect(ctx context.Context, host string, username string, password string, database string) (*MySQLResponse, error) {
	config := gomysql.Config{
		User 	: username,
		Passwd 	: password,
		Net 	: "tcp",
		Addr	: host,
		DBName	: database,
		AllowNativePasswords : true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())

	if err == nil { err = db.Ping() }
	if err != nil {
		return nil, err
	}

	databaseConnection := &MySQLConnection{
		connection : db,
	}

	sqlResponse := MySQLResponse{
		Status 			: true,
		Db 		 		: databaseConnection,
		ctx 			: ctx,
	}

	return &sqlResponse, nil
}