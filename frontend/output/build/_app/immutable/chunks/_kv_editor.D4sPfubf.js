import{a as p,t as u}from"./disclose-version.BOfWGsg9.js";import{p as rt,a as st,c as a,s as i,r,b as s,h as l,f as ot,t as S,i as y}from"./runtime.CmPLUjJJ.js";import{s as lt}from"./render.CZQ66Y_F.js";import{i as Q}from"./if.DoU_1wcm.js";import{k as dt}from"./key.B3I6Q0YW.js";import{e as it,i as nt}from"./each.BJPUi07-.js";import{r as x,b as R,s as pt}from"./attributes.yGGqmlPO.js";import{d as ut}from"./events.DnYA5zBv.js";import{b as T}from"./input.BRsTNa69.js";import{p as ct}from"./proxy.DKs42Gzo.js";import{p as g}from"./props.IjRlZquB.js";import{r as U}from"./legacy-client.BxBrBVYa.js";var vt=(m,o,d)=>o(d(),m.target.value),_t=u('<input type="text" class="border border-slate-500 rounded-sm w-full">'),bt=u("<input>"),mt=(m,o,d)=>{l(o,ct(d()))},ft=u('<button class="font-medium text-blue-600 hover:underline">edit</button>'),ht=u('<tr class="bg-white border-b hover:bg-gray-50"><th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"><div> </div></th><td class="px-6 py-4"><!></td><td class="px-6 py-4 text-right"><!> <button class="font-medium text-blue-600 hover:underline">delete</button></td></tr>'),yt=u('<!> <tr class="bg-gray-50 border-b hover:bg-gray-100"><th scope="row" class="font-medium text-gray-900 whitespace-nowrap"><input type="text" class="border border-slate-500 rounded-sm w-full"></th><td><input type="text" class="border border-slate-500 rounded-sm w-full"></td><td class="text-right"><button class="font-medium text-blue-600 hover:underline">add</button></td></tr>',1),xt=u('<div class="border p-2 shadow rounded"><table class="w-full text-sm text-left text-gray-500"><thead class="text-xs text-gray-700 uppercase bg-gray-50"><tr><th scope="col" class="px-6 py-3">Key</th><th scope="col" class="px-6 py-3">Value</th><th scope="col" class="px-6 py-3"><span class="sr-only">delete</span></th></tr></thead><tbody><!></tbody></table></div>');function qt(m,o){rt(o,!0);let d=g(o,"modified",15,!1),w=g(o,"data",19,()=>({})),V=g(o,"onChange",3,void 0),W=g(o,"sensitive",3,!1),e;U(()=>{e={...typeof w()=="string"?JSON.parse(w()):w()}});const X=()=>({...e});let k=y(""),c=y(""),f=y(""),v=y(!1);const Y=(O,h)=>{d(!0),l(v,!1),e={...e,[O]:h}};U(()=>{d()&&V()&&!s(v)&&(V()(e),l(v,!0))});var C=xt(),q=a(C),z=i(a(q)),Z=a(z);return dt(Z,()=>e,O=>{var h=yt(),A=ot(h);it(A,17,()=>Object.entries(e),nt,(_,H)=>{let n=()=>s(H)[0],I=()=>s(H)[1];var J=ht(),K=a(J),L=a(K),tt=a(L);r(L),r(K);var N=i(K),et=a(N);Q(et,()=>s(k)===n(),b=>{var t=_t();x(t),t.__change=[vt,Y,n],S(()=>R(t,I()+"")),p(b,t)},b=>{var t=bt();x(t),t.disabled=!0,S(()=>{R(t,I()),pt(t,"type",W()?"password":"text")}),p(b,t)}),r(N);var M=i(N),P=a(M);Q(P,()=>s(k)!==n(),b=>{var t=ft();t.__click=[mt,k,n],p(b,t)});var at=i(P,2);at.__click=()=>{delete e[n()],e={...e},d(!0),l(v,!1)},r(M),r(J),S(()=>lt(tt,n())),p(_,J)});var B=i(A,2),j=a(B),E=a(j);x(E),r(j);var D=i(j),F=a(D);x(F),r(D);var G=i(D),$=a(G);$.__click=()=>{s(c)&&(d(!0),l(v,!1),e={...e,[s(c)]:s(f)},l(c,""),l(f,""))},r(G),r(B),T(E,()=>s(c),_=>l(c,_)),T(F,()=>s(f),_=>l(f,_)),p(O,h)}),r(z),r(q),r(C),p(m,C),st({getData:X})}ut(["change","click"]);export{qt as _};
