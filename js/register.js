$(document).ready(function() {

    $("#submit_register_form_btn").click(function() {
        var Firstname = $("#firstname_register").val();
        var Lastname = $("#lastname_register").val();
        var Username = $("#username_register").val();
        var Email = $("#email_register").val();
        var Password = $("#password_register").val();
        var ConfirmPassword = $("#confirm_password_register").val();
        var Country = $("#country_register").val();

        $("#message_register_alert").remove();

        if($.trim(Firstname).length > 0 && $.trim(Lastname).length > 0 && $.trim(Username).length > 0 && $.trim(Email).length > 0 && $.trim(Password).length > 0 && $.trim(ConfirmPassword).length > 0 && $.trim(Country).length > 0) {
            if(Password == ConfirmPassword) {
                if(/^([a-zA-Z0-9_.-])+@(([a-zA-Z0-9-])+.)+([a-zA-Z0-9]{2,4})+$/.test(Email)) {
                    $.ajax({
                        url: "/register",
                        method: "POST",
                        data: {firstname:Firstname, lastname:Lastname, username:Username, email:Email, password:Password, country:Country},
                        beforeSend:function(){
                            $("#submit_register_form_btn").attr('disabled', true);
                            $("#submit_register_form_btn").val('Conectando...');
                        },
                        success:function(response) {
                            $('#submit_register_form_btn').removeAttr('disabled');
                            $("#submit_register_form_btn").val('Register');
                            if (response == "1") {
                                $("#messages").append("<div class='alert alert-success' role='alert' id='message_register_alert'> ¡Account created successfully, you'll be redirected in a few seconds! </div>");
                                redirectLogin();
                            } else if (response == "2") {
                                $("#messages").append("<div class='alert alert-info' role='alert' id='message_register_alert'> ¡It's seems like already exists an user with that Username, try another! </div>");
                            } else if (response == "3") {
                                $("#messages").append("<div class='alert alert-info' role='alert' id='message_register_alert'> ¡It's seems like already exists an user with that Email, try another! </div>");
                            }
                        }
                    });
                } else {
                    $("#messages").append("<div class='alert alert-warning' role='alert' id='message_register_alert'> ¡That is not a valid email! </div>");
                }
            } else {
                $("#messages").append("<div class='alert alert-warning' role='alert' id='message_register_alert'> ¡The passwords don't match! </div>");
            }
        } else {
            $("#messages").append("<div class='alert alert-danger' role='alert' id='message_register_alert'> ¡You have to fill all the fields below! </div>");
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