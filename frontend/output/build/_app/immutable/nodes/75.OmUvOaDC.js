import{a as i,t as v,b as rt}from"../chunks/disclose-version.BOfWGsg9.js";import{p as st,g as ot,i as m,u as lt,t as R,a as nt,h as s,b as t,f as dt,c as r,s as h,d as A,r as e,e as it}from"../chunks/runtime.CmPLUjJJ.js";import{s as F}from"../chunks/render.CZQ66Y_F.js";import{i as M}from"../chunks/if.DoU_1wcm.js";import{e as z,i as D}from"../chunks/each.BJPUi07-.js";import{d as pt}from"../chunks/events.DnYA5zBv.js";import{p}from"../chunks/proxy.DKs42Gzo.js";import{p as vt}from"../chunks/props.IjRlZquB.js";import{s as ut,a as ft}from"../chunks/store.BN9B0iWR.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{T as ct,a as K}from"../chunks/Tab.DYwUpmHc.js";import{C as U}from"../chunks/index.63AqBVKj.js";import{s as mt,S as _t}from"../chunks/index.BOkQpbOq.js";import{S as gt}from"../chunks/SvgIcon.Cn9REWmg.js";import{p as xt}from"../chunks/stores.P_GwiDr4.js";import{j as bt}from"../chunks/index.BW_8NPNh.js";import{L as $t}from"../chunks/loader._7BlSOYJ.js";const ht=async(G,g,y,L,j,x,w)=>{s(g,!0),s(y,"");const _=await L.runProjectSQL2(j,{qstr:t(x),data:[]});if(_.status!==200){s(y,`${_.statusText} : ${_.data}`);return}s(w,p(_.data)),s(g,!1)};var yt=v("<span>SQL</span>"),jt=v("<span>Formatter</span>"),wt=v("<!> <!>",1),Ct=v('<div class="min-h-10 card overflow-auto resize-y p-2"><!></div>'),Lt=v('<td class="p-1"> </td>'),St=v('<td class="p-1"><!></td>'),qt=v("<tr></tr>"),Pt=v('<table class="table table-hover overflow-auto input"><thead><tr></tr></thead><tbody></tbody></table>'),Qt=v('<div class="flex flex-col w-full h-[94vh] p-2"><div class="w-full p-2 rounded"><!></div> <div><p class="text-red-600"> </p></div> <div class="flex justify-end p-2"><button class="inline-flex rounded p.05 variant-filled self-center"><!> Run</button></div> <div class="w-full h-full p-2 overflow-auto"><!></div></div>');function Ut(G,g){st(g,!0);const y=ut(),L=()=>ft(xt,"$page",y);let j=vt(g,"isReadOnly",3,!1),x=m("select * from __project__Accounts;"),w=m(`const process = (ctx) => {
}`);const _=L().params.pid,V=ot("__api__");let C=m(p([])),I=m(!1),b=m("sql"),N=m(""),$=m(p([]));lt(()=>{if(!t(C).length){s($,p([]));return}const o=Object.keys(t(C)[0]);o.includes("id")?s($,p(["id",...o.filter(f=>f!=="id")])):s($,p(o))});var S=Qt(),q=r(S),W=r(q);ct(W,{children:(o,f)=>{var n=wt(),c=dt(n);K(c,{get group(){return t(b)},set group(a){s(b,p(a))},name:"SQL Query",value:"sql",children:(a,d)=>{var l=yt();i(a,l)},$$slots:{default:!0}});var u=h(c,2);K(u,{get group(){return t(b)},set group(a){s(b,p(a))},name:"Formatter",value:"formatter",children:(a,d)=>{var l=jt();i(a,l)},$$slots:{default:!0}}),i(o,n)},$$slots:{default:!0,panel:(o,f)=>{var n=Ct(),c=r(n);M(c,()=>t(b)==="sql",u=>{var a=A(()=>mt({dialect:_t}));U(u,{get value(){return t(x)},set value(d){s(x,p(d))},get readonly(){return j()},extensions:[],get lang(){return t(a)}})},u=>{var a=A(bt);U(u,{get value(){return t(w)},set value(d){s(w,p(d))},get readonly(){return j()},extensions:[],get lang(){return t(a)}})}),e(n),i(o,n)}}}),e(q);var P=h(q,2),B=r(P),X=r(B);e(B),e(P);var Q=h(P,2),k=r(Q);k.__click=[ht,I,N,V,_,x,C];var Y=r(k);gt(Y,{className:"w-4 h-4 mt-1",name:"play"}),it(),e(k),e(Q);var E=h(Q,2),Z=r(E);M(Z,()=>t(I),o=>{$t(o,{})},o=>{var f=Pt(),n=r(f),c=r(n);z(c,21,()=>t($),D,(a,d)=>{var l=Lt(),T=r(l);e(l),R(()=>F(T,t(d))),i(a,l)}),e(c),e(n);var u=h(n);z(u,21,()=>t(C),D,(a,d)=>{var l=qt();z(l,21,()=>t($),D,(T,H)=>{var O=St();const tt=A(()=>t(d)[t(H)]);var at=r(O);M(at,()=>t(tt)!==void 0,et=>{var J=rt();R(()=>F(J,t(d)[t(H)])),i(et,J)}),e(O),i(T,O)}),e(l),i(a,l)}),e(u),e(f),i(o,f)}),e(E),e(S),R(()=>F(X,t(N))),i(G,S),nt()}pt(["click"]);export{Ut as component};
