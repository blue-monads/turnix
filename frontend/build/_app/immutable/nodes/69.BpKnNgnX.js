import{a as m,t as v,c as q}from"../chunks/disclose-version.wcsHG7oF.js";import{p as C,g as D,i as w,f as A,a as E,h as d,b as c,s as u,c as i,r as l,e as _}from"../chunks/runtime.D6lgFUio.js";import{i as F}from"../chunks/if.lkQ-IRrr.js";import{k as M}from"../chunks/key.Dx7O0vkj.js";import{s as N}from"../chunks/attributes.BFGPVQjG.js";import{p as f}from"../chunks/proxy.r5-P5EHK.js";import{s as R,a as y}from"../chunks/store.BaapSzAp.js";import{S as I}from"../chunks/SvgIcon.Cj01WnO8.js";import{f as W,T as G}from"../chunks/format.CvLTkFFY.js";import{g as H}from"../chunks/stores.CebXFGJ4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as J}from"../chunks/AppBar.3srT6C5i.js";import{p as K}from"../chunks/index.BTu0Thf0.js";import{N as O}from"../chunks/index.Duto3aX0.js";import"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";import{L as Q}from"../chunks/loader.BeY12cOT.js";import{p as U}from"../chunks/stores.Dnldv-Vm.js";var V=v('<div slot="lead" class="flex gap-2"><ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li class="crumb"><a class="anchor">Account</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li class="crumb">Transactions</li></ol></div>'),X=v('<div slot="trail" class="flex gap-2"><button class="inline-flex gap-1"><!> Find</button> <a class="inline-flex gap-1"><!> New</a></div>'),Y=v("<!> <!>",1);function ha(T,j){C(j,!0);const b=R(),z=()=>y(U,"$page",b),B=()=>y(K,"$params",b),n=z().params.pid,L=B().aid,h=O(D("__api__"));H();let g=w(!0),t=w(f([])),$=f({});(async()=>{const o=h.listAccount(n);d(g,!0);const r=await h.listAccTxnWithLines(n,L);if(r.status!==200)return;console.log("@data",r.data),d(t,f(W(r.data))),console.log("@data_____",c(t)),(await o).data.forEach(a=>{const e=a.id;$[e]=a.name}),d(t,f(c(t))),d(g,!1)})();var x=Y(),k=A(x);J(k,{$$slots:{lead:(o,r)=>{var s=V(),a=i(s),e=u(i(a),4),p=i(e);N(p,"href",`/z/pages/portal/projects/books/${n}`),l(e),_(4),l(a),l(s),m(o,s)},trail:(o,r)=>{var s=X(),a=i(s),e=i(a);I(e,{name:"magnifying-glass",className:"w-5 h-5"}),_(),l(a);var p=u(a,2);N(p,"href",`/z/pages/portal/projects/books/${n}/txn/new`);var S=i(p);I(S,{name:"plus",className:"w-5 h-5"}),_(),l(p),l(s),m(o,s)}}});var P=u(k,2);F(P,()=>c(g),o=>{Q(o,{})},o=>{var r=q(),s=A(r);M(s,()=>c(t),a=>{G(a,{get accountsIndex(){return $},get lineData(){return c(t)},pid:n})}),m(o,r)}),m(T,x),E()}export{ha as component};