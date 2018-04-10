// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("CatchUps", testCatchUps)
	t.Run("Options", testOptions)
	t.Run("Votes", testVotes)
}

func TestDelete(t *testing.T) {
	t.Run("CatchUps", testCatchUpsDelete)
	t.Run("Options", testOptionsDelete)
	t.Run("Votes", testVotesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("CatchUps", testCatchUpsQueryDeleteAll)
	t.Run("Options", testOptionsQueryDeleteAll)
	t.Run("Votes", testVotesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("CatchUps", testCatchUpsSliceDeleteAll)
	t.Run("Options", testOptionsSliceDeleteAll)
	t.Run("Votes", testVotesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("CatchUps", testCatchUpsExists)
	t.Run("Options", testOptionsExists)
	t.Run("Votes", testVotesExists)
}

func TestFind(t *testing.T) {
	t.Run("CatchUps", testCatchUpsFind)
	t.Run("Options", testOptionsFind)
	t.Run("Votes", testVotesFind)
}

func TestBind(t *testing.T) {
	t.Run("CatchUps", testCatchUpsBind)
	t.Run("Options", testOptionsBind)
	t.Run("Votes", testVotesBind)
}

func TestOne(t *testing.T) {
	t.Run("CatchUps", testCatchUpsOne)
	t.Run("Options", testOptionsOne)
	t.Run("Votes", testVotesOne)
}

func TestAll(t *testing.T) {
	t.Run("CatchUps", testCatchUpsAll)
	t.Run("Options", testOptionsAll)
	t.Run("Votes", testVotesAll)
}

func TestCount(t *testing.T) {
	t.Run("CatchUps", testCatchUpsCount)
	t.Run("Options", testOptionsCount)
	t.Run("Votes", testVotesCount)
}

func TestHooks(t *testing.T) {
	t.Run("CatchUps", testCatchUpsHooks)
	t.Run("Options", testOptionsHooks)
	t.Run("Votes", testVotesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("CatchUps", testCatchUpsInsert)
	t.Run("CatchUps", testCatchUpsInsertWhitelist)
	t.Run("Options", testOptionsInsert)
	t.Run("Options", testOptionsInsertWhitelist)
	t.Run("Votes", testVotesInsert)
	t.Run("Votes", testVotesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("OptionToCatchUpUsingCatchUp", testOptionToOneCatchUpUsingCatchUp)
	t.Run("VoteToOptionUsingOption", testVoteToOneOptionUsingOption)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CatchUpToOptions", testCatchUpToManyOptions)
	t.Run("OptionToVotes", testOptionToManyVotes)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("OptionToCatchUpUsingCatchUp", testOptionToOneSetOpCatchUpUsingCatchUp)
	t.Run("VoteToOptionUsingOption", testVoteToOneSetOpOptionUsingOption)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CatchUpToOptions", testCatchUpToManyAddOpOptions)
	t.Run("OptionToVotes", testOptionToManyAddOpVotes)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("CatchUps", testCatchUpsReload)
	t.Run("Options", testOptionsReload)
	t.Run("Votes", testVotesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("CatchUps", testCatchUpsReloadAll)
	t.Run("Options", testOptionsReloadAll)
	t.Run("Votes", testVotesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("CatchUps", testCatchUpsSelect)
	t.Run("Options", testOptionsSelect)
	t.Run("Votes", testVotesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("CatchUps", testCatchUpsUpdate)
	t.Run("Options", testOptionsUpdate)
	t.Run("Votes", testVotesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("CatchUps", testCatchUpsSliceUpdateAll)
	t.Run("Options", testOptionsSliceUpdateAll)
	t.Run("Votes", testVotesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("CatchUps", testCatchUpsUpsert)
	t.Run("Options", testOptionsUpsert)
	t.Run("Votes", testVotesUpsert)
}
