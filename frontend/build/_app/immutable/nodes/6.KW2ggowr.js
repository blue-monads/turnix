import{a as i,t as v,c as w}from"../chunks/disclose-version.wcsHG7oF.js";import{p as O,g as Q,f as d,a as R,s as m,c as s,b as c,r as o,d as P,e as B,t as u,n as T}from"../chunks/runtime.D6lgFUio.js";import{s as C}from"../chunks/render.CIW-w3Th.js";import{i as H}from"../chunks/if.lkQ-IRrr.js";import{k as I}from"../chunks/key.Dx7O0vkj.js";import{e as V,i as W}from"../chunks/each.CWrzZiCN.js";import{s as X}from"../chunks/snippet.CDCY8qbR.js";import{s as M}from"../chunks/attributes.BFGPVQjG.js";import{d as Y}from"../chunks/utils.Dh-XvBAd.js";import{s as Z,a as aa}from"../chunks/store.BaapSzAp.js";import{r as ra}from"../chunks/legacy-client.B7AzNrU7.js";import{p as ea}from"../chunks/index.xGEmn6Hr.js";import{S as U}from"../chunks/SvgIcon.Cj01WnO8.js";import{g as sa}from"../chunks/stores.CebXFGJ4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as oa}from"../chunks/AppBar.3srT6C5i.js";import"../chunks/entry.Y6HcnCx3.js";import"../chunks/_kv_editor.SKGkUTMX.js";var ta=v('<h4 class="h4">Files</h4>'),la=(k,_,g)=>{_.trigger({type:"component",component:"new_folder_panel",meta:{onNewName:g}})},ia=v('<div class="flex flex-row gap-1"><a class="btn btn-sm variant-filled"><!> <span class="hidden md:inline">Upload</span></a> <button class="btn btn-sm variant-filled-primary"><!> <span class="hidden md:inline">Folder</span></button></div>'),na=v('<li class="crumb-separator" aria-hidden="">/</li> <li class="crumb"><a class="anchor"> </a></li>',1),ma=v('<li class="crumb-separator" aria-hidden="">/</li> <li class="crumb"> </li>',1),pa=v('<!> <div class="flex justify-between p-2 pl-4"><ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/self/files">Home</a></li> <!> <!></ol></div> <div class="px-2 h-full"><div class="card p-2 h-full"><!></div></div>',1);function ja(k,_){O(_,!0);const g=Z(),f=()=>aa(ea,"$params",g),q=sa(),D=Q("__api__");let y=P(()=>(f().folder||"").split("/")),h;ra(()=>{h=1});let b=P(()=>f().filename);const E=async a=>{await D.addSelfFolder(f().folder||"",a),h=h+1};var N=pa(),S=d(N);oa(S,{$$slots:{lead:(a,t)=>{var r=ta();i(a,r)},trail:(a,t)=>{var r=w(),p=d(r);H(p,()=>!f().file,n=>{var e=ia(),l=s(e),K=s(l);U(K,{name:"cloud-arrow-up",className:"h-4 w-4"}),B(2),o(l);var $=m(l,2);$.__click=[la,q,E];var L=s($);U(L,{name:"folder-plus",className:"h-4 w4"}),B(2),o($),o(e),u(()=>M(l,"href",`/z/pages/portal/self/files/upload?folder=${(f().folder||"")??""}`)),i(n,e)}),i(a,r)}}});var x=m(S,2),z=s(x),F=m(s(z),2);V(F,17,()=>c(y),W,(a,t,r)=>{var p=na(),n=m(d(p),2),e=s(n);u(()=>M(e,"href",`/z/pages/portal/self/files?folder=${c(y).slice(0,r+1).join("/")??""}`));var l=s(e);o(e),o(n),u(()=>C(l,c(t))),i(a,p)});var G=m(F,2);I(G,()=>c(b),a=>{var t=w(),r=d(t);H(r,()=>c(b),p=>{var n=ma(),e=m(d(n),2),l=s(e);o(e),u(()=>C(l,c(b))),i(p,n)}),i(a,t)}),o(z),o(x);var j=m(x,2),A=s(j),J=s(A);I(J,()=>h,a=>{var t=w(),r=d(t);X(r,()=>_.children??T),i(a,t)}),o(A),o(j),i(k,N),R()}Y(["click"]);export{ja as component};
