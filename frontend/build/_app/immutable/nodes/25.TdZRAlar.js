import{a as n,t as c}from"../chunks/disclose-version.wcsHG7oF.js";import{p as j,g as C,i as $,f as N,a as w,h as l,s as x,b as m,c as S,e as z,r as I,d as L}from"../chunks/runtime.D6lgFUio.js";import{d as B}from"../chunks/utils.Dh-XvBAd.js";import{i as D}from"../chunks/if.lkQ-IRrr.js";import{p as y}from"../chunks/proxy.r5-P5EHK.js";import{s as M,a as R}from"../chunks/store.BaapSzAp.js";import{p as U}from"../chunks/index.xGEmn6Hr.js";import"../chunks/auto_form.CxbPX2mo.js";import{A as q}from"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";import{L as E}from"../chunks/loader.BeY12cOT.js";import{g as h}from"../chunks/entry.Y6HcnCx3.js";import{g as F}from"../chunks/stores.CebXFGJ4.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as G}from"../chunks/AppBar.3srT6C5i.js";import{S as H}from"../chunks/SvgIcon.Cj01WnO8.js";var J=c('<h4 class="h4">Projects Plugins</h4>'),K=(d,e,i)=>{e.trigger({type:"component",component:"project_new_plugin_picker",meta:{pid:i}})},O=c('<button class="btn btn-sm variant-ghost-secondary"><!> Plugin</button>'),Q=c("<!> <!>",1);function mt(d,e){j(e,!0);const i=M(),b=()=>R(U,"$params",i),g=C("__api__");let o=b().pid,u=$(y([])),p=$(!0);const f=async()=>{l(p,!0);const t=await g.listProjectPlugins(o);t.status===200&&(l(u,y(t.data)),l(p,!1))};f();const A=F(),P=async t=>{(await g.removeProjectPlugin(o,t)).status===200&&f()};var _=Q(),v=N(_);G(v,{$$slots:{lead:(t,s)=>{var a=J();n(t,a)},trail:(t,s)=>{var a=O();a.__click=[K,A,o];var r=S(a);H(r,{name:"plus",className:"w-6 h-6"}),z(),I(a),n(t,a)}}});var k=x(v,2);D(k,()=>m(p),t=>{E(t,{})},t=>{var s=L(()=>[{Name:"Run",Class:"variant-filled-success",Action:async(a,r)=>{console.log("@run",a,r),h(`/z/pages/portal/project/plugins/run?pid=${o}&id=${a}`)}},{Name:"edit",Class:"variant-filled-primary",Action:async(a,r)=>{console.log("@edit",a,r),h(`/z/pages/portal/project/plugins/edit?pid=${o}&id=${a}`)}},{Name:"remove",Class:"variant-filled-error",Action:P}]);q(t,{get datas(){return m(u)},action_key:"id",key_names:[["id","ID"],["name","Name"],["ptype","type"],["created_at","Created At"],["updated_at","Updated At"]],color:["ptype"],get actions(){return m(s)}})}),n(d,_),w()}B(["click"]);export{mt as component};
