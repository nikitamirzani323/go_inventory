package entities

type Model_admin struct {
	Username      string `json:"admin_username"`
	Nama          string `json:"admin_nama"`
	Rule          string `json:"admin_rule"`
	Joindate      string `json:"admin_joindate"`
	Timezone      string `json:"admin_timezone"`
	Lastlogin     string `json:"admin_lastlogin"`
	LastIpaddress string `json:"admin_lastipaddres"`
	Status        string `json:"admin_status"`
}
type Model_adminrule struct {
	Idrule string `json:"adminrule_idruleadmin"`
}

type Responseredis_adminhome struct {
	Admin_username     string `json:"admin_username"`
	Admin_nama         string `json:"admin_nama"`
	Admin_rule         string `json:"admin_rule"`
	Admin_joindate     string `json:"admin_joindate"`
	Admin_timezone     string `json:"admin_timezone"`
	Admin_lastlogin    string `json:"admin_lastlogin"`
	Admin_lastipaddres string `json:"admin_lastipaddres"`
	Admin_status       string `json:"admin_status"`
}
type Responseredis_adminrule struct {
	Adminrule_idrule string `json:"adminrule_idruleadmin"`
}
