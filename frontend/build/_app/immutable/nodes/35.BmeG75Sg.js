import{a as c,t as l}from"../chunks/disclose-version.wcsHG7oF.js";import{p as I,g as z,i as B,f as C,a as P,h as S,s as u,b as g,c as i,e as v,r as p,d as W}from"../chunks/runtime.D6lgFUio.js";import{s as h}from"../chunks/attributes.BFGPVQjG.js";import{p as _}from"../chunks/proxy.r5-P5EHK.js";import{s as q,a as D}from"../chunks/store.BaapSzAp.js";import{g as r}from"../chunks/entry.dkjBYykr.js";import"../chunks/auto_form.CxbPX2mo.js";import{A as M}from"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";/* empty css                                                     */import{S as $}from"../chunks/SvgIcon.Cj01WnO8.js";import{N as E}from"../chunks/index.Duto3aX0.js";import{g as F}from"../chunks/stores.CebXFGJ4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as G}from"../chunks/AppBar.3srT6C5i.js";import{p as H}from"../chunks/stores.Dnldv-Vm.js";var J=l('<div class="flex flex-wrap justify-end gap-2"><a class="btn variant-filled btn-sm"><!> account</a> <a class="btn variant-filled-secondary btn-sm"><!> transaction</a></div>'),K=l('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Account</li></ol>'),L=l("<!> <!>",1);function pa(b,y){I(y,!0);const x=q(),o=D(H,"$page",x).params.pid;F();const A=E(z("__api__"));let m=B(_([]));(async()=>{const a=await A.listAccount(o);a.status===200&&S(m,_(a.data))})();var d=L(),f=C(d);G(f,{children:(a,t)=>{var s=J(),e=i(s);h(e,"href",`/z/pages/portal/projects/books/${o}/account/new`);var N=i(e);$(N,{className:"h-4 w-4",name:"plus"}),v(),p(e);var n=u(e,2);h(n,"href",`/z/pages/portal/projects/books/${o}/txn/new`);var j=i(n);$(j,{className:"h-4 w-4",name:"plus"}),v(),p(n),p(s),c(a,s)},$$slots:{default:!0,lead:(a,t)=>{var s=K();c(a,s)}}});var k=u(f,2),w=W(()=>[{Name:"explore txns",Class:"variant-filled-primary",icon:"plus",Action:async a=>{const t=`/txn/account?pid=${o}&aid=${a}`;location.pathname.endsWith("/account")?r(location.pathname.replace("/account",t)):r(location.pathname+t)}},{Name:"edit",Class:"variant-filled-secondary",Action:async a=>{const t=`/account/edit?pid=${o}&aid=${a}`;location.pathname.endsWith("/account")?r(location.pathname.replace("/account",t)):r(location.pathname+t)}}]);M(k,{action_key:"id",key_names:[["id","ID"],["name","Name"],["info","Info"],["acc_type","Account type"]],get datas(){return g(m)},color:["acc_type"],get actions(){return g(w)}}),c(b,d),P()}export{pa as component};