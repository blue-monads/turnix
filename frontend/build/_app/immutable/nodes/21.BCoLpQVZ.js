import{s as L,a as z,g as B,i as C,f,k as M,z as q,e as _,t as I,c as h,b as N,d as S,m,h as $,n as E,q as D}from"../chunks/scheduler.DVTwEvQJ.js";import{S as H,i as P,c as g,b,m as v,t as y,a as A,e as w}from"../chunks/index.DkBQfKoF.js";import{g as d}from"../chunks/entry.B0930nEj.js";import{S as j}from"../chunks/SvgIcon.BFG_zjfn.js";import{A as T}from"../chunks/autotable.Bb0nJkJi.js";import{g as W}from"../chunks/stores.DEog-POi.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Dy7ATQYE.js";import{A as O}from"../chunks/AppBar.Cm2KLNpw.js";/* empty css                                                     */import{N as V}from"../chunks/index.D3jOKT1N.js";import{p as F}from"../chunks/stores.10CBQxEg.js";function G(i){let t,s,e,r,a,o,l,p,u;return e=new j({props:{className:"h-4 w-4",name:"plus"}}),l=new j({props:{className:"h-4 w-4",name:"plus"}}),{c(){t=_("div"),s=_("a"),g(e.$$.fragment),r=I(`

      account`),a=z(),o=_("a"),g(l.$$.fragment),p=I(`

      transaction`),this.h()},l(n){t=h(n,"DIV",{class:!0});var c=N(t);s=h(c,"A",{href:!0,class:!0});var x=N(s);b(e.$$.fragment,x),r=S(x,`

      account`),x.forEach(f),a=B(c),o=h(c,"A",{href:!0,class:!0});var k=N(o);b(l.$$.fragment,k),p=S(k,`

      transaction`),k.forEach(f),c.forEach(f),this.h()},h(){m(s,"href",`/z/pages/portal/projects/books/${i[1]}/account/new`),m(s,"class","btn variant-filled btn-sm"),m(o,"href",`/z/pages/portal/projects/books/${i[1]}/txn/new`),m(o,"class","btn variant-filled-secondary btn-sm"),m(t,"class","flex flex-wrap justify-end gap-2")},m(n,c){C(n,t,c),$(t,s),v(e,s,null),$(s,r),$(t,a),$(t,o),v(l,o,null),$(o,p),u=!0},p:E,i(n){u||(y(e.$$.fragment,n),y(l.$$.fragment,n),u=!0)},o(n){A(e.$$.fragment,n),A(l.$$.fragment,n),u=!1},d(n){n&&f(t),w(e),w(l)}}}function J(i){let t,s='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Account</li>';return{c(){t=_("ol"),t.innerHTML=s,this.h()},l(e){t=h(e,"OL",{class:!0,"data-svelte-h":!0}),D(t)!=="svelte-12n8fe6"&&(t.innerHTML=s),this.h()},h(){m(t,"class","breadcrumb")},m(e,r){C(e,t,r)},p:E,d(e){e&&f(t)}}}function K(i){let t,s,e,r;return t=new O({props:{$$slots:{lead:[J],default:[G]},$$scope:{ctx:i}}}),e=new T({props:{action_key:"id",key_names:[["id","ID"],["name","Name"],["info","Info"],["acc_type","Account type"]],datas:i[0],color:["acc_type"],actions:[{Name:"explore txns",Class:"variant-filled-primary",icon:"plus",Action:i[2]},{Name:"edit",Class:"variant-filled-secondary",Action:i[3]}]}}),{c(){g(t.$$.fragment),s=z(),g(e.$$.fragment)},l(a){b(t.$$.fragment,a),s=B(a),b(e.$$.fragment,a)},m(a,o){v(t,a,o),C(a,s,o),v(e,a,o),r=!0},p(a,[o]){const l={};o&256&&(l.$$scope={dirty:o,ctx:a}),t.$set(l);const p={};o&1&&(p.datas=a[0]),e.$set(p)},i(a){r||(y(t.$$.fragment,a),y(e.$$.fragment,a),r=!0)},o(a){A(t.$$.fragment,a),A(e.$$.fragment,a),r=!1},d(a){a&&f(s),w(t,a),w(e,a)}}}function Q(i,t,s){let e;M(i,F,n=>s(4,e=n));const r=e.params.pid;W();const a=V(q("__api__"));let o=[];return(async()=>{const n=await a.listAccount(r);n.status===200&&s(0,o=n.data)})(),[o,r,async n=>{const c=`/txn/account?pid=${r}&aid=${n}`;location.pathname.endsWith("/account")?d(location.pathname.replace("/account",c)):d(location.pathname+c)},async n=>{const c=`/account/edit?pid=${r}&aid=${n}`;location.pathname.endsWith("/account")?d(location.pathname.replace("/account",c)):d(location.pathname+c)}]}class rt extends H{constructor(t){super(),P(this,t,Q,K,L,{})}}export{rt as component};
