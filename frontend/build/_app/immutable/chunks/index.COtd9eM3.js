import{d as i}from"./index.FsgtF6qT.js";import{n as m,p as t}from"./stores.DZzzPzvs.js";const c=i([m,t],([a,e])=>{let r;a&&a.to?r=a.to.url.searchParams:r=new URLSearchParams(location.search);let s={};return e.params&&(s.pid=e.params.pid),r.forEach((o,p)=>{s[p]=o}),s});c.subscribe(a=>{console.log("@params",a)});export{c as p};
