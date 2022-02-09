package models

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/inventory_api/config"
	"github.com/nikitamirzani323/inventory_api/db"
	"github.com/nikitamirzani323/inventory_api/entities"
	"github.com/nikitamirzani323/inventory_api/helpers"
	"github.com/nleeper/goment"
)

func Fetch_adminHome() (helpers.ResponseAdmin, error) {
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var res helpers.ResponseAdmin
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			username , name, idadmin, statuslogin, 
			to_char(lastlogin,'YYYY-MM-DD HH24:mm:ss') as lastlogin, 
			to_char(joindate,'YYYY-MM-DD') as joindate, 
			ipaddress, timezone  
			FROM ` + config.DB_tbl_mst_admin + ` 
			ORDER BY lastlogin DESC 
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			username_db, name_db, idadminlevel_db                                string
			statuslogin_db, lastlogin_db, joindate_db, ipaddress_db, timezone_db string
		)

		err = row.Scan(
			&username_db, &name_db, &idadminlevel_db,
			&statuslogin_db, &lastlogin_db, &joindate_db,
			&ipaddress_db, &timezone_db)

		helpers.ErrorCheck(err)
		if statuslogin_db == "Y" {
			statuslogin_db = "ACTIVE"
		}
		if lastlogin_db == "0000-00-00 00:00:00" {
			lastlogin_db = ""
		}
		obj.Admin_username = username_db
		obj.Admin_nama = name_db
		obj.Admin_rule = idadminlevel_db
		obj.Admin_joindate = joindate_db
		obj.Admin_timezone = timezone_db
		obj.Admin_lastlogin = lastlogin_db
		obj.Admin_lastipaddres = ipaddress_db
		obj.Admin_status = statuslogin_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objRule entities.Model_adminrule
	var arraobjRule []entities.Model_adminrule
	sql_listrule := `SELECT 
		idadmin 	
		FROM ` + config.DB_tbl_mst_admingroup + ` 
	`
	row_listrule, err_listrule := con.QueryContext(ctx, sql_listrule)

	helpers.ErrorCheck(err_listrule)
	for row_listrule.Next() {
		var (
			idruleadmin_db string
		)

		err = row_listrule.Scan(&idruleadmin_db)

		helpers.ErrorCheck(err)

		objRule.Idrule = idruleadmin_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listrule = arraobjRule
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_admin(admin, username, password, idadminrule, name, status, ipaddress, timezone, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(config.DB_tbl_mst_admin, "username", username)
		if !flag {
			sql_insert := `
				insert into
				` + config.DB_tbl_mst_admin + ` (
					username , password, idadmin, name, statuslogin, 
					joindate, ipaddress, timezone, createadmin, createdateadmin 
				) values (
					$1, $2, $3, $4, $5,
					$6, $7, $8, $9, $10 
				)
			`
			hashpass := helpers.HashPasswordMD5(password)
			flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_admin, "INSERT",
				username, hashpass, idadminrule, name, status,
				tglnow.Format("YYYY-MM-DD"), ipaddress, timezone,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				flag = true
				msg = "Succes"
				log.Println(msg_insert)
			} else {
				log.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		if password != "" {
			hashpass := helpers.HashPasswordMD5(password)
			sql_update := `
				UPDATE
				` + config.DB_tbl_mst_admin + `
				SET password =$1, idadmin=$2, name=$3, statuslogin=$4, 
				updateadmin=$5, updatedateadmin=$6 
				WHERE username =$7 
			`

			flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_admin, "UPDATE",
				hashpass, idadminrule, name, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

			if flag_update {
				flag = true
				msg = "Succes"
				log.Println(msg_update)
			} else {
				log.Println(msg_update)
			}
		} else {
			sql_update := `
				UPDATE
				` + config.DB_tbl_mst_admin + `
				SET idadmin=$2, name=$3, statuslogin=$4, 
				updateadmin=$5, updatedateadmin=$6 
				WHERE username =$7 
			`

			flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_mst_admin, "UPDATE",
				idadminrule, name, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

			if flag_update {
				flag = true
				msg = "Succes"
				log.Println(msg_update)
			} else {
				log.Println(msg_update)
			}
		}
	}
	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Delete_admin(username string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()
	flag := false

	flag = CheckDB(config.DB_tbl_mst_admin, "username", username)
	if !flag {
		sql_delete := `
			DELETE FROM
			` + config.DB_tbl_mst_admin + ` 
			WHERE username = ? 
		`
		flag_delete, msg_delete := Exec_SQL(sql_delete, config.DB_tbl_mst_admin, "DELETE", username)

		if flag_delete {
			msg = "Succes"
			log.Println(msg_delete)
		} else {
			log.Println(msg_delete)
		}
	} else {
		msg = "Cannot Delete"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
