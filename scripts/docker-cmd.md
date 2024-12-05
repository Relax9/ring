```shell
docker run --name mysql8 --env=MYSQL_ROOT_PASSWORD=123456 --volume=D:\Software\mysql:/var/lib/mysql -p 3306:3306 --restart=always -d mysql:8.4.3


docker run --name redis -p 6379:6379 -v D:\Software\redis:/usr/local/etc/redis -v D:\Software\redis\data:/usr/local/etc/redis/data --restart=always -d redis redis-server /usr/local/etc/redis/redis.conf
```