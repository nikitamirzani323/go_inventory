package entities

type Model_admin struct {
	Admin_username     string `json:"admin_username"`
	Admin_nama         string `json:"admin_nama"`
	Admin_rule         string `json:"admin_rule"`
	Admin_joindate     string `json:"admin_joindate"`
	Admin_timezone     string `json:"admin_timezone"`
	Admin_lastlogin    string `json:"admin_lastlogin"`
	Admin_lastipaddres string `json:"admin_lastipaddres"`
	Admin_status       string `json:"admin_status"`
}
type Model_adminrule struct {
	Idrule string `json:"adminrule_idruleadmin"`
}

type Responseredis_adminrule struct {
	Adminrule_idrule string `json:"adminrule_idruleadmin"`
}

type Controller_adminsave struct {
	Sdata                 string `json:"sData" validate:"required"`
	Admin_client          string `json:"client" validate:"required"`
	Admin_username        string `json:"admin_username" validate:"required"`
	Admin_idadminrule     string `json:"admin_idadminrule" validate:"required"`
	Admin_password        string `json:"admin_password" `
	Admin_name            string `json:"admin_name" validate:"required"`
	Admin_statuslogin     string `json:"admin_statuslogin" validate:"required"`
	Admin_lastlogin       string `json:"admin_lastlogin"`
	Admin_joindate        string `json:"admin_joindate"`
	Admin_ipaddres        string `json:"admin_ipaddres"`
	Admin_timezone        string `json:"admin_timezone"`
	Admin_createadmin     string `json:"admin_createadmin"`
	Admin_createdateadmin string `json:"admin_createdateadmin"`
	Admin_updateadmin     string `json:"admin_updateadmin"`
	Admin_updatedateadmin string `json:"admin_updatedateadmin"`
}
