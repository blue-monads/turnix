import{a as f,t as h}from"../chunks/disclose-version.BOfWGsg9.js";import{p as w,g as x,i as u,f as C,a as S,h as o,b as i,s as g,c as v,r as _,e as B,d as R}from"../chunks/runtime.CmPLUjJJ.js";import{i as j}from"../chunks/if.DoU_1wcm.js";import{s as q}from"../chunks/attributes.yGGqmlPO.js";import{p as b}from"../chunks/proxy.DKs42Gzo.js";import{s as z,a as I}from"../chunks/store.BN9B0iWR.js";import{N as L}from"../chunks/index.Dq_7tmzI.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B5GFseBd.js";import{A as P}from"../chunks/AppBar.DRZCjzsx.js";import{p as D}from"../chunks/stores.DkdCKZTz.js";import"../chunks/entry.CyBKS7ox.js";import"../chunks/index.DnNofn7z.js";import{A as E}from"../chunks/autotable.D5y0wA8l.js";import{L as F}from"../chunks/loader._7BlSOYJ.js";var G=h('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li><a class="anchor">Reports</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Saved</li></ol>'),H=h("<!> <!>",1);function sa($,y){w(y,!0);const A=z(),s=I(D,"$page",A).params.pid,p=L(x("__api__"));let e=u(!0),l=u(b([]));const m=async()=>{o(e,!0);const a=await p.listReportSaved(s);a.status===200&&(o(l,b(a.data)),o(e,!1))};m();var n=H(),c=C(n);P(c,{$$slots:{lead:(a,d)=>{var r=G(),t=g(v(r),4),N=v(t);q(N,"href",`/z/pages/portal/projects/books/${s}/report`),_(t),B(4),_(r),f(a,r)}}});var k=g(c,2);j(k,()=>i(e),a=>{F(a,{})},a=>{var d=R(()=>[{Name:"view",Class:"variant-filled-primary",Action:async(r,t)=>{}},{Name:"delete",Class:"variant-filled-error",Action:async r=>{(await p.deleteReportSaved(s,r)).status===200&&m()}}]);E(a,{action_key:"id",key_names:[["id","ID"],["name","Name"],["created_at","Created At"]],hashSeed:35,get datas(){return i(l)},color:["report_type"],get actions(){return i(d)}})}),f($,n),S()}export{sa as component};