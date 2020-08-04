// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// BackupArchive is an object representing the database table.
type BackupArchive struct {
	ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountID   string      `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Data        null.String `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	RecoveredAt null.Time   `boil:"recovered_at" json:"recovered_at,omitempty" toml:"recovered_at" yaml:"recovered_at,omitempty"`
	DeletedAt   null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *backupArchiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L backupArchiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BackupArchiveColumns = struct {
	ID          string
	AccountID   string
	Data        string
	CreatedAt   string
	RecoveredAt string
	DeletedAt   string
}{
	ID:          "id",
	AccountID:   "account_id",
	Data:        "data",
	CreatedAt:   "created_at",
	RecoveredAt: "recovered_at",
	DeletedAt:   "deleted_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BackupArchiveWhere = struct {
	ID          whereHelperstring
	AccountID   whereHelperstring
	Data        whereHelpernull_String
	CreatedAt   whereHelpertime_Time
	RecoveredAt whereHelpernull_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperstring{field: "\"backup_archive\".\"id\""},
	AccountID:   whereHelperstring{field: "\"backup_archive\".\"account_id\""},
	Data:        whereHelpernull_String{field: "\"backup_archive\".\"data\""},
	CreatedAt:   whereHelpertime_Time{field: "\"backup_archive\".\"created_at\""},
	RecoveredAt: whereHelpernull_Time{field: "\"backup_archive\".\"recovered_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"backup_archive\".\"deleted_at\""},
}

// BackupArchiveRels is where relationship names are stored.
var BackupArchiveRels = struct {
	Account string
}{
	Account: "Account",
}

// backupArchiveR is where relationships are stored.
type backupArchiveR struct {
	Account *Account
}

// NewStruct creates a new relationship struct
func (*backupArchiveR) NewStruct() *backupArchiveR {
	return &backupArchiveR{}
}

// backupArchiveL is where Load methods for each relationship are stored.
type backupArchiveL struct{}

var (
	backupArchiveAllColumns            = []string{"id", "account_id", "data", "created_at", "recovered_at", "deleted_at"}
	backupArchiveColumnsWithoutDefault = []string{"id", "account_id", "data", "recovered_at", "deleted_at"}
	backupArchiveColumnsWithDefault    = []string{"created_at"}
	backupArchivePrimaryKeyColumns     = []string{"id"}
)

type (
	// BackupArchiveSlice is an alias for a slice of pointers to BackupArchive.
	// This should generally be used opposed to []BackupArchive.
	BackupArchiveSlice []*BackupArchive

	backupArchiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	backupArchiveType                 = reflect.TypeOf(&BackupArchive{})
	backupArchiveMapping              = queries.MakeStructMapping(backupArchiveType)
	backupArchivePrimaryKeyMapping, _ = queries.BindMapping(backupArchiveType, backupArchiveMapping, backupArchivePrimaryKeyColumns)
	backupArchiveInsertCacheMut       sync.RWMutex
	backupArchiveInsertCache          = make(map[string]insertCache)
	backupArchiveUpdateCacheMut       sync.RWMutex
	backupArchiveUpdateCache          = make(map[string]updateCache)
	backupArchiveUpsertCacheMut       sync.RWMutex
	backupArchiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single backupArchive record from the query.
func (q backupArchiveQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BackupArchive, error) {
	o := &BackupArchive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: failed to execute a one query for backup_archive")
	}

	return o, nil
}

// All returns all BackupArchive records from the query.
func (q backupArchiveQuery) All(ctx context.Context, exec boil.ContextExecutor) (BackupArchiveSlice, error) {
	var o []*BackupArchive

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler: failed to assign all query results to BackupArchive slice")
	}

	return o, nil
}

// Count returns the count of all BackupArchive records in the query.
func (q backupArchiveQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to count backup_archive rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q backupArchiveQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: failed to check if backup_archive exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *BackupArchive) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"account\"")

	return query
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (backupArchiveL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBackupArchive interface{}, mods queries.Applicator) error {
	var slice []*BackupArchive
	var object *BackupArchive

	if singular {
		object = maybeBackupArchive.(*BackupArchive)
	} else {
		slice = *maybeBackupArchive.(*[]*BackupArchive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &backupArchiveR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &backupArchiveR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`account`), qm.WhereIn(`account.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for account")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Account = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.BackupArchives = append(foreign.R.BackupArchives, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.ID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.BackupArchives = append(foreign.R.BackupArchives, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the backupArchive to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.BackupArchives.
func (o *BackupArchive) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"backup_archive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 2, backupArchivePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.ID
	if o.R == nil {
		o.R = &backupArchiveR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			BackupArchives: BackupArchiveSlice{o},
		}
	} else {
		related.R.BackupArchives = append(related.R.BackupArchives, o)
	}

	return nil
}

// BackupArchives retrieves all the records using an executor.
func BackupArchives(mods ...qm.QueryMod) backupArchiveQuery {
	mods = append(mods, qm.From("\"backup_archive\""))
	return backupArchiveQuery{NewQuery(mods...)}
}

// FindBackupArchive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBackupArchive(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BackupArchive, error) {
	backupArchiveObj := &BackupArchive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"backup_archive\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, backupArchiveObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboiler: unable to select from backup_archive")
	}

	return backupArchiveObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BackupArchive) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no backup_archive provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(backupArchiveColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	backupArchiveInsertCacheMut.RLock()
	cache, cached := backupArchiveInsertCache[key]
	backupArchiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			backupArchiveAllColumns,
			backupArchiveColumnsWithDefault,
			backupArchiveColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(backupArchiveType, backupArchiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(backupArchiveType, backupArchiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"backup_archive\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"backup_archive\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to insert into backup_archive")
	}

	if !cached {
		backupArchiveInsertCacheMut.Lock()
		backupArchiveInsertCache[key] = cache
		backupArchiveInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BackupArchive.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BackupArchive) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	backupArchiveUpdateCacheMut.RLock()
	cache, cached := backupArchiveUpdateCache[key]
	backupArchiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			backupArchiveAllColumns,
			backupArchivePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboiler: unable to update backup_archive, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"backup_archive\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, backupArchivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(backupArchiveType, backupArchiveMapping, append(wl, backupArchivePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update backup_archive row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by update for backup_archive")
	}

	if !cached {
		backupArchiveUpdateCacheMut.Lock()
		backupArchiveUpdateCache[key] = cache
		backupArchiveUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q backupArchiveQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all for backup_archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected for backup_archive")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BackupArchiveSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlboiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), backupArchivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"backup_archive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, backupArchivePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to update all in backupArchive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to retrieve rows affected all in update all backupArchive")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BackupArchive) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboiler: no backup_archive provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(backupArchiveColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	backupArchiveUpsertCacheMut.RLock()
	cache, cached := backupArchiveUpsertCache[key]
	backupArchiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			backupArchiveAllColumns,
			backupArchiveColumnsWithDefault,
			backupArchiveColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			backupArchiveAllColumns,
			backupArchivePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboiler: unable to upsert backup_archive, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(backupArchivePrimaryKeyColumns))
			copy(conflict, backupArchivePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"backup_archive\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(backupArchiveType, backupArchiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(backupArchiveType, backupArchiveMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to upsert backup_archive")
	}

	if !cached {
		backupArchiveUpsertCacheMut.Lock()
		backupArchiveUpsertCache[key] = cache
		backupArchiveUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BackupArchive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BackupArchive) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboiler: no BackupArchive provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), backupArchivePrimaryKeyMapping)
	sql := "DELETE FROM \"backup_archive\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete from backup_archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by delete for backup_archive")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q backupArchiveQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboiler: no backupArchiveQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from backup_archive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for backup_archive")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BackupArchiveSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), backupArchivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"backup_archive\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, backupArchivePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: unable to delete all from backupArchive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboiler: failed to get rows affected by deleteall for backup_archive")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BackupArchive) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBackupArchive(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BackupArchiveSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BackupArchiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), backupArchivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"backup_archive\".* FROM \"backup_archive\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, backupArchivePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboiler: unable to reload all in BackupArchiveSlice")
	}

	*o = slice

	return nil
}

// BackupArchiveExists checks if the BackupArchive row exists.
func BackupArchiveExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"backup_archive\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboiler: unable to check if backup_archive exists")
	}

	return exists, nil
}
