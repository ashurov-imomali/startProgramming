package models

import (
	"time"
)

type Employer struct {
	Id          int       `json:"id,omitempty"`
	RoleId      int       `json:"role_id,omitempty"`
	FullName    string    `json:"name,omitempty"gorm:"column:name"`
	Fines       float64   `json:"fines"`
	Allowances  float64   `json:"allowances"`
	Salary      float64   `json:"salary"`
	TotalSalary float64   `json:"total_salary"`
	Login       string    `json:"login,omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"-"`
	Active      bool      `json:"active,omitempty"`
}

type Token struct {
	Id             int       `json:"id,omitempty"`
	StrToken       string    `gorm:"column:token" json:"token"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at"`
	ExpirationTime time.Time `json:"expiration_time"`
	EmployerId     int       `json:"employer_id"`
}

type Menu struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cost       float64   `json:"cost"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"-"`
	Active     bool      `json:"active"`
	CategoryId int       `json:"category_id"`
}

type Table struct {
	Id          int    `json:"id"`
	TableNumber int    `json:"table_number"`
	ZoneId      int    `json:"zone_id"`
	Zone        string `json:"zone"gorm:"column:name"`
	Reserved    bool   `json:"reserved"`
}

type OrderItem struct {
	MenuId int     `json:"id"gorm:"column:menu_id"`
	Cost   float64 `json:"cost"`
	Name   string  `json:"name"`
}

type Order struct {
	Id             int       `json:"id"`
	WaiterId       int       `json:"waiter_id"`
	WaiterName     string    `json:"name"gorm:"column:name"`
	TableId        int       `json:"table_id"`
	OrderType      bool      `json:"order_type"`
	Cost           float64   `json:"cost"`
	TotalCost      float64   `json:"total_cost"`
	ServiceCharges int       `json:"service_charges"gorm:"column:percent"`
	CreatedAt      time.Time `json:"created_at"`
	Pay            bool      `json:"pay"`
}

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Zone struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
