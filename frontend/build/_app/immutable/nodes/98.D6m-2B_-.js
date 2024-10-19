import{a as I,t as J}from"../chunks/disclose-version.wcsHG7oF.js";import{p as K,g as L,a as M,s,c as l,h as r,b as a,i as p,r as t,e as O,d as Q}from"../chunks/runtime.D6lgFUio.js";import{f as R}from"../chunks/utils.Dh-XvBAd.js";import{s as V,r as d}from"../chunks/attributes.BFGPVQjG.js";import{r as W}from"../chunks/misc.CFrgioPu.js";import{b as o}from"../chunks/input.DAoTAKLA.js";import{b as X}from"../chunks/select.BfGKzECA.js";import{p as Y}from"../chunks/event-modifiers.Bfc47y5P.js";import{g as Z}from"../chunks/entry.dkjBYykr.js";var $=J('<div class="p-2"><div class="card p-4"><header class="flex justify-center"><div class="flex flex-col gap-2"><h3 class="h3">Create New User</h3></div></header> <section><form class="flex flex-col gap-4"><label class="label"><span>Name</span> <input class="input p-1" type="text" required></label> <label class="label"><span>Email</span> <input class="input p-1" type="email" required></label> <label class="label"><span>Phone</span> <input class="input p-1" type="text" required></label> <label class="label"><span>Password</span> <input class="input p-1" type="password"></label> <label class="label"><span>User Type</span> <select class="select"><option>User</option><option>Bot</option><option>Super User</option></select></label> <label class="label"><span>Bio</span> <textarea class="textarea p-1"></textarea></label> <div class="flex justify-end pt-8 gap-2"><button type="submit" class="btn btn-sm variant-filled">Create</button></div></form></section></div></div>');function na(T,k){K(k,!0);const F=L("__api__");let i=p(""),n=p(""),u=p("user"),v=p(""),c=p(""),b=p("");const G=async()=>{console.log(a(i),a(n),a(u),a(v),a(c),a(b)),(await F.addSelfUser({name:a(i),bio:a(n),utype:a(u),email:a(v),phone:a(c),password:a(b)})).status===200&&Z("/z/pages/portal/self/users")};var m=$(),j=l(m),B=s(l(j),2),_=l(B),H=Q(()=>Y(G)),f=l(_),N=s(l(f),2);d(N),t(f);var x=s(f,2),P=s(l(x),2);d(P),t(x);var h=s(x,2),S=s(l(h),2);d(S),t(h);var g=s(h,2),z=s(l(g),2);d(z),t(g);var y=s(g,2),w=s(l(y),2),U=l(w);U.value=(U.__value="real")==null?"":"real";var q=s(U);q.value=(q.__value="bot")==null?"":"bot";var A=s(q);A.value=(A.__value="super")==null?"":"super",t(w),t(y);var D=s(y,2),C=s(l(D),2);W(C),V(C,"rows",4),t(D),O(2),t(_),t(B),t(j),t(m),R("submit",_,function(...e){var E;(E=a(H))==null||E.apply(this,e)}),o(N,()=>a(i),e=>r(i,e)),o(P,()=>a(v),e=>r(v,e)),o(S,()=>a(c),e=>r(c,e)),o(z,()=>a(b),e=>r(b,e)),X(w,()=>a(u),e=>r(u,e)),o(C,()=>a(n),e=>r(n,e)),I(T,m),M()}export{na as component};