package main;

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
 // go mod download github.com/go-sql-driver/mysql
var db *sql.DB
// 定义结构体类型 来接受数据库查到的值
type user struct{
	id int
	name string
	age int
}
func initDB()(err error){
	// 数据库信息
	dsn := "root:haosql@tcp(127.0.0.1:3306)/golang"
	// Open方法 不会校验账号密码是否正确 只会校验格式
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验数据库信息是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	// defer db.Close()  // 注意这行代码要写在上面err判断的下面
	// db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	// db.SetMaxIdleConns(3) //设置连接池中的最大闲置连接数
	return 
}
// 单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果
func queryOne(id int){
	sqlStr := "select * from user where id = ?"
	var u user
	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr,id).Scan(&u.id,&u.name,&u.age)
	if err != nil{
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
// 多行查询db.Query()执行一次查询，返回多行结果（即Rows
func queryMore(n int){
	sqlStr := "select  * from user where age > ?"
	rows,err := db.Query(sqlStr,n)
	if err != nil{
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next(){
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
// 插入数据
func insertRow(name string,age int) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, name,age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
// 更新数据
func update(age ,id int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
// 删除数据
func delete(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
// sql预处理
// Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。
// 返回值可以同时执行多个查询和命令。
func sqlPrepare(id int){
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr) //提前将sql语句发给mysql
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id) //传值执行sql 修改 更新 删除同理
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=age - 2 where id=?"
	sqlStr2 := "Update user set age=age + 2 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Print("initDB err:%v",err)
	}
	fmt.Println("数据库连接成功")
	// queryOne(2)
	// queryMore(18)
	// insertRow("ll",22)
	// update(23,2)
	// delete(2)
	// sqlPrepare(2)  //sql预处理
	transactionDemo()  //事务
}