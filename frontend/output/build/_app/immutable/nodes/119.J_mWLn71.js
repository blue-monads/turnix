import{a as t,t as r}from"../chunks/disclose-version.BOfWGsg9.js";import{p as u,f as g,a as _,b as y,c as n,e as h,r as l,s as x,ae as A}from"../chunks/runtime.CmPLUjJJ.js";import{s as $}from"../chunks/attributes.yGGqmlPO.js";import{i as b}from"../chunks/lifecycle.46CgxKeU.js";import{g as S}from"../chunks/entry.CytXzrzW.js";import{A as z}from"../chunks/autotable.D5y0wA8l.js";import{S as N}from"../chunks/SvgIcon.Cn9REWmg.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as w}from"../chunks/AppBar.pSMxueMf.js";var C=r('<div class="flex flex-wrap justify-end gap-2"><a class="btn variant-filled btn-sm"><!> peer</a></div>'),P=r("<h4>Peers</h4>"),k=r("<!> <!>",1);function H(p,m){u(m,!1),b();var o=k(),i=g(o);w(i,{children:(a,c)=>{var s=C(),e=n(s);$(e,"href","/z/pages/startpage/peers/new");var v=n(e);N(v,{className:"h-4 w-4",name:"plus"}),h(),l(e),l(s),t(a,s)},$$slots:{default:!0,lead:(a,c)=>{var s=P();t(a,s)}}});var d=x(i,2),f=A(()=>[{Action:async a=>{},Name:"explore",Class:"variant-filled-primary"},{Action:async a=>{S(`/z/pages/startpage/peers/info?id=${a}`)},Name:"info",Class:"variant-filled-secondary"},{Action:async a=>{},Name:"remove",Class:"variant-filled-error"}]);z(d,{get actions(){return y(f)},color:["status"],key_names:[["id","ID"],["lastSeen","Last seen"],["status","Status"]],action_key:"id",datas:[{id:"xx-yy-zz",lastSeen:"2022-12-31",status:"online"}]}),t(p,o),_()}export{H as component};
