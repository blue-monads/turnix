import{V as q,Y as O,a5 as B,a6 as N,a7 as D,a8 as T,b as f,h as U,a9 as S,aa as Y,ab as j,ac as y,a3 as M,d as R,ad as V,m as $,ae as _,af as z}from"./runtime.D6lgFUio.js";import{p as C}from"./proxy.r5-P5EHK.js";const G={get(e,r){if(!e.exclude.includes(r))return e.props[r]},set(e,r){return!1},getOwnPropertyDescriptor(e,r){if(!e.exclude.includes(r)&&r in e.props)return{enumerable:!0,configurable:!0,value:e.props[r]}},has(e,r){return e.exclude.includes(r)?!1:r in e.props},ownKeys(e){return Reflect.ownKeys(e.props).filter(r=>!e.exclude.includes(r))}};function W(e,r,n){return new Proxy({props:e,exclude:r},G)}const Z={get(e,r){if(!e.exclude.includes(r))return f(e.version),r in e.special?e.special[r]():e.props[r]},set(e,r,n){return r in e.special||(e.special[r]=H({get[r](){return e.props[r]}},r,D)),e.special[r](n),S(e.version),!0},getOwnPropertyDescriptor(e,r){if(!e.exclude.includes(r)&&r in e.props)return{enumerable:!0,configurable:!0,value:e.props[r]}},deleteProperty(e,r){return e.exclude.includes(r)||(e.exclude.push(r),S(e.version)),!0},has(e,r){return e.exclude.includes(r)?!1:r in e.props},ownKeys(e){return Reflect.ownKeys(e.props).filter(r=>!e.exclude.includes(r))}};function X(e,r){return new Proxy({props:e,exclude:r,special:{},version:q(0)},Z)}const F={get(e,r){let n=e.props.length;for(;n--;){let u=e.props[n];if(_(u)&&(u=u()),typeof u=="object"&&u!==null&&r in u)return u[r]}},getOwnPropertyDescriptor(e,r){let n=e.props.length;for(;n--;){let u=e.props[n];if(_(u)&&(u=u()),typeof u=="object"&&u!==null&&r in u){const i=O(u,r);return i&&!i.configurable&&(i.configurable=!0),i}}},has(e,r){for(let n of e.props)if(_(n)&&(n=n()),n!=null&&r in n)return!0;return!1},ownKeys(e){const r=[];for(let n of e.props){_(n)&&(n=n());for(const u in n)r.includes(u)||r.push(u)}return r}};function g(...e){return new Proxy({props:e},F)}function H(e,r,n,u){var I;var i=(n&Y)!==0,c=(n&j)!==0,E=(n&y)!==0,A=(n&z)!==0,d=e[r],o=(I=O(e,r))==null?void 0:I.set,a=u,v=!0,P=!1,w=()=>(P=!0,v&&(v=!1,A?a=M(u):a=u),a);d===void 0&&u!==void 0&&(o&&c&&B(),d=w(),o&&o(d));var l;if(c)l=()=>{var s=e[r];return s===void 0?w():(v=!0,P=!1,s)};else{var m=(i?R:V)(()=>e[r]);m.f|=N,l=()=>{var s=f(m);return s!==void 0&&(a=void 0),s===void 0?a:s}}if(!(n&D))return l;if(o){var K=e.$$legacy;return function(s,t){return arguments.length>0?((!c||!t||K)&&o(t?l():s),s):l()}}var h=!1,b=$(d),p=R(()=>{var s=l(),t=f(b);return h?(h=!1,t):b.v=s});return i||(p.equals=T),function(s,t){var L=f(p);if(arguments.length>0){const x=t?f(p):c&&E?C(s):s;return p.equals(x)||(h=!0,U(b,x),P&&a!==void 0&&(a=x),f(p)),s}return L}}export{X as l,H as p,W as r,g as s};