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

	w := b.app.AsApiAction

	accGrp := group.Group("/:pid/account")
	report := group.Group("/:pid/report")
	inventry := group.Group("/:pid/inventory")
	txnGrp := group.Group("/:pid/txn")
	contactGrp := group.Group("/:pid/contacts")
	salesGrp := group.Group("/:pid/sales")
	taxGrp := group.Group("/:pid/tax")
	stocksGrp := group.Group("/:pid/stocks")
	estimateGrp := group.Group("/:pid/estimates")

	reportTpl := group.Group("/:pid/reportTemplate")
	reportSaved := group.Group("/:pid/reportSaved")
	notepad := group.Group("/:pid/notepad")

	// accounts
	accGrp.GET("/", w("listAccount", b.listAccount))
	accGrp.POST("/", w("addAccount", b.addAccount))
	accGrp.GET("/:id", w("getAccount", b.getAccount))
	accGrp.POST("/:id", w("updateAccount", b.updateAccount))
	accGrp.DELETE("/:id", w("deleteAccount", b.deleteAccount))

	// transactions
	txnGrp.GET("/", w("listTxn", b.listTxn))
	txnGrp.GET("/line/list", w("listTxnWithLines", b.listTxnWithLines))
	txnGrp.GET("/line/:aid/list", w("listAccountTxnWithLines", b.listAccountTxnWithLines))
	txnGrp.POST("/", w("addTxn", b.addTxn))
	txnGrp.GET("/:id", w("getTxn", b.getTxn))
	txnGrp.GET("/:id/line", w("getTxnWithLine", b.getTxnWithLine))
	txnGrp.POST("/:id", w("updateTxn", b.updateTxn))
	txnGrp.POST("/:id/line", w("updateTxnWithLine", b.updateTxnWithLine))
	txnGrp.DELETE("/:id", w("deleteTxn", b.deleteTxn))
	txnGrp.POST("/export", w("exportData", b.exportData))
	txnGrp.POST("/import", w("importData", b.importData))

	// catagories

	inventry.GET("/catagories", w("listCatagories", b.listCatagories))
	inventry.POST("/catagories", w("addCatagory", b.addCatagory))
	inventry.GET("/catagories/:id", w("getCatagory", b.getCatagory))
	inventry.POST("/catagories/:id", w("updateCatagory", b.updateCatagory))
	inventry.DELETE("/catagories/:id", w("deleteCatagory", b.deleteCatagory))

	// products
	inventry.GET("/products", w("listProducts", b.listProducts))
	inventry.POST("/products", w("addProduct", b.addProduct))
	inventry.GET("/products/:id", w("getProduct", b.getProduct))
	inventry.POST("/products/:id", w("updateProduct", b.updateProduct))
	inventry.DELETE("/products/:id", w("deleteProduct", b.deleteProduct))

	// contacts
	contactGrp.GET("/", w("listContacts", b.listContacts))
	contactGrp.POST("/", w("addContact", b.addContact))
	contactGrp.GET("/:id", w("getContact", b.getContact))
	contactGrp.POST("/:id", w("updateContact", b.updateContact))
	contactGrp.DELETE("/:id", w("deleteContact", b.deleteContact))

	// tax
	taxGrp.GET("/", w("listTax", b.listTax))
	taxGrp.POST("/", w("addTax", b.addTax))
	taxGrp.GET("/:id", w("getTax", b.getTax))
	taxGrp.POST("/:id", w("updateTax", b.updateTax))
	taxGrp.DELETE("/:id", w("deleteTax", b.deleteTax))
	taxGrp.POST("/:id/product", w("addTaxProduct", b.addTaxProduct))
	taxGrp.DELETE("/:id/product/:tpid", w("deleteTaxProduct", b.deleteTaxProduct))
	taxGrp.GET("/:id/product", w("listTaxProduct", b.listTaxProduct))

	// report

	report.POST("/live", w("reportLiveGenerate", b.reportLiveGenerate))

	// report template
	reportTpl.GET("/", w("reportTemplateList", b.reportTemplateList))
	reportTpl.POST("/", w("reportTemplateAdd", b.reportTemplateAdd))
	reportTpl.GET("/:id", w("reportTemplateGet", b.reportTemplateGet))
	reportTpl.POST("/:id", w("reportTemplateUpdate", b.reportTemplateUpdate))
	reportTpl.DELETE("/:id", w("reportTemplateDelete", b.reportTemplateDelete))

	// report saved
	reportSaved.GET("/", w("reportSavedList", b.reportSavedList))
	reportSaved.POST("/", w("reportSavedAdd", b.reportSavedAdd))
	reportSaved.GET("/:id", w("reportSavedGet", b.reportSavedGet))
	reportSaved.POST("/:id", w("reportSavedUpdate", b.reportSavedUpdate))
	reportSaved.DELETE("/:id", w("reportSavedDelete", b.reportSavedDelete))

	// sales
	salesGrp.GET("/", w("listSales", b.listSales))
	salesGrp.POST("/", w("addSales", b.addSales))
	salesGrp.GET("/:id", w("getSales", b.getSales))
	salesGrp.DELETE("/:id", w("deleteSales", b.deleteSales))

	// stocks

	stocksGrp.GET("/", w("listProductStockIn", b.listProductStockIn))
	stocksGrp.POST("/", w("addProductStockIn", b.addProductStockIn))
	stocksGrp.GET("/:id", w("getProductStockIn", b.getProductStockIn))
	stocksGrp.DELETE("/:id", w("deleteProductStockIn", b.deleteProductStockIn))

	// estimates
	estimateGrp.GET("/", w("listEstimate", b.listEstimate))
	estimateGrp.POST("/", w("addEstimate", b.addEstimate))
	estimateGrp.GET("/:id", w("getEstimate", b.getEstimate))
	estimateGrp.DELETE("/:id", w("deleteEstimate", b.deleteEstimate))

	// notepad

	notepad.GET("/", w("listNotepad", b.listNotepad))
	notepad.POST("/", w("addNotepad", b.addNotepad))
	notepad.GET("/:id", w("getNotepad", b.getNotepad))
	notepad.POST("/:id", w("updateNotepad", b.updateNotepad))
	notepad.DELETE("/:id", w("deleteNotepad", b.deleteNotepad))

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

// report template

func (b *BookModule) reportTemplateList(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.ReportTemplateList(pid, ctx.Claim.UserId, offset)
}

func (b *BookModule) reportTemplateAdd(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.ReportTemplate{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.ReportTemplateAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) reportTemplateGet(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.ReportTemplateGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) reportTemplateUpdate(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.ReportTemplateUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) reportTemplateDelete(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.ReportTemplateDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	return nil, err
}

// report saved

func (b *BookModule) reportSavedList(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.SavedTemplateList(pid, ctx.Claim.UserId, offset)
}

func (b *BookModule) reportSavedAdd(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.SavedReport{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.SavedTemplateAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) reportSavedGet(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.SavedTemplateGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) reportSavedUpdate(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.SavedTemplateUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) reportSavedDelete(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.SavedTemplateDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	return nil, err
}

// project

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

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.ProductList(pid, ctx.Claim.UserId, offset)
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

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.ContactList(pid, ctx.Claim.UserId, offset)
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

// estimates

func (b *BookModule) listEstimate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	userId := ctx.Claim.UserId
	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	list, err := b.dbOpts.EstimateList(pid, userId, offset)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (b *BookModule) addEstimate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	userId := ctx.Claim.UserId

	var req dbops.EstimateData

	if err := ctx.Http.Bind(&req); err != nil {
		return nil, err
	}

	id, err := b.dbOpts.EstimateAdd(pid, userId, &req)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (b *BookModule) getEstimate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	id := ctx.ParamInt64("id")
	userId := ctx.Claim.UserId

	return b.dbOpts.EstimateGet(pid, userId, id)
}

func (b *BookModule) deleteEstimate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	id := ctx.ParamInt64("id")
	userId := ctx.Claim.UserId
	err := b.dbOpts.EstimateDelete(pid, userId, id)
	return nil, err
}

// notepad

func (b *BookModule) listNotepad(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return b.dbOpts.NotepadList(pid, ctx.Claim.UserId, offset)
}

func (b *BookModule) addNotepad(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &models.Notepad{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpts.NotepadAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getNotepad(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpts.NotepadGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateNotepad(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpts.NotepadUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)

	return nil, err

}

func (b *BookModule) deleteNotepad(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpts.NotepadDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	return nil, err
}
