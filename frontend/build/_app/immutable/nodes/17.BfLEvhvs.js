import{a as f,t as u}from"../chunks/disclose-version.wcsHG7oF.js";import{p as x,s as e,f as _,a as k,c as r,r as a,b as g,t as y}from"../chunks/runtime.D6lgFUio.js";import{s as z}from"../chunks/render.CIW-w3Th.js";import{e as S,i as j}from"../chunks/each.CWrzZiCN.js";import{d as q}from"../chunks/utils.Dh-XvBAd.js";import{i as D}from"../chunks/lifecycle.ZPOFI82y.js";import"../chunks/auto_form.CxbPX2mo.js";import{g as I}from"../chunks/stores.CebXFGJ4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{S as M}from"../chunks/SvgIcon.Cj01WnO8.js";var N=(n,t)=>{t.trigger({type:"component",component:"file_picker",meta:{}})},P=u('<span class="flex gap-1 justify-start items-center hover:bg-zinc-100 p-1 rounded-lg"><!> <p> </p></span>'),A=u('<h4 class="h4">Playground</h4> <button class="btn btn-sm variant-filled">click me</button> <button class="relative group transition-all duration-200 focus:overflow-visible w-max h-max p-1 overflow-hidden flex flex-row items-center justify-center bg-white gap-2 rounded-lg border border-zinc-200"><span>Dropdown</span> <div class="absolute shadow-lg top-10 left-0 w-32 h-max p-1 bg-white border border-zinc-200 rounded-lg flex flex-col gap-2"></div></button>',1);function Q(n,t){x(t,!1);const v=I(),b=[{name:"rename",icon:"pencil-square"},{name:"download",icon:"arrow-down-on-square"},{name:"delete",icon:"trash"}];D();var s=A(),i=e(_(s),2);i.__click=[N,v];var l=e(i,2),p=e(r(l),2);S(p,5,()=>b,j,(h,c)=>{var o=P(),m=r(o);M(m,{get name(){return g(c).icon},className:"w-4 h-4"});var d=e(m,2),w=r(d);a(d),a(o),y(()=>z(w,g(c).name)),f(h,o)}),a(p),a(l),f(n,s),k()}q(["click"]);export{Q as component};