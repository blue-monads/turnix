import{b as z,a as d,t as U,c as pt}from"./disclose-version.BOfWGsg9.js";import{p as et,b as F,h as D,t as dt,a as st,i as O,c as a,r as e,s,e as J,d as ut,g as mt,f as bt}from"./runtime.CmPLUjJJ.js";import{i as rt}from"./if.DoU_1wcm.js";import{p as T}from"./proxy.DKs42Gzo.js";import{p as o,s as _t,r as ft}from"./props.IjRlZquB.js";import{s as vt,a as xt}from"./store.BN9B0iWR.js";import"./entry.CytXzrzW.js";import"./index.DQ4XhlMg.js";import"./autotable.D5y0wA8l.js";import{L as gt}from"./loader._7BlSOYJ.js";import{d as ht,e as yt}from"./events.DnYA5zBv.js";import{s as kt,r as h,b as at}from"./attributes.yGGqmlPO.js";import{r as wt}from"./misc.CveY8KIp.js";import{b as S}from"./input.BRsTNa69.js";import{p as At}from"./event-modifiers.CWmzcjye.js";import{S as M}from"./SvgIcon.Cn9REWmg.js";import"./ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{F as Pt}from"./FileDropzone.hy4qW1aM.js";import{N as Nt}from"./index.Dq_7tmzI.js";import{g as It}from"./stores.C0Vv8-i4.js";import{p as St}from"./stores.BgpLsI8J.js";const Dt=async(b,t,c)=>{const l=await t.onPickAccount();l&&c(l)},Tt=async(b,t,c)=>{const l=await t.onPickAccount();l&&c(l)};var Ft=U('<div slot="lead" class="flex justify-center"><!></div>'),jt=U("<strong>Upload a file</strong> or drag and drop",1),Ct=U('<div class="p-2"><form class="card"><header class="card-header"><h4 class="h4"><!></h4></header> <section class="p-4 flex flex-col gap-4"><label class="label"><span>Title</span> <input name="title" class="input p-1" type="text" placeholder="Opening balance"></label> <div class="label border p-2 flex flex-col"><span>Lines</span> <div class="flex flex-row justify-start gap-4 items-center"><label class="label"><span>Debit Account</span> <input class="input p-1" type="text" disabled placeholder="Petty cash"></label> <button class="pt-4 cursor-pointer" type="button"><!></button> <label class="label"><span>Amount</span> <input name="debit amount" class="input p-1" type="number"></label></div> <div class="flex flex-row justify-start gap-4"><label class="label"><span>Credit Account</span> <input class="input p-1" type="text" disabled placeholder="Bank XYZ"></label> <button class="pt-4 cursor-pointer" type="button"><!></button> <label class="label"><span>Amount</span> <input name="credit amount" class="input p-1" type="number"></label></div></div> <label class="label"><span>Notes</span> <textarea class="textarea p-1" rows="4"></textarea></label> <label class="label"><span>Attachments</span> <!></label> <label class="label"><span>Reference ID</span> <input class="input p-1" type="text" placeholder="REF-123"></label></section> <footer class="card-footer flex justify-end"><button type="submit" class="btn variant-filled">save</button></footer></form></div>');function Et(b,t){et(t,!0);let c=O(void 0),l=o(t,"edit",3,!1),y=o(t,"title",15,""),k=o(t,"notes",15,""),_=o(t,"debit_account_id",15,0),f=o(t,"credit_account_id",15,0),u=o(t,"debit_amount",15,100),v=o(t,"credit_amount",15,100),w=o(t,"reference_id",3,""),X=o(t,"attachments",19,()=>({})),x=o(t,"accountsIndex",19,()=>({}));var g=Ct(),n=a(g),m=ut(()=>At(()=>{t.onSubmit({title:y(),notes:k(),debit_account_id:_(),credit_account_id:f(),debit_amount:u(),credit_amount:v(),reference_id:w(),attachments:X()})})),A=a(n),P=a(A),nt=a(P);rt(nt,l,r=>{var i=z("Edit Transaction");d(r,i)},r=>{var i=z("Add Transaction");d(r,i)}),e(P),e(A);var Y=s(A,2),j=a(Y),Z=s(a(j),2);h(Z),e(j);var C=s(j,2),E=s(a(C),2),L=a(E),q=s(a(L),2);h(q),e(L);var N=s(L,2);N.__click=[Dt,t,_];var ot=a(N);M(ot,{name:"arrow-up-right",className:"w-5 h-5"}),e(N);var H=s(N,2),K=s(a(H),2);h(K),e(H),e(E);var Q=s(E,2),B=a(Q),V=s(a(B),2);h(V),e(B);var I=s(B,2);I.__click=[Tt,t,f];var lt=a(I);M(lt,{name:"arrow-up-right",className:"w-5 h-5"}),e(I);var W=s(I,2),$=s(a(W),2);h($),e(W),e(Q),e(C);var G=s(C,2),R=s(a(G),2);wt(R),kt(R,"placeholder","Extra information about txn"),e(G);var tt=s(G,2),it=s(a(tt),2);Pt(it,{get files(){return F(c)},set files(r){D(c,T(r))},name:"files",$$slots:{lead:(r,i)=>{var p=Ft(),ct=a(p);M(ct,{name:"arrow-up-tray",className:"w-6 h-6"}),e(p),d(r,p)},message:(r,i)=>{var p=jt();J(),d(r,p)},meta:(r,i)=>{var p=z("PNG, JPG or PDF");d(r,p)}}}),e(tt),J(2),e(Y),J(2),e(n),e(g),dt(()=>{at(q,x()[_()]||""),at(V,x()[f()]||"")}),yt("submit",n,function(...r){var i;(i=F(m))==null||i.apply(this,r)}),S(Z,y),S(K,u),S($,v),S(R,k),d(b,g),st()}ht(["click"]);function ea(b,t){et(t,!0);const c=vt(),l=()=>xt(St,"$page",c);let y=ft(t,["$$slots","$$events","$$legacy","onSubmit"]);const k=l().params.pid,_=Nt(mt("__api__")),f=It();let u=O(!0),v=T({}),w=O(T([]));(async()=>{D(u,!0);const n=await _.listAccount(k);n.status===200&&((n.data||[]).forEach(m=>{v[Number(m.id)]=m.name}),D(w,T(n.data)),D(u,!1))})();var x=pt(),g=bt(x);rt(g,()=>F(u),n=>{gt(n,{})},n=>{Et(n,_t({get onSubmit(){return t.onSubmit},get accountsIndex(){return v},onPickAccount:async()=>new Promise((m,A)=>{f.trigger({type:"component",component:"books_account_picker",meta:{data:F(w),onClick:P=>{m(P)}}})})},()=>y))}),d(b,x),st()}export{ea as T};
