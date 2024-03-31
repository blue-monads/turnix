package syncdtypes

const ROOM_SYS_DATA = "sys.data"
const ROOM_SYS_USERS = "sys.users"
const ROOM_PLUG_DEV = "plugs_dev"

const TAG_REALUSER = "sys.real_user"
const TAG_CONSOLE_CONN = "sys.console_conn"

type Hub interface {
	GetSockd() SockdCore
	GetDataSyncer() DataSyncer
	GetUserSyncer() UserSyncer
}

type DataSyncer interface {
	DataTableSyncer
	DataSheetSyncer
}

type DataTableSyncer interface {
	PushNewRow(source, tenantId, groupId, table string, ids []int64, data any) error
	PushUpdateRow(source, tenantId, groupId, table string, ids []int64, data any) error
	PushDeleteRow(source, tenantId, groupId, table string, ids []int64) error
}

type DataSheetSyncer interface {
	PushSheetNewRow(source, tenantId, groupId string, sheetId int64, ids []int64, data any) error
	PushSheetUpdateRow(source, tenantId, groupId string, sheetId int64, ids []int64, data any) error
	PushSheetDeleteRow(source, tenantId, groupId string, sheetId int64, ids []int64) error
}

type UserSyncer interface {
	//NotifyUser(msg *entities.UserMessage) error
	NotifyMessageRead(tenantId, user string, msgIds []int) error
	NotifyMessageDelete(tenantId, user string, msgIds []int) error
}
