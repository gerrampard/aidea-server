package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// QuotaN is a Quota object, all fields are nullable
type QuotaN struct {
	original   *quotaOriginal
	quotaModel *QuotaModel

	Id            null.Int    `json:"id"`
	UserId        null.Int    `json:"user_id"`
	Quota         null.Int    `json:"quota"`
	Rest          null.Int    `json:"rest"`
	Note          null.String `json:"note"`
	PaymentId     null.String `json:"payment_id"`
	PeriodStartAt null.Time   `json:"period_start_at"`
	PeriodEndAt   null.Time   `json:"period_end_at"`
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *QuotaN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for Quota
func (inst *QuotaN) SetModel(quotaModel *QuotaModel) {
	inst.quotaModel = quotaModel
}

// quotaOriginal is an object which stores original Quota from database
type quotaOriginal struct {
	Id            null.Int
	UserId        null.Int
	Quota         null.Int
	Rest          null.Int
	Note          null.String
	PaymentId     null.String
	PeriodStartAt null.Time
	PeriodEndAt   null.Time
	CreatedAt     null.Time
	UpdatedAt     null.Time
}

// Staled identify whether the object has been modified
func (inst *QuotaN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &quotaOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.UserId != inst.original.UserId {
			return true
		}
		if inst.Quota != inst.original.Quota {
			return true
		}
		if inst.Rest != inst.original.Rest {
			return true
		}
		if inst.Note != inst.original.Note {
			return true
		}
		if inst.PaymentId != inst.original.PaymentId {
			return true
		}
		if inst.PeriodStartAt != inst.original.PeriodStartAt {
			return true
		}
		if inst.PeriodEndAt != inst.original.PeriodEndAt {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					return true
				}
			case "quota":
				if inst.Quota != inst.original.Quota {
					return true
				}
			case "rest":
				if inst.Rest != inst.original.Rest {
					return true
				}
			case "note":
				if inst.Note != inst.original.Note {
					return true
				}
			case "payment_id":
				if inst.PaymentId != inst.original.PaymentId {
					return true
				}
			case "period_start_at":
				if inst.PeriodStartAt != inst.original.PeriodStartAt {
					return true
				}
			case "period_end_at":
				if inst.PeriodEndAt != inst.original.PeriodEndAt {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *QuotaN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &quotaOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.UserId != inst.original.UserId {
			kv["user_id"] = inst.UserId
		}
		if inst.Quota != inst.original.Quota {
			kv["quota"] = inst.Quota
		}
		if inst.Rest != inst.original.Rest {
			kv["rest"] = inst.Rest
		}
		if inst.Note != inst.original.Note {
			kv["note"] = inst.Note
		}
		if inst.PaymentId != inst.original.PaymentId {
			kv["payment_id"] = inst.PaymentId
		}
		if inst.PeriodStartAt != inst.original.PeriodStartAt {
			kv["period_start_at"] = inst.PeriodStartAt
		}
		if inst.PeriodEndAt != inst.original.PeriodEndAt {
			kv["period_end_at"] = inst.PeriodEndAt
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					kv["user_id"] = inst.UserId
				}
			case "quota":
				if inst.Quota != inst.original.Quota {
					kv["quota"] = inst.Quota
				}
			case "rest":
				if inst.Rest != inst.original.Rest {
					kv["rest"] = inst.Rest
				}
			case "note":
				if inst.Note != inst.original.Note {
					kv["note"] = inst.Note
				}
			case "payment_id":
				if inst.PaymentId != inst.original.PaymentId {
					kv["payment_id"] = inst.PaymentId
				}
			case "period_start_at":
				if inst.PeriodStartAt != inst.original.PeriodStartAt {
					kv["period_start_at"] = inst.PeriodStartAt
				}
			case "period_end_at":
				if inst.PeriodEndAt != inst.original.PeriodEndAt {
					kv["period_end_at"] = inst.PeriodEndAt
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *QuotaN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.quotaModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.quotaModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a quota
func (inst *QuotaN) Delete(ctx context.Context) error {
	if inst.quotaModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.quotaModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *QuotaN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type quotaScope struct {
	name  string
	apply func(builder query.Condition)
}

var quotaGlobalScopes = make([]quotaScope, 0)
var quotaLocalScopes = make([]quotaScope, 0)

// AddGlobalScopeForQuota assign a global scope to a model
func AddGlobalScopeForQuota(name string, apply func(builder query.Condition)) {
	quotaGlobalScopes = append(quotaGlobalScopes, quotaScope{name: name, apply: apply})
}

// AddLocalScopeForQuota assign a local scope to a model
func AddLocalScopeForQuota(name string, apply func(builder query.Condition)) {
	quotaLocalScopes = append(quotaLocalScopes, quotaScope{name: name, apply: apply})
}

func (m *QuotaModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range quotaGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range quotaLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *QuotaModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *QuotaModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type Quota struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	Quota         int64     `json:"quota"`
	Rest          int64     `json:"rest"`
	Note          string    `json:"note"`
	PaymentId     string    `json:"payment_id"`
	PeriodStartAt time.Time `json:"period_start_at"`
	PeriodEndAt   time.Time `json:"period_end_at"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (w Quota) ToQuotaN(allows ...string) QuotaN {
	if len(allows) == 0 {
		return QuotaN{

			Id:            null.IntFrom(int64(w.Id)),
			UserId:        null.IntFrom(int64(w.UserId)),
			Quota:         null.IntFrom(int64(w.Quota)),
			Rest:          null.IntFrom(int64(w.Rest)),
			Note:          null.StringFrom(w.Note),
			PaymentId:     null.StringFrom(w.PaymentId),
			PeriodStartAt: null.TimeFrom(w.PeriodStartAt),
			PeriodEndAt:   null.TimeFrom(w.PeriodEndAt),
			CreatedAt:     null.TimeFrom(w.CreatedAt),
			UpdatedAt:     null.TimeFrom(w.UpdatedAt),
		}
	}

	res := QuotaN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "user_id":
			res.UserId = null.IntFrom(int64(w.UserId))
		case "quota":
			res.Quota = null.IntFrom(int64(w.Quota))
		case "rest":
			res.Rest = null.IntFrom(int64(w.Rest))
		case "note":
			res.Note = null.StringFrom(w.Note)
		case "payment_id":
			res.PaymentId = null.StringFrom(w.PaymentId)
		case "period_start_at":
			res.PeriodStartAt = null.TimeFrom(w.PeriodStartAt)
		case "period_end_at":
			res.PeriodEndAt = null.TimeFrom(w.PeriodEndAt)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w Quota) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *QuotaN) ToQuota() Quota {
	return Quota{

		Id:            w.Id.Int64,
		UserId:        w.UserId.Int64,
		Quota:         w.Quota.Int64,
		Rest:          w.Rest.Int64,
		Note:          w.Note.String,
		PaymentId:     w.PaymentId.String,
		PeriodStartAt: w.PeriodStartAt.Time,
		PeriodEndAt:   w.PeriodEndAt.Time,
		CreatedAt:     w.CreatedAt.Time,
		UpdatedAt:     w.UpdatedAt.Time,
	}
}

// QuotaModel is a model which encapsulates the operations of the object
type QuotaModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var quotaTableName = "quota"

// QuotaTable return table name for Quota
func QuotaTable() string {
	return quotaTableName
}

const (
	FieldQuotaId            = "id"
	FieldQuotaUserId        = "user_id"
	FieldQuotaQuota         = "quota"
	FieldQuotaRest          = "rest"
	FieldQuotaNote          = "note"
	FieldQuotaPaymentId     = "payment_id"
	FieldQuotaPeriodStartAt = "period_start_at"
	FieldQuotaPeriodEndAt   = "period_end_at"
	FieldQuotaCreatedAt     = "created_at"
	FieldQuotaUpdatedAt     = "updated_at"
)

// QuotaFields return all fields in Quota model
func QuotaFields() []string {
	return []string{
		"id",
		"user_id",
		"quota",
		"rest",
		"note",
		"payment_id",
		"period_start_at",
		"period_end_at",
		"created_at",
		"updated_at",
	}
}

func SetQuotaTable(tableName string) {
	quotaTableName = tableName
}

// NewQuotaModel create a QuotaModel
func NewQuotaModel(db query.Database) *QuotaModel {
	return &QuotaModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           quotaTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *QuotaModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *QuotaModel) clone() *QuotaModel {
	return &QuotaModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *QuotaModel) WithoutGlobalScopes(names ...string) *QuotaModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *QuotaModel) WithLocalScopes(names ...string) *QuotaModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *QuotaModel) Condition(builder query.SQLBuilder) *QuotaModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *QuotaModel) Find(ctx context.Context, id int64) (*QuotaN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *QuotaModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *QuotaModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *QuotaModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]QuotaN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *QuotaModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]QuotaN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"user_id",
			"quota",
			"rest",
			"note",
			"payment_id",
			"period_start_at",
			"period_end_at",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "user_id":
			selectFields = append(selectFields, f)
		case "quota":
			selectFields = append(selectFields, f)
		case "rest":
			selectFields = append(selectFields, f)
		case "note":
			selectFields = append(selectFields, f)
		case "payment_id":
			selectFields = append(selectFields, f)
		case "period_start_at":
			selectFields = append(selectFields, f)
		case "period_end_at":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*QuotaN, []interface{}) {
		var quotaVar QuotaN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &quotaVar.Id)
			case "user_id":
				scanFields = append(scanFields, &quotaVar.UserId)
			case "quota":
				scanFields = append(scanFields, &quotaVar.Quota)
			case "rest":
				scanFields = append(scanFields, &quotaVar.Rest)
			case "note":
				scanFields = append(scanFields, &quotaVar.Note)
			case "payment_id":
				scanFields = append(scanFields, &quotaVar.PaymentId)
			case "period_start_at":
				scanFields = append(scanFields, &quotaVar.PeriodStartAt)
			case "period_end_at":
				scanFields = append(scanFields, &quotaVar.PeriodEndAt)
			case "created_at":
				scanFields = append(scanFields, &quotaVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &quotaVar.UpdatedAt)
			}
		}

		return &quotaVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	quotas := make([]QuotaN, 0)
	for rows.Next() {
		quotaReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		quotaReal.original = &quotaOriginal{}
		_ = query.Copy(quotaReal, quotaReal.original)

		quotaReal.SetModel(m)
		quotas = append(quotas, *quotaReal)
	}

	return quotas, nil
}

// First return first result for given query
func (m *QuotaModel) First(ctx context.Context, builders ...query.SQLBuilder) (*QuotaN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new quota to database
func (m *QuotaModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all quotas to database
func (m *QuotaModel) SaveAll(ctx context.Context, quotas []QuotaN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, quota := range quotas {
		id, err := m.Save(ctx, quota)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a quota to database
func (m *QuotaModel) Save(ctx context.Context, quota QuotaN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, quota.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new quota or update it when it has a id > 0
func (m *QuotaModel) SaveOrUpdate(ctx context.Context, quota QuotaN, onlyFields ...string) (id int64, updated bool, err error) {
	if quota.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, quota.Id.Int64, quota, onlyFields...)
		return quota.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, quota, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *QuotaModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *QuotaModel) Update(ctx context.Context, builder query.SQLBuilder, quota QuotaN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, quota.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *QuotaModel) UpdateById(ctx context.Context, id int64, quota QuotaN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, quota.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *QuotaModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *QuotaModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}