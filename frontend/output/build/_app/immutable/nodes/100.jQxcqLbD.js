import"../chunks/Celh2SBM.js";import"../chunks/B857cCpa.js";import{p as f,g as _,a as c}from"../chunks/BC9Cf8Q3.js";import{i as g}from"../chunks/D-AVqrqR.js";import{s as l,a as s}from"../chunks/BvXhjoeW.js";import{A as E}from"../chunks/CMsATdMh.js";import"../chunks/C6IzKD7A.js";/* empty css                */import{b as $}from"../chunks/C_5QEsW9.js";import{p as y}from"../chunks/DK_nkBS0.js";import{N as A}from"../chunks/DQVP37Ea.js";import{p as T}from"../chunks/C1rFHlC-.js";function K(o,p){f(p,!1);const[t,m]=l(),r=()=>s(T,"$page",t),n=()=>s(y,"$params",t);let e=r().params.pid,a=n().tid;const i=A(_("__api__"));let u="";g(),E(o,{data:{},message:u,schema:{fields:[{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"SELECT",key_name:"status",name:"Status",options:["submited","processed","denied"]},{ftype:"KEY_VALUE_TEXT",key_name:"extrameta",name:"Extra Meta"}],name:"Add Queue message",required_fields:["name"]},onSave:async d=>{(await i.addQueueMessage(e,{...d,projectId:Number(e),templateId:Number(a)})).status===200&&$(e,a)}}),c(),m()}export{K as component};
