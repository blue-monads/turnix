import{a as N,t as O}from"./disclose-version.BOfWGsg9.js";import{p as Q,al as s,aC as y,h as u,at as a,aD as T,c as E,r as I,t as q,b as i,a as U,aE as f,g as c,s as ie}from"./runtime.CmPLUjJJ.js";import{a as V}from"./slot.D3lQyPcJ.js";import{s as P,r as oe,a as ne}from"./attributes.yGGqmlPO.js";import{s as F}from"./class.BFWsl-fM.js";import{i as X}from"./lifecycle.46CgxKeU.js";import{l as H,p as t}from"./props.IjRlZquB.js";import{e as _}from"./events.DnYA5zBv.js";import{d as re}from"./input.BRsTNa69.js";import{b as de}from"./this.DlUubvPq.js";import{b as L}from"./misc.CT9jRRbj.js";var se=O('<div data-testid="radio-group" role="radiogroup"><!></div>');function we(W,e){const n=H(e,["children","$$slots","$$events","$$legacy"]);Q(e,!1);const k=f();let d=t(e,"display",8,"inline-flex"),b=t(e,"flexDirection",8,"flex-row"),v=t(e,"gap",8,"gap-1"),x=t(e,"background",8,"bg-surface-200-700-token"),w=t(e,"border",8,"border-token border-surface-400-500-token"),D=t(e,"rounded",8,"rounded-token"),g=t(e,"padding",8,"px-4 py-1"),A=t(e,"active",8,"variant-filled"),m=t(e,"hover",8,"hover:variant-soft"),G=t(e,"color",8,""),K=t(e,"fill",8,""),p=t(e,"regionLabel",8,""),B=t(e,"labelledby",8,"");s("rounded",D()),s("padding",g()),s("active",A()),s("hover",m()),s("color",G()),s("fill",K()),s("regionLabel",p());const R="p-1";y(()=>(a(d()),a(b()),a(v()),a(x()),a(w()),a(D()),a(n)),()=>{u(k,`${R} ${d()} ${b()} ${v()} ${x()} ${w()} ${D()} ${n.class??""}`)}),T(),X();var r=se(),C=E(r);V(C,e,"default",{},null),I(r),q(()=>{F(r,`radio-group ${i(k)??""}`),P(r,"aria-labelledby",B())}),N(W,r),U()}var ce=O('<label><div data-testid="radio-item" role="radio" tabindex="0"><div class="h-0 w-0 overflow-hidden"><input></div> <!></div></label>');function De(W,e){const n=H(e,["children","$$slots","$$events","$$legacy"]),k=H(n,["group","name","value","title","label","rounded","padding","active","hover","color","fill","regionLabel"]);Q(e,!1);const d=f(),b=f(),v=f(),x=f(),w=f(),D=[];let g=t(e,"group",12),A=t(e,"name",8),m=t(e,"value",8),G=t(e,"title",8,""),K=t(e,"label",8,""),p=t(e,"rounded",24,()=>c("rounded")),B=t(e,"padding",24,()=>c("padding")),R=t(e,"active",24,()=>c("active")),r=t(e,"hover",24,()=>c("hover")),C=t(e,"color",24,()=>c("color")),J=t(e,"fill",24,()=>c("fill")),Y=t(e,"regionLabel",24,()=>c("regionLabel"));const Z="flex-auto",$="text-base text-center cursor-pointer",ee="opacity-50 cursor-not-allowed";let S=f();function te(l){["Enter","Space"].includes(l.code)&&(l.preventDefault(),i(S).click())}function ae(){return delete k.class,k}y(()=>(a(m()),a(g())),()=>{u(d,m()===g())}),y(()=>(i(d),a(R()),a(C()),a(J()),a(r())),()=>{u(b,i(d)?`${R()} ${C()} ${J()}`:r())}),y(()=>a(n),()=>{u(v,n.disabled?ee:"")}),y(()=>{},()=>{u(x,`${Z}`)}),y(()=>(a(B()),a(p()),i(b),i(v),a(n)),()=>{u(w,`${$} ${B()} ${p()} ${i(b)} ${i(v)} ${n.class??""}`)}),T(),X();var z=ce(),o=E(z),j=E(o),h=E(j);de(h,l=>u(S,l),()=>i(S)),oe(h);let M;q(()=>M=ne(h,M,{type:"radio",name:A(),value:m(),...ae(),tabindex:"-1"})),I(j);var le=ie(j,2);V(le,e,"default",{},null),I(o),I(z),q(()=>{F(z,`radio-label ${i(x)??""} ${Y()??""}`),F(o,`radio-item ${i(w)??""}`),P(o,"aria-checked",i(d)),P(o,"aria-label",K()),P(o,"title",G())}),re(D,[],h,()=>(m(),g()),g),_("click",h,function(l){L.call(this,e,l)}),_("change",h,function(l){L.call(this,e,l)}),_("keydown",o,te),_("keydown",o,function(l){L.call(this,e,l)}),_("keyup",o,function(l){L.call(this,e,l)}),_("keypress",o,function(l){L.call(this,e,l)}),N(W,z),U()}export{we as R,De as a};