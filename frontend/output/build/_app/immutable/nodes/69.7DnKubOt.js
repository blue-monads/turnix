import{a as l,t as g,c as O}from"../chunks/disclose-version.BOfWGsg9.js";import{p as X,i as h,g as G,f as x,t as K,a as Z,b as o,h as m,c as d,e as N,r as p,s as I,d as T}from"../chunks/runtime.CmPLUjJJ.js";import{d as J,e as V,b as tt}from"../chunks/events.DnYA5zBv.js";import{i as L}from"../chunks/if.DoU_1wcm.js";import{s as et}from"../chunks/attributes.yGGqmlPO.js";import{p as i}from"../chunks/proxy.DKs42Gzo.js";import{b as at}from"../chunks/this.DlUubvPq.js";import{s as st,a as ot}from"../chunks/store.BN9B0iWR.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as rt}from"../chunks/AppBar.pSMxueMf.js";import{T as nt,a as Q}from"../chunks/Tab.DYwUpmHc.js";import{C as Y}from"../chunks/index.63AqBVKj.js";import{s as lt,S as it}from"../chunks/index.BOkQpbOq.js";import{h as ct}from"../chunks/index.DtK7kxBu.js";import{p as dt}from"../chunks/stores.BgpLsI8J.js";import{N as pt}from"../chunks/index.Dq_7tmzI.js";import{S as R}from"../chunks/SvgIcon.Cn9REWmg.js";let mt=`
    
-- query_name: list_accounts
select * from __project__Accounts;

































`;const gt=`

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







`;var ut=g('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Reports</li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),vt=($,f,u)=>{m(f,i(o(u)))},ht=($,f)=>{var u;(u=o(f).contentWindow)==null||u.print()},ft=g('<button class="btn variant-filled btn-sm"><!> Run</button> <button class="btn variant-filled-secondary btn-sm"><!> Print</button>',1),_t=g("<span>SQL Query</span>"),bt=g("<span>UI</span>"),yt=g("<!> <!>",1),wt=g('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),xt=g('<!> <div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><div class="card p-2 h-full w-full"><iframe title="preview" width="100%" height="100%" class="border-green-200 w-full h-full transition-all" allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking" sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"></iframe></div></div></div>',1);function Rt($,f){X(f,!0);const u=st(),U=()=>ot(dt,"$page",u);let b=h(void 0),v=h(0),M=h(i(mt)),k=h(i(gt)),S=h("");const z=U().params.pid,j=G("__api__");pt(j);let y=h(void 0);const W=async r=>{var s,e,n;const a=r.data;if(console.log("onFrameMessage",a),a.type==="sql_query"){console.log("sql_query",a);try{const t=await j.runProjectSQL(z,{input:o(M),name:a.name,data:a.data});if(t.status!==200){(s=o(y))==null||s.postMessage({msgId:a.msgId,data:{msg:"error",data:t.data}});return}(e=o(y))==null||e.postMessage({msgId:a.msgId,data:{msg:"success",data:t.data}})}catch{}}else a.type==="api_call"?console.log("api_call",a):a.type==="ping"&&((n=o(y))==null||n.postMessage({msgId:a.msgId,data:{msg:"pong"}}))};var q=xt(),A=x(q);rt(A,{$$slots:{lead:(r,a)=>{var s=ut();l(r,s)},trail:(r,a)=>{var s=ft(),e=x(s);e.__click=[vt,S,k];var n=d(e);R(n,{className:"w-4 h-4",name:"play"}),N(),p(e);var t=I(e,2);t.__click=[ht,b];var c=d(t);R(c,{className:"w-4 h-4",name:"printer"}),N(),p(t),l(r,s)}}});var C=I(A,2),P=d(C),B=d(P);nt(B,{children:(r,a)=>{var s=yt(),e=x(s);Q(e,{padding:"ml-6 p-2",get group(){return o(v)},set group(t){m(v,i(t))},name:"SQL Query",value:0,children:(t,c)=>{var _=_t();l(t,_)},$$slots:{default:!0}});var n=I(e,2);Q(n,{get group(){return o(v)},set group(t){m(v,i(t))},name:"UI",value:1,children:(t,c)=>{var _=bt();l(t,_)},$$slots:{default:!0}}),l(r,s)},$$slots:{default:!0,panel:(r,a)=>{var s=wt(),e=d(s);L(e,()=>o(v)===0,n=>{var t=T(()=>lt({dialect:it}));Y(n,{get value(){return o(M)},set value(c){m(M,i(c))},get lang(){return o(t)}})},n=>{var t=O(),c=x(t);L(c,()=>o(v)===1,_=>{var F=T(ct);Y(_,{get value(){return o(k)},set value(H){m(k,i(H))},get lang(){return o(F)}})},null,!0),l(n,t)}),p(s),l(r,s)}}}),p(P);var D=I(P,2),E=d(D),w=d(E);at(w,r=>m(b,r),()=>o(b)),p(E),p(D),p(C),K(()=>et(w,"srcdoc",o(S))),V("load",w,r=>{var a,s;try{let e=new MessageChannel;e.port1.onmessage=W,m(y,i(e.port1)),console.log("chan.port2 type:",e.port2 instanceof MessagePort),(s=(a=o(b))==null?void 0:a.contentWindow)==null||s.postMessage("transfer_port","*",[e.port2])}catch(e){console.error("Error in postMessage:",e)}}),tt(w),l($,q),Z()}J(["click"]);export{Rt as component};
