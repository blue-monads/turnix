import{n as a,o as m,az as q,ay as z}from"./BC9Cf8Q3.js";function _(e,t,n){if(e==null)return t(void 0),n&&n(void 0),a;const r=m(()=>e.subscribe(t,n));return r.unsubscribe?()=>r.unsubscribe():r}const f=[];function x(e,t){return{subscribe:A(e,t).subscribe}}function A(e,t=a){let n=null;const r=new Set;function i(u){if(z(e,u)&&(e=u,n)){const o=!f.length;for(const s of r)s[1](),f.push(s,e);if(o){for(let s=0;s<f.length;s+=2)f[s][0](f[s+1]);f.length=0}}}function b(u){i(u(e))}function l(u,o=a){const s=[u,o];return r.add(s),r.size===1&&(n=t(i,b)||a),u(e),()=>{r.delete(s),r.size===0&&n&&(n(),n=null)}}return{set:i,update:b,subscribe:l}}function B(e,t,n){const r=!Array.isArray(e),i=r?[e]:e;if(!i.every(Boolean))throw new Error("derived() expects stores as input, got a falsy value");const b=t.length<2;return x(n,(l,u)=>{let o=!1;const s=[];let d=0,p=a;const y=()=>{if(d)return;p();const c=t(r?s[0]:s,l,u);b?l(c):p=typeof c=="function"?c:a},h=i.map((c,g)=>_(c,w=>{s[g]=w,d&=~(1<<g),o&&y()},()=>{d|=1<<g}));return o=!0,y(),function(){q(h),p(),o=!1}})}function E(e){let t;return _(e,n=>t=n)(),t}export{B as d,E as g,x as r,_ as s,A as w};
