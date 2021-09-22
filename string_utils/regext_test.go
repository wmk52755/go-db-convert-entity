package string_utils

import (
	"fmt"
	"testing"
)

func TestAaa(t *testing.T) {
	a := "keyword_id   varchar(255) default '' null comment '关键\\'''\\\\\\'\\'\\'\\\\''词id',"
	fmt.Println(string(a[1]))
	s, m := ReplaceAllQuotationMarks(a)
	fmt.Println(s)
	fmt.Println(m)
}

func TestGetContentInFirstBrackets(t *testing.T) {
	a := "create table makepolo.ad_unit\n(\n    id                            bigint auto_increment\n        primary key,\n    campaign_id                   bigint                       null,\n    vendor_unit_id                varchar(128)  default ''     not null comment '媒体端adid',\n    vendor_id                     int                          not null,\n    vendor_account_id             int                          not null comment '广告主账户id',\n    account_id                    int                          not null comment '广告主id',\n    company_id                    int                          null,\n    name                          varchar(128)                 not null comment '广告名称',\n    app_os                        int                          null,\n    charge_mode                   int                          not null comment '计费方式 1：cpc',\n    bid_type                      int                          not null comment '优化目标出价类型',\n    bid                           int                          not null comment '出价',\n    cpa_bid                       int                          null,\n    ocpx_action_type              int                          not null comment '优化目标',\n    deep_conversion_type          int                          not null comment '深度转化目标',\n    deep_conversion_bid           int                          not null comment '深度转化目标出价',\n    scene_id                      varchar(64)   default ''     not null comment '资源位置',\n    unit_type                     int                          not null comment '创意制作方式',\n    schedule_type                 int                          not null comment '排期类型 1:从今天开始 2:设置开始时间和结束时间',\n    begin_time                    varchar(128)                 not null comment '投放开始时间',\n    end_time                      varchar(128)                 not null comment '投放结束时间',\n    schedule_time_type            int                          null comment '投放时间段类型  1:全天  2:特定时间',\n    schedule_time                 varchar(1024) default ''     not null comment '投放时间段',\n    budget_type                   int                          null,\n    day_budget                    bigint(11)                   not null comment '当日预算',\n    day_budget_schedule           varchar(2048) default ''     not null comment '分日预算',\n    url_type                      int                          not null comment 'url类型',\n    web_uri_type                  int           default 0      null comment '1:魔力建站，2：落地页',\n    url                           varchar(2048)                not null comment '投放链接',\n    goods_id                      int                          not null,\n    app_id                        int           default 0      not null comment '应用ID',\n    project_id                    int                          null,\n    package_id                    int                          not null comment '应用渠道包ID',\n    landing_page_id               int                          null,\n    monitor_url                   varchar(2048)                null comment '广告监测url',\n    show_mode                     int                          not null comment '创意展现方式',\n    speed                         int                          not null comment '投放方式',\n    target_id                     bigint                       null,\n    target_template_id            bigint                       null,\n    convert_id                    int                          null,\n    status                        int                          not null,\n    sync_status                   int                          not null comment '同步状态',\n    sync_time                     timestamp                    null,\n    sync_detail                   text                         not null comment '同步状态信息',\n    data_version                  varchar(128)                 not null comment '版本',\n    put_status                    int                          null,\n    create_channel                int                          null comment '0：投放后台创建；1：Marketing API创建',\n    can_copy                      int           default 1      null,\n    use_app_market                int                          not null comment '优先从系统应用商店下载 1：优先从系统应用商店下载使用 默认0,不允许',\n    delivery_platform             int           default 1      not null comment '投放平台 1：默认 2：联盟广告',\n    review_detail                 varchar(128)                 null comment '审核拒绝理由',\n    create_time                   timestamp                    null,\n    update_time                   timestamp                    null on update CURRENT_TIMESTAMP,\n    union_video_type              varchar(64)                  null,\n    smart_bid_type                varchar(64)                  null,\n    vendor_version                varchar(64)                  null,\n    open_url                      varchar(2048)                null,\n    download_type                 varchar(64)                  null,\n    deep_bid_type                 varchar(64)                  null,\n    video_landing_page            tinyint(1)    default 0      null,\n    actionbar_click_url           varchar(2048)                null comment '第三方点击按钮监测链接',\n    roi_ratio                     decimal(7, 4) default 0.0000 null comment '付费ROI系数',\n    schema_uri                    varchar(2048) default ''     null comment '调起链接',\n    smart_cover                   tinyint(1)    default 0      null comment '是否智能抽帧',\n    asset_mining                  tinyint(1)    default 0      null comment '是否素材挖掘',\n    adjust_cpa                    tinyint       default 0      not null comment '是否调整自动出价, 0 系统自动计算，1 cpa_bid 必填',\n    luban_roi_ratio               decimal(7, 4) default 0.0000 not null comment '鲁班roi系数',\n    promotion_type                varchar(32)   default ''     not null comment '投放内容',\n    optimization_goal             varchar(128)                 null comment '优化目标',\n    bid_strategy                  varchar(128)                 null comment '出价策略',\n    mobile_union_display_scene    text                         null comment '优量汇广告展示场景',\n    mobile_union_industry         text                         null comment '优量汇行业精选流量包',\n    deep_conversion_behavior_goal text                         null comment '深度优化类型',\n    user_action_sets              text                         null comment '用户行为数据源',\n    item_id                       varchar(255)  default ''     null comment '媒体端商品ID（小店通）',\n    pulled_reject_reason          tinyint       default 0      not null comment '拉取拒审原因,0 代表未拉取，1 代表已拉取',\n    aweme_account                 varchar(255)  default ''     not null comment '抖音账户，营销目标为抖音号推广时必填',\n    vendor_ies_account_id         int           default 0      not null comment '抖音号在vendor_ies_account表中的Id',\n    vendor_fiction_id             int           default 0      null comment '媒体端小说ID',\n    fiction_id                    int           default 0      null comment '小说ID',\n    open_app_url_id               int           default 0      null comment '调起链接ID',\n    is_dynamic_creative           int           default 0      null comment '是否用作动态创意',\n    dynamic_creative_id           bigint        default 0      null comment '动态创意id',\n    inventory_category            int(5)        default 0      null,\n    inventory_detail              varchar(255)  default ''     null,\n    feed_delivery_search          int           default 0      not null comment '搜索快投功能 0、不启用 1、启用',\n    intelligent_flow_switch       int           default 0      not null comment ' 智能流量开关 0、关闭 1、开启 智能流量开关仅在开启搜索快投时有效，默认关闭。',\n    keyword_set_id                bigint        default 0      not null comment '快投搜索关键词包Id 对应 local_material_keyword_set 表',\n    subscribe_url                 varchar(255)  default ''     not null comment '游戏预约链接',\n    advanced_creative_type        int(1)        default 0      not null comment '附加创意类型：0、默认值，无 1、游戏表单收集 2、游戏预约按钮',\n    app_desc                      varchar(128)  default ''     not null comment '应用描述，附加创意类型不为0时，必填',\n    app_introduction              varchar(128)  default ''     not null comment '应用介绍，附加创意类型不为0时，必填',\n    local_app_thumbnail_ids       varchar(255)  default ''     not null comment '应用图片集，图片id,附加创意类型为1时，只有两个，附加创意类型为2时，仅有3个',\n    form_id                       bigint        default 0      not null comment '落地页表单id，附加创意类型为1时，必填',\n    form_index                    bigint        default 0      not null comment '落地页表单位置，附加创意类型为1时，必填',\n    course_id                     int           default 0      null comment '课程库ID',\n    vendor_course_id              varchar(255)  default '0'    null comment '媒体端课程库ID',\n    adunit_property               text                         null,\n    first_day_begin_time          varchar(128)  default ''     not null comment '首日开始投放时间',\n    vendor_create_time            timestamp                    null comment '媒体端创建时间',\n    vendor_playable_id            varchar(255)  default ''     not null,\n    app_extend_package_id         int           default 0      null comment '应用分包Id',\n    playable_button               varchar(255)  default ''     not null,\n    playable_id                   int           default 0      null comment '本地试玩素材ID',\n    playable_url                  varchar(512)  default ''     null comment '试玩素材url',\n    playable_orientation          int(2)        default 0      null comment '试玩素材的横竖适配',\n    vendor_convert_id             bigint        default 0      not null comment '媒体端转化追踪id',\n    download_url                  varchar(2048) default ''     not null comment '下载链接',\n    package_name                  varchar(255)  default ''     not null comment '应用包名',\n    docking_type                  tinyint(2)    default 1      not null comment '头条数据追踪方式 1=转化追踪 2=事件管理',\n    event_asset_type              tinyint(2)    default 1      not null comment '事件管理推广内容 1=橙子 2=自研落地页',\n    splash_ad_switch              int(2)        default 2      null comment '1:开启开屏 2:关闭',\n    asset_ids                     varchar(100)  default '0'    not null comment '事件管理资产id',\n    constraint name\n        unique (campaign_id, name)\n)\n    collate = utf8mb4_general_ci;\n\ncreate index idx_app_extend_package_id_vendor_id_campaign_id_index\n    on makepolo.ad_unit (app_extend_package_id, vendor_id, campaign_id);\n\ncreate index idx_companyid_vendoraccountid_syncstatus_status\n    on makepolo.ad_unit (company_id, vendor_account_id, sync_status, status);\n\ncreate index idx_convert_id\n    on makepolo.ad_unit (convert_id);\n\ncreate index idx_package_id\n    on makepolo.ad_unit (package_id);\n\ncreate index idx_status\n    on makepolo.ad_unit (status);\n\ncreate index idx_sync_status\n    on makepolo.ad_unit (sync_status);\n\n"
	b, err := GetContentInFirstBrackets(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}