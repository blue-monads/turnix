import{c as O,a as _,t as N}from"../chunks/disclose-version.BOfWGsg9.js";import{p as Q,g as U,f as x,a as V,h as l,b as m,i as y,s as p,c as r,r as t,e as b,d as W}from"../chunks/runtime.CmPLUjJJ.js";import{d as X}from"../chunks/events.DnYA5zBv.js";import{i as Y}from"../chunks/if.DoU_1wcm.js";import{s as Z}from"../chunks/attributes.yGGqmlPO.js";import{p as R}from"../chunks/proxy.DKs42Gzo.js";import{s as aa,a as T}from"../chunks/store.BN9B0iWR.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B5GFseBd.js";import{A as ra}from"../chunks/AppBar.DRZCjzsx.js";import{C as ta}from"../chunks/index.63AqBVKj.js";import{h as oa}from"../chunks/index.DtK7kxBu.js";import{p as ea}from"../chunks/stores.DkdCKZTz.js";import{N as sa}from"../chunks/index.Dq_7tmzI.js";import{S as C}from"../chunks/SvgIcon.Cn9REWmg.js";import{R as la}from"../chunks/RenderBox.BJeVydrf.js";import{p as ia}from"../chunks/index.COSOhg_g.js";import{L as na}from"../chunks/loader._7BlSOYJ.js";var ma=N('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="true">&rsaquo;</li> <li><a class="anchor">Reports</a></li> <li class="crumb-separator" aria-hidden="true">&rsaquo;</li> <li>Editor</li></ol>'),pa=async(S,d,c,g,h,o)=>{l(d,!0),await c.updateReportTemplate(g,h,{template:m(o)}),l(d,!1)},da=N('<button class="btn variant-filled btn-sm"><!> Run</button> <button class="btn variant-filled-secondary btn-sm"><!> Print</button> <button class="btn variant-filled-tertiary btn-sm"><!> Save</button>',1),ca=N('<!> <div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full overflow-auto"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><!></div></div>',1);function Pa(S,d){Q(d,!0);const c=aa(),g=()=>T(ea,"$page",c),h=()=>T(ia,"$params",c);let o=y(""),A=y("");const v=g().params.pid,B=h().template_id,P=U("__api__"),q=sa(P);let u=y(!0);(async()=>{l(u,!0);const i=await q.getReportTemplate(v,B);i.status===200&&(l(o,R(i.data.template)),l(u,!1))})();let a;var j=O(),E=x(j);Y(E,()=>m(u),i=>{na(i,{})},i=>{var z=ca(),H=x(z);ra(H,{$$slots:{lead:(e,G)=>{var n=ma(),s=p(r(n),4),$=r(s);Z($,"href",`/z/pages/portal/projects/books/${v}/report`),t(s),b(4),t(n),_(e,n)},trail:(e,G)=>{var n=da(),s=x(n);s.__click=()=>{l(A,R(m(o))),a==null||a.reload()};var $=r(s);C($,{className:"w-4 h-4",name:"play"}),b(),t(s);var f=p(s,2);f.__click=()=>{a==null||a.print()};var J=r(f);C(J,{className:"w-4 h-4",name:"printer"}),b(),t(f);var k=p(f,2);k.__click=[pa,u,q,v,B,o];var K=r(k);C(K,{className:"w-4 h-4",name:"arrow-down-on-square"}),b(),t(k),_(e,n)}}});var I=p(H,2),w=r(I),M=r(w),D=W(oa);ta(M,{get value(){return m(o)},set value(e){l(o,R(e))},get lang(){return m(D)}}),t(w);var L=p(w,2),F=r(L);la(F,{pid:v,get htmlSource(){return m(A)},rootApi:P,get handle(){return a},set handle(e){a=e}}),t(L),t(I),_(i,z)}),_(S,j),V()}X(["click"]);export{Pa as component};