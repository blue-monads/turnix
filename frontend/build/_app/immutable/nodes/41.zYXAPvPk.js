import{s as b,w as u,i as h,f as E,k as d,r as $,n as L}from"../chunks/scheduler.D2OtNeTK.js";import{S as N,i as X,a as l,d as w,t as _,g as A,c as y,b as k,m as g,e as T}from"../chunks/index.DD9KLV65.js";import{g as C}from"../chunks/entry.C8cUSds7.js";import{p as S}from"../chunks/stores.D0bQ952I.js";import{N as G}from"../chunks/index.ZlEQwogg.js";import{A as I}from"../chunks/auto_form.FPf_zM_P.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B6Y8wNrO.js";import{L as O}from"../chunks/loader.zv9XVPa1.js";import{p as P}from"../chunks/index.BURjlDwe.js";function q(s){let e,n;return e=new I({props:{data:s[1],schema:{name:"Edit Contact",fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Info",ftype:"LONG_TEXT",key_name:"info"},{name:"Type",ftype:"SELECT",key_name:"ctype",options:["vendor","client"]},{name:"Email",ftype:"EMAIL",key_name:"email"},{name:"Phone",ftype:"TEXT",key_name:"phone"},{name:"Second phone",ftype:"TEXT",key_name:"phone2"},{name:"Third phone",ftype:"TEXT",key_name:"phone3"},{name:"Address",ftype:"LONG_TEXT",key_name:"address"},{name:"Second Address",ftype:"LONG_TEXT",key_name:"address2"}],required_fields:["name"]},onSave:s[4]}}),{c(){y(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){g(e,t,a),n=!0},p(t,a){const m={};a&2&&(m.data=t[1]),e.$set(m)},i(t){n||(_(e.$$.fragment,t),n=!0)},o(t){l(e.$$.fragment,t),n=!1},d(t){T(e,t)}}}function j(s){let e,n;return e=new O({}),{c(){y(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){g(e,t,a),n=!0},p:L,i(t){n||(_(e.$$.fragment,t),n=!0)},o(t){l(e.$$.fragment,t),n=!1},d(t){T(e,t)}}}function x(s){let e,n,t,a;const m=[j,q],r=[];function f(o,c){return o[0]?0:1}return e=f(s),n=r[e]=m[e](s),{c(){n.c(),t=u()},l(o){n.l(o),t=u()},m(o,c){r[e].m(o,c),h(o,t,c),a=!0},p(o,[c]){let i=e;e=f(o),e===i?r[e].p(o,c):(A(),l(r[i],1,1,()=>{r[i]=null}),w(),n=r[e],n?n.p(o,c):(n=r[e]=m[e](o),n.c()),_(n,1),n.m(t.parentNode,t))},i(o){a||(_(n),a=!0)},o(o){l(n),a=!1},d(o){o&&E(t),r[e].d(o)}}}function z(s,e,n){let t,a;d(s,P,p=>n(5,t=p)),d(s,S,p=>n(6,a=p));const m=a.params.pid,r=G($("__api__")),f=t.cid;let o=!0,c;return(async()=>{n(0,o=!0);const p=await r.getContact(m,f);p.status===200&&(n(1,c=p.data),n(0,o=!1))})(),[o,c,m,r,async p=>{(await r.addContact(m,p)).status===200&&C(`/z/pages/portal/projects/books/${m}/contacts`)}]}class V extends N{constructor(e){super(),X(this,e,z,x,b,{})}}export{V as component};