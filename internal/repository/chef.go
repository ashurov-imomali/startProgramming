package repository

import (
	"main/pkg/models"
)

func (r *Repository) SelectUnfinishedOrders(page, limit int) ([]models.OrderItem, error) {
	sqlQuery := `select oi.id as menu_id, m.name, m.cost from orders_items oi join menu m on m.id = oi.menu_id
	join orders o on o.id = oi.order_id where oi.status = 1 order by o.created_at limit ? offset ?`
	orders := make([]models.OrderItem, 0)
	err := r.Db.Raw(sqlQuery, limit, (page-1)*limit).Scan(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *Repository) SetOrderItem(id, status int) (error, int) {
	sqlQuery := `update orders_items set status = ? where id = ? returning menu_id`
	var menuId int
	err := r.Db.Raw(sqlQuery, status, id).Scan(&menuId).Error
	if err != nil {
		return err, 0
	}
	return nil, menuId
}

func (r *Repository) AddChefSalary(userId, menuId int) error {
	sqlQuery := `update employers e set salary = salary + cost*0.05 from menu m where m.id = ? and e.id = ?`
	err := r.Db.Exec(sqlQuery, menuId, userId).Error
	if err != nil {
		return err
	}
	return nil
}
