```shell
docker run --name mysql8 --env=MYSQL_ROOT_PASSWORD=123456 --volume=D:\Software\mysql:/var/lib/mysql -p 3306:3306 --restart=always -d mysql:8.4.3
```