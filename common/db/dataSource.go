package datasource

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Init 初始化MySQL连接
func init() {
	fmt.Println("===============================================")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3307)/ruoyi?charset=utf8mb4")
}
