func DaoNotExists(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func Dao() (name string, err error) {
	db, _ := sql.Open("", "")
	err = errors.Wrap(db.QueryRow(`select name from user limit 1`).Scan(&name), "dao")
	return
}

func Service() {
	name, err := Dao()
	if DaoNotExists(err) {
		// 数据不存在，处理对应逻辑
		return
	}
	if err != nil {
		log.Printf("%+v\n", err)
		// 获取数据失败，处理对应逻辑
		return
	}
	// 数据处理
}
