package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//go:generate goqueryset -in rating.go

// Rating struct represent rating model. Next line (gen:qs) is needed to autogenerate RatingQuerySet.
// gen:qs
type Rating struct {
	ID           int64      `json:"id"`
	UserID       int64      `json:"userId" validate:"required"`
	LessonID     int64      `json:"lessonId" validate:"required"`
	RatingNumber int64      `json:"rating" validate:"required"`
	UpdatedAt    time.Time  `json:"updated_at"`
	CreatedAt    time.Time  `json:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// ===== BEGIN of all query sets

// ===== BEGIN of query set RatingQuerySet

// RatingQuerySet is an queryset type for Rating
type RatingQuerySet struct {
	db *gorm.DB
}

// NewRatingQuerySet constructs new RatingQuerySet
func NewRatingQuerySet(db *gorm.DB) RatingQuerySet {
	return RatingQuerySet{
		db: db.Model(&Rating{}),
	}
}

func (qs RatingQuerySet) w(db *gorm.DB) RatingQuerySet {
	return NewRatingQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *Rating) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Rating) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) All(ret *[]Rating) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtEq(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtGt(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtGte(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtLt(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtLte(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) CreatedAtNe(createdAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) Delete() error {
	return qs.db.Delete(Rating{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Rating{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Rating{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtEq(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtGt(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtGte(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtIsNotNull() RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtIsNull() RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtLt(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtLte(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) DeletedAtNe(deletedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) GetUpdater() RatingUpdater {
	return NewRatingUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDEq(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDGt(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDGte(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDIn(ID ...int64) RatingQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDLt(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDLte(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDNe(ID int64) RatingQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) IDNotIn(ID ...int64) RatingQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// LessonIDEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDEq(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id = ?", lessonID))
}

// LessonIDGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDGt(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id > ?", lessonID))
}

// LessonIDGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDGte(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id >= ?", lessonID))
}

// LessonIDIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDIn(lessonID ...int64) RatingQuerySet {
	if len(lessonID) == 0 {
		qs.db.AddError(errors.New("must at least pass one lessonID in LessonIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("lesson_id IN (?)", lessonID))
}

// LessonIDLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDLt(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id < ?", lessonID))
}

// LessonIDLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDLte(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id <= ?", lessonID))
}

// LessonIDNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDNe(lessonID int64) RatingQuerySet {
	return qs.w(qs.db.Where("lesson_id != ?", lessonID))
}

// LessonIDNotIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) LessonIDNotIn(lessonID ...int64) RatingQuerySet {
	if len(lessonID) == 0 {
		qs.db.AddError(errors.New("must at least pass one lessonID in LessonIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("lesson_id NOT IN (?)", lessonID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) Limit(limit int) RatingQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) Offset(offset int) RatingQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs RatingQuerySet) One(ret *Rating) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByCreatedAt() RatingQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByDeletedAt() RatingQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByID() RatingQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByLessonID is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByLessonID() RatingQuerySet {
	return qs.w(qs.db.Order("lesson_id ASC"))
}

// OrderAscByRatingNumber is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByRatingNumber() RatingQuerySet {
	return qs.w(qs.db.Order("rating_number ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderAscByUpdatedAt() RatingQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByCreatedAt() RatingQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByDeletedAt() RatingQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByID() RatingQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByLessonID is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByLessonID() RatingQuerySet {
	return qs.w(qs.db.Order("lesson_id DESC"))
}

// OrderDescByRatingNumber is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByRatingNumber() RatingQuerySet {
	return qs.w(qs.db.Order("rating_number DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) OrderDescByUpdatedAt() RatingQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// RatingNumberEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberEq(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number = ?", ratingNumber))
}

// RatingNumberGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberGt(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number > ?", ratingNumber))
}

// RatingNumberGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberGte(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number >= ?", ratingNumber))
}

// RatingNumberIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberIn(ratingNumber ...int64) RatingQuerySet {
	if len(ratingNumber) == 0 {
		qs.db.AddError(errors.New("must at least pass one ratingNumber in RatingNumberIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating_number IN (?)", ratingNumber))
}

// RatingNumberLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberLt(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number < ?", ratingNumber))
}

// RatingNumberLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberLte(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number <= ?", ratingNumber))
}

// RatingNumberNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberNe(ratingNumber int64) RatingQuerySet {
	return qs.w(qs.db.Where("rating_number != ?", ratingNumber))
}

// RatingNumberNotIn is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) RatingNumberNotIn(ratingNumber ...int64) RatingQuerySet {
	if len(ratingNumber) == 0 {
		qs.db.AddError(errors.New("must at least pass one ratingNumber in RatingNumberNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating_number NOT IN (?)", ratingNumber))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtEq(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtGt(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtGte(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtLt(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtLte(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs RatingQuerySet) UpdatedAtNe(updatedAt time.Time) RatingQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetCreatedAt(createdAt time.Time) RatingUpdater {
	u.fields[string(RatingDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetDeletedAt(deletedAt *time.Time) RatingUpdater {
	u.fields[string(RatingDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetID(ID int64) RatingUpdater {
	u.fields[string(RatingDBSchema.ID)] = ID
	return u
}

// SetLessonID is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetLessonID(lessonID int64) RatingUpdater {
	u.fields[string(RatingDBSchema.LessonID)] = lessonID
	return u
}

// SetRatingNumber is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetRatingNumber(ratingNumber int64) RatingUpdater {
	u.fields[string(RatingDBSchema.RatingNumber)] = ratingNumber
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u RatingUpdater) SetUpdatedAt(updatedAt time.Time) RatingUpdater {
	u.fields[string(RatingDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u RatingUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u RatingUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set RatingQuerySet

// ===== BEGIN of Rating modifiers

// RatingDBSchemaField describes database schema field. It requires for method 'Update'
type RatingDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f RatingDBSchemaField) String() string {
	return string(f)
}

// RatingDBSchema stores db field names of Rating
var RatingDBSchema = struct {
	ID           RatingDBSchemaField
	LessonID     RatingDBSchemaField
	RatingNumber RatingDBSchemaField
	UpdatedAt    RatingDBSchemaField
	CreatedAt    RatingDBSchemaField
	DeletedAt    RatingDBSchemaField
}{

	ID:           RatingDBSchemaField("id"),
	LessonID:     RatingDBSchemaField("lesson_id"),
	RatingNumber: RatingDBSchemaField("rating_number"),
	UpdatedAt:    RatingDBSchemaField("updated_at"),
	CreatedAt:    RatingDBSchemaField("created_at"),
	DeletedAt:    RatingDBSchemaField("deleted_at"),
}

// Update updates Rating fields by primary key
// nolint: dupl
func (o *Rating) Update(db *gorm.DB, fields ...RatingDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":            o.ID,
		"lesson_id":     o.LessonID,
		"rating_number": o.RatingNumber,
		"updated_at":    o.UpdatedAt,
		"created_at":    o.CreatedAt,
		"deleted_at":    o.DeletedAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Rating %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// RatingUpdater is an Rating updates manager
type RatingUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewRatingUpdater creates new Rating updater
// nolint: dupl
func NewRatingUpdater(db *gorm.DB) RatingUpdater {
	return RatingUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Rating{}),
	}
}

// ===== END of Rating modifiers

// ===== END of all query sets
