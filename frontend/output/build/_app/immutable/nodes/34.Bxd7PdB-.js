import{a as o,t as l,c as tt}from"../chunks/Celh2SBM.js";import{p as X,i as b,h as f,d as t,g as H,a as q,f as Y,c,s as I,r as m,e as T,b as z}from"../chunks/BC9Cf8Q3.js";import{d as et}from"../chunks/x8LL7je8.js";import{s as at}from"../chunks/DMueVNEq.js";import{p as a}from"../chunks/CKwprqj2.js";import{s as B,a as Q}from"../chunks/BvXhjoeW.js";import"../chunks/CRz77-DP.js";import"../chunks/B857cCpa.js";import{A as st}from"../chunks/CnSBROsQ.js";import{C as R,h as ot}from"../chunks/Czp9zhDz.js";import{N as W}from"../chunks/Dq_7tmzI.js";import{S as rt}from"../chunks/CdW99GA5.js";import{p as F}from"../chunks/DK_nkBS0.js";import{i as O}from"../chunks/BgSg1_cg.js";import{p as nt}from"../chunks/BV18QS1o.js";import{T as it,a as U}from"../chunks/Bzk7-iQC.js";import{h as lt}from"../chunks/CZDi8ouT.js";import{j as dt,l as pt,s as ct}from"../chunks/BCwcwXWc.js";import{R as mt}from"../chunks/Hd1NOPgB.js";const G=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="//unpkg.com/alpinejs" defer><\/script>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"><\/script>
    <script>
        let pendingMessages = {}
        let messageId = 0;
        window.addEventListener("message", (ev) => {
            console.log("message", ev);
            if (ev.data !== "transfer_port") {
                console.log("wrong message", ev);
                return;
            }
            const port = ev.ports[0];
            port.onmessage = handleMessage;
            window["__parent_port__"] = port;
        }, false);
        const handleMessage = (ev) => {
            const data = ev.data;
            const messageId = data.msgId;
            const resolve = pendingMessages[messageId];
            delete pendingMessages[messageId];
            resolve(data.data);
            console.log("data", data);
        }
        const sendMessage = async (data) => {
            console.log("sendMessage", data);
            const msgId = messageId + 1;
            messageId = msgId;
            const p = new Promise((resolve) => {
                pendingMessages[msgId] = resolve;
            });
            data["msgId"] = msgId;
            const port = window["__parent_port__"]
            port?.postMessage(data);
            return p;
        }
    <\/script>
</head>
<body>
    <div class="container">
        <h1>Example Plugin UI</h1>          
    </div>
    <div class="container-xl mt-5" x-data="Loader()">
        <div class="container">
            <button id="start" class="btn btn-primary" type="button" x-on:click="load">
                <span class="icon icon-thumbs-up"></span>
                Start
            </button>    
        </div>
    </div>

    <script>
    function Loader() {
      return {
          datas: [],
          loading: true,
          async load() {
              this.loading = true;
              try {

                  const data = await sendMessage({
                      type: "ping",
                      name: "hello",
                      data: [],
                  });

                  console.log("@data", data);

                  this.data = data.data;
              } catch (error) {
                  console.error("Error fetching data:", error);
              } finally {
                  this.loading = false;
              }
          },
          formatDate(dateString) {
              return new Date(dateString).toLocaleDateString();
          }
      }
  }


    <\/script>
</body>
</html>`,J=`


const myPluginDoXyz = (ctx) => {
    console.log("myPluginDoXyz/1");
};

`;var vt=l("<span>Server Code</span>"),ut=l("<span>Client Code</span>"),gt=l("<!> <!>",1),ft=l('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),ht=l('<div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><!></div></div>');function _t(j,$){X($,!0);const[A,P]=B(),C=()=>Q(F,"$params",A);let r=b(0),y=b(a(J)),d=b(a(G)),h=b("");nt($,"runAction",15)(()=>{f(h,a(t(d)))}),C().pid;const M=H("__api__");W(M);const p={};var _=ht(),e=c(_),v=c(e);it(v,{children:(k,K)=>{var g=gt(),x=Y(g);U(x,{padding:"ml-6 p-2",name:"SQL Query",value:0,get group(){return t(r)},set group(s){f(r,a(s))},children:(s,n)=>{var i=vt();o(s,i)},$$slots:{default:!0}});var D=I(x,2);U(D,{name:"UI",value:1,get group(){return t(r)},set group(s){f(r,a(s))},children:(s,n)=>{var i=ut();o(s,i)},$$slots:{default:!0}}),o(k,g)},$$slots:{default:!0,panel:(k,K)=>{var g=ft(),x=c(g);{var D=n=>{const i=T(()=>[ot({activateOnTyping:!0,override:[pt,ct(p)]})]),E=T(dt);R(n,{get extensions(){return t(i)},get lang(){return t(E)},get value(){return t(y)},set value(N){f(y,a(N))}})},s=n=>{var i=tt(),E=Y(i);{var N=L=>{const Z=T(lt);R(L,{get lang(){return t(Z)},get value(){return t(d)},set value(V){f(d,a(V))}})};O(E,L=>{t(r)===1&&L(N)},!0)}o(n,i)};O(x,n=>{t(r)===0?n(D):n(s,!1)})}m(g),o(k,g)}}}),m(e);var u=I(e,2),w=c(u);mt(w,{get client_code(){return t(h)}}),m(u),m(_),o(j,_),q(),P()}var bt=l('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/project">Projects</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li><a class="anchor">Plugins</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),$t=l('<button class="btn variant-filled btn-sm"><!> Run</button>'),yt=l("<!> <!>",1);function Ut(j,$){X($,!0);const[A,P]=B(),C=()=>Q(F,"$params",A);a(J),a(G);const r=C().pid,y=H("__api__");W(y);let d=b(void 0);var h=yt(),S=Y(h);st(S,{$$slots:{lead:(p,_)=>{var e=bt(),v=I(c(e),4),u=c(v);at(u,"href",`/z/pages/portal/project/plugins?pid=${r}`),m(v),z(4),m(e),o(p,e)},trail:(p,_)=>{var e=$t();e.__click=function(...u){var w;(w=t(d))==null||w.apply(this,u)};var v=c(e);rt(v,{className:"w-4 h-4",name:"play"}),z(),m(e),o(p,e)}}});var M=I(S,2);_t(M,{get runAction(){return t(d)},set runAction(p){f(d,a(p))}}),o(j,h),q(),P()}et(["click"]);export{Ut as component};
