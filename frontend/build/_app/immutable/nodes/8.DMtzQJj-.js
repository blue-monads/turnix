import{s as Q,e as s,a as y,c as l,b as x,f,g as C,w as D,m as t,i as W,h as e,A as R,y as j,n as K,B as X}from"../chunks/scheduler.Dy-hT7lj.js";import{S as Y,i as Z}from"../chunks/index.jcCBAMyh.js";import{a as ee}from"../chunks/axios.Cm0UX6qg.js";const te=u=>ee.post(location.origin+"/z/auth/login",u);function ae(u){let n,a,d,v,r,o,I="Login",w,m,_,L,M="Username:",U,i,k,b,T,$="Password:",q,c,z,H,O,A,h,F="<span>Submit</span>",V,G;return{c(){n=s("div"),a=s("div"),d=s("header"),v=y(),r=s("div"),o=s("h3"),o.textContent=I,w=y(),m=s("article"),_=s("label"),L=s("span"),L.textContent=M,U=y(),i=s("input"),k=y(),b=s("label"),T=s("span"),T.textContent=$,q=y(),c=s("input"),z=y(),H=s("hr"),O=y(),A=s("footer"),h=s("button"),h.innerHTML=F,this.h()},l(p){n=l(p,"DIV",{class:!0});var E=x(n);a=l(E,"DIV",{class:!0});var g=x(a);d=l(g,"HEADER",{}),x(d).forEach(f),v=C(g),r=l(g,"DIV",{class:!0});var P=x(r);o=l(P,"H3",{class:!0,"data-toc-ignore":!0,"data-svelte-h":!0}),D(o)!=="svelte-4qdpqo"&&(o.textContent=I),w=C(P),m=l(P,"ARTICLE",{class:!0});var S=x(m);_=l(S,"LABEL",{class:!0});var N=x(_);L=l(N,"SPAN",{"data-svelte-h":!0}),D(L)!=="svelte-dexfrc"&&(L.textContent=M),U=C(N),i=l(N,"INPUT",{class:!0,type:!0,placeholder:!0}),N.forEach(f),k=C(S),b=l(S,"LABEL",{class:!0});var B=x(b);T=l(B,"SPAN",{"data-svelte-h":!0}),D(T)!=="svelte-vmwi0t"&&(T.textContent=$),q=C(B),c=l(B,"INPUT",{class:!0,type:!0,placeholder:!0}),B.forEach(f),S.forEach(f),P.forEach(f),z=C(g),H=l(g,"HR",{class:!0}),O=C(g),A=l(g,"FOOTER",{class:!0});var J=x(A);h=l(J,"BUTTON",{type:!0,class:!0,"data-svelte-h":!0}),D(h)!=="svelte-z9conx"&&(h.innerHTML=F),J.forEach(f),g.forEach(f),E.forEach(f),this.h()},h(){t(o,"class","h3"),t(o,"data-toc-ignore",""),t(i,"class","input p-2"),t(i,"type","text"),t(i,"placeholder","Email"),t(_,"class","label"),t(c,"class","input p-2"),t(c,"type","password"),t(c,"placeholder","password"),t(b,"class","label"),t(m,"class","flex flex-col"),t(r,"class","p-4 space-y-4"),t(H,"class","opacity-50"),t(h,"type","button"),t(h,"class","btn variant-filled-primary"),t(A,"class","p-4 flex justify-start items-center space-x-4"),t(a,"class","card bg-initial card-hover overflow-hidden variant-ringed"),t(n,"class","flex justify-center items-center")},m(p,E){W(p,n,E),e(n,a),e(a,d),e(a,v),e(a,r),e(r,o),e(r,w),e(r,m),e(m,_),e(_,L),e(_,U),e(_,i),R(i,u[0]),e(m,k),e(m,b),e(b,T),e(b,q),e(b,c),R(c,u[1]),e(a,z),e(a,H),e(a,O),e(a,A),e(A,h),V||(G=[j(i,"input",u[3]),j(c,"input",u[4]),j(h,"click",u[2])],V=!0)},p(p,[E]){E&1&&i.value!==p[0]&&R(i,p[0]),E&2&&c.value!==p[1]&&R(c,p[1])},i:K,o:K,d(p){p&&f(n),V=!1,X(G)}}}function se(u,n,a){let d="dev@example.com",v="dev123";const r=async()=>{const w=await te({email:d,password:v});w.status===200&&(localStorage.setItem("access_token",w.data.access_token),window.location.pathname="/z/pages/portal")};function o(){d=this.value,a(0,d)}function I(){v=this.value,a(1,v)}return[d,v,r,o,I]}class oe extends Y{constructor(n){super(),Z(this,n,se,ae,Q,{})}}export{oe as component};