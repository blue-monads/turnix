var r=Object.defineProperty;var i=(s,t,e)=>t in s?r(s,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):s[t]=e;var o=(s,t,e)=>i(s,typeof t!="symbol"?t+"":t,e);const u=s=>new c(s.projectClient);class c{constructor(t){o(this,"client");o(this,"listAccount",t=>this.client.get(`books/${t}/account`));o(this,"addAccount",(t,e)=>this.client.post(`books/${t}/account`,e));o(this,"getAccount",(t,e)=>this.client.get(`books/${t}/account/${e}`));o(this,"updateAccount",(t,e,n)=>this.client.post(`books/${t}/account/${e}`,n));o(this,"deleteAccount",(t,e)=>this.client.delete(`books/${t}/account/${e}`));o(this,"listTxn",t=>this.client.get(`books/${t}/txn`));o(this,"listTxnWithLines",(t,e)=>this.client.get(`books/${t}/txn/line/list?offset=${e||0}`));o(this,"listAccTxnWithLines",(t,e,n)=>this.client.get(`books/${t}/txn/line/${e}/list?offset=${n||0}`));o(this,"addTxn",(t,e)=>this.client.post(`books/${t}/txn`,e));o(this,"getTxn",(t,e)=>this.client.get(`books/${t}/txn/${e}`));o(this,"getTxnWithLines",(t,e)=>this.client.get(`books/${t}/txn/${e}/line`));o(this,"updateTxn",(t,e,n)=>this.client.post(`books/${t}/txn/${e}`,n));o(this,"updateTxnWithLine",(t,e,n)=>this.client.post(`books/${t}/txn/${e}/line`,n));o(this,"deleteTxn",(t,e)=>this.client.delete(`books/${t}/txn/${e}`));o(this,"generateLiveTxn",(t,e)=>this.client.post(`books/${t}/report/live`,e));o(this,"exportData",t=>this.client.post(`books/${t}/txn/export`));o(this,"importData",(t,e)=>this.client.post(`books/${t}/txn/import`,e));o(this,"listCatagories",t=>this.client.get(`books/${t}/inventory/catagories`));o(this,"addCatagory",(t,e)=>this.client.post(`books/${t}/inventory/catagories`,e));o(this,"getCatagory",(t,e)=>this.client.get(`books/${t}/inventory/catagories/${e}`));o(this,"updateCatagory",(t,e,n)=>this.client.post(`books/${t}/inventory/catagories/${e}`,n));o(this,"deleteCatagory",(t,e)=>this.client.delete(`books/${t}/inventory/catagories/${e}`));o(this,"listProducts",t=>this.client.get(`books/${t}/inventory/products`));o(this,"addProduct",(t,e)=>this.client.post(`books/${t}/inventory/products`,e));o(this,"getProduct",(t,e)=>this.client.get(`books/${t}/inventory/products/${e}`));o(this,"updateProduct",(t,e,n)=>this.client.post(`books/${t}/inventory/products/${e}`,n));o(this,"deleteProduct",(t,e)=>this.client.delete(`books/${t}/inventory/products/${e}`));o(this,"listContacts",t=>this.client.get(`books/${t}/contacts`));o(this,"addContact",(t,e)=>this.client.post(`books/${t}/contacts`,e));o(this,"getContact",(t,e)=>this.client.get(`books/${t}/contacts/${e}`));o(this,"updateContact",(t,e,n)=>this.client.post(`books/${t}/contacts/${e}`,n));o(this,"deleteContact",(t,e)=>this.client.delete(`books/${t}/contacts/${e}`));o(this,"addSale",(t,e)=>this.client.post(`books/${t}/sales`,e));o(this,"getSale",(t,e)=>this.client.get(`books/${t}/sales/${e}`));o(this,"listSales",t=>this.client.get(`books/${t}/sales`));o(this,"deleteSale",(t,e)=>this.client.delete(`books/${t}/sales/${e}`));o(this,"listTax",t=>this.client.get(`books/${t}/tax`));o(this,"addTax",(t,e)=>this.client.post(`books/${t}/tax`,e));o(this,"getTax",(t,e)=>this.client.get(`books/${t}/tax/${e}`));o(this,"updateTax",(t,e,n)=>this.client.post(`books/${t}/tax/${e}`,n));o(this,"deleteTax",(t,e)=>this.client.delete(`books/${t}/tax/${e}`));o(this,"listTaxProduct",(t,e)=>this.client.get(`books/${t}/tax/product?tid=${e||""}`));o(this,"addTaxProduct",(t,e,n)=>this.client.post(`books/${t}/tax/${e}/product`,n));o(this,"deleteTaxProduct",(t,e,n)=>this.client.delete(`books/${t}/tax/${e}/product/${n}`));o(this,"listProductStockIn",t=>this.client.get(`books/${t}/stocks`));o(this,"addProductStockIn",(t,e)=>this.client.post(`books/${t}/stocks`,e));o(this,"getProductStockIn",(t,e)=>this.client.get(`books/${t}/stocks/${e}`));o(this,"deleteProductStockIn",(t,e)=>this.client.delete(`books/${t}/stocks/${e}`));this.client=t}}export{u as N};
