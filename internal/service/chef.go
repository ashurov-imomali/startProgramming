package service

import "main/pkg/models"

func (srv *Service) ReturnUnfinishedOrders(page, limit int) ([]models.OrderItem, error) {
	orders, err := srv.Repository.SelectUnfinishedOrders(page, limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (srv *Service) PatchOrderItem(id int) error {
	err, _ := srv.Repository.SetOrderItem(id, 2)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) FinishOrder(id, userId int) error {
	err, menuId := srv.Repository.SetOrderItem(id, 3)
	if err != nil {
		return err
	}
	err = srv.Repository.AddChefSalary(userId, menuId)
	if err != nil {
		return err
	}
	return nil
}
