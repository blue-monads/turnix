import{s as d,a as h,g as b,i as $,f as _,k as g,q as k,e as w,c as H,w as v,m as x,n as j}from"../chunks/scheduler.Dy-hT7lj.js";import{S as C,i as L,c as p,b as c,m,t as l,a as f,e as u}from"../chunks/index.jcCBAMyh.js";import{p as P}from"../chunks/index.NnJejm2M.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";import{A as S}from"../chunks/AppBar.C2N-oVhp.js";import{H as q}from"../chunks/hookForm.Dh1z5Mkb.js";function y(n){let t,s='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Projects</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Hooks</li> <li class="crumb-separator" aria-hidden="">›</li> <li>New</li>';return{c(){t=w("ol"),t.innerHTML=s,this.h()},l(a){t=H(a,"OL",{class:!0,"data-svelte-h":!0}),v(t)!=="svelte-oizep5"&&(t.innerHTML=s),this.h()},h(){x(t,"class","breadcrumb")},m(a,r){$(a,t,r)},p:j,d(a){a&&_(t)}}}function z(n){let t,s,a,r;return t=new S({props:{$$slots:{lead:[y]},$$scope:{ctx:n}}}),a=new q({props:{onSave:n[2]}}),{c(){p(t.$$.fragment),s=h(),p(a.$$.fragment)},l(e){c(t.$$.fragment,e),s=b(e),c(a.$$.fragment,e)},m(e,o){m(t,e,o),$(e,s,o),m(a,e,o),r=!0},p(e,[o]){const i={};o&16&&(i.$$scope={dirty:o,ctx:e}),t.$set(i)},i(e){r||(l(t.$$.fragment,e),l(a.$$.fragment,e),r=!0)},o(e){f(t.$$.fragment,e),f(a.$$.fragment,e),r=!1},d(e){e&&_(s),u(t,e),u(a,e)}}}function A(n,t,s){let a;g(n,P,i=>s(3,a=i));const r=a.pid,e=k("__api__");return[r,e,async i=>{(await e.addProjectHook(r,i)).status===200&&location.pathname.replace("/new","")}]}class E extends C{constructor(t){super(),L(this,t,A,z,d,{})}}export{E as component};
