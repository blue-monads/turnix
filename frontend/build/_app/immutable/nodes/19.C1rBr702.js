import{s as y,w as d,i as h,f as w,k as P,r as S,n as j}from"../chunks/scheduler.D2OtNeTK.js";import{S as E,i as C,a as l,d as L,t as u,g as T,c as g,b,m as k,e as $}from"../chunks/index.DD9KLV65.js";import{b as q}from"../chunks/index.ZqtmngW7.js";import{A}from"../chunks/auto_form.FPf_zM_P.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B6Y8wNrO.js";import{L as N}from"../chunks/loader.zv9XVPa1.js";import{p as X}from"../chunks/index.BURjlDwe.js";function v(c){let t,o;return t=new A({props:{data:c[0],message:c[1],schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"SELECT",key_name:"ptype",name:"type",options:["onloop"]}],name:"Edit Project",required_fields:["name"]},onSave:c[5]}}),{c(){g(t.$$.fragment)},l(e){b(t.$$.fragment,e)},m(e,a){k(t,e,a),o=!0},p(e,a){const n={};a&1&&(n.data=e[0]),a&2&&(n.message=e[1]),a&2&&(n.onSave=e[5]),t.$set(n)},i(e){o||(u(t.$$.fragment,e),o=!0)},o(e){l(t.$$.fragment,e),o=!1},d(e){$(t,e)}}}function z(c){let t,o;return t=new N({}),{c(){g(t.$$.fragment)},l(e){b(t.$$.fragment,e)},m(e,a){k(t,e,a),o=!0},p:j,i(e){o||(u(t.$$.fragment,e),o=!0)},o(e){l(t.$$.fragment,e),o=!1},d(e){$(t,e)}}}function B(c){let t,o,e,a;const n=[z,v],s=[];function i(r,m){return r[2]?0:1}return t=i(c),o=s[t]=n[t](c),{c(){o.c(),e=d()},l(r){o.l(r),e=d()},m(r,m){s[t].m(r,m),h(r,e,m),a=!0},p(r,[m]){let p=t;t=i(r),t===p?s[t].p(r,m):(T(),l(s[p],1,1,()=>{s[p]=null}),L(),o=s[t],o?o.p(r,m):(o=s[t]=n[t](r),o.c()),u(o,1),o.m(e.parentNode,e))},i(r){a||(u(o),a=!0)},o(r){l(o),a=!1},d(r){r&&w(e),s[t].d(r)}}}function D(c,t,o){let e;P(c,X,f=>o(6,e=f));const a=S("__api__");let n=e.pid,s={},i="",r=!0;return(async()=>{o(2,r=!0);const f=await a.getProject(n);f.status===200&&(o(0,s=f.data),o(2,r=!1))})(),[s,i,r,a,n,async f=>{const _=await a.updateProject(n,f);if(_.status!==200){o(1,i=_.data);return}q()}]}class O extends E{constructor(t){super(),C(this,t,D,B,y,{})}}export{O as component};
