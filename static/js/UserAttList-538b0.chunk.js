webpackJsonp([40],{2061:function(e,t,a){"use strict";function l(e){return e&&e.__esModule?e:{default:e}}function n(e){return function(t){new g.default(t).post((0,w.getUrl)("/templates/"+e+"/submit"),{},function(){t({type:h.default.SUBMIT_TEMPLATE,id:e}),b.toastr.success(v.default.translate("actions.submittemplate"))})}}function r(e){return function(t){new g.default(t).get((0,w.getUrl)("/products/"+e+"/templates"),function(e){t({type:h.default.GET_TEMPLATES,templates:e})})}}function o(e){return function(t){new g.default(t).destroy((0,w.getUrl)("/templates/"+e),function(){b.toastr.success(v.default.translate("actions.destroy-template")),t({type:h.default.DESTROY_TEMPLATE,id:e})})}}function i(e,t){return function(a){new g.default(a).post(p("/templates/"+e),t,function(){b.toastr.success(v.default.translate("actions.update-template"))})}}function s(e){return function(t){new g.default(t).post((0,w.getUrl)("/templates/preview"),e,function(e){t({type:h.default.PREVIEW_TEMPLATE,preview:e})})}}function u(e,t,a){return function(a){new g.default(a).post((0,w.getUrl)("/products/"+e+"/templates"),{templateId:t},function(e){b.toastr.success(v.default.translate("actions.create-template")),a({type:h.default.NEW_TEMPLATE,info:e})})}}function c(e){return function(t){new g.default(t).get(p("/templates/"+e),function(e){t({type:h.default.GET_TEMPLATE,info:e})})}}function p(e){return"?fromAdmin=true"===window.location.search&&(e="/admin"+e),e}function d(e,t){return function(a){new g.default(a).post((0,w.getUrl)("/templates/"+e+"/title"),{title:t},function(){a({type:h.default.UPDATE_TEMPLATE_TITLE,info:{title:t,id:e}}),b.toastr.success(v.default.translate("actions.update-title")+""+t)})}}function f(e){return function(t){new g.default(t).post("/templates/"+e+"/clone")}}Object.defineProperty(t,"__esModule",{value:!0}),t.submitTemplate=n,t.getTemplates=r,t.destroyTemplate=o,t.updateTemplate=i,t.previewTemplate=s,t.createTemplate=u,t.getTemplate=c,t.getTempUrl=p,t.updateTitle=d,t.cloneTemplate=f;var m=a(861),h=l(m),y=a(862),g=l(y),b=a(709),E=a(708),v=l(E),w=a(860)},2072:function(e,t){e.exports="data:img/png;base64,iVBORw0KGgoAAAANSUhEUgAAABIAAAASBAMAAACk4JNkAAAAMFBMVEVHcEwKbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4Kbe4B3N9aAAAAD3RSTlMAJfbm1bwRxz+WB6d6YLEl3BpnAAAAhElEQVQI12NgAAFG8wVgmoFLaLMahMX4icE+AcySfzyB5SNY6Cvjb4ZgECu/gEE+4DyQwfYFKGz6B8habwAk7h8AGvEZyOD4CSQWgYT6JwAJ5wVQIfYPIKENICO0gTpVQYY1CzkATQOxzjOocH0H2zmFgT2oAMxy75n6BeIi/f+eF8AMAD4wIPQ0HnV0AAAAAElFTkSuQmCC"},2083:function(e,t){"use strict";function a(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}Object.defineProperty(t,"__esModule",{value:!0});var l,n,r=(n=l=function e(){a(this,e)},l.fileType={allType:{file:[".ac3",".au",".mp2",".ogg",".flac",".ape",".wav",".mp3",".aac",".wma",".3gpp",".mp4",".mpeg",".mpg",".3gp",".wmv",".asf",".asx",".rm",".rmvb",".3gp",".mov",".m4v",".avi",".dat","mkv",".flv",".vob",".3gpp",".mp4",".mpeg",".mpg",".3gp",".wmv",".asf",".asx",".rm",".rmvb",".3gp",".mov",".m4v",".avi",".dat","mkv",".flv",".vob",".dwg",".dxf",".gif",".jp2",".jpe",".jpeg",".jpg",".png",".svf",".tif",".tiff",".bmp",".webp",".bmp",".pcx",".tif",".tga",".exif",".fpx",".svg",".psd",".cdr",".ico",".doc",".docx",".dot",".dtd",".js",".json",".mpp",".pdf",".pot",".pps",".ppt",".pptx",".rtf",".wdb",".wps",".xhtml",".xlc",".xlm",".xls",".xlt",".xlw",".xml",".css",".csv",".htm",".html",".txt",".rar",".zip",".arj",".z",".iso",".jar",".bz2",".gz",".tar",".ace",".lzh",".cab",".arj"]},audio:{file:[".ac3",".au",".mp2",".ogg",".flac",".ape",".wav",".mp3",".aac",".wma"],size:500},video:{file:[".3gpp",".mp4",".mpeg",".mpg",".3gp",".wmv",".asf",".asx",".rm",".rmvb",".3gp",".mov",".m4v",".avi",".dat","mkv",".flv",".vob"],size:500},pic:{file:[".dwg",".dxf",".gif",".jp2",".jpe",".jpeg",".jpg",".png",".svf",".tif",".tiff",".bmp",".webp",".bmp",".pcx",".tif",".tga",".exif",".fpx",".svg",".psd",".cdr",".ico"],size:100},txt:{file:[".doc",".docx",".dot",".dtd",".js",".json",".mpp",".pdf",".pot",".pps",".ppt",".pptx",".rtf",".wdb",".wps",".xhtml",".xlc",".xlm",".xls",".xlt",".xlw",".xml",".css",".csv",".htm",".html",".txt"],size:20},rar:{file:[".rar",".zip",".arj",".z",".iso",".jar",".bz2",".gz",".tar",".ace",".lzh",".cab",".arj"],size:200}},l.showType=function(e){var t=[".ac3",".au",".mp2",".ogg",".flac",".ape",".wav",".mp3",".aac",".wma"],a=[".3gpp",".mp4",".mpeg",".mpg",".3gp",".wmv",".asf",".asx",".rm",".rmvb",".3gp",".mov",".m4v",".avi",".dat","mkv",".flv",".vob"],l=[".dwg",".dxf",".gif",".jp2",".jpe",".jpeg",".jpg",".png",".svf",".tif",".tiff",".bmp",".webp",".bmp",".pcx",".tif",".tga",".exif",".fpx",".svg",".psd",".cdr",".ico"],n=[".doc",".docx",".dot",".dtd",".js",".json",".mpp",".pdf",".pot",".pps",".ppt",".pptx",".rtf",".wdb",".wps",".xhtml",".xlc",".xlm",".xls",".xlt",".xlw",".xml",".css",".csv",".htm",".html",".txt"],r=[".rar","zip","arj",".z",".iso",".jar",".bz2",".gz",".tar",".ace",".lzh",".cab",".arj"];return t.indexOf(e)>=0?"音频":a.indexOf(e)>=0?"视频":l.indexOf(e)>=0?"图片":n.indexOf(e)>=0?"文档":r.indexOf(e)>=0?"压缩包":"其他"},l.showFileType=function(e){switch(e){case"AUDIO":return"音频";case"VIDEO":return"视频";case"IMAGE":return"图片";case"DOC":return"文档";case"PACKAGES":return"其他";case"OTHER":return"其他";default:return"——"}},n);t.default=r},2615:function(e,t,a){"use strict";function l(e){return e&&e.__esModule?e:{default:e}}function n(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function r(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!=typeof t&&"function"!=typeof t?e:t}function o(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t)}Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i,s,u=function(){function e(e,t){for(var a=0;a<t.length;a++){var l=t[a];l.enumerable=l.enumerable||!1,l.configurable=!0,"value"in l&&(l.writable=!0),Object.defineProperty(e,l.key,l)}}return function(t,a,l){return a&&e(t.prototype,a),l&&e(t,l),t}}(),c=a(419),p=l(c),d=a(658),f=a(728),m=l(f),h=a(2616),y=l(h),g=a(2336),b=l(g),E=a(711),v=l(E),w=a(2246),T=l(w),N=a(869),x=a(869),A=a(865),I=l(A),S=(i=(0,d.connect)(function(e){return{user:e.user.kycs,balanceHolder:e.user.balanceHolder,activityTime:e.user.activityTime}}))(s=function(e){function t(e){n(this,t);var a=r(this,(t.__proto__||Object.getPrototypeOf(t)).call(this,e));return a.state={file:"",fileName:"",isHelp:!1},a.file="",a.fileName="",a.openHelp=function(){a.setState({isHelp:!0})},a.closeHelp=function(){a.setState({isHelp:!1})},a.textNodes=p.default.createElement("div",{className:"toolTip"},p.default.createElement("span",{className:"arrow att-arrow"}),p.default.createElement("i",{className:"iconfont font-console"}),p.default.createElement("div",{className:"context"},p.default.createElement("p",{className:"title"},"数据保全"),p.default.createElement("p",{className:"tipText"},"数据保全是保全网针对原创人员推出的防止侵权利器，保全网通过对接国家授时中心，在您上传原创作品的同时，进行作品实时固化，以方便您的作品在日后发生侵权时，及时证明您的作品优先时间。"),p.default.createElement("hr",null),p.default.createElement("p",{className:"tipsMore att-tips-more"},"数据保全的功能优势"),p.default.createElement("p",{className:"tips"},"时间戳证明"),p.default.createElement("p",{className:"tipsText"},"对接中国国家授时中心、苹果NTP服务，可靠的高精度的授时服务，保障时间的精准性。"),p.default.createElement("p",{className:"tips"},"作品保全"),p.default.createElement("p",{className:"tipsText"},"结合区块链技术，对上传的作品打散分布存储，保证作品的安全性与私密性。"),p.default.createElement("p",{className:"tips"},"司法出证"),p.default.createElement("p",{className:"tipsText"},"联合司法鉴定机构，可在线快速申请出具公平公正的司法鉴定书。"))),a}return o(t,e),u(t,[{key:"componentWillMount",value:function(){this.props.dispatch((0,x.getActivityTime)()),this.props.dispatch((0,N.currentUser)()),this.props.dispatch((0,N.getUserKycs)())}},{key:"render",value:function(){var e={maxWidth:1514,width:"calc(100vw - 416px)",minHeight:32,height:542,background:"#fff",boxShadow:"0 0 7px 3px rgba(0,0,0,0.05)",marginLeft:-667,padding:"30px 50px 60px",borderRadius:"3px",marginTop:-6},t={borderBottom:"10px solid #fff",left:658},a={marginLeft:"20px",color:"#1d8dfb",cursor:"pointer"},l=this.props.balanceHolder;return p.default.createElement("div",{className:"container-wrapper"},p.default.createElement("div",{className:"container member-container"},p.default.createElement("p",{className:"table-name"},"存证确权列表",p.default.createElement("span",{className:"price-tip"},"（当前剩余存证容量",p.default.createElement("span",{className:"count",style:{color:"#fe9e00"}},l&&l.attestation?""+l.attestation.total:"0"),"G","TEAM_SON"!==I.default.userIdentity&&"CLIENT_SON"!==I.default.userIdentity?p.default.createElement(m.default,{to:"/capacity?productType="+T.default.attestatton},p.default.createElement("span",{className:"to-buy",style:a},"立即购买")):"","）")),p.default.createElement("div",{className:(0,v.default)("description")},p.default.createElement("div",{style:{display:"flex",alignItems:"center",marginTop:0}},p.default.createElement("span",null,"自行上传原创内容存证，由保全网为您进行区块链数据固定。保全网不对您上传内容的真实性作验证。"),p.default.createElement("div",{onMouseOver:this.openHelp,onMouseOut:this.closeHelp,style:{position:"relative"}},p.default.createElement("span",{className:"help",style:{color:"#0a6dee"}},"了解更多"),p.default.createElement("span",{className:"iconHelp"},"?"),p.default.createElement(b.default,{style:e,borderBottom:t,children:this.textNodes,show:this.state.isHelp}))),p.default.createElement(m.default,{to:"/attestations",className:"to-add",style:{top:"-20px"}},"新增存证")),p.default.createElement("hr",null),p.default.createElement("div",{className:(0,v.default)("ctn")},p.default.createElement(y.default,{isPrivateUpload:!0,haveInput:"0",type:"user"}))))}}]),t}(c.Component))||s;t.default=S},2616:function(e,t,a){"use strict";function l(e){return e&&e.__esModule?e:{default:e}}function n(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function r(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!=typeof t&&"function"!=typeof t?e:t}function o(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t)}Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i,s,u=function(){function e(e,t){for(var a=0;a<t.length;a++){var l=t[a];l.enumerable=l.enumerable||!1,l.configurable=!0,"value"in l&&(l.writable=!0),Object.defineProperty(e,l.key,l)}}return function(t,a,l){return a&&e(t.prototype,a),l&&e(t,l),t}}(),c=a(419),p=l(c),d=a(658),f=a(708),m=l(f),h=a(2107),y=a(949),g=l(y),b=a(2061),E=a(2083),v=(l(E),a(1087)),w=a(869),T=a(896),N=a(709),x=a(711),A=l(x),I=a(728),S=l(I),O=a(2089),k=l(O),C=a(2291),L=l(C),M=a(862),j=l(M),_=a(2233),P=(l(_),a(865)),D=l(P),z="",B=void 0,U=(i=(0,d.connect)(function(e){return{user:e.user.kycs,attestations:e.attestation.user,summary:e.attestation.allSummary,params:e.router.params,template:e.template.all,attLabels:e.user.attLabels,pathname:e.router.location.pathname,operatorList:e.user.operatorList}}))(s=function(e){function t(e){n(this,t);var a=r(this,(t.__proto__||Object.getPrototypeOf(t)).call(this,e));return a.state={loading:!0,startDate:"",endDate:"",showIdentifyModal:!1,date:null,editModal:!1,selectedItem:"",showImage:!1,showDeleteModal:!1},a.file="",a.fileName="",a.date=function(e){return new Date(e).getTime()},a.showImage=function(e){z=j.default.getEndpoint("/attestations/"+e+"/images"),a.forceUpdate(),a.setState({showImage:!0})},a.closeImage=function(){z="",a.setState({showImage:!1})},a.doQuery=function(e){var t=""===a.state.startDate?"":a.date(a.state.startDate),l=""===a.state.endDate?"":0===a.date(a.state.endDate)?"":a.date(a.state.endDate),n=a.keyWord,r=a.productId,o=a.templateId,i=a.fileType,s=a.operate,u=a.fileLabel,c=a.props.pathname,p=void 0;"/org-attestations/list"===c?p=h.findAttestationByOrg:"/attestations-list"===c&&(p=h.findAllAttestationByUser),a.props.dispatch(p({pageNo:a.pageNo,pageSize:a.pageSize,productId:r,templateId:o,fileType:i,fileLabel:u,startDate:t,endDate:l,keyWord:n,operatorId:s},function(){return a.setState({loading:!1})}))},a.queryByProduct=function(e){var t=a.props,l=t.summary,n=t.type;if(a.productId!==e){if(a.productId=e,a.pageNo=0,a.keyWord="",a.templateId="",a.forceUpdate(),a.noTemplate=!1,"user"!=n&&""!=a.productId){var r=l.filter(function(e){return e.pId===a.productId});r.length>0&&1==r[0].isProduct?a.props.dispatch((0,b.getTemplates)(a.productId)):a.noTemplate=!0}a.doQuery(a.props)}},a.queryByTemplate=function(e){a.templateId!==e&&(a.templateId=e,a.pageNo=0,a.keyWord="",a.doQuery(a.props))},a.queryByFileType=function(e){a.fileType!==e&&(a.fileType=e,a.pageNo=0,a.keyWord="",a.doQuery(a.props))},a.queryByOperate=function(e){a.operate!==e&&(a.operate=e,a.pageNo=0,a.keyWord="",a.doQuery(a.props))},a.handleKeydown=function(e){13===e.keyCode&&a.handleAttestationQuery()},a.handleAutoKeydown=function(e){13===e.keyCode&&a.handleFileLabelQuery()},a.handleAttestationQuery=function(e){var t=a.props,l=t.isPrivateUpload,n=t.type;e&&e.preventDefault(),a.state.startDate="",a.state.endDate="",a.productId=l?"org"===n?"":"__private_system_upload":"",a.templateId="",a.fileType="",a.operate="",a.fileLabel="",a.setState({date:null}),a.forceUpdate(),a.doQuery(a.props)},a.handleFileLabelQuery=function(e){e&&e.preventDefault(),a.keyWord="",a.doQuery(a.props)},a.handlePageClick=function(e){a.pageNo=e.selected,a.doQuery(a.props)},a.handleKeyWordChange=function(e){a.keyWord=e,a.pageNo=0,a.forceUpdate()},a.getProductName=function(e){for(var t=a.props.summary,l=0;l<t.length;l++)if(t[l].pId===e)return t[l].pName},a.createFilter=function(e){var t=0;return function(l){if(0===l.value.toLowerCase().indexOf(e.toLowerCase()))return t+=1,a.newTag="",a.forceUpdate(),0===l.value.toLowerCase().indexOf(e.toLowerCase());0===t&&(a.newTag=e,a.forceUpdate())}},a.handleSelect=function(e){a.fileLabel=e.value,a.forceUpdate()},a.changeTag=function(e){a.fileLabel=e,a.forceUpdate()},a.queryByPageSize=function(e){a.pageSize!==e&&(a.pageSize=e,a.pageNo=0,a.doQuery(a.props))},a.queryByPageNo=function(e){a.pageNo!==e-1&&(a.pageNo=e-1,a.doQuery(a.props))},a.openEditModal=function(e,t){a.setState({editModal:e,selectedItem:t}),t.fileLabel&&(a.label=t.fileLabel,a.forceUpdate())},a.saveLabel=function(){""===a.label?N.toastr.error("标签不能为空"):a.label===a.state.selectedItem.fileLabel?N.toastr.success("标签未修改"):a.props.dispatch((0,w.editAttestationLabel)(a.state.selectedItem.id,a.label,function(){a.doQuery(a.props),a.label="",a.forceUpdate(),a.props.dispatch((0,w.getAttLabels)()),a.openEditModal(!1,"")}))},a.handleSelects=function(e){a.label=e.value,a.forceUpdate()},a.changeTags=function(e){var t=e.replace(/\s*/g,"");""!==t&&(a.label=t,a.forceUpdate())},a.changeWebTags=function(){setTimeout(function(){a.forceUpdate()},3e3),a.forceUpdate()},a.cellClick=function(e,t){"文件名称"===t.label&&"IMAGE"===e.fileType&&a.showImage(e.id)},a.canOpen=function(e){if(e){var t=e.substr(e.lastIndexOf(".")+1),a=t.toLowerCase();return"jpg"===a||"png"===a}return!1},a.selectedItems={},a.naid="",a.attId="",a.orderId="",a.resetCondition(),a.noTemplate=!1,a.pageSize=20,a.label="",a.status="",a.processLabel=[],a}return o(t,e),u(t,[{key:"resetCondition",value:function(){this.keyWord="",this.productId="",this.templateId="",this.fileType="",this.operate="",this.fileLabel="",this.pageNo=0,this.buttonDisable=!1}},{key:"componentDidMount",value:function(){this.doQuery(this.props),this.props.dispatch((0,h.findAllSummary)()),this.props.dispatch((0,w.getAttLabels)()),B="TEAM_MAIN"===D.default.userIdentity||"BOTH_MAIN"===D.default.userIdentity}},{key:"querySearch",value:function(e,t){if(this.processLabel&&this.processLabel.length>0){t(e?this.processLabel.filter(this.createFilter(e)):this.processLabel)}else{t(e?this.processLabel.filter(this.createFilter(e)):this.processLabel),this.newTag=e}}},{key:"render",value:function(){var e=this,t=this.props,l=t.attestations,n=t.summary,r=t.template,o=t.type,i=t.isPrivateUpload,s=t.attLabels,u=t.operatorList,c=g.default.get("yyyy/mm/dd hh:MM:ss");s.length>0&&(this.processLabel=s);var d=[{label:"文件名称",prop:"fileName",minWidth:"15%",render:function(e){return null!==e.fileName&&""!==e.fileName?p.default.createElement("p",{title:e.fileName,className:(0,A.default)("att-file-box",{underlineBlue:"IMAGE"===e.fileType})},p.default.createElement("span",{className:"att-file-name"},e.fileName.substr(0,e.fileName.lastIndexOf("."))),p.default.createElement("span",{className:"att-file-type"},e.fileName.substr(e.fileName.lastIndexOf(".")))):p.default.createElement("span",null,"——")}},{label:"文件大小",prop:"fileSize",minWidth:"10%",render:function(e){return p.default.createElement("span",null,e.fileSize)}},{label:"保全号",prop:"id",minWidth:"26%",render:function(e){return p.default.createElement("span",{title:e.id},e.id)}},{label:"文件标签",prop:"fileLabel",minWidth:"12%",render:function(t){return p.default.createElement("div",{className:"editAttLabel"},p.default.createElement("span",{title:null!==t.fileLabel&&""!==t.fileLabel?t.fileLabel:"——"},t.fileLabel),p.default.createElement("img",{src:a(2072),onClick:function(){return e.openEditModal(!0,t)}}))}},{label:"存证时间",prop:"createdAt",minWidth:"15%",sortable:!0,render:function(e){return null!==e.createdAt&&""!==e.createdAt?p.default.createElement("span",{title:c.format(e.createdAt)},c.format(e.createdAt)):p.default.createElement("span",null,"——")}},B?{label:"操作人",minWidth:"10%",render:function(e){return p.default.createElement("span",null,e.realName?e.realName:"")}}:{label:"操作人",prop:"status",minWidth:"0%",render:function(e){return p.default.createElement("span",null)}},{label:"操作",minWidth:"12%",render:function(t){return p.default.createElement("div",null,"IMAGE"!==t.fileType||!e.canOpen(t.fileName)||"OPEN"!==t.monitorStatus&&"RESTART"!==t.monitorStatus?"":p.default.createElement(S.default,{className:"href-link operate-view",to:"/monitor/certificate/"+t.monitorId},k.default.status(t.monitorStatus)),"IMAGE"!==t.fileType||!e.canOpen(t.fileName)||"CLOSED"!==t.monitorStatus&&"CANCELED"!==t.monitorStatus?"":[p.default.createElement(S.default,{className:"href-link operate-view",to:"/monitor/certificate/"+t.monitorId},k.default.status(t.monitorStatus))],"IMAGE"!==t.fileType||!e.canOpen(t.fileName)||"WAIT"!==t.monitorStatus&&""!==t.monitorStatus&&null!==t.monitorStatus?"":p.default.createElement(S.default,{to:"/monitor/certificate/add?attId="+t.id,className:"operate-blue"},"开启监测"),p.default.createElement("a",{target:"_blank",href:"/attestations/"+t.id,className:"operate-blue",style:{color:"#5488f9",margin:"0 10px"}},p.default.createElement("span",null,"查看证书")))}}];return p.default.createElement("div",{className:"table-wrapper"},p.default.createElement("div",{className:"table-search-bar member-search-bar-first",style:{marginBottom:10}},p.default.createElement("div",{className:"left-search-bar",style:{paddingBottom:10}},p.default.createElement(v.DateRangePicker,{value:this.state.date,placeholder:"选择日期范围",onChange:function(t){e.setState({date:t}),e.keyWord="",e.forceUpdate(),null!==t?(e.state.startDate=e.date(t[0]),e.state.endDate=e.date(t[1])):(e.state.startDate="",e.state.endDate=""),e.doQuery()}}),i?"":[p.default.createElement(v.Select,{value:this.productId,placeholder:"",onChange:function(t){return e.queryByProduct(t)}},p.default.createElement(v.Select.Option,{key:"",label:"全部产品",value:""}),n.length>0?n.map(function(e){return p.default.createElement(v.Select.Option,{key:e.pId,label:e.pName,value:e.pId})}):""),"user"===o?"":p.default.createElement(v.Select,{value:this.templateId,placeholder:"",onChange:function(t){return e.queryByTemplate(t)},disabled:r.length<=0||""===this.productId||this.noTemplate},p.default.createElement(v.Select.Option,{key:"",label:"证书名称",value:""}),r.length>0&&""!=this.productId&&!this.noTemplate?r.map(function(e){return p.default.createElement(v.Select.Option,{key:e.id,label:e.title,value:e.id})}):"")],i&&"org"===this.props.type?p.default.createElement(v.Select,{value:this.productId,placeholder:"",onChange:function(t){return e.queryByProduct(t)}},p.default.createElement(v.Select.Option,{key:"",label:"全部产品",value:""}),n.length>0?n.map(function(e){if("网页取证"!==e.pName&&"协议保全"!==e.pName&&"证据固定"!==e.pName&&"过程取证"!==e.pName)return p.default.createElement(v.Select.Option,{key:e.pId,label:e.pName,value:e.pId})}):""):"",p.default.createElement(v.Select,{value:this.fileType,placeholder:"",onChange:function(t){return e.queryByFileType(t)}},p.default.createElement(v.Select.Option,{key:"",label:"全部文件类型",value:""}),p.default.createElement(v.Select.Option,{key:"AUDIO",label:"音频",value:"AUDIO"}),p.default.createElement(v.Select.Option,{key:"VIDEO",label:"视频",value:"VIDEO"}),p.default.createElement(v.Select.Option,{key:"IMAGE",label:"图片",value:"IMAGE"}),p.default.createElement(v.Select.Option,{key:"DOC",label:"文档",value:"DOC"}),p.default.createElement(v.Select.Option,{key:"OTHER",label:"其他",value:"OTHER"})),p.default.createElement("div",{style:{marginLeft:"20px"},className:"label-box"},p.default.createElement(v.AutoComplete,{className:"my-autocomplete",placeholder:"选择标签",value:this.fileLabel,fetchSuggestions:this.querySearch.bind(this),onSelect:function(t){return e.handleSelect(t)},onChange:function(t){return e.changeTag(t)},onKeyDown:this.handleAutoKeydown,append:p.default.createElement(v.Button,{type:"primary",icon:"search",onClick:this.handleFileLabelQuery})})),B?p.default.createElement(v.Select,{value:this.operate,placeholder:"",onChange:function(t){return e.queryByOperate(t)}},p.default.createElement(v.Select.Option,{key:"",label:"全部操作人",value:""}),u.length>0?u.map(function(e){return p.default.createElement(v.Select.Option,{key:e.code,label:e.realName,value:e.code})}):""):""),p.default.createElement("div",{style:{paddingBottom:10}},p.default.createElement(v.Input,{placeholder:m.default.translate("attestation.search"),type:"text",value:this.keyWord,onChange:this.handleKeyWordChange,onKeyDown:this.handleKeydown,append:p.default.createElement(v.Button,{type:"primary",icon:"search",onClick:this.handleAttestationQuery})}))),p.default.createElement("div",null),p.default.createElement("div",{className:"common-table-wrapper"},p.default.createElement(v.Table,{style:{width:"100%"},columns:d,data:l.list,onCellClick:function(t,a){e.cellClick(t,a)}}),l.list&&l.list.length>0?p.default.createElement("div",{className:"common-page-wrapper",style:{marginTop:"20px"}},p.default.createElement(v.Pagination,{layout:"total, sizes, prev, pager, next, jumper",pageCount:l.totalPage,pageSizes:[20,50,100,500],pageSize:this.pageSize,currentPage:this.pageNo+1,onSizeChange:function(t){return e.queryByPageSize(t)},onCurrentChange:function(t){return e.queryByPageNo(t)},total:l.totalItems})):""),this.state.editModal?p.default.createElement(T.ModalContainer,{onClose:function(){return e.openEditModal(!1,"")}},p.default.createElement(T.ModalDialog,{onClose:function(){return e.openEditModal(!1,"")},width:560,className:"example-dialog",dismissOnBackgroundClick:!0},p.default.createElement("h1",{className:"copyright-model-title"},"修改标签"),p.default.createElement("div",{className:"black-white-model"},p.default.createElement("section",null,p.default.createElement("span",{style:{marginRight:30}},"标签:"),p.default.createElement(v.AutoComplete,{className:"my-autocomplete",placeholder:"请输入标签名选择或新增",value:this.label,fetchSuggestions:this.querySearch.bind(this),onSelect:function(t){return e.handleSelects(t)},onChange:function(t){return e.changeTags(t)},onBlur:function(){return e.changeWebTags()}})),p.default.createElement("div",{className:"button-group",style:{margin:"30px 0 0 62px"}},p.default.createElement("button",{className:"blue-button",type:"button",onClick:this.saveLabel},"保存"))))):"",this.state.showImage?p.default.createElement(L.default,{hideModel:this.closeImage,imageUrl:z}):"")}}]),t}(c.Component))||s;t.default=U}});