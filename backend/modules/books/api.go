package books

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bornjre/turnix/backend/modules/books/dbops"
	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/bornjre/turnix/backend/xtypes"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (b *BookModule) register(group *gin.RouterGroup) error {

	x := b.app.AuthMiddleware

	accGrp := group.Group("/:pid/account")
	report := group.Group("/:pid/report")
	inventry := group.Group("/:pid/inventory")
	txnGrp := group.Group("/:pid/txn")
	contactGrp := group.Group("/:pid/contacts")
	salesGrp := group.Group("/:pid/sales")
	taxGrp := group.Group("/:pid/tax")
	stocksGrp := group.Group("/:pid/stocks")

	// accounts
	accGrp.GET("/", x(b.listAccount))
	accGrp.POST("/", x(b.addAccount))
	accGrp.GET("/:id", x(b.getAccount))
	accGrp.POST("/:id", x(b.updateAccount))
	accGrp.DELETE("/:id", x(b.deleteAccount))

	// transactions
	txnGrp.GET("/", x(b.listTxn))
	txnGrp.GET("/line/list", x(b.listTxnWithLines))
	txnGrp.GET("/line/:aid/list", x(b.listAccountTxnWithLines))
	txnGrp.POST("/", x(b.addTxn))
	txnGrp.GET("/:id", x(b.getTxn))
	txnGrp.GET("/:id/line", x(b.getTxnWithLine))
	txnGrp.POST("/:id", x(b.updateTxn))
	txnGrp.POST("/:id/line", x(b.updateTxnWithLine))
	txnGrp.DELETE("/:id", x(b.deleteTxn))
	txnGrp.POST("/export", x(b.exportData))
	txnGrp.POST("/import", x(b.importData))

	// catagories

	inventry.GET("/catagories", x(b.listCatagories))
	inventry.POST("/catagories", x(b.addCatagory))
	inventry.GET("/catagories/:id", x(b.getCatagory))
	inventry.POST("/catagories/:id", x(b.updateCatagory))
	inventry.DELETE("/catagories/:id", x(b.deleteCatagory))

	// products
	inventry.GET("/products", x(b.listProducts))
	inventry.POST("/products", x(b.addProduct))
	inventry.GET("/products/:id", x(b.getProduct))
	inventry.POST("/products/:id", x(b.updateProduct))
	inventry.DELETE("/products/:id", x(b.deleteProduct))

	// contacts
	contactGrp.GET("/", x(b.listContacts))
	contactGrp.POST("/", x(b.addContact))
	contactGrp.GET("/:id", x(b.getContact))
	contactGrp.POST("/:id", x(b.updateContact))
	contactGrp.DELETE("/:id", x(b.deleteContact))

	// tax
	taxGrp.GET("/", x(b.listTax))
	taxGrp.POST("/", x(b.addTax))
	taxGrp.GET("/:id", x(b.getTax))
	taxGrp.POST("/:id", x(b.updateTax))
	taxGrp.DELETE("/:id", x(b.deleteTax))
	taxGrp.POST("/:id/product", x(b.addTaxProduct))
	taxGrp.DELETE("/:id/product/:tpid", x(b.deleteTaxProduct))
	taxGrp.GET("/:id/product", x(b.listTaxProduct))

	// report

	report.POST("/live", x(b.reportLiveGenerate))

	// sales
	salesGrp.GET("/", x(b.listSales))
	salesGrp.POST("/", x(b.addSales))
	salesGrp.GET("/:id", x(b.getSales))
	salesGrp.DELETE("/:id", x(b.deleteSales))

	// stocks

	stocksGrp.GET("/", x(b.listProductStockIn))
	stocksGrp.POST("/", x(b.addProductStockIn))
	stocksGrp.GET("/:id", x(b.getProductStockIn))
	stocksGrp.DELETE("/:id", x(b.deleteProductStockIn))

	return nil
}

// account

func (b *BookModule) listAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ListAccount(pid, ctx.Claim.UserId)
}

func (b *BookModule) addAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Account{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.AddAccount(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.GetAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateAccount(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.UpdateAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) deleteAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.DeleteAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

// txn

func (b *BookModule) listTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ListTxn(pid, ctx.Claim.UserId)
}

func (b *BookModule) listTxnWithLines(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.ListTxnWithLines(pid, ctx.Claim.UserId, offset)
}

func (b *BookModule) listAccountTxnWithLines(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)
	account, _ := strconv.ParseInt(ctx.Http.Param("aid"), 10, 64)

	return b.dbOpts.ListAccountTxnWithLines(pid, ctx.Claim.UserId, account, offset)
}

func (b *BookModule) addTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.TxnRecord{}

	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	var buf strings.Builder

	for key := range data.Attachments {
		buf.WriteString(key)
		buf.WriteByte(',')
	}

	// fixme store files

	txnid, err := b.dbOpts.AddTxn(pid, ctx.Claim.UserId, &models.Transaction{
		Title:       data.Title,
		Notes:       data.Notes,
		ReferenceID: data.ReferenceID,
		Attachments: buf.String(),
		CreatedBy:   ctx.Claim.UserId,
		UpdatedBy:   ctx.Claim.UserId,
	})

	if err != nil {
		return nil, err
	}

	did, err := b.dbOpts.AddTxnLine(pid, ctx.Claim.UserId, &models.TransactionLine{
		AccountID:    data.CreditAccountID,
		DebitAmount:  0,
		CreditAmount: data.CreditAmount,
		TxnID:        txnid,
		CreatedBy:    ctx.Claim.UserId,
		UpdatedBy:    ctx.Claim.UserId,
	})

	if err != nil {
		return nil, err
	}

	cid, err := b.dbOpts.AddTxnLine(pid, ctx.Claim.UserId, &models.TransactionLine{
		AccountID:    data.DebitAccountID,
		DebitAmount:  data.DebitAmount,
		CreditAmount: 0,
		TxnID:        txnid,
		CreatedBy:    ctx.Claim.UserId,
		UpdatedBy:    ctx.Claim.UserId,
	})
	if err != nil {
		return nil, err
	}

	return models.TxnResult{
		TxnId:        txnid,
		DebitLineId:  did,
		CreditLineId: cid,
	}, nil

}

func (b *BookModule) getTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.GetTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) getTxnWithLine(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.GetTxnWithLine(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateTxn(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.UpdateTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

type UpdateTxnWithLineOptions struct {
	TxnData        map[string]any `json:"txn_data"`
	FirstLineId    int64          `json:"first_line_id"`
	FirstLineData  map[string]any `json:"first_line_data"`
	SecondLineId   int64          `json:"second_line_id"`
	SecondLineData map[string]any `json:"second_line_data"`
}

func (b *BookModule) updateTxnWithLine(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := &UpdateTxnWithLineOptions{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	pp.Println("@data", data)

	txnId := ctx.ParamInt64("id")

	if len(data.TxnData) != 0 {
		err = b.dbOpts.UpdateTxn(pid, ctx.Claim.UserId, txnId, data.TxnData)
		if err != nil {
			return nil, err
		}
	}

	if len(data.FirstLineData) != 0 {
		err = b.dbOpts.UpdateTxnLine(pid, ctx.Claim.UserId, txnId, data.FirstLineId, data.FirstLineData)
		if err != nil {
			return nil, err
		}

	}

	if len(data.SecondLineData) != 0 {
		err = b.dbOpts.UpdateTxnLine(pid, ctx.Claim.UserId, txnId, data.SecondLineId, data.SecondLineData)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil

}

func (b *BookModule) deleteTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.DeleteTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

func (b *BookModule) reportLiveGenerate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	opts := dbops.ReportOptions{}
	err := ctx.Http.Bind(&opts)
	if err != nil {
		return nil, err
	}

	switch opts.ReportType {
	case models.ReportTypeLongLedger:
		return b.dbOpts.GenerateReportLongLedger(pid, ctx.Claim.UserId, opts)
	case models.ReportTypeShortLedger:
		return b.dbOpts.GenerateReportShortLedger(pid, ctx.Claim.UserId, opts)
	case models.ReportTypeCustom:
		fallthrough

	default:
		panic(fmt.Sprintf("Report type %s not implemented", opts.ReportType))
	}

}

func (b *BookModule) importData(ctx xtypes.ContextPlus) (any, error) {

	pp.Println("@importData")

	opts := dbops.ImportOptions{}
	err := ctx.Http.Bind(&opts)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.Import(ctx.ProjectId(), ctx.Claim.UserId, opts)

	return nil, err

}

func (b *BookModule) exportData(ctx xtypes.ContextPlus) (any, error) {
	pp.Println("@exportData")

	return b.dbOpts.Export(ctx.ProjectId(), ctx.Claim.UserId)
}

func (b *BookModule) listCatagories(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.CatagoryList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Catagory{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.CatagoryAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.CatagoryGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateCatagory(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.CatagoryUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err
}

func (b *BookModule) deleteCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.CatagoryDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	return nil, err

}

// products

func (b *BookModule) listProducts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ProductList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Product{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.ProductAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.ProductGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateProduct(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.ProductUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	return nil, err
}

func (b *BookModule) deleteProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.ProductDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

// contacts

func (b *BookModule) listContacts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ContactList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Contact{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.ContactAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.ContactGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateContact(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.ContactUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) deleteContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.ContactDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

// sales

func (b *BookModule) listSales(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.SalesList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addSales(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &dbops.SalesData{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.SalesAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getSales(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.SalesGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) deleteSales(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.SalesDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

// tax

func (b *BookModule) listTax(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.TaxList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addTax(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Tax{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.TaxAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getTax(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.TaxGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateTax(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	pp.Println("@updateTax", data)

	err = b.dbOpts.TaxUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) deleteTax(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.TaxDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

func (b *BookModule) listTaxProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ProductTaxList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addTaxProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.ProductTax{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.ProductTaxAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) deleteTaxProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.ProductTaxDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}

// stocks

func (b *BookModule) listProductStockIn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpts.ProductStockInList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addProductStockIn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &dbops.StockInData{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.ProductStockInAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getProductStockIn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.ProductStockInGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) deleteProductStockIn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.ProductStockInDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))

	return nil, err
}
