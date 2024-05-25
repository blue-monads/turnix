package books

import (
	"strconv"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/gin-gonic/gin"
)

func (b *BookModule) register(group *gin.RouterGroup) error {

	x := b.app.AuthMiddleware

	accGrp := group.Group("/:pid/account")

	accGrp.GET("/", x(b.listAccount))
	accGrp.POST("/", x(b.addAccount))
	accGrp.GET("/:id", x(b.getAccount))
	accGrp.POST("/:id", x(b.updateAccount))
	accGrp.DELETE("/:id", x(b.deleteAccount))

	txnGrp := group.Group("/:pid/txn")

	txnGrp.GET("/", x(b.listTxn))

	txnGrp.GET("/line/list", x(b.listTxnWithLines))
	txnGrp.GET("/line/:aid/list", x(b.listAccountTxnWithLines))

	txnGrp.POST("/", x(b.addTxn))
	txnGrp.GET("/:id", x(b.getTxn))
	txnGrp.POST("/:id", x(b.updateTxn))
	txnGrp.DELETE("/:id", x(b.deleteTxn))

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
	return b.dbOpGetTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
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

func (b *BookModule) deleteTxn(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpDeleteTxn(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
