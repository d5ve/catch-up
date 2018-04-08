// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Vote is an object representing the database table.
type Vote struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OptionID  int       `boil:"option_id" json:"option_id" toml:"option_id" yaml:"option_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Voter     string    `boil:"voter" json:"voter" toml:"voter" yaml:"voter"`
	Ynm       string    `boil:"ynm" json:"ynm" toml:"ynm" yaml:"ynm"`

	R *voteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L voteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VoteColumns = struct {
	ID        string
	OptionID  string
	CreatedAt string
	UpdatedAt string
	Voter     string
	Ynm       string
}{
	ID:        "id",
	OptionID:  "option_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Voter:     "voter",
	Ynm:       "ynm",
}

// voteR is where relationships are stored.
type voteR struct {
}

// voteL is where Load methods for each relationship are stored.
type voteL struct{}

var (
	voteColumns               = []string{"id", "option_id", "created_at", "updated_at", "voter", "ynm"}
	voteColumnsWithoutDefault = []string{"option_id", "created_at", "updated_at", "voter", "ynm"}
	voteColumnsWithDefault    = []string{"id"}
	votePrimaryKeyColumns     = []string{"id"}
)

type (
	// VoteSlice is an alias for a slice of pointers to Vote.
	// This should generally be used opposed to []Vote.
	VoteSlice []*Vote
	// VoteHook is the signature for custom Vote hook methods
	VoteHook func(boil.Executor, *Vote) error

	voteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	voteType                 = reflect.TypeOf(&Vote{})
	voteMapping              = queries.MakeStructMapping(voteType)
	votePrimaryKeyMapping, _ = queries.BindMapping(voteType, voteMapping, votePrimaryKeyColumns)
	voteInsertCacheMut       sync.RWMutex
	voteInsertCache          = make(map[string]insertCache)
	voteUpdateCacheMut       sync.RWMutex
	voteUpdateCache          = make(map[string]updateCache)
	voteUpsertCacheMut       sync.RWMutex
	voteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var voteBeforeInsertHooks []VoteHook
var voteBeforeUpdateHooks []VoteHook
var voteBeforeDeleteHooks []VoteHook
var voteBeforeUpsertHooks []VoteHook

var voteAfterInsertHooks []VoteHook
var voteAfterSelectHooks []VoteHook
var voteAfterUpdateHooks []VoteHook
var voteAfterDeleteHooks []VoteHook
var voteAfterUpsertHooks []VoteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Vote) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range voteBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Vote) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range voteBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Vote) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range voteBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Vote) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range voteBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Vote) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range voteAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Vote) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range voteAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Vote) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range voteAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Vote) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range voteAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Vote) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range voteAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVoteHook registers your hook function for all future operations.
func AddVoteHook(hookPoint boil.HookPoint, voteHook VoteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		voteBeforeInsertHooks = append(voteBeforeInsertHooks, voteHook)
	case boil.BeforeUpdateHook:
		voteBeforeUpdateHooks = append(voteBeforeUpdateHooks, voteHook)
	case boil.BeforeDeleteHook:
		voteBeforeDeleteHooks = append(voteBeforeDeleteHooks, voteHook)
	case boil.BeforeUpsertHook:
		voteBeforeUpsertHooks = append(voteBeforeUpsertHooks, voteHook)
	case boil.AfterInsertHook:
		voteAfterInsertHooks = append(voteAfterInsertHooks, voteHook)
	case boil.AfterSelectHook:
		voteAfterSelectHooks = append(voteAfterSelectHooks, voteHook)
	case boil.AfterUpdateHook:
		voteAfterUpdateHooks = append(voteAfterUpdateHooks, voteHook)
	case boil.AfterDeleteHook:
		voteAfterDeleteHooks = append(voteAfterDeleteHooks, voteHook)
	case boil.AfterUpsertHook:
		voteAfterUpsertHooks = append(voteAfterUpsertHooks, voteHook)
	}
}

// OneP returns a single vote record from the query, and panics on error.
func (q voteQuery) OneP() *Vote {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single vote record from the query.
func (q voteQuery) One() (*Vote, error) {
	o := &Vote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for vote")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Vote records from the query, and panics on error.
func (q voteQuery) AllP() VoteSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Vote records from the query.
func (q voteQuery) All() (VoteSlice, error) {
	var o []*Vote

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Vote slice")
	}

	if len(voteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Vote records in the query, and panics on error.
func (q voteQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Vote records in the query.
func (q voteQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count vote rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q voteQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q voteQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if vote exists")
	}

	return count > 0, nil
}

// VotesG retrieves all records.
func VotesG(mods ...qm.QueryMod) voteQuery {
	return Votes(boil.GetDB(), mods...)
}

// Votes retrieves all the records using an executor.
func Votes(exec boil.Executor, mods ...qm.QueryMod) voteQuery {
	mods = append(mods, qm.From("`vote`"))
	return voteQuery{NewQuery(exec, mods...)}
}

// FindVoteG retrieves a single record by ID.
func FindVoteG(id int, selectCols ...string) (*Vote, error) {
	return FindVote(boil.GetDB(), id, selectCols...)
}

// FindVoteGP retrieves a single record by ID, and panics on error.
func FindVoteGP(id int, selectCols ...string) *Vote {
	retobj, err := FindVote(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindVote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVote(exec boil.Executor, id int, selectCols ...string) (*Vote, error) {
	voteObj := &Vote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `vote` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(voteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from vote")
	}

	return voteObj, nil
}

// FindVoteP retrieves a single record by ID with an executor, and panics on error.
func FindVoteP(exec boil.Executor, id int, selectCols ...string) *Vote {
	retobj, err := FindVote(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Vote) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Vote) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Vote) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Vote) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no vote provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(voteColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	voteInsertCacheMut.RLock()
	cache, cached := voteInsertCache[key]
	voteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			voteColumns,
			voteColumnsWithDefault,
			voteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(voteType, voteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(voteType, voteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `vote` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `vote` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `vote` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, votePrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into vote")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == voteMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for vote")
	}

CacheNoHooks:
	if !cached {
		voteInsertCacheMut.Lock()
		voteInsertCache[key] = cache
		voteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Vote record. See Update for
// whitelist behavior description.
func (o *Vote) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Vote record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Vote) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Vote, and panics on error.
// See Update for whitelist behavior description.
func (o *Vote) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Vote.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Vote) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	voteUpdateCacheMut.RLock()
	cache, cached := voteUpdateCache[key]
	voteUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			voteColumns,
			votePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update vote, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `vote` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, votePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(voteType, voteMapping, append(wl, votePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update vote row")
	}

	if !cached {
		voteUpdateCacheMut.Lock()
		voteUpdateCache[key] = cache
		voteUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q voteQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q voteQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for vote")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VoteSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o VoteSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o VoteSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VoteSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `vote` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, votePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in vote slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Vote) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Vote) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Vote) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Vote) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no vote provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(voteColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	voteUpsertCacheMut.RLock()
	cache, cached := voteUpsertCache[key]
	voteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			voteColumns,
			voteColumnsWithDefault,
			voteColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			voteColumns,
			votePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert vote, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "vote", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `vote` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(voteType, voteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(voteType, voteMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for vote")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == voteMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for vote")
	}

CacheNoHooks:
	if !cached {
		voteUpsertCacheMut.Lock()
		voteUpsertCache[key] = cache
		voteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Vote record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Vote) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Vote record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Vote) DeleteG() error {
	if o == nil {
		return errors.New("models: no Vote provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Vote record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Vote) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Vote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Vote) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Vote provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), votePrimaryKeyMapping)
	sql := "DELETE FROM `vote` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from vote")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q voteQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q voteQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no voteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from vote")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o VoteSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o VoteSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Vote slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o VoteSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VoteSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Vote slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(voteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `vote` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, votePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from vote slice")
	}

	if len(voteAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Vote) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Vote) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Vote) ReloadG() error {
	if o == nil {
		return errors.New("models: no Vote provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Vote) Reload(exec boil.Executor) error {
	ret, err := FindVote(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VoteSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VoteSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VoteSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty VoteSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VoteSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	votes := VoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `vote`.* FROM `vote` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, votePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&votes)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VoteSlice")
	}

	*o = votes

	return nil
}

// VoteExists checks if the Vote row exists.
func VoteExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `vote` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if vote exists")
	}

	return exists, nil
}

// VoteExistsG checks if the Vote row exists.
func VoteExistsG(id int) (bool, error) {
	return VoteExists(boil.GetDB(), id)
}

// VoteExistsGP checks if the Vote row exists. Panics on error.
func VoteExistsGP(id int) bool {
	e, err := VoteExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// VoteExistsP checks if the Vote row exists. Panics on error.
func VoteExistsP(exec boil.Executor, id int) bool {
	e, err := VoteExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
