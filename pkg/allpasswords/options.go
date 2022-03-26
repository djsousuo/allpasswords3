package allpasswords

type Options struct {
	FullName         []string // 姓名全拼
	ShortName        []string // 姓名简拼
	BrithDay         []string // 生日
	Phone            []string // 个人手机号
	CompanyFullName  []string // 公司全拼（不要带`有限公司`等关键字，比如：百度、新浪）
	CompanyShortName []string // 公司简拼
	CompanyTelephone []string // 公司电话 手机号或电话
	JobNumber        []string // 员工工号
}
