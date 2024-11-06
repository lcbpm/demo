### 首页数据获取

1. InitParams(构造ContextParams)
  - getRecBatchId 时间戳
  - GetCfgPoolRules 推荐-推荐池规则
  - GetGenderRules 推荐性别规则
  - GetBlackIdMap 获取拉黑用户
  - initDlParams deeplink 通过跳转连接获取信息 标签ID  角色ID  后台配置的DeeplinkID
角色信息存redis,首页存,非首页取
  - GetDefaultRecGroupCharidByBucket 根据桶获取默认推荐的角色
  - selectStrategy 性别加权策略
获取所有的聊天信息(src/mecord/talker/service/character_rec_new_pool_gender.go:202这里的逻辑等等看)
  - GetNewUserType 是否新老用户
  - resetData 重置数据(删除缓存,细看这里是干嘛的)
  - CharacterRecUserVisitLog 记录日志
  - StatUserChatInfo 数据落库

2. 根据构造ContextParams,新老用户处理角色列表
  - 聊天角色
  - 未聊角色
  - 聊天CD角色
  - 未聊天CD角色

3. PushToDqUpdateChattedPool 是否需要更新uid
4. rebuild 构造结果返回