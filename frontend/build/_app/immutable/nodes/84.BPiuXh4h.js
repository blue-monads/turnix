import"../chunks/disclose-version.wcsHG7oF.js";import{p as n}from"../chunks/proxy.r5-P5EHK.js";import{p as i,g as _,a as f,b as c,h as g,i as d}from"../chunks/runtime.D6lgFUio.js";import{s as u,a as T}from"../chunks/store.BaapSzAp.js";import{A as l}from"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";/* empty css                                                     */import{f as y}from"../chunks/route.fvf27QKG.js";import"../chunks/index.xGEmn6Hr.js";import{p as E}from"../chunks/stores.D4Dcgzbh.js";function U(s,o){i(o,!0);const p=u();let e=T(E,"$page",p).params.pid;const m=_("__api__");let t=d("");l(s,{data:{},get message(){return c(t)},schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents"},{ftype:"KEY_VALUE_TEXT",key_name:"result",name:"Result"},{ftype:"LONG_TEXT",key_name:"contentScript",name:"Content script"}],name:"Add Template",required_fields:["name"]},onSave:async r=>{const a=await m.addTemplate(e,r);if(a.status!==200){g(t,n(a.data));return}y(e)}}),f()}export{U as component};
