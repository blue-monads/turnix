import{c as x,a as E}from"../chunks/disclose-version.wcsHG7oF.js";import{p as N,g as b,f as h,a as v,h as t,b as s,i as r}from"../chunks/runtime.D6lgFUio.js";import{i as L}from"../chunks/if.lkQ-IRrr.js";import{p as g}from"../chunks/proxy.r5-P5EHK.js";import{s as w,a as u}from"../chunks/store.BaapSzAp.js";import{N as A}from"../chunks/index.Duto3aX0.js";import{p as C}from"../chunks/stores.D4Dcgzbh.js";import{A as I}from"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";import{L as S}from"../chunks/loader.BeY12cOT.js";import{g as X}from"../chunks/entry.Y6HcnCx3.js";import{p as j}from"../chunks/index.xGEmn6Hr.js";function U(_,y){N(y,!0);const p=w(),l=()=>u(C,"$page",p),T=()=>u(j,"$params",p),e=l().params.pid,m=T().tid,n=A(b("__api__"));let i=r(""),f=r(void 0),o=r(!0);(async()=>{t(o,!0);const a=await n.getTax(e,m);a.status===200&&(t(f,g(a.data)),t(o,!1))})();const k=async a=>{const c=await n.updateTax(e,m,a);if(c.status!==200){t(i,g(c.data.message));return}X(`/z/pages/portal/projects/books/${e}/inventory/tax`)};var d=x(),$=h(d);L($,()=>s(o),a=>{S(a,{})},a=>{I(a,{get message(){return s(i)},get data(){return s(f)},schema:{fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Type",ftype:"SELECT",key_name:"ttype",options:["item_percent"]},{name:"Rate",ftype:"INT",key_name:"rate"},{name:"Info",ftype:"LONG_TEXT",key_name:"info"}],name:"Edit Catagory",required_fields:[]},onSave:k})}),E(_,d),v()}export{U as component};