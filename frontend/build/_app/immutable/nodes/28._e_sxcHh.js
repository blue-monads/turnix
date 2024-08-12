import{s as b,A as u,i as h,f as E,k as d,z as $,n as A}from"../chunks/scheduler.DVTwEvQJ.js";import{S as L,i as N,a as l,d as X,t as _,g as C,c as y,b as k,m as g,e as T}from"../chunks/index.DkBQfKoF.js";import{g as S}from"../chunks/entry.B0930nEj.js";import{p as w}from"../chunks/stores.10CBQxEg.js";import{N as G}from"../chunks/index.D3jOKT1N.js";import{A as I}from"../chunks/auto_form.CafTK1tp.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Dy7ATQYE.js";import{L as O}from"../chunks/loader.B70xSjYw.js";import{p as P}from"../chunks/index.B_nN5hbX.js";function q(s){let e,n;return e=new I({props:{data:s[1],schema:{name:"Edit Contact",fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Info",ftype:"LONG_TEXT",key_name:"info"},{name:"Type",ftype:"SELECT",key_name:"ctype",options:["vendor","client"]},{name:"Email",ftype:"EMAIL",key_name:"email"},{name:"Phone",ftype:"TEXT",key_name:"phone"},{name:"Second phone",ftype:"TEXT",key_name:"phone2"},{name:"Third phone",ftype:"TEXT",key_name:"phone3"},{name:"Address",ftype:"LONG_TEXT",key_name:"address"},{name:"Second Address",ftype:"LONG_TEXT",key_name:"address2"}],required_fields:["name"]},onSave:s[4]}}),{c(){y(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){g(e,t,a),n=!0},p(t,a){const m={};a&2&&(m.data=t[1]),e.$set(m)},i(t){n||(_(e.$$.fragment,t),n=!0)},o(t){l(e.$$.fragment,t),n=!1},d(t){T(e,t)}}}function z(s){let e,n;return e=new O({}),{c(){y(e.$$.fragment)},l(t){k(e.$$.fragment,t)},m(t,a){g(e,t,a),n=!0},p:A,i(t){n||(_(e.$$.fragment,t),n=!0)},o(t){l(e.$$.fragment,t),n=!1},d(t){T(e,t)}}}function j(s){let e,n,t,a;const m=[z,q],r=[];function f(o,c){return o[0]?0:1}return e=f(s),n=r[e]=m[e](s),{c(){n.c(),t=u()},l(o){n.l(o),t=u()},m(o,c){r[e].m(o,c),h(o,t,c),a=!0},p(o,[c]){let i=e;e=f(o),e===i?r[e].p(o,c):(C(),l(r[i],1,1,()=>{r[i]=null}),X(),n=r[e],n?n.p(o,c):(n=r[e]=m[e](o),n.c()),_(n,1),n.m(t.parentNode,t))},i(o){a||(_(n),a=!0)},o(o){l(n),a=!1},d(o){o&&E(t),r[e].d(o)}}}function x(s,e,n){let t,a;d(s,P,p=>n(5,t=p)),d(s,w,p=>n(6,a=p));const m=a.params.pid,r=G($("__api__")),f=t.cid;let o=!0,c;return(async()=>{n(0,o=!0);const p=await r.getContact(m,f);p.status===200&&(n(1,c=p.data),n(0,o=!1))})(),[o,c,m,r,async p=>{(await r.addContact(m,p)).status===200&&S(`/z/pages/portal/projects/books/${m}/contacts`)}]}class V extends L{constructor(e){super(),N(this,e,x,j,b,{})}}export{V as component};
