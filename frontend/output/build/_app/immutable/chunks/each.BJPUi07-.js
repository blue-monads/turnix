import{I as j,N as ee,au as Q,K as N,y as k,z as ae,a1 as re,F as U,O as ne,P as Y,x as D,A as R,Q as W,J as X,R as le,av as m,a5 as fe,L as F,aw as L,h as J,ax as O,ay as ie,m as ue,W as K,az as se,E as te,aA as ve,ak as de,q as _e,B as oe,aB as ce,v as he}from"./runtime.CmPLUjJJ.js";let M=null;function Ie(i,e){return e}function Ee(i,e,a,s){for(var v=[],d=e.length,u=0;u<d;u++)se(e[u].e,v,!0);var A=d>0&&v.length===0&&a!==null;if(A){var E=a.parentNode;te(E),E.append(a),s.clear(),C(i,e[0].prev,e[d-1].next)}ve(v,()=>{for(var c=0;c<d;c++){var _=e[c];A||(s.delete(_.k),C(i,_.prev,_.next)),de(_.e,!A)}})}function Te(i,e,a,s,v,d=null){var u=i,A={flags:e,items:new Map,first:null},E=(e&Q)!==0;if(E){var c=i;u=N?k(_e(c)):c.appendChild(j())}N&&ae();var _=null;ee(()=>{var n=a(),o=re(n)?n:n==null?[]:U(n),l=o.length;let t=!1;if(N){var p=u.data===ne;p!==(l===0)&&(u=Y(),k(u),D(!1),t=!0)}if(N){for(var x=null,I,h=0;h<l;h++){if(R.nodeType===8&&R.data===oe){u=R,t=!0,D(!1);break}var r=o[h],f=s(r,h);I=G(R,A,x,null,r,f,h,v,e),A.items.set(f,I),x=I}l>0&&k(Y())}N||pe(o,A,u,v,e,s),d!==null&&(l===0?_?W(_):_=X(()=>d(u)):_!==null&&le(_,()=>{_=null})),t&&D(!0)}),N&&(u=R)}function pe(i,e,a,s,v,d){var y,B,q,z;var u=(v&ce)!==0,A=(v&(L|O))!==0,E=i.length,c=e.items,_=e.first,n=_,o,l=null,t,p=[],x=[],I,h,r,f;if(u)for(f=0;f<E;f+=1)I=i[f],h=d(I,f),r=c.get(h),r!==void 0&&((y=r.a)==null||y.measure(),(t??(t=new Set)).add(r));for(f=0;f<E;f+=1){if(I=i[f],h=d(I,f),r=c.get(h),r===void 0){var Z=n?n.e.nodes_start:a;l=G(Z,e,l,l===null?e.first:l.next,I,h,f,s,v),c.set(h,l),p=[],x=[],n=l.next;continue}if(A&&Ae(r,I,f,v),r.e.f&m&&(W(r.e),u&&((B=r.a)==null||B.unfix(),(t??(t=new Set)).delete(r))),r!==n){if(o!==void 0&&o.has(r)){if(p.length<x.length){var g=x[0],T;l=g.prev;var b=p[0],H=p[p.length-1];for(T=0;T<p.length;T+=1)P(p[T],g,a);for(T=0;T<x.length;T+=1)o.delete(x[T]);C(e,b.prev,H.next),C(e,l,b),C(e,H,g),n=g,l=H,f-=1,p=[],x=[]}else o.delete(r),P(r,n,a),C(e,r.prev,r.next),C(e,r,l===null?e.first:l.next),C(e,l,r),l=r;continue}for(p=[],x=[];n!==null&&n.k!==h;)n.e.f&m||(o??(o=new Set)).add(n),x.push(n),n=n.next;if(n===null)continue;r=n}p.push(r),l=r,n=r.next}if(n!==null||o!==void 0){for(var w=o===void 0?[]:U(o);n!==null;)w.push(n),n=n.next;var S=w.length;if(S>0){var $=v&Q&&E===0?a:null;if(u){for(f=0;f<S;f+=1)(q=w[f].a)==null||q.measure();for(f=0;f<S;f+=1)(z=w[f].a)==null||z.fix()}Ee(e,w,$,c)}}u&&fe(()=>{var V;if(t!==void 0)for(r of t)(V=r.a)==null||V.apply()}),F.first=e.first&&e.first.e,F.last=l&&l.e}function Ae(i,e,a,s){s&L&&J(i.v,e),s&O?J(i.i,a):i.i=a}function G(i,e,a,s,v,d,u,A,E){var c=M;try{var _=(E&L)!==0,n=(E&ie)===0,o=_?n?ue(v):K(v):v,l=E&O?K(u):u,t={i:l,v:o,k:d,a:null,e:null,prev:a,next:s};return M=t,t.e=X(()=>A(i,o,l),N),t.e.prev=a&&a.e,t.e.next=s&&s.e,a===null?e.first=t:(a.next=t,a.e.next=t.e),s!==null&&(s.prev=t,s.e.prev=t.e),t}finally{M=c}}function P(i,e,a){for(var s=i.next?i.next.e.nodes_start:a,v=e?e.e.nodes_start:a,d=i.e.nodes_start;d!==s;){var u=he(d);v.before(d),d=u}}function C(i,e,a){e===null?i.first=a:(e.next=a,e.e.next=a&&a.e),a!==null&&(a.prev=e,a.e.prev=e&&e.e)}export{Te as e,Ie as i};
