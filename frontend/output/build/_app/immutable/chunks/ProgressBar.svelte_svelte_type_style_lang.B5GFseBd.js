import{K as S,z as p}from"./runtime.CmPLUjJJ.js";import{g as v,w as M,r as L}from"./index.CdZvgsHB.js";function R(e,t,r,n,d){var a;S&&p();var s=(a=t.$$slots)==null?void 0:a[r],o=!1;s===!0&&(s=t[r==="default"?"children":r],o=!0),s===void 0?d!==null&&d(e):s(e,o?()=>n:n)}function x(e){const t={};e.children&&(t.default=!0);for(const r in e.$$slots)t[r]=!0;return t}const f={};function w(e){return e==="local"?localStorage:sessionStorage}function l(e,t,r){const n=JSON,d="local";function s(o,a){w(d).setItem(o,n.stringify(a))}if(!f[e]){const o=M(t,i=>{const c=w(d).getItem(e);c&&i(n.parse(c));{const m=u=>{u.key===e&&i(u.newValue?n.parse(u.newValue):null)};return window.addEventListener("storage",m),()=>window.removeEventListener("storage",m)}}),{subscribe:a,set:g}=o;f[e]={set(i){s(e,i),g(i)},update(i){const c=i(v(o));s(e,c),g(c)},subscribe:a}}return f[e]}l("modeOsPrefers",!1);l("modeUserPrefers",void 0);l("modeCurrent",!1);const h="(prefers-reduced-motion: reduce)";function z(){return window.matchMedia(h).matches}const y=L(z(),e=>{{const t=n=>{e(n.matches)},r=window.matchMedia(h);return r.addEventListener("change",t),()=>{r.removeEventListener("change",t)}}});export{R as a,y as p,x as s};