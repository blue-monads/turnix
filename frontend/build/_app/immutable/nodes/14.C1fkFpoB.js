import{s as y,a as l,q as b,f as i,g as $,i as f,y as v,e as C,c as P,r as w,m as S,n as q}from"../chunks/scheduler.CPmtgLt4.js";import{S as A,i as B,c as u,b as _,m as h,t as g,a as d,d as j}from"../chunks/index.CsITmn2N.js";import{g as F}from"../chunks/stores.DmnC4CU0.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.XpbzWZVH.js";import{A as L}from"../chunks/AppBar.D1f81ABc.js";/* empty css                                                     */import{F as k}from"../chunks/index.CjEmxZ4R.js";import{L as E}from"../chunks/listProject.hslISDH_.js";function H(c){let e,s="Projects";return{c(){e=C("h4"),e.textContent=s,this.h()},l(a){e=P(a,"H4",{class:!0,"data-svelte-h":!0}),w(e)!=="svelte-1g1lc9j"&&(e.textContent=s),this.h()},h(){S(e,"class","h4")},m(a,n){f(a,e,n)},p:q,d(a){a&&i(e)}}}function M(c){let e,s,a,n,r,m,p;return s=new L({props:{$$slots:{lead:[H]},$$scope:{ctx:c}}}),n=new E({}),m=new k({props:{handler:c[2]}}),{c(){e=l(),u(s.$$.fragment),a=l(),u(n.$$.fragment),r=l(),u(m.$$.fragment),this.h()},l(t){b("svelte-1en3a0b",document.head).forEach(i),e=$(t),_(s.$$.fragment,t),a=$(t),_(n.$$.fragment,t),r=$(t),_(m.$$.fragment,t),this.h()},h(){document.title="Projects"},m(t,o){f(t,e,o),h(s,t,o),f(t,a,o),h(n,t,o),f(t,r,o),h(m,t,o),p=!0},p(t,[o]){const x={};o&32&&(x.$$scope={dirty:o,ctx:t}),s.$set(x)},i(t){p||(g(s.$$.fragment,t),g(n.$$.fragment,t),g(m.$$.fragment,t),p=!0)},o(t){d(s.$$.fragment,t),d(n.$$.fragment,t),d(m.$$.fragment,t),p=!1},d(t){t&&(i(e),i(a),i(r)),j(s,t),j(n,t),j(m,t)}}}function z(c){const e=v("__api__"),s=async()=>{const r=await e.listProjects();r.status===200&&r.data},a=F();return s(),[e,a,()=>{a.trigger({type:"component",component:"project_picker",meta:{api:e}})}]}class R extends A{constructor(e){super(),B(this,e,z,M,y,{})}}export{R as component};