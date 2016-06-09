/**
 * Created by rayline on 2016/6/8.
 */

var usernow;

function log(logMsg){
    $("#log").append("<div>"+logMsg+"</div>");
}

function actionLogin(){
    var userid = $("#textLoginID").val();
    var userpass = $("#textLoginPassword").val();
    $.get("api/user/"+userid+"/login",{password:userpass},function(data){
        try{
            if(data.result=="success"){
                usernow = userid;
            }
        }catch(err){log("请求失败");}
        log(userid+"登陆成功，请自行获取用户信息");
    },"json");
}

function actionRegiter(){
    var userpass = $("#textRegisterPassword").val();
    $.get("api/user/new",function(data){
        try{
            if(data.result=="success"){
                usernow = data.data.userid;
                log(userid+"注册成功，已经是登录状态");
            }else{
                log("注册失败"+data.error);
            }
        }catch(err){log("请求失败");}
        $.get("api/user/"+usernow,{password:userpass},function(data){
            try{
                if(data.result=="success"){
                    log("首次修改密码成功");
                }else{
                    log("首次修改密码失败"+data.error);
                }
            }catch(err){log("请求失败");}
            log(userid+"注册成功，已经是登录状态");
        },"json");
    },"json");
}

function actionFindIDByName(){
    var username = $("#textFindIDByName").val();
    $.get("api/username/"+username,function(data){
        try{
            log("查询结果为:"+data.userid);
        }catch(err){log("请求失败");}
    },"json");
}

function actionFindIDByMail(){
    var usermail = $("#textFindIDByMail").val();
    $.get("api/usermail/"+username,function(data){
        try{
            log("查询结果为:"+data.userid);
        }catch(err){log("请求失败");}
    },"json");
}

function actionFriendRequest(){
    var target = $("#textFriendRequest");
    var description = $("#textFriendDescription");
    $.get("api/user/"+target+"/friendrequest",{message:description},function(data){
        try{
            if(data.result=="success"){
                log("已发送");
            }else{
                log("请求失败"+data.error);
            }
        }catch(err){log("请求失败");}

    },"json");
}

function actionFriendRequest(){
    var target = $("#textFriendRequest");
    var description = $("#textFriendDescription");
    $.get("api/user/"+target+"/message",{message:description},function(data){
        try{
            if(data.result=="success"){
                log("已发送");
            }else{
                log("请求失败"+data.error);
            }
        }catch(err){log("请求失败");}

    },"json");
}

function actionFriendList(){
    $.get("api/user/"+usernow+"/friendlist",function(data){
        try{
            if(data.result=="success"){
                log("请求成功：好友列表为:"+data.data.toJSON());
            }
        }catch(err){log("请求失败");}
    },"json");
}

function actionUserInfo(){
    $.get("api/user/"+usernow,function(data){
         try{
            if(data.result=="success"){
                log("请求成功：用户信息为:"+data.data.toJSON());
            }
         }catch(err){log("请求失败");}
    },"json");
}

function actionModify(){
    var username = $("#textModifyName");
    var usermail = $("#textModifyMail");
    var userpassword = $("#textModifyPassword");
    var userdescription = $("#textModifyDescription");
    $.get("api/user/"+usernow,{password:userpassword,mail:usermail,name:username,description:userdescription},function(data){
        try{
            if(data.result=="success"){
                log("修改成功");
            }else{
                log("修改失败"+data.error);
            }
        }catch(err){log("请求失败");}
    },"json");
}

function actionGetMessage(){
    $.get("api/user/"+usernow,function(data){
    try{
        if(data.Type=="U"){
            log("来自"+data.Source+"的信息"+"</div><div>"+data.Content);
        }
        if(data.Type=="F"){
            log("来自"+data.Source+"的好友请求"+"</div><div>"+data.Content);
        }
        if(data.Type=="F"){
            log("系统消息"+"</div><div>"+data.Content);
        }
    }catch(err){log("请求失败");}
},"json");
}