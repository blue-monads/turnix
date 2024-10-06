import{a as x,t as q,c as we}from"../chunks/disclose-version.wcsHG7oF.js";import{p as _e,i as j,g as Ce,f as N,t as Se,a as Pe,s as A,b as p,h as U,c as P,e as ee,r as T,d as te}from"../chunks/runtime.D6lgFUio.js";import{d as Te,f as Ie,g as Ue}from"../chunks/utils.Dh-XvBAd.js";import{i as ae}from"../chunks/if.lkQ-IRrr.js";import{s as $e}from"../chunks/attributes.BFGPVQjG.js";import{p as Q}from"../chunks/proxy.r5-P5EHK.js";import{b as qe}from"../chunks/this.OuUT1k8T.js";import{s as Be,a as je}from"../chunks/store.BaapSzAp.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as Le}from"../chunks/AppBar.3srT6C5i.js";import{T as Re,a as re,h as Xe}from"../chunks/index.iPCAaS7f.js";import{i as De,c as Ne,f as Ae,b as Ee,t as n,L as Me,d as Ze,e as ze,E as We,g as Ve,h as Ye,k as Fe,C as se}from"../chunks/index.1_1nzF9G.js";import{p as Ge}from"../chunks/stores.D4Dcgzbh.js";import{N as Ke}from"../chunks/index.Duto3aX0.js";import{S as oe}from"../chunks/SvgIcon.Cj01WnO8.js";const He=36,ne=1,Je=2,L=3,W=4,et=5,tt=6,at=7,rt=8,st=9,ot=10,nt=11,lt=12,it=13,ct=14,dt=15,ft=16,mt=17,le=18,ut=19,he=20,ge=21,ie=22,pt=23,ht=24;function Y(t){return t>=65&&t<=90||t>=97&&t<=122||t>=48&&t<=57}function gt(t){return t>=48&&t<=57||t>=97&&t<=102||t>=65&&t<=70}function $(t,e,r){for(let a=!1;;){if(t.next<0)return;if(t.next==e&&!a){t.advance();return}a=r&&!a&&t.next==92,t.advance()}}function vt(t,e){e:for(;;){if(t.next<0)return;if(t.next==36){t.advance();for(let r=0;r<e.length;r++){if(t.next!=e.charCodeAt(r))continue e;t.advance()}if(t.next==36){t.advance();return}}else t.advance()}}function Ot(t,e){let r="[{<(".indexOf(String.fromCharCode(e)),a=r<0?e:"]}>)".charCodeAt(r);for(;;){if(t.next<0)return;if(t.next==a&&t.peek(1)==39){t.advance(2);return}t.advance()}}function F(t,e){for(;!(t.next!=95&&!Y(t.next));)e!=null&&(e+=String.fromCharCode(t.next)),t.advance();return e}function bt(t){if(t.next==39||t.next==34||t.next==96){let e=t.next;t.advance(),$(t,e,!1)}else F(t)}function ce(t,e){for(;t.next==48||t.next==49;)t.advance();e&&t.next==e&&t.advance()}function de(t,e){for(;;){if(t.next==46){if(e)break;e=!0}else if(t.next<48||t.next>57)break;t.advance()}if(t.next==69||t.next==101)for(t.advance(),(t.next==43||t.next==45)&&t.advance();t.next>=48&&t.next<=57;)t.advance()}function fe(t){for(;!(t.next<0||t.next==10);)t.advance()}function I(t,e){for(let r=0;r<e.length;r++)if(e.charCodeAt(r)==t)return!0;return!1}const V=` 	\r
`;function ve(t,e,r){let a=Object.create(null);a.true=a.false=et,a.null=a.unknown=tt;for(let s of t.split(" "))s&&(a[s]=he);for(let s of e.split(" "))s&&(a[s]=ge);for(let s of(r||"").split(" "))s&&(a[s]=ht);return a}const Oe="array binary bit boolean char character clob date decimal double float int integer interval large national nchar nclob numeric object precision real smallint time timestamp varchar varying ",be="absolute action add after all allocate alter and any are as asc assertion at authorization before begin between both breadth by call cascade cascaded case cast catalog check close collate collation column commit condition connect connection constraint constraints constructor continue corresponding count create cross cube current current_date current_default_transform_group current_transform_group_for_type current_path current_role current_time current_timestamp current_user cursor cycle data day deallocate declare default deferrable deferred delete depth deref desc describe descriptor deterministic diagnostics disconnect distinct do domain drop dynamic each else elseif end end-exec equals escape except exception exec execute exists exit external fetch first for foreign found from free full function general get global go goto grant group grouping handle having hold hour identity if immediate in indicator initially inner inout input insert intersect into is isolation join key language last lateral leading leave left level like limit local localtime localtimestamp locator loop map match method minute modifies module month names natural nesting new next no none not of old on only open option or order ordinality out outer output overlaps pad parameter partial path prepare preserve primary prior privileges procedure public read reads recursive redo ref references referencing relative release repeat resignal restrict result return returns revoke right role rollback rollup routine row rows savepoint schema scroll search second section select session session_user set sets signal similar size some space specific specifictype sql sqlexception sqlstate sqlwarning start state static system_user table temporary then timezone_hour timezone_minute to trailing transaction translation treat trigger under undo union unique unnest until update usage user using value values view when whenever where while with without work write year zone ",G={backslashEscapes:!1,hashComments:!1,spaceAfterDashes:!1,slashComments:!1,doubleQuotedStrings:!1,doubleDollarQuotedStrings:!1,unquotedBitLiterals:!1,treatBitsAsBytes:!1,charSetCasts:!1,plsqlQuotingMechanism:!1,operatorChars:"*+-%<>!=&|~^/",specialVar:"?",identifierQuotes:'"',caseInsensitiveIdentifiers:!1,words:ve(be,Oe)};function kt(t,e,r,a){let s={};for(let o in G)s[o]=(t.hasOwnProperty(o)?t:G)[o];return e&&(s.words=ve(e,r||"",a)),s}function ke(t){return new We(e=>{var r;let{next:a}=e;if(e.advance(),I(a,V)){for(;I(e.next,V);)e.advance();e.acceptToken(He)}else if(a==36&&t.doubleDollarQuotedStrings){let s=F(e,"");e.next==36&&(e.advance(),vt(e,s),e.acceptToken(L))}else if(a==39||a==34&&t.doubleQuotedStrings)$(e,a,t.backslashEscapes),e.acceptToken(L);else if(a==35&&t.hashComments||a==47&&e.next==47&&t.slashComments)fe(e),e.acceptToken(ne);else if(a==45&&e.next==45&&(!t.spaceAfterDashes||e.peek(1)==32))fe(e),e.acceptToken(ne);else if(a==47&&e.next==42){e.advance();for(let s=1;;){let o=e.next;if(e.next<0)break;if(e.advance(),o==42&&e.next==47){if(s--,e.advance(),!s)break}else o==47&&e.next==42&&(s++,e.advance())}e.acceptToken(Je)}else if((a==101||a==69)&&e.next==39)e.advance(),$(e,39,!0),e.acceptToken(L);else if((a==110||a==78)&&e.next==39&&t.charSetCasts)e.advance(),$(e,39,t.backslashEscapes),e.acceptToken(L);else if(a==95&&t.charSetCasts)for(let s=0;;s++){if(e.next==39&&s>1){e.advance(),$(e,39,t.backslashEscapes),e.acceptToken(L);break}if(!Y(e.next))break;e.advance()}else if(t.plsqlQuotingMechanism&&(a==113||a==81)&&e.next==39&&e.peek(1)>0&&!I(e.peek(1),V)){let s=e.peek(1);e.advance(2),Ot(e,s),e.acceptToken(L)}else if(a==40)e.acceptToken(at);else if(a==41)e.acceptToken(rt);else if(a==123)e.acceptToken(st);else if(a==125)e.acceptToken(ot);else if(a==91)e.acceptToken(nt);else if(a==93)e.acceptToken(lt);else if(a==59)e.acceptToken(it);else if(t.unquotedBitLiterals&&a==48&&e.next==98)e.advance(),ce(e),e.acceptToken(ie);else if((a==98||a==66)&&(e.next==39||e.next==34)){const s=e.next;e.advance(),t.treatBitsAsBytes?($(e,s,t.backslashEscapes),e.acceptToken(pt)):(ce(e,s),e.acceptToken(ie))}else if(a==48&&(e.next==120||e.next==88)||(a==120||a==88)&&e.next==39){let s=e.next==39;for(e.advance();gt(e.next);)e.advance();s&&e.next==39&&e.advance(),e.acceptToken(W)}else if(a==46&&e.next>=48&&e.next<=57)de(e,!0),e.acceptToken(W);else if(a==46)e.acceptToken(ct);else if(a>=48&&a<=57)de(e,!1),e.acceptToken(W);else if(I(a,t.operatorChars)){for(;I(e.next,t.operatorChars);)e.advance();e.acceptToken(dt)}else if(I(a,t.specialVar))e.next==a&&e.advance(),bt(e),e.acceptToken(mt);else if(I(a,t.identifierQuotes))$(e,a,!1),e.acceptToken(ut);else if(a==58||a==44)e.acceptToken(ft);else if(Y(a)){let s=F(e,String.fromCharCode(a));e.acceptToken(e.next==46||e.peek(-s.length-1)==46?le:(r=t.words[s.toLowerCase()])!==null&&r!==void 0?r:le)}})}const xe=ke(G),xt=Ze.deserialize({version:14,states:"%vQ]QQOOO#wQRO'#DSO$OQQO'#CwO%eQQO'#CxO%lQQO'#CyO%sQQO'#CzOOQQ'#DS'#DSOOQQ'#C}'#C}O'UQRO'#C{OOQQ'#Cv'#CvOOQQ'#C|'#C|Q]QQOOQOQQOOO'`QQO'#DOO(xQRO,59cO)PQQO,59cO)UQQO'#DSOOQQ,59d,59dO)cQQO,59dOOQQ,59e,59eO)jQQO,59eOOQQ,59f,59fO)qQQO,59fOOQQ-E6{-E6{OOQQ,59b,59bOOQQ-E6z-E6zOOQQ,59j,59jOOQQ-E6|-E6|O+VQRO1G.}O+^QQO,59cOOQQ1G/O1G/OOOQQ1G/P1G/POOQQ1G/Q1G/QP+kQQO'#C}O+rQQO1G.}O)PQQO,59cO,PQQO'#Cw",stateData:",[~OtOSPOSQOS~ORUOSUOTUOUUOVROXSOZTO]XO^QO_UO`UOaPObPOcPOdUOeUOfUOgUOhUO~O^]ORvXSvXTvXUvXVvXXvXZvX]vX_vX`vXavXbvXcvXdvXevXfvXgvXhvX~OsvX~P!jOa_Ob_Oc_O~ORUOSUOTUOUUOVROXSOZTO^tO_UO`UOa`Ob`Oc`OdUOeUOfUOgUOhUO~OWaO~P$ZOYcO~P$ZO[eO~P$ZORUOSUOTUOUUOVROXSOZTO^QO_UO`UOaPObPOcPOdUOeUOfUOgUOhUO~O]hOsoX~P%zOajObjOcjO~O^]ORkaSkaTkaUkaVkaXkaZka]ka_ka`kaakabkackadkaekafkagkahka~Oska~P'kO^]O~OWvXYvX[vX~P!jOWnO~P$ZOYoO~P$ZO[pO~P$ZO^]ORkiSkiTkiUkiVkiXkiZki]ki_ki`kiakibkickidkiekifkigkihki~Oski~P)xOWkaYka[ka~P'kO]hO~P$ZOWkiYki[ki~P)xOasObsOcsO~O",goto:"#hwPPPPPPPPPPPPPPPPPPPPPPPPPPx||||!Y!^!d!xPPP#[TYOZeUORSTWZbdfqT[OZQZORiZSWOZQbRQdSQfTZgWbdfqQ^PWk^lmrQl_Qm`RrseVORSTWZbdfq",nodeNames:"⚠ LineComment BlockComment String Number Bool Null ( ) { } [ ] ; . Operator Punctuation SpecialVar Identifier QuotedIdentifier Keyword Type Bits Bytes Builtin Script Statement CompositeIdentifier Parens Braces Brackets Statement",maxTerm:38,nodeProps:[["isolate",-4,1,2,3,19,""]],skippedNodes:[0,1,2],repeatNodeCount:3,tokenData:"RORO",tokenizers:[0,xe],topRules:{Script:[0,25]},tokenPrec:0});function K(t){let e=t.cursor().moveTo(t.from,-1);for(;/Comment/.test(e.name);)e.moveTo(e.from,-1);return e.node}function X(t,e){let r=t.sliceString(e.from,e.to),a=/^([`'"])(.*)\1$/.exec(r);return a?a[2]:r}function E(t){return t&&(t.name=="Identifier"||t.name=="QuotedIdentifier")}function Qt(t,e){if(e.name=="CompositeIdentifier"){let r=[];for(let a=e.firstChild;a;a=a.nextSibling)E(a)&&r.push(X(t,a));return r}return[X(t,e)]}function me(t,e){for(let r=[];;){if(!e||e.name!=".")return r;let a=K(e);if(!E(a))return r;r.unshift(X(t,a)),e=K(a)}}function yt(t,e){let r=Fe(t).resolveInner(e,-1),a=_t(t.doc,r);return r.name=="Identifier"||r.name=="QuotedIdentifier"||r.name=="Keyword"?{from:r.from,quoted:r.name=="QuotedIdentifier"?t.doc.sliceString(r.from,r.from+1):null,parents:me(t.doc,K(r)),aliases:a}:r.name=="."?{from:e,quoted:null,parents:me(t.doc,r),aliases:a}:{from:e,quoted:null,parents:[],empty:!0,aliases:a}}const wt=new Set("where group having order union intersect except all distinct limit offset fetch for".split(" "));function _t(t,e){let r;for(let s=e;!r;s=s.parent){if(!s)return null;s.name=="Statement"&&(r=s)}let a=null;for(let s=r.firstChild,o=!1,m=null;s;s=s.nextSibling){let c=s.name=="Keyword"?t.sliceString(s.from,s.to).toLowerCase():null,d=null;if(!o)o=c=="from";else if(c=="as"&&m&&E(s.nextSibling))d=X(t,s.nextSibling);else{if(c&&wt.has(c))break;m&&E(s)&&(d=X(t,s))}d&&(a||(a=Object.create(null)),a[d]=Qt(t,m)),m=/Identifier$/.test(s.name)?s:null}return a}function Ct(t,e){return t?e.map(r=>Object.assign(Object.assign({},r),{label:r.label[0]==t?r.label:t+r.label+t,apply:void 0})):e}const St=/^\w*$/,Pt=/^[`'"]?\w*[`'"]?$/;function ue(t){return t.self&&typeof t.self.label=="string"}class H{constructor(e,r){this.idQuote=e,this.idCaseInsensitive=r,this.list=[],this.children=void 0}child(e){let r=this.children||(this.children=Object.create(null)),a=r[e];return a||(e&&!this.list.some(s=>s.label==e)&&this.list.push(pe(e,"type",this.idQuote,this.idCaseInsensitive)),r[e]=new H(this.idQuote,this.idCaseInsensitive))}maybeChild(e){return this.children?this.children[e]:null}addCompletion(e){let r=this.list.findIndex(a=>a.label==e.label);r>-1?this.list[r]=e:this.list.push(e)}addCompletions(e){for(let r of e)this.addCompletion(typeof r=="string"?pe(r,"property",this.idQuote,this.idCaseInsensitive):r)}addNamespace(e){Array.isArray(e)?this.addCompletions(e):ue(e)?this.addNamespace(e.children):this.addNamespaceObject(e)}addNamespaceObject(e){for(let r of Object.keys(e)){let a=e[r],s=null,o=r.replace(/\\?\./g,c=>c=="."?"\0":c).split("\0"),m=this;ue(a)&&(s=a.self,a=a.children);for(let c=0;c<o.length;c++)s&&c==o.length-1&&m.addCompletion(s),m=m.child(o[c].replace(/\\\./g,"."));m.addNamespace(a)}}}function pe(t,e,r,a){return new RegExp("^[a-z_][a-z_\\d]*$",a?"i":"").test(t)?{label:t,type:e}:{label:t,type:e,apply:r+t+r}}function Tt(t,e,r,a,s,o){var m;let c=((m=o==null?void 0:o.spec.identifierQuotes)===null||m===void 0?void 0:m[0])||'"',d=new H(c,!!(o!=null&&o.spec.caseInsensitiveIdentifiers)),k=s?d.child(s):null;return d.addNamespace(t),e&&(k||d).addCompletions(e),r&&d.addCompletions(r),k&&d.addCompletions(k.list),a&&d.addCompletions((k||d).child(a).list),v=>{let{parents:O,from:Z,quoted:y,empty:D,aliases:w}=yt(v.state,v.pos);if(D&&!v.explicit)return null;w&&O.length==1&&(O=w[O[0]]||O);let h=d;for(let _ of O){for(;!h.children||!h.children[_];)if(h==d&&k)h=k;else if(h==k&&a)h=h.child(a);else return null;let C=h.maybeChild(_);if(!C)return null;h=C}let z=y&&v.state.sliceDoc(v.pos,v.pos+1)==y,B=h.list;return h==d&&w&&(B=B.concat(Object.keys(w).map(_=>({label:_,type:"constant"})))),{from:Z,to:z?v.pos+1:void 0,options:Ct(y,B),validFor:y?Pt:St}}}function It(t,e){let r=Object.keys(t).map(a=>({label:e?a.toUpperCase():a,type:t[a]==ge?"type":t[a]==he?"keyword":"variable",boost:-1}));return Ve(["QuotedIdentifier","SpecialVar","String","LineComment","BlockComment","."],Ye(r))}let Ut=xt.configure({props:[De.add({Statement:Ne()}),Ae.add({Statement(t,e){return{from:Math.min(t.from+100,e.doc.lineAt(t.from).to),to:t.to}},BlockComment(t){return{from:t.from+2,to:t.to-2}}}),Ee({Keyword:n.keyword,Type:n.typeName,Builtin:n.standard(n.name),Bits:n.number,Bytes:n.string,Bool:n.bool,Null:n.null,Number:n.number,String:n.string,Identifier:n.name,QuotedIdentifier:n.special(n.string),SpecialVar:n.special(n.name),LineComment:n.lineComment,BlockComment:n.blockComment,Operator:n.operator,"Semi Punctuation":n.punctuation,"( )":n.paren,"{ }":n.brace,"[ ]":n.squareBracket})]});class M{constructor(e,r,a){this.dialect=e,this.language=r,this.spec=a}get extension(){return this.language.extension}static define(e){let r=kt(e,e.keywords,e.types,e.builtin),a=ze.define({name:"sql",parser:Ut.configure({tokenizers:[{from:xe,to:ke(r)}]}),languageData:{commentTokens:{line:"--",block:{open:"/*",close:"*/"}},closeBrackets:{brackets:["(","[","{","'",'"',"`"]}}});return new M(r,a,e)}}function $t(t,e=!1){return It(t.dialect.words,e)}function qt(t,e=!1){return t.language.data.of({autocomplete:$t(t,e)})}function Bt(t){return t.schema?Tt(t.schema,t.tables,t.schemas,t.defaultTable,t.defaultSchema,t.dialect||J):()=>null}function jt(t){return t.schema?(t.dialect||J).language.data.of({autocomplete:Bt(t)}):[]}function Lt(t={}){let e=t.dialect||J;return new Me(e.language,[jt(t),qt(e,!!t.upperCaseKeywords)])}const J=M.define({}),Rt=M.define({keywords:be+"abort analyze attach autoincrement conflict database detach exclusive fail glob ignore index indexed instead isnull notnull offset plan pragma query raise regexp reindex rename replace temp vacuum virtual",types:Oe+"bool blob long longblob longtext medium mediumblob mediumint mediumtext tinyblob tinyint tinytext text bigint int2 int8 unsigned signed real",builtin:"auth backup bail changes clone databases dbinfo dump echo eqp explain fullschema headers help import imposter indexes iotrace lint load log mode nullvalue once print prompt quit restore save scanstats separator shell show stats system tables testcase timeout timer trace vfsinfo vfslist vfsname width",operatorChars:"*+-%<>!=&|/~",identifierQuotes:'`"',specialVar:"@:?$"});let Xt=`
    
-- query_name: list_accounts
select * from __project__Accounts;

































`;const Dt=`

<!DOCTYPE html>
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
        <h1>Reports Editor Example</h1> 
        <p>This is a sample preview</p>
        
    </div>
    <div class="container-xl mt-5" x-data="accountTable()" id="ac_table">
        
    <div class="container">
        <button id="start" class="btn btn-primary" type="button" x-on:click="fetchAccounts">
            <span class="icon icon-thumbs-up"></span>
            Start
        </button>    
    </div>


        <div x-show="!loading" class="table-responsive">
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Info</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                </thead>
                <tbody>
                    <template x-for="account in accounts" :key="account.id">
                        <tr>
                            <td x-text="account.id"></td>
                            <td x-text="account.name"></td>
                            <td x-text="account.acc_type" class="text-capitalize"></td>
                            <td x-text="account.info"></td>
                            <td x-text="formatDate(account.created_at)"></td>
                            <td x-text="formatDate(account.updated_at)"></td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </div>

    <script>
    function accountTable() {
      return {
          accounts: [],
          loading: true,
          async fetchAccounts() {
              this.loading = true;
              try {
                  const data = await sendMessage({
                      type: "sql_query",
                      name: "list_accounts",
                      data: [],
                  });
                  console.log("accounts", data);

                  this.accounts = data.data;
              } catch (error) {
                  console.error("Error fetching accounts:", error);
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
</html>







`;var Nt=q('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Reports</li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),At=(t,e,r)=>{U(e,Q(p(r)))},Et=(t,e)=>{var r;(r=p(e).contentWindow)==null||r.print()},Mt=q('<button class="btn variant-filled btn-sm"><!> Run</button> <button class="btn variant-filled-secondary btn-sm"><!> Print</button>',1),Zt=q("<span>SQL Query</span>"),zt=q("<span>UI</span>"),Wt=q("<!> <!>",1),Vt=q('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),Yt=q('<!> <div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><div class="card p-2 h-full w-full"><iframe title="preview" width="100%" height="100%" class="border-green-200 w-full h-full transition-all" allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking" sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"></iframe></div></div></div>',1);function da(t,e){_e(e,!0);const r=Be(),a=()=>je(Ge,"$page",r);let s=j(void 0),o=j(0),m=j(Q(Xt)),c=j(Q(Dt)),d=j("");const k=a().params.pid,v=Ce("__api__");Ke(v);let O=j(void 0);const Z=async g=>{var u,i,b;const f=g.data;if(console.log("onFrameMessage",f),f.type==="sql_query"){console.log("sql_query",f);try{const l=await v.runProjectSQL(k,{input:p(m),name:f.name,data:f.data});if(l.status!==200){(u=p(O))==null||u.postMessage({msgId:f.msgId,data:{msg:"error",data:l.data}});return}(i=p(O))==null||i.postMessage({msgId:f.msgId,data:{msg:"success",data:l.data}})}catch{}}else f.type==="api_call"?console.log("api_call",f):f.type==="ping"&&((b=p(O))==null||b.postMessage({msgId:f.msgId,data:{msg:"pong"}}))};var y=Yt(),D=N(y);Le(D,{$$slots:{lead:(g,f)=>{var u=Nt();x(g,u)},trail:(g,f)=>{var u=Mt(),i=N(u);i.__click=[At,d,c];var b=P(i);oe(b,{className:"w-4 h-4",name:"play"}),ee(),T(i);var l=A(i,2);l.__click=[Et,s];var S=P(l);oe(S,{className:"w-4 h-4",name:"printer"}),ee(),T(l),x(g,u)}}});var w=A(D,2),h=P(w),z=P(h);Re(z,{children:(g,f)=>{var u=Wt(),i=N(u);re(i,{padding:"ml-6 p-2",get group(){return p(o)},set group(l){U(o,Q(l))},name:"SQL Query",value:0,children:(l,S)=>{var R=Zt();x(l,R)},$$slots:{default:!0}});var b=A(i,2);re(b,{get group(){return p(o)},set group(l){U(o,Q(l))},name:"UI",value:1,children:(l,S)=>{var R=zt();x(l,R)},$$slots:{default:!0}}),x(g,u)},$$slots:{default:!0,panel:(g,f)=>{var u=Vt(),i=P(u);ae(i,()=>p(o)===0,b=>{var l=te(()=>Lt({dialect:Rt}));se(b,{get value(){return p(m)},set value(S){U(m,Q(S))},get lang(){return p(l)}})},b=>{var l=we(),S=N(l);ae(S,()=>p(o)===1,R=>{var Qe=te(Xe);se(R,{get value(){return p(c)},set value(ye){U(c,Q(ye))},get lang(){return p(Qe)}})},null,!0),x(b,l)}),T(u),x(g,u)}}}),T(h);var B=A(h,2),_=P(B),C=P(_);qe(C,g=>U(s,g),()=>p(s)),T(_),T(B),T(w),Se(()=>$e(C,"srcdoc",p(d))),Ie("load",C,g=>{var f,u;try{let i=new MessageChannel;i.port1.onmessage=Z,U(O,Q(i.port1)),console.log("chan.port2 type:",i.port2 instanceof MessagePort),(u=(f=p(s))==null?void 0:f.contentWindow)==null||u.postMessage("transfer_port","*",[i.port2])}catch(i){console.error("Error in postMessage:",i)}}),Ue(C),x(t,y),Pe()}Te(["click"]);export{da as component};
