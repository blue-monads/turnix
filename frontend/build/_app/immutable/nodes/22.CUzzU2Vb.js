import{s as C,a as x,z as w,g as L,i as _,f as $,k as B,y as N,e as z,c as H,r as I,m as y,n as g}from"../chunks/scheduler.CPmtgLt4.js";import{S as M,i as P,c as d,b,m as h,a as p,e as S,t as m,d as k,g as T}from"../chunks/index.CsITmn2N.js";import{A as j}from"../chunks/Account.BLR84kXS.js";import{N as q}from"../chunks/index.17k-MAAo.js";import"../chunks/index.COtd9eM3.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.XpbzWZVH.js";import{A as D}from"../chunks/AppBar.D1f81ABc.js";import{L as V}from"../chunks/loader.DJ5aES98.js";import{g as E}from"../chunks/entry.CNB9SHga.js";import{p as F}from"../chunks/stores.DZzzPzvs.js";function G(n){let t,r='<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li class="crumb">Add Account</li></ol>';return{c(){t=z("div"),t.innerHTML=r,this.h()},l(e){t=H(e,"DIV",{slot:!0,class:!0,"data-svelte-h":!0}),I(t)!=="svelte-1v4xnye"&&(t.innerHTML=r),this.h()},h(){y(t,"slot","lead"),y(t,"class","flex gap-2")},m(e,o){_(e,t,o)},p:g,d(e){e&&$(t)}}}function J(n){let t,r;return t=new j({props:{onChange:n[1]}}),{c(){d(t.$$.fragment)},l(e){b(t.$$.fragment,e)},m(e,o){h(t,e,o),r=!0},p:g,i(e){r||(m(t.$$.fragment,e),r=!0)},o(e){p(t.$$.fragment,e),r=!1},d(e){k(t,e)}}}function K(n){let t,r;return t=new V({}),{c(){d(t.$$.fragment)},l(e){b(t.$$.fragment,e)},m(e,o){h(t,e,o),r=!0},p:g,i(e){r||(m(t.$$.fragment,e),r=!0)},o(e){p(t.$$.fragment,e),r=!1},d(e){k(t,e)}}}function O(n){let t,r,e,o,l,i;t=new D({props:{$$slots:{lead:[G]},$$scope:{ctx:n}}});const f=[K,J],s=[];function A(a,c){return a[0]?0:1}return e=A(n),o=s[e]=f[e](n),{c(){d(t.$$.fragment),r=x(),o.c(),l=w()},l(a){b(t.$$.fragment,a),r=L(a),o.l(a),l=w()},m(a,c){h(t,a,c),_(a,r,c),s[e].m(a,c),_(a,l,c),i=!0},p(a,[c]){const v={};c&32&&(v.$$scope={dirty:c,ctx:a}),t.$set(v);let u=e;e=A(a),e===u?s[e].p(a,c):(T(),p(s[u],1,1,()=>{s[u]=null}),S(),o=s[e],o?o.p(a,c):(o=s[e]=f[e](a),o.c()),m(o,1),o.m(l.parentNode,l))},i(a){i||(m(t.$$.fragment,a),m(o),i=!0)},o(a){p(t.$$.fragment,a),p(o),i=!1},d(a){a&&($(r),$(l)),k(t,a),s[e].d(a)}}}function Q(n,t,r){let e;B(n,F,s=>r(2,e=s));const o=e.params.pid,l=q(N("__api__"));let i=!1;return[i,async s=>{r(0,i=!0),await l.addAccount(o,s),E(location.pathname.replace("/new","")),r(0,i=!1)}]}class oe extends M{constructor(t){super(),P(this,t,Q,O,C,{})}}export{oe as component};
