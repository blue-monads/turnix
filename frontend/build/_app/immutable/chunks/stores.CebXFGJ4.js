import{g as a,ak as s}from"./runtime.D6lgFUio.js";import{w as l}from"./index.BxIQjJFd.js";const n="modalStore";function f(){const t=a(n);if(!t)throw new Error("modalStore is not initialized. Please ensure that `initializeStores()` is invoked in the root layout file of this app!");return t}function g(){const t=u();return s(n,t)}function u(){const{subscribe:t,set:i,update:o}=l([]);return{subscribe:t,set:i,update:o,trigger:e=>o(r=>(r.push(e),r)),close:()=>o(e=>(e.length>0&&e.shift(),e)),clear:()=>i([])}}export{f as g,g as i};
