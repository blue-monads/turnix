import{s as E,a as F,e as g,g as I,c as b,b as v,f as c,m as h,i as p,h as T,y as x,k as w,q as O,n as U,x as D,t as $,d as k,w as P,j as V}from"../chunks/scheduler.Dy-hT7lj.js";import{S as B,i as G,c as j,b as z,m as C,t as S,a as N,e as q}from"../chunks/index.jcCBAMyh.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";import{F as R}from"../chunks/FileDropzone.T7HLKVtM.js";import{p as A}from"../chunks/stores.YmkBQNJ3.js";/* empty css                                                     */import{p as H}from"../chunks/index.NnJejm2M.js";import{S as J}from"../chunks/SvgIcon.D5ofxCQw.js";import{g as K}from"../chunks/entry.Cp0zZRwk.js";function L(r){let e,a,n;return a=new J({props:{name:"arrow-up-tray",className:"w-6 h-6"}}),{c(){e=g("div"),j(a.$$.fragment),this.h()},l(t){e=b(t,"DIV",{slot:!0,class:!0});var o=v(e);z(a.$$.fragment,o),o.forEach(c),this.h()},h(){h(e,"slot","lead"),h(e,"class","flex justify-center")},m(t,o){p(t,e,o),C(a,e,null),n=!0},p:U,i(t){n||(S(a.$$.fragment,t),n=!0)},o(t){N(a.$$.fragment,t),n=!1},d(t){t&&c(e),q(a)}}}function M(r){let e,a="Upload a file",n;return{c(){e=g("strong"),e.textContent=a,n=$(" or drag and drop")},l(t){e=b(t,"STRONG",{"data-svelte-h":!0}),P(e)!=="svelte-13uz6lq"&&(e.textContent=a),n=k(t," or drag and drop")},m(t,o){p(t,e,o),p(t,n,o)},p:U,d(t){t&&(c(e),c(n))}}}function Q(r){let e=r[1].name+"",a;return{c(){a=$(e)},l(n){a=k(n,e)},m(n,t){p(n,a,t)},p(n,t){t&2&&e!==(e=n[1].name+"")&&V(a,e)},d(n){n&&c(a)}}}function W(r){let e;function a(o,i){return o[1]?Q:M}let n=a(r),t=n(r);return{c(){t.c(),e=D()},l(o){t.l(o),e=D()},m(o,i){t.m(o,i),p(o,e,i)},p(o,i){n===(n=a(o))&&t?t.p(o,i):(t.d(1),t=n(o),t&&(t.c(),t.m(e.parentNode,e)))},d(o){o&&c(e),t.d(o)}}}function X(r){let e;return{c(){e=$("Upload")},l(a){e=k(a,"Upload")},m(a,n){p(a,e,n)},d(a){a&&c(e)}}}function Y(r){let e;return{c(){e=$("Uploading...")},l(a){e=k(a,"Uploading...")},m(a,n){p(a,e,n)},d(a){a&&c(e)}}}function Z(r){let e,a,n,t,o,i,m,u;e=new R({props:{name:"files",$$slots:{message:[W],lead:[L]},$$scope:{ctx:r}}}),e.$on("change",r[5]);function y(s,f){return s[0]?Y:X}let d=y(r),l=d(r);return{c(){j(e.$$.fragment),a=F(),n=g("div"),t=g("button"),l.c(),this.h()},l(s){z(e.$$.fragment,s),a=I(s),n=b(s,"DIV",{class:!0});var f=v(n);t=b(f,"BUTTON",{class:!0});var _=v(t);l.l(_),_.forEach(c),f.forEach(c),this.h()},h(){h(t,"class","btn btn-sm variant-filled"),t.disabled=o=!r[1]||r[0],h(n,"class","flex justify-end py-2")},m(s,f){C(e,s,f),p(s,a,f),p(s,n,f),T(n,t),l.m(t,null),i=!0,m||(u=x(t,"click",r[6]),m=!0)},p(s,[f]){const _={};f&258&&(_.$$scope={dirty:f,ctx:s}),e.$set(_),d!==(d=y(s))&&(l.d(1),l=d(s),l&&(l.c(),l.m(t,null))),(!i||f&3&&o!==(o=!s[1]||s[0]))&&(t.disabled=o)},i(s){i||(S(e.$$.fragment,s),i=!0)},o(s){N(e.$$.fragment,s),i=!1},d(s){s&&(c(a),c(n)),q(e,s),l.d(),m=!1,u()}}}function ee(r,e,a){let n,t;w(r,A,l=>a(7,n=l)),w(r,H,l=>a(2,t=l));const o=n.params.pid,i=O("__api__");let m=!1,u;return[m,u,t,o,i,async l=>{console.log("onDrop",l),a(1,u=l.target.files[0])},async()=>{a(0,m=!0);const l=t.folder||"";await i.addProjectFile(o,u.name,l,u),a(0,m=!1),K(`/z/pages/portal/project/files/${o}?folder=${l}`)}]}class fe extends B{constructor(e){super(),G(this,e,ee,Z,E,{})}}export{fe as component};
