// Code generated by core tool ; DO NOT EDIT.

package i_user_item

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"strconv"
	"time"

	"apb-gitlab.abot.sh/apbgo/golib/pkg/db/bean"
	"apb-gitlab.abot.sh/apbgo/golib/pkg/exception"
	"gopkg.in/guregu/null.v3"
)

const (
	TableName       = "i_user_item"
	ConnectionGroup = "default"
)

const (
	ColumnUserId    = "user_id"
	ColumnItemId    = "item_id"
	ColumnCount     = "count"
	ColumnCreatedAt = "created_at"
	ColumnUpdatedAt = "updated_at"
	ColumnDeletedAt = "deleted_at"

	ColumnManageMentComment    = "management_comment"
	ColumnLatestUpdateDatetime = "latest_update_datetime"
)

var (
	PrimaryKeys = []string{
		ColumnUserId,
		ColumnItemId,
	}
	Columns = []string{
		ColumnUserId,
		ColumnItemId,
		ColumnCount,
		ColumnCreatedAt,
		ColumnUpdatedAt,
		ColumnDeletedAt,
	}
)

type IUserItem struct {
	// UserId ユーザーID
	UserId int64 `xorm:"user_id" db:"user_id" json:"user_id"`
	// ItemId アイテムID
	ItemId int64 `xorm:"item_id" db:"item_id" json:"item_id"`
	// Count 所持数
	Count int64 `xorm:"count" db:"count" json:"count"`
	// CreatedAt 作成日時
	CreatedAt time.Time `xorm:"created_at" db:"created_at" json:"created_at"`
	// UpdatedAt 更新日時
	UpdatedAt time.Time `xorm:"updated_at" db:"updated_at" json:"updated_at"`
	// DeletedAt 削除日時
	DeletedAt null.Time `xorm:"deleted_at" db:"deleted_at" json:"deleted_at"`

	original   Original            `xorm:"-" db:"-"`
	cacheState bean.TypeCacheState `xorm:"-" db:"-"`
	cacheOrder int                 `xorm:"-" db:"-"`
}

type Original struct {
	UserId    int64
	ItemId    int64
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (m *IUserItem) Table() string {
	return TableName
}

func (m *IUserItem) ConnectionGroup() string {
	return ConnectionGroup
}

func (m *IUserItem) PrimaryKeys() []string {
	return PrimaryKeys
}

func (m *IUserItem) Columns() []string {
	return Columns
}

func (m *IUserItem) ShardKey() string {
	return ""
}

func (m *IUserItem) ShardValue() interface{} {
	return nil
}

func (m *IUserItem) UpdatedColumns() []string {
	cols := make([]string, 0, 4)

	// 主キーは変更不可

	if m.Count != m.original.Count {
		cols = append(cols, ColumnCount)
	}

	if m.CreatedAt != m.original.CreatedAt {
		cols = append(cols, ColumnCreatedAt)
	}

	if m.UpdatedAt != m.original.UpdatedAt {
		cols = append(cols, ColumnUpdatedAt)
	}

	if m.DeletedAt != m.original.DeletedAt {
		cols = append(cols, ColumnDeletedAt)
	}

	return cols
}

func (m *IUserItem) IsCreated() bool {
	return m.CreatedAt.IsZero()
}

func (m *IUserItem) AsCreated(t time.Time) {
	m.CreatedAt = t
}

func (m *IUserItem) IsUpdated() bool {
	return !m.UpdatedAt.IsZero()
}

func (m *IUserItem) AsUpdated(t time.Time) {
	m.UpdatedAt = t
}

func (m *IUserItem) IsDeleted() bool {
	return !m.DeletedAt.IsZero()
}

func (m *IUserItem) AsDeleted(t time.Time) {
	m.DeletedAt = null.NewTime(t, true)
}

func (m *IUserItem) DeleteColumn() string {
	return ColumnDeletedAt
}

func (m *IUserItem) GetValue(col string) interface{} {
	switch col {
	case ColumnUserId:
		return m.UserId
	case ColumnItemId:
		return m.ItemId
	case ColumnCount:
		return m.Count
	case ColumnCreatedAt:
		return m.CreatedAt
	case ColumnUpdatedAt:
		return m.UpdatedAt
	case ColumnDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func (m *IUserItem) ToStringMap() map[string]string {
	var (
		userId    string
		itemId    string
		count     string
		createdAt string
		updatedAt string
		deletedAt string
	)

	userId = strconv.FormatInt(m.UserId, 10)
	itemId = strconv.FormatInt(m.ItemId, 10)
	count = strconv.FormatInt(m.Count, 10)
	createdAt = m.CreatedAt.Format(time.RFC3339)
	updatedAt = m.UpdatedAt.Format(time.RFC3339)
	if !m.DeletedAt.IsZero() {
		deletedAt = m.DeletedAt.Time.Format(time.RFC3339)
	}
	return map[string]string{
		"user_id":    userId,
		"item_id":    itemId,
		"count":      count,
		"created_at": createdAt,
		"updated_at": updatedAt,
		"deleted_at": deletedAt,
	}
}

func (m *IUserItem) WriteValue(col string, buf *bytes.Buffer) (bool, error) {
	var (
		b []byte
		n bool
	)
	switch col {
	case ColumnUserId:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(m.UserId))
	case ColumnItemId:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(m.ItemId))
	case ColumnCount:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(m.Count))
	case ColumnCreatedAt:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(m.CreatedAt.UnixNano()))
	case ColumnUpdatedAt:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(m.UpdatedAt.UnixNano()))
	case ColumnDeletedAt:
		if !m.DeletedAt.IsZero() {
			b = make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(m.DeletedAt.Time.UnixNano()))
		} else {
			n = true
		}
	}
	_, err := buf.Write(b)
	if err != nil {
		return false, exception.New(exception.UnExpected, "can not write to buffer").Wrap(err)
	}
	return n, nil
}

func (m *IUserItem) Bind(rows *sql.Rows) error {
	if m == nil {
		return exception.New(exception.UnExpected, "model is nil")
	}

	err := rows.Scan(
		&m.UserId,
		&m.ItemId,
		&m.Count,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
	)
	if err != nil {
		return err
	}

	m.UpdateOriginal()
	return nil
}

func (m *IUserItem) UpdateOriginal() {
	m.original.UserId = m.UserId
	m.original.ItemId = m.ItemId
	m.original.Count = m.Count
	m.original.CreatedAt = m.CreatedAt
	m.original.UpdatedAt = m.UpdatedAt
	m.original.DeletedAt = m.DeletedAt
}

func (m *IUserItem) ToBeans() bean.IFBeans {
	return &IUserItems{m}
}

func (m IUserItem) Equal(target IUserItem) bool {

	if m.UserId != target.UserId {
		return false
	}

	if m.ItemId != target.ItemId {
		return false
	}

	if m.Count != target.Count {
		return false
	}

	if m.CreatedAt != target.CreatedAt {
		return false
	}

	if m.UpdatedAt != target.UpdatedAt {
		return false
	}

	if m.DeletedAt != target.DeletedAt {
		return false
	}

	return true
}

func (m *IUserItem) Update(src bean.IFBean) error {
	if m == nil {
		return exception.New(exception.UnExpected, "model is nil")
	}

	if src == nil {
		return exception.New(exception.UnExpected, "src is nil")
	}

	srcModel, ok := src.(*IUserItem)
	if !ok {
		return exception.New(exception.UnExpected, "src is not *IUserItem.")
	}

	m.UserId = srcModel.UserId
	m.ItemId = srcModel.ItemId
	m.Count = srcModel.Count
	m.CreatedAt = srcModel.CreatedAt
	m.UpdatedAt = srcModel.UpdatedAt
	m.DeletedAt = srcModel.DeletedAt

	return nil
}

func (m *IUserItem) Validate() error {
	// floatやdouble, decimal, 細かいテキスト型などは未対応なので、必要な場合はtmplにValidationを追加すること
	if m.UserId < -9223372036854775808 || 9223372036854775807 < m.UserId {
		return exception.New(exception.UnExpected, "validation error. invalid column value.").WithValues(map[string]interface{}{"column": "UserId", "value": m.UserId})
	}
	if m.ItemId < -9223372036854775808 || 9223372036854775807 < m.ItemId {
		return exception.New(exception.UnExpected, "validation error. invalid column value.").WithValues(map[string]interface{}{"column": "ItemId", "value": m.ItemId})
	}
	if m.Count < -9223372036854775808 || 9223372036854775807 < m.Count {
		return exception.New(exception.UnExpected, "validation error. invalid column value.").WithValues(map[string]interface{}{"column": "Count", "value": m.Count})
	}
	return nil
}

func (m *IUserItem) UniqueKey() string {
	return strconv.FormatInt(m.UserId, 10) +
		m.UniqueKeySeparator() +
		strconv.FormatInt(m.ItemId, 10)
}

func (m *IUserItem) UniqueKeySeparator() string {
	return ":"
}

func (m *IUserItem) CacheState() bean.TypeCacheState {
	return m.cacheState
}

func (m *IUserItem) SetCacheState(state bean.TypeCacheState) {
	m.cacheState = state
}

func (m *IUserItem) CacheOrder() int {
	return m.cacheOrder
}

func (m *IUserItem) SetCacheOrder(o int) {
	m.cacheOrder = o
}

func (m *IUserItem) ToCacheableBeans() bean.IFCacheableBeans {
	return &IUserItems{m}
}

type IUserItems []*IUserItem

func (ms IUserItems) Table() string {
	return TableName
}

func (ms IUserItems) ShardKey() string {
	return ""
}

func (ms IUserItems) ConnectionGroup() string {
	return ConnectionGroup
}

func (ms *IUserItems) PrimaryKeys() []string {
	return PrimaryKeys
}

func (ms IUserItems) Columns() []string {
	return Columns
}

func (ms *IUserItems) Reset() {
	v := *ms
	*ms = v[:0]
}

func (ms *IUserItems) Bind(rows *sql.Rows) error {
	m := &IUserItem{}
	err := m.Bind(rows)
	if err != nil {
		return err
	}

	*ms = append(*ms, m)
	return nil
}

func (ms *IUserItems) Add(bean interface{}) error {
	m, ok := bean.(*IUserItem)
	if !ok {
		return exception.New(exception.UnExpected, "引数が型と一致しません. 型=iUserItem")
	}

	*ms = append(*ms, m)
	return nil
}

func (ms IUserItems) ForeachBean(f func(bean bean.IFBean) error) error {
	for i := range ms {
		if err := f(ms[i]); err != nil {
			return err
		}
	}
	return nil
}

func (ms IUserItems) ForeachCacheableBean(f func(bean bean.IFCacheableBean) error) error {
	for i := range ms {
		if err := f(ms[i]); err != nil {
			return err
		}
	}
	return nil
}

func (ms IUserItems) Len() int {
	return len(ms)
}

func (ms IUserItems) DeleteColumn() string {
	return ColumnDeletedAt
}
