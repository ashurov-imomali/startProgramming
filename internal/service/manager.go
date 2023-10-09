package service

import (
	"main/pkg/models"
)

func (srv *Service) CreateTable(table *models.Table) error {
	err := srv.Repository.InsertTable(table)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CheckNewTable(table *models.Table) error {
	err := srv.Repository.CheckZone(table)
	if err != nil {
		return err
	}
	err = srv.Repository.CheckCurrentTable(table)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CheckNewZone(zone *models.Zone) error {
	err := srv.Repository.CheckCurrentZone(zone)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) AddNewZone(zone *models.Zone) error {
	err := srv.Repository.InsertZone(zone)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CheckNewCategory(category *models.Categories) error {
	err := srv.Repository.CheckCurrentCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CreateNewCategory(category *models.Categories) error {
	err := srv.Repository.InsertCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CheckNewMenuItem(menuItem *models.Menu) error {
	err := srv.Repository.CheckCategory(menuItem)
	if err != nil {
		return err
	}
	err = srv.Repository.CheckCurrentMenuItem(menuItem)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CreateNewMenuItem(menuItem *models.Menu) error {
	err := srv.Repository.InsertMenuItem(menuItem)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) UpdateMenuItem(menuItem *models.Menu) error {
	err := srv.Repository.SetMenuItem(menuItem)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) GetEmployer(userId int) (*models.Employer, error) {
	employer, err := srv.Repository.SelectEmployer(userId)
	if err != nil {
		return nil, err
	}
	return employer, nil
}

func (srv *Service) GiveAllowance(userId int, allowance float64) error {
	err := srv.Repository.InsertAllowance(userId, allowance)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) GiveFines(userId int, fines float64) error {
	err := srv.Repository.InsertFines(userId, fines)
	if err != nil {
		return err
	}
	return nil
}
