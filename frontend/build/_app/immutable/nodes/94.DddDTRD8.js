import{s as de,a as A,x as ne,g as C,i as D,f as d,q as _e,e as p,c as h,w as q,m as f,n as U,t as P,b,d as T,h as c,v as pe,y as he,j as oe}from"../chunks/scheduler.Dy-hT7lj.js";import{S as me,i as ve,c as Q,b as W,m as X,a as x,d as ge,t as I,e as Y,g as be}from"../chunks/index.jcCBAMyh.js";import{e as ie}from"../chunks/each.D6YF6ztN.js";import{S as $e}from"../chunks/SvgIcon.D5ofxCQw.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";import{A as ke}from"../chunks/AppBar.C2N-oVhp.js";import{L as we}from"../chunks/loader.DwCv7VNw.js";function ce(o,e,l){const a=o.slice();return a[6]=e[l],a}function Ee(o){let e,l="Users";return{c(){e=p("h4"),e.textContent=l,this.h()},l(a){e=h(a,"H4",{class:!0,"data-svelte-h":!0}),q(e)!=="svelte-1hx7hst"&&(e.textContent=l),this.h()},h(){f(e,"class","h4")},m(a,t){D(a,e,t)},p:U,d(a){a&&d(e)}}}function ye(o){let e,l,a,t;return l=new $e({props:{className:"h-4 w-4",name:"plus"}}),{c(){e=p("a"),Q(l.$$.fragment),a=P(`

            Add User`),this.h()},l(s){e=h(s,"A",{class:!0,href:!0});var r=b(e);W(l.$$.fragment,r),a=T(r,`

            Add User`),r.forEach(d),this.h()},h(){f(e,"class","btn btn-sm variant-soft-secondary"),f(e,"href","/z/pages/portal/self/users/new")},m(s,r){D(s,e,r),X(l,e,null),c(e,a),t=!0},p:U,i(s){t||(I(l.$$.fragment,s),t=!0)},o(s){x(l.$$.fragment,s),t=!1},d(s){s&&d(e),Y(l)}}}function Ae(o){let e;function l(s,r){return s[0].length===0?Se:De}let a=l(o),t=a(o);return{c(){e=p("div"),t.c(),this.h()},l(s){e=h(s,"DIV",{class:!0});var r=b(e);t.l(r),r.forEach(d),this.h()},h(){f(e,"class","p-2")},m(s,r){D(s,e,r),t.m(e,null)},p(s,r){a===(a=l(s))&&t?t.p(s,r):(t.d(1),t=a(s),t&&(t.c(),t.m(e,null)))},i:U,o:U,d(s){s&&d(e),t.d()}}}function Ce(o){let e,l;return e=new we({}),{c(){Q(e.$$.fragment)},l(a){W(e.$$.fragment,a)},m(a,t){X(e,a,t),l=!0},p:U,i(a){l||(I(e.$$.fragment,a),l=!0)},o(a){x(e.$$.fragment,a),l=!1},d(a){Y(e,a)}}}function De(o){let e,l,a=ie(o[0]),t=[];for(let s=0;s<a.length;s+=1)t[s]=fe(ce(o,a,s));return{c(){e=p("div"),l=p("dl");for(let s=0;s<t.length;s+=1)t[s].c();this.h()},l(s){e=h(s,"DIV",{class:!0});var r=b(e);l=h(r,"DL",{class:!0});var n=b(l);for(let u=0;u<t.length;u+=1)t[u].l(n);n.forEach(d),r.forEach(d),this.h()},h(){f(l,"class","list-dl"),f(e,"class","card p-2")},m(s,r){D(s,e,r),c(e,l);for(let n=0;n<t.length;n+=1)t[n]&&t[n].m(l,null)},p(s,r){if(r&5){a=ie(s[0]);let n;for(n=0;n<a.length;n+=1){const u=ce(s,a,n);t[n]?t[n].p(u,r):(t[n]=fe(u),t[n].c(),t[n].m(l,null))}for(;n<t.length;n+=1)t[n].d(1);t.length=a.length}},d(s){s&&d(e),pe(t,s)}}}function Se(o){let e,l='<div class="flex-auto"><span class="ml-2">No users</span></div>';return{c(){e=p("div"),e.innerHTML=l,this.h()},l(a){e=h(a,"DIV",{class:!0,"data-svelte-h":!0}),q(e)!=="svelte-1sar4ww"&&(e.innerHTML=l),this.h()},h(){f(e,"class","alert alert-info")},m(a,t){D(a,e,t)},p:U,d(a){a&&d(e)}}}function fe(o){let e,l,a='<i class="fa-solid fa-book"></i>',t,s,r,n=o[6].name+"",u,v,i,_=o[6].bio+"",S,z,g,$,B,N,R,k,j,H,O,w,F,V,G,E,Z="Delete",J,K,ee;function ue(){return o[3](o[6])}return{c(){e=p("div"),l=p("span"),l.innerHTML=a,t=A(),s=p("span"),r=p("dt"),u=P(n),v=A(),i=p("dd"),S=P(_),z=A(),g=p("div"),$=p("a"),B=P("Profile"),R=A(),k=p("a"),j=P("Reset Password"),O=A(),w=p("a"),F=P("Edit"),G=A(),E=p("button"),E.textContent=Z,J=A(),this.h()},l(L){e=h(L,"DIV",{});var m=b(e);l=h(m,"SPAN",{class:!0,"data-svelte-h":!0}),q(l)!=="svelte-1h4m9k4"&&(l.innerHTML=a),t=C(m),s=h(m,"SPAN",{class:!0});var M=b(s);r=h(M,"DT",{class:!0});var te=b(r);u=T(te,n),te.forEach(d),v=C(M),i=h(M,"DD",{class:!0});var se=b(i);S=T(se,_),se.forEach(d),M.forEach(d),z=C(m),g=h(m,"DIV",{class:!0});var y=b(g);$=h(y,"A",{href:!0,class:!0});var ae=b($);B=T(ae,"Profile"),ae.forEach(d),R=C(y),k=h(y,"A",{href:!0,class:!0});var le=b(k);j=T(le,"Reset Password"),le.forEach(d),O=C(y),w=h(y,"A",{href:!0,class:!0});var re=b(w);F=T(re,"Edit"),re.forEach(d),G=C(y),E=h(y,"BUTTON",{class:!0,"data-svelte-h":!0}),q(E)!=="svelte-157zwge"&&(E.textContent=Z),y.forEach(d),J=C(m),m.forEach(d),this.h()},h(){f(l,"class","badge-icon p-4 variant-soft-secondary"),f(r,"class","font-bold"),f(i,"class","text-sm opacity-50"),f(s,"class","flex-auto"),f($,"href",N="/z/pages/portal/user/"+o[6].id),f($,"class","btn btn-sm variant-filled-tertiary"),f(k,"href",H="/z/pages/portal/self/users/reset-password?uid="+o[6].id),f(k,"class","btn btn-sm variant-filled-primary"),f(w,"href",V="/z/pages/portal/self/users/edit?uid="+o[6].id),f(w,"class","btn btn-sm variant-filled-warning"),f(E,"class","btn btn-sm variant-filled-error"),f(g,"class","flex gap-0 sm:gap-1 flex-wrap space-y-2")},m(L,m){D(L,e,m),c(e,l),c(e,t),c(e,s),c(s,r),c(r,u),c(s,v),c(s,i),c(i,S),c(e,z),c(e,g),c(g,$),c($,B),c(g,R),c(g,k),c(k,j),c(g,O),c(g,w),c(w,F),c(g,G),c(g,E),c(e,J),K||(ee=he(E,"click",ue),K=!0)},p(L,m){o=L,m&1&&n!==(n=o[6].name+"")&&oe(u,n),m&1&&_!==(_=o[6].bio+"")&&oe(S,_),m&1&&N!==(N="/z/pages/portal/user/"+o[6].id)&&f($,"href",N),m&1&&H!==(H="/z/pages/portal/self/users/reset-password?uid="+o[6].id)&&f(k,"href",H),m&1&&V!==(V="/z/pages/portal/self/users/edit?uid="+o[6].id)&&f(w,"href",V)},d(L){L&&d(e),K=!1,ee()}}}function ze(o){let e,l,a,t,s,r;e=new ke({props:{$$slots:{trail:[ye],lead:[Ee]},$$scope:{ctx:o}}});const n=[Ce,Ae],u=[];function v(i,_){return i[1]?0:1}return a=v(o),t=u[a]=n[a](o),{c(){Q(e.$$.fragment),l=A(),t.c(),s=ne()},l(i){W(e.$$.fragment,i),l=C(i),t.l(i),s=ne()},m(i,_){X(e,i,_),D(i,l,_),u[a].m(i,_),D(i,s,_),r=!0},p(i,[_]){const S={};_&512&&(S.$$scope={dirty:_,ctx:i}),e.$set(S);let z=a;a=v(i),a===z?u[a].p(i,_):(be(),x(u[z],1,1,()=>{u[z]=null}),ge(),t=u[a],t?t.p(i,_):(t=u[a]=n[a](i),t.c()),I(t,1),t.m(s.parentNode,s))},i(i){r||(I(e.$$.fragment,i),I(t),r=!0)},o(i){x(e.$$.fragment,i),x(t),r=!1},d(i){i&&(d(l),d(s)),Y(e,i),u[a].d(i)}}}function Le(o,e,l){const a=_e("__api__");let t=[],s=!0;const r=async()=>{l(1,s=!0);const v=await a.listSelfUsers();v.status===200&&(l(0,t=v.data),l(1,s=!1))},n=async v=>{await a.deleteSelfUser(String(v.id)),r()};return r(),[t,s,n,v=>n(v)]}class Ve extends me{constructor(e){super(),ve(this,e,Le,ze,de,{})}}export{Ve as component};