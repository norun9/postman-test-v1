// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPosts(t *testing.T) {
	t.Parallel()

	query := Posts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPostsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Posts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PostExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Post exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PostExists to return true, but got false.")
	}
}

func testPostsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	postFound, err := FindPost(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if postFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPostsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Posts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPostsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Posts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPostsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Posts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPostsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	postOne := &Post{}
	postTwo := &Post{}
	if err = randomize.Struct(seed, postOne, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}
	if err = randomize.Struct(seed, postTwo, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPostsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(postColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPostsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPostsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Posts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	postDBTypes = map[string]string{`ID`: `bigint`, `Content`: `text`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_           = bytes.MinRead
)

func testPostsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postDBTypes, true, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPostsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Post{}
	if err = randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postDBTypes, true, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(postAllColumns, postPrimaryKeyColumns) {
		fields = postAllColumns
	} else {
		fields = strmangle.SetComplement(
			postAllColumns,
			postPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PostSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPostsUpsert(t *testing.T) {
	t.Parallel()

	if len(postAllColumns) == len(postPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPostUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Post{}
	if err = randomize.Struct(seed, &o, postDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err := Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, postDBTypes, false, postPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Post struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Post: %s", err)
	}

	count, err = Posts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
