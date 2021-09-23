# go-db2entity : 通过建表sql，生成entity小工具

在编写Golang代码时，我们有时候会有将数据表转换为 对应的 entity struct 的需求

有时候表的字段特别多的时候，写起来很麻烦，很慢

使用此工具即可快速将 建表sql转换为 entity

---

## Installing

```
git clone https://git.fancydsp.com/scm/v/go-makepolo.git 
cd go-db2entity
go build .
mv go-db2entity $GOPATH/bin
```

---

## 使用方法
- 新建一个文件，形如 go-makepolo/entity/ad_unit.go
- 将建表语句copy入文件内
- pwd：go-makepolo
- 执行命令：go-db2entity ./entity/ad_unit.go
- go-makepolo/entity/ad_unit.go 该文件内容将会自动变成 生成好的 entity结构

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

const TableNameTestAbc = "test_abc"

type TestAbc struct{
	Id                    int64     `json:"id" xorm:"Id pk autoincr"`
	AdunitProperty        string    `json:"adunit_property" xorm:"null  text 'adunit_property'"`
	FirstDayBeginTime     string    `json:"first_day_begin_time" xorm:"not null default '' varchar(128) COMMENT(测试1) 'first_day_begin_time'"`
	VendorCreateTime      time.Time `json:"vendor_create_time" xorm:"null  timestamp COMMENT(测试2) 'vendor_create_time'"`
	VendorPlayableId      string    `json:"vendor_playable_id" xorm:"not null default '' varchar(255) 'vendor_playable_id'"`
	AppExtendPackageId    int32     `json:"app_extend_package_id" xorm:"null default 0 int COMMENT(测试3) 'app_extend_package_id'"`
	PlayableButton        string    `json:"playable_button" xorm:"not null default '' varchar(255) 'playable_button'"`
	PackageName           string    `json:"package_name" xorm:"not null default '' varchar(255) COMMENT(测试4) 'package_name'"`
	DockingType           int32     `json:"docking_type" xorm:"not null default 1 tinyint(2) COMMENT(测试5 ) 'docking_type'"`
	EventAssetType        int32     `json:"event_asset_type" xorm:"not null default 1 tinyint(2) COMMENT(测试6) 'event_asset_type'"`
	SplashAdSwitch        int32     `json:"splash_ad_switch" xorm:"null default 2 int(2) COMMENT(1:开启开屏 2:关闭) 'splash_ad_switch'"`
	AssetIds              string    `json:"asset_ids" xorm:"not null default '0' varchar(100) COMMENT(事件管理资产id) 'asset_ids'"`
}


func (t *TestAbc)TableName() string {
	return TableNameTestAbc
}

func (t *TestAbc) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

type TestAbcList []*TestAbc

type TestAbcMapById map[int64]*TestAbc

func (list TestAbcList) GetIds() []int64 {
	res := make([]int64, 0)
	for _, l := range list {
		res = append(res, l.Id)
	}
	return res
}

func (list TestAbcList) GetMapById() TestAbcMapById {
	res := make(TestAbcMapById)
	for _, l := range list {
		res[l.Id] = l
	}
	return res
}

func (list TestAbcList) string() string {
	res, _ := json.Marshal(list)
	return string(res)
}

// 无论何种类型的值，都可转化为string类型来构建map
func (list TestAbcList) GroupBy(visitor func(item *TestAbc) string) map[string]TestAbcList {
	res := make(map[string]TestAbcList)
	for _, l := range list {
		key := visitor(l)
		value := res[key]
		if value == nil {
			value = make(TestAbcList, 0)
		}
		res[key] = append(value, l)
	}
	return res
}



```
