import{s as O,k as D,r as F,e as h,c as _,v as G,m as c,i as u,n as P,f as o,t as N,a as H,b as j,d as L,g as I,h as R}from"../chunks/scheduler.D2OtNeTK.js";import{S as J,i as K,c as b,b as v,m as w,t as k,a as A,e as y}from"../chunks/index.DD9KLV65.js";import{S as q}from"../chunks/SvgIcon.DW0uSnC-.js";import{N as Q}from"../chunks/index.ZlEQwogg.js";import{g as U}from"../chunks/stores.BSAXtruD.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B6Y8wNrO.js";import{A as V}from"../chunks/AppBar.C3Ur4xTI.js";import{p as W}from"../chunks/stores.D0bQ952I.js";function X(p){let e,r='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Reports</li>';return{c(){e=h("ol"),e.innerHTML=r,this.h()},l(a){e=_(a,"OL",{class:!0,"data-svelte-h":!0}),G(e)!=="svelte-yfgz0q"&&(e.innerHTML=r),this.h()},h(){c(e,"class","breadcrumb")},m(a,s){u(a,e,s)},p:P,d(a){a&&o(e)}}}function Y(p){let e,r,a,s,n,f,$,S,i,g,x,z,m,d,T,E;return r=new q({props:{className:"h-4 w-4",name:"pencil-square"}}),f=new q({props:{className:"h-4 w-4",name:"bolt"}}),g=new q({props:{className:"h-4 w-4",name:"cloud-arrow-up"}}),d=new q({props:{className:"h-4 w-4",name:"code-bracket-square"}}),{c(){e=h("a"),b(r.$$.fragment),a=N(`

    Editor`),s=H(),n=h("a"),b(f.$$.fragment),$=N(`

      Live Report`),S=H(),i=h("a"),b(g.$$.fragment),x=N(`

      Saved Reports`),z=H(),m=h("a"),b(d.$$.fragment),T=N(`

      Templates`),this.h()},l(t){e=_(t,"A",{href:!0,class:!0});var l=j(e);v(r.$$.fragment,l),a=L(l,`

    Editor`),l.forEach(o),s=I(t),n=_(t,"A",{href:!0,class:!0});var B=j(n);v(f.$$.fragment,B),$=L(B,`

      Live Report`),B.forEach(o),S=I(t),i=_(t,"A",{href:!0,class:!0});var C=j(i);v(g.$$.fragment,C),x=L(C,`

      Saved Reports`),C.forEach(o),z=I(t),m=_(t,"A",{href:!0,class:!0});var M=j(m);v(d.$$.fragment,M),T=L(M,`

      Templates`),M.forEach(o),this.h()},h(){c(e,"href",`/z/pages/portal/projects/books/${p[0]}/report/editor`),c(e,"class","btn variant-filled-primary btn-sm"),c(n,"href",`/z/pages/portal/projects/books/${p[0]}/report/live`),c(n,"class","btn variant-filled btn-sm"),c(i,"href",`/z/pages/portal/projects/books/${p[0]}/report/saved`),c(i,"class","btn variant-filled-tertiary btn-sm"),c(m,"href",`/z/pages/portal/projects/books/${p[0]}/report/template`),c(m,"class","btn variant-filled-secondary btn-sm")},m(t,l){u(t,e,l),w(r,e,null),R(e,a),u(t,s,l),u(t,n,l),w(f,n,null),R(n,$),u(t,S,l),u(t,i,l),w(g,i,null),R(i,x),u(t,z,l),u(t,m,l),w(d,m,null),R(m,T),E=!0},p:P,i(t){E||(k(r.$$.fragment,t),k(f.$$.fragment,t),k(g.$$.fragment,t),k(d.$$.fragment,t),E=!0)},o(t){A(r.$$.fragment,t),A(f.$$.fragment,t),A(g.$$.fragment,t),A(d.$$.fragment,t),E=!1},d(t){t&&(o(e),o(s),o(n),o(S),o(i),o(z),o(m)),y(r),y(f),y(g),y(d)}}}function Z(p){let e,r;return e=new V({props:{$$slots:{trail:[Y],lead:[X]},$$scope:{ctx:p}}}),{c(){b(e.$$.fragment)},l(a){v(e.$$.fragment,a)},m(a,s){w(e,a,s),r=!0},p(a,[s]){const n={};s&64&&(n.$$scope={dirty:s,ctx:a}),e.$set(n)},i(a){r||(k(e.$$.fragment,a),r=!0)},o(a){A(e.$$.fragment,a),r=!1},d(a){y(e,a)}}}function ee(p,e,r){let a;D(p,W,$=>r(2,a=$));const s=a.params.pid;U();const n=Q(F("__api__"));return(async()=>{const $=await n.listAccount(s);$.status===200&&$.data})(),[s]}class ie extends J{constructor(e){super(),K(this,e,ee,Z,O,{})}}export{ie as component};
