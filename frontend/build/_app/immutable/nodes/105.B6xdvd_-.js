import{a as t,t as s,c as L}from"../chunks/disclose-version.wcsHG7oF.js";import{p as R,g as q,a as A,h as m,f as O,b as j,s as e,t as B,i as z,c as i,r as p,e as G}from"../chunks/runtime.D6lgFUio.js";import{s as J}from"../chunks/render.CIW-w3Th.js";import{i as _}from"../chunks/if.lkQ-IRrr.js";import{d as K}from"../chunks/utils.Dh-XvBAd.js";import{p as T}from"../chunks/proxy.r5-P5EHK.js";import{p as M}from"../chunks/props.C7PlKWmz.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{S as Q}from"../chunks/SvgIcon.Cj01WnO8.js";var U=s('<div class="flex items-center justify-center h-96 flex-col gap-4"><div class="flex items-center justify-center flex-col gap-4"><h2 class="h2">Loading...</h2></div></div>'),V=s("<span>options</span>"),W=s("<span>sidebar</span>"),X=s('<a href="/z/pages/startpage/home" class="btn variant-soft-secondary">Show Options</a>'),Y=s('<h2 class="h2">Turnix is running.</h2> <span class="text-indigo-500 chip"> </span> <p class="p text-xl">click explore to start using it. Or click on <!></p> <div class="flex gap-2"><button class="btn variant-soft-primary">Explore</button> <!></div>',1),Z=s("<span>options</span>"),$=s("<span>sidebar</span>"),aa=async(I,g,c,u)=>{m(g,!0),await c.startInstance(),u()},ta=s('<a href="/z/pages/startpage/home" class="btn variant-soft-secondary">Show Options</a>'),ra=s('<h2 class="h2">Turnix is not running</h2> <p class="p text-xl">Click start to run a instance in current directory.</p> <p class="p text-xl">Or click on <!></p> <div class="flex gap-2"><button class="btn variant-filled"><!> Start</button> <!></div>',1),sa=s('<div class="flex items-center justify-center min-h-screen flex-col gap-4 pt-20"><!></div>');function da(I,g){R(g,!0);let c=M(g,"isHomePage",3,!0);const u=q("__start_api__");let C=z(!1),H=z(""),h=z(!0),N="";const P=async()=>{m(h,!0);const r=await u.getStatus();r.status===200&&(m(C,T(r.data.is_running)),m(H,T(r.data.working_dir)),N=r.data.port,m(h,!1))};P();const D=()=>{const r=window.__ebrowser_rpc__;r("local-navigate",{url:`http://localhost${N}/z/pages`})};var w=sa(),E=i(w);_(E,()=>j(h),r=>{var x=U();t(r,x)},r=>{var x=L(),F=O(x);_(F,()=>j(C),y=>{var d=Y(),o=e(O(d),2),S=i(o);p(o);var l=e(o,2),v=e(i(l));_(v,()=>!c(),a=>{var f=V();t(a,f)},a=>{var f=W();t(a,f)}),p(l);var b=e(l,2),k=i(b);k.__click=D;var n=e(k,2);_(n,()=>!c(),a=>{var f=X();t(a,f)}),p(b),B(()=>J(S,j(H))),t(y,d)},y=>{var d=ra(),o=e(O(d),4),S=e(i(o));_(S,()=>!c(),n=>{var a=Z();t(n,a)},n=>{var a=$();t(n,a)}),p(o);var l=e(o,2),v=i(l);v.__click=[aa,h,u,P];var b=i(v);Q(b,{className:"h-4 w-4",name:"bolt"}),G(),p(v);var k=e(v,2);_(k,()=>!c(),n=>{var a=ta();t(n,a)}),p(l),t(y,d)},!0),t(r,x)}),p(w),t(I,w),A()}K(["click"]);export{da as component};
