import"../chunks/disclose-version.wcsHG7oF.js";import{p as u,g as f,a as _}from"../chunks/runtime.D6lgFUio.js";import{i as c}from"../chunks/lifecycle.ZPOFI82y.js";import{s as g,a as s}from"../chunks/store.BaapSzAp.js";import{A as l}from"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";/* empty css                                                     */import{b as E}from"../chunks/route.fvf27QKG.js";import{p as y}from"../chunks/index.xGEmn6Hr.js";import{N as A}from"../chunks/index.DQVP37Ea.js";import{p as $}from"../chunks/stores.D4Dcgzbh.js";function w(o,p){u(p,!1);const t=g(),m=()=>s($,"$page",t),r=()=>s(y,"$params",t);let e=m().params.pid,a=r().tid;const i=A(f("__api__"));let n="";c(),l(o,{data:{},message:n,schema:{fields:[{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"SELECT",key_name:"status",name:"Status",options:["submited","processed","denied"]},{ftype:"KEY_VALUE_TEXT",key_name:"extrameta",name:"Extra Meta"}],name:"Add Queue message",required_fields:["name"]},onSave:async d=>{(await i.addQueueMessage(e,{...d,projectId:Number(e),templateId:Number(a)})).status===200&&E(e,a)}}),_()}export{w as component};
