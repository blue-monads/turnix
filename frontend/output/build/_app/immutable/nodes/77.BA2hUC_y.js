import"../chunks/uYeY98Ep.js";import"../chunks/BmWNUg7v.js";import{p as m,g as i,a as n}from"../chunks/CRNwc6nF.js";import{i as f}from"../chunks/1kGoK7q1.js";import{s as _,a as c}from"../chunks/xNd68Pov.js";import{g}from"../chunks/lZe3btPj.js";import{p as d}from"../chunks/bUAA_rse.js";import{N as l}from"../chunks/Dq_7tmzI.js";import{A as u}from"../chunks/D8sggRce.js";import"../chunks/g85Fs1Im.js";/* empty css                */function S(o,t){m(t,!1);const[p,a]=_(),e=c(d,"$page",p).params.pid,r=l(i("__api__"));f(),u(o,{data:{},schema:{name:"Add Report",fields:[{name:"Name",ftype:"TEXT",key_name:"name"},{name:"Type",ftype:"SELECT",key_name:"report_type",options:["html_report","sql_report"]}],required_fields:["name","report_type"]},onSave:async s=>{(await r.addReportTemplate(e,s)).status===200&&g(`/z/pages/portal/projects/books/${e}/report`)}}),n(),a()}export{S as component};
