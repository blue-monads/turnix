import{s as E,k as M,z as H,e as h,c as _,q as I,m as $,i as u,n as C,f as i,t as L,a as T,b as N,d as R,g as B,h as j}from"../chunks/scheduler.DVTwEvQJ.js";import{S as P,i as O,c as b,b as v,m as w,t as k,a as A,e as S}from"../chunks/index.DkBQfKoF.js";import{S as q}from"../chunks/SvgIcon.BFG_zjfn.js";import{N as D}from"../chunks/index.D3jOKT1N.js";import{g as F}from"../chunks/stores.DEog-POi.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Dy7ATQYE.js";import{A as G}from"../chunks/AppBar.Cm2KLNpw.js";import{p as J}from"../chunks/stores.10CBQxEg.js";function K(o){let e,s='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Reports</li>';return{c(){e=h("ol"),e.innerHTML=s,this.h()},l(t){e=_(t,"OL",{class:!0,"data-svelte-h":!0}),I(e)!=="svelte-yfgz0q"&&(e.innerHTML=s),this.h()},h(){$(e,"class","breadcrumb")},m(t,r){u(t,e,r)},p:C,d(t){t&&i(e)}}}function Q(o){let e,s,t,r,n,c,m,d,l,f,y,g;return s=new q({props:{className:"h-4 w-4",name:"bolt"}}),c=new q({props:{className:"h-4 w-4",name:"cloud-arrow-up"}}),f=new q({props:{className:"h-4 w-4",name:"code-bracket-square"}}),{c(){e=h("a"),b(s.$$.fragment),t=L(`

      Live Report`),r=T(),n=h("a"),b(c.$$.fragment),m=L(`

      Saved Reports`),d=T(),l=h("a"),b(f.$$.fragment),y=L(`

      Templates`),this.h()},l(a){e=_(a,"A",{href:!0,class:!0});var p=N(e);v(s.$$.fragment,p),t=R(p,`

      Live Report`),p.forEach(i),r=B(a),n=_(a,"A",{href:!0,class:!0});var z=N(n);v(c.$$.fragment,z),m=R(z,`

      Saved Reports`),z.forEach(i),d=B(a),l=_(a,"A",{href:!0,class:!0});var x=N(l);v(f.$$.fragment,x),y=R(x,`

      Templates`),x.forEach(i),this.h()},h(){$(e,"href",`/z/pages/portal/projects/books/${o[0]}/report/live`),$(e,"class","btn variant-filled btn-sm"),$(n,"href",`/z/pages/portal/projects/books/${o[0]}/report/saved`),$(n,"class","btn variant-filled-tertiary btn-sm"),$(l,"href",`/z/pages/portal/projects/books/${o[0]}/report/template`),$(l,"class","btn variant-filled-secondary btn-sm")},m(a,p){u(a,e,p),w(s,e,null),j(e,t),u(a,r,p),u(a,n,p),w(c,n,null),j(n,m),u(a,d,p),u(a,l,p),w(f,l,null),j(l,y),g=!0},p:C,i(a){g||(k(s.$$.fragment,a),k(c.$$.fragment,a),k(f.$$.fragment,a),g=!0)},o(a){A(s.$$.fragment,a),A(c.$$.fragment,a),A(f.$$.fragment,a),g=!1},d(a){a&&(i(e),i(r),i(n),i(d),i(l)),S(s),S(c),S(f)}}}function U(o){let e,s;return e=new G({props:{$$slots:{trail:[Q],lead:[K]},$$scope:{ctx:o}}}),{c(){b(e.$$.fragment)},l(t){v(e.$$.fragment,t)},m(t,r){w(e,t,r),s=!0},p(t,[r]){const n={};r&64&&(n.$$scope={dirty:r,ctx:t}),e.$set(n)},i(t){s||(k(e.$$.fragment,t),s=!0)},o(t){A(e.$$.fragment,t),s=!1},d(t){S(e,t)}}}function V(o,e,s){let t;M(o,J,m=>s(2,t=m));const r=t.params.pid;F();const n=D(H("__api__"));return(async()=>{const m=await n.listAccount(r);m.status===200&&m.data})(),[r]}class re extends P{constructor(e){super(),O(this,e,V,U,E,{})}}export{re as component};