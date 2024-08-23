import{s as ue,e as f,a as k,c as h,b as A,f as _,g as L,v as w,m as u,F as fe,i as z,h as c,z as j,G as ce,x as U,D as he,n as Q,q as _e,A as de,t as W,d as Z,E as me}from"./scheduler.D2OtNeTK.js";import{e as re}from"./each.D6YF6ztN.js";import{S as be,i as ve}from"./index.DD9KLV65.js";function oe(n,e,t){const s=n.slice();return s[10]=e[t],s}function pe(n){let e;return{c(){e=W("Add Account")},l(t){e=Z(t,"Add Account")},m(t,s){z(t,e,s)},d(t){t&&_(e)}}}function Ee(n){let e;return{c(){e=W("Edit Account")},l(t){e=Z(t,"Edit Account")},m(t,s){z(t,e,s)},d(t){t&&_(e)}}}function ie(n){let e,t=n[10][1]+"",s;return{c(){e=f("option"),s=W(t),this.h()},l(r){e=h(r,"OPTION",{});var d=A(e);s=Z(d,t),d.forEach(_),this.h()},h(){e.__value=n[10][0],j(e,e.__value)},m(r,d){z(r,e,d),c(e,s)},p:Q,d(r){r&&_(e)}}}function Ae(n){let e,t,s,r,d,m,b,C,H="Name",I,v,P,o,O,$="Type",G,p,V,T,N,ee="Info",X,g,Y,x,te='<button type="submit" class="btn variant-filled">save</button>',J,le;function ae(l,E){return l[3]?Ee:pe}let B=ae(n),y=B(n),S=re(Object.entries(n[5])),i=[];for(let l=0;l<S.length;l+=1)i[l]=ie(oe(n,S,l));return{c(){e=f("form"),t=f("div"),s=f("header"),r=f("h4"),y.c(),d=k(),m=f("section"),b=f("label"),C=f("span"),C.textContent=H,I=k(),v=f("input"),P=k(),o=f("label"),O=f("span"),O.textContent=$,G=k(),p=f("select");for(let l=0;l<i.length;l+=1)i[l].c();V=k(),T=f("label"),N=f("span"),N.textContent=ee,X=k(),g=f("textarea"),Y=k(),x=f("footer"),x.innerHTML=te,this.h()},l(l){e=h(l,"FORM",{class:!0});var E=A(e);t=h(E,"DIV",{class:!0});var a=A(t);s=h(a,"HEADER",{class:!0});var R=A(s);r=h(R,"H4",{class:!0});var ne=A(r);y.l(ne),ne.forEach(_),R.forEach(_),d=L(a),m=h(a,"SECTION",{class:!0});var q=A(m);b=h(q,"LABEL",{class:!0});var D=A(b);C=h(D,"SPAN",{"data-svelte-h":!0}),w(C)!=="svelte-15ueaex"&&(C.textContent=H),I=L(D),v=h(D,"INPUT",{class:!0,type:!0,placeholder:!0}),D.forEach(_),P=L(q),o=h(q,"LABEL",{class:!0});var F=A(o);O=h(F,"SPAN",{"data-svelte-h":!0}),w(O)!=="svelte-12grs2q"&&(O.textContent=$),G=L(F),p=h(F,"SELECT",{class:!0});var se=A(p);for(let K=0;K<i.length;K+=1)i[K].l(se);se.forEach(_),F.forEach(_),V=L(q),T=h(q,"LABEL",{class:!0});var M=A(T);N=h(M,"SPAN",{"data-svelte-h":!0}),w(N)!=="svelte-8iqa3e"&&(N.textContent=ee),X=L(M),g=h(M,"TEXTAREA",{class:!0,rows:!0,placeholder:!0}),A(g).forEach(_),M.forEach(_),q.forEach(_),Y=L(a),x=h(a,"FOOTER",{class:!0,"data-svelte-h":!0}),w(x)!=="svelte-1mj6k96"&&(x.innerHTML=te),a.forEach(_),E.forEach(_),this.h()},h(){u(r,"class","h4"),u(s,"class","card-header"),u(v,"class","input p-1"),u(v,"type","text"),u(v,"placeholder","Input"),u(b,"class","label"),u(p,"class","select"),n[1]===void 0&&fe(()=>n[7].call(p)),u(o,"class","label"),u(g,"class","textarea p-1"),u(g,"rows","4"),u(g,"placeholder","information about account"),u(T,"class","label"),u(m,"class","p-4 flex flex-col gap-4"),u(x,"class","card-footer flex justify-end"),u(t,"class","card"),u(e,"class","p-2")},m(l,E){z(l,e,E),c(e,t),c(t,s),c(s,r),y.m(r,null),c(t,d),c(t,m),c(m,b),c(b,C),c(b,I),c(b,v),j(v,n[0]),c(m,P),c(m,o),c(o,O),c(o,G),c(o,p);for(let a=0;a<i.length;a+=1)i[a]&&i[a].m(p,null);ce(p,n[1],!0),c(m,V),c(m,T),c(T,N),c(T,X),c(T,g),j(g,n[2]),c(t,Y),c(t,x),J||(le=[U(v,"input",n[6]),U(p,"change",n[7]),U(g,"input",n[8]),U(e,"submit",he(n[9]))],J=!0)},p(l,[E]){if(B!==(B=ae(l))&&(y.d(1),y=B(l),y&&(y.c(),y.m(r,null))),E&1&&v.value!==l[0]&&j(v,l[0]),E&32){S=re(Object.entries(l[5]));let a;for(a=0;a<S.length;a+=1){const R=oe(l,S,a);i[a]?i[a].p(R,E):(i[a]=ie(R),i[a].c(),i[a].m(p,null))}for(;a<i.length;a+=1)i[a].d(1);i.length=S.length}E&34&&ce(p,l[1]),E&4&&j(g,l[2])},i:Q,o:Q,d(l){l&&_(e),y.d(),_e(i,l),J=!1,de(le)}}}function ge(n,e,t){let{name:s=""}=e,{acc_type:r="revenue"}=e,{info:d=""}=e,{edit:m=!1}=e,{onChange:b}=e;const C={expenses:"Expenses",revenue:"Revenue",assets:"Assets",liabilities:"Liabilities",equity:"Equity"};function H(){s=this.value,t(0,s)}function I(){r=me(this),t(1,r),t(5,C)}function v(){d=this.value,t(2,d)}const P=()=>{b({name:s,acc_type:r,info:d})};return n.$$set=o=>{"name"in o&&t(0,s=o.name),"acc_type"in o&&t(1,r=o.acc_type),"info"in o&&t(2,d=o.info),"edit"in o&&t(3,m=o.edit),"onChange"in o&&t(4,b=o.onChange)},[s,r,d,m,b,C,H,I,v,P]}class xe extends be{constructor(e){super(),ve(this,e,ge,Ae,ue,{name:0,acc_type:1,info:2,edit:3,onChange:4})}}export{xe as A};
