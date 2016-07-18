# Gophr

A beautiful web app written in go with bootstrap

## Usage

## Installing

### Prepare Directory
```
mkdir data
mkdir -p data/images
```

### Go Get Packages
```
export GOPATH=/usr/local/go/
go get "github.com/go-sql-driver/mysql"
go get "github.com/julienschmidt/httprouter"
```

### MySQL
Reference:
1. [Tutorial Points: MySQL](http://www.tutorialspoint.com/mysql/mysql-installation.htm)
2. [MariaDB Installation](https://mariadb.com/kb/en/mariadb/yum/) 

```bash
#Install MariaDB server and client
yum groupinstall mariadb
#Start MariaDB
systemctl start mariadb.service
#Test install 
mysqladmin version
mysql
>SHOW DATABASES;
>exit
#Update password for database
#MySQL by default has an empty password
mysqladmin -u root password "Letmein123";
```

### MySQL Requirements
```
#mysql -u root -p "Letmein123"
CREATE DATABASE gophr;
USE gopher;
CREATE TABLE `images` (
	`id` varchar(255) NOT NULL DEFAULT '',
	`user_id` varchar(255) NOT NULL,
	`name` varchar(255) NOT NULL DEFAULT '',
	`location` varchar(255) NOT NULL DEFAULT '',
	`description` text NOT NULL,
	`size` int(11) NOT NULL,
	`created_at` datetime NOT NULL,
	PRIMARY KEY (`id`),
	KEY `user_id_idx` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
#DROP TABLE images
```

## Reference:

1. [Level Up Your Web Apps With Go](https://github.com/spbooks/go1)
2. [Mastering Go Web Services] (https://www.safaribooksonline.com/library/view/mastering-go-web/)
