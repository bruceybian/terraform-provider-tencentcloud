package tencentcloud

const (
	ADDRESS_TEMPLATE_TYPE_1 = 1
	ADDRESS_TEMPLATE_TYPE_5 = 5
)

var ADDRESS_TEMPLATE_TYPE = []int{
	ADDRESS_TEMPLATE_TYPE_1,
	ADDRESS_TEMPLATE_TYPE_5,
}

const (
	RULE_TYPE_1 = 1
	RULE_TYPE_2 = 2
)

var RULE_TYPE = []int{
	RULE_TYPE_1,
	RULE_TYPE_2,
}

const (
	DIRECTION_0 = "0"
	DIRECTION_1 = "1"
	DIRECTION_3 = "3"
)

var DIRECTION = []string{
	DIRECTION_0,
	DIRECTION_1,
	DIRECTION_3,
}

const (
	MODE_0 = 0
	MODE_1 = 1
)

var MODE = []int{
	MODE_0,
	MODE_1,
}

const (
	CROSS_A_ZONE_0 = 0
	CROSS_A_ZONE_1 = 1
)

var CROSS_A_ZONE = []int{
	CROSS_A_ZONE_0,
	CROSS_A_ZONE_1,
}

var ZONE_MAP_EN2CN = map[string]string{
	"ap-guangzhou-1":     "广州一区",
	"ap-guangzhou-2":     "广州二区",
	"ap-guangzhou-3":     "广州三区",
	"ap-guangzhou-4":     "广州四区",
	"ap-guangzhou-5":     "广州五区",
	"ap-guangzhou-6":     "广州六区",
	"ap-guangzhou-7":     "广州七区",
	"ap-guangzhou-8":     "广州八区",
	"ap-guangzhou-9":     "广州九区",
	"ap-shenzhen-fsi-1":  "深圳金融一区",
	"ap-shenzhen-fsi-2":  "深圳金融二区",
	"ap-shenzhen-fsi-3":  "深圳金融三区",
	"ap-shenzhen-fsi-4":  "深圳金融四区",
	"ap-shenzhen-fsi-5":  "深圳金融五区",
	"ap-shenzhen-fsi-6":  "深圳金融六区",
	"ap-shenzhen-fsi-7":  "深圳金融七区",
	"ap-shenzhen-fsi-8":  "深圳金融八区",
	"ap-shenzhen-fsi-9":  "深圳金融九区",
	"ap-shanghai-1":      "上海一区",
	"ap-shanghai-2":      "上海二区",
	"ap-shanghai-3":      "上海三区",
	"ap-shanghai-4":      "上海四区",
	"ap-shanghai-5":      "上海五区",
	"ap-shanghai-6":      "上海六区",
	"ap-shanghai-7":      "上海七区",
	"ap-shanghai-8":      "上海八区",
	"ap-shanghai-9":      "上海九区",
	"ap-shanghai-fsi-1":  "上海金融一区",
	"ap-shanghai-fsi-2":  "上海金融二区",
	"ap-shanghai-fsi-3":  "上海金融三区",
	"ap-shanghai-fsi-4":  "上海金融四区",
	"ap-shanghai-fsi-5":  "上海金融五区",
	"ap-shanghai-fsi-6":  "上海金融六区",
	"ap-shanghai-fsi-7":  "上海金融七区",
	"ap-shanghai-fsi-8":  "上海金融八区",
	"ap-shanghai-fsi-9":  "上海金融九区",
	"ap-nanjing-1":       "南京一区",
	"ap-nanjing-2":       "南京二区",
	"ap-nanjing-3":       "南京三区",
	"ap-nanjing-4":       "南京四区",
	"ap-nanjing-5":       "南京五区",
	"ap-nanjing-6":       "南京六区",
	"ap-nanjing-7":       "南京七区",
	"ap-nanjing-8":       "南京八区",
	"ap-nanjing-9":       "南京九区",
	"ap-beijing-1":       "北京一区",
	"ap-beijing-2":       "北京二区",
	"ap-beijing-3":       "北京三区",
	"ap-beijing-4":       "北京四区",
	"ap-beijing-5":       "北京五区",
	"ap-beijing-6":       "北京六区",
	"ap-beijing-7":       "北京七区",
	"ap-beijing-8":       "北京八区",
	"ap-beijing-9":       "北京九区",
	"ap-chengdu-1":       "成都一区",
	"ap-chengdu-2":       "成都二区",
	"ap-chengdu-3":       "成都三区",
	"ap-chengdu-4":       "成都四区",
	"ap-chengdu-5":       "成都五区",
	"ap-chengdu-6":       "成都六区",
	"ap-chengdu-7":       "成都七区",
	"ap-chengdu-8":       "成都八区",
	"ap-chengdu-9":       "成都九区",
	"ap-chongqing-1":     "重庆一区",
	"ap-chongqing-2":     "重庆二区",
	"ap-chongqing-3":     "重庆三区",
	"ap-chongqing-4":     "重庆四区",
	"ap-chongqing-5":     "重庆五区",
	"ap-chongqing-6":     "重庆六区",
	"ap-chongqing-7":     "重庆七区",
	"ap-chongqing-8":     "重庆八区",
	"ap-chongqing-9":     "重庆九区",
	"ap-hongkong-1":      "香港一区",
	"ap-hongkong-2":      "香港二区",
	"ap-hongkong-3":      "香港三区",
	"ap-hongkong-4":      "香港四区",
	"ap-hongkong-5":      "香港五区",
	"ap-hongkong-6":      "香港六区",
	"ap-hongkong-7":      "香港七区",
	"ap-hongkong-8":      "香港八区",
	"ap-hongkong-9":      "香港九区",
	"ap-singapore-1":     "新加坡一区",
	"ap-singapore-2":     "新加坡二区",
	"ap-singapore-3":     "新加坡三区",
	"ap-singapore-4":     "新加坡四区",
	"ap-singapore-5":     "新加坡五区",
	"ap-singapore-6":     "新加坡六区",
	"ap-singapore-7":     "新加坡七区",
	"ap-singapore-8":     "新加坡八区",
	"ap-singapore-9":     "新加坡九区",
	"ap-seoul-1":         "首尔一区",
	"ap-seoul-2":         "首尔二区",
	"ap-seoul-3":         "首尔三区",
	"ap-seoul-4":         "首尔四区",
	"ap-seoul-5":         "首尔五区",
	"ap-seoul-6":         "首尔六区",
	"ap-seoul-7":         "首尔七区",
	"ap-seoul-8":         "首尔八区",
	"ap-seoul-9":         "首尔九区",
	"ap-tokyo-1":         "东京一区",
	"ap-tokyo-2":         "东京二区",
	"ap-tokyo-3":         "东京三区",
	"ap-tokyo-4":         "东京四区",
	"ap-tokyo-5":         "东京五区",
	"ap-tokyo-6":         "东京六区",
	"ap-tokyo-7":         "东京七区",
	"ap-tokyo-8":         "东京八区",
	"ap-tokyo-9":         "东京九区",
	"ap-mumbai-1":        "孟买一区",
	"ap-mumbai-2":        "孟买二区",
	"ap-mumbai-3":        "孟买三区",
	"ap-mumbai-4":        "孟买四区",
	"ap-mumbai-5":        "孟买五区",
	"ap-mumbai-6":        "孟买六区",
	"ap-mumbai-7":        "孟买七区",
	"ap-mumbai-8":        "孟买八区",
	"ap-mumbai-9":        "孟买九区",
	"ap-bangkok-1":       "曼谷一区",
	"ap-bangkok-2":       "曼谷二区",
	"ap-bangkok-3":       "曼谷三区",
	"ap-bangkok-4":       "曼谷四区",
	"ap-bangkok-5":       "曼谷五区",
	"ap-bangkok-6":       "曼谷六区",
	"ap-bangkok-7":       "曼谷七区",
	"ap-bangkok-8":       "曼谷八区",
	"ap-bangkok-9":       "曼谷九区",
	"ap-jakarta-1":       "雅加达一区",
	"ap-jakarta-2":       "雅加达二区",
	"ap-jakarta-3":       "雅加达三区",
	"ap-jakarta-4":       "雅加达四区",
	"ap-jakarta-5":       "雅加达五区",
	"ap-jakarta-6":       "雅加达六区",
	"ap-jakarta-7":       "雅加达七区",
	"ap-jakarta-8":       "雅加达八区",
	"sa-saopaulo-1":      "圣保罗一区",
	"sa-saopaulo-2":      "圣保罗二区",
	"sa-saopaulo-3":      "圣保罗三区",
	"sa-saopaulo-4":      "圣保罗四区",
	"sa-saopaulo-5":      "圣保罗五区",
	"sa-saopaulo-6":      "圣保罗六区",
	"sa-saopaulo-7":      "圣保罗七区",
	"sa-saopaulo-8":      "圣保罗八区",
	"na-toronto-1":       "多伦多一区",
	"na-toronto-2":       "多伦多二区",
	"na-toronto-3":       "多伦多三区",
	"na-toronto-4":       "多伦多四区",
	"na-toronto-5":       "多伦多五区",
	"na-toronto-6":       "多伦多六区",
	"na-toronto-7":       "多伦多七区",
	"na-toronto-8":       "多伦多八区",
	"na-toronto-9":       "多伦多九区",
	"na-siliconvalley-1": "硅谷一区",
	"na-siliconvalley-2": "硅谷二区",
	"na-siliconvalley-3": "硅谷三区",
	"na-siliconvalley-4": "硅谷四区",
	"na-siliconvalley-5": "硅谷五区",
	"na-siliconvalley-6": "硅谷六区",
	"na-siliconvalley-7": "硅谷七区",
	"na-siliconvalley-8": "硅谷八区",
	"na-siliconvalley-9": "硅谷九区",
	"na-ashburn-1":       "弗吉尼亚一区",
	"na-ashburn-2":       "弗吉尼亚二区",
	"na-ashburn-3":       "弗吉尼亚三区",
	"na-ashburn-4":       "弗吉尼亚四区",
	"na-ashburn-5":       "弗吉尼亚五区",
	"na-ashburn-6":       "弗吉尼亚六区",
	"na-ashburn-7":       "弗吉尼亚七区",
	"na-ashburn-8":       "弗吉尼亚八区",
	"na-ashburn-9":       "弗吉尼亚九区",
	"eu-frankfurt-1":     "法兰克福一区",
	"eu-frankfurt-2":     "法兰克福二区",
	"eu-frankfurt-3":     "法兰克福三区",
	"eu-frankfurt-4":     "法兰克福四区",
	"eu-frankfurt-5":     "法兰克福五区",
	"eu-frankfurt-6":     "法兰克福六区",
	"eu-frankfurt-7":     "法兰克福七区",
	"eu-frankfurt-8":     "法兰克福八区",
	"eu-frankfurt-9":     "法兰克福九区",
	"eu-moscow-1":        "莫斯科一区",
	"eu-moscow-2":        "莫斯科二区",
	"eu-moscow-3":        "莫斯科三区",
	"eu-moscow-4":        "莫斯科四区",
	"eu-moscow-5":        "莫斯科五区",
	"eu-moscow-6":        "莫斯科六区",
	"eu-moscow-7":        "莫斯科七区",
	"eu-moscow-8":        "莫斯科八区",
	"eu-moscow-9":        "莫斯科九区",
	"ap-shanghai-adc-1":  "上海自动驾驶云一区",
	"ap-taipei-1":        "中国台北一区",
	"ap-taipei-2":        "中国台北二区",
	"ap-taipei-3":        "中国台北三区",
	"ap-taipei-4":        "中国台北四区",
	"ap-taipei-5":        "中国台北五区",
	"ap-taipei-6":        "中国台北六区",
}

var ZONE_MAP_CN2EN = map[string]string{
	"广州一区":      "ap-guangzhou-1",
	"广州二区":      "ap-guangzhou-2",
	"广州三区":      "ap-guangzhou-3",
	"广州四区":      "ap-guangzhou-4",
	"广州五区":      "ap-guangzhou-5",
	"广州六区":      "ap-guangzhou-6",
	"广州七区":      "ap-guangzhou-7",
	"广州八区":      "ap-guangzhou-8",
	"广州九区":      "ap-guangzhou-9",
	"深圳金融一区":    "ap-shenzhen-fsi-1",
	"深圳金融二区":    "ap-shenzhen-fsi-2",
	"深圳金融三区":    "ap-shenzhen-fsi-3",
	"深圳金融四区":    "ap-shenzhen-fsi-4",
	"深圳金融五区":    "ap-shenzhen-fsi-5",
	"深圳金融六区":    "ap-shenzhen-fsi-6",
	"深圳金融七区":    "ap-shenzhen-fsi-7",
	"深圳金融八区":    "ap-shenzhen-fsi-8",
	"深圳金融九区":    "ap-shenzhen-fsi-9",
	"上海一区":      "ap-shanghai-1",
	"上海二区":      "ap-shanghai-2",
	"上海三区":      "ap-shanghai-3",
	"上海四区":      "ap-shanghai-4",
	"上海五区":      "ap-shanghai-5",
	"上海六区":      "ap-shanghai-6",
	"上海七区":      "ap-shanghai-7",
	"上海八区":      "ap-shanghai-8",
	"上海九区":      "ap-shanghai-9",
	"上海金融一区":    "ap-shanghai-fsi-1",
	"上海金融二区":    "ap-shanghai-fsi-2",
	"上海金融三区":    "ap-shanghai-fsi-3",
	"上海金融四区":    "ap-shanghai-fsi-4",
	"上海金融五区":    "ap-shanghai-fsi-5",
	"上海金融六区":    "ap-shanghai-fsi-6",
	"上海金融七区":    "ap-shanghai-fsi-7",
	"上海金融八区":    "ap-shanghai-fsi-8",
	"上海金融九区":    "ap-shanghai-fsi-9",
	"南京一区":      "ap-nanjing-1",
	"南京二区":      "ap-nanjing-2",
	"南京三区":      "ap-nanjing-3",
	"南京四区":      "ap-nanjing-4",
	"南京五区":      "ap-nanjing-5",
	"南京六区":      "ap-nanjing-6",
	"南京七区":      "ap-nanjing-7",
	"南京八区":      "ap-nanjing-8",
	"南京九区":      "ap-nanjing-9",
	"北京一区":      "ap-beijing-1",
	"北京二区":      "ap-beijing-2",
	"北京三区":      "ap-beijing-3",
	"北京四区":      "ap-beijing-4",
	"北京五区":      "ap-beijing-5",
	"北京六区":      "ap-beijing-6",
	"北京七区":      "ap-beijing-7",
	"北京八区":      "ap-beijing-8",
	"北京九区":      "ap-beijing-9",
	"成都一区":      "ap-chengdu-1",
	"成都二区":      "ap-chengdu-2",
	"成都三区":      "ap-chengdu-3",
	"成都四区":      "ap-chengdu-4",
	"成都五区":      "ap-chengdu-5",
	"成都六区":      "ap-chengdu-6",
	"成都七区":      "ap-chengdu-7",
	"成都八区":      "ap-chengdu-8",
	"成都九区":      "ap-chengdu-9",
	"重庆一区":      "ap-chongqing-1",
	"重庆二区":      "ap-chongqing-2",
	"重庆三区":      "ap-chongqing-3",
	"重庆四区":      "ap-chongqing-4",
	"重庆五区":      "ap-chongqing-5",
	"重庆六区":      "ap-chongqing-6",
	"重庆七区":      "ap-chongqing-7",
	"重庆八区":      "ap-chongqing-8",
	"重庆九区":      "ap-chongqing-9",
	"香港一区":      "ap-hongkong-1",
	"香港二区":      "ap-hongkong-2",
	"香港三区":      "ap-hongkong-3",
	"香港四区":      "ap-hongkong-4",
	"香港五区":      "ap-hongkong-5",
	"香港六区":      "ap-hongkong-6",
	"香港七区":      "ap-hongkong-7",
	"香港八区":      "ap-hongkong-8",
	"香港九区":      "ap-hongkong-9",
	"新加坡一区":     "ap-singapore-1",
	"新加坡二区":     "ap-singapore-2",
	"新加坡三区":     "ap-singapore-3",
	"新加坡四区":     "ap-singapore-4",
	"新加坡五区":     "ap-singapore-5",
	"新加坡六区":     "ap-singapore-6",
	"新加坡七区":     "ap-singapore-7",
	"新加坡八区":     "ap-singapore-8",
	"新加坡九区":     "ap-singapore-9",
	"首尔一区":      "ap-seoul-1",
	"首尔二区":      "ap-seoul-2",
	"首尔三区":      "ap-seoul-3",
	"首尔四区":      "ap-seoul-4",
	"首尔五区":      "ap-seoul-5",
	"首尔六区":      "ap-seoul-6",
	"首尔七区":      "ap-seoul-7",
	"首尔八区":      "ap-seoul-8",
	"首尔九区":      "ap-seoul-9",
	"东京一区":      "ap-tokyo-1",
	"东京二区":      "ap-tokyo-2",
	"东京三区":      "ap-tokyo-3",
	"东京四区":      "ap-tokyo-4",
	"东京五区":      "ap-tokyo-5",
	"东京六区":      "ap-tokyo-6",
	"东京七区":      "ap-tokyo-7",
	"东京八区":      "ap-tokyo-8",
	"东京九区":      "ap-tokyo-9",
	"孟买一区":      "ap-mumbai-1",
	"孟买二区":      "ap-mumbai-2",
	"孟买三区":      "ap-mumbai-3",
	"孟买四区":      "ap-mumbai-4",
	"孟买五区":      "ap-mumbai-5",
	"孟买六区":      "ap-mumbai-6",
	"孟买七区":      "ap-mumbai-7",
	"孟买八区":      "ap-mumbai-8",
	"孟买九区":      "ap-mumbai-9",
	"曼谷一区":      "ap-bangkok-1",
	"曼谷二区":      "ap-bangkok-2",
	"曼谷三区":      "ap-bangkok-3",
	"曼谷四区":      "ap-bangkok-4",
	"曼谷五区":      "ap-bangkok-5",
	"曼谷六区":      "ap-bangkok-6",
	"曼谷七区":      "ap-bangkok-7",
	"曼谷八区":      "ap-bangkok-8",
	"曼谷九区":      "ap-bangkok-9",
	"雅加达一区":     "ap-jakarta-1",
	"雅加达二区":     "ap-jakarta-2",
	"雅加达三区":     "ap-jakarta-3",
	"雅加达四区":     "ap-jakarta-4",
	"雅加达五区":     "ap-jakarta-5",
	"雅加达六区":     "ap-jakarta-6",
	"雅加达七区":     "ap-jakarta-7",
	"雅加达八区":     "ap-jakarta-8",
	"圣保罗一区":     "sa-saopaulo-1",
	"圣保罗二区":     "sa-saopaulo-2",
	"圣保罗三区":     "sa-saopaulo-3",
	"圣保罗四区":     "sa-saopaulo-4",
	"圣保罗五区":     "sa-saopaulo-5",
	"圣保罗六区":     "sa-saopaulo-6",
	"圣保罗七区":     "sa-saopaulo-7",
	"圣保罗八区":     "sa-saopaulo-8",
	"多伦多一区":     "na-toronto-1",
	"多伦多二区":     "na-toronto-2",
	"多伦多三区":     "na-toronto-3",
	"多伦多四区":     "na-toronto-4",
	"多伦多五区":     "na-toronto-5",
	"多伦多六区":     "na-toronto-6",
	"多伦多七区":     "na-toronto-7",
	"多伦多八区":     "na-toronto-8",
	"多伦多九区":     "na-toronto-9",
	"硅谷一区":      "na-siliconvalley-1",
	"硅谷二区":      "na-siliconvalley-2",
	"硅谷三区":      "na-siliconvalley-3",
	"硅谷四区":      "na-siliconvalley-4",
	"硅谷五区":      "na-siliconvalley-5",
	"硅谷六区":      "na-siliconvalley-6",
	"硅谷七区":      "na-siliconvalley-7",
	"硅谷八区":      "na-siliconvalley-8",
	"硅谷九区":      "na-siliconvalley-9",
	"弗吉尼亚一区":    "na-ashburn-1",
	"弗吉尼亚二区":    "na-ashburn-2",
	"弗吉尼亚三区":    "na-ashburn-3",
	"弗吉尼亚四区":    "na-ashburn-4",
	"弗吉尼亚五区":    "na-ashburn-5",
	"弗吉尼亚六区":    "na-ashburn-6",
	"弗吉尼亚七区":    "na-ashburn-7",
	"弗吉尼亚八区":    "na-ashburn-8",
	"弗吉尼亚九区":    "na-ashburn-9",
	"法兰克福一区":    "eu-frankfurt-1",
	"法兰克福二区":    "eu-frankfurt-2",
	"法兰克福三区":    "eu-frankfurt-3",
	"法兰克福四区":    "eu-frankfurt-4",
	"法兰克福五区":    "eu-frankfurt-5",
	"法兰克福六区":    "eu-frankfurt-6",
	"法兰克福七区":    "eu-frankfurt-7",
	"法兰克福八区":    "eu-frankfurt-8",
	"法兰克福九区":    "eu-frankfurt-9",
	"莫斯科一区":     "eu-moscow-1",
	"莫斯科二区":     "eu-moscow-2",
	"莫斯科三区":     "eu-moscow-3",
	"莫斯科四区":     "eu-moscow-4",
	"莫斯科五区":     "eu-moscow-5",
	"莫斯科六区":     "eu-moscow-6",
	"莫斯科七区":     "eu-moscow-7",
	"莫斯科八区":     "eu-moscow-8",
	"莫斯科九区":     "eu-moscow-9",
	"上海自动驾驶云一区": "ap-shanghai-adc-1",
	"中国台北一区":    "ap-taipei-1",
	"中国台北二区":    "ap-taipei-2",
	"中国台北三区":    "ap-taipei-3",
	"中国台北四区":    "ap-taipei-4",
	"中国台北五区":    "ap-taipei-5",
	"中国台北六区":    "ap-taipei-6",
}

const (
	BAND_WIDTH = 20
)

const (
	SWITCH_MODE_1 = 1
	SWITCH_MODE_2 = 2
	SWITCH_MODE_4 = 4
)

var SWITCH_MODE = []int{
	SWITCH_MODE_1,
	SWITCH_MODE_2,
	SWITCH_MODE_4,
}

const (
	FW_TYPE_NAT = "nat"
	FW_TYPE_EW  = "ew"
)

var FW_TYPE = []string{
	FW_TYPE_NAT,
	FW_TYPE_EW,
}

const (
	POLICY_ENABLE_TRUE  = "true"
	POLICY_ENABLE_FALSE = "false"
)

var POLICY_ENABLE = []string{
	POLICY_ENABLE_TRUE,
	POLICY_ENABLE_FALSE,
}

const (
	POLICY_SCOPE_SERIAL = "serial"
	POLICY_SCOPE_SIDE   = "side"
	POLICY_SCOPE_ALL    = "all"
)

var POLICY_SCOPE = []string{
	POLICY_SCOPE_SERIAL,
	POLICY_SCOPE_SIDE,
	POLICY_SCOPE_ALL,
}

const (
	POLICY_RULE_ACTION_ACCEPT = "accept"
	POLICY_RULE_ACTION_DROP   = "drop"
	POLICY_RULE_ACTION_LOG    = "log"
)

var POLICY_RULE_ACTION = []string{
	POLICY_RULE_ACTION_ACCEPT,
	POLICY_RULE_ACTION_DROP,
	POLICY_RULE_ACTION_LOG,
}

type SourceContentJson struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type TargetContentJson struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
