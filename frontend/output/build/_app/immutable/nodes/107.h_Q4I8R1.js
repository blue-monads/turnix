import{a as c,t as m}from"../chunks/disclose-version.BOfWGsg9.js";import{p as D,g as E,i as F,f as P,a as G,h as H,c as a,e as b,r as t,s as r,b as p,t as J}from"../chunks/runtime.CmPLUjJJ.js";import{s as x}from"../chunks/render.CZQ66Y_F.js";import{e as K,i as L}from"../chunks/each.BJPUi07-.js";import{d as O}from"../chunks/events.DnYA5zBv.js";import{p as S}from"../chunks/proxy.DKs42Gzo.js";import{S as d}from"../chunks/SvgIcon.Cn9REWmg.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as Q}from"../chunks/AppBar.pSMxueMf.js";var R=m('<h4 class="h4">Inbox</h4>'),V=m('<button class="btn btn-sm variant-filled"><!> Clear</button> <button class="btn btn-sm variant-filled-secondary"><!> refresh</button>',1),W=m('<tr class="bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100"><td class="p-2 whitespace-nowrap text-sm font-medium text-gray-900">1</td><td class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"> </td><td class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"><span class="chip variant-filled-secondary"> </span></td><td class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"> </td><td class="p-2 text-sm text-gray-900 font-light whitespace-nowrap"></td></tr>'),X=m('<!> <div class="flex flex-col"><div class="overflow-auto"><div class="inline-block min-w-full p-2 "><table class="min-w-full border-gray-300 border"><thead class="bg-gray-200 border-b"><tr><th scope="col" class="text-sm font-medium text-gray-900 p-2 text-left">#</th><th scope="col" class="text-sm font-medium text-gray-900 p-2 text-left">Title</th><th scope="col" class="text-sm font-medium text-gray-900 p-2 text-left">Type</th><th scope="col" class="text-sm font-medium text-gray-900 p-2 text-left">Content</th><th scope="col" class="text-sm font-medium text-gray-900 p-2 text-left"></th></tr></thead><tbody></tbody></table> <div class="flex justify-between py-2"><button class="btn btn-sm variant-ghost-secondary"><!></button> <button class="btn btn-sm variant-ghost-secondary"><!></button></div></div></div></div>',1);function lt(T,j){D(j,!0);const z=E("__api__");let g=F(S([]));const u=async()=>{const s=await z.listUserMessages();s.status===200&&H(g,S(s.data))};u();var _=X(),y=P(_);Q(y,{$$slots:{lead:(s,l)=>{var e=R();c(s,e)},trail:(s,l)=>{var e=V(),o=P(e),h=a(o);d(h,{className:"w-4 h-4",name:"ellipsis-horizontal"}),b(),t(o);var n=r(o,2);n.__click=u;var i=a(n);d(i,{className:"w-4 h-4",name:"arrow-path"}),b(),t(n),c(s,e)}}});var w=r(y,2),$=a(w),N=a($),v=a(N),k=r(a(v));K(k,21,()=>p(g),L,(s,l)=>{var e=W(),o=r(a(e)),h=a(o);t(o);var n=r(o),i=a(n),U=a(i);t(i),t(n);var I=r(n),q=a(I);t(I),b(),t(e),J(()=>{x(h,p(l).title||""),x(U,p(l).type||""),x(q,p(l).contents||"")}),c(s,e)}),t(k),t(v);var C=r(v,2),f=a(C),B=a(f);d(B,{className:"h-4 w-4",name:"arrow-left"}),t(f);var A=r(f,2),M=a(A);d(M,{className:"h-4 w-4",name:"arrow-right"}),t(A),t(C),t(N),t($),t(w),c(T,_),G()}O(["click"]);export{lt as component};
