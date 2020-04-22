package dao

import "time"

// Posts 文章
type Posts struct {
	ID         int64     `db:"id" form:"id" json:"id"`
	Title      string    `db:"title" form:"title" json:"title"`
	Content    string    `db:"content" form:"content" json:"content"`
	ClickTimes int64     `db:"click_times" json:"clickTimes"`
	PostTime   time.Time `db:"post_time" json:"postTime"`
}

// InsertPost 添加文章
func InsertPost(post *Posts) (int64, error) {
	logger.Infof("post对象: ", post)
	ret, err := db.NamedExec("INSERT INTO posts (title, content) VALUES (:title, :content)", *post)
	if err != nil {
		logger.Errorf("添加文章错误:", err)
		return 0, err
	}
	pid, _ := ret.LastInsertId()
	pcount, _ := ret.RowsAffected()
	logger.Infof("发布了文章:%+v, id: %d", post, pid)
	return pcount, nil
}

// SelectPosts 查询所有内容
func SelectPosts(start int, offset int) (*[]Posts, error) {
	var posts = []Posts{}
	err := db.Select(&posts, "select * from posts limit ?, ?", start, offset)
	if err != nil {
		logger.Errorf("查询文章错误:", err)
		return &posts, err
	}
	return &posts, nil
}

// UpdatePost 修改文章
func UpdatePost(post *Posts) (int64, error) {
	ret, err := db.NamedExec("update posts set title = :title, content = :content where id = :id", *post)
	if err != nil {
		logger.Error("修改文章错误:", err)
		return 0, err
	}
	pid, _ := ret.LastInsertId()
	pcount, _ := ret.RowsAffected()
	logger.Infof("修改了文章:%+v, id: %d", post, pid)
	return pcount, nil
}

// DletePost 删除文章
func DletePost(post *Posts) (int64, error) {
	ret, err := db.Exec("delete from posts where id = ?", post.ID)
	if err != nil {
		logger.Error("删除文章错误:", err)
		return 0, err
	}
	pid, _ := ret.LastInsertId()
	pcount, _ := ret.RowsAffected()
	logger.Infof("删除了文章:%+v, id: %d", post, pid)
	return pcount, nil
}

// SelectPost 查询单个文章
func SelectPost(id int64) (*Posts, error) {
	post := Posts{}
	err := db.Get(&post, "select * from posts where id = ?", id)
	if err != nil {
		logger.Error("查询文章错误:", err)
		return &post, err
	}
	return &post, nil
}

// UpdateClickTime 更新点击次数
func UpdateClickTime(post *Posts) error {
	_, err := db.Exec("update posts set click_times = click_times + 1 where id = ?", post.ID)
	if err != nil {
		logger.Error("点击错误:", err)
		return err
	}
	logger.Info("点击了文章,", post)
	return nil
}
