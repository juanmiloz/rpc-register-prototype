$(document).ready(function() {

    $("#submit_login_form_btn").click(function() {
        var Email = $("#email_login").val();
        var Password = $("#password_login").val();

        $("#message_login_alert").remove();

        if($.trim(Email).length > 0 && $.trim(Password).length > 0) {
            if(/^([a-zA-Z0-9_.-])+@(([a-zA-Z0-9-])+.)+([a-zA-Z0-9]{2,4})+$/.test(Email)) {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {email:Email, password:Password},
                    beforeSend:function(){
                        $("#submit_login_form_btn").attr('disabled', true);
                        $("#submit_login_form_btn").val('Conectando...');
                    },
                    success:function(response) {
                        $('#submit_login_form_btn').removeAttr('disabled');
                        $("#submit_login_form_btn").val('Log In');
                        if (response == "1") {
                            $("#messages_login").append("<div class='alert alert-success' role='alert' id='message_login_alert'> ¡You've logged successfully, you'll be redirected in a few seconds! </div>");
                            redirectLogin();
                        } else if (response == "2") {
                            $("#messages_login").append("<div class='alert alert-info' role='alert' id='message_login_alert'> ¡Do not exists an account with that email! </div>");
                        } else if (response == "3") {
                            $("#messages_login").append("<div class='alert alert-danger' role='alert' id='message_login_alert'> ¡Incorrect password! </div>");
                        }
                    }
                });
            } else {
                $("#messages_login").append("<div class='alert alert-warning' role='alert' id='message_register_alert'> ¡That is not a valid email! </div>");
            }
        } else {
            $("#messages_login").append("<div class='alert alert-danger' role='alert' id='message_register_alert'> ¡You have to fill all the fields below! </div>");
        }
    });

    async function redirectLogin() {
        await sleep(5000);
        window.location.replace("/userIndex");
    }

    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

});