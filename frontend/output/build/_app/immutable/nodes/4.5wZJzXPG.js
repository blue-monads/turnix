import{a as i,t as v,c as k}from"../chunks/disclose-version.BOfWGsg9.js";import{p as W,g as X,f as c,a as Y,s as n,c as r,r as e,b as d,d as H,e as I,t as b,n as Z}from"../chunks/runtime.CmPLUjJJ.js";import{s as M}from"../chunks/render.CZQ66Y_F.js";import{i as U}from"../chunks/if.DoU_1wcm.js";import{k as q}from"../chunks/key.B3I6Q0YW.js";import{e as aa,i as ra}from"../chunks/each.BJPUi07-.js";import{s as ea}from"../chunks/snippet.UtI6fqho.js";import{s as y}from"../chunks/attributes.yGGqmlPO.js";import{d as sa}from"../chunks/events.DnYA5zBv.js";import{s as oa,a as D}from"../chunks/store.BN9B0iWR.js";import{r as ta}from"../chunks/legacy-client.BxBrBVYa.js";import{p as la}from"../chunks/index.NL8p-q19.js";import{S as E}from"../chunks/SvgIcon.Cn9REWmg.js";import{g as ia}from"../chunks/stores.C0Vv8-i4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as pa}from"../chunks/AppBar.pSMxueMf.js";import{p as na}from"../chunks/stores.BgpLsI8J.js";var ma=v('<h4 class="h4">Project Files</h4>'),ca=(N,_,g)=>{_.trigger({type:"component",component:"new_folder_panel",meta:{onNewName:g}})},da=v('<div class="flex flex-row gap-1"><a class="btn btn-sm variant-filled"><!> <span class="hidden md:inline">Upload</span></a> <button class="btn btn-sm variant-filled-primary"><!> <span class="hidden md:inline">Folder</span></button></div>'),fa=v('<li class="crumb-separator" aria-hidden="">/</li> <li class="crumb"><a class="anchor"> </a></li>',1),va=v('<li class="crumb-separator" aria-hidden="">/</li> <li class="crumb"> </li>',1),_a=v('<!> <div class="flex justify-between p-2 pl-4"><ol class="breadcrumb"><li class="crumb"><a class="anchor">Home</a></li> <!> <!></ol></div> <div class="px-2 h-full"><div class="card p-2 h-full"><!></div></div>',1);function Ca(N,_){W(_,!0);const g=oa(),G=()=>D(na,"$page",g),f=()=>D(la,"$params",g),h=G().params.pid,J=ia(),K=X("__api__");let P=H(()=>(f().folder||"").split("/")),$=H(()=>f().filename),u;ta(()=>{u=1});const L=async a=>{await K.addProjectFolder(h,a,f().folder),u=u+1};var z=_a(),F=c(z);pa(F,{$$slots:{lead:(a,t)=>{var s=ma();i(a,s)},trail:(a,t)=>{var s=k(),m=c(s);U(m,()=>!f().file,p=>{var o=da(),l=r(o),T=r(l);E(T,{name:"cloud-arrow-up",className:"h-4 w-4"}),I(2),e(l);var j=n(l,2);j.__click=[ca,J,L];var V=r(j);E(V,{name:"folder-plus",className:"h-4 w4"}),I(2),e(j),e(o),b(()=>y(l,"href",`/z/pages/portal/project/files/${h??""}/upload?folder=${(f().folder||"")??""}`)),i(p,o)}),i(a,s)}}});var x=n(F,2),S=r(x),w=r(S),O=r(w);y(O,"href",`/z/pages/portal/project/files/${h??""}`),e(w);var A=n(w,2);aa(A,17,()=>d(P),ra,(a,t,s)=>{var m=fa(),p=n(c(m),2),o=r(p);b(()=>y(o,"href",`/z/pages/portal/project/files/${h??""}?folder=${d(P).slice(0,s+1).join("/")??""}`));var l=r(o);e(o),e(p),b(()=>M(l,d(t)||"")),i(a,m)});var Q=n(A,2);q(Q,()=>d($),a=>{var t=k(),s=c(t);U(s,()=>d($),m=>{var p=va(),o=n(c(p),2),l=r(o);e(o),b(()=>M(l,d($))),i(m,p)}),i(a,t)}),e(S),e(x);var B=n(x,2),C=r(B),R=r(C);q(R,()=>u,a=>{var t=k(),s=c(t);ea(s,()=>_.children??Z),i(a,t)}),e(C),e(B),i(N,z),Y()}sa(["click"]);export{Ca as component};
