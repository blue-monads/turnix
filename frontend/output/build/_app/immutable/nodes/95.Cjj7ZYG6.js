import"../chunks/disclose-version.BOfWGsg9.js";import{p as u,g as f,a as _}from"../chunks/runtime.CmPLUjJJ.js";import{i as c}from"../chunks/lifecycle.46CgxKeU.js";import{s as g,a as s}from"../chunks/store.BN9B0iWR.js";import{A as l}from"../chunks/index.DnNofn7z.js";import"../chunks/autotable.D5y0wA8l.js";/* empty css                                                     */import{b as E}from"../chunks/route.Q5tnqHfN.js";import{p as y}from"../chunks/index.COSOhg_g.js";import{N as A}from"../chunks/index.DQVP37Ea.js";import{p as $}from"../chunks/stores.DkdCKZTz.js";function j(o,p){u(p,!1);const t=g(),m=()=>s($,"$page",t),r=()=>s(y,"$params",t);let e=m().params.pid,a=r().tid;const n=A(f("__api__"));let i="";c(),l(o,{data:{},message:i,schema:{fields:[{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"SELECT",key_name:"status",name:"Status",options:["submited","processed","denied"]},{ftype:"KEY_VALUE_TEXT",key_name:"extrameta",name:"Extra Meta"}],name:"Add Queue message",required_fields:["name"]},onSave:async d=>{(await n.addQueueMessage(e,{...d,projectId:Number(e),templateId:Number(a)})).status===200&&E(e,a)}}),_()}export{j as component};