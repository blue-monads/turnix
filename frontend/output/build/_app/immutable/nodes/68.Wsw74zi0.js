import{a as x,t as y}from"../chunks/uYeY98Ep.js";import{p as h,g as k,a as w,i as p,d,h as n,c as b,r as L}from"../chunks/CRNwc6nF.js";import{i as j}from"../chunks/DPugg-eY.js";import{p as z}from"../chunks/Ycl9y-hy.js";import{s as A,a as c}from"../chunks/xNd68Pov.js";import{N as B}from"../chunks/Diz-v8IQ.js";import{p as C}from"../chunks/bUAA_rse.js";import{N as E}from"../chunks/Dq_7tmzI.js";import{g as I}from"../chunks/lZe3btPj.js";import{p as P}from"../chunks/BIb1nXgJ.js";import{L as S}from"../chunks/DDEYG0vs.js";var q=y('<div class="px-2 py-10 md:px-20"><!></div>');function V(f,l){h(l,!0);const[r,g]=A(),u=()=>c(C,"$page",r),e=()=>c(P,"$params",r),t=u().params.pid,i=E(k("__api__"));let o=n(!0),m=n(null);(async()=>{p(o,!0);const a=await i.getNotepad(t,e().id);a.status===200&&(p(m,z(a.data)),p(o,!1))})();var s=q(),v=b(s);{var _=a=>{S(a,{})},$=a=>{B(a,{get data(){return d(m)},onSave:async N=>{await i.updateNotepad(t,e().id,N),I(`/z/pages/portal/projects/books/${t}/notepad`)}})};j(v,a=>{d(o)?a(_):a($,!1)})}L(s),x(f,s),w(),g()}export{V as component};
