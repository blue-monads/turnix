import{s as g,a as b,g as h,i as _,f as d,k as x,y as v,e as T,c as k,r as w,m as p,n as S}from"../chunks/scheduler.CPmtgLt4.js";import{S as A,i as y,c as m,b as c,m as l,t as f,a as u,d as $}from"../chunks/index.CsITmn2N.js";import{g as B}from"../chunks/entry.CNB9SHga.js";import{N as C}from"../chunks/index.17k-MAAo.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.XpbzWZVH.js";import{A as z}from"../chunks/AppBar.D1f81ABc.js";import{p as H}from"../chunks/stores.DZzzPzvs.js";import{T as I}from"../chunks/TxnForm.CZq6tAoy.js";function L(n){let t,s='<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li class="crumb">Add Transaction</li></ol>';return{c(){t=T("div"),t.innerHTML=s,this.h()},l(a){t=k(a,"DIV",{slot:!0,class:!0,"data-svelte-h":!0}),w(t)!=="svelte-k5c16z"&&(t.innerHTML=s),this.h()},h(){p(t,"slot","lead"),p(t,"class","flex gap-2")},m(a,r){_(a,t,r)},p:S,d(a){a&&d(t)}}}function M(n){let t,s,a,r;return t=new z({props:{$$slots:{lead:[L]},$$scope:{ctx:n}}}),a=new I({props:{onSubmit:n[0]}}),{c(){m(t.$$.fragment),s=b(),m(a.$$.fragment)},l(e){c(t.$$.fragment,e),s=h(e),c(a.$$.fragment,e)},m(e,o){l(t,e,o),_(e,s,o),l(a,e,o),r=!0},p(e,[o]){const i={};o&16&&(i.$$scope={dirty:o,ctx:e}),t.$set(i)},i(e){r||(f(t.$$.fragment,e),f(a.$$.fragment,e),r=!0)},o(e){u(t.$$.fragment,e),u(a.$$.fragment,e),r=!1},d(e){e&&d(s),$(t,e),$(a,e)}}}function N(n,t,s){let a;x(n,H,i=>s(1,a=i));const r=a.params.pid,e=C(v("__api__"));return[async i=>{(await e.addTxn(r,i)).status===200&&B(location.pathname.replace("/new",""))}]}class K extends A{constructor(t){super(),y(this,t,N,M,g,{})}}export{K as component};