import"../chunks/disclose-version.wcsHG7oF.js";import{p as m,g as s,a as r}from"../chunks/runtime.D6lgFUio.js";import{i}from"../chunks/lifecycle.ZPOFI82y.js";import{s as f,a as d}from"../chunks/store.BaapSzAp.js";import{g as y}from"../chunks/entry.Y6HcnCx3.js";import{p as _}from"../chunks/stores.D4Dcgzbh.js";import{N as c}from"../chunks/index.Duto3aX0.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as T}from"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";/* empty css                                                     */function O(a,o){m(o,!1);const t=f(),e=d(_,"$page",t).params.pid,n=c(s("__api__"));i(),T(a,{data:{},schema:{name:"Add Contact",fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Info",ftype:"LONG_TEXT",key_name:"info"},{name:"Type",ftype:"SELECT",key_name:"ctype",options:["vendor","client"]},{name:"Email",ftype:"EMAIL",key_name:"email"},{name:"Phone",ftype:"TEXT",key_name:"phone"},{name:"Second phone",ftype:"TEXT",key_name:"phone2"},{name:"Third phone",ftype:"TEXT",key_name:"phone3"},{name:"Address",ftype:"LONG_TEXT",key_name:"address"},{name:"Address2",ftype:"LONG_TEXT",key_name:"address2"}],required_fields:["name"]},onSave:async p=>{(await n.addContact(e,p)).status===200&&y(`/z/pages/portal/projects/books/${e}/contacts`)}}),r()}export{O as component};