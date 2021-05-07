function login() {
    var username = document.getElementById("username");
    var pass = document.getElementById("password");
    if (username.value == "" && pass.value  == ""){
        alert("請輸入使用者名稱及密碼");
    }
    else if (username.value == "") {
        alert("請輸入使用者名稱");
    } 
    else if (pass.value  == "") {
        alert("請輸入密碼");
    } 
    else{
        window.location.href="search.html";
    }
    } 

