import{c as w,a as x}from"../chunks/disclose-version.BOfWGsg9.js";import{p as N,g as k,i as m,f as v,a as I,h as n,b as e,d as c}from"../chunks/runtime.CmPLUjJJ.js";import{i as h}from"../chunks/if.DoU_1wcm.js";import{p as d}from"../chunks/proxy.DKs42Gzo.js";import{s as y,a as C}from"../chunks/store.BN9B0iWR.js";import{N as D}from"../chunks/index.Dq_7tmzI.js";import{p as L}from"../chunks/stores.BgpLsI8J.js";import"../chunks/index.DQ4XhlMg.js";import"../chunks/autotable.D5y0wA8l.js";import{L as j}from"../chunks/loader._7BlSOYJ.js";import{S as z}from"../chunks/SaleMaker.BDe_Ewx1.js";import{g as A}from"../chunks/entry.CytXzrzW.js";function U(l,g){N(g,!0);const f=y(),s=C(L,"$page",f).params.pid,o=D(k("__api__"));let p=m(d([])),r=m(!0);(async()=>{n(r,!0);const t=await o.listContacts(s);t.status===200&&(n(p,d(t.data)),n(r,!1))})();let u=c(()=>e(p).reduce((t,a)=>{const S=a.id,$=a.name;return t[S]=$,t},{}));const _=async t=>{console.log("@submit/data",t),t.sale.sales_date=new Date(t.sale.sales_date).toISOString();const a=await o.addSale(s,t);return a.status!==200?a.data.message:(A(`/z/pages/portal/projects/books/${s}/sales`),"")};var i=w(),b=v(i);h(b,()=>e(r),t=>{j(t,{})},t=>{var a=c(()=>Number(s));z(t,{name:"New Sale",get pid(){return e(a)},api:o,get contactsNameIndex(){return e(u)},submit:_})}),x(l,i),I()}export{U as component};
