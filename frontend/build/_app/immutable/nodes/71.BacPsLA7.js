import{a as r,t}from"../chunks/disclose-version.wcsHG7oF.js";import{p as g,g as u,f as _,a as v,s as $}from"../chunks/runtime.D6lgFUio.js";import{i as b}from"../chunks/lifecycle.ZPOFI82y.js";import{s as h,a as x}from"../chunks/store.BaapSzAp.js";import{g as A}from"../chunks/entry.dkjBYykr.js";import{N as T}from"../chunks/index.Duto3aX0.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as k}from"../chunks/AppBar.3srT6C5i.js";import{p as w}from"../chunks/stores.Dnldv-Vm.js";import{T as B}from"../chunks/TxnForm.B7Gt1Lkq.js";var N=t('<div slot="lead" class="flex gap-2"><ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li class="crumb">Add Transaction</li></ol></div>'),P=t("<!> <!>",1);function H(p,e){g(e,!1);const i=h(),m=x(w,"$page",i).params.pid,n=T(u("__api__")),l=async a=>{(await n.addTxn(m,a)).status===200&&A(location.pathname.replace("/new",""))};b();var o=P(),s=_(o);k(s,{$$slots:{lead:(a,d)=>{var f=N();r(a,f)}}});var c=$(s,2);B(c,{onSubmit:l}),r(p,o),v()}export{H as component};