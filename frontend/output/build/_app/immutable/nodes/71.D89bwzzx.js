import{a as c,t as u,c as Z}from"../chunks/Celh2SBM.js";import{p as J,i as f,g as V,f as x,t as tt,a as et,d as o,h as g,c as p,b as L,r as m,s as I,e as Q}from"../chunks/BC9Cf8Q3.js";import{d as at,e as st,c as ot}from"../chunks/x8LL7je8.js";import{i as Y}from"../chunks/BgSg1_cg.js";import{s as rt}from"../chunks/DMueVNEq.js";import{p as d}from"../chunks/CKwprqj2.js";import{b as nt}from"../chunks/B96pt0vU.js";import{s as lt,a as it}from"../chunks/BvXhjoeW.js";import"../chunks/CRz77-DP.js";import"../chunks/B857cCpa.js";import{A as ct}from"../chunks/CnSBROsQ.js";import{T as dt,a as R}from"../chunks/Bzk7-iQC.js";import{C as U}from"../chunks/Czp9zhDz.js";import{s as pt,S as mt}from"../chunks/DPRgOf0J.js";import{h as gt}from"../chunks/CZDi8ouT.js";import{p as ut}from"../chunks/C1rFHlC-.js";import{N as vt}from"../chunks/Dq_7tmzI.js";import{S as z}from"../chunks/CdW99GA5.js";let ht=`
    
-- query_name: list_accounts
select * from __project__Accounts;

































`;const ft=`

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="//unpkg.com/alpinejs" defer><\/script>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"><\/script>
    <script>
        let pendingMessages = {}
        let messageId = 0;
        window.addEventListener("message", (ev) => {
            console.log("message", ev);
            if (ev.data !== "transfer_port") {
                console.log("wrong message", ev);
                return;
            }
            const port = ev.ports[0];
            port.onmessage = handleMessage;
            window["__parent_port__"] = port;
        }, false);
        const handleMessage = (ev) => {
            const data = ev.data;
            const messageId = data.msgId;
            const resolve = pendingMessages[messageId];
            delete pendingMessages[messageId];
            resolve(data.data);
            console.log("data", data);
        }
        const sendMessage = async (data) => {
            console.log("sendMessage", data);
            const msgId = messageId + 1;
            messageId = msgId;
            const p = new Promise((resolve) => {
                pendingMessages[msgId] = resolve;
            });
            data["msgId"] = msgId;
            const port = window["__parent_port__"]
            port?.postMessage(data);
            return p;
        }
    <\/script>
</head>
<body>
    <div class="container">
        <h1>Reports Editor Example</h1> 
        <p>This is a sample preview</p>
        
    </div>
    <div class="container-xl mt-5" x-data="accountTable()" id="ac_table">
        
    <div class="container">
        <button id="start" class="btn btn-primary" type="button" x-on:click="fetchAccounts">
            <span class="icon icon-thumbs-up"></span>
            Start
        </button>    
    </div>


        <div x-show="!loading" class="table-responsive">
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Info</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                </thead>
                <tbody>
                    <template x-for="account in accounts" :key="account.id">
                        <tr>
                            <td x-text="account.id"></td>
                            <td x-text="account.name"></td>
                            <td x-text="account.acc_type" class="text-capitalize"></td>
                            <td x-text="account.info"></td>
                            <td x-text="formatDate(account.created_at)"></td>
                            <td x-text="formatDate(account.updated_at)"></td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </div>

    <script>
    function accountTable() {
      return {
          accounts: [],
          loading: true,
          async fetchAccounts() {
              this.loading = true;
              try {
                  const data = await sendMessage({
                      type: "sql_query",
                      name: "list_accounts",
                      data: [],
                  });
                  console.log("accounts", data);

                  this.accounts = data.data;
              } catch (error) {
                  console.error("Error fetching accounts:", error);
              } finally {
                  this.loading = false;
              }
          },
          formatDate(dateString) {
              return new Date(dateString).toLocaleDateString();
          }
      }
  }


    <\/script>
</body>
</html>







`;var _t=u('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Reports</li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),bt=($,_,v)=>{g(_,d(o(v)))},yt=($,_)=>{var v;(v=o(_).contentWindow)==null||v.print()},wt=u('<button class="btn variant-filled btn-sm"><!> Run</button> <button class="btn variant-filled-secondary btn-sm"><!> Print</button>',1),xt=u("<span>SQL Query</span>"),It=u("<span>UI</span>"),$t=u("<!> <!>",1),Mt=u('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),kt=u('<!> <div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><div class="card p-2 h-full w-full"><iframe title="preview" width="100%" height="100%" class="border-green-200 w-full h-full transition-all" allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking" sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"></iframe></div></div></div>',1);function Ft($,_){J(_,!0);const[v,W]=lt(),B=()=>it(ut,"$page",v);let b=f(void 0),h=f(0),M=f(d(ht)),k=f(d(ft)),q=f("");const F=B().params.pid,A=V("__api__");vt(A);let y=f(void 0);const H=async r=>{var s,t,l;const e=r.data;if(console.log("onFrameMessage",e),e.type==="sql_query"){console.log("sql_query",e);try{const a=await A.runProjectSQL(F,{input:o(M),name:e.name,data:e.data});if(a.status!==200){(s=o(y))==null||s.postMessage({msgId:e.msgId,data:{msg:"error",data:a.data}});return}(t=o(y))==null||t.postMessage({msgId:e.msgId,data:{msg:"success",data:a.data}})}catch{}}else e.type==="api_call"?console.log("api_call",e):e.type==="ping"&&((l=o(y))==null||l.postMessage({msgId:e.msgId,data:{msg:"pong"}}))};var C=kt(),D=x(C);ct(D,{$$slots:{lead:(r,e)=>{var s=_t();c(r,s)},trail:(r,e)=>{var s=wt(),t=x(s);t.__click=[bt,q,k];var l=p(t);z(l,{className:"w-4 h-4",name:"play"}),L(),m(t);var a=I(t,2);a.__click=[yt,b];var n=p(a);z(n,{className:"w-4 h-4",name:"printer"}),L(),m(a),c(r,s)}}});var E=I(D,2),P=p(E),O=p(P);dt(O,{children:(r,e)=>{var s=$t(),t=x(s);R(t,{padding:"ml-6 p-2",name:"SQL Query",value:0,get group(){return o(h)},set group(a){g(h,d(a))},children:(a,n)=>{var i=xt();c(a,i)},$$slots:{default:!0}});var l=I(t,2);R(l,{name:"UI",value:1,get group(){return o(h)},set group(a){g(h,d(a))},children:(a,n)=>{var i=It();c(a,i)},$$slots:{default:!0}}),c(r,s)},$$slots:{default:!0,panel:(r,e)=>{var s=Mt(),t=p(s);{var l=n=>{const i=Q(()=>pt({dialect:mt}));U(n,{get lang(){return o(i)},get value(){return o(M)},set value(S){g(M,d(S))}})},a=n=>{var i=Z(),S=x(i);{var X=j=>{const G=Q(gt);U(j,{get lang(){return o(G)},get value(){return o(k)},set value(K){g(k,d(K))}})};Y(S,j=>{o(h)===1&&j(X)},!0)}c(n,i)};Y(t,n=>{o(h)===0?n(l):n(a,!1)})}m(s),c(r,s)}}}),m(P);var N=I(P,2),T=p(N),w=p(T);nt(w,r=>g(b,r),()=>o(b)),m(T),m(N),m(E),tt(()=>rt(w,"srcdoc",o(q))),st("load",w,r=>{var e,s;try{let t=new MessageChannel;t.port1.onmessage=H,g(y,d(t.port1)),console.log("chan.port2 type:",t.port2 instanceof MessagePort),(s=(e=o(b))==null?void 0:e.contentWindow)==null||s.postMessage("transfer_port","*",[t.port2])}catch(t){console.error("Error in postMessage:",t)}}),ot(w),c($,C),et(),W()}at(["click"]);export{Ft as component};
