import{s as u,e as h,c as _,w as b,m as v,i as w,n as r,f as x,k as y,q as L,D as C}from"../chunks/scheduler.Dy-hT7lj.js";import{S as R,i as j}from"../chunks/index.jcCBAMyh.js";import{p as U}from"../chunks/index.NnJejm2M.js";function D(s){let t,o='<a href="/" class="hidden">a</a>';return{c(){t=h("div"),t.innerHTML=o,this.h()},l(n){t=_(n,"DIV",{class:!0,"data-svelte-h":!0}),b(t)!=="svelte-1uv7z32"&&(t.innerHTML=o),this.h()},h(){v(t,"class","overflow-auto svelte-1ayc50x")},m(n,l){w(n,t,l),s[1](t)},p:r,i:r,o:r,d(n){n&&x(t),s[1](null)}}}function E(s,t,o){let n;y(s,U,a=>o(2,n=a));const l=L("__api__"),p=n.fid;n.folder;const c=n.filename;let i;const m=["png","jpg","jpeg","gif","bmp","webp","svg"];(async()=>{const a=await l.getSelfFile(p);if(a.status!==200)return;const d=c.split(".").pop();if(d)if(m.includes(d)){console.log("image");const e=new Image;e.src=URL.createObjectURL(a.data),e.onload=()=>{console.log("img loaded")},e.onerror=g=>{console.log("img error",g)},i.appendChild(e)}else{const e=document.createElement("a");e.href=URL.createObjectURL(a.data),e.download=c,e.innerText="download",e.style.color="while",e.style.padding="5%",e.style.paddingLeft="10%",e.style.textDecoration="underline",i.appendChild(e)}})();function f(a){C[a?"unshift":"push"](()=>{i=a,o(0,i)})}return[i,f]}class q extends R{constructor(t){super(),j(this,t,E,D,u,{})}}export{q as component};