import{s as $,x as _,i as E,f as w,k as d,q as h,n as x}from"../chunks/scheduler.Dy-hT7lj.js";import{S as A,i as L,a as p,d as S,t as f,g as T,c as g,b,m as k,e as y}from"../chunks/index.jcCBAMyh.js";import{A as C}from"../chunks/auto_form.cwLFn19I.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";import{L as q}from"../chunks/loader.DwCv7VNw.js";import{p as N}from"../chunks/index.NnJejm2M.js";import{N as P}from"../chunks/index.B_G37SEf.js";import{p as U}from"../chunks/stores.YmkBQNJ3.js";function K(n){let e,a;return e=new C({props:{data:n[0],message:V,schema:{fields:[{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents",disabled:!0},{ftype:"SELECT",key_name:"status",name:"Status",options:["submited","processed","denied"],disabled:!0},{ftype:"KEY_VALUE_TEXT",key_name:"extrameta",name:"Extra Meta",disabled:!0}],name:"Preview Queue message",required_fields:[]},onSave:X}}),{c(){g(e.$$.fragment)},l(t){b(e.$$.fragment,t)},m(t,o){k(e,t,o),a=!0},p(t,o){const i={};o&1&&(i.data=t[0]),e.$set(i)},i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){p(e.$$.fragment,t),a=!1},d(t){y(e,t)}}}function M(n){let e,a;return e=new q({}),{c(){g(e.$$.fragment)},l(t){b(e.$$.fragment,t)},m(t,o){k(e,t,o),a=!0},p:x,i(t){a||(f(e.$$.fragment,t),a=!0)},o(t){p(e.$$.fragment,t),a=!1},d(t){y(e,t)}}}function Q(n){let e,a,t,o;const i=[M,K],s=[];function l(r,m){return r[1]?0:1}return e=l(n),a=s[e]=i[e](n),{c(){a.c(),t=_()},l(r){a.l(r),t=_()},m(r,m){s[e].m(r,m),E(r,t,m),o=!0},p(r,[m]){let u=e;e=l(r),e===u?s[e].p(r,m):(T(),p(s[u],1,1,()=>{s[u]=null}),S(),a=s[e],a?a.p(r,m):(a=s[e]=i[e](r),a.c()),f(a,1),a.m(t.parentNode,t))},i(r){o||(f(a),o=!0)},o(r){p(a),o=!1},d(r){r&&w(t),s[e].d(r)}}}let V="";const X=async n=>{};function Y(n,e,a){let t,o;d(n,N,c=>a(2,t=c)),d(n,U,c=>a(3,o=c));let i=o.params.pid,s=t.mid,l={},r=!0;const m=P(h("__api__"));return(async()=>{const c=await m.getQueueMessage(i,s);c.status===200&&(a(0,l=c.data),a(1,r=!1))})(),[l,r]}class H extends A{constructor(e){super(),L(this,e,Y,Q,$,{})}}export{H as component};
