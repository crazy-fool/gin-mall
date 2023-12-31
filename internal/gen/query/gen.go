// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:        db,
		Category:  newCategory(db, opts...),
		Customer:  newCustomer(db, opts...),
		Sku:       newSku(db, opts...),
		SkuDetail: newSkuDetail(db, opts...),
		SkuStock:  newSkuStock(db, opts...),
		Spec:      newSpec(db, opts...),
		SpecGroup: newSpecGroup(db, opts...),
		Spu:       newSpu(db, opts...),
		SpuDetail: newSpuDetail(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Category  category
	Customer  customer
	Sku       sku
	SkuDetail skuDetail
	SkuStock  skuStock
	Spec      spec
	SpecGroup specGroup
	Spu       spu
	SpuDetail spuDetail
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Category:  q.Category.clone(db),
		Customer:  q.Customer.clone(db),
		Sku:       q.Sku.clone(db),
		SkuDetail: q.SkuDetail.clone(db),
		SkuStock:  q.SkuStock.clone(db),
		Spec:      q.Spec.clone(db),
		SpecGroup: q.SpecGroup.clone(db),
		Spu:       q.Spu.clone(db),
		SpuDetail: q.SpuDetail.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Category:  q.Category.replaceDB(db),
		Customer:  q.Customer.replaceDB(db),
		Sku:       q.Sku.replaceDB(db),
		SkuDetail: q.SkuDetail.replaceDB(db),
		SkuStock:  q.SkuStock.replaceDB(db),
		Spec:      q.Spec.replaceDB(db),
		SpecGroup: q.SpecGroup.replaceDB(db),
		Spu:       q.Spu.replaceDB(db),
		SpuDetail: q.SpuDetail.replaceDB(db),
	}
}

type queryCtx struct {
	Category  ICategoryDo
	Customer  ICustomerDo
	Sku       ISkuDo
	SkuDetail ISkuDetailDo
	SkuStock  ISkuStockDo
	Spec      ISpecDo
	SpecGroup ISpecGroupDo
	Spu       ISpuDo
	SpuDetail ISpuDetailDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Category:  q.Category.WithContext(ctx),
		Customer:  q.Customer.WithContext(ctx),
		Sku:       q.Sku.WithContext(ctx),
		SkuDetail: q.SkuDetail.WithContext(ctx),
		SkuStock:  q.SkuStock.WithContext(ctx),
		Spec:      q.Spec.WithContext(ctx),
		SpecGroup: q.SpecGroup.WithContext(ctx),
		Spu:       q.Spu.WithContext(ctx),
		SpuDetail: q.SpuDetail.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
