package models

import "time"

type User struct {
	ID                     string     `json:"id" db:"id" valid:"required,uuid"`
	Username               string     `json:"username" db:"username" valid:"required,stringlength(5|50),matches(^[a-zA-Z0-9_]+$)"`
	CodeStudent            string     `json:"code_student" db:"code_student" valid:"required,stringlength(5|50),matches(^[a-zA-Z0-9_]+$)"`
	Dni                    string     `json:"dni" db:"dni" valid:"required,stringlength(5|50),matches(^[a-zA-Z0-9_]+$)"`
	Names                  string     `json:"names,omitempty" db:"names" valid:"required,stringlength(0|255)"`
	LastNameFather         string     `json:"lastname_father,omitempty" db:"lastname_father" valid:"required,stringlength(0|255)"`
	LastNameMother         string     `json:"lastname_mother,omitempty" db:"lastname_mother" valid:"required,stringlength(0|255)"`
	Email                  string     `json:"email,omitempty" db:"email" valid:"required,email,stringlength(5|255)"`
	Password               string     `json:"password,omitempty" db:"password" valid:"-"`
	EmailNotifications     string     `json:"email_notifications,omitempty" db:"email_notifications" valid:"required,email,stringlength(5|255)"`
	IdentificationNumber   string     `json:"identification_number,omitempty" db:"identification_number" valid:"required,stringlength(0|255)"`
	IdentificationType     string     `json:"identification_type,omitempty" db:"identification_type" valid:"required"`
	Status                 int        `json:"status,omitempty" db:"status" valid:"-"`
	FailedAttempts         int        `json:"failed_attempts,omitempty" db:"failed_attempts" valid:"-"`
	LastChangePassword     *time.Time `json:"last_change_password,omitempty" db:"last_change_password" valid:"-"`
	BlockDate              *time.Time `json:"block_date,omitempty" db:"block_date" valid:"-"`
	DisabledDate           *time.Time `json:"disabled_date,omitempty" db:"disabled_date" valid:"-"`
	ChangePassword         *bool      `json:"change_password,omitempty" db:"change_password" valid:"-"`
	ChangePasswordDaysLeft *int       `json:"change_password_days_left,omitempty" db:"change_password_days_left" valid:"-"`
	IsBlock                *bool      `json:"is_block,omitempty" db:"is_block" valid:"-"`
	IsDisabled             *bool      `json:"is_disabled,omitempty" db:"is_disabled" valid:"-"`
	LastLogin              *time.Time `json:"last_login,omitempty" db:"last_login" valid:"-"`
	TimeOut                int        `json:"time_out,omitempty" valid:"-"`
	ClientID               int        `json:"client_id,omitempty" bson:"client_id"`
	HostName               string     `json:"host_name,omitempty" bson:"host_name"`
	RealIP                 string     `json:"real_ip,omitempty" bson:"real_ip"`
	Token                  string     `json:"token,omitempty" bson:"token"`
	SessionID              string     `json:"session_id" bson:"session_id"`
	Colors                 Color      `json:"colors,omitempty" bson:"colors"`
	Roles                  []*string  `json:"roles,omitempty" bson:"roles"`
	DocTypes               []*int     `json:"doc_types,omitempty" bson:"doc_types"`
	Projects               []*string  `json:"projects,omitempty" bson:"projects"`
	CreatedAt              time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	UserId                 string     `json:"user_id,omitempty" db:"user_id"`
}

type Color struct {
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
	Tertiary  string `json:"tertiary" bson:"tertiary"`
}

type ResAuthentication struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	UserId      string `json:"userId"`
}
