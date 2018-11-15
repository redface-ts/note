1. 下边这种用法 checkTime变量乘以一秒要用time.Duration(）格式化下
time.NewTicker(time.Second * time.Duration(checkTime))
