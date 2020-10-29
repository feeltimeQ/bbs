package logic

import (
	"backend/bluebell/dao/mysql"
	"backend/bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查数据库 查找所以的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
