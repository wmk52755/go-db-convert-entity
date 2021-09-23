
# go-db-convert-entity
This tool can help you convert sql to entities.


# go-db-convert-entity : 通过建表sql，生成entity小工具

在编写Golang代码时，我们有时候会有将数据表转换为 对应的 entity struct 的需求

有时候表的字段特别多的时候，写起来很麻烦，很慢

使用此工具即可快速将 建表sql转换为 entity

---

## Installing

```
git clone https://git.fancydsp.com/scm/v/go-makepolo.git 

cd go-db-convert-entity
go build .
mv go-db-convert-entity $GOPATH/bin
```
你也可以重命名可执行文件，以简化你的命令


## 使用方法
- 新建一个文件，形如 project/entity/ad_unit.go
- 将建表语句copy入文件内
- pwd：go-makepolo
- 执行命令：go-db-convert-entity ./entity/ad_unit.go
- project/entity/ad_unit.go 该文件内容将会自动变成 生成好的 entity结构

## 如图所示：

### 生成前：

```sql
create table dbname.test_abc
(
    id                            bigint auto_increment
        primary key,
    adunit_property               text                         null,
    first_day_begin_time          varchar(128)  default ''     not null comment '测试1',
    vendor_create_time            timestamp                    null comment '测试2',
    vendor_playable_id            varchar(255)  default ''     not null,
    app_extend_package_id         int           default 0      null comment '测试3',
    playable_button               varchar(255)  default ''     not null,
    package_name                  varchar(255)  default ''     not null comment '测试4',
    docking_type                  tinyint(2)    default 1      not null comment '测试5 ',
    event_asset_type              tinyint(2)    default 1      not null comment '测试6',
    splash_ad_switch              int(2)        default 2      null comment '1:开启开屏 2:关闭',
    asset_ids                     varchar(100)  default '0'    not null comment '事件管理资产id',
    constraint name
        unique (campaign_id, name)
)
    collate = utf8mb4_general_ci;

create index idx_status
on dbname.test_abc (app_extend_package_id);

create index idx_sync_status
on dbname.test_abc (package_name);

```

### 生成后：
```go
package entity

import (
	"encoding/json"
	"time"
)

const TableNameAdUnit = "ad_unit"

type AdUnit struct{
	Id                            int64     `json:"id" xorm:"Id pk autoincr"`
	CampaignId                    int64     `json:"campaign_id" xorm:"null  bigint 'campaign_id'"`
	VendorUnitId                  string    `json:"vendor_unit_id" xorm:"not null default '' varchar(128) COMMENT(媒体端adid) 'vendor_unit_id'"`
	VendorId                      int32     `json:"vendor_id" xorm:"not null  int 'vendor_id'"`
	VendorAccountId               int32     `json:"vendor_account_id" xorm:"not null  int COMMENT(广告主账户id) 'vendor_account_id'"`
	AccountId                     int32     `json:"account_id" xorm:"not null  int COMMENT(广告主id) 'account_id'"`
	CompanyId                     int32     `json:"company_id" xorm:"null  int 'company_id'"`
	Name                          string    `json:"name" xorm:"not null  varchar(128) COMMENT(广告名称) 'name'"`
	AppOs                         int32     `json:"app_os" xorm:"null  int 'app_os'"`
	ChargeMode                    int32     `json:"charge_mode" xorm:"not null  int COMMENT(计费方式 1：cpc) 'charge_mode'"`
	BidType                       int32     `json:"bid_type" xorm:"not null  int COMMENT(优化目标出价类型) 'bid_type'"`
	Bid                           int32     `json:"bid" xorm:"not null  int COMMENT(出价) 'bid'"`
	CpaBid                        int32     `json:"cpa_bid" xorm:"null  int 'cpa_bid'"`
	OcpxActionType                int32     `json:"ocpx_action_type" xorm:"not null  int COMMENT(优化目标) 'ocpx_action_type'"`
	DeepConversionType            int32     `json:"deep_conversion_type" xorm:"not null  int COMMENT(深度转化目标) 'deep_conversion_type'"`
	DeepConversionBid             int32     `json:"deep_conversion_bid" xorm:"not null  int COMMENT(深度转化目标出价) 'deep_conversion_bid'"`
	SceneId                       string    `json:"scene_id" xorm:"not null default '' varchar(64) COMMENT(资源位置) 'scene_id'"`
	UnitType                      int32     `json:"unit_type" xorm:"not null  int COMMENT(创意制作方式) 'unit_type'"`
	ScheduleType                  int32     `json:"schedule_type" xorm:"not null  int COMMENT(排期类型 1:从今天开始 2:设置开始时间和结束时间) 'schedule_type'"`
	BeginTime                     string    `json:"begin_time" xorm:"not null  varchar(128) COMMENT(投放开始时间) 'begin_time'"`
	EndTime                       string    `json:"end_time" xorm:"not null  varchar(128) COMMENT(投放结束时间) 'end_time'"`
	ScheduleTimeType              int32     `json:"schedule_time_type" xorm:"null  int COMMENT(投放时间段类型  1:全天  2:特定时间) 'schedule_time_type'"`
	ScheduleTime                  string    `json:"schedule_time" xorm:"not null default '' varchar(1024) COMMENT(投放时间段) 'schedule_time'"`
	BudgetType                    int32     `json:"budget_type" xorm:"null  int 'budget_type'"`
	DayBudget                     int64     `json:"day_budget" xorm:"not null  bigint(11) COMMENT(当日预算) 'day_budget'"`
	DayBudgetSchedule             string    `json:"day_budget_schedule" xorm:"not null default '' varchar(2048) COMMENT(分日预算) 'day_budget_schedule'"`
	UrlType                       int32     `json:"url_type" xorm:"not null  int COMMENT(url类型) 'url_type'"`
	WebUriType                    int32     `json:"web_uri_type" xorm:"null default 0 int COMMENT(1:魔力建站，2：落地页) 'web_uri_type'"`
	Url                           string    `json:"url" xorm:"not null  varchar(2048) COMMENT(投放链接) 'url'"`
	GoodsId                       int32     `json:"goods_id" xorm:"not null  int 'goods_id'"`
	AppId                         int32     `json:"app_id" xorm:"not null default 0 int COMMENT(应用ID) 'app_id'"`
	ProjectId                     int32     `json:"project_id" xorm:"null  int 'project_id'"`
	PackageId                     int32     `json:"package_id" xorm:"not null  int COMMENT(应用渠道包ID) 'package_id'"`
	LandingPageId                 int32     `json:"landing_page_id" xorm:"null  int 'landing_page_id'"`
	MonitorUrl                    string    `json:"monitor_url" xorm:"null  varchar(2048) COMMENT(广告监测url) 'monitor_url'"`
	ShowMode                      int32     `json:"show_mode" xorm:"not null  int COMMENT(创意展现方式) 'show_mode'"`
	Speed                         int32     `json:"speed" xorm:"not null  int COMMENT(投放方式) 'speed'"`
	TargetId                      int64     `json:"target_id" xorm:"null  bigint 'target_id'"`
	TargetTemplateId              int64     `json:"target_template_id" xorm:"null  bigint 'target_template_id'"`
	ConvertId                     int32     `json:"convert_id" xorm:"null  int 'convert_id'"`
	Status                        int32     `json:"status" xorm:"not null  int 'status'"`
	SyncStatus                    int32     `json:"sync_status" xorm:"not null  int COMMENT(同步状态) 'sync_status'"`
	SyncTime                      time.Time `json:"sync_time" xorm:"null  timestamp 'sync_time'"`
	SyncDetail                    string    `json:"sync_detail" xorm:"not null  text COMMENT(同步状态信息) 'sync_detail'"`
	DataVersion                   string    `json:"data_version" xorm:"not null  varchar(128) COMMENT(版本) 'data_version'"`
	PutStatus                     int32     `json:"put_status" xorm:"null  int 'put_status'"`
	CreateChannel                 int32     `json:"create_channel" xorm:"null  int COMMENT(0：投放后台创建；1：Marketing API创建) 'create_channel'"`
	CanCopy                       int32     `json:"can_copy" xorm:"null default 1 int 'can_copy'"`
	UseAppMarket                  int32     `json:"use_app_market" xorm:"not null  int COMMENT(优先从系统应用商店下载 1：优先从系统应用商店下载使用 默认0,不允许) 'use_app_market'"`
	DeliveryPlatform              int32     `json:"delivery_platform" xorm:"not null default 1 int COMMENT(投放平台 1：默认 2：联盟广告) 'delivery_platform'"`
	ReviewDetail                  string    `json:"review_detail" xorm:"null  varchar(128) COMMENT(审核拒绝理由) 'review_detail'"`
	CreateTime                    time.Time `json:"create_time" xorm:"null  timestamp 'create_time'"`
	UpdateTime                    time.Time `json:"update_time" xorm:"null  timestamp 'update_time'"`
	UnionVideoType                string    `json:"union_video_type" xorm:"null  varchar(64) 'union_video_type'"`
	SmartBidType                  string    `json:"smart_bid_type" xorm:"null  varchar(64) 'smart_bid_type'"`
	VendorVersion                 string    `json:"vendor_version" xorm:"null  varchar(64) 'vendor_version'"`
	OpenUrl                       string    `json:"open_url" xorm:"null  varchar(2048) 'open_url'"`
	DownloadType                  string    `json:"download_type" xorm:"null  varchar(64) 'download_type'"`
	DeepBidType                   string    `json:"deep_bid_type" xorm:"null  varchar(64) 'deep_bid_type'"`
	VideoLandingPage              int32     `json:"video_landing_page" xorm:"null default 0 tinyint(1) 'video_landing_page'"`
	ActionbarClickUrl             string    `json:"actionbar_click_url" xorm:"null  varchar(2048) COMMENT(第三方点击按钮监测链接) 'actionbar_click_url'"`
	RoiRatio                      float64   `json:"roi_ratio" xorm:"null default 0 decimal(7, 4) COMMENT(付费ROI系数) 'roi_ratio'"`
	SchemaUri                     string    `json:"schema_uri" xorm:"null default '' varchar(2048) COMMENT(调起链接) 'schema_uri'"`
	SmartCover                    int32     `json:"smart_cover" xorm:"null default 0 tinyint(1) COMMENT(是否智能抽帧) 'smart_cover'"`
	AssetMining                   int32     `json:"asset_mining" xorm:"null default 0 tinyint(1) COMMENT(是否素材挖掘) 'asset_mining'"`
	AdjustCpa                     int32     `json:"adjust_cpa" xorm:"not null default 0 tinyint COMMENT(是否调整自动出价, 0 系统自动计算，1 cpa_bid 必填) 'adjust_cpa'"`
	LubanRoiRatio                 float64   `json:"luban_roi_ratio" xorm:"not null default 0 decimal(7, 4) COMMENT(鲁班roi系数) 'luban_roi_ratio'"`
	PromotionType                 string    `json:"promotion_type" xorm:"not null default '' varchar(32) COMMENT(投放内容) 'promotion_type'"`
	OptimizationGoal              string    `json:"optimization_goal" xorm:"null  varchar(128) COMMENT(优化目标) 'optimization_goal'"`
	BidStrategy                   string    `json:"bid_strategy" xorm:"null  varchar(128) COMMENT(出价策略) 'bid_strategy'"`
	MobileUnionDisplayScene       string    `json:"mobile_union_display_scene" xorm:"null  text COMMENT(优量汇广告展示场景) 'mobile_union_display_scene'"`
	MobileUnionIndustry           string    `json:"mobile_union_industry" xorm:"null  text COMMENT(优量汇行业精选流量包) 'mobile_union_industry'"`
	DeepConversionBehaviorGoal    string    `json:"deep_conversion_behavior_goal" xorm:"null  text COMMENT(深度优化类型) 'deep_conversion_behavior_goal'"`
	UserActionSets                string    `json:"user_action_sets" xorm:"null  text COMMENT(用户行为数据源) 'user_action_sets'"`
	ItemId                        string    `json:"item_id" xorm:"null default '' varchar(255) COMMENT(媒体端商品ID（小店通）) 'item_id'"`
	PulledRejectReason            int32     `json:"pulled_reject_reason" xorm:"not null default 0 tinyint COMMENT(拉取拒审原因,0 代表未拉取，1 代表已拉取) 'pulled_reject_reason'"`
	AwemeAccount                  string    `json:"aweme_account" xorm:"not null default '' varchar(255) COMMENT(抖音账户，营销目标为抖音号推广时必填) 'aweme_account'"`
	VendorIesAccountId            int32     `json:"vendor_ies_account_id" xorm:"not null default 0 int COMMENT(抖音号在vendor_ies_account表中的Id) 'vendor_ies_account_id'"`
	VendorFictionId               int32     `json:"vendor_fiction_id" xorm:"null default 0 int COMMENT(媒体端小说ID) 'vendor_fiction_id'"`
	FictionId                     int32     `json:"fiction_id" xorm:"null default 0 int COMMENT(小说ID) 'fiction_id'"`
	OpenAppUrlId                  int32     `json:"open_app_url_id" xorm:"null default 0 int COMMENT(调起链接ID) 'open_app_url_id'"`
	IsDynamicCreative             int32     `json:"is_dynamic_creative" xorm:"null default 0 int COMMENT(是否用作动态创意) 'is_dynamic_creative'"`
	DynamicCreativeId             int64     `json:"dynamic_creative_id" xorm:"null default 0 bigint COMMENT(动态创意id) 'dynamic_creative_id'"`
	InventoryCategory             int32     `json:"inventory_category" xorm:"null default 0 int(5) 'inventory_category'"`
	InventoryDetail               string    `json:"inventory_detail" xorm:"null default '' varchar(255) 'inventory_detail'"`
	FeedDeliverySearch            int32     `json:"feed_delivery_search" xorm:"not null default 0 int COMMENT(搜索快投功能 0、不启用 1、启用) 'feed_delivery_search'"`
	IntelligentFlowSwitch         int32     `json:"intelligent_flow_switch" xorm:"not null default 0 int COMMENT( 智能流量开关 0、关闭 1、开启 智能流量开关仅在开启搜索快投时有效，默认关闭。) 'intelligent_flow_switch'"`
	KeywordSetId                  int64     `json:"keyword_set_id" xorm:"not null default 0 bigint COMMENT(快投搜索关键词包Id 对应 local_material_keyword_set 表) 'keyword_set_id'"`
	SubscribeUrl                  string    `json:"subscribe_url" xorm:"not null default '' varchar(255) COMMENT(游戏预约链接) 'subscribe_url'"`
	AdvancedCreativeType          int32     `json:"advanced_creative_type" xorm:"not null default 0 int(1) COMMENT(附加创意类型：0、默认值，无 1、游戏表单收集 2、游戏预约按钮) 'advanced_creative_type'"`
	AppDesc                       string    `json:"app_desc" xorm:"not null default '' varchar(128) COMMENT(应用描述，附加创意类型不为0时，必填) 'app_desc'"`
	AppIntroduction               string    `json:"app_introduction" xorm:"not null default '' varchar(128) COMMENT(应用介绍，附加创意类型不为0时，必填) 'app_introduction'"`
	LocalAppThumbnailIds          string    `json:"local_app_thumbnail_ids" xorm:"not null default '' varchar(255) COMMENT(应用图片集，图片id,附加创意类型为1时，只有两个，附加创意类型为2时，仅有3个) 'local_app_thumbnail_ids'"`
	FormId                        int64     `json:"form_id" xorm:"not null default 0 bigint COMMENT(落地页表单id，附加创意类型为1时，必填) 'form_id'"`
	FormIndex                     int64     `json:"form_index" xorm:"not null default 0 bigint COMMENT(落地页表单位置，附加创意类型为1时，必填) 'form_index'"`
	CourseId                      int32     `json:"course_id" xorm:"null default 0 int COMMENT(课程库ID) 'course_id'"`
	VendorCourseId                string    `json:"vendor_course_id" xorm:"null default '0' varchar(255) COMMENT(媒体端课程库ID) 'vendor_course_id'"`
	AdunitProperty                string    `json:"adunit_property" xorm:"null  text 'adunit_property'"`
	FirstDayBeginTime             string    `json:"first_day_begin_time" xorm:"not null default '' varchar(128) COMMENT(首日开始投放时间) 'first_day_begin_time'"`
	VendorCreateTime              time.Time `json:"vendor_create_time" xorm:"null  timestamp COMMENT(媒体端创建时间) 'vendor_create_time'"`
	VendorPlayableId              string    `json:"vendor_playable_id" xorm:"not null default '' varchar(255) 'vendor_playable_id'"`
	AppExtendPackageId            int32     `json:"app_extend_package_id" xorm:"null default 0 int COMMENT(应用分包Id) 'app_extend_package_id'"`
	PlayableButton                string    `json:"playable_button" xorm:"not null default '' varchar(255) 'playable_button'"`
	PlayableId                    int32     `json:"playable_id" xorm:"null default 0 int COMMENT(本地试玩素材ID) 'playable_id'"`
	PlayableUrl                   string    `json:"playable_url" xorm:"null default '' varchar(512) COMMENT(试玩素材url) 'playable_url'"`
	PlayableOrientation           int32     `json:"playable_orientation" xorm:"null default 0 int(2) COMMENT(试玩素材的横竖适配) 'playable_orientation'"`
	VendorConvertId               int64     `json:"vendor_convert_id" xorm:"not null default 0 bigint COMMENT(媒体端转化追踪id) 'vendor_convert_id'"`
	DownloadUrl                   string    `json:"download_url" xorm:"not null default '' varchar(2048) COMMENT(下载链接) 'download_url'"`
	PackageName                   string    `json:"package_name" xorm:"not null default '' varchar(255) COMMENT(应用包名) 'package_name'"`
	DockingType                   int32     `json:"docking_type" xorm:"not null default 1 tinyint(2) COMMENT(头条数据追踪方式 1=转化追踪 2=事件管理) 'docking_type'"`
	EventAssetType                int32     `json:"event_asset_type" xorm:"not null default 1 tinyint(2) COMMENT(事件管理推广内容 1=橙子 2=自研落地页) 'event_asset_type'"`
	SplashAdSwitch                int32     `json:"splash_ad_switch" xorm:"null default 2 int(2) COMMENT(1:开启开屏 2:关闭) 'splash_ad_switch'"`
	AssetIds                      string    `json:"asset_ids" xorm:"not null default '0' varchar(100) COMMENT(事件管理资产id) 'asset_ids'"`
}


func (t *AdUnit)TableName() string {
	return TableNameAdUnit
}

func (t *AdUnit) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

type AdUnitList []*AdUnit

type AdUnitMapById map[int64]*AdUnit

func (list AdUnitList) GetIds() []int64 {
	res := make([]int64, 0)
	for _, l := range list {
		res = append(res, l.Id)
	}
	return res
}

func (list AdUnitList) GetMapById() AdUnitMapById {
	res := make(AdUnitMapById)
	for _, l := range list {
		res[l.Id] = l
	}
	return res
}

func (list AdUnitList) string() string {
	res, _ := json.Marshal(list)
	return string(res)
}

// 无论何种类型的值，都可转化为string类型来构建map
func (list AdUnitList) GroupBy(visitor func(item *AdUnit) string) map[string]AdUnitList {
	res := make(map[string]AdUnitList)
	for _, l := range list {
		key := visitor(l)
		value := res[key]
		if value == nil {
			value = make(AdUnitList, 0)
		}
		res[key] = append(value, l)
	}
	return res
}

```
