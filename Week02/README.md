学习笔记

## Week02 作业题目：

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 解

错误信息只消费一遍，package级别的err一般透传到最外层处理，日志在一个地方集中打印
```
func Biz() error {
	id := uint64(2333)
	user, err := Dao(id)
	if err != nil {
        // 调试的时候可以使用不同的日志级别
		log.Info("Error: Dao has error.userID=%d,err=%+v\n", id, err)
		return err
	}

    //handle user
    println(user)
}

// Dao query user info
func Dao(id uint64) (*User, error) {
    user:= new(User)
	// access DB...
    err := orm.Where("id = ?", id).Get(&user)
	// sql.ErrNoRows
	return nil, errors.Errorf("query user info has error.userID=%d,err=%v", id, err)
}
//查不到数据，先检查error!=nil,再检查User!=nil,查不到数据也算是一种错误，继续往外抛
```

