import{a as o,t as r}from"../chunks/Celh2SBM.js";import{p as _,g as b,i as y,f as x,a as h,h as k,d as A,c as N,b as w,r as T,s as I}from"../chunks/BC9Cf8Q3.js";import{s as j}from"../chunks/DMueVNEq.js";import{p as m}from"../chunks/CKwprqj2.js";import{s as z,a as B}from"../chunks/BvXhjoeW.js";import{S as C}from"../chunks/CdW99GA5.js";import{N as P}from"../chunks/Dq_7tmzI.js";import"../chunks/CRz77-DP.js";import"../chunks/B857cCpa.js";import{A as S}from"../chunks/CnSBROsQ.js";import{p as q}from"../chunks/C1rFHlC-.js";import"../chunks/CMsATdMh.js";import{A as D}from"../chunks/C6IzKD7A.js";/* empty css                */import{g as R}from"../chunks/BW30Hxfo.js";var E=r('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li class="crumb">Taxes</li></ol>'),F=r('<a class="btn variant-ghost-primary btn-sm"><!> Tax</a>'),G=r("<!> <!>",1);function oa(c,d){_(d,!0);const[f,g]=z(),s=B(q,"$page",f).params.pid,e=P(b("__api__"));let p=y(m([]));const i=async()=>{const a=await e.listTax(s);a.status===200&&k(p,m(a.data))};i();var n=G(),l=x(n);S(l,{$$slots:{lead:(a,u)=>{var t=E();o(a,t)},trail:(a,u)=>{var t=F();j(t,"href",`/z/pages/portal/projects/books/${s}/inventory/tax/new`);var v=N(t);C(v,{className:"h-4 w-4",name:"plus"}),w(),T(t),o(a,t)}}});var $=I(l,2);D($,{action_key:"id",key_names:[["id","ID"],["name","Name"],["ttype","Type"],["info","Info"],["rate","Rate"]],get datas(){return A(p)},color:["ttype"],actions:[{Name:"edit",Class:"variant-filled-secondary",Action:async a=>{R(`/z/pages/portal/projects/books/${s}/inventory/tax/edit?pid=${s}&tid=${a}`)}},{Name:"delete",Class:"variant-filled-error",Action:async a=>{await e.deleteTax(s,a),i()}}]}),o(c,n),h(),g()}export{oa as component};
