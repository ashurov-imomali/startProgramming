package service

import (
	"errors"
	"github.com/jung-kurt/gofpdf"
	"main/pkg/models"
	"strconv"
)

func (srv *Service) FindTables(zone string, page, limit int) ([]models.Table, error) {
	tables, err := srv.Repository.SelectTables(zone, page, limit)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (srv *Service) BookingTheTable(id int) error {
	err := srv.Repository.SetTable(id, true)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CancelReservation(id int) error {
	err := srv.Repository.SetTable(id, false)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) ReturnCategories(categories string, page, limit int) ([]models.Categories, error) {
	menuCategories, err := srv.Repository.SelectCategories(categories, page, limit)
	if err != nil {
		return nil, err
	}
	return menuCategories, nil
}

func (srv *Service) ReturnMenu(dishesName string, page, limit int) ([]models.Menu, error) {
	menu, err := srv.Repository.SelectMenu(dishesName, page, limit)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (srv *Service) AddOrder(order *models.Order, orderItems []models.OrderItem) (int, error) {
	cost, err := srv.Repository.GetCost(orderItems)
	if err != nil {
		return 0, err
	}
	order.Cost = cost
	if order.OrderType {
		order.TotalCost = cost
		order.ServiceCharges = 2
	} else {
		totalCost := srv.GetTotalCost(cost)
		order.TotalCost = totalCost
		order.ServiceCharges = 1
	}
	err, orderId := srv.Repository.InsertOrder(order)
	if err != nil {
		return 0, err
	}
	err = srv.BookingTheTable(order.TableId)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}

func (srv *Service) AddOrderItems(orderId int, orderItems []models.OrderItem) error {
	err := srv.Repository.InsertOrderItems(orderId, orderItems)
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) CheckTable(id int) error {
	table, err := srv.Repository.SelectTable(id)
	if err != nil {
		return err
	}
	if table.Reserved == true {
		return errors.New("Table is reserved :(")
	}
	return nil
}

func (srv *Service) ReturnOrder(id, page, limit int) ([]models.OrderItem, error) {
	orderItems, err := srv.Repository.SelectOrderItems(id, page, limit)
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (srv *Service) GetTotalCost(cost float64) float64 {
	return cost + cost*0.1
}

func (srv *Service) CreateCheck(id int) error {
	orderItems, err := srv.Repository.SelectOrderItems(id, 1, 100)
	if err != nil {
		return err
	}
	order, err := srv.Repository.SelectOrder(id)
	if err != nil {
		return err
	}
	file := srv.GetPdfFile(orderItems, order)
	err = file.OutputFileAndClose("output.pdf")
	if err != nil {
		return err
	}
	return nil
}

func (srv *Service) GetPdfFile(orderItems []models.OrderItem, order *models.Order) *gofpdf.Fpdf {
	pdfFile := gofpdf.New("P", "mm", "", "")
	height := len(orderItems) + 7
	flHeight := float64(height) * 16
	newPage := gofpdf.SizeType{Wd: 120, Ht: flHeight}
	pdfFile.AddPageFormat("P", newPage)
	pdfFile.SetFont("Arial", "B", 16)
	pdfFile.Cell(40, 10, "===============ILF===============")
	pdfFile.Ln(10)
	pdfFile.SetFont("Arial", "", 14)
	pdfFile.Cell(40, 10, "order time: ")
	pdfFile.Cell(20, 10, "")
	strTime := order.CreatedAt.Format("2006-01-02 15:04:05")
	pdfFile.Cell(40, 10, strTime)
	pdfFile.Ln(10)
	pdfFile.SetFont("Arial", "", 14)
	pdfFile.Cell(40, 10, "waiter: ")
	pdfFile.Cell(40, 10, "")
	pdfFile.Cell(40, 10, order.WaiterName)
	pdfFile.Ln(10)
	pdfFile.Cell(40, 10, "===================================================")
	pdfFile.Ln(10)
	pdfFile.SetFont("Arial", "", 16)
	for _, item := range orderItems {
		pdfFile.Cell(40, 10, item.Name)
		strCost := strconv.FormatFloat(item.Cost, 'f', 2, 64)
		pdfFile.Cell(40, 10, "")
		pdfFile.Cell(40, 10, strCost)
		pdfFile.Ln(10)
	}
	pdfFile.Cell(40, 10, "===================================================")
	pdfFile.Ln(10)
	pdfFile.Cell(40, 10, "cost: ")
	pdfFile.Cell(40, 10, "")
	strCost := strconv.FormatFloat(order.Cost, 'f', 2, 64)
	pdfFile.Cell(40, 10, strCost)
	pdfFile.Ln(10)
	pdfFile.Cell(40, 10, "service charge: ")
	pdfFile.Cell(40, 10, "")
	strServiceCharges := strconv.Itoa(order.ServiceCharges)
	strServiceCharges += "%"
	pdfFile.Cell(40, 10, strServiceCharges)
	pdfFile.Ln(10)
	pdfFile.SetFont("Arial", "B", 16)
	strTotalCost := strconv.FormatFloat(order.TotalCost, 'f', 2, 64)
	pdfFile.Cell(40, 10, "total cost: ")
	pdfFile.Cell(40, 10, "")
	pdfFile.Cell(40, 10, strTotalCost)
	return pdfFile
}

func (srv *Service) CheckOrder(id int) (error, int) {
	err, userId := srv.Repository.CheckPay(id)
	if err != nil {
		return err, 0
	}
	return nil, userId
}

func (srv *Service) CloseCheck(userId, id int) error {
	err := srv.Repository.SetOrder(userId, id)
	if err != nil {
		return err
	}
	err = srv.Repository.AddWaiterSalary(userId, id)
	if err != nil {
		return err
	}
	return nil
}
