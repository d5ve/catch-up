// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testCatchUps(t *testing.T) {
	t.Parallel()

	query := CatchUps(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCatchUpsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = catchUp.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCatchUpsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CatchUps(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCatchUpsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CatchUpSlice{catchUp}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCatchUpsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CatchUpExists(tx, catchUp.ID)
	if err != nil {
		t.Errorf("Unable to check if CatchUp exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CatchUpExistsG to return true, but got false.")
	}
}
func testCatchUpsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	catchUpFound, err := FindCatchUp(tx, catchUp.ID)
	if err != nil {
		t.Error(err)
	}

	if catchUpFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCatchUpsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = CatchUps(tx).Bind(catchUp); err != nil {
		t.Error(err)
	}
}

func testCatchUpsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := CatchUps(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCatchUpsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUpOne := &CatchUp{}
	catchUpTwo := &CatchUp{}
	if err = randomize.Struct(seed, catchUpOne, catchUpDBTypes, false, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}
	if err = randomize.Struct(seed, catchUpTwo, catchUpDBTypes, false, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUpOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = catchUpTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CatchUps(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCatchUpsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	catchUpOne := &CatchUp{}
	catchUpTwo := &CatchUp{}
	if err = randomize.Struct(seed, catchUpOne, catchUpDBTypes, false, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}
	if err = randomize.Struct(seed, catchUpTwo, catchUpDBTypes, false, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUpOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = catchUpTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func catchUpBeforeInsertHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpAfterInsertHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpAfterSelectHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpBeforeUpdateHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpAfterUpdateHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpBeforeDeleteHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpAfterDeleteHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpBeforeUpsertHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func catchUpAfterUpsertHook(e boil.Executor, o *CatchUp) error {
	*o = CatchUp{}
	return nil
}

func testCatchUpsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &CatchUp{}
	o := &CatchUp{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, catchUpDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CatchUp object: %s", err)
	}

	AddCatchUpHook(boil.BeforeInsertHook, catchUpBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	catchUpBeforeInsertHooks = []CatchUpHook{}

	AddCatchUpHook(boil.AfterInsertHook, catchUpAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	catchUpAfterInsertHooks = []CatchUpHook{}

	AddCatchUpHook(boil.AfterSelectHook, catchUpAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	catchUpAfterSelectHooks = []CatchUpHook{}

	AddCatchUpHook(boil.BeforeUpdateHook, catchUpBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	catchUpBeforeUpdateHooks = []CatchUpHook{}

	AddCatchUpHook(boil.AfterUpdateHook, catchUpAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	catchUpAfterUpdateHooks = []CatchUpHook{}

	AddCatchUpHook(boil.BeforeDeleteHook, catchUpBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	catchUpBeforeDeleteHooks = []CatchUpHook{}

	AddCatchUpHook(boil.AfterDeleteHook, catchUpAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	catchUpAfterDeleteHooks = []CatchUpHook{}

	AddCatchUpHook(boil.BeforeUpsertHook, catchUpBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	catchUpBeforeUpsertHooks = []CatchUpHook{}

	AddCatchUpHook(boil.AfterUpsertHook, catchUpAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	catchUpAfterUpsertHooks = []CatchUpHook{}
}
func testCatchUpsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCatchUpsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx, catchUpColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCatchUpToManyOptions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CatchUp
	var b, c Option

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, optionDBTypes, false, optionColumnsWithDefault...)
	randomize.Struct(seed, &c, optionDBTypes, false, optionColumnsWithDefault...)

	b.CatchUpID = a.ID
	c.CatchUpID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	option, err := a.Options(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range option {
		if v.CatchUpID == b.CatchUpID {
			bFound = true
		}
		if v.CatchUpID == c.CatchUpID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CatchUpSlice{&a}
	if err = a.L.LoadOptions(tx, false, (*[]*CatchUp)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Options); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Options = nil
	if err = a.L.LoadOptions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Options); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", option)
	}
}

func testCatchUpToManyAddOpOptions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a CatchUp
	var b, c, d, e Option

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, catchUpDBTypes, false, strmangle.SetComplement(catchUpPrimaryKeyColumns, catchUpColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Option{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Option{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOptions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CatchUpID {
			t.Error("foreign key was wrong value", a.ID, first.CatchUpID)
		}
		if a.ID != second.CatchUpID {
			t.Error("foreign key was wrong value", a.ID, second.CatchUpID)
		}

		if first.R.CatchUp != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CatchUp != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Options[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Options[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Options(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCatchUpsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = catchUp.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCatchUpsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CatchUpSlice{catchUp}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCatchUpsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := CatchUps(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	catchUpDBTypes = map[string]string{`CreatedAt`: `datetime`, `Details`: `varchar`, `EndDate`: `datetime`, `ID`: `int`, `Name`: `varchar`, `StartDate`: `datetime`, `UpdatedAt`: `datetime`}
	_              = bytes.MinRead
)

func testCatchUpsUpdate(t *testing.T) {
	t.Parallel()

	if len(catchUpColumns) == len(catchUpPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	if err = catchUp.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCatchUpsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(catchUpColumns) == len(catchUpPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	catchUp := &CatchUp{}
	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, catchUp, catchUpDBTypes, true, catchUpPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(catchUpColumns, catchUpPrimaryKeyColumns) {
		fields = catchUpColumns
	} else {
		fields = strmangle.SetComplement(
			catchUpColumns,
			catchUpPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(catchUp))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CatchUpSlice{catchUp}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCatchUpsUpsert(t *testing.T) {
	t.Parallel()

	if len(catchUpColumns) == len(catchUpPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	catchUp := CatchUp{}
	if err = randomize.Struct(seed, &catchUp, catchUpDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = catchUp.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert CatchUp: %s", err)
	}

	count, err := CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &catchUp, catchUpDBTypes, false, catchUpPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CatchUp struct: %s", err)
	}

	if err = catchUp.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert CatchUp: %s", err)
	}

	count, err = CatchUps(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
