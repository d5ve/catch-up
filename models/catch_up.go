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

// CatchUp is an object representing the database table.
type CatchUp struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Details   string    `boil:"details" json:"details" toml:"details" yaml:"details"`
	StartDate time.Time `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate   time.Time `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`

	R *catchUpR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L catchUpL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CatchUpColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	Name      string
	Details   string
	StartDate string
	EndDate   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Name:      "name",
	Details:   "details",
	StartDate: "start_date",
	EndDate:   "end_date",
}

// catchUpR is where relationships are stored.
type catchUpR struct {
	Options OptionSlice
}

// catchUpL is where Load methods for each relationship are stored.
type catchUpL struct{}

var (
	catchUpColumns               = []string{"id", "created_at", "updated_at", "name", "details", "start_date", "end_date"}
	catchUpColumnsWithoutDefault = []string{"created_at", "updated_at", "name", "details", "start_date", "end_date"}
	catchUpColumnsWithDefault    = []string{"id"}
	catchUpPrimaryKeyColumns     = []string{"id"}
)

type (
	// CatchUpSlice is an alias for a slice of pointers to CatchUp.
	// This should generally be used opposed to []CatchUp.
	CatchUpSlice []*CatchUp
	// CatchUpHook is the signature for custom CatchUp hook methods
	CatchUpHook func(boil.Executor, *CatchUp) error

	catchUpQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	catchUpType                 = reflect.TypeOf(&CatchUp{})
	catchUpMapping              = queries.MakeStructMapping(catchUpType)
	catchUpPrimaryKeyMapping, _ = queries.BindMapping(catchUpType, catchUpMapping, catchUpPrimaryKeyColumns)
	catchUpInsertCacheMut       sync.RWMutex
	catchUpInsertCache          = make(map[string]insertCache)
	catchUpUpdateCacheMut       sync.RWMutex
	catchUpUpdateCache          = make(map[string]updateCache)
	catchUpUpsertCacheMut       sync.RWMutex
	catchUpUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var catchUpBeforeInsertHooks []CatchUpHook
var catchUpBeforeUpdateHooks []CatchUpHook
var catchUpBeforeDeleteHooks []CatchUpHook
var catchUpBeforeUpsertHooks []CatchUpHook

var catchUpAfterInsertHooks []CatchUpHook
var catchUpAfterSelectHooks []CatchUpHook
var catchUpAfterUpdateHooks []CatchUpHook
var catchUpAfterDeleteHooks []CatchUpHook
var catchUpAfterUpsertHooks []CatchUpHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CatchUp) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CatchUp) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CatchUp) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CatchUp) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CatchUp) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CatchUp) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CatchUp) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CatchUp) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CatchUp) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range catchUpAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCatchUpHook registers your hook function for all future operations.
func AddCatchUpHook(hookPoint boil.HookPoint, catchUpHook CatchUpHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		catchUpBeforeInsertHooks = append(catchUpBeforeInsertHooks, catchUpHook)
	case boil.BeforeUpdateHook:
		catchUpBeforeUpdateHooks = append(catchUpBeforeUpdateHooks, catchUpHook)
	case boil.BeforeDeleteHook:
		catchUpBeforeDeleteHooks = append(catchUpBeforeDeleteHooks, catchUpHook)
	case boil.BeforeUpsertHook:
		catchUpBeforeUpsertHooks = append(catchUpBeforeUpsertHooks, catchUpHook)
	case boil.AfterInsertHook:
		catchUpAfterInsertHooks = append(catchUpAfterInsertHooks, catchUpHook)
	case boil.AfterSelectHook:
		catchUpAfterSelectHooks = append(catchUpAfterSelectHooks, catchUpHook)
	case boil.AfterUpdateHook:
		catchUpAfterUpdateHooks = append(catchUpAfterUpdateHooks, catchUpHook)
	case boil.AfterDeleteHook:
		catchUpAfterDeleteHooks = append(catchUpAfterDeleteHooks, catchUpHook)
	case boil.AfterUpsertHook:
		catchUpAfterUpsertHooks = append(catchUpAfterUpsertHooks, catchUpHook)
	}
}

// OneP returns a single catchUp record from the query, and panics on error.
func (q catchUpQuery) OneP() *CatchUp {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single catchUp record from the query.
func (q catchUpQuery) One() (*CatchUp, error) {
	o := &CatchUp{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for catch_up")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all CatchUp records from the query, and panics on error.
func (q catchUpQuery) AllP() CatchUpSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CatchUp records from the query.
func (q catchUpQuery) All() (CatchUpSlice, error) {
	var o []*CatchUp

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CatchUp slice")
	}

	if len(catchUpAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all CatchUp records in the query, and panics on error.
func (q catchUpQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CatchUp records in the query.
func (q catchUpQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count catch_up rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q catchUpQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q catchUpQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if catch_up exists")
	}

	return count > 0, nil
}

// OptionsG retrieves all the option's option.
func (o *CatchUp) OptionsG(mods ...qm.QueryMod) optionQuery {
	return o.Options(boil.GetDB(), mods...)
}

// Options retrieves all the option's option with an executor.
func (o *CatchUp) Options(exec boil.Executor, mods ...qm.QueryMod) optionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`option`.`catch_up_id`=?", o.ID),
	)

	query := Options(exec, queryMods...)
	queries.SetFrom(query.Query, "`option`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`option`.*"})
	}

	return query
}

// LoadOptions allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (catchUpL) LoadOptions(e boil.Executor, singular bool, maybeCatchUp interface{}) error {
	var slice []*CatchUp
	var object *CatchUp

	count := 1
	if singular {
		object = maybeCatchUp.(*CatchUp)
	} else {
		slice = *maybeCatchUp.(*[]*CatchUp)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &catchUpR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &catchUpR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from `option` where `catch_up_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load option")
	}
	defer results.Close()

	var resultSlice []*Option
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice option")
	}

	if len(optionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Options = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CatchUpID {
				local.R.Options = append(local.R.Options, foreign)
				break
			}
		}
	}

	return nil
}

// AddOptionsG adds the given related objects to the existing relationships
// of the catch_up, optionally inserting them as new records.
// Appends related to o.R.Options.
// Sets related.R.CatchUp appropriately.
// Uses the global database handle.
func (o *CatchUp) AddOptionsG(insert bool, related ...*Option) error {
	return o.AddOptions(boil.GetDB(), insert, related...)
}

// AddOptionsP adds the given related objects to the existing relationships
// of the catch_up, optionally inserting them as new records.
// Appends related to o.R.Options.
// Sets related.R.CatchUp appropriately.
// Panics on error.
func (o *CatchUp) AddOptionsP(exec boil.Executor, insert bool, related ...*Option) {
	if err := o.AddOptions(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOptionsGP adds the given related objects to the existing relationships
// of the catch_up, optionally inserting them as new records.
// Appends related to o.R.Options.
// Sets related.R.CatchUp appropriately.
// Uses the global database handle and panics on error.
func (o *CatchUp) AddOptionsGP(insert bool, related ...*Option) {
	if err := o.AddOptions(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddOptions adds the given related objects to the existing relationships
// of the catch_up, optionally inserting them as new records.
// Appends related to o.R.Options.
// Sets related.R.CatchUp appropriately.
func (o *CatchUp) AddOptions(exec boil.Executor, insert bool, related ...*Option) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CatchUpID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `option` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"catch_up_id"}),
				strmangle.WhereClause("`", "`", 0, optionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CatchUpID = o.ID
		}
	}

	if o.R == nil {
		o.R = &catchUpR{
			Options: related,
		}
	} else {
		o.R.Options = append(o.R.Options, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &optionR{
				CatchUp: o,
			}
		} else {
			rel.R.CatchUp = o
		}
	}
	return nil
}

// CatchUpsG retrieves all records.
func CatchUpsG(mods ...qm.QueryMod) catchUpQuery {
	return CatchUps(boil.GetDB(), mods...)
}

// CatchUps retrieves all the records using an executor.
func CatchUps(exec boil.Executor, mods ...qm.QueryMod) catchUpQuery {
	mods = append(mods, qm.From("`catch_up`"))
	return catchUpQuery{NewQuery(exec, mods...)}
}

// FindCatchUpG retrieves a single record by ID.
func FindCatchUpG(id int, selectCols ...string) (*CatchUp, error) {
	return FindCatchUp(boil.GetDB(), id, selectCols...)
}

// FindCatchUpGP retrieves a single record by ID, and panics on error.
func FindCatchUpGP(id int, selectCols ...string) *CatchUp {
	retobj, err := FindCatchUp(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCatchUp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCatchUp(exec boil.Executor, id int, selectCols ...string) (*CatchUp, error) {
	catchUpObj := &CatchUp{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `catch_up` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(catchUpObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from catch_up")
	}

	return catchUpObj, nil
}

// FindCatchUpP retrieves a single record by ID with an executor, and panics on error.
func FindCatchUpP(exec boil.Executor, id int, selectCols ...string) *CatchUp {
	retobj, err := FindCatchUp(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CatchUp) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CatchUp) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CatchUp) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CatchUp) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no catch_up provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(catchUpColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	catchUpInsertCacheMut.RLock()
	cache, cached := catchUpInsertCache[key]
	catchUpInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			catchUpColumns,
			catchUpColumnsWithDefault,
			catchUpColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(catchUpType, catchUpMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(catchUpType, catchUpMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `catch_up` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `catch_up` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `catch_up` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, catchUpPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into catch_up")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == catchUpMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for catch_up")
	}

CacheNoHooks:
	if !cached {
		catchUpInsertCacheMut.Lock()
		catchUpInsertCache[key] = cache
		catchUpInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single CatchUp record. See Update for
// whitelist behavior description.
func (o *CatchUp) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CatchUp record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CatchUp) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CatchUp, and panics on error.
// See Update for whitelist behavior description.
func (o *CatchUp) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CatchUp.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CatchUp) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	catchUpUpdateCacheMut.RLock()
	cache, cached := catchUpUpdateCache[key]
	catchUpUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			catchUpColumns,
			catchUpPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update catch_up, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `catch_up` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, catchUpPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(catchUpType, catchUpMapping, append(wl, catchUpPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update catch_up row")
	}

	if !cached {
		catchUpUpdateCacheMut.Lock()
		catchUpUpdateCache[key] = cache
		catchUpUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q catchUpQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q catchUpQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for catch_up")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CatchUpSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CatchUpSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CatchUpSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CatchUpSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catchUpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `catch_up` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catchUpPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in catchUp slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CatchUp) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CatchUp) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CatchUp) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CatchUp) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no catch_up provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(catchUpColumnsWithDefault, o)

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

	catchUpUpsertCacheMut.RLock()
	cache, cached := catchUpUpsertCache[key]
	catchUpUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			catchUpColumns,
			catchUpColumnsWithDefault,
			catchUpColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			catchUpColumns,
			catchUpPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert catch_up, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "catch_up", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `catch_up` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(catchUpType, catchUpMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(catchUpType, catchUpMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for catch_up")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == catchUpMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for catch_up")
	}

CacheNoHooks:
	if !cached {
		catchUpUpsertCacheMut.Lock()
		catchUpUpsertCache[key] = cache
		catchUpUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single CatchUp record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CatchUp) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CatchUp record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CatchUp) DeleteG() error {
	if o == nil {
		return errors.New("models: no CatchUp provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CatchUp record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CatchUp) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CatchUp record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CatchUp) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CatchUp provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), catchUpPrimaryKeyMapping)
	sql := "DELETE FROM `catch_up` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from catch_up")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q catchUpQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q catchUpQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no catchUpQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from catch_up")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CatchUpSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CatchUpSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no CatchUp slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CatchUpSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CatchUpSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CatchUp slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(catchUpBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catchUpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `catch_up` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catchUpPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from catchUp slice")
	}

	if len(catchUpAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CatchUp) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CatchUp) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CatchUp) ReloadG() error {
	if o == nil {
		return errors.New("models: no CatchUp provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CatchUp) Reload(exec boil.Executor) error {
	ret, err := FindCatchUp(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CatchUpSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CatchUpSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CatchUpSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CatchUpSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CatchUpSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	catchUps := CatchUpSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), catchUpPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `catch_up`.* FROM `catch_up` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, catchUpPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&catchUps)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CatchUpSlice")
	}

	*o = catchUps

	return nil
}

// CatchUpExists checks if the CatchUp row exists.
func CatchUpExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `catch_up` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if catch_up exists")
	}

	return exists, nil
}

// CatchUpExistsG checks if the CatchUp row exists.
func CatchUpExistsG(id int) (bool, error) {
	return CatchUpExists(boil.GetDB(), id)
}

// CatchUpExistsGP checks if the CatchUp row exists. Panics on error.
func CatchUpExistsGP(id int) bool {
	e, err := CatchUpExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CatchUpExistsP checks if the CatchUp row exists. Panics on error.
func CatchUpExistsP(exec boil.Executor, id int) bool {
	e, err := CatchUpExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
