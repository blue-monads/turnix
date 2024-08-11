package books

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	report.POST("/live", x(b.reportLiveGenerate))

	return nil
}

// account

func (b *BookModule) listAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpListAccount(pid, ctx.Claim.UserId)
}

func (b *BookModule) addAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &Account{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpAddAccount(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpGetAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateAccount(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpUpdateAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (b *BookModule) deleteAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpDeleteAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// txn

func (b *BookModule) listTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpListTxn(pid, ctx.Claim.UserId)
}

func (b *BookModule) listTxnWithLines(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Param("offset"), 10, 64)

	return b.dbOpListTxnWithLines(pid, ctx.Claim.UserId, offset)
}

func (b *BookModule) listAccountTxnWithLines(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)
	account, _ := strconv.ParseInt(ctx.Http.Param("aid"), 10, 64)

	return b.dbOpListAccountTxnWithLines(pid, ctx.Claim.UserId, account, offset)
}

func (b *BookModule) addTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &TxnRecord{}

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

	txnid, err := b.dbOpAddTxn(pid, ctx.Claim.UserId, &Transaction{
		Title:           data.Title,
		Notes:           data.Notes,
		LinkedSalesID:   data.LinkedSalesID,
		LinkedInvoiceID: data.LinkedSalesID,
		ReferenceID:     data.ReferenceID,
		Attachments:     buf.String(),
		CreatedBy:       ctx.Claim.UserId,
		UpdatedBy:       ctx.Claim.UserId,
	})

	if err != nil {
		return nil, err
	}

	did, err := b.dbOpAddTxnLine(pid, ctx.Claim.UserId, &TransactionLine{
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

	cid, err := b.dbOpAddTxnLine(pid, ctx.Claim.UserId, &TransactionLine{
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

	return TxnResult{
		TxnId:        txnid,
		DebitLineId:  did,
		CreditLineId: cid,
	}, nil

}

func (b *BookModule) getTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	pp.Println("ctx.Http.FullPath()", ctx.Http.FullPath())
	pp.Println("ctx.Http.Request.URL.Path()", ctx.Http.Request.URL.Path)

	return b.dbOpGetTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) getTxnWithLine(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpGetTxnWithLine(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateTxn(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpUpdateTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

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
		err = b.dbOpUpdateTxn(pid, ctx.Claim.UserId, txnId, data.TxnData)
		if err != nil {
			return nil, err
		}
	}

	if len(data.FirstLineData) != 0 {
		err = b.dbOpUpdateTxnLine(pid, ctx.Claim.UserId, txnId, data.FirstLineId, data.FirstLineData)
		if err != nil {
			return nil, err
		}

	}

	if len(data.SecondLineData) != 0 {
		err = b.dbOpUpdateTxnLine(pid, ctx.Claim.UserId, txnId, data.SecondLineId, data.SecondLineData)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil

}

func (b *BookModule) deleteTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpDeleteTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type ReportOptions struct {
	ReportType string     `json:"report_type"`
	TemplateId int64      `json:"template_id"`
	FromDate   *time.Time `json:"from_date"`
	ToDate     *time.Time `json:"to_date"`
}

func (b *BookModule) reportLiveGenerate(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	opts := ReportOptions{}
	err := ctx.Http.Bind(&opts)
	if err != nil {
		return nil, err
	}

	switch opts.ReportType {
	case ReportTypeLongLedger:
		return b.dbOptsGenerateReportLongLedger(pid, ctx.Claim.UserId, opts)
	case ReportTypeShortLedger:
		return b.dbOptsGenerateReportShortLedger(pid, ctx.Claim.UserId, opts)
	case ReportTypeCustom:
		fallthrough

	default:
		panic(fmt.Sprintf("Report type %s not implemented", opts.ReportType))
	}

}

type ExportData struct {
	Accounts         []Account         `json:"accounts"`
	Transactions     []Transaction     `json:"transactions"`
	TransactionLines []TransactionLine `json:"transaction_lines"`
}

type ImportOptions struct {
	Data     ExportData `json:"data"`
	AsNewTxn bool       `json:"as_new_txn"`
}

func (b *BookModule) importData(ctx xtypes.ContextPlus) (any, error) {

	pp.Println("@importData")

	opts := ImportOptions{}
	err := ctx.Http.Bind(&opts)
	if err != nil {
		return nil, err
	}

	err = b.dbOpsImport(ctx.ProjectId(), ctx.Claim.UserId, opts)

	return nil, err

}

func (b *BookModule) exportData(ctx xtypes.ContextPlus) (any, error) {
	pp.Println("@exportData")

	return b.dbOpsExport(ctx.ProjectId(), ctx.Claim.UserId)
}

func (b *BookModule) listCatagories(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpsCatagoryList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &Catagory{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpsCatagoryAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpsCatagoryGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateCatagory(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpsCatagoryUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (b *BookModule) deleteCatagory(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpsCatagoryDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// products

func (b *BookModule) listProducts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpsProductList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &Product{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpsProductAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpsProductGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateProduct(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpsProductUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (b *BookModule) deleteProduct(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpsProductDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// contacts

func (b *BookModule) listContacts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpsContactList(pid, ctx.Claim.UserId)
}

func (b *BookModule) addContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &Contact{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpsContactAdd(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpsContactGet(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateContact(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpsContactUpdate(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (b *BookModule) deleteContact(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	err := b.dbOpsContactDelete(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
