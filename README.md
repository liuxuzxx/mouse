# Mouse
Go语言实验基地

cmake \
-DCMAKE_INSTALL_PREFIX=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30 \
-DMYSQL_DATADIR=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/data \
-DSYSCONFDIR=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30 \
-DMYSQL_UNIX_ADDR=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/data/mysql.sock \
-DWITH_DEBUG=1 \
-DDOWNLOAD_BOOST=1 -DWITH_BOOST=/media/liuxu/data/rondo/project/cpp/boost_1_59_0


--port=9527
--basedir=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30
--datadir=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/data
--general_log=0
--general_log_file=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/log/mysql/mysql-general.log
--log_error_verbosity=3
--log_error=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/log/mysql/mysql-error.log
--socket=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/log/mysql/mysql.sock
--pid-file=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/log/mysql/mysqld.pid
--debug=d:t:i:o,/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/log/trace/mysqld.trace
--console


./mysqld --basedir=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30 --datadir=/media/liuxu/data/rondo/project/cpp/mysql-5.7.30/data --initialize-insecure --user=shiyibo