/**
 * Created by rayline on 2016/6/8.
 */

$(document).ready(function(){
    $("#buttonLogin").click(actionLogin);
    $("#buttonRegister").click(actionRegister);
    $("#buttonFindIDbyName").click(actionFindIDByName);
    $("#buttonUserInfo").click(actionUserInfo);
    $("#buttonFriendList").click(actionFriendList);
    $("#buttonSendMessage").click(actionSendMessage);
    $("#buttonFriendRequest").click(actionFriendRequest);
    $("#buttonFindIDbyMail").click(actionFindIDByMail);
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
                log(userid+"登陆成功，请自行获取用户信息");
            }else{
                log("登陆失败"+data.error);
            }
        }catch(err){log("请求失败"+err);}
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
            if(data.result=="success"){
                log("查询结果为:"+data.data.userid);
            }else{
                log("查询失败，可能没有");
            }
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionFindIDByMail(){
    var usermail = $("#textFindIDByMail").val();
    $.get("api/usermail/"+usermail,function(data){
        try{
            if(data.result=="success"){
                log("查询结果为:"+data.data.userid);
            }else{
                log("查询失败，可能没有");
            }
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionFriendRequest(){
    var target = $("#textFriendRequest").val();
    var description = $("#textFriendDescription").val();
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
    var target = $("#textMessageTarget").val();
    var description = $("#textMessageDescription").val();
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
                log("请求成功：好友列表为:"+JSON.stringify(data.data));
            }
        }catch(err){log("请求失败"+err);}
    },"json");
}

function actionUserInfo(){
    $.get("api/user/"+usernow,function(data){
         try{
            if(data.result=="success"){
                log("请求成功：用户信息为:"+data.JSON.stringify(data.data));
            }
         }catch(err){log("请求失败"+err);}
    },"json");
}

function actionModify(){
    var username = $("#textModifyName").val();
    var usermail = $("#textModifyMail").val();
    var userpassword = $("#textModifyPassword").val();
    var userdescription = $("#textModifyDescription").val();
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