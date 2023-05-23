package entity

// 请求对象
type Obj struct {
	// 消息id
	ID uint
	// 负载类型, 0:文本, 1:图片, 2:音频
	Type uint
	// 操作id
	ProcessId uint
	// 请求负载
	Payload string
}

// id
const (
    // 获取好友会话
    GetFriendSession = 0
    // 获取群聊会话
    GetGroupSession = 1
    // 获取新好友信息
    GetNewFriend = 2
    // 获取新群友消息
    GetNewGroup = 3
    // 获取好友列表
    GetFriendList = 4
    // 获取群聊列表
    GetGroupList = 5
    // 添加好友
    AddFriend = 6
    // 添加群聊
    AddGroup = 7
	// 发送好友消息
	SendFriendMsg = 8
	// 发送群聊消息
	SendGroupMsg = 9
	// 通过用户名称模糊查询用户
	GetFuzzyUserByUserName = 10
	// 通过群聊名称模糊查询群聊
	GetFuzzyGroupByGroupName = 11
	// 同意新好友请求
	AgreeNewFriend = 12
	// 同意新群友请求
	AgreeNewGroup = 13
	// 拒绝新好友请求
	RefuseNewFriend = 14
	// 拒绝新群友请求
	RefuseNewGroup = 15
	// 请求五子棋对战
	RequestFiveChess = 16
	// 同意五子棋对战
	AgreeFiveChess = 17
	// 拒绝五子棋对战
	RefuseFiveChess = 18
	// 发送下棋位置
	DownChessPosition = 19
)

// type
const (
	// 文本
	Text = 0
	// 图片
	Pic = 1
	// 音频
	Audio = 2
)