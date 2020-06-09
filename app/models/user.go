package models

import "github.com/Gogotchuri/GoFast/database"

//Default interface methods forcefully
var _ DeleterSaver = &User{}

/*User is a base model of user*/
type User struct {
	BaseModel

	GoogleID        string `gorm:"unique_index;default:null;"`
	FacebookID      string `gorm:"unique_index;default:null;"`
	Password        string `gorm:"default:null;"`
	FirstName       string `gorm:"not null"`
	LastName        string `gorm:"not null"`
	Email           string `gorm:"type:varchar(100);unique_index;not null;"`
	EmailVerifiedAt TimestampT
	Role            uint `gorm:"default:1;"`
	//TODO setup relationships here
}

/*GetUserByID returns user with given ID or null*/
func GetUserByID(ID uint) *User {
	var u User
	database.GetInstance().First(&u, ID)
	if u.ID == ID {
		return &u
	}
	return nil
}

/*Delete deletes user, softDelete parameter specifies whether value should be deleted permanently
 *Primary key (ID) must be set
 */
func (u *User) Delete(softDelete bool) {
	db := database.GetInstance()
	if !softDelete {
		db = db.Unscoped()
	}
	db.Where("ID = ?", u.ID).Delete(u)
}

/*Save updates model in the database, if primary key isn't set creates model*/
func (u *User) Save() {
	database.GetInstance().Save(u)
}

/*SaveUserIfDoesntExist saves user to database if emails isn't occupied*/
func (u *User) SaveUserIfDoesntExist() {
	if existing := GetUserByEmail(u.Email); existing == nil {
		u.Save() //Write user to database if doesn't exist
	}
}

/*GetUserByEmail returns user with given email if exists, otherwise null*/
func GetUserByEmail(email string) *User {
	var u User
	database.GetInstance().Where("email = ?", email).First(&u)
	if u.Email == email {
		return &u
	}
	return nil
}

/*CheckCredentials returns user if email-password combination exists, otherwise null*/
func CheckCredentials(email, password string) *User {
	u := GetUserByEmail(email)
	// TODO: Change after adding encryption
	if u != nil && u.Password == password {
		return u
	}
	return nil
}
