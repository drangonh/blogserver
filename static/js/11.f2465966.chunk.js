(this.webpackJsonpmyblog=this.webpackJsonpmyblog||[]).push([[11],{234:function(e,t,a){},251:function(e,t,a){"use strict";a.r(t);var n,r=a(53),c=a.n(r),s=a(54),i=a(4),o=a(10),u=a(12),l=a(11),p=a(0),d=a.n(p),f=(a(234),a(50)),m=a(15),h=a(235),v=a.n(h);function g(e,t,a,n,r){return b.apply(this,arguments)}function b(){return(b=Object(s.a)(c.a.mark((function e(t,a,n,r,s){var i,o,u,l,p;return c.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Object(f.a)("oss/ossInfo",{});case 2:return(o=e.sent)&&o.data&&(u=o.data,i=new v.a({region:u.endpoint,accessKeyId:u.accessKeyId,accessKeySecret:u.accessKeySecret,bucket:u.bucketName})),l="photo/",2==s?l="language/":3==s&&(l="article/"),e.prev=6,e.next=9,i.put(l+a,t);case 9:p=e.sent,n(p),e.next=18;break;case 14:e.prev=14,e.t0=e.catch(6),r(e.t0);case 18:case"end":return e.stop()}}),e,null,[[6,14]])})))).apply(this,arguments)}var k=Object(m.b)("header")(n=Object(m.c)(n=function(e){Object(u.a)(n,e);var t=Object(l.a)(n);function n(e){var a;return Object(i.a)(this,n),(a=t.call(this,e)).addImage=function(){a.input.click()},a.handleImageChange=function(){var e=Object(s.a)(c.a.mark((function e(t){var n;return c.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:g(n=t.target.files[0],n.name,(function(e){a.setState({photo:e.url})}));case 2:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),a.upload=Object(s.a)(c.a.mark((function e(){var t,n,r,s;return c.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=a.props.header.person,n={avatar:a.state.photo,email:a.state.email,nickName:a.state.nickName,description:a.state.description,uid:t&&t.uid},e.next=5,Object(f.b)("/profile/editUserProfile",n);case 5:(r=e.sent).data&&(s=a.props.header,a.email.value="",s.changeInfo(r.data),a.props.history.push("/home"),localStorage.setItem("userInfo",JSON.stringify(r.data)));case 8:case"end":return e.stop()}}),e)}))),a.state={login:!0,photo:"",email:"",nickName:"",description:""},a}return Object(o.a)(n,[{key:"componentDidMount",value:function(){var e=this.props.header.person;e&&this.setState({photo:e.avatar,email:e.email,nickName:e.nickName,description:e.description})}},{key:"render",value:function(){var e=this,t=this.state.photo;return d.a.createElement("div",{id:"updateInfo"},d.a.createElement("div",{className:"box_6747X3"},d.a.createElement("p",{className:"text_6746X3"},"\u8bf7\u586b\u5199\u57fa\u672c\u4fe1\u606f"),d.a.createElement("div",{onClick:this.addImage},d.a.createElement("input",{style:{display:"none"},ref:function(t){e.input=t},type:"file",accept:"image/*",onChange:this.handleImageChange}),d.a.createElement("img",{className:"default",src:0!=t.length?t:a(62)})),d.a.createElement("div",{className:"text_12X1"},d.a.createElement("input",{placeholder:"\u8bf7\u8f93\u5165\u6635\u79f0",type:"text",className:"editValue",ref:function(t){return e.nickName=t},value:this.state.nickName,onChange:function(t){e.setState({nickName:t.target.value})}})),d.a.createElement("div",{className:"text_12X1"},d.a.createElement("input",{placeholder:"\u8bf7\u8f93\u5165\u90ae\u7bb1",type:"text",className:"editValue",ref:function(t){return e.email=t},value:this.state.email,onChange:function(t){e.setState({email:t.target.value})}})),d.a.createElement("div",{className:"text_12X2"},d.a.createElement("textarea",{placeholder:"\u8bf7\u8f93\u5165\u4e2a\u4eba\u4ecb\u7ecd",type:"text",className:"editValue",ref:function(t){return e.info=t},value:this.state.description,onChange:function(t){e.setState({description:t.target.value})}})),d.a.createElement("div",{onClick:this.upload,className:"text_18X1 uac_jc"},"\u5b8c \u6210")))}}]),n}(d.a.Component))||n)||n;t.default=k},50:function(e,t,a){"use strict";a.d(t,"a",(function(){return c})),a.d(t,"b",(function(){return s}));var n=a(60),r=a.n(n);a(61);function c(e,t){return new Promise((function(a,n){r.a.get(e,{params:t}).then((function(e){a(e.data)})).catch((function(e){n(e.data)}))}))}function s(e,t){return new Promise((function(a,n){r.a.post(e,t).then((function(e){a(e.data)})).catch((function(e){n(e.data)}))}))}r.a.defaults.baseURL="https://www.dontheway.com:8880/blog/",r.a.defaults.withCredentials=!0,r.a.defaults.timeout=1e4,r.a.defaults.headers={"Content-Type":"application/x-www-form-urlencoded"},r.a.interceptors.request.use((function(e){return e}),(function(e){return Promise.error(e)})),r.a.interceptors.response.use((function(e){return 200===e.status?Promise.resolve(e):Promise.reject(e)}),(function(e){if(e.response.status)return e.response.status,Promise.reject(e.response)}))},62:function(e,t,a){e.exports=a.p+"static/media/defualt.8e05f6fe.jpeg"}}]);
//# sourceMappingURL=11.f2465966.chunk.js.map