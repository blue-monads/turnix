import{a as t,t as l,c as y}from"./disclose-version.wcsHG7oF.js";import{p as T,a as U,h as c,f as k,t as x,b as s,i as F,c as E,r as K}from"./runtime.D6lgFUio.js";import{d as q}from"./utils.Dh-XvBAd.js";import{i as m}from"./if.lkQ-IRrr.js";import{s as f}from"./attributes.BFGPVQjG.js";import{p as z}from"./proxy.r5-P5EHK.js";import{b as A}from"./this.OuUT1k8T.js";var B=l('<div class="svelte-vtwvgl">loading...</div>'),D=l('<img class="h-fit w-auto overflow-auto">'),G=l('<video controls><source type="video/mp4"></video>'),H=l('<audio controls><source type="audio/mp3"></audio>'),J=l('<button class="btn btn-sm variant-filled">download</button>'),M=l('<div class="h-inherit w-inherit overflow-auto svelte-vtwvgl"><!></div>');function ee(j,a){T(a,!0);let u=F(!1),S=F(void 0);const R=["png","jpg","jpeg","gif","bmp","webp","svg"],p=a.filename.split(".").pop();let g=R.includes(p)?"image":p==="mp4"?"video":p==="mp3"?"audio":"other",d=F("");(async()=>{c(u,!0);const e=await a.api.getFileShortKey(a.fileId);e.status===200&&(c(d,z(a.api.getFileWithShortKeyURL(e.data))),c(u,!1))})();const W=async()=>{const e=document.createElement("a");e.href=s(d),e.download=a.filename,document.body.appendChild(e),e.click()};var n=M();A(n,e=>c(S,e),()=>s(S));var C=E(n);m(C,()=>s(u),e=>{var v=B();t(e,v)},e=>{var v=y(),I=k(v);m(I,()=>g==="image",_=>{var o=D();x(()=>{f(o,"src",s(d)),f(o,"alt",a.filename)}),t(_,o)},_=>{var o=y(),L=k(o);m(L,()=>g==="video",h=>{var i=G(),b=E(i);K(i),x(()=>f(b,"src",s(d))),t(h,i)},h=>{var i=y(),b=k(i);m(b,()=>g==="audio",w=>{var r=H(),P=E(r);K(r),x(()=>f(P,"src",s(d))),t(w,r)},w=>{var r=J();r.__click=W,t(w,r)},!0),t(h,i)},!0),t(_,o)},!0),t(e,v)}),K(n),t(j,n),U()}q(["click"]);export{ee as F};
