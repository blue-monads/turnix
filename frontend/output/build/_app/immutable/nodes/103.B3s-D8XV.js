import"../chunks/Celh2SBM.js";import{p as i,g as _,a as f,d as c,h as g,i as d}from"../chunks/BC9Cf8Q3.js";import{p as u}from"../chunks/CKwprqj2.js";import{s as l,a as T}from"../chunks/BvXhjoeW.js";import{A as y}from"../chunks/CMsATdMh.js";import"../chunks/C6IzKD7A.js";/* empty css                */import{f as E}from"../chunks/C_5QEsW9.js";import"../chunks/DK_nkBS0.js";import{p as $}from"../chunks/C1rFHlC-.js";function U(s,o){i(o,!0);const[p,m]=l();let e=T($,"$page",p).params.pid;const r=_("__api__");let t=d("");y(s,{data:{},get message(){return c(t)},schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"KEY_VALUE_TEXT",key_name:"result",name:"Result"},{ftype:"LONG_TEXT",key_name:"contentScript",name:"Content script"}],name:"Add Template",required_fields:["name"]},onSave:async n=>{const a=await r.addTemplate(e,n);if(a.status!==200){g(t,u(a.data));return}E(e)}}),f(),m()}export{U as component};
