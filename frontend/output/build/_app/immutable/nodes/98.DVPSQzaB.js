import{c as k,a as A}from"../chunks/disclose-version.BOfWGsg9.js";import{p as L,g as h,i as o,f as N,a as X,h as e,b as r}from"../chunks/runtime.CmPLUjJJ.js";import{i as v}from"../chunks/if.DoU_1wcm.js";import{p}from"../chunks/proxy.DKs42Gzo.js";import{s as w,a as _}from"../chunks/store.BN9B0iWR.js";import{A as x}from"../chunks/index.DQ4XhlMg.js";import"../chunks/autotable.D5y0wA8l.js";import{L as C}from"../chunks/loader._7BlSOYJ.js";import{f as U}from"../chunks/route.CqicUvmm.js";import{p as b}from"../chunks/index.NL8p-q19.js";import{N as K}from"../chunks/index.DQVP37Ea.js";import{p as O}from"../chunks/stores.BgpLsI8J.js";function H(u,g){L(g,!0);const m=w(),T=()=>_(O,"$page",m),y=()=>_(b,"$params",m),n=K(h("__api__"));let a=T().params.pid,i=y().tid,f=o(""),s=o(!0),c=o(p({}));(async()=>{e(s,!0);const t=await n.getTemplate(a,i);t.status===200&&(e(c,p(t.data)),e(s,!1))})();var l=k(),E=N(l);v(E,()=>r(s),t=>{C(t,{})},t=>{x(t,{get data(){return r(c)},get message(){return r(f)},schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"KEY_VALUE_TEXT",key_name:"result",name:"Result"},{ftype:"LONG_TEXT",key_name:"contentScript",name:"Content script"}],name:"Edit Template",required_fields:["name"]},onSave:async $=>{const d=await n.updateTemplate(a,i,$);if(d.status!==200){e(f,p(d.data));return}U(a)}})}),A(u,l),X()}export{H as component};
