import{o as g,q as H,H as I,v as O,w as v,x as h,y as T,z as C,A as _,B as M,C as Y,D as b,E as S,F as V,G as $,I as k,J as q,p as z,K as w,L as B,a as F,M as G}from"./runtime.CmPLUjJJ.js";import{a as J,r as A,h as m}from"./events.DnYA5zBv.js";import{r as K}from"./svelte-head.CnkRZSTV.js";import{d as W}from"./disclose-version.BOfWGsg9.js";import{i as j}from"./utils.DxcAWsds.js";let D=!0;function ee(t,e){e!==(t.__t??(t.__t=t.nodeValue))&&(t.__t=e,t.nodeValue=e==null?"":e+"")}function P(t,e){return L(t,e)}function te(t,e){g(),e.intro=e.intro??!1;const i=e.target,l=w,u=_;try{for(var r=H(i);r&&(r.nodeType!==8||r.data!==I);)r=O(r);if(!r)throw v;h(!0),T(r),C();const o=L(t,{...e,anchor:r});if(_===null||_.nodeType!==8||_.data!==M)throw Y(),v;return h(!1),o}catch(o){if(o===v)return e.recover===!1&&b(),g(),S(i),h(!1),P(t,e);throw o}finally{h(l),T(u),K()}}const d=new Map;function L(t,{target:e,anchor:i,props:l={},events:u,context:r,intro:o=!0}){g();var p=new Set,y=n=>{for(var a=0;a<n.length;a++){var s=n[a];if(!p.has(s)){p.add(s);var f=j(s);e.addEventListener(s,m,{passive:f});var R=d.get(s);R===void 0?(document.addEventListener(s,m,{passive:f}),d.set(s,1)):d.set(s,R+1)}}};y(V(J)),A.add(y);var c=void 0,N=$(()=>{var n=i??e.appendChild(k());return q(()=>{if(r){z({});var a=G;a.c=r}u&&(l.$$events=u),w&&W(n,null),D=o,c=t(n,l)||{},D=!0,w&&(B.nodes_end=_),r&&F()}),()=>{var f;for(var a of p){e.removeEventListener(a,m);var s=d.get(a);--s===0?(document.removeEventListener(a,m),d.delete(a)):d.set(a,s)}A.delete(y),E.delete(c),n!==i&&((f=n.parentNode)==null||f.removeChild(n))}});return E.set(c,N),c}let E=new WeakMap;function re(t){const e=E.get(t);e&&e()}export{D as a,te as h,P as m,ee as s,re as u};