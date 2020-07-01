# 启动 redis 服务
redis-server ./conf/redis.conf
# 启动 trackerd
fdfs_trackerd /home/lpgit/go/src/ihome/ihomeWeb/conf/tracker.conf restart
# 启动 storaged
fdfs_storaged /home/lpgit/go/src/ihome/ihomeWeb/conf/storage.conf restart
