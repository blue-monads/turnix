import{s as O,k as D,y as F,e as h,c as _,r as G,m,i as u,n as P,f as o,t as j,a as H,b as L,d as R,g as I,h as x}from"../chunks/scheduler.CPmtgLt4.js";import{S as J,i as K,c as b,b as v,m as w,t as A,a as k,d as y}from"../chunks/index.CsITmn2N.js";import{S as E}from"../chunks/SvgIcon.sHYuI0dM.js";import{N as Q}from"../chunks/index.17k-MAAo.js";import{g as U}from"../chunks/stores.DmnC4CU0.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.XpbzWZVH.js";import{A as V}from"../chunks/AppBar.D1f81ABc.js";import{p as W}from"../chunks/stores.DZzzPzvs.js";function X(p){let e,s='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Reports</li>';return{c(){e=h("ol"),e.innerHTML=s,this.h()},l(a){e=_(a,"OL",{class:!0,"data-svelte-h":!0}),G(e)!=="svelte-yfgz0q"&&(e.innerHTML=s),this.h()},h(){m(e,"class","breadcrumb")},m(a,r){u(a,e,r)},p:P,d(a){a&&o(e)}}}function Y(p){let e,s,a,r,n,f,$,S,i,g,T,z,c,d,q,N;return s=new E({props:{className:"h-4 w-4",name:"bolt"}}),f=new E({props:{className:"h-4 w-4",name:"cloud-arrow-up"}}),g=new E({props:{className:"h-4 w-4",name:"code-bracket-square"}}),d=new E({props:{className:"h-4 w-4",name:"users"}}),{c(){e=h("a"),b(s.$$.fragment),a=j(`

      Live Report`),r=H(),n=h("a"),b(f.$$.fragment),$=j(`

      Saved Reports`),S=H(),i=h("a"),b(g.$$.fragment),T=j(`

      Templates`),z=H(),c=h("a"),b(d.$$.fragment),q=j(`

      Accounts`),this.h()},l(t){e=_(t,"A",{href:!0,class:!0});var l=L(e);v(s.$$.fragment,l),a=R(l,`

      Live Report`),l.forEach(o),r=I(t),n=_(t,"A",{href:!0,class:!0});var B=L(n);v(f.$$.fragment,B),$=R(B,`

      Saved Reports`),B.forEach(o),S=I(t),i=_(t,"A",{href:!0,class:!0});var C=L(i);v(g.$$.fragment,C),T=R(C,`

      Templates`),C.forEach(o),z=I(t),c=_(t,"A",{href:!0,class:!0});var M=L(c);v(d.$$.fragment,M),q=R(M,`

      Accounts`),M.forEach(o),this.h()},h(){m(e,"href",`/z/pages/portal/projects/books/${p[0]}/report/live`),m(e,"class","btn variant-filled btn-sm"),m(n,"href",`/z/pages/portal/projects/books/${p[0]}/report/saved`),m(n,"class","btn variant-filled-tertiary btn-sm"),m(i,"href",`/z/pages/portal/projects/books/${p[0]}/report/template`),m(i,"class","btn variant-filled-secondary btn-sm"),m(c,"href",`/z/pages/portal/projects/books/${p[0]}`),m(c,"class","btn variant-glass-secondary btn-sm")},m(t,l){u(t,e,l),w(s,e,null),x(e,a),u(t,r,l),u(t,n,l),w(f,n,null),x(n,$),u(t,S,l),u(t,i,l),w(g,i,null),x(i,T),u(t,z,l),u(t,c,l),w(d,c,null),x(c,q),N=!0},p:P,i(t){N||(A(s.$$.fragment,t),A(f.$$.fragment,t),A(g.$$.fragment,t),A(d.$$.fragment,t),N=!0)},o(t){k(s.$$.fragment,t),k(f.$$.fragment,t),k(g.$$.fragment,t),k(d.$$.fragment,t),N=!1},d(t){t&&(o(e),o(r),o(n),o(S),o(i),o(z),o(c)),y(s),y(f),y(g),y(d)}}}function Z(p){let e,s;return e=new V({props:{$$slots:{trail:[Y],lead:[X]},$$scope:{ctx:p}}}),{c(){b(e.$$.fragment)},l(a){v(e.$$.fragment,a)},m(a,r){w(e,a,r),s=!0},p(a,[r]){const n={};r&64&&(n.$$scope={dirty:r,ctx:a}),e.$set(n)},i(a){s||(A(e.$$.fragment,a),s=!0)},o(a){k(e.$$.fragment,a),s=!1},d(a){y(e,a)}}}function ee(p,e,s){let a;D(p,W,$=>s(2,a=$));const r=a.params.pid;U();const n=Q(F("__api__"));return(async()=>{const $=await n.listAccount(r);$.status===200&&$.data})(),[r]}class ie extends J{constructor(e){super(),K(this,e,ee,Z,O,{})}}export{ie as component};
