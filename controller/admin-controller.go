package controller

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/inventory_api/entities"
	"github.com/nikitamirzani323/inventory_api/helpers"
	"github.com/nikitamirzani323/inventory_api/models"
)

const Fieldadmin_home_redis = "LISTADMIN_BACKEND"

func Adminhome(c *fiber.Ctx) error {
	var obj entities.Responseredis_adminhome
	var arraobj []entities.Responseredis_adminhome
	var obj_listruleadmin entities.Responseredis_adminrule
	var arraobj_listruleadmin []entities.Responseredis_adminrule
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldadmin_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listruleadmin_RD, _, _, _ := jsonparser.Get(jsonredis, "listruleadmin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		admin_username, _ := jsonparser.GetString(value, "admin_username")
		admin_nama, _ := jsonparser.GetString(value, "admin_nama")
		admin_rule, _ := jsonparser.GetString(value, "admin_rule")
		admin_joindate, _ := jsonparser.GetString(value, "admin_joindate")
		admin_timezone, _ := jsonparser.GetString(value, "admin_timezone")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_lastipaddres, _ := jsonparser.GetString(value, "admin_lastipaddres")
		admin_status, _ := jsonparser.GetString(value, "admin_status")

		obj.Admin_username = admin_username
		obj.Admin_nama = admin_nama
		obj.Admin_rule = admin_rule
		obj.Admin_joindate = admin_joindate
		obj.Admin_timezone = admin_timezone
		obj.Admin_lastlogin = admin_lastlogin
		obj.Admin_lastipaddres = admin_lastipaddres
		obj.Admin_status = admin_status
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listruleadmin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_name, _ := jsonparser.GetString(value, "adminrule_idruleadmin")

		obj_listruleadmin.Adminrule_idrule = adminrule_name
		arraobj_listruleadmin = append(arraobj_listruleadmin, obj_listruleadmin)
	})
	if !flag {
		result, err := models.Fetch_adminHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldadmin_home_redis, result, 30*time.Minute)
		log.Println("ADMIN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("ADMIN CACHE")
		return c.JSON(fiber.Map{
			"status":        fiber.StatusOK,
			"message":       "Success",
			"record":        arraobj,
			"listruleadmin": arraobj_listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	}
}
