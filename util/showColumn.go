	// DB接続確認
	ping := todoDB.Ping()
	if ping != nil {
		return err
	}

	// show table info query
	rows, err := todoDB.Query("PRAGMA table_info(todos)")
	if err != nil {
		return err
	}
	defer rows.Close()

	// 結果の取得
	for rows.Next() {
		var cid int
		var name string
		var typ string
		var notnull int
		var dflt_value interface{}
		var pk int
		if err := rows.Scan(&cid, &name, &typ, &notnull, &dflt_value, &pk); err != nil {
			return err
		}
		fmt.Println(cid, name, typ, notnull, dflt_value, pk)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}