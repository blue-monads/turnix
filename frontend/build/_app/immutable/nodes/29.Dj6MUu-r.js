import{s as r,n as p,k as i,z as c}from"../chunks/scheduler.DVTwEvQJ.js";import{S as f,i as _,c as d,b as y,m as u,t as T,a as l,e as g}from"../chunks/index.DkBQfKoF.js";import{g as k}from"../chunks/entry.B0930nEj.js";import{p as $}from"../chunks/stores.10CBQxEg.js";import{N as E}from"../chunks/index.D3jOKT1N.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Dy7ATQYE.js";import{A as h}from"../chunks/auto_form.CafTK1tp.js";/* empty css                                                     */function A(o){let n,t;return n=new h({props:{data:{},schema:{name:"Add Contact",fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Info",ftype:"LONG_TEXT",key_name:"info"},{name:"Type",ftype:"SELECT",key_name:"ctype",options:["vendor","client"]},{name:"Email",ftype:"EMAIL",key_name:"email"},{name:"Phone",ftype:"TEXT",key_name:"phone"},{name:"Second phone",ftype:"TEXT",key_name:"phone2"},{name:"Third phone",ftype:"TEXT",key_name:"phone3"},{name:"Address",ftype:"LONG_TEXT",key_name:"address"},{name:"Address2",ftype:"LONG_TEXT",key_name:"address2"}],required_fields:["name"]},onSave:o[2]}}),{c(){d(n.$$.fragment)},l(e){y(n.$$.fragment,e)},m(e,a){u(n,e,a),t=!0},p,i(e){t||(T(n.$$.fragment,e),t=!0)},o(e){l(n.$$.fragment,e),t=!1},d(e){g(n,e)}}}function X(o,n,t){let e;i(o,$,s=>t(3,e=s));const a=e.params.pid,m=E(c("__api__"));return[a,m,async s=>{(await m.addContact(a,s)).status===200&&k(`/z/pages/portal/projects/books/${a}/contacts`)}]}class O extends f{constructor(n){super(),_(this,n,X,A,r,{})}}export{O as component};