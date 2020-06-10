package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*TimestampT for timestamp references in database*/
type TimestampT time.Time

/*DateT for date references in database*/
type DateT time.Time

/*BaseModel is base class for all models, created for flexibility later,
 * currently same as gorm.Model
 */
type BaseModel struct {
	gorm.Model
}

/*DeleterSaver interface to define methods for base models with ID*/
type DeleterSaver interface {
	Delete(bool)
	Save()
}
