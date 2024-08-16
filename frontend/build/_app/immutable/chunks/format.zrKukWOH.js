import{s as ie,e as g,a as N,c as v,b as p,q as se,g as S,f as i,m,i as V,h as l,n as Z,v as ne,t as M,A as $,d as O,j as Y}from"./scheduler.CW5wdEZY.js";import{e as J}from"./each.D6YF6ztN.js";import{S as ce,i as de}from"./index.CHcwNeV9.js";const K=a=>{const e=typeof a=="number"?Number(a):parseFloat(a);return isNaN(e)?a:e.toLocaleString(void 0,{minimumFractionDigits:2,maximumFractionDigits:2})};function ee(a,e,t){const r=a.slice();return r[3]=e[t],r}function te(a,e,t){const r=a.slice();return r[6]=e[t],r}function le(a){let e,t=K(a[6].credit_amount)+"",r;return{c(){e=g("span"),r=M(t),this.h()},l(n){e=v(n,"SPAN",{class:!0});var d=p(e);r=O(d,t),d.forEach(i),this.h()},h(){m(e,"class","chip variant-ghost")},m(n,d){V(n,e,d),l(e,r)},p(n,d){d&1&&t!==(t=K(n[6].credit_amount)+"")&&Y(r,t)},d(n){n&&i(e)}}}function re(a){let e,t=K(a[6].debit_amount)+"",r;return{c(){e=g("span"),r=M(t),this.h()},l(n){e=v(n,"SPAN",{class:!0});var d=p(e);r=O(d,t),d.forEach(i),this.h()},h(){m(e,"class","chip variant-ghost")},m(n,d){V(n,e,d),l(e,r)},p(n,d){d&1&&t!==(t=K(n[6].debit_amount)+"")&&Y(r,t)},d(n){n&&i(e)}}}function ae(a){let e,t,r,n,d,u,y,s,c,D=(a[1][a[6].account_id]||"")+"",o,x,I,L,C,P,j,w,F,H,T=a[6].credit_amount&&le(a),b=a[6].debit_amount&&re(a);return{c(){e=g("tr"),t=g("td"),r=N(),n=g("td"),d=N(),u=g("td"),y=N(),s=g("td"),c=g("span"),o=M(D),x=N(),I=g("td"),T&&T.c(),L=N(),C=g("td"),b&&b.c(),P=N(),j=g("td"),w=N(),F=g("td"),H=N(),this.h()},l(E){e=v(E,"TR",{});var _=p(e);t=v(_,"TD",{role:!0}),p(t).forEach(i),r=S(_),n=v(_,"TD",{role:!0}),p(n).forEach(i),d=S(_),u=v(_,"TD",{role:!0}),p(u).forEach(i),y=S(_),s=v(_,"TD",{role:!0});var z=p(s);c=v(z,"SPAN",{class:!0});var B=p(c);o=O(B,D),B.forEach(i),z.forEach(i),x=S(_),I=v(_,"TD",{role:!0});var R=p(I);T&&T.l(R),R.forEach(i),L=S(_),C=v(_,"TD",{role:!0});var q=p(C);b&&b.l(q),q.forEach(i),P=S(_),j=v(_,"TD",{role:!0}),p(j).forEach(i),w=S(_),F=v(_,"TD",{role:!0}),p(F).forEach(i),H=S(_),_.forEach(i),this.h()},h(){m(t,"role","gridcell"),m(n,"role","gridcell"),m(u,"role","gridcell"),m(c,"class","chip variant-filled"),m(s,"role","gridcell"),m(I,"role","gridcell"),m(C,"role","gridcell"),m(j,"role","gridcell"),m(F,"role","gridcell")},m(E,_){V(E,e,_),l(e,t),l(e,r),l(e,n),l(e,d),l(e,u),l(e,y),l(e,s),l(s,c),l(c,o),l(e,x),l(e,I),T&&T.m(I,null),l(e,L),l(e,C),b&&b.m(C,null),l(e,P),l(e,j),l(e,w),l(e,F),l(e,H)},p(E,_){_&3&&D!==(D=(E[1][E[6].account_id]||"")+"")&&Y(o,D),E[6].credit_amount?T?T.p(E,_):(T=le(E),T.c(),T.m(I,null)):T&&(T.d(1),T=null),E[6].debit_amount?b?b.p(E,_):(b=re(E),b.c(),b.m(C,null)):b&&(b.d(1),b=null)},d(E){E&&i(e),T&&T.d(),b&&b.d()}}}function oe(a){let e,t,r=a[3].txn.id+"",n,d,u,y=a[3].txn.title+"",s,c,D,o=a[3].txn.notes+"",x,I,L,C,P,j,w,F,H,T,b,E,_,z,B,R,q=J(a[3].lines),A=[];for(let h=0;h<q.length;h+=1)A[h]=ae(te(a,q,h));return{c(){e=g("tr"),t=g("td"),n=M(r),d=N(),u=g("td"),s=M(y),c=N(),D=g("td"),x=M(o),I=N(),L=g("td"),C=N(),P=g("td"),j=N(),w=g("td"),F=N(),H=g("td"),T=N(),b=g("td"),E=g("a"),_=M("edit"),B=N();for(let h=0;h<A.length;h+=1)A[h].c();R=$(),this.h()},l(h){e=v(h,"TR",{});var f=p(e);t=v(f,"TD",{role:!0});var k=p(t);n=O(k,r),k.forEach(i),d=S(f),u=v(f,"TD",{role:!0});var G=p(u);s=O(G,y),G.forEach(i),c=S(f),D=v(f,"TD",{role:!0});var U=p(D);x=O(U,o),U.forEach(i),I=S(f),L=v(f,"TD",{role:!0}),p(L).forEach(i),C=S(f),P=v(f,"TD",{role:!0}),p(P).forEach(i),j=S(f),w=v(f,"TD",{role:!0}),p(w).forEach(i),F=S(f),H=v(f,"TD",{role:!0}),p(H).forEach(i),T=S(f),b=v(f,"TD",{role:!0});var W=p(b);E=v(W,"A",{class:!0,href:!0});var X=p(E);_=O(X,"edit"),X.forEach(i),W.forEach(i),f.forEach(i),B=S(h);for(let Q=0;Q<A.length;Q+=1)A[Q].l(h);R=$(),this.h()},h(){m(t,"role","gridcell"),m(u,"role","gridcell"),m(D,"role","gridcell"),m(L,"role","gridcell"),m(P,"role","gridcell"),m(w,"role","gridcell"),m(H,"role","gridcell"),m(E,"class","underline"),m(E,"href",z=`/z/pages/portal/projects/books/${a[2]}/txn/edit?tid=${a[3].txn.id}`),m(b,"role","gridcell")},m(h,f){V(h,e,f),l(e,t),l(t,n),l(e,d),l(e,u),l(u,s),l(e,c),l(e,D),l(D,x),l(e,I),l(e,L),l(e,C),l(e,P),l(e,j),l(e,w),l(e,F),l(e,H),l(e,T),l(e,b),l(b,E),l(E,_),V(h,B,f);for(let k=0;k<A.length;k+=1)A[k]&&A[k].m(h,f);V(h,R,f)},p(h,f){if(f&1&&r!==(r=h[3].txn.id+"")&&Y(n,r),f&1&&y!==(y=h[3].txn.title+"")&&Y(s,y),f&1&&o!==(o=h[3].txn.notes+"")&&Y(x,o),f&5&&z!==(z=`/z/pages/portal/projects/books/${h[2]}/txn/edit?tid=${h[3].txn.id}`)&&m(E,"href",z),f&3){q=J(h[3].lines);let k;for(k=0;k<q.length;k+=1){const G=te(h,q,k);A[k]?A[k].p(G,f):(A[k]=ae(G),A[k].c(),A[k].m(R.parentNode,R))}for(;k<A.length;k+=1)A[k].d(1);A.length=q.length}},d(h){h&&(i(e),i(B),i(R)),ne(A,h)}}}function he(a){let e,t,r,n='<tr><th role="columnheader">#</th> <th role="columnheader">Title</th> <th role="columnheader">Notes</th> <th role="columnheader">Account</th> <th role="columnheader">Debit</th> <th role="columnheader">Credit</th> <th role="columnheader">Attachments</th> <th role="columnheader">Actions</th></tr>',d,u,y=J(a[0]),s=[];for(let c=0;c<y.length;c+=1)s[c]=oe(ee(a,y,c));return{c(){e=g("div"),t=g("table"),r=g("thead"),r.innerHTML=n,d=N(),u=g("tbody");for(let c=0;c<s.length;c+=1)s[c].c();this.h()},l(c){e=v(c,"DIV",{class:!0});var D=p(e);t=v(D,"TABLE",{class:!0});var o=p(t);r=v(o,"THEAD",{class:!0,"data-svelte-h":!0}),se(r)!=="svelte-4ay68a"&&(r.innerHTML=n),d=S(o),u=v(o,"TBODY",{class:!0});var x=p(u);for(let I=0;I<s.length;I+=1)s[I].l(x);x.forEach(i),o.forEach(i),D.forEach(i),this.h()},h(){m(r,"class","table-head"),m(u,"class","table-body p-0 overflow-y-scroll w-full max-h-screen"),m(t,"class","table p-0 table-compact"),m(e,"class","card m-2 overflow-x-auto")},m(c,D){V(c,e,D),l(e,t),l(t,r),l(t,d),l(t,u);for(let o=0;o<s.length;o+=1)s[o]&&s[o].m(u,null)},p(c,[D]){if(D&7){y=J(c[0]);let o;for(o=0;o<y.length;o+=1){const x=ee(c,y,o);s[o]?s[o].p(x,D):(s[o]=oe(x),s[o].c(),s[o].m(u,null))}for(;o<s.length;o+=1)s[o].d(1);s.length=y.length}},i:Z,o:Z,d(c){c&&i(e),ne(s,c)}}}function fe(a,e,t){let{lineData:r=[]}=e,{accountsIndex:n}=e,{pid:d}=e;return console.log("@lines",r),a.$$set=u=>{"lineData"in u&&t(0,r=u.lineData),"accountsIndex"in u&&t(1,n=u.accountsIndex),"pid"in u&&t(2,d=u.pid)},[r,n,d]}class ge extends ce{constructor(e){super(),de(this,e,fe,he,ie,{lineData:0,accountsIndex:1,pid:2})}}const ve=a=>{let e={};return a.lines.forEach(t=>{let r=e[t.txn_id];r||(r=[],e[t.txn_id]=r),r.push(t)}),a.transactions.map(t=>({lines:e[t.id]||[],txn:t}))};export{ge as T,ve as f};