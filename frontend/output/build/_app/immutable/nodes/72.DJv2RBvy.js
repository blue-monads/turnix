import{a as v,t as A,c as O}from"../chunks/disclose-version.BOfWGsg9.js";import{c as a,s as r,f as B,r as t,b as e,t as J,d as et,p as lt,g as st,i as H,a as dt,h as I,e as U}from"../chunks/runtime.CmPLUjJJ.js";import{d as nt}from"../chunks/events.DnYA5zBv.js";import{i as V}from"../chunks/if.DoU_1wcm.js";import{k as at}from"../chunks/key.B3I6Q0YW.js";import{s as W}from"../chunks/attributes.yGGqmlPO.js";import{p as X}from"../chunks/proxy.DKs42Gzo.js";import{b as it}from"../chunks/select.BNLwbYi8.js";import{s as ct,a as vt}from"../chunks/store.BN9B0iWR.js";import{r as pt}from"../chunks/legacy-client.BxBrBVYa.js";import{S as _t}from"../chunks/SvgIcon.Cn9REWmg.js";import{N as ut}from"../chunks/index.Dq_7tmzI.js";import{g as mt}from"../chunks/stores.C0Vv8-i4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as ht}from"../chunks/AppBar.pSMxueMf.js";import{p as bt}from"../chunks/stores.DkdCKZTz.js";import{s as i}from"../chunks/render.CZQ66Y_F.js";import{e as Y,i as Z}from"../chunks/each.BJPUi07-.js";var ft=A("<tr><td> </td><td> </td><td> </td><td> </td></tr>"),gt=A('<tr><td colspan="5"><span class="font-semibold text-lg"> </span></td></tr> <!> <tr><th colspan="2">Total</th><td> </td><td> </td></tr>',1),xt=A('<div class="table-container h-[calc(100vh-14rem)] overflow-y-auto"><table class="table table-hover h-fit"><thead><tr><th>Txn</th><th>Title</th><th>Debit Amount</th><th>Credit Amount</th></tr></thead><tbody></tbody></table></div>');function yt(L,_){const u=n=>{let o={};return n.forEach(s=>{let l=o[s.account_id];l||(l={account_id:s.account_id,account_name:s.account_name,account_type:s.acc_type,lines:[],total_credit:s.total_credit,total_debit:s.total_debit},o[s.account_id]=l),l.lines.push(s)}),console.log("@after_format",o),o};let b=et(()=>u(_.data));var c=xt(),f=a(c),d=r(a(f));Y(d,21,()=>Object.values(e(b)),Z,(n,o)=>{var s=gt(),l=B(s),N=a(l),g=a(N),S=a(g);t(g),t(N),t(l);var m=r(l,2);Y(m,17,()=>e(o).lines,Z,(q,x)=>{var T=ft(),C=a(T),M=a(C);t(C);var y=r(C),K=a(y);t(y);var D=r(y),Q=a(D);t(D);var $=r(D),z=a($);t($),t(T),J(()=>{i(M,e(x).txn_id),i(K,e(x).title),i(Q,e(x).debit_amount),i(z,e(x).credit_amount)}),v(q,T)});var k=r(m,2),p=r(a(k)),w=a(p);t(p);var j=r(p),F=a(j);t(j),t(k),J(()=>{i(S,e(o).account_name),i(w,e(o).total_debit),i(F,e(o).total_credit)}),v(n,s)}),t(d),t(f),t(c),v(L,c)}var $t=A("<tr><th> </th><td> </td><td> </td><td> </td><td> </td></tr>"),kt=A('<div class="table-container h-[calc(100vh-14rem)] overflow-y-auto"><table class="table table-hover h-fit"><thead><tr><th>Id</th><th>Account</th><th>Account Type</th><th>Debit Total</th><th>Credit Total</th></tr></thead><tbody></tbody></table></div>');function Tt(L,_){var u=kt(),b=a(u),c=r(a(b));Y(c,21,()=>_.data,Z,(f,d)=>{var n=$t(),o=a(n),s=a(o);t(o);var l=r(o),N=a(l);t(l);var g=r(l),S=a(g);t(g);var m=r(g),k=a(m);t(m);var p=r(m),w=a(p);t(p),t(n),J(()=>{i(s,e(d).account_id),i(N,e(d).account_name),i(S,e(d).acc_type),i(k,e(d).total_debit),i(w,e(d).total_credit)}),v(f,n)}),t(c),t(b),t(u),v(L,u)}const At=async(L,_,u,b,c,f,d,n)=>{I(_,!0),I(u,X(e(b)));const o=await c.generateLiveTxn(f,{report_type:e(b),template_id:d});if(o.status!==200){I(_,!1);return}I(n,X(o.data)),console.log("@data",e(n)),I(_,!1)};var Lt=A('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator">&rsaquo;</li> <li><a class="anchor">Reports</a></li> <li class="crumb-separator">&rsaquo;</li> <li>Static Report</li></ol>'),St=A('<!> <div class="flex flex-col p-2 gap-2"><div class="card p-2"><header class="card-header"><h4 class="h4">Static Reports</h4></header> <section class="pl-4 py-2 flex justify-between items-center flex-wrap gap-2"><div class="flex gap-2 items-center flex-wrap"><label class="label"><span>Report Type</span> <select class="select"><option>Ledger short</option><option>Ledger long</option><option>Custom</option></select></label> <label class="label"><span>From Date</span> <input class="input" type="datetime-local"></label> <label class="label"><span>To Date</span> <input class="input" type="datetime-local"></label></div> <div><button class="btn btn-sm variant-filled disabled:variant-ghost disabled:text-slate-800"><!> Generate</button></div></section></div> <div class="p-2 card"><!></div></div>',1);function Kt(L,_){lt(_,!0);const u=ct(),c=vt(bt,"$page",u).params.pid;mt();const f=ut(st("__api__"));let d=H(!1),n=0,o=H("short_ledger"),s=H(X(e(o))),l=H(void 0);(async()=>{})();let g=et(()=>e(o)==="custom"&&n===0);pt(()=>{console.log("@data/2",e(l))});var S=St(),m=B(S);ht(m,{$$slots:{lead:($,z)=>{var G=Lt(),R=r(a(G),2);W(R,"aria-hidden",!0);var h=r(R,2),P=a(h);W(P,"href",`/z/pages/portal/projects/books/${c}/report`),t(h);var E=r(h,2);W(E,"aria-hidden",!0),U(2),t(G),v($,G)}}});var k=r(m,2),p=a(k),w=r(a(p),2),j=a(w),F=a(j),q=r(a(F),2),x=a(q);x.value=(x.__value="short_ledger")==null?"":"short_ledger";var T=r(x);T.value=(T.__value="long_ledger")==null?"":"long_ledger";var C=r(T);C.value=(C.__value="custom")==null?"":"custom",t(q),t(F),U(4),t(j);var M=r(j,2),y=a(M);y.__click=[At,d,s,o,f,c,n,l];var K=a(y);_t(K,{className:"h-4 w-4",name:"bolt"}),U(),t(y),t(M),t(w),t(p);var D=r(p,2),Q=a(D);V(Q,()=>e(l),$=>{var z=O(),G=B(z);V(G,()=>e(s)==="long_ledger",R=>{var h=O(),P=B(h);at(P,()=>e(l),E=>{yt(E,{get data(){return e(l)}})}),v(R,h)},R=>{var h=O(),P=B(h);V(P,()=>e(s)==="short_ledger",E=>{var tt=O(),rt=B(tt);at(rt,()=>e(l),ot=>{Tt(ot,{get data(){return e(l)}})}),v(E,tt)},null,!0),v(R,h)}),v($,z)}),t(D),t(k),J(()=>y.disabled=e(d)||e(g)),it(q,()=>e(o),$=>I(o,$)),v(L,S),dt()}nt(["click"]);export{Kt as component};