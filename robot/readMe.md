





https://openwechat.readthedocs.io/zh/latest/user.html



User的结构体

```go
type User struct {
	Uin               int //用户uin
	HideInputBarFlag  int
	StarFriend        int // 是否是星标好友
	Sex               int // 性别 1-男 2-女
	AppAccountFlag    int // 是否为公众号，0为正常联系人，其余都为公众号(个人公众号/服务号:8；企业服务号24；其余存在部分特殊id)
	VerifyFlag        int
	ContactFlag       int // 是否为联系人
	WebWxPluginSwitch int // 网页版微信插件开关
	HeadImgFlag       int
	SnsFlag           int

	//好友、群组
	IsOwner         int
	MemberCount     int // 群人数
	ChatRoomId      int //群id
	UniFriend       int //共同好友
	OwnerUin        int
	Statues         int
	AttrStatus      int
	Province        string //省份
	City            string //城市
	Alias           string
	DisplayName     string //群成员的备注名称
	KeyWord         string
	EncryChatRoomId string
	UserName        string //wechat的用户名字 每次登陆随机分配的
	NickName        string //微信昵称
	HeadImgUrl      string
	RemarkName      string //你备注的名字
	PYInitial       string //昵称拼音首字母
	PYQuanPin       string //昵称拼音全拼
	RemarkPYInitial string
	RemarkPYQuanPin string
	Signature       string

	MemberList []*User // 群成员

}
```

Friend的结构题

```go
type Friend struct{ *User }
```

