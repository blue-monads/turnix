import{a as L,t as z}from"../chunks/uYeY98Ep.js";import{p as D,g as E,a as F,i as e,d as o,h,c as t,s as r,r as s}from"../chunks/CRNwc6nF.js";import{d as G,r as H}from"../chunks/CHLatDmB.js";import{i as I}from"../chunks/DPugg-eY.js";import{r as J,s as K}from"../chunks/Bo8IgBjd.js";import{b as S}from"../chunks/C6nodsJR.js";import{p as q}from"../chunks/Ycl9y-hy.js";import{s as M,a as O}from"../chunks/xNd68Pov.js";import{g as P}from"../chunks/lZe3btPj.js";import{p as Q}from"../chunks/BIb1nXgJ.js";import{L as R}from"../chunks/DDEYG0vs.js";const T=async(y,p,n,c,d)=>{(await p.updateSelfUser(n,{name:o(c),bio:o(d)})).status===200&&P("/z/pages/portal/self/users")};var V=z(`<div class="card p-4"><header class="flex justify-center"><div class="flex flex-col gap-2"><h3 class="h3">Update User</h3></div></header> <section><form class="flex flex-col gap-4"><label class="label"><span>Name</span> <input class="input p-1" type="text" required></label> <label class="label"><span>Bio</span> <textarea class="textarea p-1">
                        </textarea></label> <div class="flex justify-end pt-8 gap-2"><button type="button" class="btn btn-sm variant-filled">Update</button></div></form></section></div>`),W=z('<div class="p-2"><!></div>');function na(y,p){D(p,!0);const[n,c]=M(),d=()=>O(Q,"$params",n),m=E("__api__");let l=h(""),i=h(""),f=h(!0);const U=d().uid;(async()=>{e(f,!0);const a=await m.getSelfUser(U);a.status===200&&(e(l,q(a.data.name)),e(i,q(a.data.bio)),e(f,!1))})();var u=W(),A=t(u);{var B=a=>{R(a,{})},C=a=>{var v=V(),$=r(t(v),2),k=t($),b=t(k),w=r(t(b),2);J(w),s(b);var _=r(b,2),g=r(t(_),2);H(g),K(g,"rows",4),s(_);var j=r(_,2),N=t(j);N.__click=[T,m,U,l,i],s(j),s(k),s($),s(v),S(w,()=>o(l),x=>e(l,x)),S(g,()=>o(i),x=>e(i,x)),L(a,v)};I(A,a=>{o(f)?a(B):a(C,!1)})}s(u),L(y,u),F(),c()}G(["click"]);export{na as component};
