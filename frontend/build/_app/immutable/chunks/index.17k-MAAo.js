var s=Object.defineProperty;var r=(o,t,e)=>t in o?s(o,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):o[t]=e;var n=(o,t,e)=>(r(o,typeof t!="symbol"?t+"":t,e),e);const u=o=>new c(o.projectClient);class c{constructor(t){n(this,"client");n(this,"listAccount",t=>this.client.get(`books/${t}/account`));n(this,"addAccount",(t,e)=>this.client.post(`books/${t}/account`,e));n(this,"getAccount",(t,e)=>this.client.get(`books/${t}/account/${e}`));n(this,"updateAccount",(t,e,i)=>this.client.post(`books/${t}/account/${e}`,i));n(this,"deleteAccount",(t,e)=>this.client.delete(`books/${t}/account/${e}`));n(this,"listTxn",t=>this.client.get(`books/${t}/txn`));n(this,"listTxnWithLines",(t,e)=>this.client.get(`books/${t}/txn/line/list?offset=${e||0}`));n(this,"listAccTxnWithLines",(t,e,i)=>this.client.get(`books/${t}/txn/line/${e}/list?offset=${i||0}`));n(this,"addTxn",(t,e)=>this.client.post(`books/${t}/txn`,e));n(this,"getTxn",(t,e)=>this.client.get(`books/${t}/txn/${e}`));n(this,"getTxnWithLines",(t,e)=>this.client.get(`books/${t}/txn/${e}/line`));n(this,"updateTxn",(t,e,i)=>this.client.post(`books/${t}/txn/${e}`,i));n(this,"updateTxnWithLine",(t,e,i)=>this.client.post(`books/${t}/txn/${e}/line`,i));n(this,"deleteTxn",(t,e)=>this.client.delete(`books/${t}/txn/${e}`));n(this,"generateLiveTxn",(t,e)=>this.client.post(`books/${t}/report/live`,e));n(this,"exportData",t=>this.client.post(`books/${t}/txn/export`));n(this,"importData",(t,e)=>this.client.post(`books/${t}/txn/import`,e));this.client=t}}export{u as N};
