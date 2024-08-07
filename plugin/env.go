/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-08-07 08:58:05
 * @LastEditors: lhl
 * @LastEditTime: 2024-08-07 09:02:18
 */
package plugin

import (
	"log"
	"os"
	"strconv"
)

func GetHpName() string {

	home := os.Getenv("HP_NAME")
	if home == "" {
		return "mysql"
	}

	return home

}

func GetHpPort() int {

	home := os.Getenv("HP_PORT")
	if home == "" {
		return 3306
	}

	num, _ := strconv.Atoi(home)
	return num

}

func GetHpPortStr() string {

	home := os.Getenv("HP_PORT")
	if home == "" {
		return "53"
	}

	return home

}

func GetHpLogPath() string {

	home := os.Getenv("HP_LOG_PATH")

	log.Println("lllllllll", home)

	if home == "" {
		return "./log/mysql.json"
	}

	return home

}

func GetLoginName() string {

	home := os.Getenv("LOGIN_NAME")
	if home == "" {
		return "root"
	}

	return home

}

func GetLoginPwd() string {

	home := os.Getenv("LOGIN_PWD")
	if home == "" {
		return "root"
	}

	return home

}
