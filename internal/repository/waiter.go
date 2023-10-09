package repository

import (
	"errors"
	"main/pkg/models"
)

func (r *Repository) SelectTables(zone string, page, limit int) ([]models.Table, error) {
	sqlQuery := `select hm.id, hm.table_number, z.name, hm.reserved, z.id as zone_id from hall_map hm 
    join zones z on hm.zone_id = z.id where lower(name) like concat(concat('%',lower(?)),'%') and not (hm.reserved) 
	order by hm.id limit ? offset ?`
	tables := make([]models.Table, 0)
	err := r.Db.Raw(sqlQuery, "%"+zone+"%", limit, (page-1)*limit).Scan(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *Repository) SetTable(id int, ok bool) error {
	sqlQuery := `update hall_map set reserved = ? where id = ?`
	err := r.Db.Exec(sqlQuery, ok, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectCategories(categories string, page, limit int) ([]models.Categories, error) {
	///Можно заменить на ||
	sqlQuery := `select * from menu_categories where lower(name) like concat(concat('%', lower(?)), '%') limit ? offset ?`
	menuCategories := make([]models.Categories, 0)
	err := r.Db.Raw(sqlQuery, categories, limit, (page-1)*limit).Scan(&menuCategories).Error
	if err != nil {
		return nil, err
	}
	return menuCategories, nil
}

func (r *Repository) SelectMenu(dishesName string, page, limit int) ([]models.Menu, error) {
	sqlQuery := `select * from menu where lower(name) like concat(concat('%',lower(?)), '%') and active = true
	limit ? offset ?`
	menu := make([]models.Menu, 0)

	err := r.Db.Raw(sqlQuery, dishesName, limit, (page-1)*limit).Scan(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *Repository) InsertOrder(order *models.Order) (error, int) {
	sqlQuery := `insert into orders (waiter_id, type_order, place, cost, total_cost,service_charge)
values (?,?,?,?,?,?) returning id`
	var orderId int
	err := r.Db.Raw(sqlQuery, order.WaiterId, order.OrderType, order.TableId, order.Cost,
		order.TotalCost, order.ServiceCharges).Scan(&orderId).Error
	if err != nil {
		return err, 0
	}
	return nil, orderId
}

func (r *Repository) InsertOrderItems(orderId int, orderItems []models.OrderItem) error {
	sqlQuery := `insert into orders_items (order_id, menu_id, status)
values (?,?,?)`
	for _, item := range orderItems {
		err := r.Db.Exec(sqlQuery, orderId, item.MenuId, 1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) SelectTable(id int) (*models.Table, error) {
	sqlQuery := `select * from hall_map where id = ?`
	var table models.Table
	err := r.Db.Raw(sqlQuery, id).Scan(&table).Error
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *Repository) GetCost(orderItems []models.OrderItem) (float64, error) {
	sqlQuery := `select cost from menu where id = ?`
	var sum float64 = 0
	for _, order := range orderItems {
		var cost float64
		err := r.Db.Raw(sqlQuery, order.MenuId).Scan(&cost).Error
		if err != nil {
			return 0, err
		}
		sum += cost
	}
	return sum, nil
}

func (r *Repository) SelectOrderItems(id, page, limit int) ([]models.OrderItem, error) {
	sqlQuery := `select menu_id, m.name, m.cost from menu m join orders_items oi on m.id = oi.menu_id join orders o
    on o.id = oi.order_id where o.id = ? and pay = false limit ? offset ?`
	orderItems := make([]models.OrderItem, 0)
	err := r.Db.Raw(sqlQuery, id, limit, (page-1)*limit).Scan(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *Repository) SelectOrder(id int) (*models.Order, error) {
	sqlQuery := `select o.created_at, e.name, o.cost, sc.percent, o.total_cost from 
    orders o, employers e, service_charges sc where e.id = o.waiter_id and o.service_charge = sc.id and o.id = ? and o.pay = false`
	var order models.Order
	err := r.Db.Raw(sqlQuery, id).Scan(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) SetOrder(userId, id int) error {
	sqlQuery := `update orders set pay = true where id = ? and waiter_id = ?`
	err := r.Db.Exec(sqlQuery, id, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddWaiterSalary(userId, id int) error {
	sqlQuery := `update employers e set salary = salary + total_cost*0.08 from orders o where o.id = ? and e.id = ?`
	err := r.Db.Exec(sqlQuery, id, userId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckPay(id int) (error, int) {
	sqlQuery := `select pay, waiter_id from orders where id = ? and pay = false`
	var order models.Order
	order.Pay = true
	err := r.Db.Raw(sqlQuery, id).Scan(&order).Error
	if err != nil {
		return err, 0
	}
	if order.Pay {
		return errors.New("Check Closet yet"), 0
	}
	return nil, order.WaiterId
}
