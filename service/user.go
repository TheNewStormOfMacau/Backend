package service

import (
	"backend/global"
	"backend/model"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

func GetUserByAddr(addr string) *model.User {
	var user model.User
	result := global.G_DB.Where("chain_addr = ?", addr).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func GetRecordsByAddr(addr string) *[]model.Record {
	var records []model.Record
	result := global.G_DB.Where("chain_addr = ?", addr).Find(&records)
	if result.Error != nil {
		return nil
	}
	return &records
}

func UpdateUserInfo(from string, to string, amount string) {
	if (from == common.Address{0}.Hex()) {
		var user model.User
		result := global.G_DB.Where("chain_addr = ?", from).First(&user)
		if result.Error != nil {
			panic(result.Error)
		}
		amountInt, _ := strconv.Atoi(amount)
		user.Balance = user.Balance + amountInt
		user.Total = user.Total + amountInt
		global.G_DB.Save(&user)
	} else if (to == common.Address{0}.Hex()) {
		var user model.User
		result := global.G_DB.Where("chain_addr = ?", to).First(&user)
		if result.Error != nil {
			panic(result.Error)
		}
		amountInt, _ := strconv.Atoi(amount)
		user.Balance = user.Balance - amountInt
		user.Total = user.Total - amountInt
		global.G_DB.Save(&user)
	}
}

func UpdateGameResult(addr string, amount string) {
	var user model.User
	result := global.G_DB.Where("chain_addr = ?", addr).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	amountInt, _ := strconv.Atoi(amount)
	user.Balance = user.Balance + amountInt
	user.Total = user.Total + amountInt

	global.G_DB.Save(&user)
}
