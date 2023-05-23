<script setup lang="ts">
import {ref} from 'vue'
import Nav from './components/Nav.vue'
import Set from './components/Set.vue'
  // 请求id
  enum ID {
    // 获取好友会话
    GetFriendSession = 0,
    // 获取群聊会话
    GetGroupSession = 1,
    // 获取新好友信息
    GetNewFriend = 2,
    // 获取新群友消息
    GetNewGroup = 3,
    // 获取好友列表
    GetFriendList = 4,
    // 获取群聊列表
    GetGroupList = 5,
    // 添加好友
    AddFriend = 6,
    // 添加群聊
    AddGroup = 7,
    // 发送好友消息
    SendFriendMsg = 8,
    // 发送群聊消息
    SendGroupMsg = 9,
    // 通过用户名称模糊查询用户
    GetFuzzyUserByUserName = 10,
    // 通过群聊名称模糊查询群聊
    GetFuzzyGroupByGroupName = 11,
    // 同意新好友请求
    AgreeNewFriend = 12,
    // 同意新群友请求
    AgreeNewGroup = 13,
    // 拒绝新好友请求
    RefuseNewFriend = 14,
    // 拒绝新群友请求
    RefuseNewGroup = 15,
    // 请求五子棋对战
    RequestFiveChess = 16,
    // 同意五子棋对战
    AgreeFiveChess = 17,
    // 拒绝五子棋对战
    RefuseFiveChess = 18,
    // 发送下棋位置
    DownChessPosition = 19,
  }
  // 请求负载类型
  enum Type {
    // 文本
    Text = 0,
    // 图片
    Pic = 1,
    // 音频
    Audio = 2,
    // 文件
    File = 3,
  }
  // 请求
  class Request {  
    // 消息id
    ID: ID
    // 负载类型
    Type: Type
    // 接收方id
    ProcessId:number
    // 请求负载
    Payload:string
    constructor(id:ID, type:Type, processId:number, payload:string) {
      this.ID = id
      this.Type = type
      this.ProcessId = processId
      this.Payload = payload
    }
    toJson(): string {
      let s = JSON.stringify(this)
      return s
    }
  }
  // 用户
  class User {
    ID:number
    // 名称
    Name:string
    // 个人介绍
    Introduce:string
    // 构造函数 
    constructor(ID:number,Name:string,Introduce:string) { 
      this.ID = ID
      this.Name = Name
      this.Introduce =Introduce
    }
  }
  // 群聊
  class Group {
    // 群聊id
    ID:number
	  // 群聊名称
	  Name:string
	  // 群主id
	  MasterId:number
	  // 群聊简介
	  Introduce:string
    // 构造函数
    constructor(id:number, name:string, masterId:number, introduce:string) {
      this.ID = id
      this.Name = name
      this.MasterId = masterId
      this.Introduce = introduce
    }
  }
  // 用户聊天消息
  class UserMsg {
    ID:number
    // 发送方id
    SenderId:number
    // 接收方id
    ReceiverId:number
    // 消息类型 文字为0, 图片为1, 音频为2
    Type:number
    // 消息内容
    Content:string
    constructor(id:number, senderId:number, receiverId:number,type:number,content:string) {
      this.ID = id
      this.SenderId = senderId
      this.ReceiverId = receiverId
      this.Type = type
      this.Content = content
    }
  }

  // 群聊聊天消息
  class GroupMsg {
    ID:number
    // 群聊id
    GroupId:number
    // 发送方id
    SenderId:number
    // 发送方名称
    SenderName:string
    // 消息类型 文字为0, 图片为1, 音频为2
    Type:number
    // 消息内容
    Content:string
    constructor(id:number, groupId:number, senderId:number, senderName:string, type:number, content:string) {
      this.ID = id
      this.GroupId = groupId
      this.SenderId = senderId
      this.SenderName = senderName
      this.Type = type
      this.Content = content
    }
  }

  function setPage(page:number) {
    index.value = page
    if (page == 2) {
      otherId = sessionId.value
      initFiveChess()
    }
  }

  // 好友列表
  let friendListValue:User[] = []
  var friendList = ref(friendListValue)
  // 群聊列表
  let groupListValue:Group[] = []
  var groupList = ref(groupListValue)
  // 用户消息列表
  let userMsgListValue:UserMsg[] = []
  var userMsgList = ref(userMsgListValue)
  // 群聊消息列表
  let groupMsgListValue:GroupMsg[] = []
  var groupMsgList = ref(groupMsgListValue)

  // 新会话模块 ================================================================================================================
  var isNewGroup = ref(false)
  // 新好友列表
  let newFriendListValue:User[] = []
  var newFriendList = ref(newFriendListValue)
  // 新群友列表
  class NewGroupMsg {
    GroupId:number
    Username: string
    GroupName: string
    constructor(groupId:number,username:string,groupName:string) {
      this.GroupId = groupId
      this.Username= username
      this.GroupName= groupName
    }
  }
  let newGroupListValue:NewGroupMsg[] = []
  var newGroupList = ref(newGroupListValue)
  // 搜索用户列表
  let searchUserListValue:User[] = []
  var searchUserList = ref(searchUserListValue)
  // 搜索群聊列表
  let searchGroupListValue:Group[] = []
  var searchGroupList = ref(searchGroupListValue)
  var index = ref(3)
  var sessionId = ref(0)
  var sessionName = ref('')
  // 0是用户会话, 1是群聊会话
  var sessionType = ref(0)
  var searchInput:string

  var msgInput:string
  var emoji = ref([
    "https://s1.chu0.com/src/img/png/ff/ffdc9e6ebc3b4f89aaf08eb2e58ba96e.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:xpIhE9lnwo2iM1gaUGl6JHF65CQ=",
    "https://s1.chu0.com/src/img/png/ea/ea3af3b2dd8742d8b69fcc0aceaf2bae.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:cBQMLFE6wz44UUXj8GyDIR2vxUs=",
    "https://s1.chu0.com/src/img/png/b0/b0884a8b4081463cae8e4d0a1b9422fd.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:ul6Ctv-yzpqMF5oM9_8fR5nJY4k=",
    "https://s1.chu0.com/src/img/png/2f/2f46f89f49a9447ca4d0b3207f57c1bb.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:Dupe_Fzdlxv0dvXPbKHRadT9NGQ=",
    "https://s1.chu0.com/src/img/png/a8/a81d642487dd4a5ba723a647d9e02eb7.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:QYALC0L6SxPqvr2E-c7ywN2ZVbU=",
    "https://s1.chu0.com/src/img/png/a8/a8a0e3152bb44100b051b589cc545d19.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:x2w_Se-UR3V-ez6blMqaqTLU2YM=",
    "https://s1.chu0.com/src/img/png/60/608f0e9f66d7455388af07827bf78bbd.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:NwDOow8bpIIKqIjb3wnorjqFHuM=",
    "https://s1.chu0.com/src/img/png/fc/fc422659017444779c764e7971229553.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:HVWaqokQHw_i_S_-OMdtfYWtKok=",
    "https://s1.chu0.com/src/img/png/f9/f961cbb4aac84c5990602913d1d0e582.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:upAO9En9PJoaOokhkqPrI8dKw_4=",
    "https://s1.chu0.com/src/img/png/77/77d128dc772541f48971c06199705d78.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:hzLF47b2vWtqGxu2JxCY8VFvt7E=",
    "https://s1.chu0.com/src/img/png/27/2751030cad7a492cad64cfac5fdb7730.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:gDlq_nlj9KkU3g5jKlhmAu7qlOM=",
    "https://s1.chu0.com/src/img/png/63/636a6c50899c4a4483a717a0865b6e31.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:hSlrZgY2nn3h4YomFNWioPX0zlM=",
    "https://s1.chu0.com/src/img/png/b8/b8e66c0832664a838a1f2348c4ef733a.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:KbUPFoWoblgXJO9WTSciNLImTHA=",
    "https://s1.chu0.com/src/img/png/e6/e6594a15e58b441686ab71b9ec3b66d4.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:DOiRq5QBi0hwRNVW2Ho-F9dRiUE=",
    "https://s1.chu0.com/src/img/png/91/91bd8d89e3a1480f810b71b294546c4f.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:-yQodZxx1sUp7WdSDPYZToGyXZM=",
    "https://s1.chu0.com/src/img/png/7a/7a29127664f143e5a8f46d37e7e3e74b.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:qLI-LYfUNxJ2DbYzwMJpOUT3nvo=",
    "https://s1.chu0.com/src/img/png/90/904844bb0bbe4f85a1e54a5073122692.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:xHOyUL57RpNDVOjeZt5V2AJ4iPE=",
    "https://s1.chu0.com/src/img/png/76/76f6157cdd7f4d6fb9ef6611de3fa5ba.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:hWGLbteCFgL9HDTyFQ-AAHR37Wo=",
    "https://s1.chu0.com/src/img/png/b9/b922bd54dcc24faaaacfa1966d759721.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:wByck7o1dJsEWT_xeeN1-l34I2M=",
    "https://s1.chu0.com/src/img/png/ba/baac6a6d674745f58e5416f0bb6c0ca3.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:4xUzEtGeklfyZHTUmxQAqUNDn8M=",
    "https://s1.chu0.com/src/img/png/16/162933c730f44b79ad94c43d4c83f0bf.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:JKH1b-F6ntogiEfVljFDU90I0gc=",
    "https://s1.chu0.com/src/img/png/b7/b70d5b2e4bcb41b1ab11e4f392e3403b.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:D-0BzJVQtrm1x30DguqqGvCtqc8=",
    "https://s1.chu0.com/src/img/png/cc/cc078db137194e479039d5f5568842a0.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:z2EMo9ODTBP8c92yR5J8VV-y66s=",
    "https://s1.chu0.com/src/img/png/a1/a15919447b0b449b8a623776ec2018da.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:ALaO8xh1sLiTMZd2W-L2ZM5YIX4=",
    "https://s1.chu0.com/src/img/png/7f/7f3f03594868462b846d905a8ec55d0c.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:a3w4ieklhmvSsSjd6nF-wXxPd6Y=",
    "https://s1.chu0.com/src/img/png/b9/b90c8d41861a4a598f4fa441c65ed566.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:lKQ1qW8vo3CNEMyIloxXw6n2dcE=",
    "https://s1.chu0.com/src/img/png/6a/6a7adb35ed7a481a83b76f71c90603aa.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:jR013xd9zAXtZqsLI_CtL42bOBU=",
    "https://s1.chu0.com/src/img/png/a8/a8e657b51d474bdcb6f687686bc75e27.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:m59SXe2CLOTpKGADtdPC4Vz7eC4=",
    "https://s1.chu0.com/src/img/png/e3/e38b1219a6324072926f8219e59c5346.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:aPNiroYvu0wGE4vBpGUik8QfP5g=",
    "https://s1.chu0.com/src/img/png/f2/f2c880e8bd6e4fb2b3569cc8b2858e6c.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:5jABavKa6hatB-wVQCbkygva6Bc=",
    "https://s1.chu0.com/src/img/png/ae/ae68da59ccb74a1a87bf70543b281454.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:G6kwZh_r6xSe3aInK3YbfiVHLn0=",
    "https://s1.chu0.com/src/img/png/14/1408057aa0af42d79c29e49f5366479f.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:pt5ng4jjGwxQVfUdBD_LOlOISYY=",
    "https://s1.chu0.com/src/img/png/d7/d79b5f5c5e624651b845b0dbdf460047.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:0VTloBjF8otmeOVDmBCJ0ze5nyg=",
    "https://s1.chu0.com/src/img/png/46/46f01707dfc843b59352547d7f150025.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:N8HPWUfbnMw5c8R4oGVDLrck43o=",
    "https://s1.chu0.com/src/img/png/fa/fa0112bfcf1046e3865b5529c20a238e.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:30E40chqwPQJ7WK5mtKJ9ACswNU=",
    "https://s1.chu0.com/src/img/png/39/39c48a5ef0bf4d3788643559b496c80f.png?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:yVEIKCF-ZbQhTcqLeZWtPRvh2KU=",
    "https://s1.chu0.com/src/img/gif/fe/fed1040f232a4f1bb8be016c41656e25.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:FpyE6RfkSLXnmheOD2OPoc8p-zU=",
    "https://s1.chu0.com/src/img/gif/92/9251c9ccd39b4e20835293300fc90818.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:tRI1oNUkRB7F5z9I0uH9X00jb1c=",
    "https://s1.chu0.com/src/img/gif/0e/0e4d4ffa74694781a906e1c0adecb733.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:eqwzP29VUY_cWNvYk4Gx0NE42es=",
    "https://s1.chu0.com/src/img/gif/fa/faee6b522e68448781b782aaa19cde8e.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:O6NMpWN0BhNk_nfHe94k1CyZ73k=",
    "https://s1.aigei.com/src/img/gif/90/902e1490ef0543328b6e69074d3524ca.gif?e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:gkHaNWlPd7-ZIXA6IfIGLWYe2sc=",
    "https://s1.chu0.com/src/img/gif/74/748efdddcb8d4d64a9706cefc84418a2.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:yz644GNJfnUeKU-6j7l80oViYDk=",
    "https://s1.chu0.com/src/img/gif/6b/6b55917a7d214261bd97cf6f2757f132.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:uh8e_msiAGsRGk4OWP_oeF4Ye3o=",
    "https://s1.aigei.com/src/img/gif/2b/2b572958c86f4ef9b4332acdd229881b.gif?e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:h5gv7N--oMwAcFKy3O188WcIlH4=",
    "https://s1.chu0.com/src/img/gif/78/7818b40a97b4430fbe496d8ccfbfcc93.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:Co3PsArhHr_pKBHWoJ0pMmcF0rE=",
    "https://s1.chu0.com/src/img/gif/66/6687a055cf93429ab1fac6e642f47abb.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:pGnBWYI8wuJCV7__xXe0yk0vGl0=",
    "https://s1.chu0.com/src/img/gif/8a/8a6a2f0a7a114e26bcef69ace4118d21.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:8PbJxnI_qrsYNs2EwUfs7oe_Adw=",
    "https://s1.chu0.com/src/img/gif/4d/4d4cb52063e84c84bbc9be9217ad80fa.gif?e=1735488000&token=1srnZGLKZ0Aqlz6dk7yF4SkiYf4eP-YrEOdM1sob:4ygFD9LQ90GiPssISYnxJgTQAXI=",
  ])
  var emojiDisplay = ref(false)
  var audioText = ref('')
  var audioData:BlobPart[] = []
  var audioUrl = ref('')
  var mediaRecorder:MediaRecorder
  const myid = parseInt(localStorage.getItem('id') as string)
  const myname = localStorage.getItem('name') as string
  // 群聊模块 ================================================================================================================
  var isCreateGroup = ref(false)


  // websocket模块 ================================================================================================================
  const ip = window.location.hostname
  const ws:WebSocket=new WebSocket('ws://' + ip + ':' + window.location.port + '/upgrade')
  ws.onopen = function (evt) {
      // 获取好友列表
      send(ID.GetFriendList, Type.Text, 0, '')
      // 获取群聊列表
      send(ID.GetGroupList, Type.Text, 0, '')
      // 获取新好友列表
      send(ID.GetNewFriend, Type.Text, 0, '')
      // 获取新群友列表
      send(ID.GetNewGroup, Type.Text, 0, '')
  }
  ws.onclose = function (evt) {
    alert('连接断开')
  }
  ws.onmessage = function (evt) {
      let req:Request = JSON.parse(evt.data)
      console.log(req)
      switch (req.ID) {
        case ID.GetFriendSession:
          // 获取好友会话
          userMsgList.value = JSON.parse(req.Payload)
          console.log(userMsgList)
          break
        case ID.GetGroupSession:
          // 获取群聊会话
          groupMsgList.value = JSON.parse(req.Payload)
          console.log(groupMsgList)
          break
        case ID.GetNewFriend:
          // 获取新好友列表
          newFriendList.value = JSON.parse(req.Payload)
          console.log('new friend list: ', JSON.parse(req.Payload))
          break
        case ID.GetNewGroup:
          // 获取新群友列表
          newGroupList.value = JSON.parse(req.Payload)
          break
        case ID.GetFriendList:
          // 更新好友列表
          friendList.value = JSON.parse(req.Payload)
          break
        case ID.GetGroupList:
          // 更新群聊列表
          groupList.value = JSON.parse(req.Payload)
          break
        case ID.AddFriend:
          // 添加好友
          newFriendList.value.push(JSON.parse(req.Payload) as User)
          break
        case ID.AddGroup:
          // 添加群聊
          newGroupList.value.push(JSON.parse(req.Payload) as NewGroupMsg)
          break
        case ID.SendFriendMsg:
          // 发送好友消息
          if (sessionType.value == 0 && sessionId.value == req.ProcessId) {
            // 是发给当前好友会话的消息
            userMsgList.value.push(JSON.parse(req.Payload) as UserMsg)
          }
          break
        case ID.SendGroupMsg:
          // 发送群聊消息
          console.log(sessionType.value,sessionId.value, req.ProcessId,req.Payload)
          if (sessionType.value == 1 && sessionId.value == req.ProcessId) {
            // 是发给当前群聊会话的消息
            groupMsgList.value.push(JSON.parse(req.Payload) as GroupMsg)
          }
          break
        case ID.GetFuzzyUserByUserName:
          // 模糊查询的用户列表
          searchUserList.value = JSON.parse(req.Payload)
          break
        case ID.GetFuzzyGroupByGroupName:
          // 模糊查询的群聊列表
          searchGroupList.value = JSON.parse(req.Payload)
          console.log(req.Payload)
          break
        case ID.AgreeNewFriend:
          // 同意新好友请求
          friendList.value.push(JSON.parse(req.Payload) as User)
          break
        case ID.AgreeNewGroup:  
          // 同意新群友请求
          groupList.value.push(JSON.parse(req.Payload) as Group)
          break
        case ID.RequestFiveChess:
          // 请求五子棋对战
          requestChessList.value.push(new User(req.ProcessId, req.Payload, ''))
          break
        case ID.AgreeFiveChess:
          // 同意五子棋对战
          
          // 1.设置棋子颜色
          console.log('chess color is ' + req.Payload)
          chessColor.value = Number(req.Payload)
          console.log('chess color is ', chessColor)
          // 2.设置对方id
          otherId = req.ProcessId
          // 3.设置黑棋先手
          curColor.value = 2
          break
        case ID.RefuseFiveChess:
          // 拒绝五子棋对战
          break
        case ID.DownChessPosition:
          // 发送下棋位置
          if (req.ProcessId == otherId && curColor.value != chessColor.value) {
            // 获取到对方下棋的信息
            let xy = req.Payload.split("\"")[1].split(",")
            let x = parseInt(xy[0])
            let y = parseInt(xy[1])
            console.log(xy)
            downChess(x, y, curColor.value)
            
          }
          break
      }
  }
  ws.onerror = function (evt) {
      // alert("ERROR")
  }
  function send(id:ID, type:Type, processId:number, payload:string) {
    let data = new Request(id, type, processId, payload).toJson()
    console.log(data)
    ws.send(data)
  }
  // ============================================================================================================================
  // 获取用户会话
  function friendSessionFn(friend:User) {
    send(ID.GetFriendSession, Type.Text, friend.ID, '')
    sessionId.value = friend.ID
    sessionName.value = friend.Name
    sessionType.value = 0
  }

  // 获取群聊会话
  function groupSessionFn(group:Group) {
    send(ID.GetGroupSession, Type.Text, group.ID, '')
    sessionId.value = group.ID
    sessionName.value = group.Name
    sessionType.value = 1
  }

  function searchFn(name:string) {
    send(ID.GetFuzzyUserByUserName, Type.Text, 0, name)
  }

  function searchGroupFn(name:string) {
    send(ID.GetFuzzyGroupByGroupName, Type.Text, 0, name)
  }

  function addUserFn(id:number) {
    send(ID.AddFriend, Type.Text, id, '')
  }

  function addGroupFn(id:number) {
    send(ID.AddGroup, Type.Text, id, '')
  }

  function agreeNewFriend(id:number, i:number) {
    send(ID.AgreeNewFriend, Type.Text, id, '')
    friendList.value.push(newFriendList.value[i])
    let tmp:User[] = [] 
    for (var j = 0; j < newFriendList.value.length;j++) {
      if (i == j) {
        continue
      }
      tmp.push(friendList.value[j])
    }
    newFriendList.value = tmp
  }

  function refuseNewFriend(id:number, i:number) {
    send(ID.RefuseNewFriend, Type.Text, id, '')
    let tmp:User[] = [] 
    for (var j = 0; j < newFriendList.value.length;j++) {
      if (i == j) {
        continue
      }
      tmp.push(friendList.value[j])
    }
    newFriendList.value = tmp
  }

  function agreeNewGroup(id:number, username:string, i:number) {
    send(ID.AgreeNewGroup, Type.Text, id, username)
    let tmp:NewGroupMsg[] = []
    console.log(i)
    for (var j = 0; j < newGroupList.value.length;j++) {
      console.log(j)
      if (i == j) {
        continue
      }
      tmp.push(newGroupList.value[j])
    }
    newGroupList.value = tmp
  }

  function refuseNewGroup(id:number, username:string, i:number) {
    send(ID.RefuseNewGroup, Type.Text, id, username)
    delete newGroupList.value[i]
  }

  function audioFn() {
    switch (audioText.value) {
      case '开始录音':{
        audioUrl.value = ''
        audioData = []
        navigator.mediaDevices.getUserMedia({audio: true}).then(stream=>{
          mediaRecorder = new MediaRecorder(stream)
          mediaRecorder.start()
          mediaRecorder.ondataavailable = evt => {
            audioData.push(evt.data)
          }
          mediaRecorder.onstop = evt=>{
            let tracks = mediaRecorder.stream.getTracks()
            for(var i = 0 ; i< tracks.length ; i++) {
                tracks[i].stop()
            }
            const blob = new Blob(audioData, { type: "audio/mp3; codecs=opus" })
            audioUrl.value = window.URL.createObjectURL(blob)    
          }
        })
        audioText.value = '结束录音'
        break
      }
      case '结束录音':{
        mediaRecorder.stop()
        audioText.value = '开始录音'
        break
      }
    }
  }

  function blobToBase64(blob:Blob, callback:(s:string)=>void) {
      var reader = new FileReader()
      reader.onload = e=> {
          callback(e.target?.result?.toString() as string)
      }
      reader.readAsDataURL(blob);
  }
  function base642Url(data:string) {
      var parts = data.split(';base64,'),
          contentType = parts[0].split(':')[1],
          raw = window.atob(parts[1]),
          length = raw.length,
          arr = new Uint8Array(length)
      for (var i = 0; i < length; i++) {
          arr[i] = raw.charCodeAt(i)
      }
      var blob = new Blob([arr], { type: contentType })
      return URL.createObjectURL(blob)
  }
  function sendAudio() {
    const blob = new Blob(audioData, {type: "audio/mp3; codecs=opus"})
    blobToBase64(blob, s=>{
      userMsgList.value.push(new UserMsg(0, myid, sessionId.value, Type.Audio, s))
      send(sessionType.value==0?ID.SendFriendMsg:ID.SendGroupMsg, Type.Audio, sessionId.value, s)
    })
    audioUrl.value = ''
    audioText.value = ''
  }

  // ============================================================================================================================
  //五子棋模块
  var chessColor = ref(0)
  var ctx:CanvasRenderingContext2D
  var curColor = ref(0)
  var otherId:number
  let requestChessListValue:User[] = []
  var requestChessList = ref(requestChessListValue)
  var isOver = ref(false)
  var chessBoard:number[][]

  function initFiveChess() {
    chessColor.value = 0
    curColor.value = 0
    otherId = 0
    requestChessListValue = []
    requestChessList.value = requestChessListValue
    isOver.value = false
    chessBoard = [
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],

      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],

      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
      [0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0],
    ]
    setTimeout(function(){
      showChessBoard()
    }, 50)
  }

  // 显示棋盘
  function showChessBoard(){
    let div = document.getElementsByClassName('chess-board')[0] as HTMLDivElement
    // let res:FiveChessRes = new FiveChessRes(myid, otherId, 0, 0, true)
    // let s = JSON.stringify(res)
    // send(ID.FiveChess, Type.Text, otherId, s)
    div.innerHTML = '<canvas width="450px" height="450px"></canvas>'
    let chess = div.getElementsByTagName('canvas')[0]
    chess.style.backgroundColor = 'antiquewhite'
    chess.style.zIndex = '1'
    ctx = chess.getContext('2d') as CanvasRenderingContext2D
    drawChessBoard()

    chess.onclick = chessBoardClickFn
  }

  // 画棋盘
  function drawChessBoard() {
    ctx.strokeStyle = "#b9b9b9"
    for (let i = 0; i < 15; i++) {
      //设置水平方向的起始点
      ctx.moveTo(15, 15+i*30)
      //设置水平方向的终止点
      ctx.lineTo(435, 15+i*30)
      //连接水平方向的两点
      ctx.stroke()
      //设置竖直方向的起始点
      ctx.moveTo(15+i*30, 15)
      //设置竖直方向的终止点
      ctx.lineTo(15+i*30, 435)
      //连接竖直方向的两点
      ctx.stroke()
    }
  }

  // 棋盘点击事件
  function chessBoardClickFn(evt:MouseEvent) {

    let x = Math.floor(evt.offsetX/30)
    let y = Math.floor(evt.offsetY/30)
    if (chessBoard[x][y] != 0 || curColor.value != chessColor.value) {
        return
    }

    // 刷新本地棋盘
    downChess(x, y, chessColor.value)

    // 坐标信息
    let pos = x + "," +y

    // 将落子信息发送给对手
    send(ID.DownChessPosition, Type.Text, otherId, pos)
  }

  // 下棋
  function downChess(x:number, y:number, color:number) {
    if (isOver.value) {
      alert("对局已结束")
      return
    }

    console.log("下棋的坐标:(" + x + ", " + y + "), 棋子的颜色: " + color + "\n")

    chessBoard[x][y] = color

    // 画棋
    ctx.beginPath()
    ctx.arc(15+x*30, 15+y*30, 13, 0, 2*Math.PI)
    ctx.closePath()
    ctx.fillStyle = color == 1 ? "white": "black"
    ctx.fill()

    // 检查游戏是否结束
    if (checkOver(x, y, color)) {
      isOver.value = true
      if (color == 1) {
          alert("白棋获胜")
      } else {
          alert("黑棋获胜")
      }
    }

    // 切换下棋回合
    if (curColor.value == 1) {
      curColor.value = 2
    } else {
      curColor.value = 1
    }
  }

  // 检查游戏是否结束
  function checkOver(x:number, y:number, color:number) {
    let count = 0
    let i = 1
    while (i<5&&y>=i&&chessBoard[x][y-i]==color) {
        count++
        i++
    }
    i = 1
    while (i<5&&y+i<15&&chessBoard[x][y+i]==color) {
        count++
        i++
    }
    if (count > 3) {
        return true
    }
    // |
    count = 0
    i = 1
    while (i<5&&x>=i&&chessBoard[x-i][y]==color) {
        count++
        i++
    }
    i = 1
    while (i<5&&x+i<15&&chessBoard[x+i][y]==color) {
        count++
        i++
    }
    if (count > 3) {
        return true
    }
    // \
    count = 0
    i = 1
    while (i<5&&x>=i&&y>=i&&chessBoard[x-i][y-i]==color) {
        count++
        i++
    }
    i = 1
    while (i<5&&x+i<15&&y+i<15&&chessBoard[x+i][y+i]==color) {
        count++
        i++
    }
    if (count > 3) {
        return true
    }
    // /
    count = 0
    i = 1
    while (i<5&&x+i<15&&y>=i&&chessBoard[x+i][y-i]==color) {
        count++
        i++
    }
    i = 1
    while (i<5&&x>=i&&y+i<15&&chessBoard[x-i][y+i]==color) {
        count++
        i++
    }
    return count > 3
  }

  // ============================================================================================================================
  // 群聊模块
  var groupName = ref('')
  var introduction = ref('')
  function createGroup() {
    if (groupName.value.includes(' ') || introduction.value.includes(' ')) {
      alert('群聊名称或群聊简介不能有空格')
    } else {
      let xhr = new XMLHttpRequest()
      xhr.open('post', 'http://' + window.location.hostname + ':' + window.location.port + '/group/regist', false)
      let data = new FormData()
      data.append('name', groupName.value)
      data.append('introduce', introduction.value)
      data.append('head', (document.getElementById('head') as HTMLInputElement).files?.item(0) as File)
      xhr.send(data)
      if (xhr.status == 200) {
        // 创建成功
        groupList.value.push(new Group(parseInt(xhr.responseText), groupName.value, myid, introduction.value,))
        let text = xhr.responseText
        alert(text)
      } else {
        // 创建失败
        let text = xhr.responseText
        alert(text)
      }
    }
  }
  function chooseHead() {
    document.getElementById('head')?.click()
  }
</script>

<template>
  <div class="home">
    <div class="main">
      <Nav class="nav" @page="setPage"></Nav>
      <!-- <div class="nav">
        <img class="head-icon" :src="'http://'+ip+':9999/head/' + myname + '.png'">
        <img class="head-icon" :src="img.url" @click="img.fn()" v-for="img in nav">
      </div> -->
      <div class="content">
        <div class="list">

          <div class="list-left">
            {{ myname }}

            <h3>好友</h3>
            <div class="session-list">
              <div class="msg" v-for="friend in friendList" @click="friendSessionFn(friend)">
                <img class="msg-img" :src="'http://'+ip+':9999/head/'+friend.Name+'.png'">
                <p class="msg-title">
                  {{ friend.Name }}
                </p>
              </div>
            </div>
            <h3>群聊</h3>
            <div class="session-list">
              <div class="msg" v-for="group in groupList" @click="groupSessionFn(group)">
                <img class="msg-img" :src="'http://'+ip+':9999/group/head/'+group.Name+'.png'">
                <p class="msg-title">
                  {{ group.Name }}
                </p>
              </div>
            </div>
          </div>

          <div class="list-right" v-if="index==1">
            <div class="chat-title">
              <h2>{{ sessionName }}</h2>
            </div>
            <div class="chats">
              <!-- 用户消息 -->
              <div :class="msg.SenderId==myid?'chat-right':'chat-left'" v-if="sessionType==0" v-for="msg in userMsgList">
                <!-- 发送方头像 -->
                <img :src="'http://'+ip+':9999/head/' + (msg.SenderId==myid?myname:sessionName) +'.png'" class="list-img">
                <!-- 文本消息 -->
                <p v-if="msg.Type==Type.Text" :class="msg.SenderId==myid?'right-msg':'left-msg'">
                  {{ msg.Content }}
                </p>
                <!-- 图片消息 -->
                <img v-else-if="msg.Type==Type.Pic" :src="msg.Content" :class="msg.SenderId==myid?'right-msg':'left-msg'">
                <!-- 音频消息 -->
                <audio v-else-if="msg.Type==Type.Audio" :src="base642Url(msg.Content)" :class="msg.SenderId==myid?'right-audio':'left-audio'" controls></audio>

              </div>

              <!-- 群聊消息 -->
              <div :class="msg.SenderId==myid?'chat-right':'chat-left'" v-if="sessionType==1" v-for="msg in groupMsgList">
                <!-- 发送方头像 -->
                <img :src="'http://'+ip+':9999/head/' + msg.SenderName +'.png'" class="list-img">

                <div>
                <p :style="'font-size: 12px; line-height: 0px; color: gray; text-align: ' + (msg.SenderId==myid ? 'right;':'left;')">{{ msg.SenderName }}</p>
                <!-- 文本消息 -->
                <p v-if="msg.Type==Type.Text" :class="msg.SenderId==myid?'right-msg':'left-msg'">
                  {{ msg.Content }}
                </p>
                <!-- 图片消息 -->
                <img v-else-if="msg.Type==Type.Pic" :src="msg.Content" :class="msg.SenderId==myid?'right-msg':'left-msg'">
                <!-- 音频消息 -->
                <audio v-else-if="msg.Type==Type.Audio" :src="base642Url(msg.Content)" :class="msg.SenderId==myid?'right-audio':'left-audio'" controls></audio>
                </div>              
              </div>
              <!-- 表情 -->
              <ul v-if="emojiDisplay">
                <li v-for="emo in emoji">
                  <img :src="emo" @click="function(){
                    send(sessionType==0?ID.SendFriendMsg:ID.SendGroupMsg, Type.Pic, sessionId, emo)
                    userMsgList.push(new UserMsg(0, myid, sessionId, Type.Pic, emo))
                    emojiDisplay = false
                  }">
                </li>
              </ul>
              <!-- 语音 -->
              <div v-if="audioText!=''" class="audioDiv">
                <button class="audio-btn" @click="audioFn()">
                  {{ audioText }}
                </button>
                <div v-if="audioUrl!=''">
                  <audio controls :src="audioUrl"></audio>
                  <button class="audio-btn" @click="sendAudio()">发送</button>
                </div>
              </div>
            </div>
            <div class="chat-boom">
              <input type="text" v-model="msgInput">
              <!-- 表情 -->
              <img class="input-btn" @click="emojiDisplay=!emojiDisplay" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0yNCA0NEMzNS4wNDU3IDQ0IDQ0IDM1LjA0NTcgNDQgMjRDNDQgMTIuOTU0MyAzNS4wNDU3IDQgMjQgNEMxMi45NTQzIDQgNCAxMi45NTQzIDQgMjRDNCAzNS4wNDU3IDEyLjk1NDMgNDQgMjQgNDRaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0zMSAxOFYxOSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0xNyAxOFYxOSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0zMSAzMUMzMSAzMSAyOSAzNSAyNCAzNUMxOSAzNSAxNyAzMSAxNyAzMSIgc3Ryb2tlPSIjMzMzIiBzdHJva2Utd2lkdGg9IjQiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjwvc3ZnPg==">
              <!-- 语音 -->
              <img class="input-btn" @click="audioText=audioText==''?'开始录音':''" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik0yNCA0NEMzNS4wNDU3IDQ0IDQ0IDM1LjA0NTcgNDQgMjRDNDQgMTIuOTU0MyAzNS4wNDU3IDQgMjQgNEMxMi45NTQzIDQgNCAxMi45NTQzIDQgMjRDNCAzNS4wNDU3IDEyLjk1NDMgNDQgMjQgNDRaIiBmaWxsPSJub25lIiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCIvPjxwYXRoIGQ9Ik0xNyAyNS44OTk0QzE4LjEwNDYgMjUuODk5NCAxOSAyNS4wMDQgMTkgMjMuODk5NEMxOSAyMi43OTQ4IDE4LjEwNDYgMjEuODk5NCAxNyAyMS44OTk0QzE1Ljg5NTQgMjEuODk5NCAxNSAyMi43OTQ4IDE1IDIzLjg5OTRDMTUgMjUuMDA0IDE1Ljg5NTQgMjUuODk5NCAxNyAyNS44OTk0WiIgZmlsbD0iIzMzMyIvPjxwYXRoIGQ9Ik0yMS45NDk3IDI4Ljg0OTJDMjMuMjE2NSAyNy41ODI1IDI0IDI1LjgzMjUgMjQgMjMuODk5NUMyNCAyMS45NjY1IDIzLjIxNjUgMjAuMjE2NSAyMS45NDk3IDE4Ljk0OTciIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48cGF0aCBkPSJNMjYuODk5NCAzMy43OTlDMjkuNDMyOSAzMS4yNjU1IDMwLjk5OTkgMjcuNzY1NSAzMC45OTk5IDIzLjg5OTVDMzAuOTk5OSAyMC4wMzM1IDI5LjQzMjkgMTYuNTMzNSAyNi44OTk0IDE0IiBzdHJva2U9IiMzMzMiIHN0cm9rZS13aWR0aD0iNCIgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PC9zdmc+">
              <!-- 发送 -->
              <img class="input-btn" @click="function(){
                if (sessionType==0) {
                  send(ID.SendFriendMsg, Type.Text, sessionId, msgInput)
                  userMsgList.push(new UserMsg(0, myid, sessionId, Type.Text, msgInput))
                } else {
                  send(ID.SendGroupMsg, Type.Text, sessionId, msgInput)
                  // 不需要push, 因为服务端会推送消息
                }
                msgInput = ''
                }" src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48c3ZnIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxwYXRoIGQ9Ik00MiA2TDQgMjAuMTM4M0wyNCAyNC4wMDgzTDI5LjAwNTIgNDRMNDIgNloiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWpvaW49InJvdW5kIi8+PHBhdGggZD0iTTI0LjAwODMgMjQuMDA4NEwyOS42NjUxIDE4LjM1MTYiIHN0cm9rZT0iIzMzMyIgc3Ryb2tlLXdpZHRoPSI0IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz48L3N2Zz4=">
            </div>
          </div>

          <div class="list-right" v-else-if="index==2">
            <!-- 五子棋 -->
            <span class="chess-board"></span>
            <span v-if="isOver">
              对局已结束
            </span>
            <span v-else-if="chessColor == 1 && chessColor == curColor">
              我方执白棋, 我方回合
            </span>
            <span v-else-if="chessColor == 1 && chessColor != curColor">
              我方执白棋, 对方回合
            </span>
            <span v-else-if="chessColor == 2 && chessColor == curColor">
              我方执黑棋, 我方回合
            </span>
            <span v-else-if="chessColor == 2 && chessColor != curColor">
              我方执黑棋, 对方回合
            </span>
            <span v-else>
              对局未开始
            </span>
            <br>
            <input type="text" v-model="searchInput">
            <button @click="searchFn(searchInput)">搜索</button>
            <!-- 搜索列表 -->
            <div class="chess-user-list">
              <div v-for="user in searchUserList">
                <button @click="send(ID.RequestFiveChess, Type.Text, user.ID, '')">请求五子棋对战</button>
                <img class="msg-img" :src="'http://'+ip+':9999/head/'+user.Name+'.png'">
                {{ user.Name }}
              </div>
            </div>
            <!-- 请求对战的列表 -->
            <div class="chess-user-list">
              <div v-for="(user, i) in requestChessList">
                <button @click="send(ID.AgreeFiveChess, Type.Text, user.ID, '')">同意</button>
                <button @click="send(ID.RefuseFiveChess, Type.Text, user.ID, '')">拒绝</button>
                <img class="msg-img" :src="'http://'+ip+':9999/head/'+user.Name+'.png'">
                {{ user.Name }}
              </div>
            </div>
          </div>

          <div class="list-right" v-else-if="index==3">
            <input type="text" v-model="searchInput">
            <button @click="isNewGroup?searchGroupFn(searchInput):searchFn(searchInput)">搜索</button>
            <span v-if="isNewGroup">
              <button @click="isNewGroup=!isNewGroup">切换用户模式</button>
              <button @click="isCreateGroup=!isCreateGroup">创建群聊</button>
              <div style="z-index: 2; position: absolute; " v-if="isCreateGroup">
                <br>
                <input type="text" v-model="groupName" placeholder="群聊名称">
                <br>
                <input type="text" v-model="introduction" placeholder="群聊简介">
                <input style="display: none;" id="head" accept="image/*" type="file" >
                <br>
                <button @click="chooseHead()">选择头像</button>
                <button @click="createGroup()">确认创建</button>
              </div>
              <div class="new-list" v-else>
                <div class="search-list">
                  群聊搜索列表
                  <div v-for="group in searchGroupList">
                    <button @click="addGroupFn(group.ID)">添加</button>
                    <img class="msg-img" :src="'http://'+ip+':9999/group/head/'+group.Name+'.png'">
                    {{ group.Name }}
                  </div>
                </div>
                <div class="search-list">
                  入群请求列表
                  <div v-for="(group, i) in newGroupList">
                    <button @click="agreeNewGroup(group.GroupId, group.Username, i)">同意</button>
                    <button @click="refuseNewGroup(group.GroupId, group.Username, i)">拒绝</button>
                    {{ group.Username }}
                    <img class="msg-img" :src="'http://'+ip+':9999/head/'+group.Username+'.png'">
                    {{ group.GroupName }}
                  </div>
                </div>
              </div>
            </span>

            <span v-else>
              <button @click="isNewGroup=!isNewGroup">切换群聊模式</button>
              <div class="new-list">
                <div class="search-list">
                  用户搜索列表
                  <div v-for="user in searchUserList">
                    <button @click="addUserFn(user.ID)">添加</button>
                    <img class="msg-img" :src="'http://'+ip+':9999/head/'+user.Name+'.png'">
                    {{ user.Name }}
                  </div>
                </div>
                <div class="search-list">
                  好友请求列表
                  <div v-for="(user, i) in newFriendList">
                    <button @click="agreeNewFriend(user.ID, i)">同意</button>
                    <button @click="refuseNewFriend(user.ID, i)">拒绝</button>
                    <img class="msg-img" :src="'http://'+ip+':9999/head/'+user.Name+'.png'">
                    {{ user.Name }}
                  </div>
                </div>
              </div>
            </span>
          </div>
          <Set class="list-right" v-else-if="index==4"></Set>
        </div>

      </div>
    </div>
  </div>
</template>