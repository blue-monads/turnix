import{c as I,a as d,t as A}from"../chunks/disclose-version.BOfWGsg9.js";import{p as H,al as Q,f as b,n as M,a as T,h as U,i as L,c as t,b as e,s as u,r,d as z,t as P}from"../chunks/runtime.CmPLUjJJ.js";import{s as V}from"../chunks/render.CZQ66Y_F.js";import{i as N}from"../chunks/if.DoU_1wcm.js";import{e as W,i as X}from"../chunks/each.BJPUi07-.js";import{s as j}from"../chunks/snippet.UtI6fqho.js";import{s as Y}from"../chunks/attributes.yGGqmlPO.js";import{s as R}from"../chunks/class.BFWsl-fM.js";import{d as Z}from"../chunks/events.DnYA5zBv.js";import{p as ee}from"../chunks/proxy.DKs42Gzo.js";import{s as ae,a as te}from"../chunks/store.BN9B0iWR.js";import{r as re}from"../chunks/legacy-client.BxBrBVYa.js";import{p as oe}from"../chunks/stores.P_GwiDr4.js";import{S}from"../chunks/SvgIcon.Cn9REWmg.js";import{o as se}from"../chunks/index-client.Caj9tY0k.js";import{a as ne}from"../chunks/axios.BimPEqV4.js";const ie=async o=>{console.log("buildApi",o);const a=ne.create({baseURL:o,headers:{NodeCTRLKey:"fixme"}});return{getStatus:async()=>await a.get("/z/eapi/status"),startInstance:async()=>await a.post("/z/eapi/start",{})}};function le(o,a){H(a,!0),Q("__start_api__",a.api);var p=I(),c=b(p);N(c,()=>a.api,i=>{var v=I(),m=b(v);j(m,()=>a.children??M),d(i,v)}),d(o,p),T()}var de=A("<a><!> </a>"),pe=(o,a)=>{U(a,!e(a))},ce=A('<div id="sidebar"><div class="p-2 border text-center variant-outline uppercase font-medium font-token"><h4>Startpage</h4></div> <!></div> <div class="absolute md:bottom-8"><button class=" rounded border-2 bg-slate-100 p-1"><span class="hidden md:block"><!></span> <span class="block md:hidden"><!></span></button></div>',1),me=A('<div class="flex relative h-screen flex-col md:flex-row md:w-full border-r border-gray-800"><!> <main class="grow w-full h-screen overflow-auto"><!></main></div>');function $e(o,a){H(a,!0);const p=ae(),c=()=>te(oe,"$page",p);c().params.pid;let i=L(!0),v=[{name:"Home",link:"/z/pages/startpage",icon:"home",key:"home"},{name:"Peers",link:"/z/pages/startpage/peers",icon:"rss",key:"peers"},{name:"Extend",link:"/z/pages/startpage/extend",icon:"puzzle-piece",key:"extend"},{name:"Easy Share",link:"/z/pages/startpage/share",icon:"share",key:"share"},{name:"Setting",link:"/z/pages/startpage/setting",icon:"cog",key:"setting"}],m=z(()=>c().url.pathname.split("/").at(4));re(()=>{console.log(e(m))}),se(()=>{const s=window;s.__handle_rpc__=function(l,n){console.log("@bang/bang"),console.log(l,n)}});let x=L(void 0);ie(location.origin).then(s=>{U(x,ee(s))});var y=me(),$=t(y);N($,()=>e(m),s=>{var l=ce(),n=b(l),f=u(t(n),2);W(f,17,()=>v,X,(J,h)=>{var g=de(),K=t(g);S(K,{className:"h-4 w-4 mr-1",get name(){return e(h).icon}});var O=u(K);r(g),P(()=>{Y(g,"href",e(h).link),R(g,`flex items-center px-2 py-2 w-full text-gray-600 transition-colors duration-300 transform rounded-lg hover:bg-gray-300 hover:text-gray-700 ${(e(m)===e(h).key?"bg-gray-400 text-gray-900":"")??""}`),V(O,` ${e(h).name??""}`)}),d(J,g)}),r(n);var _=u(n,2),k=t(_);k.__click=[pe,i];var w=t(k),B=t(w),D=z(()=>e(i)?"chevron-double-left":"chevron-double-right");S(B,{className:"h-4 w-4",get name(){return e(D)}}),r(w);var E=u(w,2),F=t(E),G=z(()=>e(i)?"chevron-double-up":"chevron-double-down");S(F,{className:"h-4 w-4",get name(){return e(G)}}),r(E),r(k),r(_),P(()=>R(n,`p-4 flex flex-col gap-2 transition-all duration-300 ease-in-out bg-blend-darken bg-gradient-to-b from-purple-50 shadow
${(e(i)?"w-48":"hidden")??""}`)),d(s,l)});var C=u($,2),q=t(C);N(q,()=>e(x),s=>{le(s,{get api(){return e(x)},children:(l,n)=>{var f=I(),_=b(f);j(_,()=>a.children??M),d(l,f)},$$slots:{default:!0}})}),r(C),r(y),d(o,y),T()}Z(["click"]);export{$e as component};
