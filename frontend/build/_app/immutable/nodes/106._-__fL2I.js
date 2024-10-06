import{a as c,t as f}from"../chunks/disclose-version.wcsHG7oF.js";import{t as b,s as e,c as a,r as t,e as z,b as v}from"../chunks/runtime.D6lgFUio.js";import{s as o}from"../chunks/render.CIW-w3Th.js";import{e as C,i as G}from"../chunks/each.CWrzZiCN.js";import{r as lt,a as N}from"../chunks/attributes.BFGPVQjG.js";import{r as st}from"../chunks/misc.CFrgioPu.js";import{b as pt}from"../chunks/input.DAoTAKLA.js";import{b as ct}from"../chunks/select.BfGKzECA.js";import{p as bt}from"../chunks/proxy.r5-P5EHK.js";import{S as rt}from"../chunks/SvgIcon.Cj01WnO8.js";var ft=f('<tr><td><span class="chip variant-ghost-secondary text-lg"> </span></td></tr>'),ut=f("<tr><td> </td><td> </td><td> </td></tr>"),_t=f("<tr><td> </td><td> </td><td> </td><td> </td></tr>"),mt=f('<div class="flex p-2 flex-col gap-2 overflow-auto max-h-screen"><div class="card p-4 flex flex-col gap-2"><header class="card-header"><h4 class="h4">Peer Info</h4></header> <section class="p-4 flex flex-col gap-4"><label class="label flex flex-col"><span>Id</span> <input class="input p-0.5 w-64" disabled type="text" placeholder="name"></label> <div class="label flex flex-col"><span>Status</span> <div><span class="chip variant-filled"> </span></div></div> <div class="label flex flex-col"><span>Last seen</span> <div><span class="chip variant-filled"> </span></div></div> <div class="label flex flex-col"><span>Peer type</span> <select class="select"><option>Guest</option><option>Admin</option></select></div> <div class="label flex flex-col"><span>Full public key</span> <textarea class="input textarea" disabled></textarea></div> <div class="label flex flex-col"><span>Admin key</span> <textarea class="input textarea" disabled></textarea></div> <div class="label flex flex-col"><span>Pre connect</span> <input type="checkbox" class="checkbox"></div> <div class="label flex flex-col"><span>Multiaddrs</span> <table class="table table-hover overflow-auto input"><thead><tr><th>Addr</th></tr></thead><tbody></tbody></table></div> <div class="label flex flex-col"><span>Static addrs</span> <table class="table table-hover overflow-auto input"><thead><tr><th>Protocol</th><th>IP</th><th>Port</th></tr></thead><tbody></tbody></table> <div class="label flex justify-start p-1"><button class="btn btn-sm variant-filled"><!> add</button></div></div> <div class="label flex flex-col"><span>Static relays</span> <table class="table table-hover overflow-auto input"><thead><tr><th>Protocol</th><th>IP</th><th>Port</th><th>Relay token</th></tr></thead><tbody></tbody></table> <div class="label flex justify-start p-1"><button class="btn btn-sm variant-filled"><!> add</button></div></div> <footer class="card-footer flex justify-end gap-2"><button class="btn btn-sm variant-filled-primary">shell</button> <button class="btn btn-sm variant-filled-secondary">share</button> <button class="btn btn-sm variant-filled-tertiary mr-8">chat</button> <button class="btn btn-sm variant-filled">save</button> <button class="btn btn-sm variant-filled-error">delete</button></footer></section></div></div>');function Tt(xt){const l=bt({id:"xx-yy-zz",added:"2022-12-31",lastSeen:"2022-12-31",status:"online",fullPublicKey:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",peerType:"guest",preConnect:!1,adminKey:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",multiAddrs:["/ip6/2604:a880:1:20::204:4001/udp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"],staticAddrs:[{ip:"44.55.66.77",port:1234,protocol:"tcp"}],staticRelays:[{ip:"44.55.66.77",port:1234,protocol:"tcp",relayToken:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}]});var u=mt(),Q=a(u),E=e(a(Q),2),_=a(E),L=e(a(_),2);lt(L),t(_);var m=e(_,2),M=e(a(m),2),U=a(M),ot=a(U);t(U),t(M),t(m);var h=e(m,2),q=e(a(h),2),B=a(q),dt=a(B);t(B),t(q),t(h);var y=e(h,2),g=e(a(y),2),P=a(g);P.value=(P.__value="guest")==null?"":"guest";var D=e(P);D.value=(D.__value="admin")==null?"":"admin",t(g),t(y);var w=e(y,2),F=e(a(w),2);st(F),t(w);var A=e(w,2),H=e(a(A),2);st(H),t(A);var S=e(A,2),J=e(a(S),2);lt(J),t(S);var k=e(S,2),V=e(a(k),2),Y=e(a(V));C(Y,21,()=>l.multiAddrs,G,(d,r)=>{var s=ft(),x=a(s),n=a(x),i=a(n);t(n),t(x),t(s),b(()=>o(i,v(r))),c(d,s)}),t(Y),t(V),t(k);var I=e(k,2),R=e(a(I),2),O=e(a(R));C(O,21,()=>l.staticAddrs,G,(d,r)=>{var s=ut(),x=a(s),n=a(x);t(x);var i=e(x),K=a(i);t(i);var p=e(i),j=a(p);t(p),t(s),b(()=>{o(n,v(r).protocol),o(K,v(r).ip),o(j,v(r).port)}),c(d,s)}),t(O),t(R);var W=e(R,2),X=a(W),it=a(X);rt(it,{name:"plus",className:"w-4 h-4"}),z(),t(X),t(W),t(I);var Z=e(I,2),T=e(a(Z),2),$=e(a(T));C($,21,()=>l.staticRelays,G,(d,r)=>{var s=_t(),x=a(s),n=a(x);t(x);var i=e(x),K=a(i);t(i);var p=e(i),j=a(p);t(p);var et=e(p),nt=a(et);t(et),t(s),b(()=>{o(n,v(r).protocol),o(K,v(r).ip),o(j,v(r).port),o(nt,v(r).relayToken)}),c(d,s)}),t($),t(T);var tt=e(T,2),at=a(tt),vt=a(at);rt(vt,{name:"plus",className:"w-4 h-4"}),z(),t(at),t(tt),t(Z),z(2),t(E),t(Q),t(u),b(()=>{o(ot,l.status),o(dt,l.lastSeen),N(F,l.fullPublicKey),N(H,l.adminKey),N(J,l.preConnect)}),pt(L,()=>l.id,d=>l.id=d),ct(g,()=>l.peerType,d=>l.peerType=d),c(xt,u)}export{Tt as component};
