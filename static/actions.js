/**
 * Created by rayline on 2016/6/8.
 */

$(document).ready(function(){
    $("#buttonLogin").click(actionLogin);
    $("#buttonRegister").click(actionRegister);
    $("#buttonFindIDByName").click(actionFindIDByName);
    $("#buttonUserInfo").click(actionUserInfo);
    $("#buttonFriendList").click(actionFriendList);
    $("#buttonSendMessage").click(actionSendMessage);
    $("#buttonFriendRequest").click(actionFriendRequest);
    $("#buttonFindIDByMail").click(actionFindIDByMail);
    $("#buttonModify").click(actionModify);
    $("#buttonGetMessage").click(actionGetMessage);
});

var usernow;

function log(logMsg){
    $("#log").append("<div>"+logMsg+"</div>");
}

function actionLogin(){
    var userid = $("#textLoginID").val();
    var userpass = $("#textLoginPassword").val();
    $.post("api/user/"+userid+"/login",{password:userpass},function(data){
        try{
            if(data.result=="success"){
                usernow = userid;
            }
        }catch(err){log("请求失败"+err);}
        log(userid+"登陆成功，请自行获取用户信息");
    },"json");
}

function actionRegister(){
    var userpass = $("#textRegisterPassword").val();
    $.get("api/user/new",function(data){
        try{
            if(data.result=="success"){
                usernow = data.data.userid;
                log(usernow+"注册成功，已经是登录状态");
            }else{
                log("注册失败"+data.error);
            }
        }catch(err){log("请求失败"+err);}
        $.post("api/user/"+usernow,{password:userpass},function(data){
            try{
                if(data.result=="success"){
                    log("首次修改密码成功");
                }else{
                    log("首次修改密码失败"+data.error);
                }
            }catch(err){log("请求失败"+err);}
        },"json");
    },"json");
}

function actionFindIDByName(){
    var username = $("#textFindIDByName").val();
    $.get("api/username/"+username,function(data){
        try{
            log("查询结果为:"+data.userid);
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionFindIDByMail(){
    var usermail = $("#textFindIDByMail").val();
    $.get("api/usermail/"+username,function(data){
        try{
            log("查询结果为:"+data.userid);
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionFriendRequest(){
    var target = $("#textFriendRequest");
    var description = $("#textFriendDescription");
    $.post("api/user/"+target+"/friendrequest",{message:description},function(data){
        try{
            if(data.result=="success"){
                log("已发送");
            }else{
                log("请求失败"+data.error);
            }
        }catch(err){log("请求失败"+err);}

    },"json");
}

function actionSendMessage(){
    var target = $("#textMessageTarget");
    var description = $("#textMessageDescription");
    $.post("api/user/"+target+"/message",{message:description},function(data){
        try{
            if(data.result=="success"){
                log("已发送");
            }else{
                log("请求失败"+data.error);
            }
        }catch(err){log("请求失败"+err);}

    },"json");
}

function actionFriendList(){
    $.get("api/user/"+usernow+"/friendlist",function(data){
        try{
            if(data.result=="success"){
                log("请求成功：好友列表为:"+data.data.toJSON());
            }
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionUserInfo(){
    $.get("api/user/"+usernow,function(data){
         try{
            if(data.result=="success"){
                log("请求成功：用户信息为:"+data.data.toJSON());
            }
         }catch(err){log("请求失败"+err);}
    },"json");
}

function actionModify(){
    var username = $("#textModifyName");
    var usermail = $("#textModifyMail");
    var userpassword = $("#textModifyPassword");
    var userdescription = $("#textModifyDescription");
    $.post("api/user/"+usernow,{password:userpassword,mail:usermail,name:username,description:userdescription},function(data){
        try{
            if(data.result=="success"){
                log("修改成功");
            }else{
                log("修改失败"+data.error);
            }
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionGetMessage(){
    $.get("api/user/"+usernow,function(data){
    try{
        if(data.data.Type=="U"){
            log("来自"+data.data.Source+"的信息"+"</div><div>"+data.data.Content);
        }
        if(data.data.Type=="F"){
            log("来自"+data.data.Source+"的好友请求"+"</div><div>"+data.data.Content);
        }
        if(data.data.Type=="F"){
            log("系统消息"+"</div><div>"+data.data.Content);
        }
    }catch(err){log("并没有什么消息");}
},"json");
}