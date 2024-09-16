# Gin

https://goproxy.cn,direct

https://goproxy.io

```
go mod tidy
```



## Mysql

```bash
//下载依赖
go get -u github.com/go-sql-driver/mysql
```

```
//使用mysql驱动
func Open(driverName,dataSourceName string)(*DB error)
```

```
go run name.go //启动
```

# sqlx库使用指南

在项目中我们通常可能会使用`database/sql`连接MySQL数据库。本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和`DB.NamedExec`方法。

```
go get github.com/jmoiron/sqlx
```

## Redis

```
go get -u github.com/go-redis/redis
```



# SQL

1. **使用命令行连接 MySQL**：

   ```
   bashCopy Codemysql -u root -p
   ```

2. **创建数据库 `sql_test`**：

   ```
   sqlCopy CodeCREATE DATABASE sql_name;
   ```

3. **确认数据库创建成功**：

   ```sql
   sqlCopy CodeSHOW DATABASES;
   ```

```sql
USE sql_name;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
);
```



# zap

```
go get -u go.uber.org/zap
```

# viper

```
go get github.com/spf13/viper
```

