import{s as I,e as g,a as C,c as v,w as E,g as $,b as P,f as h,m as y,i as w,h as b,y as O,v as U,t as H,d as M,n as V}from"../chunks/scheduler.Dy-hT7lj.js";import{S as F,i as G,t as N,d as J,a as T,c as K,b as L,m as Q,e as R,g as W}from"../chunks/index.jcCBAMyh.js";import{e as B}from"../chunks/each.D6YF6ztN.js";import{S as X}from"../chunks/SvgIcon.D5ofxCQw.js";import{g as Y}from"../chunks/stores.C8TdjPC3.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";function j(m,t,s){const i=m.slice();return i[3]=t[s],i}function A(m){let t,s,i,a,k=m[3].name+"",d,r,u;return s=new X({props:{name:m[3].icon,className:"w-4 h-4"}}),{c(){t=g("span"),K(s.$$.fragment),i=C(),a=g("p"),d=H(k),r=C(),this.h()},l(c){t=v(c,"SPAN",{class:!0});var f=P(t);L(s.$$.fragment,f),i=$(f),a=v(f,"P",{});var p=P(a);d=M(p,k),p.forEach(h),r=$(f),f.forEach(h),this.h()},h(){y(t,"class","flex gap-1 justify-start items-center hover:bg-zinc-100 p-1 rounded-lg")},m(c,f){w(c,t,f),Q(s,t,null),b(t,i),b(t,a),b(a,d),b(t,r),u=!0},p:V,i(c){u||(N(s.$$.fragment,c),u=!0)},o(c){T(s.$$.fragment,c),u=!1},d(c){c&&h(t),R(s)}}}function Z(m){let t,s="Playground",i,a,k="click me",d,r,u,c="Dropdown",f,p,S,z,q,_=B(m[1]),n=[];for(let e=0;e<_.length;e+=1)n[e]=A(j(m,_,e));const D=e=>T(n[e],1,1,()=>{n[e]=null});return{c(){t=g("h4"),t.textContent=s,i=C(),a=g("button"),a.textContent=k,d=C(),r=g("button"),u=g("span"),u.textContent=c,f=C(),p=g("div");for(let e=0;e<n.length;e+=1)n[e].c();this.h()},l(e){t=v(e,"H4",{class:!0,"data-svelte-h":!0}),E(t)!=="svelte-1p7vrku"&&(t.textContent=s),i=$(e),a=v(e,"BUTTON",{class:!0,"data-svelte-h":!0}),E(a)!=="svelte-1yzgo35"&&(a.textContent=k),d=$(e),r=v(e,"BUTTON",{class:!0});var o=P(r);u=v(o,"SPAN",{"data-svelte-h":!0}),E(u)!=="svelte-15nc1p9"&&(u.textContent=c),f=$(o),p=v(o,"DIV",{class:!0});var l=P(p);for(let x=0;x<n.length;x+=1)n[x].l(l);l.forEach(h),o.forEach(h),this.h()},h(){y(t,"class","h4"),y(a,"class","btn btn-sm variant-filled"),y(p,"class","absolute shadow-lg top-10 left-0 w-32 h-max p-1 bg-white border border-zinc-200 rounded-lg flex flex-col gap-2"),y(r,"class","relative group transition-all duration-200 focus:overflow-visible w-max h-max p-1 overflow-hidden flex flex-row items-center justify-center bg-white gap-2 rounded-lg border border-zinc-200")},m(e,o){w(e,t,o),w(e,i,o),w(e,a,o),w(e,d,o),w(e,r,o),b(r,u),b(r,f),b(r,p);for(let l=0;l<n.length;l+=1)n[l]&&n[l].m(p,null);S=!0,z||(q=O(a,"click",m[2]),z=!0)},p(e,[o]){if(o&2){_=B(e[1]);let l;for(l=0;l<_.length;l+=1){const x=j(e,_,l);n[l]?(n[l].p(x,o),N(n[l],1)):(n[l]=A(x),n[l].c(),N(n[l],1),n[l].m(p,null))}for(W(),l=_.length;l<n.length;l+=1)D(l);J()}},i(e){if(!S){for(let o=0;o<_.length;o+=1)N(n[o]);S=!0}},o(e){n=n.filter(Boolean);for(let o=0;o<n.length;o+=1)T(n[o]);S=!1},d(e){e&&(h(t),h(i),h(a),h(d),h(r)),U(n,e),z=!1,q()}}}function ee(m){const t=Y();return[t,[{name:"rename",icon:"pencil-square"},{name:"download",icon:"arrow-down-on-square"},{name:"delete",icon:"trash"}],()=>{t.trigger({type:"component",component:"file_picker",meta:{}})}]}class re extends F{constructor(t){super(),G(this,t,ee,Z,I,{})}}export{re as component};