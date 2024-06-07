package action

type VlanData struct {
	TimingRstWeek         string `json:"timing_rst_week"`
	DhcpLease             int    `json:"dhcp_lease"`
	TimingRstTime         string `json:"timing_rst_time"`
	CycleRstTime          int    `json:"cycle_rst_time"`
	PppoeService          string `json:"pppoe_service"`
	PppoeAc               string `json:"pppoe_ac"`
	Mtu                   int    `json:"mtu"`
	Mru                   int    `json:"mru"`
	DefaultRoute          int    `json:"default_route"`
	DiscAutoSwitch        int    `json:"disc_auto_switch"`
	LinkTime              string `json:"link_time"`
	CheckLinkMode         int    `json:"check_link_mode"`
	CheckLinkHost         string `json:"check_link_host"`
	QosSwitch             int    `json:"qos_switch"`
	PppoeCheckErripSwitch int    `json:"pppoe_check_errip_switch"`
	PppoeCheckErripList   string `json:"pppoe_check_errip_list"`
	DhcpDNS2              string `json:"dhcp_dns2"`
	DhcpDNS1              string `json:"dhcp_dns1"`
	DhcpUpdatetime        int    `json:"dhcp_updatetime"`
	DhcpGateway           string `json:"dhcp_gateway"`
	DhcpNetmask           string `json:"dhcp_netmask"`
	DhcpIPAddr            string `json:"dhcp_ip_addr"`
	PppoeMacremote        string `json:"pppoe_macremote"`
	ID                    int    `json:"id"`
	Comment               string `json:"comment"`
	Interface             string `json:"interface"`
	PppoeDNS2             string `json:"pppoe_dns2"`
	Enabled               string `json:"enabled"`
	PppoeIPAddr           string `json:"pppoe_ip_addr"`
	VlanID                string `json:"vlan_id"`
	PppoeUpdatetime       int    `json:"pppoe_updatetime"`
	VlanName              string `json:"vlan_name"`
	Mac                   string `json:"mac"`
	PppoeNetmask          string `json:"pppoe_netmask"`
	VlanInternet          int    `json:"vlan_internet"`
	Upload                int    `json:"upload"`
	Download              int    `json:"download"`
	QosUpload             int    `json:"qos_upload"`
	QosDownload           int    `json:"qos_download"`
	Vendorclass           string `json:"vendorclass"`
	Clientid              string `json:"clientid"`
	Hostname              string `json:"hostname"`
	OptType12             int    `json:"opt_type12"`
	OptType60             int    `json:"opt_type60"`
	OptType61             int    `json:"opt_type61"`
	IPMask                string `json:"ip_mask"`
	Gateway               string `json:"gateway"`
	Username              string `json:"username"`
	PppoeGateway          string `json:"pppoe_gateway"`
	Passwd                string `json:"passwd"`
	PppoeDNS1             string `json:"pppoe_dns1"`
	TimingRstSwitch       int    `json:"timing_rst_switch"`
}

type SnapshootLan struct {
	ID        int      `json:"id"`
	Comment   string   `json:"comment"`
	Interface string   `json:"interface"`
	Bandmode  int      `json:"bandmode"`
	Linkmode  int      `json:"linkmode"`
	Mac       string   `json:"mac"`
	Member    []string `json:"member"`
	IPAddr    string   `json:"ip_addr"`
	Netmask   string   `json:"netmask"`
}

type SnapshootWan struct {
	ID             int      `json:"id"`
	Comment        string   `json:"comment"`
	Interface      string   `json:"interface"`
	Mac            string   `json:"mac"`
	Member         []string `json:"member"`
	Bandmode       int      `json:"bandmode"`
	DefaultRoute   int      `json:"default_route"`
	Internet       int      `json:"internet"`
	IPAddr         string   `json:"ip_addr"`
	Netmask        string   `json:"netmask"`
	Gateway        string   `json:"gateway"`
	DNS1           string   `json:"dns1"`
	DNS2           string   `json:"dns2"`
	CountStatic    int      `json:"count_static"`
	CountDhcp      int      `json:"count_dhcp"`
	CountPppoe     int      `json:"count_pppoe"`
	CountCheckFail int      `json:"count_check_fail"`
	Updatetime     int      `json:"updatetime"`
	CheckRes       int      `json:"check_res"`
	Errmsg         string   `json:"errmsg"`
	Power          string   `json:"power"`
	Isp            string   `json:"isp"`
	Imei           string   `json:"imei"`
	Qnw            string   `json:"qnw"`
	Ccid           string   `json:"ccid"`
	Isnr           string   `json:"isnr"`
	Pcid           string   `json:"pcid"`
}

type ShowLanResult struct {
	Result
	Data struct {
		SnapshootLan []SnapshootLan `json:"snapshoot_lan"`
		SnapshootWan []SnapshootWan `json:"snapshoot_wan"`
	} `json:"Data"`
}

type ShowWanVlanListResult struct {
	Result
	Data struct {
		VlanData  []VlanData `json:"vlan_data"`
		VlanTotal int        `json:"vlan_total"`
	} `json:"Data"`
}

type ShowActionReqParam struct {
	ID           string `json:"id"`
	Type         string `json:"TYPE"`
	OrderBy      string `json:"ORDER_BY"`
	Order        string `json:"ORDER"`
	VlanInternet int    `json:"vlan_internet"`
	Interface    string `json:"interface"`
	Limit        string `json:"limit"`
}

func NewWanVlanListAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "wan",
		Param: map[string]interface{}{
			"ORDER":         "asc",
			"ORDER_BY":      "vlan_name",
			"TYPE":          "vlan_data,vlan_total",
			"interface":     "",
			"limit":         "0,100",
			"vlan_internet": "2",
		},
	}
}

func NewLanListAction() *Action {
	return &Action{
		Action:   "show",
		FuncName: "lan",
		Param: map[string]interface{}{
			"TYPE": "snapshoot",
		},
	}
}
