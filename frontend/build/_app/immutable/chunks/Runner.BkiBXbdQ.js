import{s as w,e as u,c as g,b as h,f as m,m as n,i as y,h as b,x as M,n as f,R as v}from"./scheduler.D2OtNeTK.js";import{S as I,i as k}from"./index.DD9KLV65.js";function x(i){let a,e,d,t;return{c(){a=u("div"),e=u("iframe"),this.h()},l(s){a=g(s,"DIV",{class:!0});var o=h(a);e=g(o,"IFRAME",{title:!0,srcdoc:!0,width:!0,height:!0,class:!0,allow:!0,sandbox:!0});var p=h(e);p.forEach(m),o.forEach(m),this.h()},h(){n(e,"title","preview"),n(e,"srcdoc",i[0]),n(e,"width","100%"),n(e,"height","100%"),n(e,"class","border-green-200 w-full h-full transition-all"),n(e,"allow","accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"),n(e,"sandbox","allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"),n(a,"class","card p-2 h-full w-full")},m(s,o){y(s,a,o),b(a,e),i[5](e),d||(t=M(e,"load",i[4]),d=!0)},p(s,[o]){o&1&&n(e,"srcdoc",s[0])},i:f,o:f,d(s){s&&m(a),i[5](null),d=!1,t()}}}function E(i,a,e){let{client_code:d=""}=a,t,s;const o=async l=>{const r=l.data;if(console.log("onFrameMessage",r),r.type==="plugin_ipc"){console.log("sql_query",r);try{s==null||s.postMessage({msgId:r.msgId,data:{msg:"success",data:{}}})}catch{}}else r.type==="ping"&&(s==null||s.postMessage({msgId:r.msgId,data:{msg:"pong"}}))},p=l=>{var r;try{let c=new MessageChannel;c.port1.onmessage=o,e(2,s=c.port1),console.log("chan.port2 type:",c.port2 instanceof MessagePort),(r=t==null?void 0:t.contentWindow)==null||r.postMessage("transfer_port","*",[c.port2])}catch(c){console.error("Error in postMessage:",c)}};function _(l){v[l?"unshift":"push"](()=>{t=l,e(1,t)})}return i.$$set=l=>{"client_code"in l&&e(0,d=l.client_code)},[d,t,s,o,p,_]}class F extends I{constructor(a){super(),k(this,a,E,x,w,{client_code:0})}}export{F as R};
