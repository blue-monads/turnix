import"../chunks/disclose-version.BOfWGsg9.js";import{p as u}from"../chunks/proxy.DKs42Gzo.js";import{p as _,g as h,l as x,a as F,b as a,d as b,i as y,h as l}from"../chunks/runtime.CmPLUjJJ.js";import{p as f}from"../chunks/props.IjRlZquB.js";import{s as v,a as w}from"../chunks/store.BN9B0iWR.js";import{p as z}from"../chunks/index.CqTR9n69.js";import{F as C}from"../chunks/FileListings.B4T3VrXQ.js";function B(n,e){_(e,!0);const m=v(),c=()=>w(z,"$params",m);let r=f(e,"files",31,()=>u([])),o=f(e,"selected",15);const g=h("__api__");let p=b(()=>c().folder||""),t=y(!1);const d=async s=>{l(t,!0);const i=await g.listSelfFiles(s);i.status===200&&(r(i.data),l(t,!1))};x(()=>{d(a(p))}),C(n,{get selected(){return o()},set selected(s){o(s)},baseUrl:"/z/pages/portal/self/files",get path(){return a(p)},get files(){return r()},get loading(){return a(t)}}),F()}export{B as component};
