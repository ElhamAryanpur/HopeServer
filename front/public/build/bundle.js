var app=function(){"use strict";function t(){}function n(t){return t()}function e(){return Object.create(null)}function o(t){t.forEach(n)}function s(t){return"function"==typeof t}function r(t,n){return t!=t?n==n:t!==n||t&&"object"==typeof t||"function"==typeof t}function u(t,n,e){t.insertBefore(n,e||null)}function c(t){t.parentNode.removeChild(t)}function l(t){return document.createElement(t)}function a(){return t=" ",document.createTextNode(t);var t}function f(t,n,e){null==e?t.removeAttribute(n):t.getAttribute(n)!==e&&t.setAttribute(n,e)}let i;function b(t){i=t}const d=[],p=[],m=[],h=[],g=Promise.resolve();let $=!1;function x(t){m.push(t)}let y=!1;const T=new Set;function _(){if(!y){y=!0;do{for(let t=0;t<d.length;t+=1){const n=d[t];b(n),v(n.$$)}for(d.length=0;p.length;)p.pop()();for(let t=0;t<m.length;t+=1){const n=m[t];T.has(n)||(T.add(n),n())}m.length=0}while(d.length);for(;h.length;)h.pop()();$=!1,y=!1,T.clear()}}function v(t){if(null!==t.fragment){t.update(),o(t.before_update);const n=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,n),t.after_update.forEach(x)}}const E=new Set;function w(t,n){-1===t.$$.dirty[0]&&(d.push(t),$||($=!0,g.then(_)),t.$$.dirty.fill(0)),t.$$.dirty[n/31|0]|=1<<n%31}function k(r,u,l,a,f,d,p=[-1]){const m=i;b(r);const h=u.props||{},g=r.$$={fragment:null,ctx:null,props:d,update:t,not_equal:f,bound:e(),on_mount:[],on_destroy:[],before_update:[],after_update:[],context:new Map(m?m.$$.context:[]),callbacks:e(),dirty:p};let $=!1;if(g.ctx=l?l(r,h,(t,n,...e)=>{const o=e.length?e[0]:n;return g.ctx&&f(g.ctx[t],g.ctx[t]=o)&&(g.bound[t]&&g.bound[t](o),$&&w(r,t)),n}):[],g.update(),$=!0,o(g.before_update),g.fragment=!!a&&a(g.ctx),u.target){if(u.hydrate){const t=function(t){return Array.from(t.childNodes)}(u.target);g.fragment&&g.fragment.l(t),t.forEach(c)}else g.fragment&&g.fragment.c();u.intro&&((y=r.$$.fragment)&&y.i&&(E.delete(y),y.i(T))),function(t,e,r){const{fragment:u,on_mount:c,on_destroy:l,after_update:a}=t.$$;u&&u.m(e,r),x(()=>{const e=c.map(n).filter(s);l?l.push(...e):o(e),t.$$.on_mount=[]}),a.forEach(x)}(r,u.target,u.anchor),_()}var y,T;b(m)}function A(n){let e,o,s,r;return{c(){e=l("link"),s=a(),r=l("table"),r.innerHTML='<tr><td class="card menu svelte-13p67tl"><table class="menu-main svelte-13p67tl"><tr><div class="menu-list svelte-13p67tl"><button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button> \n            <button class="box">Test1</button></div></tr></table></td> \n    <td class="card">REEE</td></tr>',f(e,"rel","stylesheet"),f(e,"href",o="/themes/"+n[0]+".css"),f(r,"class","table unselectable svelte-13p67tl")},m(t,n){!function(t,n){t.appendChild(n)}(document.head,e),u(t,s,n),u(t,r,n)},p(t,[n]){1&n&&o!==(o="/themes/"+t[0]+".css")&&f(e,"href",o)},i:t,o:t,d(t){c(e),t&&c(s),t&&c(r)}}}function N(t,n,e){let o=localStorage.getItem("theme");return null===o&&(o="blue"),[o]}return new class extends class{$destroy(){!function(t,n){const e=t.$$;null!==e.fragment&&(o(e.on_destroy),e.fragment&&e.fragment.d(n),e.on_destroy=e.fragment=null,e.ctx=[])}(this,1),this.$destroy=t}$on(t,n){const e=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return e.push(n),()=>{const t=e.indexOf(n);-1!==t&&e.splice(t,1)}}$set(){}}{constructor(t){super(),k(this,t,N,A,r,{})}}({target:document.body})}();
//# sourceMappingURL=bundle.js.map
