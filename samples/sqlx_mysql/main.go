// ref. https://qiita.com/rihofujino/items/b69e6a23e7cef1d692c4
package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//引っ張ってきたデータを当てはめる構造体を用意。
//その際、バッククオート（`）で、どのカラムと紐づけるのかを明示する。
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Userlist []User

func main() {

	//Userデータ一件一件を格納する配列Userlistを、Userlist型で用意
	var userlist Userlist

	//Mysqlに接続。sql.Openの代わりにsqlx.Openを使う。
	//ドライバ名、データソース名を引数に渡す
	// dsn spec: "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
	db, err := sqlx.Open("mysql", "root:asn10026900@/calico")
	if err != nil {
		log.Fatal(err)
	}

	//SELECTを実行。db.Queryの代わりにdb.Queryxを使う。
	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for rows.Next() {

		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}

	fmt.Println(userlist)
	//[{1 yamada 25} {2 suzuki 28}]

}
