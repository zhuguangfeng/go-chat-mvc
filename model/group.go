package model

const TableNameGroup = "group"

type Group struct {
	Base
	OwnerID         int64  `gorm:"column:owner_id;type:bigint;index:idx_owner_status;not null;default:0;comment:群主ID" json:"ownerId"`
	GroupName       string `gorm:"column:group_name;type:varchar(255);not null;default:'';comment:活动群名" json:"name"`
	PeopleNumber    int64  `gorm:"column:people_number;type:int;not null;default:0;comment:群人数" json:"peopleNumber"`
	MaxPeopleNumber int64  `gorm:"column:max_people_number;type:int;not null;default:0;comment:最大群人数" json:"maxPeopleNumber"`
	Category        uint   `gorm:"column:category;type:tinyint;not null;default:0;comment:群聊类型" json:"category"`
	Status          uint   `gorm:"column:status;type:tinyint;index:idx_owner_status;not null;default:0;comment:群聊状态" json:"status"`
}

func (g Group) TableName() string {
	return TableNameGroup
}

type GroupCategory uint

func (g GroupCategory) Uint() uint {
	return uint(g)
}

const (
	GroupCategoryUnknown  GroupCategory = iota
	GroupCategoryActivity               //活动群
	GroupCategoryChat                   //聊天群
)

type GroupStatus uint

func (g GroupStatus) Uint() uint {
	return uint(g)
}

const (
	GroupStatusUnknown GroupStatus = iota
	GroupStatusNormal              //正常
	GroupStatusDisband             //解散
	GroupStatusBan                 //封禁
)
