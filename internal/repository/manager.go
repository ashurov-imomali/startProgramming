package repository

import (
	"errors"
	"main/pkg/models"
	"time"
)

func (r *Repository) CheckCurrentTable(table *models.Table) error {
	sqlQuery := `select * from hall_map where zone_id = ? and table_number = ?`
	var currentTable models.Table
	err := r.Db.Raw(sqlQuery, table.ZoneId, table.TableNumber).Scan(&currentTable).Error
	if err != nil {
		return err
	}
	if currentTable.ZoneId != 0 {
		return errors.New("current table")
	}
	return nil
}

func (r *Repository) InsertTable(table *models.Table) error {
	sqlQuery := `insert into hall_map(table_number, zone_id) values (?,?)`
	err := r.Db.Exec(sqlQuery, table.TableNumber, table.ZoneId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckCurrentZone(zone *models.Zone) error {
	sqlQuery := `select * from zones where name = ?`
	var currentZone models.Zone
	err := r.Db.Raw(sqlQuery, zone.Name).Scan(&currentZone).Error
	if err != nil {
		return err
	}
	if currentZone.Name == zone.Name {
		return errors.New("current zone")
	}
	return nil
}

func (r *Repository) InsertZone(zone *models.Zone) error {
	sqlQuery := `insert into zones(name) values (?)`
	err := r.Db.Exec(sqlQuery, zone.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckZone(table *models.Table) error {
	sqlQuery := `select id from zones where id = ?`
	var zoneId = 0
	err := r.Db.Raw(sqlQuery, table.ZoneId).Scan(&zoneId).Error
	if err != nil {
		return err
	}
	if zoneId == 0 {
		return errors.New("incorrect zone Id")
	}
	return nil
}

func (r *Repository) CheckCurrentCategory(categories *models.Categories) error {
	sqlQuery := `select * from menu_categories where name = ?`
	var currentCategory models.Categories
	err := r.Db.Raw(sqlQuery, categories.Name).Scan(&currentCategory).Error
	if err != nil {
		return err
	}
	if currentCategory.Id != 0 {
		return errors.New("current category:(")
	}
	return nil
}

func (r *Repository) InsertCategory(category *models.Categories) error {
	sqlQuery := `insert into menu_categories (name) values (?)`
	err := r.Db.Exec(sqlQuery, category.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckCurrentMenuItem(menuItem *models.Menu) error {
	sqlQuery := `select * from menu where category_id = ? and name = ?`
	var currentMenuItem models.Menu
	err := r.Db.Raw(sqlQuery, menuItem.CategoryId, menuItem.Name).Scan(&currentMenuItem).Error
	if err != nil {
		return err
	}
	if currentMenuItem.Id != 0 {
		return errors.New("current menuItem")
	}
	return nil
}

func (r *Repository) CheckCategory(menuItem *models.Menu) error {
	sqlQuery := `select id from menu_categories where id = ?`
	var categoryId = 0
	err := r.Db.Raw(sqlQuery, menuItem.CategoryId).Scan(&categoryId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertMenuItem(menuItem *models.Menu) error {
	sqlQuery := `insert into menu (category_id, name, cost) values (?, ?, ?)`
	err := r.Db.Exec(sqlQuery, menuItem.CategoryId, menuItem.Name, menuItem.Cost).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetMenuItem(menuItem *models.Menu) error {
	sqlQuery := `update menu set cost = ?, updated_at = ? where id = ?`
	err := r.Db.Exec(sqlQuery, menuItem.Cost, time.Now(), menuItem.Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectEmployer(userId int) (*models.Employer, error) {
	sqlQuery := `select id,salary, fines, allowances, name, created_at, updated_at,salary+allowances-fines as total_salary
	from employers where id = ?`
	var employer models.Employer
	err := r.Db.Raw(sqlQuery, userId).Scan(&employer).Error
	if err != nil {
		return nil, err
	}
	return &employer, nil
}

func (r *Repository) InsertAllowance(userId int, allowance float64) error {
	sqlQuery := `update employers set allowances = allowances+?, updated_at = ? where id = ?`
	err := r.Db.Exec(sqlQuery, allowance, time.Now(), userId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertFines(userId int, fines float64) error {
	sqlQuery := `update employers set fines = fines + ?, updated_at = ? where id = ?`
	err := r.Db.Exec(sqlQuery, fines, time.Now(), userId).Error
	if err != nil {
		return err
	}
	return nil
}
